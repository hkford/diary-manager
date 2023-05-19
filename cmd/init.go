package cmd

import (
	"fmt"
	"mydiary/pkg/util"
	"mydiary/pkg/workspace"
	"strconv"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func BuildInitCmd() *cobra.Command {
	var year int64
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generate diary template.",
		Long: `
Generate diary template.
Diary template is a collection of text files like 2020/202001.txt`,
		Example: `
Generate diary template of 2020.
$ mydiary init --year 2020`,
		Run: func(c *cobra.Command, args []string) {
			var AppFs = afero.NewOsFs()
			isLeap := util.IsLeapYear(year)
			ws := workspace.Workspace{
				DiaryDir: strconv.FormatInt(year, 10),
				IsLeap:   isLeap,
				Fs:       &afero.Afero{Fs: AppFs},
			}
			err := ws.Create()
			if err != nil {
				panic("Failed to create directory")
			}
			err = ws.WriteYearTemplates()
			if err != nil {
				err = ws.Delete()
				if err != nil {
					m := fmt.Sprintf("Failed to write yearly template and delete directory: %s\n", ws.DiaryDir)
					panic(m)
				}
				panic("Failed to write yearly template")
			}
		},
	}
	cmd.Flags().Int64Var(&year, "year", 0, "year of diary")
	_ = cmd.MarkFlagRequired("year")
	return cmd
}
