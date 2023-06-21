package cmd

import (
	"errors"

	"github.com/apigear-io/cli/pkg/cmd/cache"
	"github.com/apigear-io/cli/pkg/cmd/cfg"
	"github.com/apigear-io/cli/pkg/cmd/gen"
	"github.com/apigear-io/cli/pkg/cmd/mon"
	"github.com/apigear-io/cli/pkg/cmd/prj"
	"github.com/apigear-io/cli/pkg/cmd/registry"
	"github.com/apigear-io/cli/pkg/cmd/sim"
	"github.com/apigear-io/cli/pkg/cmd/spec"
	"github.com/apigear-io/cli/pkg/cmd/x"

	"github.com/spf13/cobra"
)

var ErrSilent = errors.New("SilentErr")

func NewRootCommand() *cobra.Command {
	// cmd represents the base command when called without any subcommands
	cmd := &cobra.Command{
		Use:   "apigear",
		Short: "apigear creates instrumented SDKs from an API description",
		Long:  `ApiGear allows you to describe interfaces and generate instrumented SDKs out of the descriptions.`,
	}
	cmd.SilenceErrors = false
	cmd.SilenceUsage = false
	cmd.AddCommand(gen.NewRootCommand())
	cmd.AddCommand(mon.NewRootCommand())
	cmd.AddCommand(cfg.NewRootCommand())
	cmd.AddCommand(registry.NewRootCommand())
	cmd.AddCommand(cache.NewRootCommand())
	cmd.AddCommand(sim.NewRootCommand())
	cmd.AddCommand(spec.NewRootCommand())
	cmd.AddCommand(prj.NewRootCommand())
	cmd.AddCommand(x.NewRootCommand())
	cmd.AddCommand(NewUpdateCommand())
	cmd.AddCommand(NewVersionCommand())
	return cmd
}
