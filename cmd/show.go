package cmd

import (
	"github.com/spf13/cobra"
)

func BuildShowCmd() *cobra.Command {
	var date string
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show diary",
	}
	cmd.Flags().StringVar(&date, "date", "", "Date")
	return cmd
}
