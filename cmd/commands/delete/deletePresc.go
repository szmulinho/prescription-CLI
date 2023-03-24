package delete

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var DeletePrescCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete prescription by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		url := fmt.Sprintf("http://localhost:8080/prescs/%s", id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Prescription with ID %s deleted successfully\n", id)
		} else {
			fmt.Printf("Error deleteing prescription with ID %s\n", id, resp.Status)
		}
	},
}
