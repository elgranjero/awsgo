package datazone

// CreateAsset is generated as a reference stub.
// Executable command wiring lives under cmd/datazone.go.
//
// Creates an asset in Amazon DataZone catalog.
//
// Before creating assets, make sure that the following requirements are met:
//
// - --domain-identifier must refer to an existing domain.
//
// - --owning-project-identifier must be a valid project within the domain.
//
// - Asset type must be created beforehand using create-asset-type , or be a
// supported system-defined type. For more information, see [create-asset-type].
//
// - --type-revision (if used) must match a valid revision of the asset type.
//
// - formsInput is required when it is associated as required in the asset-type .
// For more information, see [create-form-type].
//
// - Form content must include all required fields as per the form schema (e.g.,
// bucketArn ).
//
// You must invoke the following pre-requisite commands before invoking this API:
//
// [CreateFormType]
//
// [CreateAssetType]
//
// [create-asset-type]: https://docs.aws.amazon.com/cli/latest/reference/datazone/create-asset-type.html
// [create-form-type]: https://docs.aws.amazon.com/cli/latest/reference/datazone/create-form-type.html
// [CreateFormType]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateFormType.html
// [CreateAssetType]: https://docs.aws.amazon.com/datazone/latest/APIReference/API_CreateAssetType.html
