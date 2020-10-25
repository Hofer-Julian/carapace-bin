package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ktlint",
	Short: "An anti-bikeshedding Kotlin linter with built-in formatter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("android", "a", false, "Turn on Android Kotlin Style Guide compatibility")
	rootCmd.Flags().Bool("apply-to-idea", false, "Update Intellij IDEA")
	rootCmd.Flags().Bool("apply-to-idea-project", false, "Update Intellij IDEA project")
	rootCmd.Flags().Bool("color", false, "Make output colorful")
	rootCmd.Flags().String("color-name", "", "Customize the output color")
	rootCmd.Flags().Bool("debug", false, "Turn on debug output")
	rootCmd.Flags().String("disabled_rules", "", "Comma-separated list of rules to globally disable")
	rootCmd.Flags().String("editorconfig", "", "Path to .editorconfig")
	rootCmd.Flags().Bool("experimental", false, "Enabled experimental rules (ktlint-ruleset-experimental)")
	rootCmd.Flags().BoolP("format", "F", false, "Fix any deviations from the code style")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	rootCmd.Flags().Bool("install-git-pre-commit-hook", false, "Install git hook to automatically check files")
	rootCmd.Flags().Bool("install-git-pre-push-hook", false, "Install git hook to automatically check files")
	rootCmd.Flags().String("limit", "", "Maximum number of errors to show (default: show all)")
	rootCmd.Flags().Bool("print-ast", false, "Print AST")
	rootCmd.Flags().Bool("relative", false, "Print files relative to the working directory")
	rootCmd.Flags().String("reporter", "", "A reporter to use")
	rootCmd.Flags().StringP("ruleset", "R", "", "A path to a JAR file containing additional ruleset(s)")
	rootCmd.Flags().Bool("stdin", false, "Read file from stdin")
	rootCmd.Flags().BoolP("verbose", "v", false, "Show error codes")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"editorconfig": carapace.ActionFiles(".editorconfig"),
		"color-name":   carapace.ActionValues("WHITE", "LIGHT_GRAY", "GRAY", "DARK_GRAY", "BLACK", "RED", "PINK", "ORANGE", "YELLOW", "GREEN", "MAGENTA", "CYAN", "BLUE"),
		"disabled_rules": carapace.ActionMultiParts(",", func(args, parts []string) carapace.Action {
			return ActionRules().Invoke(args).Filter(parts).ToA()
		}),
		"reporter": carapace.ActionValues("plain", "plain?group_by_file", "json", "checkstyle", "html"),
		"ruleset":  carapace.ActionFiles(".jar"),
	})
}

func ActionRules() carapace.Action {
	return carapace.ActionValues(
		"chain-wrapping",
		"colon-spacing",
		"comma-spacing",
		"comment-spacing",
		"curly-spacing",
		"dot-spacing",
		"filename",
		"final-newline",
		"import-ordering",
		"indent",
		"keyword-spacing",
		"max-line-length",
		"modifier-order",
		"no-blank-line-before-rbrace",
		"no-consecutive-blank-lines",
		"no-empty-class-body",
		"no-line-break-after-else",
		"no-line-break-before-assignment",
		"no-multi-spaces",
		"no-semi",
		"no-trailing-spaces",
		"no-unit-return",
		"no-unused-imports",
		"no-wildcard-imports",
		"op-spacing",
		"parameter-list-wrapping",
		"paren-spacing",
		"range-spacing",
		"string-template",
	)
}