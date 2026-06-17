package location

// SearchPlaceIndexForText is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to GeocodeGeocode or SearchTextSearchText unless you require Grab data.
//
// - SearchPlaceIndexForText is part of a previous Amazon Location Service Places
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The version 2 Geocode operation gives better results in the address
// geocoding use case, while the version 2 SearchText operation gives better
// results when searching for businesses and points of interest.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// Geocodes free-form text, such as an address, name, city, or region to allow you
// to search for Places or points of interest.
//
// Optional parameters let you narrow your search results by bounding box or
// country, or bias your search toward a specific position on the globe.
//
// You can search for places near a given position using BiasPosition , or filter
// results within a bounding box using FilterBBox . Providing both parameters
// simultaneously returns an error.
//
// Search results are returned in order of highest to lowest relevance.
