package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnf",
	Short: "Install, update and remove packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("6", "6", false, "resolve to IPv6 addresses only")
	rootCmd.Flags().BoolS("4", "4", false, "resolve to IPv4 addresses only")
	rootCmd.Flags().Bool("allowerasing", false, "allow erasing of installed packages to resolve dependencies")
	rootCmd.Flags().Bool("assumeno", false, "automatically answer no for all questions")
	rootCmd.Flags().BoolP("assumeyes", "y", false, "automatically answer yes for all questions")
	rootCmd.Flags().BoolP("best", "b", false, "try the best available package versions in transactions.")
	rootCmd.Flags().Bool("bugfix", false, "Include bugfix relevant packages, in updates")
	rootCmd.Flags().BoolP("cacheonly", "C", false, "run entirely from system cache, don't update cache")
	rootCmd.Flags().String("color", "", "control whether color is used")
	rootCmd.Flags().String("comment", "", "add a comment to transaction")
	rootCmd.Flags().Bool("debugsolver", false, "dumps detailed solving results into files")
	rootCmd.Flags().Bool("disable", false, "disable repos with config-manager command")
	rootCmd.Flags().String("disablerepo", "", "Temporarily disable active repositories for the purpose of the current dnf command.")
	rootCmd.Flags().Bool("downloadonly", false, "only download packages")
	rootCmd.Flags().Bool("enable", false, "enable repos with config-manager command")
	rootCmd.Flags().String("enablerepo", "", "Temporarily enable repositories for the purpose of the current dnf command.")
	rootCmd.Flags().Bool("enhancement", false, "Include enhancement relevant packages, in updates")
	rootCmd.Flags().String("forcearch", "", "Force the use of an architecture")
	rootCmd.Flags().String("installroot", "", "set install root")
	rootCmd.Flags().Bool("newpackage", false, "Include newpackage relevant packages, in updates")
	rootCmd.Flags().Bool("noautoremove", false, "disable removal of dependencies that are no longer used")
	rootCmd.Flags().Bool("nobest", false, "do not limit the transaction to the best candidate")
	rootCmd.Flags().Bool("nodocs", false, "do not install documentations")
	rootCmd.Flags().Bool("nogpgcheck", false, "disable gpg signature checking (if RPM policy allows)")
	rootCmd.Flags().Bool("noplugins", false, "disable all plugins")
	rootCmd.Flags().Bool("obsoletes", false, "enables dnf's obsoletes processing logic for upgrade or display capabilities")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet operation")
	rootCmd.Flags().Bool("refresh", false, "set metadata as expired before running the command")
	rootCmd.Flags().Bool("security", false, "Include security relevant packages, in updates")
	rootCmd.Flags().String("setopt", "", "set arbitrary config and repo options")
	rootCmd.Flags().Bool("showduplicates", false, "show duplicates, in repos, in list/search commands")
	rootCmd.Flags().Bool("skip-broken", false, "resolve depsolve problems by skipping packages")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose operation")
	rootCmd.Flags().Bool("version", false, "show DNF version and exit")
}
