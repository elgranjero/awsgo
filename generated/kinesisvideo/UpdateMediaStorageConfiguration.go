package kinesisvideo

// UpdateMediaStorageConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// Associates a SignalingChannel to a stream to store the media. There are two
// signaling modes that you can specify :
//
// - If StorageStatus is enabled, the data will be stored in the StreamARN
// provided. In order for WebRTC Ingestion to work, the stream must have data
// retention enabled.
//
// - If StorageStatus is disabled, no data will be stored, and the StreamARN
// parameter will not be needed.
//
// If StorageStatus is enabled, direct peer-to-peer (master-viewer) connections no
// longer occur. Peers connect directly to the storage session. You must call the
// JoinStorageSession API to trigger an SDP offer send and establish a connection
// between a peer and the storage session.
