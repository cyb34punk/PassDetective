/*
Copyright © 2023 Yunus AYDIN <aydinnyunus@gmail.com>
*/
package cmd

import (
	"github.com/aydinnyunus/PassDetective/cmd/internal/extract"
	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

func getExtractCmd() *cobra.Command {
	extractCmd := &cobra.Command{
		Use:   "extract",
		Short: "Extract passwords from shell history",
		Long: `The "extract" command allows you to automatically extract passwords from shell history.
By analyzing the history of shell commands, this tool can identify and extract passwords that were used during
previous commands and display them for further inspection or use.
`,
		Run: extractFn,
	}
	extractCmd.Flags().BoolP("zsh", "z", false, "Check passwords on ZSH")
	extractCmd.Flags().BoolP("bash", "b", false, "Check passwords on BASH")
	return extractCmd
}

func extractFn(cmd *cobra.Command, args []string) {
	zsh, _ := cmd.Flags().GetBool("zsh")
	bash, _ := cmd.Flags().GetBool("bash")

	kind := extract.All
	if zsh {
		kind = extract.Zsh
	} else if bash {
		kind = extract.Bash
	}
	printStart()
	out, err := extract.Create(kind).Process()
	if err != nil {
		color.Red("error while processing history file: %v", err)
		return
	}
	for kind, values := range out {
		color.Green("Total provider detected for %s: %d", kind, len(values))
		for provider, detections := range values {
			color.Green("Detected %d entries for %s", len(detections), provider)
			for _, v := range detections {
				color.Red("Line number: %d", v.LineNum)
				color.Red("Line text: %s", v.Text)
			}
		}
	}
	printEnd()
}
