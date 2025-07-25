package ports

// CarbonIntensityAPI defines the interface for the carbon intensity service.
type CarbonIntensityAPI interface {
	// GetCarbonIntensity returns the carbon intensity for a given zone.
	GetCarbonIntensity(zone string) (*CarbonIntensity, error)
	// GetZones returns the list of available zones.
	GetZones() (*Zones, error)
}
