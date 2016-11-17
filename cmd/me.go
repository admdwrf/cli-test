package cmd

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
)

type PartialMe struct {
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
}

var me PartialMe

func meDetail() {

	client, err := ovh.NewEndpointClient("ovh-eu")
	if err != nil {
		fmt.Printf("Error: %q\n", err)
		return
	}

	client.Get("/me", &me)
	fmt.Printf("Welcome %s!\n", me.Firstname)
	fmt.Printf("Your email is %s\n", me.Email)

}

// meCmd represents the me command
var meCmd = &cobra.Command{
	Use:   "me",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("me called")
		meDetail()
	},
}

func init() {
	RootCmd.AddCommand(meCmd)

}
