package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/argcomplete"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "az",
	Short:              "Azure Command-Line Interface",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		argcomplete.ActionArgcomplete("az"),
	)
}