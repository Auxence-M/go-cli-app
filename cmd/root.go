/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"github.com/spf13/cobra"
)

// Store value of flags
var dataFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cli-app",
	Short: "go-cli-app is a todo application made in go",
	Long: `go-cli-app is a todo application made in go. 
I dicided to to this to learn more about the go programming language basics. 
It is designed to be as simple as possible to help you accomplish your goals `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	home, err := os.UserHomeDir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile")
	}

	// $HOME/.next.json
	dataFilePath := home+string(os.PathSeparator)+".todo.json"

	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", dataFilePath, "data file to store todos")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


