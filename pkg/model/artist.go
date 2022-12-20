package model

// artist nfo file
// <?xml version="1.0" encoding="UTF-8" standalone="yes" ?>
type Artist struct {
	Name                string   `xml:"name"`
	MusicBrainzArtistID string   `xml:"musicBrainzArtistID"`
	SortName            string   `xml:"sortname"`
	Type                string   `xml:"type"`
	Disambiguation      string   `xml:"disambiguation"`
	Genre               string   `xml:"genre"`
	Style               []string `xml:"style"`
	Mood                []string `xml:"mood"`
	YearsActive         string   `xml:"yearsactive"`
	Born                string   `xml:"born"`
	Formed              string   `xml:"formed"`
	Biography           string   `xml:"biography"`
	Died                string   `xml:"died"`
	Disbanded           string   `xml:"disbanded"`
	Thumb               []Thumb  `xml:"thumb"`
	Path                string   `xml:"path"`
}

type Thumb struct {
	Spoof   string `xml:"spoof,attr"`
	Cache   string `xml:"cache,attr"`
	Aspect  string `xml:"aspect,attr"`
	Preview string `xml:"preview,attr"`
	Value   string `xml:",chardata"`
}
