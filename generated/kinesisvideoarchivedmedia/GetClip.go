package kinesisvideoarchivedmedia

// GetClip is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideoarchivedmedia.go.
//
// Downloads an MP4 file (clip) containing the archived, on-demand media from the
// specified video stream over the specified time range.
//
// Both the StreamName and the StreamARN parameters are optional, but you must
// specify either the StreamName or the StreamARN when invoking this API operation.
//
// As a prerequisite to using GetCLip API, you must obtain an endpoint using
// GetDataEndpoint , specifying GET_CLIP for the APIName parameter.
//
// An Amazon Kinesis video stream has the following requirements for providing
// data through MP4:
//
// - The media must contain h.264 or h.265 encoded video and, optionally, AAC or
// G.711 encoded audio. Specifically, the codec ID of track 1 should be
// V_MPEG/ISO/AVC (for h.264) or V_MPEGH/ISO/HEVC (for H.265). Optionally, the
// codec ID of track 2 should be A_AAC (for AAC) or A_MS/ACM (for G.711).
//
// - Data retention must be greater than 0.
//
// - The video track of each fragment must contain codec private data in the
// Advanced Video Coding (AVC) for H.264 format and HEVC for H.265 format. For more
// information, see [MPEG-4 specification ISO/IEC 14496-15]. For information about adapting stream data to a given
// format, see [NAL Adaptation Flags].
//
// - The audio track (if present) of each fragment must contain codec private
// data in the AAC format ([AAC specification ISO/IEC 13818-7] ) or the [MS Wave format].
//
// You can monitor the amount of outgoing data by monitoring the
// GetClip.OutgoingBytes Amazon CloudWatch metric. For information about using
// CloudWatch to monitor Kinesis Video Streams, see [Monitoring Kinesis Video Streams]. For pricing information, see [Amazon Kinesis Video Streams Pricing]
// and [Amazon Web Services Pricing]. Charges for outgoing Amazon Web Services data apply.
//
// [Amazon Web Services Pricing]: https://aws.amazon.com/pricing/
// [MPEG-4 specification ISO/IEC 14496-15]: https://www.iso.org/standard/55980.html
// [Amazon Kinesis Video Streams Pricing]: https://aws.amazon.com/kinesis/video-streams/pricing/
// [Monitoring Kinesis Video Streams]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/monitoring.html
// [AAC specification ISO/IEC 13818-7]: https://www.iso.org/standard/43345.html
// [NAL Adaptation Flags]: http://docs.aws.amazon.com/kinesisvideostreams/latest/dg/producer-reference-nal.html
// [MS Wave format]: http://www-mmsp.ece.mcgill.ca/Documents/AudioFormats/WAVE/WAVE.html
