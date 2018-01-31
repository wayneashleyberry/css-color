package parser

import (
	"errors"
	"fmt"
	"image/color"
	"strings"
)

// Parser implementation
type Parser struct {
	keywords map[string]string
}

// New creates a new parser
func New() *Parser {
	p := Parser{}
	p.keywords = keywords()
	return &p
}

// Convert takes a CSS color value and returns an RGBA struct.
// Supported notations are HEX, RGB, RGBA, HSL, HSLA and Keywords.
func (p *Parser) Convert(s string) (color.RGBA, error) {
	// clean up
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)

	// hex
	if strings.HasPrefix(s, "#") {
		return parseHex(s)
	}

	// rgb()
	if strings.HasPrefix(s, "rgb(") && strings.HasSuffix(s, ")") {
		return parseRGB(s)
	}

	// rgba()
	if strings.HasPrefix(s, "rgba(") && strings.HasSuffix(s, ")") {
		return parseRGBA(s)
	}

	// hsl()
	if strings.HasPrefix(s, "hsl(") && strings.HasSuffix(s, ")") {
		return parseHSL(s)
	}

	// hsla()
	if strings.HasPrefix(s, "hsla(") && strings.HasSuffix(s, ")") {
		return parseHSLA(s)
	}

	if val, ok := p.keywords[s]; ok {
		return parseHex(val)
	}

	if s == "transparent" {
		return color.RGBA{0, 0, 0, 0}, nil
	}

	return color.RGBA{}, errors.New("could not parse input")
}

func parseHex(scol string) (color.RGBA, error) {
	c := color.RGBA{}

	format := "#%02x%02x%02x%02x"
	factor := 1.0 / 255.0

	if len(scol) == 7 {
		scol = scol + "ff"
	}

	if len(scol) == 4 {
		scol = scol + "f"
	}

	if len(scol) == 5 {
		format = "#%1x%1x%1x%1x"
		factor = 1.0 / 15.0
	}

	var r, g, b, a uint8
	n, err := fmt.Sscanf(scol, format, &r, &g, &b, &a)
	if err != nil {
		return c, err
	}
	if n != 4 {
		return c, errors.New("not a hex-color")
	}

	c.R = uint8(float64(r)*factor*255.0 + 0.5)
	c.G = uint8(float64(g)*factor*255.0 + 0.5)
	c.B = uint8(float64(b)*factor*255.0 + 0.5)
	c.A = uint8(float64(a)*factor*255.0 + 0.5)
	return c, nil
}

func parseRGB(s string) (color.RGBA, error) {
	var r, g, b int
	n, err := fmt.Sscanf(s, "rgb(%d,%d,%d)", &r, &g, &b)
	if err != nil {
		return color.RGBA{}, err
	}
	if n != 3 {
		return color.RGBA{}, errors.New("invalid format")
	}

	if r > 255 || g > 255 || b > 255 {
		return color.RGBA{}, errors.New("invalid value")
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}, nil
}

func parseRGBA(s string) (color.RGBA, error) {
	var r, g, b int
	var a float64
	n, err := fmt.Sscanf(s, "rgba(%d,%d,%d,%f)", &r, &g, &b, &a)
	if err != nil {
		return color.RGBA{}, err
	}
	if n != 4 {
		return color.RGBA{}, errors.New("invalid format")
	}

	if r > 255 || g > 255 || b > 255 || a > 1 {
		return color.RGBA{}, errors.New("invalid value")
	}

	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(float64(a)*255.0 + 0.5),
	}, nil
}

func parseHSL(str string) (color.RGBA, error) {
	var hI, sI, lI int
	n, err := fmt.Sscanf(str, "hsl(%d,%d%%,%d%%)", &hI, &sI, &lI)
	if n != 3 {
		return color.RGBA{}, errors.New("invalid format")
	}

	c := color.RGBA{}

	h := float64(hI)
	s := float64(sI) / 100
	l := float64(lI) / 100

	if s > 1.0 || l > 1.0 {
		return c, errors.New("invalid value")
	}

	if s == 0 {
		c.R = uint8(l*255 + 0.5)
		c.G = c.R
		c.B = c.R
		c.A = 255
		return c, err
	}

	var r, g, b float64
	var t1 float64
	var t2 float64
	var tr float64
	var tg float64
	var tb float64

	if l < 0.5 {
		t1 = l * (1.0 + s)
	} else {
		t1 = l + s - l*s
	}

	t2 = 2*l - t1
	h = h / 360
	tr = h + 1.0/3.0
	tg = h
	tb = h - 1.0/3.0

	if tr < 0 {
		tr++
	}
	if tr > 1 {
		tr--
	}
	if tg < 0 {
		tg++
	}
	if tg > 1 {
		tg--
	}
	if tb < 0 {
		tb++
	}
	if tb > 1 {
		tb--
	}

	// Red
	if 6*tr < 1 {
		r = t2 + (t1-t2)*6*tr
	} else if 2*tr < 1 {
		r = t1
	} else if 3*tr < 2 {
		r = t2 + (t1-t2)*(2.0/3.0-tr)*6
	} else {
		r = t2
	}

	// Green
	if 6*tg < 1 {
		g = t2 + (t1-t2)*6*tg
	} else if 2*tg < 1 {
		g = t1
	} else if 3*tg < 2 {
		g = t2 + (t1-t2)*(2.0/3.0-tg)*6
	} else {
		g = t2
	}

	// Blue
	if 6*tb < 1 {
		b = t2 + (t1-t2)*6*tb
	} else if 2*tb < 1 {
		b = t1
	} else if 3*tb < 2 {
		b = t2 + (t1-t2)*(2.0/3.0-tb)*6
	} else {
		b = t2
	}

	c.R = uint8(float64(r)*255.0 + 0.5)
	c.G = uint8(float64(g)*255.0 + 0.5)
	c.B = uint8(float64(b)*255.0 + 0.5)
	c.A = 255
	return c, nil
}

func parseHSLA(str string) (color.RGBA, error) {
	var h, s, l int
	var a float64

	n, err := fmt.Sscanf(str, "hsla(%d,%d%%,%d%%,%f)", &h, &s, &l, &a)
	if err != nil {
		return color.RGBA{}, err
	}
	if n != 4 {
		return color.RGBA{}, errors.New("invalid format")
	}

	hsl, err := parseHSL(fmt.Sprintf("hsl(%d, %d%%, %d%%)", h, s, l))

	hsl.A = uint8(a*255.0 + 0.5)

	return hsl, err
}
