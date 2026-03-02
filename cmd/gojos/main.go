package main

import (
	"fmt"

	figure "github.com/common-nighthawk/go-figure"
)

func showBanner() {
// Banner
	  fig := figure.NewFigure("GOJOS", "doom", true)
    fig.Print()

    // Print metadata below it
		fmt.Println()
		fmt.Println("Simple JSON Store")
    fmt.Println("─────────────────────────────────────")
    fmt.Println("Version  : v1.0.0-SNAPSHOT")
    fmt.Println("─────────────────────────────────────")
}

func main() {
    showBanner()
}