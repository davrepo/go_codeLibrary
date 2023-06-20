/*
Implement a paining program. It should support

- Circle with location (x, y), color and radius
- Rectangle with location (x, y), width, height and color

Each type should implement a `Draw(d Device)` method.

Implement an `ImageCanvas` struct which hold a slice of drawable items and has
`Draw(w io.Writer)` that writes a PNG to w (using `image/png`).
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"os"
)

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)

type Shape struct {
	X     int
	Y     int
	Color color.Color
}

// Set() is implemented by '*image.RGBA' type from 'image' package in standard library
// this type is used as the underlying type for the 'img' variable in the 'Draw()' method
// of the 'ImageCanvas' type
type Device interface {
	Set(int, int, color.Color)
}

type Drawer interface {
	Draw(d Device)
}

type Circle struct {
	Shape
	Radius int
}

func NewCircle(x, y, r int, c color.Color) *Circle {
	cr := Circle{
		Shape:  Shape{x, y, c},
		Radius: r,
	}
	return &cr
}

func (c *Circle) Draw(d Device) {
	minX, minY := c.X-c.Radius, c.Y-c.Radius
	maxX, maxY := c.X+c.Radius, c.Y+c.Radius
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			dx, dy := x-c.X, y-c.Y
			if int(math.Sqrt(float64(dx*dx+dy*dy))) <= c.Radius {
				d.Set(x, y, c.Color)
			}
		}
	}
}

type Rectangle struct {
	Shape
	Height int
	Width  int
}

func NewRectangle(x, y, h, w int, c color.Color) *Rectangle {
	r := Rectangle{
		Shape:  Shape{x, y, c},
		Height: h,
		Width:  w,
	}
	return &r
}

func (r *Rectangle) Draw(d Device) {
	minX, minY := r.X-r.Width/2, r.Y-r.Height/2
	maxX, maxY := r.X+r.Width/2, r.Y+r.Height/2
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			d.Set(x, y, r.Color)
		}
	}
}

type ImageCanvas struct {
	width  int
	height int
	shapes []Drawer
}

func NewImageCanvas(width, height int) (*ImageCanvas, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("negative size: width=%d, height=%d", width, height)
	}

	c := ImageCanvas{
		width:  width,
		height: height,
	}
	return &c, nil
}

func (ic *ImageCanvas) Add(d Drawer) {
	ic.shapes = append(ic.shapes, d)
}

func (ic *ImageCanvas) Draw(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, ic.width, ic.height))
	for _, s := range ic.shapes {
		// img implements the 'Device' interface as it has the 'Set()' method
		s.Draw(img)
	}
	return png.Encode(w, img)
}

func main() {
	ic, err := NewImageCanvas(200, 200)
	if err != nil {
		log.Fatal(err)
	}

	ic.Add(NewCircle(100, 100, 80, Green))
	ic.Add(NewCircle(60, 60, 10, Blue))
	ic.Add(NewCircle(140, 60, 10, Blue))
	ic.Add(NewRectangle(100, 130, 10, 80, Red))
	f, err := os.Create("face.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// Draw() takes a 'io.Writer' interface as an argument
	// 'f' implements the 'io.Writer' interface as it has the 'Write()' method
	// so we can pass 'f' to the 'Draw()' method, which will write the image to 'f'
	if err := ic.Draw(f); err != nil {
		log.Fatal(err)
	}
}
