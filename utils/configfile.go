package utils

import (
	"fmt"
	"os"
)

type ConfigFIle struct {
	Env        string
	BaseDir    string
	Path       string
	JsonStruct *JsonStruct
}

func (c *ConfigFIle) ConfigFileHandler() {
	c.Env = os.Getenv("GO_ENV")
	if c.Env == "" {
		c.Env = "default"
	}
	c.Path = c.BaseDir + "/" + c.Env + ".json"
	exist, err := PathExists(c.Path)
	if exist == false {
		fmt.Println(err)
	}
	c.JsonStruct = NewJsonStruct()
}
