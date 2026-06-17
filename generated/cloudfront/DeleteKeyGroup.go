package cloudfront

// DeleteKeyGroup is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Deletes a key group.
//
// You cannot delete a key group that is referenced in a cache behavior. First
// update your distributions to remove the key group from all cache behaviors, then
// delete the key group.
//
// To delete a key group, you must provide the key group's identifier and version.
// To get these values, use ListKeyGroups followed by GetKeyGroup or
// GetKeyGroupConfig .
