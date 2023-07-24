package gconf

import (
	"encoding/json"
	"fmt"
	"os"
)

type GlobalConf struct {
	Name     string
	Version  string
	Host     string
	HostPort uint
}

var Globalconf *GlobalConf

func init() {
	Globalconf = &GlobalConf{
		Name:     string("Gerver"),
		Version:  string("V0.0.1"),
		Host:     string("127.0.0.1"),
		HostPort: 7650,
	}
	data, err := os.ReadFile("./Gerver.conf")
	if err != nil || len(data) == 0 {
		fmt.Println("No find Conf")
		data, _ := json.Marshal(Globalconf)
		f, _ := os.Create("./Gerver.conf")
		defer f.Close()
		_, err := f.Write(data)
		if err != nil {
			panic(err)
		}
	}
	data, _ = os.ReadFile("./Gerver.conf")
	err = json.Unmarshal(data, Globalconf)
	if err != nil {
		fmt.Println("Load err : ", err)
	}
}
