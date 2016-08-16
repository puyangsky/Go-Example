package main

import "encoding/json"
import "fmt"

//Server is a server
type Server struct {
	Name string
	ID   string
}

func main() {
	var s Server

	// string json decode to struct
	str := `{"name":"pyt","id":"1"}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// struct encode to json string
	s = Server{"puyangsky", "2"}
	var newJSON []byte
	newJSON, err := json.Marshal(s)
	if err != nil {

	}
	fmt.Println("new json: ", string(newJSON))
}
