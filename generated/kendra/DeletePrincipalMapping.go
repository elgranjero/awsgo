package kendra

// DeletePrincipalMapping is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Deletes a group so that all users that belong to the group can no longer access
// documents only available to that group.
//
// For example, after deleting the group "Summer Interns", all interns who
// belonged to that group no longer see intern-only documents in their search
// results.
//
// If you want to delete or replace users or sub groups of a group, you need to
// use the PutPrincipalMapping operation. For example, if a user in the group
// "Engineering" leaves the engineering team and another user takes their place,
// you provide an updated list of users or sub groups that belong to the
// "Engineering" group when calling PutPrincipalMapping . You can update your
// internal list of users or sub groups and input this list when calling
// PutPrincipalMapping .
//
// DeletePrincipalMapping is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
