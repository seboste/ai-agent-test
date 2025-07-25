package ports

import (
	"context"
	"errors"
)

var ErrZoneNotFound = errors.New("zone not found")

type Api interface {
	GetCarbonIntensity(zone string, ctx context.Context) (CarbonIntensityData, error)
	GetZones(ctx context.Context) (ZonesResponse, error)
}