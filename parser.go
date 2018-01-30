package parser

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

func TODO() error {
	return errors.New("not yet implemented")
}

type Parser struct {
	keywords map[string]string
}

// New creates a new parser
func New() *Parser {
	p := Parser{}
	p.keywords = keywords()
	return &p
}

func (p *Parser) Convert(s string) (color.RGBA, error) {
	// clean up
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

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
		return c, fmt.Errorf("color: %v is not a hex-color", scol)
	}

	c.R = uint8(float64(r)*factor*255.0 + 0.5)
	c.G = uint8(float64(g)*factor*255.0 + 0.5)
	c.B = uint8(float64(b)*factor*255.0 + 0.5)
	c.A = uint8(float64(a)*factor*255.0 + 0.5)
	return c, nil
}

func parseRGB(s string) (color.RGBA, error) {
	c := color.RGBA{}

	s = strings.Replace(s, "rgb(", "", 1)
	s = strings.Replace(s, ")", "", 1)
	s = strings.Replace(s, " ", "", -1)

	parts := strings.Split(s, ",")
	r, err := strconv.Atoi(parts[0])
	if err != nil {
		return c, err
	}
	g, err := strconv.Atoi(parts[1])
	if err != nil {
		return c, err
	}
	b, err := strconv.Atoi(parts[2])
	if err != nil {
		return c, err
	}

	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)
	c.A = 255

	return c, nil
}

func parseRGBA(s string) (color.RGBA, error) {
	c := color.RGBA{}

	s = strings.Replace(s, "rgba(", "", 1)
	s = strings.Replace(s, ")", "", 1)
	s = strings.Replace(s, " ", "", -1)

	parts := strings.Split(s, ",")
	r, err := strconv.Atoi(parts[0])
	if err != nil {
		return c, err
	}
	g, err := strconv.Atoi(parts[1])
	if err != nil {
		return c, err
	}
	b, err := strconv.Atoi(parts[2])
	if err != nil {
		return c, err
	}
	a, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		return c, err
	}

	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)
	c.A = uint8(float64(a)*255.0 + 0.5)

	return c, nil
}

func parseHSL(s string) (color.RGBA, error) {
	return color.RGBA{}, TODO()
}

func parseHSLA(s string) (color.RGBA, error) {
	return color.RGBA{}, TODO()
}
