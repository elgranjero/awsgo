package synthetics

// DeleteGroup is generated as a reference stub.
// Executable command wiring lives under cmd/synthetics.go.
//
// Deletes a group. The group doesn't need to be empty to be deleted. If there are
// canaries in the group, they are not deleted when you delete the group.
//
// Groups are a global resource that appear in all Regions, but the request to
// delete a group must be made from its home Region. You can find the home Region
// of a group within its ARN.
