package parser

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
