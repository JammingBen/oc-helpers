package main

import (
	"fmt"
	"log"
	"oc-helpers/pkg/env"
	oc_os "oc-helpers/pkg/os"
	"os"
)

func main() {
	ocPath := env.Get("OC_PATH")
	configFilePath := fmt.Sprintf("%s/oc-config.json", oc_os.GetBasePath())
	configFile, _ := os.Create(configFilePath)

	err := oc_os.RunCmd("/usr/bin/php occ config:list --private", ocPath, configFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = oc_os.RemoveContents(fmt.Sprintf("%s/data", ocPath))
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = oc_os.ReplaceFileContent(fmt.Sprintf("%s/config/config.php", ocPath), "'installed' => true,", "'installed' => false,")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = oc_os.RunCmd("/usr/bin/php occ maintenance:install --admin-pass admin", ocPath, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = oc_os.RunCmd(fmt.Sprintf("/usr/bin/php occ config:import %s", configFilePath), ocPath, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	_ = os.Remove(configFilePath)
}
