package cloudfront

// ListOriginAccessControls is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Gets the list of CloudFront origin access controls (OACs) in this Amazon Web
// Services account.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send another request that specifies the NextMarker value from the
// current response as the Marker value in the next request.
//
// If you're not using origin access controls for your Amazon Web Services
// account, the ListOriginAccessControls operation doesn't return the Items
// element in the response.
