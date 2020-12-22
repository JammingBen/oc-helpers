package main

import (
	"fmt"
	"log"
	"oc-helpers/pkg/env"
	"oc-helpers/pkg/os"
)

func main() {
	ocPath := env.Get("OC_PATH")
	err := os.RemoveContents(fmt.Sprintf("%s/data", ocPath))
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = os.ReplaceFileContent(fmt.Sprintf("%s/config/config.php", ocPath), "'installed' => true,", "'installed' => false,")
	if err != nil {
		log.Fatalf(err.Error())
	}
}
