package ports

// CarbonIntensity represents the carbon intensity data for a zone.
type CarbonIntensity struct {
	Zone            string  `json:"zone"`
	CarbonIntensity float64 `json:"carbonIntensity"`
}

// Zone represents a single electricity zone.
type Zone struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Zones represents a list of electricity zones.
type Zones struct {
	Zones []Zone `json:"zones"`
}
