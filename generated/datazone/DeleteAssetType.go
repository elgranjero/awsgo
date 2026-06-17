package datazone

// DeleteAssetType is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Deletes an asset type in Amazon DataZone.
//
// Prerequisites:
//
// - The asset type must exist in the domain.
//
// - You must have DeleteAssetType permission.
//
// - The asset type must not be in use (e.g., assigned to any asset). If used,
// deletion will fail.
//
// - You should retrieve the asset type using get-asset-type to confirm its
// presence before deletion.
