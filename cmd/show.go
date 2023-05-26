package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hkford/diary-manager/pkg/show"
	"github.com/hkford/diary-manager/pkg/util"
	"github.com/hkford/diary-manager/pkg/workspace"

	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func BuildShowCmd() *cobra.Command {
	var date int64
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show diary of specified date.",
		Long:  "Show diary of specified date.",
		Example: `
Show diary of 2020/01/01
$ mydiary show --date 20200101`,
		Run: func(c *cobra.Command, args []string) {
			var AppFs = afero.NewOsFs()
			date, err := show.ValidateInput(date)
			if err != nil {
				log.Fatalf("Invalid argument: %v\n", err)
			}
			isLeap := util.IsLeapYear(date.Y)
			ws := workspace.Workspace{
				DiaryDir: strconv.FormatInt(date.Y, 10),
				IsLeap:   isLeap,
				Fs:       &afero.Afero{Fs: AppFs},
			}
			isFileExist, err := show.IsDiaryFileExists(ws, date)
			if err != nil {
				log.Fatalf("Diary file does not exists: %v\n", err)
			}
			if isFileExist {
				diary, err := show.GetDiary(ws, date)
				if err != nil {
					log.Fatalf("Failed to get diary: %v\n", err)
				}
				fmt.Println(diary)
			} else {
				fmt.Println("Diary of specified date does not exist.")
			}
		},
	}
	cmd.Flags().Int64Var(&date, "date", 0, "date to show diary")
	_ = cmd.MarkFlagRequired("date")
	return cmd
}
