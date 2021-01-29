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
	rootCmd.AddCommand(importAppConfigCmd)
}

var importAppConfigCmd = &cobra.Command{
	Use:   "import-app-config",
	Short: "Import the config for an OC app",
	Long:  `Import app config file which are located in the directory app_config.
			Can be called with one or multiple app names as param(s). When called without
			param, the import is run for all config files.`,
	Run: func(cmd *cobra.Command, args []string) {
		basepath := oc_os.GetBasePath()
		ocPath := env.Get("OC_PATH")
		apps := args
		if len(args) == 0 {
			apps = oc_os.ReadFileNames(fmt.Sprintf("%s/app_config", basepath), true)
		}

		for _, app := range apps {
			configFile := fmt.Sprintf("%s/app_config/%s.json", basepath, app)
			if _, err := os.Stat(configFile); os.IsNotExist(err) {
				log.Println(fmt.Sprintf("config file missing for app %s", app))
				continue
			}

			err := oc_os.RunCmd(fmt.Sprintf("/usr/bin/php occ config:import %s", configFile), ocPath, nil)
			if err != nil {
				log.Println(fmt.Sprintf("import failed for app %s with error: %s", app, err.Error()))
				continue
			}
		}
	},
}
