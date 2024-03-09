package command

import (
	"github.com/enuesaa/loadii/pkg/repository"
	"github.com/enuesaa/loadii/pkg/usecase"
	"github.com/urfave/cli/v2"
)

func NewRunCommand(repos repository.Repos, watchpath *string) *cli.Command {
	cmd := cli.Command{
		Name:  "run",
		Usage: "run app. Currently, this command supports golang app.",
		Before: func(c *cli.Context) error {
			path := c.Args().Get(0)
			if path == "" {
				return cli.ShowAppHelp(c)
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			path := c.Args().Get(0)

			return usecase.RunAppWatch(repos, *watchpath, path)
		},
	}

	return &cmd
}
