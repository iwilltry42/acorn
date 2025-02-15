package cli

import (
	apiv1 "github.com/acorn-io/runtime/pkg/apis/api.acorn.io/v1"
	cli "github.com/acorn-io/runtime/pkg/cli/builder"
	"github.com/acorn-io/runtime/pkg/cli/builder/table"
	"github.com/acorn-io/runtime/pkg/tables"
	"github.com/spf13/cobra"
	"k8s.io/utils/strings/slices"
)

func NewApp(c CommandContext) *cobra.Command {
	return cli.Command(&App{client: c.ClientFactory}, cobra.Command{
		Use:     "app [flags] [APP_NAME...]",
		Aliases: []string{"apps", "a", "ps"},
		Example: `
acorn app`,
		SilenceUsage:      true,
		Short:             "List or get apps",
		ValidArgsFunction: newCompletion(c.ClientFactory, appsCompletion).complete,
	})
}

type App struct {
	All    bool   `usage:"Include stopped apps" short:"a"`
	Quiet  bool   `usage:"Output only names" short:"q"`
	Output string `usage:"Output format (json, yaml, {{gotemplate}})" short:"o"`
	client ClientFactory
}

func (a *App) Run(cmd *cobra.Command, args []string) error {
	c, err := a.client.CreateDefault()
	if err != nil {
		return err
	}

	out := table.NewWriter(tables.App, a.Quiet, a.Output)

	if len(args) == 1 {
		app, err := c.AppGet(cmd.Context(), args[0])
		if err != nil {
			return err
		}
		out.Write(app)
		return out.Err()
	}

	apps, err := c.AppList(cmd.Context())
	if err != nil {
		return err
	}

	for _, app := range apps {
		if (app.Status.AppStatus.Stopped || inactive(app)) && !a.All {
			continue
		}
		if len(args) > 0 {
			if slices.Contains(args, app.Name) {
				out.Write(&app)
			}
		} else {
			out.Write(&app)
		}
	}

	return out.Err()
}

func inactive(app apiv1.App) bool {
	return app.Status.Ready &&
		app.Status.Columns.Healthy == "0" &&
		app.Status.Columns.UpToDate == "0" &&
		app.Status.Columns.Message == "OK"
}
