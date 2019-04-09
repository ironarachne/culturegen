package culturegen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/random"
)

// MusicStyle is a cultural music style
type MusicStyle struct {
	Structure   int
	Vocals      int
	Beat        int
	Tonality    int
	Descriptors []string
	Instruments []string
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

func (culture Culture) generateMusicalInstruments() []string {
	var instrument string

	hides := []string{}
	metals := []string{}
	woods := []string{}
	instrumentTypes := []string{}
	instruments := []string{}

	for _, r := range culture.HomeClimate.Resources {
		if r.Type == "metal ingot" {
			metals = append(metals, strings.TrimSuffix(r.Name, " ingot"))
		} else if r.Type == "metal bar" {
			metals = append(metals, strings.TrimSuffix(r.Name, " bar"))
		} else if r.Type == "wood planks" {
			woods = append(woods, strings.TrimSuffix(r.Name, " planks"))
		} else if r.Type == "hide" {
			hides = append(hides, r.Name)
		}
	}

	if len(hides) > 0 && len(woods) > 0 {
		instrumentTypes = append(instrumentTypes, "drum")
		instrumentTypes = append(instrumentTypes, "bagpipes")
	}

	if len(metals) > 0 {
		instrumentTypes = append(instrumentTypes, "horn")
		instrumentTypes = append(instrumentTypes, "pan pipes")
		instrumentTypes = append(instrumentTypes, "flute")
	}

	if len(woods) > 0 {
		instrumentTypes = append(instrumentTypes, "lyre")
		instrumentTypes = append(instrumentTypes, "lute")
	}

	baseMaterials := metals
	baseMaterials = append(baseMaterials, woods...)

	numberOfInstruments := rand.Intn(3) + 1

	for i := 0; i < numberOfInstruments; i++ {
		instrument = random.Item(instrumentTypes)
		if instrument == "drum" {
			instrument = random.Item(woods) + " drum skinned with " + random.Item(hides)
		} else if instrument == "lyre" || instrument == "lute" {
			instrument = random.Item(woods) + " " + instrument
		} else {
			instrument = random.Item(baseMaterials) + " " + instrument
		}

		instruments = append(instruments, instrument)
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
