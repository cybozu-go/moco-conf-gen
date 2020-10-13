package cmd

import (
	"github.com/cybozu-go/well"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	serverIDBaseFlag = "server-id-base"
)

var serverIDBase uint32

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
			serverIDBase = viper.GetUint32(serverIDBaseFlag)
			return subMain()
		},
	}
)

func init() {
	// ordinal should be increased by 1000 as default because the case server-id is 0 is not suitable for the replication purpose
	rootCmd.Flags().Uint32(serverIDBaseFlag, 1000, "Base value of server-id.")
	err := viper.BindPFlags(rootCmd.Flags())
	if err != nil {
		panic(err)
	}
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
