package cmd

import (
	"fmt"
	"time"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
	"github.com/szmulinho/prescription/pkg/model"
	"github.com/spf13/cobra"
)

func cmd() {
	var preID, expiration string
	var drugs []string

	var createPrescCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new prescription",
		Run: func(cmd *cobra.Command, args []string) {
			t, err := time.Parse(time.RFC3339, expiration)
			if err != nil {
				fmt.Println("Error parsing expiration time:", err)
				return
			}

			pre := model.CreatePrescInput{
				PreId:      preID,
				Drugs:      drugs,
				Expiration: t,
			}

			fmt.Println("Created prescription:", pre)

			resp := endpoints.CreatePrescription
			fmt.Println("Prescription created with ID:", resp)
		},
	}

	createPrescCmd.Flags().StringVar(&preID, "id", "", "ID of the prescription (required)")
	createPrescCmd.MarkFlagRequired("id")

	createPrescCmd.Flags().StringSliceVar(&drugs, "drugs", []string{}, "List of drugs (separated by commas)")

	createPrescCmd.Flags().StringVar(&expiration, "expiration", "", "Expiration time (RFC3339 format) (required)")
	createPrescCmd.MarkFlagRequired("expiration")

	var rootCmd = &cobra.Command{Use: "prescription-cli"}
	rootCmd.AddCommand(createPrescCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
