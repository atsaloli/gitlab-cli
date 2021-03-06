package create

import (
	"github.com/makkes/gitlab-cli/api"
	createproj "github.com/makkes/gitlab-cli/cmd/create/project"
	"github.com/makkes/gitlab-cli/cmd/create/vars"
	"github.com/makkes/gitlab-cli/config"
	"github.com/spf13/cobra"
)

func NewCommand(client api.Client, cfg config.Config) *cobra.Command {
	var project *string
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a resource such as a project or a variable",
	}

	project = cmd.PersistentFlags().StringP("project", "p", "", "If present, the project scope for this CLI request")

	cmd.AddCommand(createproj.NewCommand(client))
	cmd.AddCommand(vars.NewCommand(client, project))

	return cmd
}
