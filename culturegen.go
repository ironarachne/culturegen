package culturegen

import (
	"math/rand"

	"github.com/ironarachne/climategen"
)

// Culture is a fantasy culture
type Culture struct {
	Name              string
	Adjective         string
	Language          Language
	CommonMaleNames   []string
	CommonFamilyNames []string
	CommonFemaleNames []string
	MusicStyle        MusicStyle
	AttributeMax      int
	Aggression        int
	Curiosity         int
	Rigidity          int
	Superstition      int
	HomeClimate       climategen.Climate
	ClothingStyle     ClothingStyle
	FoodStyle         FoodStyle
	AlcoholicDrinks   []Drink
}

// GenerateCulture generates a culture
func GenerateCulture() Culture {
	culture := Culture{}

	culture.Language = randomLanguage()

	culture.CommonMaleNames = culture.Language.generateNameList("male")
	culture.CommonFemaleNames = culture.Language.generateNameList("female")
	culture.CommonFamilyNames = culture.Language.generateNameList("family")

	culture.Name = culture.Language.Name
	culture.Adjective = culture.Language.Adjective
	culture.MusicStyle = randomMusicStyle()

	culture.AttributeMax = 10
	culture.Aggression = rand.Intn(culture.AttributeMax) + 1
	culture.Curiosity = rand.Intn(culture.AttributeMax) + 1
	culture.Rigidity = rand.Intn(culture.AttributeMax) + 1
	culture.Superstition = rand.Intn(culture.AttributeMax) + 1

	culture.HomeClimate = climategen.Generate()
	culture.ClothingStyle = culture.generateClothingStyle()
	culture.FoodStyle = culture.generateFoodStyle()
	culture.AlcoholicDrinks = culture.generateDrinks()

	return culture
}

func (culture Culture) setClimate(climate string) {
	culture.HomeClimate = climategen.GetClimate(climate)
	culture.ClothingStyle = culture.generateClothingStyle()
	culture.FoodStyle = culture.generateFoodStyle()
	culture.AlcoholicDrinks = culture.generateDrinks()
}
