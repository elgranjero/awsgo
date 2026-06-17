package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudcontrolCmd represents the cloudcontrol command
var _cloudcontrolCmd = &cobra.Command{
	Use:   "cloudcontrol",
	Short: "AWS cloudcontrol CLI",
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
		client := cloudcontrol.NewFromConfig(cfg)
		if _cloudcontrolCancelResourceRequest {
			cloudcontrol_CancelResourceRequest(cfg, client)
			return
		}
		if _cloudcontrolCreateResource {
			cloudcontrol_CreateResource(cfg, client)
			return
		}
		if _cloudcontrolDeleteResource {
			cloudcontrol_DeleteResource(cfg, client)
			return
		}
		if _cloudcontrolGetResource {
			cloudcontrol_GetResource(cfg, client)
			return
		}
		if _cloudcontrolGetResourceRequestStatus {
			cloudcontrol_GetResourceRequestStatus(cfg, client)
			return
		}
		if _cloudcontrolListResourceRequests {
			cloudcontrol_ListResourceRequests(cfg, client)
			return
		}
		if _cloudcontrolListResources {
			cloudcontrol_ListResources(cfg, client)
			return
		}
		if _cloudcontrolUpdateResource {
			cloudcontrol_UpdateResource(cfg, client)
			return
		}

	},
}

var (
	_cloudcontrolCancelResourceRequest    bool
	_cloudcontrolCreateResource           bool
	_cloudcontrolDeleteResource           bool
	_cloudcontrolGetResource              bool
	_cloudcontrolGetResourceRequestStatus bool
	_cloudcontrolListResourceRequests     bool
	_cloudcontrolListResources            bool
	_cloudcontrolUpdateResource           bool

	_cloudcontrolClientToken                 string
	_cloudcontrolDesiredState                string
	_cloudcontrolIdentifier                  string
	_cloudcontrolMaxResults                  string
	_cloudcontrolNextToken                   string
	_cloudcontrolPatchDocument               string
	_cloudcontrolRequestToken                string
	_cloudcontrolResourceModel               string
	_cloudcontrolResourceRequestStatusFilter string
	_cloudcontrolRoleArn                     string
	_cloudcontrolTypeName                    string
	_cloudcontrolTypeVersionId               string
)

// Cancels the specified resource operation request. For more information, see [Canceling resource operation requests] in
// the Amazon Web Services Cloud Control API User Guide.
//
// Only resource operations requests with a status of PENDING or IN_PROGRESS can
// be canceled.
//
// [Canceling resource operation requests]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-cancel
func cloudcontrol_CancelResourceRequest(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.CancelResourceRequestInput{
		// RequestToken: *string, // Required
	}

	if len(_cloudcontrolRequestToken) > 0 {
		input.RequestToken = aws.String(_cloudcontrolRequestToken)
	}

	if resp, err := client.CancelResourceRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified resource. For more information, see [Creating a resource] in the Amazon Web
// Services Cloud Control API User Guide.
//
// After you have initiated a resource creation request, you can monitor the
// progress of your request by calling [GetResourceRequestStatus]using the RequestToken of the ProgressEvent
// type returned by CreateResource .
//
// [GetResourceRequestStatus]: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
// [Creating a resource]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-create.html
func cloudcontrol_CreateResource(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.CreateResourceInput{
		// DesiredState: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_cloudcontrolDesiredState) > 0 {
		input.DesiredState = aws.String(_cloudcontrolDesiredState)
	}
	if len(_cloudcontrolTypeName) > 0 {
		input.TypeName = aws.String(_cloudcontrolTypeName)
	}
	if len(_cloudcontrolClientToken) > 0 {
		input.ClientToken = aws.String(_cloudcontrolClientToken)
	}
	if len(_cloudcontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudcontrolRoleArn)
	}
	if len(_cloudcontrolTypeVersionId) > 0 {
		input.TypeVersionId = aws.String(_cloudcontrolTypeVersionId)
	}

	if resp, err := client.CreateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource. For details, see [Deleting a resource] in the Amazon Web Services
// Cloud Control API User Guide.
//
// After you have initiated a resource deletion request, you can monitor the
// progress of your request by calling [GetResourceRequestStatus]using the RequestToken of the ProgressEvent
// returned by DeleteResource .
//
// [GetResourceRequestStatus]: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
// [Deleting a resource]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-delete.html
func cloudcontrol_DeleteResource(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.DeleteResourceInput{
		// Identifier: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_cloudcontrolIdentifier) > 0 {
		input.Identifier = aws.String(_cloudcontrolIdentifier)
	}
	if len(_cloudcontrolTypeName) > 0 {
		input.TypeName = aws.String(_cloudcontrolTypeName)
	}
	if len(_cloudcontrolClientToken) > 0 {
		input.ClientToken = aws.String(_cloudcontrolClientToken)
	}
	if len(_cloudcontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudcontrolRoleArn)
	}
	if len(_cloudcontrolTypeVersionId) > 0 {
		input.TypeVersionId = aws.String(_cloudcontrolTypeVersionId)
	}

	if resp, err := client.DeleteResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the current state of the specified resource. For
// details, see [Reading a resource's current state].
//
// You can use this action to return information about an existing resource in
// your account and Amazon Web Services Region, whether those resources were
// provisioned using Cloud Control API.
//
// [Reading a resource's current state]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-read.html
func cloudcontrol_GetResource(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.GetResourceInput{
		// Identifier: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_cloudcontrolIdentifier) > 0 {
		input.Identifier = aws.String(_cloudcontrolIdentifier)
	}
	if len(_cloudcontrolTypeName) > 0 {
		input.TypeName = aws.String(_cloudcontrolTypeName)
	}
	if len(_cloudcontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudcontrolRoleArn)
	}
	if len(_cloudcontrolTypeVersionId) > 0 {
		input.TypeVersionId = aws.String(_cloudcontrolTypeVersionId)
	}

	if resp, err := client.GetResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current status of a resource operation request. For more
// information, see [Tracking the progress of resource operation requests]in the Amazon Web Services Cloud Control API User Guide.
//
// [Tracking the progress of resource operation requests]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-track
func cloudcontrol_GetResourceRequestStatus(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.GetResourceRequestStatusInput{
		// RequestToken: *string, // Required
	}

	if len(_cloudcontrolRequestToken) > 0 {
		input.RequestToken = aws.String(_cloudcontrolRequestToken)
	}

	if resp, err := client.GetResourceRequestStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns existing resource operation requests. This includes requests of all
// status types. For more information, see [Listing active resource operation requests]in the Amazon Web Services Cloud
// Control API User Guide.
//
// Resource operation requests expire after 7 days.
//
// [Listing active resource operation requests]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-manage-requests.html#resource-operations-manage-requests-list
func cloudcontrol_ListResourceRequests(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.ListResourceRequestsInput{}

	if len(_cloudcontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudcontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudcontrolNextToken) > 0 {
		input.NextToken = aws.String(_cloudcontrolNextToken)
	}
	if len(_cloudcontrolResourceRequestStatusFilter) > 0 {
		if err := assignInputField(input, "ResourceRequestStatusFilter", _cloudcontrolResourceRequestStatusFilter); err != nil {
			log.Errorf("invalid --resource-request-status-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudcontrol.ListResourceRequestsOutput
	p := cloudcontrol.NewListResourceRequestsPaginator(client, input)
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

// Returns information about the specified resources. For more information, see [Discovering resources]
// in the Amazon Web Services Cloud Control API User Guide.
//
// You can use this action to return information about existing resources in your
// account and Amazon Web Services Region, whether those resources were provisioned
// using Cloud Control API.
//
// [Discovering resources]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-list.html
func cloudcontrol_ListResources(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.ListResourcesInput{
		// TypeName: *string, // Required
	}

	if len(_cloudcontrolTypeName) > 0 {
		input.TypeName = aws.String(_cloudcontrolTypeName)
	}
	if len(_cloudcontrolMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudcontrolMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudcontrolNextToken) > 0 {
		input.NextToken = aws.String(_cloudcontrolNextToken)
	}
	if len(_cloudcontrolResourceModel) > 0 {
		input.ResourceModel = aws.String(_cloudcontrolResourceModel)
	}
	if len(_cloudcontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudcontrolRoleArn)
	}
	if len(_cloudcontrolTypeVersionId) > 0 {
		input.TypeVersionId = aws.String(_cloudcontrolTypeVersionId)
	}

	if disablePaginator() {
		if resp, err := client.ListResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudcontrol.ListResourcesOutput
	p := cloudcontrol.NewListResourcesPaginator(client, input)
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

// Updates the specified property values in the resource.
// You specify your resource property updates as a list of patch operations
// contained in a JSON patch document that adheres to the [RFC 6902 - JavaScript Object Notation (JSON) Patch]standard.
//
// For details on how Cloud Control API performs resource update operations, see [Updating a resource]
// in the Amazon Web Services Cloud Control API User Guide.
//
// After you have initiated a resource update request, you can monitor the
// progress of your request by calling [GetResourceRequestStatus]using the RequestToken of the ProgressEvent
// returned by UpdateResource .
//
// For more information about the properties of a specific resource, refer to the
// related topic for the resource in the [Resource and property types reference]in the CloudFormation Users Guide.
//
// [GetResourceRequestStatus]: https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html
// [RFC 6902 - JavaScript Object Notation (JSON) Patch]: https://datatracker.ietf.org/doc/html/rfc6902
// [Updating a resource]: https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-update.html
// [Resource and property types reference]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
func cloudcontrol_UpdateResource(cfg aws.Config, client *cloudcontrol.Client) {
	input := &cloudcontrol.UpdateResourceInput{
		// Identifier: *string, // Required
		// PatchDocument: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_cloudcontrolIdentifier) > 0 {
		input.Identifier = aws.String(_cloudcontrolIdentifier)
	}
	if len(_cloudcontrolPatchDocument) > 0 {
		input.PatchDocument = aws.String(_cloudcontrolPatchDocument)
	}
	if len(_cloudcontrolTypeName) > 0 {
		input.TypeName = aws.String(_cloudcontrolTypeName)
	}
	if len(_cloudcontrolClientToken) > 0 {
		input.ClientToken = aws.String(_cloudcontrolClientToken)
	}
	if len(_cloudcontrolRoleArn) > 0 {
		input.RoleArn = aws.String(_cloudcontrolRoleArn)
	}
	if len(_cloudcontrolTypeVersionId) > 0 {
		input.TypeVersionId = aws.String(_cloudcontrolTypeVersionId)
	}

	if resp, err := client.UpdateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudcontrolCmd)
	_cloudcontrolCmd.Flags().SortFlags = false

	_cloudcontrolCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudcontrolCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudcontrolCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolClientToken, "client-token", "", "", "Client Token")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolDesiredState, "desired-state", "", "", "Desired State")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolIdentifier, "identifier", "", "", "Identifier")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolMaxResults, "max-results", "", "", "Max Results")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolNextToken, "next-token", "", "", "Next Token")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolPatchDocument, "patch-document", "", "", "Patch Document")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolRequestToken, "request-token", "", "", "Request Token")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolResourceModel, "resource-model", "", "", "Resource Model")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolResourceRequestStatusFilter, "resource-request-status-filter", "", "", "Resource Request Status Filter")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolRoleArn, "role-arn", "", "", "Role ARN")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolTypeName, "type-name", "", "", "Type Name")
	_cloudcontrolCmd.Flags().StringVarP(&_cloudcontrolTypeVersionId, "type-version-id", "", "", "Type Version ID")

	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolCancelResourceRequest, "cancel-resource-request", "", false, "Cancel Resource Request")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolCreateResource, "create-resource", "", false, "Create Resource")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolDeleteResource, "delete-resource", "", false, "Delete Resource")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolGetResource, "get-resource", "", false, "Get Resource")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolGetResourceRequestStatus, "get-resource-request-status", "", false, "Get Resource Request Status")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolListResourceRequests, "list-resource-requests", "", false, "List Resource Requests")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolListResources, "list-resources", "", false, "List Resources")
	_cloudcontrolCmd.Flags().BoolVarP(&_cloudcontrolUpdateResource, "update-resource", "", false, "Update Resource")

}
