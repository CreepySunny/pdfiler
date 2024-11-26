/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert a markdown file to pdf.",
	Long: `
pdfiler convert [flags] <markdown_file>

Examples:

Convert a Markdown File to PDF:

	pdfiler convert example.md

This generates 'example.pdf' in the current directory.

Specify Output Filename:

	pdfiler convert example.md -o custom_name.pdf

Set Output Directory

	pdfiler convert example.md -d ./output`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("convert called")
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
