package tpl

import "github.com/spf13/cobra"

// apigear template list -- list all templates from the registry
// apigear template install  -- install a template from the registry into the cache
// apigear template cache  -- list all templates in the cache
// apigear template remove -- remove a template from the cache
// apigear template clean  -- clean the cache
// apigear template update -- update a template from the registry
// apigear template info  -- get information about a template from the registry
// apigear template create -- create a new custom template ((--lang language))
// apigear template import -- import a template from a local directory or git url
func NewRootCommand() *cobra.Command {
	// cmd represents the tpl command
	cmd := &cobra.Command{
		Use:     "template",
		Aliases: []string{"tpl", "t"},
		Short:   "manage template registry and cache and create custom templates",
	}
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewInstallCommand())
	cmd.AddCommand(NewUpdateCommand())
	cmd.AddCommand(NewInfoCommand())
	// cache commands
	cmd.AddCommand(NewCacheCommand())
	cmd.AddCommand(NewRemoveCommand())
	cmd.AddCommand(NewCleanCommand())
	cmd.AddCommand(NewImportCommand())
	// custom template commands
	cmd.AddCommand(NewCreateCommand())
	cmd.AddCommand(NewLintCommand())
	cmd.AddCommand(NewPublishCommand())
	return cmd
}
