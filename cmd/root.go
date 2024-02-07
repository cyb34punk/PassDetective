/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "PassDetective",
	Short: "Extract passwords from shell history change descriptions",
	Long: `The "extract" command allows you to automatically extract passwords from shell history change descriptions.
By analyzing the history of shell commands, this tool can identify and extract passwords that were used during
previous commands and display them for further inspection or use.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(passDetectiveAsciiArt)
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().BoolP("help", "h", false, "Help message for PassDetective")
	rootCmd.AddCommand(getExtractCmd())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
