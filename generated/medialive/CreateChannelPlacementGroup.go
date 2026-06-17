package medialive

// CreateChannelPlacementGroup is generated as a reference stub.
// Executable command wiring lives under cmd/medialive.go.
//
// Create a ChannelPlacementGroup in the specified Cluster. As part of the create
// operation, you specify the Nodes to attach the group to.After you create a
// ChannelPlacementGroup, you add Channels to the group (you do this by modifying
// the Channels to add them to a specific group). You now have an association of
// Channels to ChannelPlacementGroup, and ChannelPlacementGroup to Nodes. This
// association means that all the Channels in the group are able to run on any of
// the Nodes associated with the group.
