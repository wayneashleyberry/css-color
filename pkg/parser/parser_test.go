package parser

import (
	"fmt"
	"image/color"
	"testing"
)

func ExampleParser_Convert_hex() {
	p := New()
	c, _ := p.Convert("#bada55")
	fmt.Println(c.R, c.G, c.B, c.A)
	// Output: 186 218 85 255
}

func ExampleParser_Convert_keyword() {
	p := New()
	c, _ := p.Convert("red")
	fmt.Println(c.R, c.G, c.B, c.A)
	// Output: 255 0 0 255
}

func ExampleParser_Convert_rgba() {
	p := New()
	c, _ := p.Convert("rgba(255, 100, 0, 0.5)")
	fmt.Println(c.R, c.G, c.B, c.A)
	// Output: 255 100 0 128
}

func ExampleParser_Convert_hsl() {
	p := New()
	c, _ := p.Convert("hsl(120, 100%, 25%)")
	fmt.Println(c.R, c.G, c.B, c.A)
	// Output: 0 128 0 255
}

func TestConvert(t *testing.T) {
	testCases := []struct {
		input string
		want  color.RGBA
	}{
		// 3-digit hex codes
		{"#000", color.RGBA{0, 0, 0, 255}},
		{"#f00", color.RGBA{255, 0, 0, 255}},
		{"#abc", color.RGBA{170, 187, 204, 255}},
		// 4-digit hex codes
		{"#0000", color.RGBA{0, 0, 0, 0}},
		{"#0010", color.RGBA{0, 0, 17, 0}},
		// 6-digit hex codes
		{"#ff0000", color.RGBA{255, 0, 0, 255}},
		{"#aabbcc", color.RGBA{170, 187, 204, 255}},
		// 8-digit hex codes
		{"#ff000000", color.RGBA{255, 0, 0, 0}},
		{"#aabbccdd", color.RGBA{170, 187, 204, 221}},
		// rgb()
		{"rgb(1,1,0)", color.RGBA{1, 1, 0, 255}},
		{"rgb(1, 1, 0)", color.RGBA{1, 1, 0, 255}},
		// rgba()
		{"rgba(255, 255, 0, 0)", color.RGBA{255, 255, 0, 0}},
		{"rgba(255, 255, 0, 1)", color.RGBA{255, 255, 0, 255}},
		{"rgba(0, 0, 0, 0.5)", color.RGBA{0, 0, 0, 128}},
		// hsl()
		{"hsl(0, 0%, 100%)", color.RGBA{255, 255, 255, 255}},
		{"hsl(0, 0%, 75%)", color.RGBA{191, 191, 191, 255}},
		{"hsl(0, 0%, 50%)", color.RGBA{128, 128, 128, 255}},
		{"hsl(120, 100%, 25%)", color.RGBA{0, 128, 0, 255}},
		{"hsl(220, 25%, 50%)", color.RGBA{96, 117, 159, 255}},
		// hsla()
		{"hsla(0, 0%, 100%, 1)", color.RGBA{255, 255, 255, 255}},
		{"hsla(0, 0%, 75%, 0.5)", color.RGBA{191, 191, 191, 128}},
		{"hsla(0, 0%, 50%, 0)", color.RGBA{128, 128, 128, 0}},
		{"hsla(120, 100%, 25%, 1)", color.RGBA{0, 128, 0, 255}},
		{"hsla(220, 25%, 50%, 0.5)", color.RGBA{96, 117, 159, 128}},
		// keywords
		{"black", color.RGBA{0, 0, 0, 255}},
		{"red", color.RGBA{255, 0, 0, 255}},
		// computed values
		{"transparent", color.RGBA{0, 0, 0, 0}},
	}

	p := New()

	for _, tc := range testCases {
		got, err := p.Convert(tc.input)
		if err != nil {
			t.Errorf("input: %s, conversion error: %s", tc.input, err)
			continue
		}

		if got.R != tc.want.R || got.G != tc.want.G || got.B != tc.want.B || got.A != tc.want.A {
			t.Errorf("input: %s, got: %+v, wanted: %+v", tc.input, got, tc.want)
		}
	}
}

func TestConvertErrors(t *testing.T) {
	testCases := []struct {
		input string
	}{
		{"kjasdflasdf"},
		{"rgb(999, 1, 1)"},
		{"hsl(1, 1%, 120%)"},
		{"rgba(1, 1, 1, 2)"},
	}

	p := New()

	for _, tc := range testCases {
		got, err := p.Convert(tc.input)
		if err == nil {
			t.Errorf("expected error from: %s, got: %+v", tc.input, got)
		}
	}
}
