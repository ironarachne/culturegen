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
	Religion          Religion
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

	culture.HomeClimate = climategen.Generate()
	culture.MusicStyle = culture.randomMusicStyle()
	culture.ClothingStyle = culture.generateClothingStyle()
	culture.FoodStyle = culture.generateFoodStyle()
	culture.AlcoholicDrinks = culture.generateDrinks()

	culture.AttributeMax = 100
	culture.Aggression = rand.Intn(culture.AttributeMax) + 1
	culture.Curiosity = rand.Intn(culture.AttributeMax) + 1
	culture.Rigidity = rand.Intn(culture.AttributeMax) + 1
	culture.Superstition = rand.Intn(culture.AttributeMax) + 1

	culture.Religion = culture.generateReligion()

	return culture
}

// SetClimate sets the climate and recalculates some traits
func (culture Culture) SetClimate(climate string) Culture {
	newCulture := culture
	newCulture.HomeClimate = climategen.GetClimate(climate)
	newCulture.MusicStyle = newCulture.randomMusicStyle()
	newCulture.ClothingStyle = newCulture.generateClothingStyle()
	newCulture.FoodStyle = newCulture.generateFoodStyle()
	newCulture.AlcoholicDrinks = newCulture.generateDrinks()

	return newCulture
}
