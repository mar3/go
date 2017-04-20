package main

import "fmt"
import "io/ioutil"
import "gopkg.in/yaml.v2"

func main() {

	buf, err := ioutil.ReadFile("settings.yml")
	if err != nil {
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", m["a"])
	fmt.Printf("%d\n", m["b"].(map[interface {}]interface {})["c"])

	fmt.Println("Ok.")
}


