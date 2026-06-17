package location

// GetPlace is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the V2 GetPlaceGetPlace operation unless you require Grab data.
//
// - This version of GetPlace is part of a previous Amazon Location Service
// Places API (version 1) which has been superseded by a more intuitive, powerful,
// and complete API (version 2).
//
// - Version 2 of the GetPlace operation interoperates with the rest of the
// Places V2 API, while this version does not.
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
// Finds a place by its unique ID. A PlaceId is returned by other search
// operations.
//
// A PlaceId is valid only if all of the following are the same in the original
// search request and the call to GetPlace .
//
// - Customer Amazon Web Services account
//
// - Amazon Web Services Region
//
// - Data provider specified in the place index resource
//
// If your Place index resource is configured with Grab as your geolocation
// provider and Storage as Intended use, the GetPlace operation is unavailable. For
// more information, see [AWS service terms].
//
// [AWS service terms]: http://aws.amazon.com/service-terms
