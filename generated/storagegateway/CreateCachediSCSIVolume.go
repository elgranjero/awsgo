package storagegateway

// CreateCachediSCSIVolume is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Creates a cached volume on a specified cached volume gateway. This operation is
// only supported in the cached volume gateway type.
//
// Cache storage must be allocated to the gateway before you can create a cached
// volume. Use the AddCacheoperation to add cache storage to a gateway.
//
// In the request, you must specify the gateway, size of the volume in bytes, the
// iSCSI target name, an IP address on which to expose the target, and a unique
// client token. In response, the gateway creates the volume and returns
// information about it. This information includes the volume Amazon Resource Name
// (ARN), its size, and the iSCSI target ARN that initiators can use to connect to
// the volume target.
//
// Optionally, you can provide the ARN for an existing volume as the
// SourceVolumeARN for this cached volume, which creates an exact copy of the
// existing volume’s latest recovery point. The VolumeSizeInBytes value must be
// equal to or larger than the size of the copied volume, in bytes.
