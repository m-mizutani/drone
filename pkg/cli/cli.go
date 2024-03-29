package cli

import (
	"github.com/m-mizutani/drone/pkg/cli/config"
	"github.com/m-mizutani/drone/pkg/domain/types"
	"github.com/m-mizutani/drone/pkg/utils"
	"github.com/urfave/cli/v2"
)

func Run(args []string) error {
	var (
		logger config.Logger

		logCloser func()
	)

	app := cli.App{
		Name:    "drone",
		Flags:   mergeFlags([]cli.Flag{}, &logger),
		Version: types.AppVersion,
		Commands: []*cli.Command{
			subImport(),
		},
		Before: func(ctx *cli.Context) error {
			f, err := logger.Configure()
			if err != nil {
				return err
			}
			logCloser = f
			return nil
		},
		After: func(ctx *cli.Context) error {
			if logCloser != nil {
				logCloser()
			}
			return nil
		},
	}

	if err := app.Run(args); err != nil {
		utils.HandleError("failed to run drone", err)
		return err
	}

	return nil
}
