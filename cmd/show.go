package cmd

import (
	"fmt"
	"mydiary/pkg/show"
	"mydiary/pkg/util"
	"mydiary/pkg/workspace"
	"strconv"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func BuildShowCmd() *cobra.Command {
	var date int64
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show diary",
		Run: func(c *cobra.Command, args []string) {
			var AppFs = afero.NewOsFs()
			date, err := show.ValidateInput(date)
			if err != nil {
				fmt.Printf("Invalid argument: %v", err)
			}
			isLeap := util.IsLeapYear(date.Y)
			ws := workspace.Workspace{
				DiaryDir: strconv.FormatInt(date.Y, 10),
				IsLeap:   isLeap,
				Fs:       &afero.Afero{Fs: AppFs},
			}
			isFileExist := show.IsDiaryFileExists(ws, date)
			if isFileExist {
				diary, err := show.GetDiary(ws, date)
				if err != nil {
					fmt.Printf("Failed to get diary: %v", err)
				}
				fmt.Println(diary)
			} else {
				fmt.Println("Diary of specified date does not exist.")
			}
		},
	}
	cmd.Flags().Int64Var(&date, "date", 0, "Date")
	return cmd
}
