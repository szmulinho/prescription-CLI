package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
	"github.com/szmulinho/prescription/pkg/model"
	"time"
)

var createPrescCmd = &cobra.Command{
	Use:   "create",
	Short: "create prescription",
	Long:  "this command will create new prescription",
	Run: func(cmd *cobra.Command, args []string) {

		Presc := model.Presc{
			PreId:      "",
			Drugs:      nil,
			Expiration: time.Time{},
		}
		fmt.Println("creating new %+v\n prescription", Presc)

		resp := endpoints.CreatePrescription
		fmt.Println("Prescription created with ID:", resp)
	},
}

func init() {
	rootCmd.AddCommand(createPrescCmd)
	createPrescCmd.Flags().StringP("PreID", "p", "", "specify Presc ID")
	createPrescCmd.Flags().StringP("Drugs", "d", "", "specify Presc drugs")
	createPrescCmd.Flags().dateStringP("Expiration", "e", "", "specify Presc expiration")
}
