package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/geomaps/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"get-glyphs", "get-sprites", "get-static-map", "get-style-descriptor", "get-tile"},
		OperationSet: map[string]bool{"get-glyphs": true, "get-sprites": true, "get-static-map": true, "get-style-descriptor": true, "get-tile": true},
		OperationInputs: map[string][]string{
			"get-glyphs":           {"FontStack", "FontUnicodeRange"},
			"get-sprites":          {"ColorScheme", "FileName", "Style", "Variant"},
			"get-static-map":       {"BoundedPositions", "BoundingBox", "Center", "ColorScheme", "CompactOverlay", "CropLabels", "FileName", "GeoJsonOverlay", "Height", "Key", "LabelSize", "Language", "Padding", "PointsOfInterests", "PoliticalView", "Radius", "ScaleBarUnit", "Style", "Width", "Zoom"},
			"get-style-descriptor": {"Buildings", "ColorScheme", "ContourDensity", "Key", "PoliticalView", "Style", "Terrain", "Traffic", "TravelModes"},
			"get-tile":             {"AdditionalFeatures", "Key", "Tileset", "X", "Y", "Z"},
		},
		OperationInputTypes: map[string]map[string]string{
			"get-glyphs":           {"FontStack": "*string", "FontUnicodeRange": "*string"},
			"get-sprites":          {"ColorScheme": "types.ColorScheme", "FileName": "*string", "Style": "types.MapStyle", "Variant": "types.Variant"},
			"get-static-map":       {"BoundedPositions": "*string", "BoundingBox": "*string", "Center": "*string", "ColorScheme": "types.ColorScheme", "CompactOverlay": "*string", "CropLabels": "*bool", "FileName": "*string", "GeoJsonOverlay": "*string", "Height": "*int32", "Key": "*string", "LabelSize": "types.LabelSize", "Language": "*string", "Padding": "*int32", "PointsOfInterests": "types.MapFeatureMode", "PoliticalView": "*string", "Radius": "*int64", "ScaleBarUnit": "types.ScaleBarUnit", "Style": "types.StaticMapStyle", "Width": "*int32", "Zoom": "*float32"},
			"get-style-descriptor": {"Buildings": "types.Buildings", "ColorScheme": "types.ColorScheme", "ContourDensity": "types.ContourDensity", "Key": "*string", "PoliticalView": "*string", "Style": "types.MapStyle", "Terrain": "types.Terrain", "Traffic": "types.Traffic", "TravelModes": "[]types.TravelMode"},
			"get-tile":             {"AdditionalFeatures": "[]types.TileAdditionalFeature", "Key": "*string", "Tileset": "*string", "X": "*string", "Y": "*string", "Z": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"get-glyphs":           {"FontStack", "FontUnicodeRange"},
			"get-sprites":          {"ColorScheme", "FileName", "Style", "Variant"},
			"get-static-map":       {"FileName", "Height", "Width"},
			"get-style-descriptor": {"Style"},
			"get-tile":             {"Tileset", "X", "Y", "Z"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("geomaps", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
