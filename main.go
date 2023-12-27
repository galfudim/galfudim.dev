package main

import (
	"context"
	"os"
)

func main() {
	component := hello("Gal")
	w, _ := os.Create("./index.html")
	err := component.Render(context.Background(), w)
	if err != nil {
		return
	}
}
