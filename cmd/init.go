package cmd

import (
	"mydiary/pkg/initialize"
	"mydiary/pkg/workspace"
	"strconv"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func BuildInitCmd() *cobra.Command {
	var year int64
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generate diary template",
		Run: func(c *cobra.Command, args []string) {
			var AppFs = afero.NewOsFs()
			isLeap := initialize.IsLeapYear(year)
			ws := workspace.Workspace{
				DiaryDir: strconv.FormatInt(year, 10),
				IsLeap:   isLeap,
				Fs:       &afero.Afero{Fs: AppFs},
			}
			ws.Create()
			initialize.WriteYearTemplates(ws)
		},
	}
	cmd.Flags().Int64Var(&year, "year", 0, "Year")
	return cmd
}
