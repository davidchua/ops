package cmd

import "github.com/spf13/cobra"

var (

	// systemdCmd is a helper command that helps create/manipulate and manage
	// systemd services
	systemdCmd = &cobra.Command{
		Use:   "systemd",
		Short: "Helper toolkit for systemd related commands",
	}
)

func init() {
	rootCmd.AddCommand(systemdCmd)
}
