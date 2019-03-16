package main

import (
	"errors"
	"fmt"
)

func createError(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("出现异常")
	}
	return num, nil
}

func test() {
	result, err := createError(-1)
	if err != nil {
		fmt.Println(result)
		fmt.Println(err)
	}
}
