package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set Git build environment key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_config_setCmd).Standalone()
	git_configCmd.AddCommand(git_config_setCmd)
}
