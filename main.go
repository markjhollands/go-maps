package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}

	fmt.Println(colours)

	// Different way of declaring Maps
	var colours1 map[string]string

	fmt.Println(colours1)

	// Another Different way of declaring Maps
	colours2 := make(map[string]string)

	colours2["white"] = "#ffffff"
	colours2["blue"] = "#0000ff"

	fmt.Println(colours2)
	fmt.Println(colours2["white"])

	delete(colours2, "blue")

	fmt.Println(colours2)
}
