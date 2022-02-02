package cmd

import (
	"github.com/spf13/cobra"
)

func BuildShowCmd() *cobra.Command {
	var date int64
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show diary",
	}
	cmd.Flags().Int64Var(&date, "date", 0, "Date")
	return cmd
}
