package storagegateway

// ListVolumes is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Lists the iSCSI stored volumes of a gateway. Results are sorted by volume ARN.
// The response includes only the volume ARNs. If you want additional volume
// information, use the DescribeStorediSCSIVolumesor the DescribeCachediSCSIVolumes API.
//
// The operation supports pagination. By default, the operation returns a maximum
// of up to 100 volumes. You can optionally specify the Limit field in the body to
// limit the number of volumes in the response. If the number of volumes returned
// in the response is truncated, the response includes a Marker field. You can use
// this Marker value in your subsequent request to retrieve the next set of
// volumes. This operation is only supported in the cached volume and stored volume
// gateway types.
