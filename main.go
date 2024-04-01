package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: ob  filepath...\n")
		os.Exit(1)
	}

	var wg sync.WaitGroup

	for i := 1; i < len(os.Args); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			path := os.Args[i]
			if check_if_dir(path) {
				tmp := gen_dir(path)
				write_dir(tmp, path)
			} else {
				tmp := gen_file(path)
				write_file(tmp, path)
			}
		}()
	}

	wg.Wait()
}

func check_if_dir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		return true
	}

	return false
}

func write_file(content []byte, source string) {
	file, err := os.OpenFile(source, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		os.Exit(1)
	}

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func write_dir(content []byte, source string) {
	path := source + "/Content.md"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}
