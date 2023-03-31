package scaler

import (
	"fmt"
	"image"
	"image/color"
	"adventure-tools/utils/files"

	_ "github.com/alexflint/go-arg"
)

type Command struct {
	Source      string `arg:"-s" help:"The input file path"`
	Destination string `arg:"-d" help:"The output file path. To write to stdout, do not provide this option"`
	X           uint   `arg:"-x" default:"1" help:"Set the x scaling factor."`
	Y           uint   `arg:"-y" default:"1" help:"Set the x scaling factor."`
}

func (c *Command) Validate() error {
	if c.Source == "" {
		return fmt.Errorf("No source provided")
	}
	if c.X == 0 {
		return fmt.Errorf("Cannot scale x by a zero value")
	}
	if c.Y == 0 {
		return fmt.Errorf("Cannot scale y by a zero value")
	}

	return nil
}

func (c *Command) Run() (err error) {
	if err = c.Validate(); err != nil {
		return err
	}
	img, err := files.LoadPng(c.Source)
	if err != nil {
		return err
	}

	scaled := c.Scale(img)

	if err = files.Save(c.Destination, scaled); err != nil {
		return err
	}

	// Success!
	return nil
}

func (c *Command) scalePx(dst *image.RGBA, l, t int, xScale, yScale uint, col color.Color) {
	for i := 0; i < int(xScale); i++ {
		for j := 0; j < int(yScale); j++ {
			x, y := int(xScale)*l+i, int(yScale)*t+j
			dst.Set(x, y, col)
		}
	}
}

func (c *Command) Scale(img image.Image) image.Image {
	originalSize := img.Bounds()
	scaledSize := image.Rect(0, 0, originalSize.Max.X*int(c.X), originalSize.Max.Y*int(c.Y))
	scaled := image.NewRGBA(scaledSize)
	for sY := 0; sY < originalSize.Max.Y-originalSize.Min.Y; sY++ {
		for sX := 0; sX < originalSize.Max.X-originalSize.Min.X; sX++ {
			c.scalePx(scaled, sX, sY, c.X, c.Y, img.At(originalSize.Min.X+sX, originalSize.Min.Y+sY))
		}
	}

	return scaled
}
