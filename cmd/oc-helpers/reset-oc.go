package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"oc-helpers/pkg/env"
	oc_os "oc-helpers/pkg/os"
	"os"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset-oc",
	Short: "Reset all ownCloud data while preserving the config",
	Run: func(cmd *cobra.Command, args []string) {
		ocPath := env.Get("OC_PATH")
		configFilePath := fmt.Sprintf("%s/oc-config.json", oc_os.GetBasePath())
		configFile, _ := os.Create(configFilePath)
		defer func() {
			_ = os.Remove(configFilePath)
		}()

		err := oc_os.RunCmd("/usr/bin/php occ config:list --private", ocPath, configFile)
		if err != nil {
			log.Panicf(err.Error())
		}
		err = oc_os.RemoveContents(fmt.Sprintf("%s/data", ocPath))
		if err != nil {
			log.Panicf(err.Error())
		}
		err = oc_os.ReplaceFileContent(fmt.Sprintf("%s/config/config.php", ocPath), "'installed' => true,", "'installed' => false,")
		if err != nil {
			log.Panicf(err.Error())
		}
		err = oc_os.RunCmd("/usr/bin/php occ maintenance:install --admin-pass admin", ocPath, nil)
		if err != nil {
			log.Panicf(err.Error())
		}
		err = oc_os.RunCmd(fmt.Sprintf("/usr/bin/php occ config:import %s", configFilePath), ocPath, nil)
		if err != nil {
			log.Panicf(err.Error())
		}
	},
}
