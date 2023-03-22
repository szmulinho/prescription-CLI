package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/preAppCli/internal/structs"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
)

var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "get a list of all prescriptions",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(structs.Prescs)

		resp := endpoints.GetAllPrescriptions
		fmt.Println("There are your prescriptions", resp)
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
	getAllCmd.Flags()
}
