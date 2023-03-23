package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"github.com/szmulinho/preAppCli/internal/model"
	"time"
)

var addPrescCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new prescription",
	Long:  "create a new prescription to the drug database at localhost:8080/presc",
	Run: func(cmd *cobra.Command, args []string) {
		// read in the drug details
		prescId, _ := cmd.Flags().GetString("presc-id")
		drugs, _ := cmd.Flags().GetStringSlice("drugs")
		expiration, _ := cmd.Flags().GetString("expiration")
		expirationTime, _ := time.Parse("06-01-02-15-04-05", expiration)

		// create a new drug struct
		newPresc := model.CreatePrescInput{
			PreId:      prescId,
			Drugs:      drugs,
			Expiration: expirationTime,
		}

		// convert the drug struct to JSON
		jsonStr, _ := json.Marshal(newPresc)

		// send a POST request to the server to add the drug
		url := "http://localhost:8080/presc"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// print the response from the server
		fmt.Println("Response Status:", resp.Status)
	},
}

func init() {
	rootCmd.AddCommand(addPrescCmd)
	addPrescCmd.Flags().String("presc-id", "", "Prescription ID")
	addPrescCmd.Flags().StringSlice("drugs", []string{}, "Drugs in prescription ")
	addPrescCmd.Flags().String("expiration", "", "Prescription's expiration")
}
