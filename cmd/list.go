/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/apodhrad/tool-manager/utils"
	"github.com/rodaine/table"
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
	name, _ := cmd.Flags().GetString("name")
	installed, _ := cmd.Flags().GetBool("installed")
	tools := utils.GetTools(name, installed)
	tbl := getTable(tools)
	output := TableToString(tbl)
	fmt.Println(output)
}

func getTable(tools map[string]utils.Tool) table.Table {
	tbl := NewTable("Name", "Version", "Installed")
	for name, tool := range tools {
		for _, release := range tool.Releases {
			tbl.AddRow(name, release.Version, "")
		}
	}
	return tbl
}
