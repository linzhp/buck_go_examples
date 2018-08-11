package main

import (
	"github.com/mattn/go-sqlite3"
	"fmt"
)

func main() {
	fmt.Println(sqlite3.Version())
}
