package culturegen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/climategen"
	"github.com/ironarachne/random"
)

// ClothingStyle describes what kind of clothing the culture wears
type ClothingStyle struct {
	CommonItems     []string
	CommonJewelry   []string
	CommonColors    []string
	CommonMaterials []string
	DecorativeStyle string
}

func (culture Culture) generateClothingStyle() ClothingStyle {
	style := ClothingStyle{}

	if culture.HomeClimate.Temperature <= 3 {
		style.CommonMaterials = []string{"thick hides", "furs", "thick wool", "linen"}
		style.CommonItems = []string{"mittens", "scarves", "fur hats", "long tunics", "long fur boots", "long coats", "hooded cloaks", "light underclothes", "long pants", "long dresses"}
	} else if culture.HomeClimate.Temperature > 3 && culture.HomeClimate.Temperature <= 7 {
		style.CommonMaterials = []string{"cotton", "linen", "furs", "light hides"}
		style.CommonItems = []string{"cloaks", "short tunics", "long pants", "long dresses", "short boots", "shoes"}
	} else {
		style.CommonMaterials = []string{"light linen", "light wool"}
		style.CommonItems = []string{"short tunics", "pantaloons", "skirts", "sandals", "light shoes", "short boots"}
	}

	if climategen.IsTypeInResources("mount", culture.HomeClimate.Resources) {
		style.CommonItems = append(style.CommonItems, "riding boots")
	}

	if climategen.IsTypeInResources("oil", culture.HomeClimate.Resources) {
		style.CommonMaterials = append(style.CommonMaterials, "oiled cloth")
	}

	style.CommonJewelry = culture.generateJewelry()

	style.DecorativeStyle = culture.randomDecorativeStyle()
	style.CommonColors = randomColorSet()

	return style
}

func (culture Culture) generateJewelry() []string {
	var jewelryItem string
	var gemProbability int

	jewelry := []string{}

	descriptors := []string{
		"brilliant",
		"gaudy",
		"lustrous",
		"ornate",
		"simple",
	}

	foundations := []string{
		"anklets",
		"bracelets",
		"chokers",
		"necklaces",
		"rings",
	}

	settings := []string{
		"adorned with",
		"decorated with",
		"set with",
	}

	materials := []string{}
	gems := []string{}

	for _, r := range culture.HomeClimate.Resources {
		if r.Type == "metal ingot" {
			materials = append(materials, strings.TrimSuffix(r.Name, " ingot"))
		} else if r.Type == "metal bar" {
			materials = append(materials, strings.TrimSuffix(r.Name, " bar"))
		} else if r.Type == "gem" {
			gems = append(gems, r.Name)
		}
	}

	numberOfJewelryPieces := rand.Intn(4) + 1

	for i := 0; i < numberOfJewelryPieces; i++ {
		jewelryItem = random.Item(descriptors) + " " + random.Item(materials) + " " + random.Item(foundations)
		if len(gems) > 0 {
			gemProbability = rand.Intn(10) + 1
			if gemProbability > 5 {
				jewelryItem += " " + random.Item(settings) + " " + random.Item(gems)
			}
		}

		jewelry = append(jewelry, jewelryItem)
	}

	return jewelry
}

func (culture Culture) randomDecorativeStyle() string {
	styles := []string{
		"beads",
		"complex embroidery",
		"geometric shapes",
		"knotwork",
		"plain",
		"simple embroidery",
		"stripes",
		"stylized animals",
		"tassels",
	}

	if climategen.IsTypeInResources("ivory", culture.HomeClimate.Resources) {
		styles = append(styles, "ivory decorations")
	}

	if climategen.IsTypeInResources("feathers", culture.HomeClimate.Resources) {
		styles = append(styles, "feather decorations")
	}

	if culture.HomeClimate.Temperature < 6 {
		styles = append(styles, "many layers")
	}

	return random.Item(styles)
}

func randomColorSet() []string {
	var newColor string
	var colorSet []string
	var saturation string

	colors := []string{
		"red",
		"blue",
		"green",
		"black",
		"white",
		"yellow",
		"purple",
		"orange",
		"grey",
	}

	for i := 0; i < 3; i++ {
		newColor = random.Item(colors)

		if !inSlice(newColor, []string{"white", "black"}) {
			saturation = randomSaturation()
			newColor = saturation + " " + newColor
		}

		if !inSlice(newColor, colorSet) {
			colorSet = append(colorSet, newColor)
		}
	}

	return colorSet
}

func randomSaturation() string {
	saturations := []string{
		"bright",
		"dark",
		"light",
		"moderate",
		"pastel",
		"subdued",
	}

	return random.Item(saturations)
}
