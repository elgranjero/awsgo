package eks

// DeleteFargateProfile is generated as a reference stub.
// Executable command wiring lives under cmd/eks.go.
//
// Deletes an Fargate profile.
//
// When you delete a Fargate profile, any Pod running on Fargate that was created
// with the profile is deleted. If the Pod matches another Fargate profile, then
// it is scheduled on Fargate with that profile. If it no longer matches any
// Fargate profiles, then it's not scheduled on Fargate and may remain in a pending
// state.
//
// Only one Fargate profile in a cluster can be in the DELETING status at a time.
// You must wait for a Fargate profile to finish deleting before you can delete any
// other profiles in that cluster.
