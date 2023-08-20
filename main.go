package main

import "github.com/spf13/cobra"

var cmd = &cobra.Command{
	Use: "ssh-server-setup",
	Run: func(cmd *cobra.Command, args []string) {
		ensureSshdService()
		// make sure sshd config doesn't allow password authentication
		// make sure

	},
}

func ensureSshdService() {
	// if sshd service exists, return
	// if sshd service does not exist, create it
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
