package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func init() {
	hubRootCmd.AddCommand(
		hubStatusCmd,
		hubWorkerRootCmd,
		hubACLRootCmd,
		hubOrderRootCmd,
		hubTasksRootCmd,
		hubDeviceRootCmd,
	)
}

var hubRootCmd = &cobra.Command{
	Use:   "hub",
	Short: "Hub management",
}

var hubStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show hub status",
	Run: func(cmd *cobra.Command, _ []string) {
		hub, err := NewHubInteractor(nodeAddressFlag, timeoutFlag)
		if err != nil {
			showError(cmd, "Cannot connect to Node", err)
			os.Exit(1)
		}

		status, err := hub.Status()
		if err != nil {
			showError(cmd, "Cannot get hub status", err)
			os.Exit(1)
		}

		printHubStatus(cmd, status)
	},
}
