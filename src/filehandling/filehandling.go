package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	//1. reading file
	// data, err := ioutil.ReadFile("./test.txt")
	// if err != nil {
	// 	fmt.Println("file reading eror", err)
	// 	return
	// }
	// fmt.Println("contents of file: ", string(data))

	//2. flag

	// fptr := flag.String("fpath", "test.txt", "file path to read from")
	// flag.Parse()
	// fmt.Println("value of fpath is ", *fptr)
	// data, err := ioutil.ReadFile(*fptr)
	// if err != nil {
	// 	fmt.Println("file reading eror", err)
	// 	return
	// }
	// fmt.Println("contents of file: ", string(data))

	// 3. reading file by chunks
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	// by byte
	// r := bufio.NewReader(f)
	// b := make([]byte, 1)
	// for {
	// 	n, err := r.Read(b)
	// 	if err != nil {
	// 		fmt.Println("Error reading file:", err)
	// 		break
	// 	}
	// 	fmt.Println(string(b[0:n]))
	// }
	// by line

	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
