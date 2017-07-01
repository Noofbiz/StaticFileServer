package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func readConfig() (path, port string) {
	c := make(map[string]interface{})
	var err error
	var f *os.File
	if f, err = os.Open("/configuration/conf.json"); err != nil {
		log.Fatalf("Unable to open configuration file. Error: %v", err.Error())
	}
	defer f.Close()

	var buf []byte
	if buf, err = ioutil.ReadAll(f); err != nil {
		log.Fatalf("Failure reading configuration file. Error: %v", err.Error())
	}
	if err = json.Unmarshal(buf, &c); err != nil {
		log.Fatalf("Failure unmarshalling config.json. Error: %v", err.Error())
	}

	if c["path"] == "default" {
		var pwd string
		if pwd, err = os.Getwd(); err != nil {
			log.Fatalf("Unable to get working directory. Error: %v", err.Error())
		}
		path = filepath.Join(pwd, "static")
	} else {
		path = c["path"].(string)
	}

	port = ":" + c["port"].(string)
	return
}
