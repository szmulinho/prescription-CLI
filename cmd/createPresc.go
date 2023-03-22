package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/preAppCli/internal/structs"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
	"github.com/szmulinho/prescription/pkg/model"
	"time"
)

var createPrescCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new prescription",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := time.Parse(time.RFC3339, structs.Expiration)
		if err != nil {
			fmt.Println("Error parsing expiration time:", err)
			return
		}

		pre := model.CreatePrescInput{
			PreId:      structs.PreID,
			Drugs:      structs.Drugs,
			Expiration: t,
		}

		fmt.Println("Created prescription:", pre)

		resp := endpoints.CreatePrescription
		fmt.Println("Prescription created with ID:", resp)
	},
}

func init() {

	rootCmd.AddCommand(createPrescCmd)
	createPrescCmd.Flags().StringVar(&structs.PreID, "id", "", "ID of the prescription (required)")
	createPrescCmd.Flags().StringSliceVar(&structs.Drugs, "drugs", []string{}, "List of drugs (separated by commas)")
	createPrescCmd.Flags().StringVar(&structs.Expiration, "expiration", "", "Expiration time (RFC3339 format) (required)")
	var rootCmd = &cobra.Command{Use: "prescription-cli"}
	rootCmd.AddCommand(createPrescCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
