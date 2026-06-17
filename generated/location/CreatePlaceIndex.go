package location

// CreatePlaceIndex is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Places API V2 unless you require Grab data.
//
// - CreatePlaceIndex is part of a previous Amazon Location Service Places API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Places API version 2 has a simplified interface that can be used
// without creating or managing place index resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Creates a place index resource in your Amazon Web Services account. Use a place
// index resource to geocode addresses and other text queries by using the
// SearchPlaceIndexForText operation, and reverse geocode coordinates by using the
// SearchPlaceIndexForPosition operation, and enable autosuggestions by using the
// SearchPlaceIndexForSuggestions operation.
//
// If your application is tracking or routing assets you use in your business,
// such as delivery vehicles or employees, you must not use Esri as your
// geolocation provider. See section 82 of the [Amazon Web Services service terms]for more details.
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms
