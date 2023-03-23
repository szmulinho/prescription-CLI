package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
)

var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "get a list of all prescriptions",

	Run: func(cmd *cobra.Command, args []string) {

		resp := endpoints.GetAllPrescriptions
		fmt.Println(resp)
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
	getAllCmd.Flags()
}
