package kendra

// PutPrincipalMapping is generated as a reference stub.
// Executable command wiring lives under cmd/kendra.go.
//
// Maps users to their groups so that you only need to provide the user ID when
// you issue the query.
//
// You can also map sub groups to groups. For example, the group "Company
// Intellectual Property Teams" includes sub groups "Research" and "Engineering".
// These sub groups include their own list of users or people who work in these
// teams. Only users who work in research and engineering, and therefore belong in
// the intellectual property group, can see top-secret company documents in their
// search results.
//
// This is useful for user context filtering, where search results are filtered
// based on the user or their group access to documents. For more information, see [Filtering on user context]
// .
//
// If more than five PUT actions for a group are currently processing, a
// validation exception is thrown.
//
// [Filtering on user context]: https://docs.aws.amazon.com/kendra/latest/dg/user-context-filter.html
