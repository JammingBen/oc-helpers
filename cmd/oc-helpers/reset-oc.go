package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"oc-helpers/pkg/env"
	"oc-helpers/pkg/os"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset-oc",
	Short: "Reset your ownCloud installation (all data will be lost)",
	Run: func(cmd *cobra.Command, args []string) {
		ocPath := env.Get("OC_PATH")
		err := os.RemoveContents(fmt.Sprintf("%s/data", ocPath))
		if err != nil {
			log.Panicf(err.Error())
		}
		err = os.ReplaceFileContent(fmt.Sprintf("%s/config/config.php", ocPath), "'installed' => true,", "'installed' => false,")
		if err != nil {
			log.Panicf(err.Error())
		}
	},
}
