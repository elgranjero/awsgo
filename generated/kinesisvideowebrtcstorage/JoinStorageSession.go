package kinesisvideowebrtcstorage

// JoinStorageSession is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideowebrtcstorage.go.
//
// Before using this API, you must call the GetSignalingChannelEndpoint API to
// request the WEBRTC endpoint. You then specify the endpoint and region in your
// JoinStorageSession API request.
//
// Join the ongoing one way-video and/or multi-way audio WebRTC session as a video
// producing device for an input channel. If there’s no existing session for the
// channel, a new streaming session needs to be created, and the Amazon Resource
// Name (ARN) of the signaling channel must be provided.
//
// Currently for the SINGLE_MASTER type, a video producing device is able to
// ingest both audio and video media into a stream. Only video producing devices
// can join the session and record media.
//
// Both audio and video tracks are currently required for WebRTC ingestion.
//
// Current requirements:
//
// - Video track: H.264
//
// - Audio track: Opus
//
// The resulting ingested video in the Kinesis video stream will have the
// following parameters: H.264 video and AAC audio.
//
// Once a master participant has negotiated a connection through WebRTC, the
// ingested media session will be stored in the Kinesis video stream. Multiple
// viewers are then able to play back real-time media through our Playback APIs.
//
// You can also use existing Kinesis Video Streams features like HLS or DASH
// playback, image generation via [GetImages], and more with ingested WebRTC media.
//
// S3 image delivery and notifications are not currently supported.
//
// Assume that only one video producing device client can be associated with a
// session for the channel. If more than one client joins the session of a specific
// channel as a video producing device, the most recent client request takes
// precedence.
//
// Additional information
//
// - Idempotent - This API is not idempotent.
//
// - Retry behavior - This is counted as a new API call.
//
// - Concurrent calls - Concurrent calls are allowed. An offer is sent once per
// each call.
//
// [GetImages]: https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/gs-getImages.html
