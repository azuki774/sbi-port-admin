package main

import (
	"azuki774/sbiport-server/internal/factory"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dbRepo, err := factory.NewDBRepo(&factory.DBInfo{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer dbRepo.CloseDB()

		us, err := factory.NewUsecase(dbRepo)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		srv, err := factory.NewServer(&factory.ServerRunOption{}, us)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		srv.Start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
