package ssm

// PutResourcePolicy is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Creates or updates a Systems Manager resource policy. A resource policy helps
// you to define the IAM entity (for example, an Amazon Web Services account) that
// can manage your Systems Manager resources. The following resources support
// Systems Manager resource policies.
//
// - OpsItemGroup - The resource policy for OpsItemGroup enables Amazon Web
// Services accounts to view and interact with OpsCenter operational work items
// (OpsItems).
//
// - Parameter - The resource policy is used to share a parameter with other
// accounts using Resource Access Manager (RAM).
//
// To share a parameter, it must be in the advanced parameter tier. For
//
// information about parameter tiers, see [Managing parameter tiers]. For information about changing an
// existing standard parameter to an advanced parameter, see [Changing a standard parameter to an advanced parameter].
//
// To share a SecureString parameter, it must be encrypted with a customer managed
//
// key, and you must share the key separately through Key Management Service.
// Amazon Web Services managed keys cannot be shared. Parameters encrypted with the
// default Amazon Web Services managed key can be updated to use a customer managed
// key instead. For KMS key definitions, see [KMS concepts]in the Key Management Service
// Developer Guide.
//
// While you can share a parameter using the Systems Manager PutResourcePolicy
//
// operation, we recommend using Resource Access Manager (RAM) instead. This is
// because using PutResourcePolicy requires the extra step of promoting the
// parameter to a standard RAM Resource Share using the RAM [PromoteResourceShareCreatedFromPolicy]API operation.
// Otherwise, the parameter won't be returned by the Systems Manager [DescribeParameters]API
// operation using the --shared option.
//
// For more information, see [Sharing a parameter]in the Amazon Web Services Systems Manager User Guide
//
// [Sharing a parameter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-shared-parameters.html#share
//
// [Managing parameter tiers]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html
// [Changing a standard parameter to an advanced parameter]: https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html#parameter-store-advanced-parameters-enabling
// [PromoteResourceShareCreatedFromPolicy]: https://docs.aws.amazon.com/ram/latest/APIReference/API_PromoteResourceShareCreatedFromPolicy.html
// [KMS concepts]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html
// [DescribeParameters]: https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DescribeParameters.html
