package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
)

var fields_get_clip = []leanruntime.Field{
	{Name: "ClipFragmentSelector", Flag: "clip-fragment-selector", Type: "*types.ClipFragmentSelector", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_get_dash_streaming_session_url = []leanruntime.Field{
	{Name: "DASHFragmentSelector", Flag: "dash-fragment-selector", Type: "*types.DASHFragmentSelector", Required: false},
	{Name: "DisplayFragmentNumber", Flag: "display-fragment-number", Type: "types.DASHDisplayFragmentNumber", Required: false},
	{Name: "DisplayFragmentTimestamp", Flag: "display-fragment-timestamp", Type: "types.DASHDisplayFragmentTimestamp", Required: false},
	{Name: "Expires", Flag: "expires", Type: "*int32", Required: false},
	{Name: "MaxManifestFragmentResults", Flag: "max-manifest-fragment-results", Type: "*int64", Required: false},
	{Name: "PlaybackMode", Flag: "playback-mode", Type: "types.DASHPlaybackMode", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_get_hls_streaming_session_url = []leanruntime.Field{
	{Name: "ContainerFormat", Flag: "container-format", Type: "types.ContainerFormat", Required: false},
	{Name: "DiscontinuityMode", Flag: "discontinuity-mode", Type: "types.HLSDiscontinuityMode", Required: false},
	{Name: "DisplayFragmentTimestamp", Flag: "display-fragment-timestamp", Type: "types.HLSDisplayFragmentTimestamp", Required: false},
	{Name: "Expires", Flag: "expires", Type: "*int32", Required: false},
	{Name: "HLSFragmentSelector", Flag: "hls-fragment-selector", Type: "*types.HLSFragmentSelector", Required: false},
	{Name: "MaxMediaPlaylistFragmentResults", Flag: "max-media-playlist-fragment-results", Type: "*int64", Required: false},
	{Name: "PlaybackMode", Flag: "playback-mode", Type: "types.HLSPlaybackMode", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_get_images = []leanruntime.Field{
	{Name: "EndTimestamp", Flag: "end-timestamp", Type: "*time.Time", Required: true},
	{Name: "Format", Flag: "format", Type: "types.Format", Required: true},
	{Name: "FormatConfig", Flag: "format-config", Type: "map[string]string", Required: false},
	{Name: "HeightPixels", Flag: "height-pixels", Type: "*int32", Required: false},
	{Name: "ImageSelectorType", Flag: "image-selector-type", Type: "types.ImageSelectorType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int64", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SamplingInterval", Flag: "sampling-interval", Type: "*int32", Required: false},
	{Name: "StartTimestamp", Flag: "start-timestamp", Type: "*time.Time", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
	{Name: "WidthPixels", Flag: "width-pixels", Type: "*int32", Required: false},
}

var fields_get_media_for_fragment_list = []leanruntime.Field{
	{Name: "Fragments", Flag: "fragments", Type: "[]string", Required: true},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

var fields_list_fragments = []leanruntime.Field{
	{Name: "FragmentSelector", Flag: "fragment-selector", Type: "*types.FragmentSelector", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int64", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StreamARN", Flag: "stream-arn", Type: "*string", Required: false},
	{Name: "StreamName", Flag: "stream-name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-clip": {
			Name:   "get-clip",
			Fields: fields_get_clip,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetClipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_clip, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetClip(ctx, input)
			},
		},
		"get-dash-streaming-session-url": {
			Name:   "get-dash-streaming-session-url",
			Fields: fields_get_dash_streaming_session_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDASHStreamingSessionURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dash_streaming_session_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDASHStreamingSessionURL(ctx, input)
			},
		},
		"get-hls-streaming-session-url": {
			Name:   "get-hls-streaming-session-url",
			Fields: fields_get_hls_streaming_session_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHLSStreamingSessionURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hls_streaming_session_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHLSStreamingSessionURL(ctx, input)
			},
		},
		"get-images": {
			Name:   "get-images",
			Fields: fields_get_images,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_images, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetImages(ctx, input)
				}
				var results []*svc.GetImagesOutput
				p := svc.NewGetImagesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"get-media-for-fragment-list": {
			Name:   "get-media-for-fragment-list",
			Fields: fields_get_media_for_fragment_list,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMediaForFragmentListInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_media_for_fragment_list, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMediaForFragmentList(ctx, input)
			},
		},
		"list-fragments": {
			Name:   "list-fragments",
			Fields: fields_list_fragments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFragmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fragments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFragments(ctx, input)
				}
				var results []*svc.ListFragmentsOutput
				p := svc.NewListFragmentsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
	}
	if err := leanruntime.Execute("kinesisvideoarchivedmedia", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
