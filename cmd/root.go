/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

func BuildRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mydiary",
		Short: "Manage your diary",
		Long: `mydiary is a CLI tool for Golang that manages your diary.
This tool generates diary templates, shows diary of specified date.`,
	}

	cmd.AddCommand(BuildInitCmd())
	cmd.AddCommand(BuildShowCmd())
	return cmd
}
