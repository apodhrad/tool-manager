/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available tools",
	Long: `List all tools which can be managed by this manager.
The output will also specify if and which version is installed.`,
	Run: listCmdRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolP("installed", "i", false, "List only installed tools")
	listCmd.Flags().StringP("name", "n", "", "List only tools with a given name")
}

func listCmdRun(cmd *cobra.Command, args []string) {
	table := NewTable("Name", "Version", "Installed")
	table.AddRow("example", "1.2.1", "installed")
	table.AddRow("example", "1.2.0", "")
	table.AddRow("example", "1.1.0", "")
	output := TableToString(table)
	fmt.Println(output)
}
