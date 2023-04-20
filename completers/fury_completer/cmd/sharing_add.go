package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sharing_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a collaborator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharing_addCmd).Standalone()
	sharing_addCmd.Flags().String("role", "", "Collaborator role")
	sharingCmd.AddCommand(sharing_addCmd)
}
