/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/qgp9/pssh-go/parser"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install pssh.conf to ~/.ssh/config",
	Long:  `install pssh.conf to ~/.ssh/config`,
	Run: func(cmd *cobra.Command, args []string) {
		var pconfig = parser.ParsePConfigFromFile("example/pssh.conf")
		pconfig.WriteSshConfig("output.txt")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
