package cfg

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/cfg"
	"github.com/spf13/cobra"
)

func jsonIdent(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(data)
}
func NewEnvCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "env",
		Short: "Env prints apigear environment variables",
		Long:  `Env prints apigear environment variables`,
		Run: func(cmd *cobra.Command, args []string) {
			settings := cfg.AllSettings()
			for key, value := range settings {
				name := fmt.Sprintf("APIGEAR_%s", strings.ToUpper(key))
				valueStr := jsonIdent(value)
				cmd.Printf("%s=%v\n", name, valueStr)
			}
		},
	}
	return cmd
}
