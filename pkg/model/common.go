package model

type Thumb struct {
	Spoof   string `xml:"spoof,attr"`
	Cache   string `xml:"cache,attr"`
	Aspect  string `xml:"aspect,attr"`
	Preview string `xml:"preview,attr"`
	Value   string `xml:",chardata"`
}
