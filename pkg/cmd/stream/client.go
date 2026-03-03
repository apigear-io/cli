package stream

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/apigear-io/cli/pkg/stream"
	"github.com/apigear-io/cli/pkg/stream/config"
)

// NewClientCommand creates the client management command.
func NewClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Manage ObjectLink clients",
		Long:  `Manage ObjectLink clients (list, create, connect, disconnect, delete).`,
	}

	cmd.AddCommand(newClientListCommand())
	cmd.AddCommand(newClientCreateCommand())
	cmd.AddCommand(newClientConnectCommand())
	cmd.AddCommand(newClientDisconnectCommand())
	cmd.AddCommand(newClientDeleteCommand())
	cmd.AddCommand(newClientStatusCommand())

	return cmd
}

func newClientListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all clients",
		RunE: func(cmd *cobra.Command, args []string) error {
			services := stream.NewServices()
			defer services.Close()

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if err := services.ClientManager.LoadFromConfig(cfg.Clients); err != nil {
				return err
			}

			clients := services.ClientManager.ListClients()

			if len(clients) == 0 {
				fmt.Println("No clients configured")
				return nil
			}

			// Print table
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
			fmt.Fprintln(w, "NAME\tURL\tSTATUS\tINTERFACES")
			for _, c := range clients {
				interfaces := strings.Join(c.Interfaces, ", ")
				if len(interfaces) > 40 {
					interfaces = interfaces[:37] + "..."
				}
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
					c.Name, c.URL, c.Status, interfaces)
			}
			w.Flush()

			return nil
		},
	}
}

type clientCreateOptions struct {
	URL           string
	Interfaces    []string
	Enabled       bool
	AutoReconnect bool
}

func newClientCreateCommand() *cobra.Command {
	opts := &clientCreateOptions{}

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a new client",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			if opts.URL == "" {
				return fmt.Errorf("--url is required")
			}

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			// Add client to config
			if cfg.Clients == nil {
				cfg.Clients = make(map[string]config.ClientConfig)
			}

			cfg.Clients[name] = config.ClientConfig{
				URL:           opts.URL,
				Interfaces:    opts.Interfaces,
				Enabled:       opts.Enabled,
				AutoReconnect: opts.AutoReconnect,
			}

			// Save config
			if err := config.SaveConfig("stream.yaml", cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			fmt.Printf("Client '%s' created successfully\n", name)
			fmt.Println("Run 'apigear stream' to start the client")

			return nil
		},
	}

	cmd.Flags().StringVar(&opts.URL, "url", "", "WebSocket URL (e.g., ws://localhost:5560/ws)")
	cmd.Flags().StringSliceVar(&opts.Interfaces, "interfaces", []string{}, "ObjectLink interfaces to link")
	cmd.Flags().BoolVar(&opts.Enabled, "enabled", true, "enable client on startup")
	cmd.Flags().BoolVar(&opts.AutoReconnect, "auto-reconnect", true, "automatically reconnect on disconnect")

	return cmd
}

func newClientConnectCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "connect <name>",
		Short: "Connect a client",
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

			clientCfg, exists := cfg.Clients[name]
			if !exists {
				return fmt.Errorf("client '%s' not found", name)
			}

			if err := services.ClientManager.AddClient(name, clientCfg); err != nil {
				return err
			}

			if err := services.ClientManager.ConnectClient(name); err != nil {
				return err
			}

			fmt.Printf("Client '%s' connected\n", name)
			return nil
		},
	}
}

func newClientDisconnectCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "disconnect <name>",
		Short: "Disconnect a client",
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

			if err := services.ClientManager.LoadFromConfig(cfg.Clients); err != nil {
				return err
			}

			if err := services.ClientManager.DisconnectClient(name); err != nil {
				return err
			}

			fmt.Printf("Client '%s' disconnected\n", name)
			return nil
		},
	}
}

func newClientDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "delete <name>",
		Short: "Delete a client",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			// Load config
			cfg, _, err := config.LoadOrCreateConfig("stream.yaml")
			if err != nil {
				return fmt.Errorf("failed to load config: %w", err)
			}

			if _, exists := cfg.Clients[name]; !exists {
				return fmt.Errorf("client '%s' not found", name)
			}

			delete(cfg.Clients, name)

			// Save config
			if err := config.SaveConfig("stream.yaml", cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			fmt.Printf("Client '%s' deleted\n", name)
			return nil
		},
	}
}

func newClientStatusCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "status <name>",
		Short: "Show client status",
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

			if err := services.ClientManager.LoadFromConfig(cfg.Clients); err != nil {
				return err
			}

			client, err := services.ClientManager.GetClient(name)
			if err != nil {
				return err
			}

			info := client.Info()
			fmt.Printf("Client: %s\n", info.Name)
			fmt.Printf("URL: %s\n", info.URL)
			fmt.Printf("Status: %s\n", info.Status)
			fmt.Printf("Interfaces: %s\n", strings.Join(info.Interfaces, ", "))
			fmt.Printf("Auto-reconnect: %v\n", info.AutoReconnect)
			fmt.Printf("Enabled: %v\n", info.Enabled)
			if info.LastError != "" {
				fmt.Printf("Last Error: %s\n", info.LastError)
			}

			return nil
		},
	}
}
