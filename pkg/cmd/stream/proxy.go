package stream

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"text/tabwriter"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/apigear-io/cli/pkg/foundation/logging"
	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/config"
)

// proxyRunOptions holds configuration for running the proxy server.
type proxyRunOptions struct {
	ConfigFile string
	Verbose    bool
	Trace      bool
}

// NewProxyCommand creates the proxy management command.
func NewProxyCommand() *cobra.Command {
	opts := &proxyRunOptions{}

	cmd := &cobra.Command{
		Use:   "proxy [config.yaml]",
		Short: "Run proxy server or manage proxies",
		Long: `Run the WebSocket proxy server, or use subcommands to manage proxy configuration.

When called without a subcommand, starts the proxy server using the config file.

Examples:
  # Start proxy server with default config
  apigear stream proxy

  # Start with custom config file
  apigear stream proxy config.yaml

  # Start with verbose logging
  apigear stream proxy --verbose

  # Manage proxies
  apigear stream proxy list
  apigear stream proxy create my-proxy --listen ws://localhost:5550/ws
  apigear stream proxy delete my-proxy`,
		RunE: func(cmd *cobra.Command, args []string) error {
			configFile := opts.ConfigFile
			if len(args) > 0 {
				configFile = args[0]
			}
			if configFile == "" {
				configFile = "stream.yaml"
			}

			cfg, created, err := config.LoadOrCreateConfig(configFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if created {
				logging.Info().Msgf("Created default config file: %s", configFile)
				logging.Info().Msg("Edit the config file and restart to configure proxies and clients")
			}

			if err := cfg.Validate(); err != nil {
				return fmt.Errorf("invalid config: %w", err)
			}

			if opts.Verbose {
				cfg.Verbose = true
			}
			if opts.Trace {
				cfg.Trace = true
			}

			services := stream.NewServices()
			defer services.Close()

			if len(cfg.Proxies) > 0 {
				logging.Info().Msgf("Loading %d proxies from config", len(cfg.Proxies))
				if err := services.ProxyManager.LoadFromConfig(cfg.Proxies); err != nil {
					log.Warn().Err(err).Msg("failed to load some proxies")
				}
				// Enable console traffic output for all proxies
				for name := range cfg.Proxies {
					if p, err := services.ProxyManager.GetProxy(name); err == nil {
						p.SetOutput(cmd.ErrOrStderr())
					}
				}
			}

			if len(cfg.Clients) > 0 {
				logging.Info().Msgf("Loading %d clients from config", len(cfg.Clients))
				if err := services.ClientManager.LoadFromConfig(cfg.Clients); err != nil {
					log.Warn().Err(err).Msg("failed to load some clients")
				}
			}

			proxies := services.ProxyManager.ListProxies()
			clients := services.ClientManager.ListClients()

			logging.Info().Msgf("Stream server started with %d proxies and %d clients",
				len(proxies), len(clients))

			if len(proxies) > 0 {
				logging.Info().Msg("Active proxies:")
				for _, p := range proxies {
					logging.Info().Msgf("  - %s: %s -> %s (%s, %s)",
						p.Name, p.Listen, p.Backend, p.Mode, p.Status)
				}
			}

			if len(clients) > 0 {
				logging.Info().Msg("Active clients:")
				for _, c := range clients {
					logging.Info().Msgf("  - %s: %s (%s)",
						c.Name, c.URL, c.Status)
				}
			}

			if len(proxies) == 0 && len(clients) == 0 {
				logging.Info().Msg("No proxies or clients configured. Edit the config file to add them.")
			}

			sigCh := make(chan os.Signal, 1)
			signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

			logging.Info().Msg("Press Ctrl+C to stop")

			<-sigCh
			logging.Info().Msg("Shutting down...")

			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.ConfigFile, "config", "c", "", "config file (default: stream.yaml)")
	cmd.Flags().BoolVarP(&opts.Verbose, "verbose", "v", false, "enable verbose logging")
	cmd.Flags().BoolVarP(&opts.Trace, "trace", "t", false, "enable trace logging to files")

	cmd.AddCommand(newProxyListCommand())
	cmd.AddCommand(newProxyCreateCommand())
	cmd.AddCommand(newProxyStartCommand())
	cmd.AddCommand(newProxyStopCommand())
	cmd.AddCommand(newProxyDeleteCommand())
	cmd.AddCommand(newProxyStatsCommand())

	return cmd
}

func newProxyListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all proxies",
		RunE: func(cmd *cobra.Command, args []string) error {
			services := stream.NewServices()
			defer services.Close()

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if err := services.ProxyManager.LoadFromConfig(cfg.Proxies); err != nil {
				return err
			}

			proxies := services.ProxyManager.ListProxies()

			if len(proxies) == 0 {
				fmt.Println("No proxies configured")
				return nil
			}

			// Print table
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Fprintln(w, "NAME\tLISTEN\tBACKEND\tMODE\tSTATUS\tCONNECTIONS")
			for _, p := range proxies {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\n",
					p.Name, p.Listen, p.Backend, p.Mode, p.Status, p.ActiveConnections)
			}
			w.Flush()

			return nil
		},
	}
}

type proxyCreateOptions struct {
	Listen  string
	Backend string
	Mode    string
}

func newProxyCreateCommand() *cobra.Command {
	opts := &proxyCreateOptions{}

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a new proxy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if opts.Listen == "" {
				return fmt.Errorf("--listen is required")
			}

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			// Add proxy to config
			if cfg.Proxies == nil {
				cfg.Proxies = make(map[string]config.ProxyConfig)
			}

			cfg.Proxies[name] = config.ProxyConfig{
				Listen:  opts.Listen,
				Backend: opts.Backend,
				Mode:    opts.Mode,
			}

			// Save config
			if err := config.SaveConfig("stream.yaml", cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			fmt.Printf("Proxy '%s' created successfully\n", name)
			fmt.Println("Run 'apigear stream' to start the proxy")

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.Listen, "listen", "", "listen address (e.g., ws://localhost:5550/ws)")
	cmd.Flags().StringVar(&opts.Backend, "backend", "", "backend URL (e.g., ws://localhost:5560/ws)")
	cmd.Flags().StringVar(&opts.Mode, "mode", "proxy", "proxy mode (proxy, echo, backend, inbound-only)")

	return cmd
}

func newProxyStartCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start <name>",
		Short: "Start a proxy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			services := stream.NewServices()
			defer services.Close()

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			proxyCfg, exists := cfg.Proxies[name]
			if !exists {
				return fmt.Errorf("proxy '%s' not found", name)
			}

			if err := services.ProxyManager.AddProxy(name, proxyCfg); err != nil {
				return err
			}

			if err := services.ProxyManager.StartProxy(name); err != nil {
				return err
			}

			fmt.Printf("Proxy '%s' started\n", name)
			return nil
		},
	}
}

func newProxyStopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stop <name>",
		Short: "Stop a proxy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			services := stream.NewServices()
			defer services.Close()

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if err := services.ProxyManager.LoadFromConfig(cfg.Proxies); err != nil {
				return err
			}

			if err := services.ProxyManager.StopProxy(name); err != nil {
				return err
			}

			fmt.Printf("Proxy '%s' stopped\n", name)
			return nil
		},
	}
}

func newProxyDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete a proxy",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if _, exists := cfg.Proxies[name]; !exists {
				return fmt.Errorf("proxy '%s' not found", name)
			}

			delete(cfg.Proxies, name)

			// Save config
			if err := config.SaveConfig("stream.yaml", cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			fmt.Printf("Proxy '%s' deleted\n", name)
			return nil
		},
	}
}

func newProxyStatsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "stats [name]",
		Short: "Show proxy statistics",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			services := stream.NewServices()
			defer services.Close()

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if err := services.ProxyManager.LoadFromConfig(cfg.Proxies); err != nil {
				return err
			}

			if len(args) == 0 {
				// Show all stats
				stats := services.Stats.AllProxyStats()

				if len(stats) == 0 {
					fmt.Println("No proxies running")
					return nil
				}

				w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
				fmt.Fprintln(w, "NAME\tMSG_RECV\tMSG_SENT\tBYTES_RECV\tBYTES_SENT\tCONNS\tUPTIME")
				for _, s := range stats {
					fmt.Fprintf(w, "%s\t%d\t%d\t%d\t%d\t%d\t%ds\n",
						s.Name, s.MessagesReceived, s.MessagesSent,
						s.BytesReceived, s.BytesSent, s.ActiveConnections, s.Uptime)
				}
				w.Flush()
			} else {
				// Show specific proxy stats
				name := args[0]
				proxy, err := services.ProxyManager.GetProxy(name)
				if err != nil {
					return err
				}

				info := proxy.Info()
				fmt.Printf("Proxy: %s\n", info.Name)
				fmt.Printf("Listen: %s\n", info.Listen)
				fmt.Printf("Backend: %s\n", info.Backend)
				fmt.Printf("Mode: %s\n", info.Mode)
				fmt.Printf("Status: %s\n", info.Status)
				fmt.Printf("Messages Received: %d\n", info.MessagesReceived)
				fmt.Printf("Messages Sent: %d\n", info.MessagesSent)
				fmt.Printf("Bytes Received: %d\n", info.BytesReceived)
				fmt.Printf("Bytes Sent: %d\n", info.BytesSent)
				fmt.Printf("Active Connections: %d\n", info.ActiveConnections)
				fmt.Printf("Uptime: %ds\n", info.Uptime)
			}

			return nil
		},
	}
}
