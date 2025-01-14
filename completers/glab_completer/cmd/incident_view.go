package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var incident_viewCmd = &cobra.Command{
	Use:     "view <id>",
	Short:   "Display the title, body, and other information about an incident.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(incident_viewCmd).Standalone()

	incident_viewCmd.Flags().BoolP("comments", "c", false, "Show incident comments and activities")
	incident_viewCmd.Flags().StringP("page", "p", "", "Page number")
	incident_viewCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page")
	incident_viewCmd.Flags().BoolP("system-logs", "s", false, "Show system activities / logs")
	incident_viewCmd.Flags().BoolP("web", "w", false, "Open incident in a browser. Uses default browser or browser specified in BROWSER variable")
	incidentCmd.AddCommand(incident_viewCmd)

	// TODO positional completion
}
