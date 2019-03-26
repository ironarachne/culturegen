package culturegen

import (
	"strings"
)

// Culture is a fantasy culture
type Culture struct {
	Name                string
	Adjective           string
	LanguageName        string
	LanguageDescriptors []string
	MusicStyle          MusicStyle
}

// GenerateCulture generates a culture
func GenerateCulture() Culture {
	culture := Culture{}

	languageCategory := randomLanguageCategory()

	culture.Name = strings.Title(randomName(languageCategory))
	culture.Adjective = deriveAdjective(culture.Name)
	culture.LanguageName = culture.Adjective
	culture.LanguageDescriptors = append(culture.LanguageDescriptors, languageCategory.Name)
	culture.MusicStyle = randomMusicStyle()

	return culture
}
