package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package or packages on your system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.PersistentFlags().BoolS("6", "6", false, "resolve to IPv6 addresses only")
	installCmd.PersistentFlags().BoolS("4", "4", false, "resolve to IPv4 addresses only")
	installCmd.PersistentFlags().Bool("allowerasing", false, "a erasing of installed packages to resolve dependencies")
	installCmd.PersistentFlags().Bool("assumeno", false, "automatically answer no for all questions")
	installCmd.PersistentFlags().BoolP("assumeyes", "y", false, "automatically answer yes for all questions")
	installCmd.PersistentFlags().BoolP("best", "b", false, "try the best available package versions in transactions.")
	installCmd.PersistentFlags().Bool("bugfix", false, "Include bugfix relevant packages, in updates")
	installCmd.PersistentFlags().BoolP("cacheonly", "C", false, "run entirely from system cache, don't update cache")
	installCmd.PersistentFlags().String("color", "", "control whether color is used")
	installCmd.PersistentFlags().String("comment", "", "add a comment to transaction")
	installCmd.PersistentFlags().Bool("debugsolver", false, "dumps detailed solving results into files")
	installCmd.PersistentFlags().Bool("disable", false, "disable repos with config-manager command")
	installCmd.PersistentFlags().String("disablerepo", "", "Temporarily disable active repositories for the purpose of the current dnf command. Accepts an id, a")
	installCmd.PersistentFlags().Bool("downloadonly", false, "only download packages")
	installCmd.PersistentFlags().Bool("enable", false, "enable repos with config-manager command")
	installCmd.PersistentFlags().String("enablerepo", "", "Temporarily enable repositories for the purpose of the current dnf command. ")
	installCmd.PersistentFlags().Bool("enhancement", false, "Include enhancement relevant packages, in updates")
	installCmd.PersistentFlags().String("forcearch", "", "Force the use of an architecture")
	installCmd.PersistentFlags().String("installroot", "", "set install root")
	installCmd.PersistentFlags().Bool("newpackage", false, "Include newpackage relevant packages, in updates")
	installCmd.PersistentFlags().Bool("noautoremove", false, "disable removal of dependencies that are no longer used")
	installCmd.PersistentFlags().Bool("nobest", false, "do not limit the transaction to the best candidate")
	installCmd.PersistentFlags().Bool("nodocs", false, "do not install documentations")
	installCmd.PersistentFlags().Bool("nogpgcheck", false, "disable gpg signature checking (if RPM policy allows)")
	installCmd.PersistentFlags().Bool("noplugins", false, "disable all plugins")
	installCmd.PersistentFlags().Bool("obsoletes", false, "enables dnf's obsoletes processing logic")
	installCmd.PersistentFlags().BoolP("quiet", "q", false, "quiet operation")
	installCmd.PersistentFlags().Bool("refresh", false, "set metadata as expired before running the command")
	installCmd.PersistentFlags().Bool("security", false, "Include security relevant packages, in updates")
	installCmd.PersistentFlags().String("setopt", "", "set arbitrary config and repo options")
	installCmd.PersistentFlags().Bool("showduplicates", false, "show duplicates, in repos, in list/search commands")
	installCmd.PersistentFlags().Bool("skip-broken", false, "resolve depsolve problems by skipping packages")
	installCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose operation")
	installCmd.PersistentFlags().Bool("version", false, "show DNF version and exit")
	rootCmd.AddCommand(installCmd)
}
