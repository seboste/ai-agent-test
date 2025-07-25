package ports

type CarbonIntensityData struct {
	Zone            string  `json:"zone"`
	CarbonIntensity float64 `json:"carbonIntensity"`
}

type Zone struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ZonesResponse struct {
	Zones []Zone `json:"zones"`
}