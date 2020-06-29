package cmd

import (
	"github.com/cybozu-go/well"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "moco-conf-gen",
		Short: "Configuration generator MySQL instances managed by MOCO",
		Long:  `Configuration generator MySQL instances managed by MOCO.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// without this, each subcommand's RunE would display usage text.
			cmd.SilenceUsage = true

			err := well.LogConfig{}.Apply()
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return subMain()
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
