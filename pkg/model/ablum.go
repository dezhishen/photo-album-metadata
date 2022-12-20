package model

// ablum nfo file
// <?xml version="1.0" encoding="UTF-8" standalone="yes" ?>
type Album struct {
	Title                     string   `xml:"title"`
	Artist                    []string `xml:"artist"`
	MusicBrainzAlbumID        string   `xml:"musicbrainzalbumid"`
	MusicBrainzReleaseGroupID []string `xml:"musicbrainzreleasegroupid"`
	ScrapedmbID               string   `xml:"scrapedmbid"`
	ArtistDesc                string   `xml:"artistdesc"`
	Genre                     string   `xml:"genre"`
	Style                     []string `xml:"style"`
	Mood                      []string `xml:"mood"`
	Theme                     []string `xml:"theme"`
	Compilation               bool     `xml:"compilation"`
	Boxset                    bool     `xml:"boxset"`
	Review                    string   `xml:"review"`
	Type                      string   `xml:"type"`
	ReleaseStatus             string   `xml:"releasestatus"`
	ReleaseDate               string   `xml:"releasedate"`
	OriginalReleaseDate       string   `xml:"originalreleasedate"`
	Label                     []string `xml:"label"`
	Duration                  string   `xml:"duration"`
	Thumb                     []string `xml:"thumb"`
	Path                      string   `xml:"path"`
	Rating                    Rating   `xml:"rating"`
	UserRating                Rating   `xml:"userrating"`
	Votes                     string   `xml:"votes"`
	AlbumArtistCredits        []Artist `xml:"albumArtistCredits>artist"`
	ReleaseType               string   `xml:"releasetype"`
}

type Rating struct {
	Max    string `xml:"max,attr"`
	Rating string `xml:",chardata"`
}
