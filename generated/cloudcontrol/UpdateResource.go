package cloudcontrol

// UpdateResource is generated as a reference stub.
// Executable command wiring lives under cmd/cloudcontrol.go.
//
// Updates the specified property values in the resource.
//
// You specify your resource property updates as a list of patch operations
// contained in a JSON patch document that adheres to the [RFC 6902 - JavaScript Object Notation (JSON) Patch]standard.
//
// For details on how Cloud Control API performs resource update operations, see [Updating a resource]
// in the Amazon Web Services Cloud Control API User Guide.
//
// After you have initiated a resource update request, you can monitor the
// progress of your request by calling [GetResourceRequestStatus]using the RequestToken of the ProgressEvent
// returned by UpdateResource .
//
// For more information about the properties of a specific resource, refer to the
// related topic for the resource in the [Resource and property types reference]in the CloudFormation Users Guide.
//
// [GetResourceRequestStatus]: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
// [RFC 6902 - JavaScript Object Notation (JSON) Patch]: https://datatracker.ietf.org/doc/html/rfc6902
// [Updating a resource]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html
// [Resource and property types reference]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
