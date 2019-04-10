package culturegen

import (
	"bytes"
	"html/template"
	"math/rand"

	"github.com/ironarachne/random"
)

// MusicStyle is a cultural music style
type MusicStyle struct {
	Structure   int
	Vocals      int
	Beat        int
	Tonality    int
	Descriptors []string
	Instruments []Instrument
}

// Instrument is a musical instrument
type Instrument struct {
	Name                   string
	Description            string
	Type                   string
	BaseMaterialOptions    []string
	SupportMaterialOptions []string
	BaseMaterial           string
	SupportMaterial        string
	DescriptionTemplate    string
}

func (culture Culture) generateMusicDescriptors() []string {
	descriptors := []string{}

	possibleDescriptors := []string{
		"airy",
		"bombastic",
		"booming",
		"breathy",
		"bright",
		"cheerful",
		"driving",
		"dynamic",
		"energetic",
		"ethereal",
		"euphonic",
		"fast",
		"full-toned",
		"haunting",
		"lilting",
		"lofty",
		"mellifluous",
		"mellow",
		"melodic",
		"moody",
		"operatic",
		"orotund",
		"percussive",
		"powerful",
		"primitive",
		"regimented",
		"resonant",
		"rigid",
		"savage",
		"somber",
		"structured",
		"tumid",
		"uplifting",
		"vibrant",
		"warm",
	}

	numberOfDescriptors := rand.Intn(4) + 1

	for i := 0; i < numberOfDescriptors; i++ {
		descriptors = append(descriptors, random.Item(possibleDescriptors))
	}

	return descriptors
}

func (culture Culture) generateMusicalInstruments() []Instrument {
	var instrument Instrument
	var materialType string
	var availableBaseMaterials []string
	var availableSupportMaterials []string

	availableHides := []string{}
	availableMetals := []string{}
	availableWoods := []string{}
	availableMaterials := []string{}

	allInstruments := getAllInstruments()
	availableInstruments := []Instrument{}
	instruments := []Instrument{}

	for _, i := range culture.HomeClimate.CommonMetals {
		availableMetals = append(availableMetals, i.Name)
	}

	for _, i := range culture.HomeClimate.PreciousMetals {
		availableMetals = append(availableMetals, i.Name)
	}

	for _, i := range culture.HomeClimate.Plants {
		if i.IsTree {
			availableWoods = append(availableWoods, i.Name)
		}
	}

	for _, i := range culture.HomeClimate.Animals {
		if i.GivesHide {
			availableHides = append(availableHides, i.Name)
		}
	}

	if len(availableHides) > 0 {
		availableMaterials = append(availableMaterials, "hide")
	}
	if len(availableMetals) > 0 {
		availableMaterials = append(availableMaterials, "metal")
	}
	if len(availableWoods) > 0 {
		availableMaterials = append(availableMaterials, "wood")
	}

	for _, i := range allInstruments {
		if slicePartlyWithin(i.BaseMaterialOptions, availableMaterials) {
			if slicePartlyWithin(i.SupportMaterialOptions, availableMaterials) {
				availableInstruments = append(availableInstruments, i)
			}
		}
	}

	numberOfInstruments := rand.Intn(3) + 1

	for i := 0; i < numberOfInstruments; i++ {
		instrument = availableInstruments[rand.Intn(len(availableInstruments)-1)]
		availableBaseMaterials = []string{}
		availableSupportMaterials = []string{}

		for _, m := range instrument.BaseMaterialOptions {
			if inSlice(m, availableMaterials) {
				availableBaseMaterials = append(availableBaseMaterials, m)
			}
		}

		for _, m := range instrument.SupportMaterialOptions {
			if inSlice(m, availableMaterials) {
				availableSupportMaterials = append(availableSupportMaterials, m)
			}
		}

		materialType = random.Item(availableBaseMaterials)
		if materialType == "hide" {
			instrument.BaseMaterial = random.Item(availableHides)
		} else if materialType == "metal" {
			instrument.BaseMaterial = random.Item(availableMetals)
		} else if materialType == "wood" {
			instrument.BaseMaterial = random.Item(availableWoods)
		}

		materialType = random.Item(availableSupportMaterials)
		if materialType == "hide" {
			instrument.SupportMaterial = random.Item(availableHides)
		} else if materialType == "metal" {
			instrument.SupportMaterial = random.Item(availableMetals)
		} else if materialType == "wood" {
			instrument.SupportMaterial = random.Item(availableWoods)
		}

		instrument.Description = instrument.getDescription()

		instruments = append(instruments, instrument)
	}

	return instruments
}

func (instrument Instrument) getDescription() string {
	t := template.New("instrument description")

	var err error
	t, err = t.Parse(instrument.DescriptionTemplate)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, instrument); err != nil {
		panic(err)
	}

	result := tpl.String()

	return result
}

func getAllInstruments() []Instrument {
	instruments := []Instrument{
		Instrument{
			Name:                   "short flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		Instrument{
			Name:                   "long flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		Instrument{
			Name:                   "twin flute",
			Type:                   "flute",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} trimmed with {{.SupportMaterial}}",
		},
		Instrument{
			Name:                   "short harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "long harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "full harp",
			Type:                   "harp",
			BaseMaterialOptions:    []string{"metal", "wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "lyre",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "lijerica",
			Type:                   "lyre",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "long-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "pierced lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "short-necked lute",
			Type:                   "lute",
			BaseMaterialOptions:    []string{"wood"},
			SupportMaterialOptions: []string{"hide"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} strung with {{.SupportMaterial}} sinew",
		},
		Instrument{
			Name:                   "single-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drone",
		},
		Instrument{
			Name:                   "multiple-drone bagpipes",
			Type:                   "bagpipes",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}}-hide {{.Name}} with {{.SupportMaterial}} drones",
		},
		Instrument{
			Name:                   "hand drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		Instrument{
			Name:                   "short drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		Instrument{
			Name:                   "walking drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
		Instrument{
			Name:                   "heavy drum",
			Type:                   "drum",
			BaseMaterialOptions:    []string{"hide"},
			SupportMaterialOptions: []string{"metal", "wood"},
			DescriptionTemplate:    "{{.BaseMaterial}} {{.Name}} skinned with {{.SupportMaterial}} hide",
		},
	}

	return instruments
}

func (culture Culture) randomMusicStyle() MusicStyle {
	style := MusicStyle{}

	style.Beat = rand.Intn(3)
	style.Structure = rand.Intn(3)
	style.Tonality = rand.Intn(3)
	style.Vocals = rand.Intn(3)

	style.Descriptors = culture.generateMusicDescriptors()
	style.Instruments = culture.generateMusicalInstruments()

	return style
}
