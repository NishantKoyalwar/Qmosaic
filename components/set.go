package components

import (
	"fmt"
	"log"

	"github.com/reujab/wallpaper"
)

func Set() {
	// Specify the path to the image file you want to set as wallpaper
	imagePath := "/home/nishantk/Desktop/code/quotes/text_image.jpg"

	// Set the wallpaper using the provided image path
	err := wallpaper.SetFromFile(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wallpaper set successfully!")
}
