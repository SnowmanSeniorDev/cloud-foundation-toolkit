package cmd

import (
	"github.com/spf13/cobra"

	"github.com/GoogleCloudPlatform/cloud-foundation-toolkit/cli/launchpad"
)

func init() {
	rootCmd.AddCommand(launchpadCmd)
	launchpadCmd.AddCommand(launchpadGenerateCmd)
	launchpadCmd.AddCommand(launchpadBootstrapCmd)
}

var launchpadCmd = &cobra.Command{
	Use:     "launchpad",
	Aliases: []string{"lp"},
	Short:   "launchpad (lp)",
	Long: `Cloud Foundation Toolkit Launchpad
	bootstraps foundational GCP infrastructure by following the
	Cloud Foundation Ecosystem Convention. Taking YAML and generate opinionated
	infrastructure resources ready to be deployed in Infrastructure as Code style`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			cmd.HelpFunc()(cmd, args)
		}
	},
}

var launchpadGenerateCmd = &cobra.Command{
	Use:     "generate [YAML files]",
	Aliases: []string{"g", "gen"},
	Short:   "generate (g)",
	Long:    `Generate infrastructure foundation via defined YAML`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			cmd.HelpFunc()(cmd, args)
		} else {
			launchpad.NewGenerate(args)
		}
	},
}

var launchpadBootstrapCmd = &cobra.Command{
	Use:     "bootstrap",
	Aliases: []string{"b"},
	Short:   "bootstrap (g)",
	Long:    `Bootstrap initial seed project`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			cmd.HelpFunc()(cmd, args)
		} else {
			// TODO (@rjerrems) Bootstrap entry point
			print("Bootstrapping... (NYI)")
			launchpad.NewBootstrap()
		}
	},
}
