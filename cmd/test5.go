package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/hiswest/test5"
)

// test5Cmd represents the test5 command
var test5Cmd = &cobra.Command{
	Use:   "test5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test5.RunTest5()
	},
}

func init() {
	rootCmd.AddCommand(test5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
