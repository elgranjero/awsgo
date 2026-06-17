package mediapackagev2

// CreateChannel is generated as a reference stub.
// Executable command wiring lives under cmd/mediapackagev2.go.
//
// Create a channel to start receiving content streams. The channel represents the
// input to MediaPackage for incoming live content from an encoder such as AWS
// Elemental MediaLive. The channel receives content, and after packaging it,
// outputs it through an origin endpoint to downstream devices (such as video
// players or CDNs) that request the content. You can create only one channel with
// each request. We recommend that you spread out channels between channel groups,
// such as putting redundant channels in the same AWS Region in different channel
// groups.
