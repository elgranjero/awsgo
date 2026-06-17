package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sqsCmd represents the sqs command
var _sqsCmd = &cobra.Command{
	Use:   "sqs",
	Short: "AWS sqs CLI",
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
		client := sqs.NewFromConfig(cfg)
		if _sqsAddPermission {
			sqs_AddPermission(cfg, client)
			return
		}
		if _sqsCancelMessageMoveTask {
			sqs_CancelMessageMoveTask(cfg, client)
			return
		}
		if _sqsChangeMessageVisibility {
			sqs_ChangeMessageVisibility(cfg, client)
			return
		}
		if _sqsChangeMessageVisibilityBatch {
			sqs_ChangeMessageVisibilityBatch(cfg, client)
			return
		}
		if _sqsCreateQueue {
			sqs_CreateQueue(cfg, client)
			return
		}
		if _sqsDeleteMessage {
			sqs_DeleteMessage(cfg, client)
			return
		}
		if _sqsDeleteMessageBatch {
			sqs_DeleteMessageBatch(cfg, client)
			return
		}
		if _sqsDeleteQueue {
			sqs_DeleteQueue(cfg, client)
			return
		}
		if _sqsGetQueueAttributes {
			sqs_GetQueueAttributes(cfg, client)
			return
		}
		if _sqsGetQueueUrl {
			sqs_GetQueueUrl(cfg, client)
			return
		}
		if _sqsListDeadLetterSourceQueues {
			sqs_ListDeadLetterSourceQueues(cfg, client)
			return
		}
		if _sqsListMessageMoveTasks {
			sqs_ListMessageMoveTasks(cfg, client)
			return
		}
		if _sqsListQueueTags {
			sqs_ListQueueTags(cfg, client)
			return
		}
		if _sqsListQueues {
			sqs_ListQueues(cfg, client)
			return
		}
		if _sqsPurgeQueue {
			sqs_PurgeQueue(cfg, client)
			return
		}
		if _sqsReceiveMessage {
			sqs_ReceiveMessage(cfg, client)
			return
		}
		if _sqsRemovePermission {
			sqs_RemovePermission(cfg, client)
			return
		}
		if _sqsSendMessage {
			sqs_SendMessage(cfg, client)
			return
		}
		if _sqsSendMessageBatch {
			sqs_SendMessageBatch(cfg, client)
			return
		}
		if _sqsSetQueueAttributes {
			sqs_SetQueueAttributes(cfg, client)
			return
		}
		if _sqsStartMessageMoveTask {
			sqs_StartMessageMoveTask(cfg, client)
			return
		}
		if _sqsTagQueue {
			sqs_TagQueue(cfg, client)
			return
		}
		if _sqsUntagQueue {
			sqs_UntagQueue(cfg, client)
			return
		}

	},
}

var (
	_sqsAddPermission                bool
	_sqsCancelMessageMoveTask        bool
	_sqsChangeMessageVisibility      bool
	_sqsChangeMessageVisibilityBatch bool
	_sqsCreateQueue                  bool
	_sqsDeleteMessage                bool
	_sqsDeleteMessageBatch           bool
	_sqsDeleteQueue                  bool
	_sqsGetQueueAttributes           bool
	_sqsGetQueueUrl                  bool
	_sqsListDeadLetterSourceQueues   bool
	_sqsListMessageMoveTasks         bool
	_sqsListQueueTags                bool
	_sqsListQueues                   bool
	_sqsPurgeQueue                   bool
	_sqsReceiveMessage               bool
	_sqsRemovePermission             bool
	_sqsSendMessage                  bool
	_sqsSendMessageBatch             bool
	_sqsSetQueueAttributes           bool
	_sqsStartMessageMoveTask         bool
	_sqsTagQueue                     bool
	_sqsUntagQueue                   bool

	_sqsActions                      []string
	_sqsAttributeNames               string
	_sqsAttributes                   string
	_sqsAWSAccountIds                []string
	_sqsDelaySeconds                 string
	_sqsDestinationArn               string
	_sqsEntries                      string
	_sqsLabel                        string
	_sqsMaxNumberOfMessages          string
	_sqsMaxNumberOfMessagesPerSecond string
	_sqsMaxResults                   string
	_sqsMessageAttributeNames        []string
	_sqsMessageAttributes            string
	_sqsMessageBody                  string
	_sqsMessageDeduplicationId       string
	_sqsMessageGroupId               string
	_sqsMessageSystemAttributeNames  string
	_sqsMessageSystemAttributes      string
	_sqsNextToken                    string
	_sqsQueueName                    string
	_sqsQueueNamePrefix              string
	_sqsQueueOwnerAWSAccountId       string
	_sqsQueueUrl                     string
	_sqsReceiptHandle                string
	_sqsReceiveRequestAttemptId      string
	_sqsSourceArn                    string
	_sqsTagKeys                      []string
	_sqsTags                         string
	_sqsTaskHandle                   string
	_sqsVisibilityTimeout            string
	_sqsWaitTimeSeconds              string
)

// Adds a permission to a queue for a specific [principal]. This allows sharing access to the
// queue.
//
// When you create a queue, you have full control access rights for the queue.
// Only you, the owner of the queue, can grant or deny permissions to the queue.
// For more information about these permissions, see [Allow Developers to Write Messages to a Shared Queue]in the Amazon SQS Developer
// Guide.
//
// - AddPermission generates a policy for you. You can use SetQueueAttributesto upload your
// policy. For more information, see [Using Custom Policies with the Amazon SQS Access Policy Language]in the Amazon SQS Developer Guide.
//
// - An Amazon SQS policy can have a maximum of seven actions per statement.
//
// - To remove the ability to change queue permissions, you must deny permission
// to the AddPermission , RemovePermission , and SetQueueAttributes actions in
// your IAM policy.
//
// - Amazon SQS AddPermission does not support adding a non-account principal.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Allow Developers to Write Messages to a Shared Queue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue
// [Using Custom Policies with the Amazon SQS Access Policy Language]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
// [principal]: https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P
func sqs_AddPermission(cfg aws.Config, client *sqs.Client) {
	input := &sqs.AddPermissionInput{
		// AWSAccountIds: []string, // Required
		// Actions: []string, // Required
		// Label: *string, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsAWSAccountIds) > 0 {
		input.AWSAccountIds = append([]string(nil), _sqsAWSAccountIds...)
	}
	if len(_sqsActions) > 0 {
		input.Actions = append([]string(nil), _sqsActions...)
	}
	if len(_sqsLabel) > 0 {
		input.Label = aws.String(_sqsLabel)
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.AddPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a specified message movement task. A message movement can only be
// cancelled when the current status is RUNNING. Cancelling a message movement task
// does not revert the messages that have already been moved. It can only stop the
// messages that have not been moved yet.
//
// - This action is currently limited to supporting message redrive from [dead-letter queues (DLQs)]only.
// In this context, the source queue is the dead-letter queue (DLQ), while the
// destination queue can be the original source queue (from which the messages were
// driven to the dead-letter-queue), or a custom destination queue.
//
// - Only one active message movement task is supported per queue at any given
// time.
//
// [dead-letter queues (DLQs)]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
func sqs_CancelMessageMoveTask(cfg aws.Config, client *sqs.Client) {
	input := &sqs.CancelMessageMoveTaskInput{
		// TaskHandle: *string, // Required
	}

	if len(_sqsTaskHandle) > 0 {
		input.TaskHandle = aws.String(_sqsTaskHandle)
	}

	if resp, err := client.CancelMessageMoveTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the visibility timeout of a specified message in a queue to a new
// value. The default visibility timeout for a message is 30 seconds. The minimum
// is 0 seconds. The maximum is 12 hours. For more information, see [Visibility Timeout]in the Amazon
// SQS Developer Guide.
//
// For example, if the default timeout for a queue is 60 seconds, 15 seconds have
// elapsed since you received the message, and you send a ChangeMessageVisibility
// call with VisibilityTimeout set to 10 seconds, the 10 seconds begin to count
// from the time that you make the ChangeMessageVisibility call. Thus, any attempt
// to change the visibility timeout or to delete that message 10 seconds after you
// initially change the visibility timeout (a total of 25 seconds) might result in
// an error.
//
// An Amazon SQS message has three basic states:
//
// - Sent to a queue by a producer.
//
// - Received from the queue by a consumer.
//
// - Deleted from the queue.
//
// A message is considered to be stored after it is sent to a queue by a producer,
// but not yet received from the queue by a consumer (that is, between states 1 and
// 2). There is no limit to the number of stored messages. A message is considered
// to be in flight after it is received from a queue by a consumer, but not yet
// deleted from the queue (that is, between states 2 and 3). There is a limit to
// the number of in flight messages.
//
// Limits that apply to in flight messages are unrelated to the unlimited number
// of stored messages.
//
// For most standard queues (depending on queue traffic and message backlog),
// there can be a maximum of approximately 120,000 in flight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns the OverLimit error message. To avoid reaching
// the limit, you should delete messages from the queue after they're processed.
// You can also increase the number of queues you use to process your messages. To
// request a limit increase, [file a support request].
//
// For FIFO queues, there can be a maximum of 120,000 in flight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns no error messages.
//
// If you attempt to set the VisibilityTimeout to a value greater than the maximum
// time left, Amazon SQS returns an error. Amazon SQS doesn't automatically
// recalculate and increase the timeout to the maximum remaining time.
//
// Unlike with a queue, when you change the visibility timeout for a specific
// message the timeout value is applied immediately but isn't saved in memory for
// that message. If you don't delete a message after it is received, the visibility
// timeout for the message reverts to the original timeout value (not to the value
// you set using the ChangeMessageVisibility action) the next time the message is
// received.
//
// [Visibility Timeout]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html
// [file a support request]: https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-sqs
func sqs_ChangeMessageVisibility(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ChangeMessageVisibilityInput{
		// QueueUrl: *string, // Required
		// ReceiptHandle: *string, // Required
		// VisibilityTimeout: int32, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsReceiptHandle) > 0 {
		input.ReceiptHandle = aws.String(_sqsReceiptHandle)
	}
	if len(_sqsVisibilityTimeout) > 0 {
		if err := assignInputField(input, "VisibilityTimeout", _sqsVisibilityTimeout); err != nil {
			log.Errorf("invalid --visibility-timeout: %s", err.Error())
			return
		}
	}

	if resp, err := client.ChangeMessageVisibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes the visibility timeout of multiple messages. This is a batch version of ChangeMessageVisibility
// . The result of the action on each message is reported individually in the
// response. You can send up to 10 ChangeMessageVisibilityrequests with each ChangeMessageVisibilityBatch
// action.
//
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
func sqs_ChangeMessageVisibilityBatch(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ChangeMessageVisibilityBatchInput{
		// Entries: []types.ChangeMessageVisibilityBatchRequestEntry, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsEntries) > 0 {
		if err := assignInputField(input, "Entries", _sqsEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.ChangeMessageVisibilityBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new standard or FIFO queue. You can pass one or more attributes in
// the request. Keep the following in mind:
//
// - If you don't specify the FifoQueue attribute, Amazon SQS creates a standard
// queue.
//
// You can't change the queue type after you create it and you can't convert an
//
// existing standard queue into a FIFO queue. You must either create a new FIFO
// queue for your application or delete your existing standard queue and recreate
// it as a FIFO queue. For more information, see [Moving From a standard queue to a FIFO queue]in the Amazon SQS Developer
// Guide.
//
// - If you don't provide a value for an attribute, the queue is created with
// the default value for the attribute.
//
// - If you delete a queue, you must wait at least 60 seconds before creating a
// queue with the same name.
//
// To successfully create a new queue, you must provide a queue name that adheres
// to the [limits related to queues]and is unique within the scope of your queues.
//
// After you create a queue, you must wait at least one second after the queue is
// created to be able to use the queue.
//
// To retrieve the URL of a queue, use the [GetQueueUrl]GetQueueUrl action. This action only
// requires the [QueueName]QueueName parameter.
//
// When creating queues, keep the following points in mind:
//
// - If you specify the name of an existing queue and provide the exact same
// names and values for all its attributes, the [CreateQueue]CreateQueue action will return
// the URL of the existing queue instead of creating a new one.
//
// - If you attempt to create a queue with a name that already exists but with
// different attribute names or values, the CreateQueue action will return an
// error. This ensures that existing queues are not inadvertently altered.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [limits related to queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html
// [CreateQueue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
// [QueueName]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html#API_CreateQueue_RequestSyntax
// [GetQueueUrl]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_GetQueueUrl.html
// [Moving From a standard queue to a FIFO queue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-moving
func sqs_CreateQueue(cfg aws.Config, client *sqs.Client) {
	input := &sqs.CreateQueueInput{
		// QueueName: *string, // Required
	}

	if len(_sqsQueueName) > 0 {
		input.QueueName = aws.String(_sqsQueueName)
	}
	if len(_sqsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _sqsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_sqsTags) > 0 {
		if err := assignInputField(input, "Tags", _sqsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified message from the specified queue. To select the message
// to delete, use the ReceiptHandle of the message (not the MessageId which you
// receive when you send the message). Amazon SQS can delete a message from a queue
// even if a visibility timeout setting causes the message to be locked by another
// consumer. Amazon SQS automatically deletes messages left in a queue longer than
// the retention period configured for the queue.
//
// Each time you receive a message, meaning when a consumer retrieves a message
// from the queue, it comes with a unique ReceiptHandle . If you receive the same
// message more than once, you will get a different ReceiptHandle each time. When
// you want to delete a message using the DeleteMessage action, you must use the
// ReceiptHandle from the most recent time you received the message. If you use an
// old ReceiptHandle , the request will succeed, but the message might not be
// deleted.
//
// For standard queues, it is possible to receive a message even after you delete
// it. This might happen on rare occasions if one of the servers which stores a
// copy of the message is unavailable when you send the request to delete the
// message. The copy remains on the server and might be returned to you during a
// subsequent receive request. You should ensure that your application is
// idempotent, so that receiving a message more than once does not cause issues.
func sqs_DeleteMessage(cfg aws.Config, client *sqs.Client) {
	input := &sqs.DeleteMessageInput{
		// QueueUrl: *string, // Required
		// ReceiptHandle: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsReceiptHandle) > 0 {
		input.ReceiptHandle = aws.String(_sqsReceiptHandle)
	}

	if resp, err := client.DeleteMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes up to ten messages from the specified queue. This is a batch version of DeleteMessage
// . The result of the action on each message is reported individually in the
// response.
//
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
func sqs_DeleteMessageBatch(cfg aws.Config, client *sqs.Client) {
	input := &sqs.DeleteMessageBatchInput{
		// Entries: []types.DeleteMessageBatchRequestEntry, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsEntries) > 0 {
		if err := assignInputField(input, "Entries", _sqsEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.DeleteMessageBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the queue specified by the QueueUrl , regardless of the queue's contents.
// Be careful with the DeleteQueue action: When you delete a queue, any messages
// in the queue are no longer available.
//
// When you delete a queue, the deletion process takes up to 60 seconds. Requests
// you send involving that queue during the 60 seconds might succeed. For example,
// a SendMessagerequest might succeed, but after 60 seconds the queue and the message you
// sent no longer exist.
//
// When you delete a queue, you must wait at least 60 seconds before creating a
// queue with the same name.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// The delete operation uses the HTTP GET verb.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_DeleteQueue(cfg aws.Config, client *sqs.Client) {
	input := &sqs.DeleteQueueInput{
		// QueueUrl: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.DeleteQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets attributes for the specified queue.
// To determine whether a queue is [FIFO], you can check whether QueueName ends with the
// .fifo suffix.
//
// [FIFO]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html
func sqs_GetQueueAttributes(cfg aws.Config, client *sqs.Client) {
	input := &sqs.GetQueueAttributesInput{
		// QueueUrl: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsAttributeNames) > 0 {
		if err := assignInputField(input, "AttributeNames", _sqsAttributeNames); err != nil {
			log.Errorf("invalid --attribute-names: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetQueueAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetQueueUrl API returns the URL of an existing Amazon SQS queue. This is
// useful when you know the queue's name but need to retrieve its URL for further
// operations.
//
// To access a queue owned by another Amazon Web Services account, use the
// QueueOwnerAWSAccountId parameter to specify the account ID of the queue's owner.
// Note that the queue owner must grant you the necessary permissions to access the
// queue. For more information about accessing shared queues, see the AddPermissionAPI or [Allow developers to write messages to a shared queue] in
// the Amazon SQS Developer Guide.
//
// [Allow developers to write messages to a shared queue]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue
func sqs_GetQueueUrl(cfg aws.Config, client *sqs.Client) {
	input := &sqs.GetQueueUrlInput{
		// QueueName: *string, // Required
	}

	if len(_sqsQueueName) > 0 {
		input.QueueName = aws.String(_sqsQueueName)
	}
	if len(_sqsQueueOwnerAWSAccountId) > 0 {
		input.QueueOwnerAWSAccountId = aws.String(_sqsQueueOwnerAWSAccountId)
	}

	if resp, err := client.GetQueueUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of your queues that have the RedrivePolicy queue attribute
// configured with a dead-letter queue.
//
// The ListDeadLetterSourceQueues methods supports pagination. Set parameter
// MaxResults in the request to specify the maximum number of results to be
// returned in the response. If you do not set MaxResults , the response includes a
// maximum of 1,000 results. If you set MaxResults and there are additional
// results to display, the response includes a value for NextToken . Use NextToken
// as a parameter in your next request to ListDeadLetterSourceQueues to receive
// the next page of results.
//
// For more information about using dead-letter queues, see [Using Amazon SQS Dead-Letter Queues] in the Amazon SQS
// Developer Guide.
//
// [Using Amazon SQS Dead-Letter Queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
func sqs_ListDeadLetterSourceQueues(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ListDeadLetterSourceQueuesInput{
		// QueueUrl: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sqsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sqsNextToken) > 0 {
		input.NextToken = aws.String(_sqsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDeadLetterSourceQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sqs.ListDeadLetterSourceQueuesOutput
	p := sqs.NewListDeadLetterSourceQueuesPaginator(client, input)
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

// Gets the most recent message movement tasks (up to 10) under a specific source
// queue.
//
// - This action is currently limited to supporting message redrive from [dead-letter queues (DLQs)]only.
// In this context, the source queue is the dead-letter queue (DLQ), while the
// destination queue can be the original source queue (from which the messages were
// driven to the dead-letter-queue), or a custom destination queue.
//
// - Only one active message movement task is supported per queue at any given
// time.
//
// [dead-letter queues (DLQs)]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
func sqs_ListMessageMoveTasks(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ListMessageMoveTasksInput{
		// SourceArn: *string, // Required
	}

	if len(_sqsSourceArn) > 0 {
		input.SourceArn = aws.String(_sqsSourceArn)
	}
	if len(_sqsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sqsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListMessageMoveTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all cost allocation tags added to the specified Amazon SQS queue. For an
// overview, see [Tagging Your Amazon SQS Queues]in the Amazon SQS Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Tagging Your Amazon SQS Queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_ListQueueTags(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ListQueueTagsInput{
		// QueueUrl: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.ListQueueTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of your queues in the current region. The response includes a
// maximum of 1,000 results. If you specify a value for the optional
// QueueNamePrefix parameter, only queues with a name that begins with the
// specified value are returned.
//
// The listQueues methods supports pagination. Set parameter MaxResults in the
// request to specify the maximum number of results to be returned in the response.
// If you do not set MaxResults , the response includes a maximum of 1,000 results.
// If you set MaxResults and there are additional results to display, the response
// includes a value for NextToken . Use NextToken as a parameter in your next
// request to listQueues to receive the next page of results.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_ListQueues(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ListQueuesInput{}

	if len(_sqsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sqsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sqsNextToken) > 0 {
		input.NextToken = aws.String(_sqsNextToken)
	}
	if len(_sqsQueueNamePrefix) > 0 {
		input.QueueNamePrefix = aws.String(_sqsQueueNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListQueues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sqs.ListQueuesOutput
	p := sqs.NewListQueuesPaginator(client, input)
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

// Deletes available messages in a queue (including in-flight messages) specified
// by the QueueURL parameter.
//
// When you use the PurgeQueue action, you can't retrieve any messages deleted
// from a queue.
//
// The message deletion process takes up to 60 seconds. We recommend waiting for
// 60 seconds regardless of your queue's size.
//
// Messages sent to the queue before you call PurgeQueue might be received but are
// deleted within the next minute.
//
// Messages sent to the queue after you call PurgeQueue might be deleted while the
// queue is being purged.
func sqs_PurgeQueue(cfg aws.Config, client *sqs.Client) {
	input := &sqs.PurgeQueueInput{
		// QueueUrl: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.PurgeQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves one or more messages (up to 10), from the specified queue. Using the
// WaitTimeSeconds parameter enables long-poll support. For more information, see [Amazon SQS Long Polling]
// in the Amazon SQS Developer Guide.
//
// Short poll is the default behavior where a weighted random set of machines is
// sampled on a ReceiveMessage call. Therefore, only the messages on the sampled
// machines are returned. If the number of messages in the queue is small (fewer
// than 1,000), you most likely get fewer messages than you requested per
// ReceiveMessage call. If the number of messages in the queue is extremely small,
// you might not receive any messages in a particular ReceiveMessage response. If
// this happens, repeat the request.
//
// For each message returned, the response includes the following:
//
// - The message body.
//
// - An MD5 digest of the message body. For information about MD5, see [RFC1321].
//
// - The MessageId you received when you sent the message to the queue.
//
// - The receipt handle.
//
// - The message attributes.
//
// - An MD5 digest of the message attributes.
//
// The receipt handle is the identifier you must provide when deleting the
// message. For more information, see [Queue and Message Identifiers]in the Amazon SQS Developer Guide.
//
// You can provide the VisibilityTimeout parameter in your request. The parameter
// is applied to the messages that Amazon SQS returns in the response. If you don't
// include the parameter, the overall visibility timeout for the queue is used for
// the returned messages. The default visibility timeout for a queue is 30 seconds.
//
// In the future, new attributes might be added. If you write code that calls this
// action, we recommend that you structure your code so that it can handle new
// attributes gracefully.
//
// [Queue and Message Identifiers]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html
// [Amazon SQS Long Polling]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html
// [RFC1321]: https://www.ietf.org/rfc/rfc1321.txt
func sqs_ReceiveMessage(cfg aws.Config, client *sqs.Client) {
	input := &sqs.ReceiveMessageInput{
		// QueueUrl: *string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsAttributeNames) > 0 {
		if err := assignInputField(input, "AttributeNames", _sqsAttributeNames); err != nil {
			log.Errorf("invalid --attribute-names: %s", err.Error())
			return
		}
	}
	if len(_sqsMaxNumberOfMessages) > 0 {
		if err := assignInputField(input, "MaxNumberOfMessages", _sqsMaxNumberOfMessages); err != nil {
			log.Errorf("invalid --max-number-of-messages: %s", err.Error())
			return
		}
	}
	if len(_sqsMessageAttributeNames) > 0 {
		input.MessageAttributeNames = append([]string(nil), _sqsMessageAttributeNames...)
	}
	if len(_sqsMessageSystemAttributeNames) > 0 {
		if err := assignInputField(input, "MessageSystemAttributeNames", _sqsMessageSystemAttributeNames); err != nil {
			log.Errorf("invalid --message-system-attribute-names: %s", err.Error())
			return
		}
	}
	if len(_sqsReceiveRequestAttemptId) > 0 {
		input.ReceiveRequestAttemptId = aws.String(_sqsReceiveRequestAttemptId)
	}
	if len(_sqsVisibilityTimeout) > 0 {
		if err := assignInputField(input, "VisibilityTimeout", _sqsVisibilityTimeout); err != nil {
			log.Errorf("invalid --visibility-timeout: %s", err.Error())
			return
		}
	}
	if len(_sqsWaitTimeSeconds) > 0 {
		if err := assignInputField(input, "WaitTimeSeconds", _sqsWaitTimeSeconds); err != nil {
			log.Errorf("invalid --wait-time-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReceiveMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes any permissions in the queue policy that matches the specified Label
// parameter.
//
// - Only the owner of a queue can remove permissions from it.
//
// - Cross-account permissions don't apply to this action. For more information,
// see [Grant cross-account permissions to a role and a username]in the Amazon SQS Developer Guide.
//
// - To remove the ability to change queue permissions, you must deny permission
// to the AddPermission , RemovePermission , and SetQueueAttributes actions in
// your IAM policy.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_RemovePermission(cfg aws.Config, client *sqs.Client) {
	input := &sqs.RemovePermissionInput{
		// Label: *string, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsLabel) > 0 {
		input.Label = aws.String(_sqsLabel)
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.RemovePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delivers a message to the specified queue.
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed. For more information, see the [W3C specification for characters].
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// If a message contains characters outside the allowed set, Amazon SQS rejects
// the message and returns an InvalidMessageContents error. Ensure that your
// message body includes only valid characters to avoid this exception.
//
// [W3C specification for characters]: http://www.w3.org/TR/REC-xml/#charsets
func sqs_SendMessage(cfg aws.Config, client *sqs.Client) {
	input := &sqs.SendMessageInput{
		// MessageBody: *string, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsMessageBody) > 0 {
		input.MessageBody = aws.String(_sqsMessageBody)
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsDelaySeconds) > 0 {
		if err := assignInputField(input, "DelaySeconds", _sqsDelaySeconds); err != nil {
			log.Errorf("invalid --delay-seconds: %s", err.Error())
			return
		}
	}
	if len(_sqsMessageAttributes) > 0 {
		if err := assignInputField(input, "MessageAttributes", _sqsMessageAttributes); err != nil {
			log.Errorf("invalid --message-attributes: %s", err.Error())
			return
		}
	}
	if len(_sqsMessageDeduplicationId) > 0 {
		input.MessageDeduplicationId = aws.String(_sqsMessageDeduplicationId)
	}
	if len(_sqsMessageGroupId) > 0 {
		input.MessageGroupId = aws.String(_sqsMessageGroupId)
	}
	if len(_sqsMessageSystemAttributes) > 0 {
		if err := assignInputField(input, "MessageSystemAttributes", _sqsMessageSystemAttributes); err != nil {
			log.Errorf("invalid --message-system-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use SendMessageBatch to send up to 10 messages to the specified queue
// by assigning either identical or different values to each message (or by not
// assigning values at all). This is a batch version of SendMessage. For a FIFO queue,
// multiple messages within a single batch are enqueued in the order they are sent.
//
// The result of sending each message is reported individually in the response.
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200 .
//
// The maximum allowed individual message size and the maximum total payload size
// (the sum of the individual lengths of all of the batched messages) are both 1
// MiB 1,048,576 bytes.
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed. For more information, see the [W3C specification for characters].
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// If a message contains characters outside the allowed set, Amazon SQS rejects
// the message and returns an InvalidMessageContents error. Ensure that your
// message body includes only valid characters to avoid this exception.
//
// If you don't specify the DelaySeconds parameter for an entry, Amazon SQS uses
// the default value for the queue.
//
// [W3C specification for characters]: http://www.w3.org/TR/REC-xml/#charsets
func sqs_SendMessageBatch(cfg aws.Config, client *sqs.Client) {
	input := &sqs.SendMessageBatchInput{
		// Entries: []types.SendMessageBatchRequestEntry, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsEntries) > 0 {
		if err := assignInputField(input, "Entries", _sqsEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.SendMessageBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the value of one or more queue attributes, like a policy. When you change
// a queue's attributes, the change can take up to 60 seconds for most of the
// attributes to propagate throughout the Amazon SQS system. Changes made to the
// MessageRetentionPeriod attribute can take up to 15 minutes and will impact
// existing messages in the queue potentially causing them to be expired and
// deleted if the MessageRetentionPeriod is reduced below the age of existing
// messages.
//
// - In the future, new attributes might be added. If you write code that calls
// this action, we recommend that you structure your code so that it can handle new
// attributes gracefully.
//
// - Cross-account permissions don't apply to this action. For more information,
// see [Grant cross-account permissions to a role and a username]in the Amazon SQS Developer Guide.
//
// - To remove the ability to change queue permissions, you must deny permission
// to the AddPermission , RemovePermission , and SetQueueAttributes actions in
// your IAM policy.
//
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_SetQueueAttributes(cfg aws.Config, client *sqs.Client) {
	input := &sqs.SetQueueAttributesInput{
		// Attributes: map[string]string, // Required
		// QueueUrl: *string, // Required
	}

	if len(_sqsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _sqsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}

	if resp, err := client.SetQueueAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an asynchronous task to move messages from a specified source queue to a
// specified destination queue.
//
// - This action is currently limited to supporting message redrive from queues
// that are configured as [dead-letter queues (DLQs)]of other Amazon SQS queues only. Non-SQS queue sources
// of dead-letter queues, such as Lambda or Amazon SNS topics, are currently not
// supported.
//
// - In dead-letter queues redrive context, the StartMessageMoveTask the source
// queue is the DLQ, while the destination queue can be the original source queue
// (from which the messages were driven to the dead-letter-queue), or a custom
// destination queue.
//
// - Only one active message movement task is supported per queue at any given
// time.
//
// [dead-letter queues (DLQs)]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html
func sqs_StartMessageMoveTask(cfg aws.Config, client *sqs.Client) {
	input := &sqs.StartMessageMoveTaskInput{
		// SourceArn: *string, // Required
	}

	if len(_sqsSourceArn) > 0 {
		input.SourceArn = aws.String(_sqsSourceArn)
	}
	if len(_sqsDestinationArn) > 0 {
		input.DestinationArn = aws.String(_sqsDestinationArn)
	}
	if len(_sqsMaxNumberOfMessagesPerSecond) > 0 {
		if err := assignInputField(input, "MaxNumberOfMessagesPerSecond", _sqsMaxNumberOfMessagesPerSecond); err != nil {
			log.Errorf("invalid --max-number-of-messages-per-second: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartMessageMoveTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add cost allocation tags to the specified Amazon SQS queue. For an overview,
// see [Tagging Your Amazon SQS Queues]in the Amazon SQS Developer Guide.
//
// When you use queue tags, keep the following guidelines in mind:
//
// - Adding more than 50 tags to a queue isn't recommended.
//
// - Tags don't have any semantic meaning. Amazon SQS interprets tags as
// character strings.
//
// - Tags are case-sensitive.
//
// - A new tag with a key identical to that of an existing tag overwrites the
// existing tag.
//
// For a full list of tag restrictions, see [Quotas related to queues] in the Amazon SQS Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Tagging Your Amazon SQS Queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html
// [Quotas related to queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_TagQueue(cfg aws.Config, client *sqs.Client) {
	input := &sqs.TagQueueInput{
		// QueueUrl: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsTags) > 0 {
		if err := assignInputField(input, "Tags", _sqsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove cost allocation tags from the specified Amazon SQS queue. For an
// overview, see [Tagging Your Amazon SQS Queues]in the Amazon SQS Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information, see [Grant cross-account permissions to a role and a username]
// in the Amazon SQS Developer Guide.
//
// [Tagging Your Amazon SQS Queues]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html
// [Grant cross-account permissions to a role and a username]: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name
func sqs_UntagQueue(cfg aws.Config, client *sqs.Client) {
	input := &sqs.UntagQueueInput{
		// QueueUrl: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_sqsQueueUrl) > 0 {
		input.QueueUrl = aws.String(_sqsQueueUrl)
	}
	if len(_sqsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _sqsTagKeys...)
	}

	if resp, err := client.UntagQueue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sqsCmd)
	_sqsCmd.Flags().SortFlags = false

	_sqsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_sqsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sqsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_sqsCmd.Flags().StringSliceVarP(&_sqsActions, "actions", "", nil, "Actions")
	_sqsCmd.Flags().StringVarP(&_sqsAttributeNames, "attribute-names", "", "", "Attribute Names")
	_sqsCmd.Flags().StringVarP(&_sqsAttributes, "attributes", "", "", "Attributes")
	_sqsCmd.Flags().StringSliceVarP(&_sqsAWSAccountIds, "aws-account-ids", "", nil, "AWS Account Ids")
	_sqsCmd.Flags().StringVarP(&_sqsDelaySeconds, "delay-seconds", "", "", "Delay Seconds")
	_sqsCmd.Flags().StringVarP(&_sqsDestinationArn, "destination-arn", "", "", "Destination ARN")
	_sqsCmd.Flags().StringVarP(&_sqsEntries, "entries", "", "", "Entries")
	_sqsCmd.Flags().StringVarP(&_sqsLabel, "label", "", "", "Label")
	_sqsCmd.Flags().StringVarP(&_sqsMaxNumberOfMessages, "max-number-of-messages", "", "", "Max Number Of Messages")
	_sqsCmd.Flags().StringVarP(&_sqsMaxNumberOfMessagesPerSecond, "max-number-of-messages-per-second", "", "", "Max Number Of Messages Per Second")
	_sqsCmd.Flags().StringVarP(&_sqsMaxResults, "max-results", "", "", "Max Results")
	_sqsCmd.Flags().StringSliceVarP(&_sqsMessageAttributeNames, "message-attribute-names", "", nil, "Message Attribute Names")
	_sqsCmd.Flags().StringVarP(&_sqsMessageAttributes, "message-attributes", "", "", "Message Attributes")
	_sqsCmd.Flags().StringVarP(&_sqsMessageBody, "message-body", "", "", "Message Body")
	_sqsCmd.Flags().StringVarP(&_sqsMessageDeduplicationId, "message-deduplication-id", "", "", "Message Deduplication ID")
	_sqsCmd.Flags().StringVarP(&_sqsMessageGroupId, "message-group-id", "", "", "Message Group ID")
	_sqsCmd.Flags().StringVarP(&_sqsMessageSystemAttributeNames, "message-system-attribute-names", "", "", "Message System Attribute Names")
	_sqsCmd.Flags().StringVarP(&_sqsMessageSystemAttributes, "message-system-attributes", "", "", "Message System Attributes")
	_sqsCmd.Flags().StringVarP(&_sqsNextToken, "next-token", "", "", "Next Token")
	_sqsCmd.Flags().StringVarP(&_sqsQueueName, "queue-name", "", "", "Queue Name")
	_sqsCmd.Flags().StringVarP(&_sqsQueueNamePrefix, "queue-name-prefix", "", "", "Queue Name Prefix")
	_sqsCmd.Flags().StringVarP(&_sqsQueueOwnerAWSAccountId, "queue-owner-aws-account-id", "", "", "Queue Owner AWS Account ID")
	_sqsCmd.Flags().StringVarP(&_sqsQueueUrl, "queue-url", "", "", "Queue URL")
	_sqsCmd.Flags().StringVarP(&_sqsReceiptHandle, "receipt-handle", "", "", "Receipt Handle")
	_sqsCmd.Flags().StringVarP(&_sqsReceiveRequestAttemptId, "receive-request-attempt-id", "", "", "Receive Request Attempt ID")
	_sqsCmd.Flags().StringVarP(&_sqsSourceArn, "source-arn", "", "", "Source ARN")
	_sqsCmd.Flags().StringSliceVarP(&_sqsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_sqsCmd.Flags().StringVarP(&_sqsTags, "tags", "", "", "Tags")
	_sqsCmd.Flags().StringVarP(&_sqsTaskHandle, "task-handle", "", "", "Task Handle")
	_sqsCmd.Flags().StringVarP(&_sqsVisibilityTimeout, "visibility-timeout", "", "", "Visibility Timeout")
	_sqsCmd.Flags().StringVarP(&_sqsWaitTimeSeconds, "wait-time-seconds", "", "", "Wait Time Seconds")

	_sqsCmd.Flags().BoolVarP(&_sqsAddPermission, "add-permission", "", false, "Add Permission")
	_sqsCmd.Flags().BoolVarP(&_sqsCancelMessageMoveTask, "cancel-message-move-task", "", false, "Cancel Message Move Task")
	_sqsCmd.Flags().BoolVarP(&_sqsChangeMessageVisibility, "change-message-visibility", "", false, "Change Message Visibility")
	_sqsCmd.Flags().BoolVarP(&_sqsChangeMessageVisibilityBatch, "change-message-visibility-batch", "", false, "Change Message Visibility Batch")
	_sqsCmd.Flags().BoolVarP(&_sqsCreateQueue, "create-queue", "", false, "Create Queue")
	_sqsCmd.Flags().BoolVarP(&_sqsDeleteMessage, "delete-message", "", false, "Delete Message")
	_sqsCmd.Flags().BoolVarP(&_sqsDeleteMessageBatch, "delete-message-batch", "", false, "Delete Message Batch")
	_sqsCmd.Flags().BoolVarP(&_sqsDeleteQueue, "delete-queue", "", false, "Delete Queue")
	_sqsCmd.Flags().BoolVarP(&_sqsGetQueueAttributes, "get-queue-attributes", "", false, "Get Queue Attributes")
	_sqsCmd.Flags().BoolVarP(&_sqsGetQueueUrl, "get-queue-url", "", false, "Get Queue URL")
	_sqsCmd.Flags().BoolVarP(&_sqsListDeadLetterSourceQueues, "list-dead-letter-source-queues", "", false, "List Dead Letter Source Queues")
	_sqsCmd.Flags().BoolVarP(&_sqsListMessageMoveTasks, "list-message-move-tasks", "", false, "List Message Move Tasks")
	_sqsCmd.Flags().BoolVarP(&_sqsListQueueTags, "list-queue-tags", "", false, "List Queue Tags")
	_sqsCmd.Flags().BoolVarP(&_sqsListQueues, "list-queues", "", false, "List Queues")
	_sqsCmd.Flags().BoolVarP(&_sqsPurgeQueue, "purge-queue", "", false, "Purge Queue")
	_sqsCmd.Flags().BoolVarP(&_sqsReceiveMessage, "receive-message", "", false, "Receive Message")
	_sqsCmd.Flags().BoolVarP(&_sqsRemovePermission, "remove-permission", "", false, "Remove Permission")
	_sqsCmd.Flags().BoolVarP(&_sqsSendMessage, "send-message", "", false, "Send Message")
	_sqsCmd.Flags().BoolVarP(&_sqsSendMessageBatch, "send-message-batch", "", false, "Send Message Batch")
	_sqsCmd.Flags().BoolVarP(&_sqsSetQueueAttributes, "set-queue-attributes", "", false, "Set Queue Attributes")
	_sqsCmd.Flags().BoolVarP(&_sqsStartMessageMoveTask, "start-message-move-task", "", false, "Start Message Move Task")
	_sqsCmd.Flags().BoolVarP(&_sqsTagQueue, "tag-queue", "", false, "Tag Queue")
	_sqsCmd.Flags().BoolVarP(&_sqsUntagQueue, "untag-queue", "", false, "Untag Queue")

}
