package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kinesisvideoarchivedmedia/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-clip", "get-dash-streaming-session-url", "get-hls-streaming-session-url", "get-images", "get-media-for-fragment-list", "list-fragments"},
		OperationSet: map[string]bool{"get-clip": true, "get-dash-streaming-session-url": true, "get-hls-streaming-session-url": true, "get-images": true, "get-media-for-fragment-list": true, "list-fragments": true},
		OperationInputs: map[string][]string{
			"get-clip":                       {"ClipFragmentSelector", "StreamARN", "StreamName"},
			"get-dash-streaming-session-url": {"DASHFragmentSelector", "DisplayFragmentNumber", "DisplayFragmentTimestamp", "Expires", "MaxManifestFragmentResults", "PlaybackMode", "StreamARN", "StreamName"},
			"get-hls-streaming-session-url":  {"ContainerFormat", "DiscontinuityMode", "DisplayFragmentTimestamp", "Expires", "HLSFragmentSelector", "MaxMediaPlaylistFragmentResults", "PlaybackMode", "StreamARN", "StreamName"},
			"get-images":                     {"EndTimestamp", "Format", "FormatConfig", "HeightPixels", "ImageSelectorType", "MaxResults", "NextToken", "SamplingInterval", "StartTimestamp", "StreamARN", "StreamName", "WidthPixels"},
			"get-media-for-fragment-list":    {"Fragments", "StreamARN", "StreamName"},
			"list-fragments":                 {"FragmentSelector", "MaxResults", "NextToken", "StreamARN", "StreamName"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-clip":                       {"ClipFragmentSelector": "*types.ClipFragmentSelector", "StreamARN": "*string", "StreamName": "*string"},
			"get-dash-streaming-session-url": {"DASHFragmentSelector": "*types.DASHFragmentSelector", "DisplayFragmentNumber": "types.DASHDisplayFragmentNumber", "DisplayFragmentTimestamp": "types.DASHDisplayFragmentTimestamp", "Expires": "*int32", "MaxManifestFragmentResults": "*int64", "PlaybackMode": "types.DASHPlaybackMode", "StreamARN": "*string", "StreamName": "*string"},
			"get-hls-streaming-session-url":  {"ContainerFormat": "types.ContainerFormat", "DiscontinuityMode": "types.HLSDiscontinuityMode", "DisplayFragmentTimestamp": "types.HLSDisplayFragmentTimestamp", "Expires": "*int32", "HLSFragmentSelector": "*types.HLSFragmentSelector", "MaxMediaPlaylistFragmentResults": "*int64", "PlaybackMode": "types.HLSPlaybackMode", "StreamARN": "*string", "StreamName": "*string"},
			"get-images":                     {"EndTimestamp": "*time.Time", "Format": "types.Format", "FormatConfig": "map[string]string", "HeightPixels": "*int32", "ImageSelectorType": "types.ImageSelectorType", "MaxResults": "*int64", "NextToken": "*string", "SamplingInterval": "*int32", "StartTimestamp": "*time.Time", "StreamARN": "*string", "StreamName": "*string", "WidthPixels": "*int32"},
			"get-media-for-fragment-list":    {"Fragments": "[]string", "StreamARN": "*string", "StreamName": "*string"},
			"list-fragments":                 {"FragmentSelector": "*types.FragmentSelector", "MaxResults": "*int64", "NextToken": "*string", "StreamARN": "*string", "StreamName": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-clip":                       {"ClipFragmentSelector"},
			"get-dash-streaming-session-url": {},
			"get-hls-streaming-session-url":  {},
			"get-images":                     {"EndTimestamp", "Format", "ImageSelectorType", "StartTimestamp"},
			"get-media-for-fragment-list":    {"Fragments"},
			"list-fragments":                 {},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kinesisvideoarchivedmedia", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
