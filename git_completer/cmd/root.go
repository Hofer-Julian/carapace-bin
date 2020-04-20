package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
  carapace.Gen(rootCmd).Standalone()
  rootCmd.Flags().Bool("bare", false, "use $PWD as repository")
  rootCmd.Flags().StringP("c", "c", "", "pass configuration parameter to command")
  rootCmd.Flags().StringP("C", "C", "", "run as if git was started in given path")
  rootCmd.Flags().String("exec-path", "", "path containing core git-programs")
  rootCmd.Flags().String("git-dir", "", "path to repository")
  rootCmd.Flags().Bool("help", false, "display help message")
  rootCmd.Flags().Bool("html-path", false, "display path to HTML documentation and exit")
  rootCmd.Flags().Bool("info-path", false, "print the path where the info files are installed and exit")
  rootCmd.Flags().Bool("literal-pathspecs", false, "treat pathspecs literally, rather than as glob patterns")
  rootCmd.Flags().Bool("man-path", false, "print the manpath for the man pages for this version of Git and exit")
  rootCmd.Flags().String("namespace", "", "set the Git namespace")
  rootCmd.Flags().BoolP("no-pager", "P", false, "don't pipe git output into a pager")
  rootCmd.Flags().Bool("no-replace-objects", false, "do not use replacement refs to replace git objects")
  rootCmd.Flags().BoolP("paginate", "p", false, "pipe output into a pager")
  rootCmd.Flags().Bool("version", false, "display version information")
  rootCmd.Flags().String("work-tree", "", "path to working tree")

  carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
      "c": carapace.ActionMessage("TODO"), // TODO ActionMultiparts with callback for config values (git help --config)
      "C": carapace.ActionDirectories(),
      "exec-path": carapace.ActionDirectories(),
      "git-dir": carapace.ActionDirectories(),
      "work-tree": carapace.ActionDirectories(),
  })
}
