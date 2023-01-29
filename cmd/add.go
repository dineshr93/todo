/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/dineshr93/todo/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a todo Format: Project - task - JIRA link",
	Long: `add a todo.
Format: Project - task - JIRA link

1. Name of project
2. Task
3. Jira link`,
	Run: func(cmd *cobra.Command, args []string) {

		dataFile := string(viper.ConfigFileUsed())
		// fmt.Println(dataFile)
		todos := &model.Todos{}
		if err := todos.Load(dataFile); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if len(args) > 0 {
			todos.Add(strings.Join(args, " "))
			err := todos.Store(dataFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.Print()
		} else {
			fmt.Println("Invalid args. space is important in around eifen. Format: Project - task - JIRA link")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
