# Carbon Intensity Provider Service

The carbon intensity provider service provides real-time carbon intensity data by electricity zone.

> **WARNING**
> The implementation is in an early stage. Many things are still missing. Use with care.

## API

The service implements the following endpoints:

- `GET /carbon-intensity/{zone}` - Get carbon intensity data for a specific zone
- `GET /carbon-intensity/zones` - Get list of available zones

All endpoints require a valid JWT Bearer token for authentication.

## Usage

API is defined in `api.yaml`. Service structure follows the ports & adapters architecture pattern.

## Development Status

- [x] Basic service structure
- [x] API interface definition in ports
- [ ] Core business logic implementation  
- [ ] HTTP adapter implementation
- [ ] External data provider adapter implementation