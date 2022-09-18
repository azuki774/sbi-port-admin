package main

import (
	"azuki774/sbiport-server/internal/factory"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var DBInfo factory.DBInfo

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
		dbRepo, err := factory.NewDBRepo(&DBInfo)
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

	startCmd.Flags().StringVar(&DBInfo.Host, "db-host", "", "DB Host")
	startCmd.Flags().StringVar(&DBInfo.Port, "db-port", "", "DB Port")
	startCmd.Flags().StringVar(&DBInfo.DBName, "db-name", "", "DB Name")
	startCmd.Flags().StringVar(&DBInfo.UserName, "db-user", "", "DB User")
	startCmd.Flags().StringVar(&DBInfo.UserPass, "db-pass", "", "DB Pass")
}
