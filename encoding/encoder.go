package encoding

import (
	"fmt"
	"image"
	"log"

	"github.com/yahya-lafhal/steganography/utils"
)

func Encode(filename string, message string) {
	img := utils.LoadImage(filename)
	bounds := img.Bounds()

	rgbaImg := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgbaImg.Set(x, y, img.At(x, y))
		}
	}

	binaryMessage := ""
	for _, char := range message {
		binaryMessage += fmt.Sprintf("%08b", char)
	}
	binaryMessage += "00000000"

	index := 0
	for y := bounds.Min.Y; y < bounds.Max.Y && index < len(binaryMessage); y++ {
		for x := bounds.Min.X; x < bounds.Max.X && index < len(binaryMessage); x++ {
			pixel := rgbaImg.RGBAAt(x, y)
			pixel.R = (pixel.R & 0xFE) | (binaryMessage[index] - '0')
			index++
			rgbaImg.SetRGBA(x, y, pixel)
		}
	}

	if index < len(binaryMessage) {
		log.Fatal("Message too large for the image.")
	}

	outputFilename := "encoded_" + filename
	utils.SaveImage(rgbaImg, outputFilename)
	fmt.Printf("Message encoded successfully! Saved to %s\n", outputFilename)
}
