package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Config struct {
	Url string `json:"url"`
	Username string `json:"username"`
	Password string `json:password`
	Action string `json:"action"`
}

func main() {
	dir, _ := os.Executable()
	exPath := filepath.Dir(dir)
	fmt.Println(`exPath:`+exPath)
	config,err := decoderCfgJson()
	if err!=nil {
		return
	}
	httpGet(config.Url)
}

func httpGet(url string)  {
	resp,err := http.Get(url)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func decoderCfgJson() (Config,error) {
	var configStruct Config
	filePtr,err := os.Open("./config.json")
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return configStruct,err
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&configStruct)
	if err != nil {
		fmt.Println("Decoder failed", err.Error())
		return configStruct,err
	} else {
		fmt.Println("Decoder success")
		fmt.Println(configStruct)
	}
	return configStruct,nil
}