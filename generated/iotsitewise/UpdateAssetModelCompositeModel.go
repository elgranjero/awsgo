package iotsitewise

// UpdateAssetModelCompositeModel is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Updates a composite model and all of the assets that were created from the
// model. Each asset created from the model inherits the updated asset model's
// property and hierarchy definitions. For more information, see [Updating assets and models]in the IoT
// SiteWise User Guide.
//
// If you remove a property from a composite asset model, IoT SiteWise deletes all
// previous data for that property. You can’t change the type or data type of an
// existing property.
//
// To replace an existing composite asset model property with a new one with the
// same name , do the following:
//
// - Submit an UpdateAssetModelCompositeModel request with the entire existing
// property removed.
//
// - Submit a second UpdateAssetModelCompositeModel request that includes the new
// property. The new asset property will have the same name as the previous one
// and IoT SiteWise will generate a new unique id .
//
// [Updating assets and models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html
