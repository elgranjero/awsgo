package iotsitewise

// CreateAssetModelCompositeModel is generated as a reference stub.
// Executable command wiring lives under cmd/iotsitewise.go.
//
// Creates a custom composite model from specified property and hierarchy
// definitions. There are two types of custom composite models, inline and
// component-model-based .
//
// Use component-model-based custom composite models to define standard, reusable
// components. A component-model-based custom composite model consists of a name, a
// description, and the ID of the component model it references. A
// component-model-based custom composite model has no properties of its own; its
// referenced component model provides its associated properties to any created
// assets. For more information, see [Custom composite models (Components)]in the IoT SiteWise User Guide.
//
// Use inline custom composite models to organize the properties of an asset
// model. The properties of inline custom composite models are local to the asset
// model where they are included and can't be used to create multiple assets.
//
// To create a component-model-based model, specify the composedAssetModelId of an
// existing asset model with assetModelType of COMPONENT_MODEL .
//
// To create an inline model, specify the assetModelCompositeModelProperties and
// don't include an composedAssetModelId .
//
// [Custom composite models (Components)]: https://docs.aws.amazon.com/iot-sitewise/latest/userguide/custom-composite-models.html
