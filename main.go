// HelloWorld project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		if os.Args[1] == "-path" {
			var stringNum int = 0
			file, err := os.Open(os.Args[2])
			if err != nil {
				fmt.Println("File does not Exist")
				return
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				stringNum++
			}

			fmt.Println(stringNum)
		} else {
			fmt.Println("Wrong parametr name")
		}
	} else {
		fmt.Println("Parametrs error")
	}
}
