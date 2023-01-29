/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/dineshr93/todo/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pendingCmd represents the pending command
var pendingCmd = &cobra.Command{
	Use:   "pending",
	Short: "Mark a task as pending(supports space separated multiple params)",
	Long: `Mark a task as pending.  Can give multiple id with space separated:

tdv pending taskID1 taskID2 `,
	Run: func(cmd *cobra.Command, args []string) {

		dataFile := string(viper.ConfigFileUsed())
		todos := &model.Todos{}
		if err := todos.Load(dataFile); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if len(args) > 0 {
			err := todos.PendingSA(args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err = todos.Store(dataFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.Print()
		} else {
			fmt.Println("Invalid args. Should be space separated id1 id2")
			os.Exit(1)

		}

	},
}

func init() {
	rootCmd.AddCommand(pendingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pendingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pendingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
