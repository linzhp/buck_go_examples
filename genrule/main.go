package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
