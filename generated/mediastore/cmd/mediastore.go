package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mediastoreCmd represents the mediastore command
var _mediastoreCmd = &cobra.Command{
	Use:   "mediastore",
	Short: "AWS mediastore CLI",
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
		client := mediastore.NewFromConfig(cfg)
		if _mediastoreCreateContainer {
			mediastore_CreateContainer(cfg, client)
			return
		}
		if _mediastoreDeleteContainer {
			mediastore_DeleteContainer(cfg, client)
			return
		}
		if _mediastoreDeleteContainerPolicy {
			mediastore_DeleteContainerPolicy(cfg, client)
			return
		}
		if _mediastoreDeleteCorsPolicy {
			mediastore_DeleteCorsPolicy(cfg, client)
			return
		}
		if _mediastoreDeleteLifecyclePolicy {
			mediastore_DeleteLifecyclePolicy(cfg, client)
			return
		}
		if _mediastoreDeleteMetricPolicy {
			mediastore_DeleteMetricPolicy(cfg, client)
			return
		}
		if _mediastoreDescribeContainer {
			mediastore_DescribeContainer(cfg, client)
			return
		}
		if _mediastoreGetContainerPolicy {
			mediastore_GetContainerPolicy(cfg, client)
			return
		}
		if _mediastoreGetCorsPolicy {
			mediastore_GetCorsPolicy(cfg, client)
			return
		}
		if _mediastoreGetLifecyclePolicy {
			mediastore_GetLifecyclePolicy(cfg, client)
			return
		}
		if _mediastoreGetMetricPolicy {
			mediastore_GetMetricPolicy(cfg, client)
			return
		}
		if _mediastoreListContainers {
			mediastore_ListContainers(cfg, client)
			return
		}
		if _mediastoreListTagsForResource {
			mediastore_ListTagsForResource(cfg, client)
			return
		}
		if _mediastorePutContainerPolicy {
			mediastore_PutContainerPolicy(cfg, client)
			return
		}
		if _mediastorePutCorsPolicy {
			mediastore_PutCorsPolicy(cfg, client)
			return
		}
		if _mediastorePutLifecyclePolicy {
			mediastore_PutLifecyclePolicy(cfg, client)
			return
		}
		if _mediastorePutMetricPolicy {
			mediastore_PutMetricPolicy(cfg, client)
			return
		}
		if _mediastoreStartAccessLogging {
			mediastore_StartAccessLogging(cfg, client)
			return
		}
		if _mediastoreStopAccessLogging {
			mediastore_StopAccessLogging(cfg, client)
			return
		}
		if _mediastoreTagResource {
			mediastore_TagResource(cfg, client)
			return
		}
		if _mediastoreUntagResource {
			mediastore_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_mediastoreCreateContainer       bool
	_mediastoreDeleteContainer       bool
	_mediastoreDeleteContainerPolicy bool
	_mediastoreDeleteCorsPolicy      bool
	_mediastoreDeleteLifecyclePolicy bool
	_mediastoreDeleteMetricPolicy    bool
	_mediastoreDescribeContainer     bool
	_mediastoreGetContainerPolicy    bool
	_mediastoreGetCorsPolicy         bool
	_mediastoreGetLifecyclePolicy    bool
	_mediastoreGetMetricPolicy       bool
	_mediastoreListContainers        bool
	_mediastoreListTagsForResource   bool
	_mediastorePutContainerPolicy    bool
	_mediastorePutCorsPolicy         bool
	_mediastorePutLifecyclePolicy    bool
	_mediastorePutMetricPolicy       bool
	_mediastoreStartAccessLogging    bool
	_mediastoreStopAccessLogging     bool
	_mediastoreTagResource           bool
	_mediastoreUntagResource         bool

	_mediastoreContainerName   string
	_mediastoreCorsPolicy      string
	_mediastoreLifecyclePolicy string
	_mediastoreMaxResults      string
	_mediastoreMetricPolicy    string
	_mediastoreNextToken       string
	_mediastorePolicy          string
	_mediastoreResource        string
	_mediastoreTagKeys         []string
	_mediastoreTags            string
)

// Creates a storage container to hold objects. A container is similar to a bucket
// in the Amazon S3 service.
func mediastore_CreateContainer(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.CreateContainerInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}
	if len(_mediastoreTags) > 0 {
		if err := assignInputField(input, "Tags", _mediastoreTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContainer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified container. Before you make a DeleteContainer request,
// delete any objects in the container or in any folders in the container. You can
// delete only empty containers.
func mediastore_DeleteContainer(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.DeleteContainerInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.DeleteContainer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the access policy that is associated with the specified container.
func mediastore_DeleteContainerPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.DeleteContainerPolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.DeleteContainerPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the cross-origin resource sharing (CORS) configuration information that
// is set for the container.
//
// To use this operation, you must have permission to perform the
// MediaStore:DeleteCorsPolicy action. The container owner has this permission by
// default and can grant this permission to others.
func mediastore_DeleteCorsPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.DeleteCorsPolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.DeleteCorsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an object lifecycle policy from a container. It takes up to 20 minutes
// for the change to take effect.
func mediastore_DeleteLifecyclePolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.DeleteLifecyclePolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.DeleteLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the metric policy that is associated with the specified container. If
// there is no metric policy associated with the container, MediaStore doesn't send
// metrics to CloudWatch.
func mediastore_DeleteMetricPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.DeleteMetricPolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.DeleteMetricPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the properties of the requested container. This request is commonly
// used to retrieve the endpoint of a container. An endpoint is a value assigned by
// the service when a new container is created. A container's endpoint does not
// change after it has been assigned. The DescribeContainer request returns a
// single Container object based on ContainerName . To return all Container
// objects that are associated with a specified AWS account, use ListContainers.
func mediastore_DescribeContainer(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.DescribeContainerInput{}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.DescribeContainer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the access policy for the specified container. For information about
// the data that is included in an access policy, see the [AWS Identity and Access Management User Guide].
//
// [AWS Identity and Access Management User Guide]: https://aws.amazon.com/documentation/iam/
func mediastore_GetContainerPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.GetContainerPolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.GetContainerPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the cross-origin resource sharing (CORS) configuration information that
// is set for the container.
//
// To use this operation, you must have permission to perform the
// MediaStore:GetCorsPolicy action. By default, the container owner has this
// permission and can grant it to others.
func mediastore_GetCorsPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.GetCorsPolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.GetCorsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the object lifecycle policy that is assigned to a container.
func mediastore_GetLifecyclePolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.GetLifecyclePolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.GetLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metric policy for the specified container.
func mediastore_GetMetricPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.GetMetricPolicyInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.GetMetricPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the properties of all containers in AWS Elemental MediaStore.
// You can query to receive all the containers in one response. Or you can include
// the MaxResults parameter to receive a limited number of containers in each
// response. In this case, the response includes a token. To get the next set of
// containers, send the command again, this time with the NextToken parameter
// (with the returned token as its value). The next set of responses appears, with
// a token if there are still more containers to receive.
//
// See also DescribeContainer, which gets the properties of one container.
func mediastore_ListContainers(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.ListContainersInput{}

	if len(_mediastoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mediastoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mediastoreNextToken) > 0 {
		input.NextToken = aws.String(_mediastoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListContainers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mediastore.ListContainersOutput
	p := mediastore.NewListContainersPaginator(client, input)
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

// Returns a list of the tags assigned to the specified container.
func mediastore_ListTagsForResource(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.ListTagsForResourceInput{
		// Resource: *string, // Required
	}

	if len(_mediastoreResource) > 0 {
		input.Resource = aws.String(_mediastoreResource)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access policy for the specified container to restrict the users and
// clients that can access it. For information about the data that is included in
// an access policy, see the [AWS Identity and Access Management User Guide].
//
// For this release of the REST API, you can create only one policy for a
// container. If you enter PutContainerPolicy twice, the second command modifies
// the existing policy.
//
// [AWS Identity and Access Management User Guide]: https://aws.amazon.com/documentation/iam/
func mediastore_PutContainerPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.PutContainerPolicyInput{
		// ContainerName: *string, // Required
		// Policy: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}
	if len(_mediastorePolicy) > 0 {
		input.Policy = aws.String(_mediastorePolicy)
	}

	if resp, err := client.PutContainerPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the cross-origin resource sharing (CORS) configuration on a container so
// that the container can service cross-origin requests. For example, you might
// want to enable a request whose origin is http://www.example.com to access your
// AWS Elemental MediaStore container at my.example.container.com by using the
// browser's XMLHttpRequest capability.
//
// To enable CORS on a container, you attach a CORS policy to the container. In
// the CORS policy, you configure rules that identify origins and the HTTP methods
// that can be executed on your container. The policy can contain up to 398,000
// characters. You can add up to 100 rules to a CORS policy. If more than one rule
// applies, the service uses the first applicable rule listed.
//
// To learn more about CORS, see [Cross-Origin Resource Sharing (CORS) in AWS Elemental MediaStore].
//
// [Cross-Origin Resource Sharing (CORS) in AWS Elemental MediaStore]: https://docs.aws.amazon.com/mediastore/latest/ug/cors-policy.html
func mediastore_PutCorsPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.PutCorsPolicyInput{
		// ContainerName: *string, // Required
		// CorsPolicy: []types.CorsRule, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}
	if len(_mediastoreCorsPolicy) > 0 {
		if err := assignInputField(input, "CorsPolicy", _mediastoreCorsPolicy); err != nil {
			log.Errorf("invalid --cors-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutCorsPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Writes an object lifecycle policy to a container. If the container already has
// an object lifecycle policy, the service replaces the existing policy with the
// new policy. It takes up to 20 minutes for the change to take effect.
//
// For information about how to construct an object lifecycle policy, see [Components of an Object Lifecycle Policy].
//
// [Components of an Object Lifecycle Policy]: https://docs.aws.amazon.com/mediastore/latest/ug/policies-object-lifecycle-components.html
func mediastore_PutLifecyclePolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.PutLifecyclePolicyInput{
		// ContainerName: *string, // Required
		// LifecyclePolicy: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}
	if len(_mediastoreLifecyclePolicy) > 0 {
		input.LifecyclePolicy = aws.String(_mediastoreLifecyclePolicy)
	}

	if resp, err := client.PutLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The metric policy that you want to add to the container. A metric policy allows
// AWS Elemental MediaStore to send metrics to Amazon CloudWatch. It takes up to 20
// minutes for the new policy to take effect.
func mediastore_PutMetricPolicy(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.PutMetricPolicyInput{
		// ContainerName: *string, // Required
		// MetricPolicy: *types.MetricPolicy, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}
	if len(_mediastoreMetricPolicy) > 0 {
		if err := assignInputField(input, "MetricPolicy", _mediastoreMetricPolicy); err != nil {
			log.Errorf("invalid --metric-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMetricPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts access logging on the specified container. When you enable access
// logging on a container, MediaStore delivers access logs for objects stored in
// that container to Amazon CloudWatch Logs.
func mediastore_StartAccessLogging(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.StartAccessLoggingInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.StartAccessLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops access logging on the specified container. When you stop access logging
// on a container, MediaStore stops sending access logs to Amazon CloudWatch Logs.
// These access logs are not saved and are not retrievable.
func mediastore_StopAccessLogging(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.StopAccessLoggingInput{
		// ContainerName: *string, // Required
	}

	if len(_mediastoreContainerName) > 0 {
		input.ContainerName = aws.String(_mediastoreContainerName)
	}

	if resp, err := client.StopAccessLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds tags to the specified AWS Elemental MediaStore container. Tags are
// key:value pairs that you can associate with AWS resources. For example, the tag
// key might be "customer" and the tag value might be "companyA." You can specify
// one or more tags to add to each container. You can add up to 50 tags to each
// container. For more information about tagging, including naming and usage
// conventions, see [Tagging Resources in MediaStore].
//
// [Tagging Resources in MediaStore]: https://docs.aws.amazon.com/mediastore/latest/ug/tagging.html
func mediastore_TagResource(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.TagResourceInput{
		// Resource: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_mediastoreResource) > 0 {
		input.Resource = aws.String(_mediastoreResource)
	}
	if len(_mediastoreTags) > 0 {
		if err := assignInputField(input, "Tags", _mediastoreTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from the specified container. You can specify one or more tags to
// remove.
func mediastore_UntagResource(cfg aws.Config, client *mediastore.Client) {
	input := &mediastore.UntagResourceInput{
		// Resource: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mediastoreResource) > 0 {
		input.Resource = aws.String(_mediastoreResource)
	}
	if len(_mediastoreTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mediastoreTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mediastoreCmd)
	_mediastoreCmd.Flags().SortFlags = false

	_mediastoreCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_mediastoreCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mediastoreCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_mediastoreCmd.Flags().StringVarP(&_mediastoreContainerName, "container-name", "", "", "Container Name")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreCorsPolicy, "cors-policy", "", "", "Cors Policy")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreLifecyclePolicy, "lifecycle-policy", "", "", "Lifecycle Policy")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreMaxResults, "max-results", "", "", "Max Results")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreMetricPolicy, "metric-policy", "", "", "Metric Policy")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreNextToken, "next-token", "", "", "Next Token")
	_mediastoreCmd.Flags().StringVarP(&_mediastorePolicy, "policy", "", "", "Policy")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreResource, "resource", "", "", "Resource")
	_mediastoreCmd.Flags().StringSliceVarP(&_mediastoreTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mediastoreCmd.Flags().StringVarP(&_mediastoreTags, "tags", "", "", "Tags")

	_mediastoreCmd.Flags().BoolVarP(&_mediastoreCreateContainer, "create-container", "", false, "Create Container")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreDeleteContainer, "delete-container", "", false, "Delete Container")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreDeleteContainerPolicy, "delete-container-policy", "", false, "Delete Container Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreDeleteCorsPolicy, "delete-cors-policy", "", false, "Delete Cors Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreDeleteLifecyclePolicy, "delete-lifecycle-policy", "", false, "Delete Lifecycle Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreDeleteMetricPolicy, "delete-metric-policy", "", false, "Delete Metric Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreDescribeContainer, "describe-container", "", false, "Describe Container")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreGetContainerPolicy, "get-container-policy", "", false, "Get Container Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreGetCorsPolicy, "get-cors-policy", "", false, "Get Cors Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreGetLifecyclePolicy, "get-lifecycle-policy", "", false, "Get Lifecycle Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreGetMetricPolicy, "get-metric-policy", "", false, "Get Metric Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreListContainers, "list-containers", "", false, "List Containers")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mediastoreCmd.Flags().BoolVarP(&_mediastorePutContainerPolicy, "put-container-policy", "", false, "Put Container Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastorePutCorsPolicy, "put-cors-policy", "", false, "Put Cors Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastorePutLifecyclePolicy, "put-lifecycle-policy", "", false, "Put Lifecycle Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastorePutMetricPolicy, "put-metric-policy", "", false, "Put Metric Policy")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreStartAccessLogging, "start-access-logging", "", false, "Start Access Logging")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreStopAccessLogging, "stop-access-logging", "", false, "Stop Access Logging")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreTagResource, "tag-resource", "", false, "Tag Resource")
	_mediastoreCmd.Flags().BoolVarP(&_mediastoreUntagResource, "untag-resource", "", false, "Untag Resource")

}
