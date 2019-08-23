package main

import (
	"fmt"
	"os"
)

func main() {

	// 写入文件
	// f, err := os.Create("lines")
	// if err != nil {
	// 	fmt.Println(err)
	// 	f.Close()
	// 	return
	// }

	// d := []string{"Wlecome", "go is amazing", "I love li li"}

	// for _, v := range d {
	// 	fmt.Fprintln(f, v)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	// err = f.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("file written successfully")

	// append 文件

	f, err := os.OpenFile("lines", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "file handling is easy"
	_, err = fmt.Fprintln(f, newLine)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file append successfully!!")
}
