package canvas

import (
	"image"
	"image/color"
	"image/draw"
)

// Gradient describes a gradient between two colors
type Gradient struct {
	baseObject

	Generator func(w, h int) image.Image

	Translucency float64 // Set a translucency value > 0.0 to fade the raster

	img draw.Image // internal cache for pixel generator
}

// Alpha is a convenience function that returns the alpha value for a raster
// based on its Translucency value. The result is 1.0 - Translucency.
func (g *Gradient) Alpha() float64 {
	return 1.0 - g.Translucency
}

// LinearGradientColor defines start and end color
type LinearGradientColor struct {
	Start color.Color
	End   color.Color
}

func (gc *LinearGradientColor) linearGradient(w, h, x, y int) *color.RGBA64 {
	d := float64(x) / float64(w) // horizontal
	//d := float64(w) / float64(x)
	//d := float64(y) / float64(h) // top down

	// fetch RGBA values
	aR, aG, aB, aA := gc.Start.RGBA()
	bR, bG, bB, bA := gc.End.RGBA()

	// Get difference
	dR := (float64(bR) - float64(aR))
	dG := (float64(bG) - float64(aG))
	dB := (float64(bB) - float64(aB))
	dA := (float64(bA) - float64(aA))

	// Return with applied gradiation
	pixel := &color.RGBA64{
		R: uint16(float64(aR) + d*dR),
		B: uint16(float64(aB) + d*dB),
		G: uint16(float64(aG) + d*dG),
		A: uint16(float64(aA) + d*dA),
	}

	return pixel

}

type pixelGradient struct {
	g   *Gradient
	img draw.Image
}

// NewRectangleLinearGradient returns a new Image instance that dynamically
// renders a rectangular linear gradient
func NewRectangleLinearGradient(start color.Color, end color.Color) *Gradient {

	gc := &LinearGradientColor{
		Start: start,
		End:   end,
	}
	pix := &pixelGradient{}
	pix.g = &Gradient{
		Generator: func(w, h int) image.Image {

			if pix.img == nil || pix.img.Bounds().Size().X != w || pix.img.Bounds().Size().Y != h {
				rect := image.Rect(0, 0, w, h)
				pix.img = image.NewRGBA(rect)
			}

			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					pix.img.Set(x, y, gc.linearGradient(w, h, x, y))
				}
			}
			return pix.img
		},
	}
	return pix.g
}
