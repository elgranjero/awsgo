package qbusiness

// PutGroup is generated as a reference stub.
// Executable command wiring lives under cmd/qbusiness.go.
//
// Create, or updates, a mapping of users—who have access to a document—to groups.
//
// You can also map sub groups to groups. For example, the group "Company
// Intellectual Property Teams" includes sub groups "Research" and "Engineering".
// These sub groups include their own list of users or people who work in these
// teams. Only users who work in research and engineering, and therefore belong in
// the intellectual property group, can see top-secret company documents in their
// Amazon Q Business chat results.
//
// There are two options for creating groups, either passing group members inline
// or using an S3 file via the S3PathForGroupMembers field. For inline groups,
// there is a limit of 1000 members per group and for provided S3 files there is a
// limit of 100 thousand members. When creating a group using an S3 file, you
// provide both an S3 file and a RoleArn for Amazon Q Buisness to access the file.
