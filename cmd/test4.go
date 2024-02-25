package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/hiswest/test4"
)

// test4Cmd represents the test4 command
var test4Cmd = &cobra.Command{
	Use:   "test4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		test4.RunTest4()
	},
}

func init() {
	rootCmd.AddCommand(test4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// test4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// test4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
