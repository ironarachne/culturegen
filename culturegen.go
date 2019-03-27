package culturegen

import (
	"math/rand"

	"github.com/ironarachne/climategen"
)

// Culture is a fantasy culture
type Culture struct {
	Name          string
	Adjective     string
	Language      Language
	MusicStyle    MusicStyle
	Aggression    int
	Curiosity     int
	Rigidity      int
	Superstition  int
	HomeClimate   climategen.Climate
	ClothingStyle ClothingStyle
	FoodStyle     FoodStyle
}

// GenerateCulture generates a culture
func GenerateCulture() Culture {
	culture := Culture{}

	culture.Language = randomLanguage()

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective
	culture.MusicStyle = randomMusicStyle()

	culture.Aggression = rand.Intn(10) + 1
	culture.Curiosity = rand.Intn(10) + 1
	culture.Rigidity = rand.Intn(10) + 1
	culture.Superstition = rand.Intn(10) + 1

	culture.HomeClimate = climategen.Generate()
	culture.ClothingStyle = culture.generateClothingStyle()
	culture.FoodStyle = culture.generateFoodStyle()

	return culture
}

func (culture Culture) setClimate(climate string) {
	culture.HomeClimate = climategen.GetClimate(climate)
	culture.ClothingStyle = culture.generateClothingStyle()
	culture.FoodStyle = culture.generateFoodStyle()
}
