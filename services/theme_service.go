package services

import (
	"DBDock/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type ThemeService struct {
	configDir string
	fileName  string
}

var (
	hexRegex = regexp.MustCompile(`^#([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)
	rgbRegex = regexp.MustCompile(`(?i)^rgb\(\s*(\d{1,3})\s*,?\s*(\d{1,3})\s*,?\s*(\d{1,3})\s*\)$`)
)

func NewThemeService(configDir string) *ThemeService {
	return &ThemeService{
		configDir: configDir,
		fileName:  "theme.json",
	}
}

func (ts *ThemeService) LoadTheme() models.ThemeConfig {
	path := filepath.Join(ts.configDir, ts.fileName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return ts.defaultTheme()
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return ts.defaultTheme()
	}

	var rawThemeConfig models.ThemeConfig
	if err := json.Unmarshal(data, &rawThemeConfig); err != nil {
		return ts.defaultTheme()
	}

	themeConfig := models.ThemeConfig{}
	for k, v := range rawThemeConfig {
		normalized, err := normalizeColor(v)
		if err != nil {
			normalized = v
		}
		themeConfig[k] = normalized
	}

	return themeConfig
}

func (ts *ThemeService) defaultTheme() models.ThemeConfig {
	return models.ThemeConfig{
		"primary":                                        "30 59 138",
		"primary-light":                                  "245 247 249",
		"primary-hover":                                  "51 76 149",
		"primary-dark":                                   "30 58 138",
		"primary-dark-hover":                             "30 51 109",
		"secondary":                                      "96 165 250",
		"secondary-dark-hover":                           "90 152 229",
		"danger":                                         "228 41 51",
		"danger-light":                                   "253 239 239",
		"danger-hover":                                   "217 83 83",
		"danger-darker":                                  "194 31 43",
		"accent-dark":                                    "75 207 250",
		"surface-light":                                  "255 255 255",
		"surface-dark":                                   "42 43 56",
		"background-light":                               "248 249 250",
		"background-dark":                                "31 32 41",
		"text-light":                                     "31 41 55",
		"text-dark":                                      "245 245 245",
		"text-secondary-light":                           "75 85 99",
		"text-secondary-dark":                            "209 213 219",
		"text-accent-light":                              "30 58 138",
		"text-accent-dark":                               "75 207 250",
		"text-success":                                   "34 197 94",
		"text-success-light":                             "236 253 245",
		"text-success-hover":                             "22 163 74",
		"text-success-dark":                              "27 58 42",
		"text-success-darker":                            "21 128 61",
		"text-info":                                      "75 207 250",
		"text-info-light":                                "230 249 255",
		"text-info-hover":                                "43 183 230",
		"text-info-dark":                                 "18 54 64",
		"text-info-darker":                               "29 174 216",
		"text-danger":                                    "228 41 51",
		"text-danger-light":                              "253 239 239",
		"text-danger-hover":                              "217 83 83",
		"text-danger-dark":                               "76 28 30",
		"text-danger-darker":                             "194 31 43",
		"text-warning":                                   "245 158 11",
		"text-warning-light":                             "255 251 235",
		"text-warning-hover":                             "217 119 6",
		"text-warning-dark":                              "59 42 5",
		"text-warning-darker":                            "180 83 9",
		"gray-light":                                     "209 213 219",
		"gray-lighter":                                   "241 243 246",
		"gray-placeholder":                               "156 163 175",
		"textfield-empty-light":                          "241 243 246",
		"textfield-empty-dark":                           "42 43 56",
		"textfield-filled-light":                         "245 247 249",
		"textfield-filled-dark":                          "42 43 56",
		"textfield-error-light":                          "253 239 239",
		"textfield-error-dark":                           "76 28 30",
		"textfield-border-light":                         "209 213 219",
		"textfield-border-dark":                          "209 213 219",
		"textfield-border-error-light":                   "228 41 51",
		"textfield-border-error-dark":                    "194 31 43",
		"textfield-border-focused-border-light":          "30 58 138",
		"textfield-border-focused-border-dark":           "51 76 149",
		"textfield-placeholder-light":                    "156 163 175",
		"textfield-placeholder-dark":                     "209 213 219",
		"textfield-placeholder-error-light":              "228 41 51",
		"textfield-placeholder-error-dark":               "194 31 43",
		"select-options-text-light":                      "31 41 55",
		"select-options-text-dark":                       "245 245 245",
		"select-options-text-selected-light":             "30 58 138",
		"select-options-text-selected-dark":              "231 227 255",
		"select-options-background-light":                "255 255 255",
		"select-options-background-dark":                 "42 43 56",
		"select-options-background-hovered-light":        "241 243 246",
		"select-options-background-hovered-dark":         "58 59 78",
		"select-options-background-selected-light":       "244 245 248",
		"select-options-background-selected-dark":        "30 58 138",
		"select-options-background-selected-hover-light": "245 247 249",
		"select-options-background-selected-hover-dark":  "30 51 109",
	}
}

func normalizeColor(input string) (string, error) {
	input = strings.TrimSpace(input)

	if hexRegex.MatchString(input) {
		return hexToRgbString(input)
	}

	if matches := rgbRegex.FindStringSubmatch(input); matches != nil {
		return fmt.Sprintf("%s %s %s", matches[1], matches[2], matches[3]), nil
	}

	parts := strings.Fields(input)
	if len(parts) == 3 {
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil || n < 0 || n > 255 {
				return "", fmt.Errorf("invalid RGB value: %s", input)
			}
		}
		return input, nil
	}

	return "", fmt.Errorf("unrecognized color format: %s", input)
}

func hexToRgbString(hex string) (string, error) {
	hex = strings.TrimPrefix(hex, "#")

	if len(hex) == 3 {
		hex = fmt.Sprintf("%c%c%c%c%c%c",
			hex[0], hex[0],
			hex[1], hex[1],
			hex[2], hex[2],
		)
	}

	if len(hex) != 6 {
		return "", fmt.Errorf("invalid hex color: %s", hex)
	}

	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)

	return fmt.Sprintf("%d %d %d", r, g, b), nil
}
