package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/preAppCli/internal/structs"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
	"github.com/szmulinho/prescription/pkg/model"
	"time"
)

var createCmd = &cobra.Command{
	Use:   "create1",
	Short: "Create a presc",

	Run: func(cmd *cobra.Command, args []string) {

		t, err := time.Parse(time.RFC3339, structs.Expiration)
		if err != nil {
			fmt.Println("Error parsing expiration time:", err)
			return
		}

		preID, _ := cmd.Flags().GetString("preID")
		preDrugs, _ := cmd.Flags().GetStringArray("preDrugs")
		pre := model.CreatePrescInput{
			PreId:      preID,
			Drugs:      preDrugs,
			Expiration: t,
		}
		fmt.Printf("Creating presc %+v\n", pre)

		resp := endpoints.CreatePrescription
		fmt.Println("Task created with ID:", resp)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("preID", "p", "", "ID")
	createCmd.Flags().StringP("preDrugs", "d", "", "Drugs")
	createCmd.Flags().StringP("expiration", "e", "", "Expiration")
}
