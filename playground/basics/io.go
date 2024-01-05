package basics

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func FileIOBasics() {

	fmt.Println("File IO Basics")
	fmt.Println("--------------")

	writeFile("data/fileIO.txt")

}

func BufferedScan() {
	// Create a new scanner that reads from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Continue scanning until there is no more input
	for scanner.Scan() {
		// Print the text of the current line
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
func writeFile(fname string) {

	// file, err := os.Create(fname)

	// If the file doesn't exist, create it, or append to the file
	file, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintln(writer, "Hello, world!")
	fmt.Fprintln(writer, "Another line.")
	writer.Flush()
}

func readFile(fname string) {
	file, err := os.Open(fname) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
