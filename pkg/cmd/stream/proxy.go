package stream

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/config"
)

// NewProxyCommand creates the proxy management command.
func NewProxyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Manage WebSocket proxies",
		Long:  `Manage WebSocket proxies (list, create, start, stop, delete).`,
	}

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
