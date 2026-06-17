package codeartifact

// DeletePackageGroup is generated as a reference stub.
// Executable command wiring lives under cmd/codeartifact.go.
//
// Deletes a package group. Deleting a package group does not delete packages or
// package versions associated with the package group. When a package group is
// deleted, the direct child package groups will become children of the package
// group's direct parent package group. Therefore, if any of the child groups are
// inheriting any settings from the parent, those settings could change.
