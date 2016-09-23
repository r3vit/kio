package main

import "github.com/r3vit/kio/kio"

func main() {
	newTemplate := kio.NewKio("example", "input.yml", "output.txt")
	newTemplate.Execute()

	newTemplate.List()
}
