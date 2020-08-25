package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	arg2, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	res := arg1 + arg2

	fmt.Println(res)

}
