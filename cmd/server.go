package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// serverCmd represents the command that starts the host agetn server
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Server configuration",
	// Run the server in the main thread
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		log.Infof("Start server in %v:%v", host, port)
	},
}

// init adds this command to the root cmd
func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().String("host", "127.0.0.1", "Address to attach this server")
	serverCmd.Flags().Int("port", 8080, "Port to listen for new connections")
}
