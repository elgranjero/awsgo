package iotsitewise

// CreateAssetModel is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Creates an asset model from specified property and hierarchy definitions. You
// create assets from asset models. With asset models, you can easily create assets
// of the same type that have standardized definitions. Each asset created from a
// model inherits the asset model's property and hierarchy definitions. For more
// information, see [Defining asset models]in the IoT SiteWise User Guide.
//
// You can create three types of asset models, ASSET_MODEL , COMPONENT_MODEL , or
// an INTERFACE .
//
// - ASSET_MODEL – (default) An asset model that you can use to create assets.
// Can't be included as a component in another asset model.
//
// - COMPONENT_MODEL – A reusable component that you can include in the
// composite models of other asset models. You can't create assets directly from
// this type of asset model.
//
// - INTERFACE – An interface is a type of model that defines a standard
// structure that can be applied to different asset models.
//
// [Defining asset models]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/define-models.html
