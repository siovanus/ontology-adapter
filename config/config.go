package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	DEFAULT_CONFIG_FILE_NAME = "./config.json"
)

//Default config instance
var DefConfig = NewConfig()

//Config object used by instance
type Config struct {
	OracleContractAddress string
	OntologyWalletPath    string
	OntologyRpc           string
	Listening             string
}

//NewConfig retuen a Config instance
func NewConfig() *Config {
	return &Config{}
}

//Init Config with a config file
func (this *Config) Init(fileName string) error {
	err := this.loadConfig(fileName)
	if err != nil {
		return fmt.Errorf("loadConfig error:%s", err)
	}
	return nil
}

func (this *Config) loadConfig(fileName string) error {
	data, err := this.readFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, this)
	if err != nil {
		return fmt.Errorf("json.Unmarshal Config:%s error:%s", data, err)
	}
	return nil
}

func (this *Config) readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("file %s close error %s", fileName, err))
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}
