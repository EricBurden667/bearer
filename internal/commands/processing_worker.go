package commands

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bearer/bearer/internal/commands/debugprofile"
	"github.com/bearer/bearer/internal/commands/process/orchestrator/worker"
	"github.com/bearer/bearer/internal/flag"
	"github.com/bearer/bearer/internal/util/output"
)

func NewProcessingWorkerCommand() *cobra.Command {
	flags := &flag.Flags{
		ProcessFlagGroup: flag.NewProcessGroup(),
		GeneralFlagGroup: flag.NewGeneralFlagGroup(),
		ScanFlagGroup:    flag.NewScanFlagGroup(),
	}

	cmd := &cobra.Command{
		Use:   "processing-worker [flags] PATH",
		Short: "start scan processing server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := flags.Bind(cmd); err != nil {
				return fmt.Errorf("flag bind error: %w", err)
			}

			output.Setup(cmd, output.SetupRequest{
				LogLevel:  viper.GetString(flag.LogLevelFlag.ConfigName),
				Quiet:     viper.GetBool(flag.QuietFlag.ConfigName),
				ProcessID: viper.GetString(flag.WorkerIDFlag.ConfigName),
			})

			if viper.GetBool(flag.DebugProfileFlag.ConfigName) {
				debugprofile.Start()
			}

			processOptions, err := flags.ProcessFlagGroup.ToOptions()
			if err != nil {
				return fmt.Errorf("options binding error: %w", err)
			}

			log.Debug().Msgf("running scan worker on port `%s`", processOptions.Port)

			err = worker.Start(processOptions.ParentProcessID, processOptions.Port)
			return err
		},
		Hidden:        true,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmd.SetFlagErrorFunc(flagErrorFunc)
	flags.AddFlags(cmd)

	return cmd
}