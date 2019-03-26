package culturegen

// ClothingStyle describes what kind of clothing the culture wears
type ClothingStyle struct {
	CommonItems     []string
	CommonJewelry   []string
	CommonDyes      []string
	CommonMaterials []string
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

	if inSlice("horses", culture.HomeClimate.Resources) {
		style.CommonItems = append(style.CommonItems, "riding boots")
	}

	if inSlice("oil", culture.HomeClimate.Resources) {
		style.CommonMaterials = append(style.CommonMaterials, "oiled leather")
	}

	return style
}
