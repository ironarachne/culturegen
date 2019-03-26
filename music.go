package culturegen

import "math/rand"

// MusicStyle is a cultural music style
type MusicStyle struct {
	Structure   int
	Vocals      int
	Beat        int
	Tonality    int
	Descriptors []string
	Instruments []string
}

func randomMusicStyle() MusicStyle {
	styles := []MusicStyle{
		MusicStyle{
			3,
			0,
			1,
			3,
			[]string{"melodic", "rigid", "structured", "regimented"},
			[]string{"lyre", "lute", "flute", "violin", "viola", "cello", "bass"},
		},
		MusicStyle{
			3,
			3,
			1,
			3,
			[]string{"ethereal", "haunting", "operatic", "moody"},
			[]string{"flute", "pan pipes"},
		},
		MusicStyle{
			1,
			2,
			3,
			2,
			[]string{"primitive", "driving", "savage", "powerful"},
			[]string{"drums"},
		},
	}

	return styles[rand.Intn(len(styles)-1)]
}
