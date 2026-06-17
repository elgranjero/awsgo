package kinesisvideowebrtcstorage

// JoinStorageSessionAsViewer is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideowebrtcstorage.go.
//
// Join the ongoing one way-video and/or multi-way audio WebRTC session as a
//
// viewer for an input channel. If there’s no existing session for the channel,
// create a new streaming session and provide the Amazon Resource Name (ARN) of the
// signaling channel ( channelArn ) and client id ( clientId ).
//
// Currently for SINGLE_MASTER type, a video producing device is able to ingest
// both audio and video media into a stream, while viewers can only ingest audio.
// Both a video producing device and viewers can join a session first and wait for
// other participants. While participants are having peer to peer conversations
// through WebRTC, the ingested media session will be stored into the Kinesis Video
// Stream. Multiple viewers are able to playback real-time media.
//
// Customers can also use existing Kinesis Video Streams features like HLS or DASH
// playback, Image generation, and more with ingested WebRTC media. If there’s an
// existing session with the same clientId that's found in the join session
// request, the new request takes precedence.
