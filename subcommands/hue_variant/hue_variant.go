package huevariant

import (
	"adventure-tools/utils/files"
	"adventure-tools/utils/rotato"
	"image"
	"image/color"
	"image/draw"
	"math"

	_ "github.com/alexflint/go-arg"
)

type Command struct {
	Source      string `arg:"-s" help:"The source png containing the original sprites"`
	Destination string `arg:"-d" default:"" help:"The destination path to write the resulting png to. Leave blank to write to stdout"`
	Width       int    `arg:"-w" default:"16" help:"The width of a column in pixels"`
	Height      int    `arg:"-h" default:"16" help:"The height of a row in pixels"`
	Begin       int    `arg:"-b" default:"0" help:"The index of the sprite to begin at. The begin index is inclusive."`
	End         int    `arg:"-e" default:"0" help:"The index of the sprite to end at. The end index is exclusive and a value of 0 denotes the whole span"`
	Variants    int    `arg:"-v" default:"3" help:"The number of variants to produce in the output file"`
}

func (c *Command) GetSpan() [2]int {
	return [2]int{c.Begin, c.End}
}

func (c *Command) Run() (err error) {
	img, err := files.LoadPng(c.Source)
	if err != nil {
		return err
	}

	subImage, err := c.SliceSheet(img)
	if err != nil {
		return err
	}

	variants := c.ComposeVariants(subImage)

	if err := files.Save(c.Destination, variants); err != nil {
		return err
	}
	return nil
}

func (c *Command) CellCounts(img image.Image) (cW, cH int) {
	return img.Bounds().Dx() / c.Width, img.Bounds().Dy() / c.Height
}

func (c *Command) SourcePt(img image.Image, cellIdx int) image.Point {
	cW, _ := c.CellCounts(img)
	left := (cellIdx % cW) * c.Width
	top := (cellIdx / cW) * c.Height

	return image.Pt(left, top)
}

func (c *Command) DstRect(cellIdx int) image.Rectangle {
	topLeft := image.Pt(
		cellIdx*c.Width,
		0,
	)
	bottomRight := topLeft.Add(image.Pt(
		c.Width,
		c.Height,
	))

	return image.Rectangle{Min: topLeft, Max: bottomRight}
}

func (c *Command) InitSliceDstImage(start, stop int) draw.Image {
	span := stop - start
	width := span * c.Width
	height := c.Height

	bounds := image.Rect(0, 0, width, height)

	return image.NewRGBA(bounds)
}

func (c *Command) SliceSheet(img image.Image) (image.Image, error) {
	cW, cH := c.CellCounts(img)
	maxCell := cW * cH
	start := 0
	if c.Begin > 0 && c.Begin < maxCell {
		start = c.Begin
	}
	stop := maxCell - 1
	if c.End > c.Begin && c.End < maxCell {
		stop = c.End
	}

	output := c.InitSliceDstImage(start, stop)

	for i := start; i <= stop; i++ {
		draw.Draw(output, c.DstRect(i), img, c.SourcePt(img, i), draw.Src)
	}

	return output, nil
}

func (c *Command) InitCompositionDst(src image.Image) draw.Image {
	dstBounds := image.Rect(
		0, 0,
		src.Bounds().Dx(),
		c.Variants*src.Bounds().Dy(),
	)

	return image.NewRGBA(dstBounds)
}

func (c *Command) ComposeVariants(img image.Image) image.Image {
	dst := c.InitCompositionDst(img)
	for i := 0; i <= c.Variants; i++ {
		angle := float32(i) * 2.0 * float32(math.Pi) / float32(c.Variants)
		hueRotated := &hueRotatedImage{rotation: angle, Image: img}
		draw.Draw(
			dst,
			img.Bounds().Add(image.Pt(0, img.Bounds().Dy()*i)),
			hueRotated,
			hueRotated.Bounds().Min,
			draw.Src,
		)
	}

	return dst
}

type hueRotatedImage struct {
	rotation float32
	image.Image
}

func (r *hueRotatedImage) At(x, y int) color.Color {
	actual := r.Image.At(x, y).(color.RGBA)
	rotated := &color.RGBA{
		actual.R, actual.G, actual.B, actual.A,
	}

	rotato.RotateHue(rotated, r.rotation)
	return rotated
}
