package efs

// PutLifecycleConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/efs.go.
//
// Use this action to manage storage for your file system. A LifecycleConfiguration
// consists of one or more LifecyclePolicy objects that define the following:
//
// - TransitionToIA – When to move files in the file system from primary storage
// (Standard storage class) into the Infrequent Access (IA) storage.
//
// - TransitionToArchive – When to move files in the file system from their
// current storage class (either IA or Standard storage) into the Archive storage.
//
// File systems cannot transition into Archive storage before transitioning into
//
// IA storage. Therefore, TransitionToArchive must either not be set or must be
// later than TransitionToIA.
//
// The Archive storage class is available only for file systems that use the
//
// Elastic throughput mode and the General Purpose performance mode.
//
// - TransitionToPrimaryStorageClass – Whether to move files in the file system
// back to primary storage (Standard storage class) after they are accessed in IA
// or Archive storage.
//
// For more information, see [Managing file system storage].
//
// Each Amazon EFS file system supports one lifecycle configuration, which applies
// to all files in the file system. If a LifecycleConfiguration object already
// exists for the specified file system, a PutLifecycleConfiguration call modifies
// the existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration . In the request, specify the following:
//
// - The ID for the file system for which you are enabling, disabling, or
// modifying lifecycle management.
//
// - A LifecyclePolicies array of LifecyclePolicy objects that define when to
// move files to IA storage, to Archive storage, and back to primary storage.
//
// Amazon EFS requires that each LifecyclePolicy object have only have a single
//
// transition, so the LifecyclePolicies array needs to be structured with
// separate LifecyclePolicy objects. See the example requests in the following
// section for more information.
//
// This operation requires permissions for the
// elasticfilesystem:PutLifecycleConfiguration operation.
//
// To apply a LifecycleConfiguration object to an encrypted file system, you need
// the same Key Management Service permissions as when you created the encrypted
// file system.
//
// [Managing file system storage]: https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html
