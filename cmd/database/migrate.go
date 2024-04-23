package main

import "fmt"

const (
	migrationsDir = `file://database/migration`
)

func main() {
	fmt.Println(migrationsDir)
}