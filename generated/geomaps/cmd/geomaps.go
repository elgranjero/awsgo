package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/geomaps"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// geomapsCmd represents the geomaps command
var _geomapsCmd = &cobra.Command{
	Use:   "geomaps",
	Short: "AWS geomaps CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := geomaps.NewFromConfig(cfg)
		if _geomapsGetGlyphs {
			geomaps_GetGlyphs(cfg, client)
			return
		}
		if _geomapsGetSprites {
			geomaps_GetSprites(cfg, client)
			return
		}
		if _geomapsGetStaticMap {
			geomaps_GetStaticMap(cfg, client)
			return
		}
		if _geomapsGetStyleDescriptor {
			geomaps_GetStyleDescriptor(cfg, client)
			return
		}
		if _geomapsGetTile {
			geomaps_GetTile(cfg, client)
			return
		}

	},
}

var (
	_geomapsGetGlyphs          bool
	_geomapsGetSprites         bool
	_geomapsGetStaticMap       bool
	_geomapsGetStyleDescriptor bool
	_geomapsGetTile            bool

	_geomapsAdditionalFeatures string
	_geomapsBoundedPositions   string
	_geomapsBoundingBox        string
	_geomapsBuildings          string
	_geomapsCenter             string
	_geomapsColorScheme        string
	_geomapsCompactOverlay     string
	_geomapsContourDensity     string
	_geomapsCropLabels         string
	_geomapsFileName           string
	_geomapsFontStack          string
	_geomapsFontUnicodeRange   string
	_geomapsGeoJsonOverlay     string
	_geomapsHeight             string
	_geomapsKey                string
	_geomapsLabelSize          string
	_geomapsLanguage           string
	_geomapsPadding            string
	_geomapsPointsOfInterests  string
	_geomapsPoliticalView      string
	_geomapsRadius             string
	_geomapsScaleBarUnit       string
	_geomapsStyle              string
	_geomapsTerrain            string
	_geomapsTileset            string
	_geomapsTraffic            string
	_geomapsTravelModes        string
	_geomapsVariant            string
	_geomapsWidth              string
	_geomapsX                  string
	_geomapsY                  string
	_geomapsZ                  string
	_geomapsZoom               string
)

// GetGlyphs returns the map's glyphs.
// For more information, see [Style labels with glyphs] in the Amazon Location Service Developer Guide.
//
// [Style labels with glyphs]: https://docs.aws.amazon.com/location/latest/developerguide/styling-labels-with-glyphs.html
func geomaps_GetGlyphs(cfg aws.Config, client *geomaps.Client) {
	input := &geomaps.GetGlyphsInput{
		// FontStack: *string, // Required
		// FontUnicodeRange: *string, // Required
	}

	if len(_geomapsFontStack) > 0 {
		input.FontStack = aws.String(_geomapsFontStack)
	}
	if len(_geomapsFontUnicodeRange) > 0 {
		input.FontUnicodeRange = aws.String(_geomapsFontUnicodeRange)
	}

	if resp, err := client.GetGlyphs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// GetSprites returns the map's sprites.
// For more information, see [Style iconography with sprites] in the Amazon Location Service Developer Guide.
//
// [Style iconography with sprites]: https://docs.aws.amazon.com/location/latest/developerguide/styling-iconography-with-sprites.html
func geomaps_GetSprites(cfg aws.Config, client *geomaps.Client) {
	input := &geomaps.GetSpritesInput{
		// ColorScheme: types.ColorScheme, // Required
		// FileName: *string, // Required
		// Style: types.MapStyle, // Required
		// Variant: types.Variant, // Required
	}

	if len(_geomapsColorScheme) > 0 {
		if err := assignInputField(input, "ColorScheme", _geomapsColorScheme); err != nil {
			log.Errorf("invalid --color-scheme: %s", err.Error())
			return
		}
	}
	if len(_geomapsFileName) > 0 {
		input.FileName = aws.String(_geomapsFileName)
	}
	if len(_geomapsStyle) > 0 {
		if err := assignInputField(input, "Style", _geomapsStyle); err != nil {
			log.Errorf("invalid --style: %s", err.Error())
			return
		}
	}
	if len(_geomapsVariant) > 0 {
		if err := assignInputField(input, "Variant", _geomapsVariant); err != nil {
			log.Errorf("invalid --variant: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSprites(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// GetStaticMap provides high-quality static map images with customizable options.
// You can modify the map's appearance and overlay additional information. It's an
// ideal solution for applications requiring tailored static map snapshots.
//
// For more information, see the following topics in the Amazon Location Service
// Developer Guide:
//
// [Static maps]
//
// [Customize static maps]
//
// [Overlay on the static map]
//
// [Overlay on the static map]: https://docs.aws.amazon.com/location/latest/developerguide/overlaying-static-map.html
// [Customize static maps]: https://docs.aws.amazon.com/location/latest/developerguide/customizing-static-maps.html
// [Static maps]: https://docs.aws.amazon.com/location/latest/developerguide/static-maps.html
func geomaps_GetStaticMap(cfg aws.Config, client *geomaps.Client) {
	input := &geomaps.GetStaticMapInput{
		// FileName: *string, // Required
		// Height: *int32, // Required
		// Width: *int32, // Required
	}

	if len(_geomapsFileName) > 0 {
		input.FileName = aws.String(_geomapsFileName)
	}
	if len(_geomapsHeight) > 0 {
		if err := assignInputField(input, "Height", _geomapsHeight); err != nil {
			log.Errorf("invalid --height: %s", err.Error())
			return
		}
	}
	if len(_geomapsWidth) > 0 {
		if err := assignInputField(input, "Width", _geomapsWidth); err != nil {
			log.Errorf("invalid --width: %s", err.Error())
			return
		}
	}
	if len(_geomapsBoundedPositions) > 0 {
		input.BoundedPositions = aws.String(_geomapsBoundedPositions)
	}
	if len(_geomapsBoundingBox) > 0 {
		input.BoundingBox = aws.String(_geomapsBoundingBox)
	}
	if len(_geomapsCenter) > 0 {
		input.Center = aws.String(_geomapsCenter)
	}
	if len(_geomapsColorScheme) > 0 {
		if err := assignInputField(input, "ColorScheme", _geomapsColorScheme); err != nil {
			log.Errorf("invalid --color-scheme: %s", err.Error())
			return
		}
	}
	if len(_geomapsCompactOverlay) > 0 {
		input.CompactOverlay = aws.String(_geomapsCompactOverlay)
	}
	if len(_geomapsCropLabels) > 0 {
		if err := assignInputField(input, "CropLabels", _geomapsCropLabels); err != nil {
			log.Errorf("invalid --crop-labels: %s", err.Error())
			return
		}
	}
	if len(_geomapsGeoJsonOverlay) > 0 {
		input.GeoJsonOverlay = aws.String(_geomapsGeoJsonOverlay)
	}
	if len(_geomapsKey) > 0 {
		input.Key = aws.String(_geomapsKey)
	}
	if len(_geomapsLabelSize) > 0 {
		if err := assignInputField(input, "LabelSize", _geomapsLabelSize); err != nil {
			log.Errorf("invalid --label-size: %s", err.Error())
			return
		}
	}
	if len(_geomapsLanguage) > 0 {
		input.Language = aws.String(_geomapsLanguage)
	}
	if len(_geomapsPadding) > 0 {
		if err := assignInputField(input, "Padding", _geomapsPadding); err != nil {
			log.Errorf("invalid --padding: %s", err.Error())
			return
		}
	}
	if len(_geomapsPointsOfInterests) > 0 {
		if err := assignInputField(input, "PointsOfInterests", _geomapsPointsOfInterests); err != nil {
			log.Errorf("invalid --points-of-interests: %s", err.Error())
			return
		}
	}
	if len(_geomapsPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geomapsPoliticalView)
	}
	if len(_geomapsRadius) > 0 {
		if err := assignInputField(input, "Radius", _geomapsRadius); err != nil {
			log.Errorf("invalid --radius: %s", err.Error())
			return
		}
	}
	if len(_geomapsScaleBarUnit) > 0 {
		if err := assignInputField(input, "ScaleBarUnit", _geomapsScaleBarUnit); err != nil {
			log.Errorf("invalid --scale-bar-unit: %s", err.Error())
			return
		}
	}
	if len(_geomapsStyle) > 0 {
		if err := assignInputField(input, "Style", _geomapsStyle); err != nil {
			log.Errorf("invalid --style: %s", err.Error())
			return
		}
	}
	if len(_geomapsZoom) > 0 {
		if err := assignInputField(input, "Zoom", _geomapsZoom); err != nil {
			log.Errorf("invalid --zoom: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetStaticMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// GetStyleDescriptor returns information about the style.
// For more information, see [Style dynamic maps] in the Amazon Location Service Developer Guide.
//
// [Style dynamic maps]: https://docs.aws.amazon.com/location/latest/developerguide/styling-dynamic-maps.html
func geomaps_GetStyleDescriptor(cfg aws.Config, client *geomaps.Client) {
	input := &geomaps.GetStyleDescriptorInput{
		// Style: types.MapStyle, // Required
	}

	if len(_geomapsStyle) > 0 {
		if err := assignInputField(input, "Style", _geomapsStyle); err != nil {
			log.Errorf("invalid --style: %s", err.Error())
			return
		}
	}
	if len(_geomapsBuildings) > 0 {
		if err := assignInputField(input, "Buildings", _geomapsBuildings); err != nil {
			log.Errorf("invalid --buildings: %s", err.Error())
			return
		}
	}
	if len(_geomapsColorScheme) > 0 {
		if err := assignInputField(input, "ColorScheme", _geomapsColorScheme); err != nil {
			log.Errorf("invalid --color-scheme: %s", err.Error())
			return
		}
	}
	if len(_geomapsContourDensity) > 0 {
		if err := assignInputField(input, "ContourDensity", _geomapsContourDensity); err != nil {
			log.Errorf("invalid --contour-density: %s", err.Error())
			return
		}
	}
	if len(_geomapsKey) > 0 {
		input.Key = aws.String(_geomapsKey)
	}
	if len(_geomapsPoliticalView) > 0 {
		input.PoliticalView = aws.String(_geomapsPoliticalView)
	}
	if len(_geomapsTerrain) > 0 {
		if err := assignInputField(input, "Terrain", _geomapsTerrain); err != nil {
			log.Errorf("invalid --terrain: %s", err.Error())
			return
		}
	}
	if len(_geomapsTraffic) > 0 {
		if err := assignInputField(input, "Traffic", _geomapsTraffic); err != nil {
			log.Errorf("invalid --traffic: %s", err.Error())
			return
		}
	}
	if len(_geomapsTravelModes) > 0 {
		if err := assignInputField(input, "TravelModes", _geomapsTravelModes); err != nil {
			log.Errorf("invalid --travel-modes: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetStyleDescriptor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// GetTile returns a tile. Map tiles are used by clients to render a map. they're
// addressed using a grid arrangement with an X coordinate, Y coordinate, and Z
// (zoom) level.
//
// For more information, see [Tiles] in the Amazon Location Service Developer Guide.
//
// [Tiles]: https://docs.aws.amazon.com/location/latest/developerguide/tiles.html
func geomaps_GetTile(cfg aws.Config, client *geomaps.Client) {
	input := &geomaps.GetTileInput{
		// Tileset: *string, // Required
		// X: *string, // Required
		// Y: *string, // Required
		// Z: *string, // Required
	}

	if len(_geomapsTileset) > 0 {
		input.Tileset = aws.String(_geomapsTileset)
	}
	if len(_geomapsX) > 0 {
		input.X = aws.String(_geomapsX)
	}
	if len(_geomapsY) > 0 {
		input.Y = aws.String(_geomapsY)
	}
	if len(_geomapsZ) > 0 {
		input.Z = aws.String(_geomapsZ)
	}
	if len(_geomapsAdditionalFeatures) > 0 {
		if err := assignInputField(input, "AdditionalFeatures", _geomapsAdditionalFeatures); err != nil {
			log.Errorf("invalid --additional-features: %s", err.Error())
			return
		}
	}
	if len(_geomapsKey) > 0 {
		input.Key = aws.String(_geomapsKey)
	}

	if resp, err := client.GetTile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_geomapsCmd)
	_geomapsCmd.Flags().SortFlags = false

	_geomapsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_geomapsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_geomapsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_geomapsCmd.Flags().StringVarP(&_geomapsAdditionalFeatures, "additional-features", "", "", "Additional Features")
	_geomapsCmd.Flags().StringVarP(&_geomapsBoundedPositions, "bounded-positions", "", "", "Bounded Positions")
	_geomapsCmd.Flags().StringVarP(&_geomapsBoundingBox, "bounding-box", "", "", "Bounding Box")
	_geomapsCmd.Flags().StringVarP(&_geomapsBuildings, "buildings", "", "", "Buildings")
	_geomapsCmd.Flags().StringVarP(&_geomapsCenter, "center", "", "", "Center")
	_geomapsCmd.Flags().StringVarP(&_geomapsColorScheme, "color-scheme", "", "", "Color Scheme")
	_geomapsCmd.Flags().StringVarP(&_geomapsCompactOverlay, "compact-overlay", "", "", "Compact Overlay")
	_geomapsCmd.Flags().StringVarP(&_geomapsContourDensity, "contour-density", "", "", "Contour Density")
	_geomapsCmd.Flags().StringVarP(&_geomapsCropLabels, "crop-labels", "", "", "Crop Labels")
	_geomapsCmd.Flags().StringVarP(&_geomapsFileName, "file-name", "", "", "File Name")
	_geomapsCmd.Flags().StringVarP(&_geomapsFontStack, "font-stack", "", "", "Font Stack")
	_geomapsCmd.Flags().StringVarP(&_geomapsFontUnicodeRange, "font-unicode-range", "", "", "Font Unicode Range")
	_geomapsCmd.Flags().StringVarP(&_geomapsGeoJsonOverlay, "geo-json-overlay", "", "", "Geo JSON Overlay")
	_geomapsCmd.Flags().StringVarP(&_geomapsHeight, "height", "", "", "Height")
	_geomapsCmd.Flags().StringVarP(&_geomapsKey, "key", "", "", "Key")
	_geomapsCmd.Flags().StringVarP(&_geomapsLabelSize, "label-size", "", "", "Label Size")
	_geomapsCmd.Flags().StringVarP(&_geomapsLanguage, "language", "", "", "Language")
	_geomapsCmd.Flags().StringVarP(&_geomapsPadding, "padding", "", "", "Padding")
	_geomapsCmd.Flags().StringVarP(&_geomapsPointsOfInterests, "points-of-interests", "", "", "Points Of Interests")
	_geomapsCmd.Flags().StringVarP(&_geomapsPoliticalView, "political-view", "", "", "Political View")
	_geomapsCmd.Flags().StringVarP(&_geomapsRadius, "radius", "", "", "Radius")
	_geomapsCmd.Flags().StringVarP(&_geomapsScaleBarUnit, "scale-bar-unit", "", "", "Scale Bar Unit")
	_geomapsCmd.Flags().StringVarP(&_geomapsStyle, "style", "", "", "Style")
	_geomapsCmd.Flags().StringVarP(&_geomapsTerrain, "terrain", "", "", "Terrain")
	_geomapsCmd.Flags().StringVarP(&_geomapsTileset, "tileset", "", "", "Tileset")
	_geomapsCmd.Flags().StringVarP(&_geomapsTraffic, "traffic", "", "", "Traffic")
	_geomapsCmd.Flags().StringVarP(&_geomapsTravelModes, "travel-modes", "", "", "Travel Modes")
	_geomapsCmd.Flags().StringVarP(&_geomapsVariant, "variant", "", "", "Variant")
	_geomapsCmd.Flags().StringVarP(&_geomapsWidth, "width", "", "", "Width")
	_geomapsCmd.Flags().StringVarP(&_geomapsX, "x", "", "", "X")
	_geomapsCmd.Flags().StringVarP(&_geomapsY, "y", "", "", "Y")
	_geomapsCmd.Flags().StringVarP(&_geomapsZ, "z", "", "", "Z")
	_geomapsCmd.Flags().StringVarP(&_geomapsZoom, "zoom", "", "", "Zoom")

	_geomapsCmd.Flags().BoolVarP(&_geomapsGetGlyphs, "get-glyphs", "", false, "Get Glyphs")
	_geomapsCmd.Flags().BoolVarP(&_geomapsGetSprites, "get-sprites", "", false, "Get Sprites")
	_geomapsCmd.Flags().BoolVarP(&_geomapsGetStaticMap, "get-static-map", "", false, "Get Static Map")
	_geomapsCmd.Flags().BoolVarP(&_geomapsGetStyleDescriptor, "get-style-descriptor", "", false, "Get Style Descriptor")
	_geomapsCmd.Flags().BoolVarP(&_geomapsGetTile, "get-tile", "", false, "Get Tile")

}
