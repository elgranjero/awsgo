package amp

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/amp.go.
//
// Creates or updates a resource-based policy for an Amazon Managed Service for
// Prometheus workspace. Use resource-based policies to grant permissions to other
// AWS accounts or services to access your workspace.
//
// Only Prometheus-compatible APIs can be used for workspace sharing. You can add
// non-Prometheus-compatible APIs to the policy, but they will be ignored. For more
// information, see [Prometheus-compatible APIs]in the Amazon Managed Service for Prometheus User Guide.
//
// If your workspace uses customer-managed KMS keys for encryption, you must grant
// the principals in your resource-based policy access to those KMS keys. You can
// do this by creating KMS grants. For more information, see [CreateGrant]in the AWS Key
// Management Service API Reference and [Encryption at rest]in the Amazon Managed Service for
// Prometheus User Guide.
//
// For more information about working with IAM, see [Using Amazon Managed Service for Prometheus with IAM] in the Amazon Managed Service
// for Prometheus User Guide.
//
// [Prometheus-compatible APIs]: https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-APIReference-Prometheus-Compatible-Apis.html
// [Using Amazon Managed Service for Prometheus with IAM]: https://docs.aws.amazon.com/prometheus/latest/userguide/security_iam_service-with-iam.html
// [Encryption at rest]: https://docs.aws.amazon.com/prometheus/latest/userguide/encryption-at-rest-Amazon-Service-Prometheus.html
// [CreateGrant]: https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateGrant.html
