package cmd

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
)

type Projects []string

var listprojects Projects

func listProjects() {

	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	client.Get("/cloud/project", &listprojects)
	fmt.Printf("Your projects are %s\n", listprojects)

}

// listprojectsCmd represents the listprojects command
var listprojectsCmd = &cobra.Command{
	Use:   "listprojects",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		listProjects()
	},
}

func init() {
	RootCmd.AddCommand(listprojectsCmd)

}
