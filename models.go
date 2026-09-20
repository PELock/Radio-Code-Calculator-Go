/******************************************************************************
 * Radio Code Calculator API - models
 *
 * Version      : v1.1.6
 * Language     : Go
 * Author       : Bartosz Wójcik (support@pelock.com)
 * Homepage     : https://www.pelock.com
 *
 *****************************************************************************/

package radiocodecalculator

import (
	"regexp"
	"strings"
)

// RadioModel describes one calculator (name, lengths, regex patterns).
type RadioModel struct {
	Name                   string
	SerialMaxLen           int
	SerialRegexPatterns    map[string]string
	ExtraMaxLen            int
	ExtraRegexPatterns     map[string]string
	DefaultLanguage        string
}

// Built-in offline validation templates (PHP SDK patterns; Eclipse ESN name fixed).
var (
	RenaultDacia          = NewRadioModel("renault-dacia", 4, `^([A-Z]{1}[0-9]{3})$`, 0, nil)
	ChryslerPanasonicTM9  = NewRadioModel("chrysler-panasonic-tm9", 4, `^([0-9]{4})$`, 0, nil)
	ChryslerDodgeVP       = NewRadioModel("chrysler-dodge-vp", 4, `^([a-zA-Z0-9]{4})$`, 0, nil)
	FordMSeries           = NewRadioModel("ford-m-series", 6, `^([0-9]{6})$`, 0, nil)
	FordVSeries           = NewRadioModel("ford-v-series", 6, `^([0-9]{6})$`, 0, nil)
	FordTravelpilot       = NewRadioModel("ford-travelpilot", 7, `^([0-9]{7})$`, 0, nil)
	FiatStiloBravoVisteon = NewRadioModel("fiat-stilo-bravo-visteon", 6, `^([a-zA-Z0-9]{6})$`, 0, nil)
	FiatDaiichi           = NewRadioModel("fiat-daiichi", 4, `^([0-9]{4})$`, 0, nil)
	FiatVP                = NewRadioModel("fiat-vp", 4, `^([0-9]{4})$`, 0, nil)
	ToyotaERC             = NewRadioModel("toyota-erc", 16, `^([a-zA-Z0-9]{16})$`, 0, nil)
	JeepCherokee          = NewRadioModel("jeep-cherokee", 14, `^([a-zA-Z0-9]{10}[0-9]{4})$`, 0, nil)
	NissanGloveBox        = NewRadioModel("nissan-glove-box", 12, `^([a-zA-Z0-9]{12})$`, 0, nil)
	EclipseESN            = NewRadioModel("eclipse-esn", 6, `^([a-zA-Z0-9]{6})$`, 0, nil)
	JaguarAlpine          = NewRadioModel("jaguar-alpine", 5, `^([0-9]{5})$`, 0, nil)
)

// NewRadioModel builds a model. serialRegex / extraRegex may be a string or map[string]string.
func NewRadioModel(name string, serialMaxLen int, serialRegex any, extraMaxLen int, extraRegex any) *RadioModel {
	m := &RadioModel{
		Name:                name,
		SerialMaxLen:        serialMaxLen,
		SerialRegexPatterns: map[string]string{},
		ExtraMaxLen:         extraMaxLen,
		DefaultLanguage:     "php",
	}
	assignPatterns(m.SerialRegexPatterns, serialRegex, m.DefaultLanguage)
	if extraMaxLen != 0 && extraRegex != nil {
		m.ExtraRegexPatterns = map[string]string{}
		assignPatterns(m.ExtraRegexPatterns, extraRegex, m.DefaultLanguage)
	}
	return m
}

func assignPatterns(dst map[string]string, src any, defaultLang string) {
	switch v := src.(type) {
	case string:
		if v != "" {
			dst[defaultLang] = v
		}
	case map[string]string:
		for k, p := range v {
			dst[k] = p
		}
	case map[string]any:
		for k, p := range v {
			if s, ok := p.(string); ok {
				dst[k] = s
			}
		}
	}
}

// SerialRegexPattern returns the regex for the default language (PCRE slashes stripped).
func (m *RadioModel) SerialRegexPattern() string {
	if m == nil {
		return ""
	}
	if p, ok := m.SerialRegexPatterns["go"]; ok {
		return normalizePCRE(p)
	}
	if p, ok := m.SerialRegexPatterns[m.DefaultLanguage]; ok {
		return normalizePCRE(p)
	}
	for _, p := range m.SerialRegexPatterns {
		return normalizePCRE(p)
	}
	return ""
}

// ExtraRegexPattern returns the extra-field regex, or "" if unused.
func (m *RadioModel) ExtraRegexPattern() string {
	if m == nil || m.ExtraRegexPatterns == nil {
		return ""
	}
	if p, ok := m.ExtraRegexPatterns["go"]; ok {
		return normalizePCRE(p)
	}
	if p, ok := m.ExtraRegexPatterns[m.DefaultLanguage]; ok {
		return normalizePCRE(p)
	}
	for _, p := range m.ExtraRegexPatterns {
		return normalizePCRE(p)
	}
	return ""
}

// Validate checks serial / extra length and regex offline.
func (m *RadioModel) Validate(serial string, extra string) int {
	if len(serial) != m.SerialMaxLen {
		return ErrorInvalidSerialLength
	}
	if matched, err := regexp.MatchString(m.SerialRegexPattern(), serial); err != nil || !matched {
		return ErrorInvalidSerialPattern
	}
	if extra != "" {
		if len(extra) != m.ExtraMaxLen {
			return ErrorInvalidExtraLength
		}
		if matched, err := regexp.MatchString(m.ExtraRegexPattern(), extra); err != nil || !matched {
			return ErrorInvalidExtraPattern
		}
	}
	return ErrorSuccess
}

func normalizePCRE(p string) string {
	p = strings.TrimSpace(p)
	if len(p) >= 2 && p[0] == '/' {
		if i := strings.LastIndex(p[1:], "/"); i >= 0 {
			return p[1 : 1+i]
		}
	}
	return p
}

func modelName(model any) string {
	switch v := model.(type) {
	case string:
		return v
	case *RadioModel:
		if v != nil {
			return v.Name
		}
	case RadioModel:
		return v.Name
	}
	return ""
}
