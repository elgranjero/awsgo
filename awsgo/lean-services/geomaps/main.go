package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/geomaps"
)

var fields_get_glyphs = []leanruntime.Field{
	{Name: "FontStack", Flag: "font-stack", Type: "*string", Required: true},
	{Name: "FontUnicodeRange", Flag: "font-unicode-range", Type: "*string", Required: true},
}

var fields_get_sprites = []leanruntime.Field{
	{Name: "ColorScheme", Flag: "color-scheme", Type: "types.ColorScheme", Required: true},
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: true},
	{Name: "Style", Flag: "style", Type: "types.MapStyle", Required: true},
	{Name: "Variant", Flag: "variant", Type: "types.Variant", Required: true},
}

var fields_get_static_map = []leanruntime.Field{
	{Name: "BoundedPositions", Flag: "bounded-positions", Type: "*string", Required: false},
	{Name: "BoundingBox", Flag: "bounding-box", Type: "*string", Required: false},
	{Name: "Center", Flag: "center", Type: "*string", Required: false},
	{Name: "ColorScheme", Flag: "color-scheme", Type: "types.ColorScheme", Required: false},
	{Name: "CompactOverlay", Flag: "compact-overlay", Type: "*string", Required: false},
	{Name: "CropLabels", Flag: "crop-labels", Type: "*bool", Required: false},
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: true},
	{Name: "GeoJsonOverlay", Flag: "geo-json-overlay", Type: "*string", Required: false},
	{Name: "Height", Flag: "height", Type: "*int32", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "LabelSize", Flag: "label-size", Type: "types.LabelSize", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "Padding", Flag: "padding", Type: "*int32", Required: false},
	{Name: "PointsOfInterests", Flag: "points-of-interests", Type: "types.MapFeatureMode", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "Radius", Flag: "radius", Type: "*int64", Required: false},
	{Name: "ScaleBarUnit", Flag: "scale-bar-unit", Type: "types.ScaleBarUnit", Required: false},
	{Name: "Style", Flag: "style", Type: "types.StaticMapStyle", Required: false},
	{Name: "Width", Flag: "width", Type: "*int32", Required: true},
	{Name: "Zoom", Flag: "zoom", Type: "*float32", Required: false},
}

var fields_get_style_descriptor = []leanruntime.Field{
	{Name: "Buildings", Flag: "buildings", Type: "types.Buildings", Required: false},
	{Name: "ColorScheme", Flag: "color-scheme", Type: "types.ColorScheme", Required: false},
	{Name: "ContourDensity", Flag: "contour-density", Type: "types.ContourDensity", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "PoliticalView", Flag: "political-view", Type: "*string", Required: false},
	{Name: "Style", Flag: "style", Type: "types.MapStyle", Required: true},
	{Name: "Terrain", Flag: "terrain", Type: "types.Terrain", Required: false},
	{Name: "Traffic", Flag: "traffic", Type: "types.Traffic", Required: false},
	{Name: "TravelModes", Flag: "travel-modes", Type: "[]types.TravelMode", Required: false},
}

var fields_get_tile = []leanruntime.Field{
	{Name: "AdditionalFeatures", Flag: "additional-features", Type: "[]types.TileAdditionalFeature", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Tileset", Flag: "tileset", Type: "*string", Required: true},
	{Name: "X", Flag: "x", Type: "*string", Required: true},
	{Name: "Y", Flag: "y", Type: "*string", Required: true},
	{Name: "Z", Flag: "z", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"get-glyphs": {
			Name:   "get-glyphs",
			Fields: fields_get_glyphs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlyphsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_glyphs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlyphs(ctx, input)
			},
		},
		"get-sprites": {
			Name:   "get-sprites",
			Fields: fields_get_sprites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSpritesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sprites, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSprites(ctx, input)
			},
		},
		"get-static-map": {
			Name:   "get-static-map",
			Fields: fields_get_static_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStaticMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_static_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStaticMap(ctx, input)
			},
		},
		"get-style-descriptor": {
			Name:   "get-style-descriptor",
			Fields: fields_get_style_descriptor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStyleDescriptorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_style_descriptor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStyleDescriptor(ctx, input)
			},
		},
		"get-tile": {
			Name:   "get-tile",
			Fields: fields_get_tile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_tile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTile(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("geomaps", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
