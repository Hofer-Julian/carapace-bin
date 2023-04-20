package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove Git build environment key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_config_unsetCmd).Standalone()
	git_configCmd.AddCommand(git_config_unsetCmd)
}
