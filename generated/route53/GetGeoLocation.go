package route53

// GetGeoLocation is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Gets information about whether a specified geographic location is supported for
// Amazon Route 53 geolocation resource record sets.
//
// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public.
//
// Use the following syntax to determine whether a continent is supported for
// geolocation:
//
// GET /2013-04-01/geolocation?continentcode=two-letter abbreviation for a
// continent
//
// Use the following syntax to determine whether a country is supported for
// geolocation:
//
// GET /2013-04-01/geolocation?countrycode=two-character country code
//
// Use the following syntax to determine whether a subdivision of a country is
// supported for geolocation:
//
// GET /2013-04-01/geolocation?countrycode=two-character country
// code&subdivisioncode=subdivision code
