package culturegen

import (
	"math/rand"
	"strings"

	"github.com/ironarachne/random"
)

// Language is a fantasy language
type Language struct {
	Name          string
	Adjective     string
	Descriptors   []string
	Category      LanguageCategory
	IsTonal       bool
	WritingSystem WritingSystem
}

// LanguageCategory is a style of language
type LanguageCategory struct {
	Name             string
	WordLength       int
	Initiators       []string
	Connectors       []string
	Finishers        []string
	MasculineEndings []string
	FeminineEndings  []string
}

// LanguageMutation is a word mutation
type LanguageMutation struct {
	From string
	To   string
}

// WritingSystem is a system of writing
type WritingSystem struct {
	Name           string
	Classification string
	StrokeType     string
}

var (
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	vowels     = []string{"a", "e", "i", "o", "u"}
	glottals   = []string{"g", "k", "ch"}
	velars     = []string{"k", "g", "ng", "w"}
	sibilants  = []string{"s", "f"}
	fricatives = []string{"f", "v", "th", "ð", "s", "z", "∫", "zh"}
	liquids    = []string{"l", "r"}
	nasals     = []string{"m", "n", "ng"}
	glides     = []string{"j", "w"}
	growls     = []string{"br", "tr", "gr", "dr", "kr"}
	stops      = []string{"p", "b", "t", "d", "k", "g"}
)

func deriveAdjective(name string) string {
	var suffix string

	adjective := name
	lastCharacter := adjective[len(adjective)-1:]

	potentialSuffixes := []string{"n", "lese", "ish"}

	if inSlice(lastCharacter, consonants) {
		potentialSuffixes = []string{"ish", "ian", "an", "i", "ese"}
	}

	suffix = random.Item(potentialSuffixes)

	adjective += suffix

	return adjective
}

func (language Language) generateNameList(nameType string) []string {
	var names []string
	var name string
	var endings []string

	if nameType == "male" {
		endings = language.Category.MasculineEndings
	} else if nameType == "female" {
		endings = language.Category.FeminineEndings
	} else {
		for i := 0; i < 5; i++ {
			endings = append(endings, randomSyllable(language.Category, "finisher"))
		}
	}

	for i := 0; i < 10; i++ {
		name = language.RandomName() + random.Item(endings)
		if !inSlice(name, names) {
			names = append(names, name)
		}
	}

	return names
}

func mutateName(name string) string {
	mutation := randomMutation()

	name = strings.Replace(name, mutation.From, mutation.To, 1)

	return name
}

func randomLanguageCategory() LanguageCategory {
	languageCategories := []LanguageCategory{
		LanguageCategory{
			"musical",
			2,
			fricatives,
			liquids,
			sibilants,
			[]string{"ion", "on", "en", "o"},
			[]string{"i", "a", "ia", "ila"},
		},
		LanguageCategory{
			"guttural",
			1,
			append(glottals, growls...),
			velars,
			velars,
			[]string{"ur", "ar", "ach", "ag"},
			[]string{"a", "agi"},
		},
		LanguageCategory{
			"abrupt",
			2,
			stops,
			fricatives,
			liquids,
			[]string{"on", "en", "an"},
			[]string{"a", "e", "et"},
		},
		LanguageCategory{
			"nasal",
			2,
			glottals,
			stops,
			nasals,
			[]string{"een", "oon", "in", "en"},
			[]string{"ini", "nia", "mia", "mi"},
		},
	}
	return languageCategories[rand.Intn(len(languageCategories)-1)]
}

func randomLanguage() Language {
	var language Language

	language.Category = randomLanguageCategory()
	language.Name = strings.Title(randomLanguageName(language.Category))
	language.Descriptors = append(language.Descriptors, language.Category.Name)
	language.Adjective = deriveAdjective(language.Name)

	tonalChance := rand.Intn(10) + 1
	if tonalChance > 7 {
		language.IsTonal = true
	} else {
		language.IsTonal = false
	}

	language.WritingSystem = randomWritingSystem()
	language.WritingSystem.Name = language.Adjective

	return language
}

func randomLanguageName(languageCategory LanguageCategory) string {
	var name string
	var syllables []string
	skewLonger := false

	if rand.Intn(10) > 3 {
		skewLonger = true
	}

	randomLength := rand.Intn(languageCategory.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllables = append(syllables, randomSyllable(languageCategory, role))
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := rand.Intn(10) + 1
	if chance > 3 {
		name = mutateName(name)
	}

	return name
}

func randomMutation() LanguageMutation {
	rules := []LanguageMutation{
		LanguageMutation{
			"s",
			"ss",
		},
		LanguageMutation{
			"s",
			"sh",
		},
		LanguageMutation{
			"f",
			"ff",
		},
		LanguageMutation{
			"f",
			"fh",
		},
		LanguageMutation{
			"g",
			"gh",
		},
		LanguageMutation{
			"l",
			"l'",
		},
	}

	return rules[rand.Intn(len(rules)-1)]
}

// RandomGenderedName generates a random gendered first name
func (language Language) RandomGenderedName(gender string) string {
	var endings []string

	name := language.RandomName()

	if gender == "male" {
		endings = language.Category.MasculineEndings
	} else {
		endings = language.Category.FeminineEndings
	}

	name = name + random.Item(endings)

	return name
}

// RandomName generates a random name using the language
func (language Language) RandomName() string {
	var name string
	var syllables []string
	skewLonger := false

	if rand.Intn(10) > 7 {
		skewLonger = true
	}

	randomLength := rand.Intn(language.Category.WordLength) + 1

	if skewLonger {
		randomLength++
	}

	role := "connector"

	for i := 0; i < randomLength; i++ {
		if randomLength-i == 1 {
			role = "finisher"
		}
		syllables = append(syllables, randomSyllable(language.Category, role))
	}

	for _, syllable := range syllables {
		name += syllable
	}

	chance := rand.Intn(10) + 1
	if chance > 3 {
		name = mutateName(name)
	}

	name = strings.Title(name)

	return name
}

func randomSyllable(category LanguageCategory, role string) string {
	syllable := random.Item(category.Initiators) + random.Item(vowels)
	expand := rand.Intn(10) + 1
	if expand > 2 {
		if role == "connector" {
			syllable += random.Item(category.Connectors)
		} else {
			syllable += random.Item(category.Finishers)
		}
	}

	return syllable
}

func randomStrokeType() string {
	strokeTypes := []string{
		"arcs and loops",
		"dots and slashes",
		"flowing",
		"geometric shapes",
		"right angles",
		"runic",
	}

	return random.Item(strokeTypes)
}

func randomWritingSystem() WritingSystem {
	var writingSystem WritingSystem

	classifications := []string{
		"abjad",
		"abugida",
		"alphabet",
		"ideograms",
		"pictograms",
		"semanto-phonetic",
		"syllabary",
	}

	writingSystem.Classification = random.Item(classifications)
	writingSystem.StrokeType = randomStrokeType()

	return writingSystem
}
