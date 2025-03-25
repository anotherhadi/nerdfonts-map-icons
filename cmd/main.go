package main

import (
	"fmt"

	icons "github.com/anotherhadi/nerdfonts-map-icons"
)

func main() {
	icons := icons.GetAllIcons()

	fmt.Println(icons)
}
