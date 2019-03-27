package culturegen

import (
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

	if climategen.IsTypeInResources("precious metal", culture.HomeClimate.Resources) {
		for _, r := range culture.HomeClimate.Resources {
			if inSlice("precious metal", r.Types) {
				style.CommonJewelry = append(style.CommonJewelry, r.Name+" necklaces")
				style.CommonJewelry = append(style.CommonJewelry, r.Name+" rings")
				style.CommonJewelry = append(style.CommonJewelry, r.Name+" bracelets")
			}
			if inSlice("gemstone", r.Types) {
				style.CommonJewelry = append(style.CommonJewelry, "necklaces set with "+r.Name)
				style.CommonJewelry = append(style.CommonJewelry, "pendants set with "+r.Name)
				style.CommonJewelry = append(style.CommonJewelry, "rings set with "+r.Name)
			}
		}
	}

	style.DecorativeStyle = culture.randomDecorativeStyle()
	style.CommonColors = randomColorSet()

	return style
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
