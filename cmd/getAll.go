/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/prescription/pkg/api/endpoints"
)

var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:`,

	Run: func(cmd *cobra.Command, args []string) {
		resp := endpoints.GetAllPrescriptions
		fmt.Println("Here are your prescriptions", resp)
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
	getAllCmd.Flags()
}
