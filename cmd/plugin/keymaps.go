/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package plugin

import (
	"fmt"

	"github.com/spf13/cobra"
)

// keymapsCmd represents the keymaps command
var keymapsCmd = &cobra.Command{
	Use:   "keymaps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("keymaps called")
	},
}

func init() {
	PluginCmd.AddCommand(keymapsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// keymapsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// keymapsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
