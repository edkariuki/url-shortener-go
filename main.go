package main

import (
	"fmt"
)

func main() {
	db := connectDB()

	_ = db

	fmt.Println("Hello, World!")
}
