package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/prescription/pkg/model"
	"net/http"
)

var deletePrescriptionCmd = &cobra.Command{
	Use:   "deletePrescription",
	Short: "Deletes a prescription",

	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest("DELETE", fmt.Sprintf("/prescriptions/%s", model.CreatePrescInput{PreId: ""}), nil)
		if err != nil {
			fmt.Println("Error creating delete request:", err)
			return
		}

		// Send the HTTP request to delete the prescription
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error deleting prescription:", err)
			return
		}
		defer resp.Body.Close()

		// Check if the HTTP response was successful
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error deleting prescription: server returned status code %d\n", resp.StatusCode)
			return
		}

		// Print a success message
		fmt.Printf("The prescription with ID %s has been deleted successfully\n", model.CreatePrescInput{PreId: ""})
	},
}

func init() {
	rootCmd.AddCommand(deletePrescriptionCmd)
}
