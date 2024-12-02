package decoding

import (
	"fmt"
	"image/color"

	"github.com/yahya-lafhal/steganography/utils"
)

func Decode(filename string) string {
	img := utils.LoadImage(filename)
	bounds := img.Bounds()

	binaryMessage := ""
	message := ""
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := img.At(x, y).(color.RGBA)
			binaryMessage += fmt.Sprintf("%b", pixel.R&1)

			if len(binaryMessage) >= 8 {
				byteValue := binaryMessage[:8]
				binaryMessage = binaryMessage[8:]

				if byteValue == "00000000" {
					return message
				}

				char := 0
				for _, bit := range byteValue {
					char = (char << 1) | int(bit-'0')
				}
				message += string(char)
			}
		}
	}

	return message
}
