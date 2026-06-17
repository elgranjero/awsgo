package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediastoredataCmd represents the mediastoredata command
var _mediastoredataCmd = &cobra.Command{
	Use:   "mediastoredata",
	Short: "AWS mediastoredata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mediastoredata.NewFromConfig(cfg)
		if _mediastoredataDeleteObject {
			mediastoredata_DeleteObject(cfg, client)
			return
		}
		if _mediastoredataDescribeObject {
			mediastoredata_DescribeObject(cfg, client)
			return
		}
		if _mediastoredataGetObject {
			mediastoredata_GetObject(cfg, client)
			return
		}
		if _mediastoredataListItems {
			mediastoredata_ListItems(cfg, client)
			return
		}
		if _mediastoredataPutObject {
			mediastoredata_PutObject(cfg, client)
			return
		}

	},
}

var (
	_mediastoredataDeleteObject   bool
	_mediastoredataDescribeObject bool
	_mediastoredataGetObject      bool
	_mediastoredataListItems      bool
	_mediastoredataPutObject      bool

	_mediastoredataBody               string
	_mediastoredataCacheControl       string
	_mediastoredataContentType        string
	_mediastoredataMaxResults         string
	_mediastoredataNextToken          string
	_mediastoredataPath               string
	_mediastoredataRange              string
	_mediastoredataStorageClass       string
	_mediastoredataUploadAvailability string
)

// Deletes an object at the specified path.
func mediastoredata_DeleteObject(cfg aws.Config, client *mediastoredata.Client) {
	input := &mediastoredata.DeleteObjectInput{
		// Path: *string, // Required
	}

	if len(_mediastoredataPath) > 0 {
		input.Path = aws.String(_mediastoredataPath)
	}

	if resp, err := client.DeleteObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the headers for an object at the specified path.
func mediastoredata_DescribeObject(cfg aws.Config, client *mediastoredata.Client) {
	input := &mediastoredata.DescribeObjectInput{
		// Path: *string, // Required
	}

	if len(_mediastoredataPath) > 0 {
		input.Path = aws.String(_mediastoredataPath)
	}

	if resp, err := client.DescribeObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Downloads the object at the specified path. If the object’s upload availability
// is set to streaming , AWS Elemental MediaStore downloads the object even if it’s
// still uploading the object.
func mediastoredata_GetObject(cfg aws.Config, client *mediastoredata.Client) {
	input := &mediastoredata.GetObjectInput{
		// Path: *string, // Required
	}

	if len(_mediastoredataPath) > 0 {
		input.Path = aws.String(_mediastoredataPath)
	}
	if len(_mediastoredataRange) > 0 {
		input.Range = aws.String(_mediastoredataRange)
	}

	if resp, err := client.GetObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of metadata entries about folders and objects in the specified
// folder.
func mediastoredata_ListItems(cfg aws.Config, client *mediastoredata.Client) {
	input := &mediastoredata.ListItemsInput{}

	if len(_mediastoredataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediastoredataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediastoredataNextToken) > 0 {
		input.NextToken = aws.String(_mediastoredataNextToken)
	}
	if len(_mediastoredataPath) > 0 {
		input.Path = aws.String(_mediastoredataPath)
	}

	if disablePaginator() {
		if resp, err := client.ListItems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediastoredata.ListItemsOutput
	p := mediastoredata.NewListItemsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Uploads an object to the specified path. Object sizes are limited to 25 MB for
// standard upload availability and 10 MB for streaming upload availability.
func mediastoredata_PutObject(cfg aws.Config, client *mediastoredata.Client) {
	input := &mediastoredata.PutObjectInput{
		// Body: io.Reader, // Required
		// Path: *string, // Required
	}

	if len(_mediastoredataBody) > 0 {
		if err := assignInputField(input, "Body", _mediastoredataBody); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_mediastoredataPath) > 0 {
		input.Path = aws.String(_mediastoredataPath)
	}
	if len(_mediastoredataCacheControl) > 0 {
		input.CacheControl = aws.String(_mediastoredataCacheControl)
	}
	if len(_mediastoredataContentType) > 0 {
		input.ContentType = aws.String(_mediastoredataContentType)
	}
	if len(_mediastoredataStorageClass) > 0 {
		if err := assignInputField(input, "StorageClass", _mediastoredataStorageClass); err != nil {
			log.Errorf("invalid --storage-class: %s", err.Error())
			return
		}
	}
	if len(_mediastoredataUploadAvailability) > 0 {
		if err := assignInputField(input, "UploadAvailability", _mediastoredataUploadAvailability); err != nil {
			log.Errorf("invalid --upload-availability: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediastoredataCmd)
	_mediastoredataCmd.Flags().SortFlags = false

	_mediastoredataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_mediastoredataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediastoredataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataBody, "body", "", "", "Body")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataCacheControl, "cache-control", "", "", "Cache Control")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataContentType, "content-type", "", "", "Content Type")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataMaxResults, "max-results", "", "", "Max Results")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataNextToken, "next-token", "", "", "Next Token")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataPath, "path", "", "", "Path")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataRange, "range", "", "", "Range")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataStorageClass, "storage-class", "", "", "Storage Class")
	_mediastoredataCmd.Flags().StringVarP(&_mediastoredataUploadAvailability, "upload-availability", "", "", "Upload Availability")

	_mediastoredataCmd.Flags().BoolVarP(&_mediastoredataDeleteObject, "delete-object", "", false, "Delete Object")
	_mediastoredataCmd.Flags().BoolVarP(&_mediastoredataDescribeObject, "describe-object", "", false, "Describe Object")
	_mediastoredataCmd.Flags().BoolVarP(&_mediastoredataGetObject, "get-object", "", false, "Get Object")
	_mediastoredataCmd.Flags().BoolVarP(&_mediastoredataListItems, "list-items", "", false, "List Items")
	_mediastoredataCmd.Flags().BoolVarP(&_mediastoredataPutObject, "put-object", "", false, "Put Object")

}
