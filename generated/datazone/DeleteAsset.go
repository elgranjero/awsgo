package datazone

// DeleteAsset is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Deletes an asset in Amazon DataZone.
//
// - --domain-identifier must refer to a valid and existing domain.
//
// - --identifier must refer to an existing asset in the specified domain.
//
// - Asset must not be referenced in any existing asset filters.
//
// - Asset must not be linked to any draft or published data product.
//
// - User must have delete permissions for the domain and project.
