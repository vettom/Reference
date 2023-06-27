package main

func main() {
	// Read file and print
	// content, err := os.ReadFile("go.mod")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(content))

	// Read file by line using buffio
	/*
		F, err := os.Open("main.go")
		defer F.Close()

		// Read lineby line using BuffIo scanner
		data := bufio.NewScanner(F)
		for data.Scan() {
			fmt.Printf("Line: %s\n", data.Text())
		}
		if err != nil {
			log.Fatal(err)
		}
	*/

}
