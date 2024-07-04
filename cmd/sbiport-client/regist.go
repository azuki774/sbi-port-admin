package main

import (
	"azuki774/sbiport-server/internal/factory"
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RegistInfo factory.RegistInfo
var HTTPClientInfo factory.HTTPClientInfo

// loadCmd represents the load command
var registCmd = &cobra.Command{
	Use:   "regist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		hc := factory.NewHTTPClient(&HTTPClientInfo)
		us, err := factory.NewUsecaseClient(hc)
		if err != nil {
			fmt.Printf("failed to create usecaseClient struct: %v\n", err)
			os.Exit(1)
		}

		var fileNameList []string
		RegistInfo.TargetDir = args[0]
		fInfo, _ := os.Stat(RegistInfo.TargetDir)
		if fInfo.IsDir() {
			files, _ := os.ReadDir(RegistInfo.TargetDir)
			for _, f := range files {
				fileNameList = append(fileNameList, RegistInfo.TargetDir+f.Name())
			}
		} else {
			// file
			fileNameList = append(fileNameList, RegistInfo.TargetDir)
		}

		for _, v := range fileNameList {
			ctx := context.Background()
			us.RegistJob(ctx, v)
		}

	},
}

func init() {
	rootCmd.AddCommand(registCmd)

	registCmd.Flags().StringVar(&HTTPClientInfo.Scheme, "scheme", "", "http or https")
	registCmd.Flags().StringVar(&HTTPClientInfo.Host, "host", "", "server host")
	registCmd.Flags().StringVar(&HTTPClientInfo.Port, "port", "", "server port")
}
