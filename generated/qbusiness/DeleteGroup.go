package qbusiness

// DeleteGroup is generated as a reference stub.
// Executable command wiring lives under cmd/qbusiness.go.
//
// Deletes a group so that all users and sub groups that belong to the group can
// no longer access documents only available to that group. For example, after
// deleting the group "Summer Interns", all interns who belonged to that group no
// longer see intern-only documents in their chat results.
//
// If you want to delete, update, or replace users or sub groups of a group, you
// need to use the PutGroup operation. For example, if a user in the group
// "Engineering" leaves the engineering team and another user takes their place,
// you provide an updated list of users or sub groups that belong to the
// "Engineering" group when calling PutGroup .
