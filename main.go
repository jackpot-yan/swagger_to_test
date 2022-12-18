package main

import (
	"github.com/flopp/go-findfont"
	"os"
	"strings"
	"swagger_to_test/ui"
)

func init() {
	fontPath := findfont.List()
	for _, path := range fontPath {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			err := os.Setenv("FYNE_FONT", path)
			if err != nil {
				return
			}
		}
	}
}

func main() {
	ui.Home()
}
