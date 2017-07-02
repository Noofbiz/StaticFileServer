package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//ReadConfig reads the confiuration file
func ReadConfig() (path, port string) {
	c := make(map[string]interface{})
	var err error
	var pwd string
	if pwd, err = os.Getwd(); err != nil {
		log.Fatalf("Unable to get working directory. Error: %v", err.Error())
	}

	var f *os.File
	if f, err = os.Open(filepath.Join(pwd, "configuration", "conf.json")); err != nil {
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
		path = filepath.Join(pwd, "static")
	} else {
		path = c["path"].(string)
	}

	port = c["port"].(string)
	return
}

//UpdateConfig updates the configuration file
func UpdateConfig(path, port string) {
	c := make(map[string]interface{})
	c["path"] = path
	c["port"] = strings.TrimLeft(port, ":")
	var err error
	var pwd string
	if pwd, err = os.Getwd(); err != nil {
		log.Fatalf("Unable to get working directory. Error: %v", err.Error())
	}

	if err = os.Remove(filepath.Join(pwd, "configuration", "conf.json")); err != nil {
		log.Fatalf("Unable to remove conf file. Error: %v", err.Error())
	}

	var f *os.File
	if f, err = os.Create(filepath.Join(pwd, "configuration", "conf.json")); err != nil {
		UpdateConfig(filepath.Join(pwd, "static"), "8080")
		log.Fatalf("Unable to create new configuration file. Error: %v", err.Error())
	}
	defer f.Close()

	var buf []byte
	if buf, err = json.Marshal(c); err != nil {
		UpdateConfig(filepath.Join(pwd, "static"), "8080")
		log.Fatalf("Unable to marshal config. Error: %v", err.Error())
	}
	if _, err = f.Write(buf); err != nil {
		UpdateConfig(filepath.Join(pwd, "static"), "8080")
		log.Fatalf("Unable to write to config. Error: %v", err.Error())
	}
}
