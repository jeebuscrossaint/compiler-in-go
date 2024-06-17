package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/jeebuscrossaint/constrict/src/parser"
)

func main() {
	file_exists := locate_constrict("confile")
	if file_exists {
		fmt.Println("file found")
	} else {
		fmt.Println("file not found")
	}
}

}
