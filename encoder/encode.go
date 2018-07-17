// +build !jpeg

// Package encoder.
package encoder

import (
	"image"
	"io"

	jpeg "github.com/antonini/golibjpegturbo"
)

// New returns a new Encoder.
func New(w io.Writer) *Encoder {
	return &Encoder{w}
}

// Encoder struct.
type Encoder struct {
	w io.Writer
}

// Encode encodes image to JPEG.
func (e Encoder) Encode(img image.Image) error {
	err := jpeg.Encode(e.w, img, &jpeg.Options{Quality: 75})
	if err != nil {
		return err
	}

	return nil
}
