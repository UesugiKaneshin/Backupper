package main

import (
	"backupper"
)

func main() {
	routiner := backupper.NewRoutiner()
	routiner.Run()
}
