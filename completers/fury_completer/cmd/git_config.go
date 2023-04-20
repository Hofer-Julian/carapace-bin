package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var git_configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Git build",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_configCmd).Standalone()
	gitCmd.AddCommand(git_configCmd)
}
