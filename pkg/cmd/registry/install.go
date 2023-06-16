package registry

import (
	"github.com/apigear-io/cli/pkg/repos"

	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	// cmd represents the pkgInstall command
	var version string
	var cmd = &cobra.Command{
		Use:     "install [name]",
		Short:   "install template into cache",
		Long:    `install template from registry using a name`,
		Aliases: []string{"i"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			cmd.Printf("installing template from %s\n", name)
			info, err := repos.Registry.Get(name)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			version := ""
			switch len(args) {
			case 1:
				if info.Latest.Name == "" {
					cmd.PrintErrln("no version specified and no latest version available")
					return
				}
				version = info.Latest.Name
			case 2:
				version = args[1]
			default:
				cmd.PrintErrln("invalid number of arguments")
				return
			}
			fqn, err := repos.Cache.Install(info.Git, version)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Printf("installed template %s\n", fqn)
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", "latest", "template version to install")
	return cmd
}
