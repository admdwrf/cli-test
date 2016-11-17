package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	statusTimeout = 200
)

// createinstanceCmd represents the createinstance command
var createinstanceCmd = &cobra.Command{
	Use:   "createinstance",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("createinstance called")
	},
}

func init() {
	RootCmd.AddCommand(createinstanceCmd)
	createinstanceCmd.Flags().BoolP("project", "", false, "OVH Cloud Project Name, required=True")
	createinstanceCmd.Flags().BoolP("name", "", false, "OVH Cloud VM Instance Name, default=test0")
	createinstanceCmd.Flags().BoolP("region", "", false, "OVH Cloud Region, default=SBG1")
	createinstanceCmd.Flags().BoolP("image", "", false, "OVH Cloud OS Image, default=Ubuntu 16.04")
	createinstanceCmd.Flags().BoolP("flavor", "", false, "OVH Cloud Machine Type, default=vps-ssd-1")
	createinstanceCmd.Flags().BoolP("ssh-key", "", false, "OVH Cloud SSH Key Name, required=True")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createinstanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createinstanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
