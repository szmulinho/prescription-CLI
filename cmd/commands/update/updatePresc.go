package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/preAppCli/internal/model"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var UpdateDrugCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a drug with a specific ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		expirationTime, _ := time.Parse("06-01-02-15-04-05", model.Expiration)
		drug := model.CreatePrescInput{
			PreId:      model.PrescriptionID,
			Drugs:      model.Drugs,
			Expiration: expirationTime,
		}

		url := fmt.Sprintf("http://localhost:8080/prescs/%s", model.PrescriptionID)
		reqBody, err := json.Marshal(drug)
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to update drug: %s", resp.Status)
		}

		fmt.Println("Prescription updated successfully!")
		return nil
	},
}

func init() {
	UpdateDrugCmd.Flags().StringVarP(&model.PrescriptionID, "id", "i", "", "Prescription ID to update (required)")
	UpdateDrugCmd.MarkFlagRequired("presc-id")
	UpdateDrugCmd.Flags().StringSliceVarP(&model.Drugs, "drugs", "d", []string{}, "New drugs in the prescription")
	UpdateDrugCmd.Flags().StringVarP(&model.Expiration, "expiration", "e", "", "New expiration of the prescription")
}
