package main

import "io/ioutil"
import "fmt"
import "encoding/json"

func main() {

	//
	// JSON 文字列があったとして
	//
	data, _ := ioutil.ReadFile("sample.json")
	json_text := string(data)

	//
	// これを parse
	//
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(json_text), &dat)
	if err != nil {
		panic(err)
	}

	//
	// 確認
	//
	fmt.Printf("%v\n", dat)
}
