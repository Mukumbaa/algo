package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Palette struct
type Palette struct {
	TextColor        string
	SelectedBg       string
	SelectedText     string
	Indicator        string
	InputText        string
	InputPlaceholder string
	InputPrompt      string
}

// DefaultPalette
func DefaultPalette() Palette {
	return Palette{
		TextColor:        "#ffffff",
		SelectedBg:       "#ffffff",
		SelectedText:     "#000000",
		Indicator:        "#ffffff",
		InputText:        "#ffffff",
		InputPlaceholder: "#808080",
		InputPrompt:      "#ffffff",
	}
}
func RosePinePalette() Palette {
	return Palette{
		TextColor:        "#908caa", // Subtle
		SelectedBg:       "#26233a", // Overlay
		SelectedText:     "#e0def4", // Text
		Indicator:        "#c4a7e7", // Iris
		InputText:        "#e0def4", // Text
		InputPlaceholder: "#6e6a86", // Muted
		InputPrompt:      "#c4a7e7", // Iris
	}
}

//structs for lipgloss
type InputStyle struct {
	Text        lipgloss.Style
	Placeholder lipgloss.Style
	Prompt      lipgloss.Style
}

//returned obj
type Theme struct {
	Box       lipgloss.Style
	Text      lipgloss.Style
	Selected  lipgloss.Style
	Indicator lipgloss.Style
	Input     InputStyle
}

// Palette -> Theme
func BuildTheme(p Palette) Theme {
	return Theme{
		Box: lipgloss.NewStyle().
			Padding(1, 2).
			Margin(1, 0),
		Text: lipgloss.NewStyle().
			Foreground(lipgloss.Color(p.TextColor)),
		Selected: lipgloss.NewStyle().
			Background(lipgloss.Color(p.SelectedBg)).
			Foreground(lipgloss.Color(p.SelectedText)).
			Bold(true),
		Indicator: lipgloss.NewStyle().
			Foreground(lipgloss.Color(p.Indicator)),
		Input: InputStyle{
			Text:        lipgloss.NewStyle().Foreground(lipgloss.Color(p.InputText)),
			Placeholder: lipgloss.NewStyle().Foreground(lipgloss.Color(p.InputPlaceholder)),
			Prompt:      lipgloss.NewStyle().Foreground(lipgloss.Color(p.InputPrompt)),
		},
	}
}


func loadThemeConfig(name string) (Palette, error) {
	// init
	p := DefaultPalette()

	switch name{
		case "":
			break
		case "rose-pine":
			return RosePinePalette(),nil
		case "default":
			return p,nil
	}

	// base user path
	configDir, err := os.UserConfigDir()
	if err != nil {
		return p, err // default if no file
	}

	//path build
	path := filepath.Join(configDir, "algo", "theme.conf")

	file, err := os.Open(path)
	if err != nil {
		// default if no file
		return p, nil
	}
	defer file.Close()

	// parse
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// ignore empty string or comment
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "text_color":
			p.TextColor = value
		case "selected_bg":
			p.SelectedBg = value
		case "selected_text":
			p.SelectedText = value
		case "indicator":
			p.Indicator = value
		case "input_text":
			p.InputText = value
		case "input_placeholder":
			p.InputPlaceholder = value
		case "input_prompt":
			p.InputPrompt = value
		}
	}

	return p, scanner.Err()
}

