package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ochelper",
		Short: "A small helper project for ownCloud including handy commands and services to make developing easier",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
