package components

import (
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/fogleman/gg"
)

const (
	imageWidth  = 1920
	imageHeight = 1080
	fontSize    = 48
)

type user struct {
	BgImagePath string
	OutPath     string
	FontPath    string
	fontSize    float64
	Text        string
}

func Img() {

	var Request user

	Request.BgImagePath = "/home/nishantk/Downloads/peakpx.jpg"
	Request.FontPath = "/home/nishantk/Downloads/oswald/Oswald-Bold.ttf"
	Request.fontSize = 45
	Request.Text = rename
	Request.OutPath = "/home/nishantk/Desktop/code/quotes/text_image.jpg"

	var TextOnImg = func(request user) {

		bgImage, err := gg.LoadImage(request.BgImagePath)
		if err != nil {
			panic(err)
		}

		imgWidth := bgImage.Bounds().Dx()
		imgHeight := bgImage.Bounds().Dy()

		dc := gg.NewContext(imgWidth, imgHeight)
		dc.DrawImage(bgImage, 0, 0)
		if err := dc.LoadFontFace(request.FontPath, request.fontSize); err != nil {
			panic(err)
		}

		x := float64(imgWidth / 2)
		y := float64((imgHeight / 2) - 180)
		maxWidth := float64(imgWidth) - 60.0
		dc.SetColor(color.White)
		dc.DrawStringWrapped(request.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

		outputFile, err := os.Create(request.OutPath)

		if err != nil {
			panic(err)
		}
		defer outputFile.Close()

		if err := jpeg.Encode(outputFile, dc.Image(), nil); err != nil {
			log.Fatal(err)
		}
		log.Println("Text added to the image and saved as", request.OutPath)

	}
	TextOnImg(Request)
}
