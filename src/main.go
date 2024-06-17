package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"strings"
)

func locate_constrict(filename string) bool {
	files, err := ioutil.ReadDir(".") // read current dir
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error reading directory. Ensure you have read permissions.")
	}

	for _, file := range files {
		if file.Name() == filename {
			return true
		}
	}
	
	return false
}

func parse_vars(filename string) (map[string]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    vars := make(map[string]string)
    scanner := bufio.NewScanner(file)
    lineNumber := 0
    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        parts := strings.Split(line, "=")
        if len(parts) != 2 {
            continue
        }
        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])

        // Check if the key already exists
        if _, exists := vars[key]; exists {
            return nil, fmt.Errorf("variable %s is assigned more than once at line %d", key, lineNumber)
        }

        // Check if the value starts with another variable name
        if len(value) > 0 && value[0] == '!' {
            otherVar := value[1:]
            otherValue, exists := vars[otherVar]
            if !exists {
                return nil, fmt.Errorf("variable %s is not defined at line %d", otherVar, lineNumber)
            }
            value = otherValue
        }

        vars[key] = value
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return vars, nil
} 


func parse_functions() {

}

func print_fucking_everything() {

}

func possible_incremental_building() {

}

func detect_shit() {

}

func main() {
	file_exists := locate_constrict("confile")
	if file_exists {
		fmt.Println("file found")
		vars, err := parse_vars("confile")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(vars)
	} else {
		fmt.Println("file not found")
	}
}

