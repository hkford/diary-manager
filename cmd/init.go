package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func BuildInitCmd() *cobra.Command {
	var year int64
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generate diary template",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(year)
		},
	}
	cmd.Flags().Int64Var(&year, "year", 0, "Year")
	return cmd
}
