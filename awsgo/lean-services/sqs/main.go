package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/sqs"
)

var fields_add_permission = []leanruntime.Field{
	{Name: "AWSAccountIds", Flag: "aws-account-ids", Type: "[]string", Required: true},
	{Name: "Actions", Flag: "actions", Type: "[]string", Required: true},
	{Name: "Label", Flag: "label", Type: "*string", Required: true},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_cancel_message_move_task = []leanruntime.Field{
	{Name: "TaskHandle", Flag: "task-handle", Type: "*string", Required: true},
}

var fields_change_message_visibility = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
	{Name: "ReceiptHandle", Flag: "receipt-handle", Type: "*string", Required: true},
	{Name: "VisibilityTimeout", Flag: "visibility-timeout", Type: "int32", Required: true},
}

var fields_change_message_visibility_batch = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.ChangeMessageVisibilityBatchRequestEntry", Required: true},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_create_queue = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "QueueName", Flag: "queue-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_message = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
	{Name: "ReceiptHandle", Flag: "receipt-handle", Type: "*string", Required: true},
}

var fields_delete_message_batch = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.DeleteMessageBatchRequestEntry", Required: true},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_delete_queue = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_get_queue_attributes = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]types.QueueAttributeName", Required: false},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_get_queue_url = []leanruntime.Field{
	{Name: "QueueName", Flag: "queue-name", Type: "*string", Required: true},
	{Name: "QueueOwnerAWSAccountId", Flag: "queue-owner-aws-account-id", Type: "*string", Required: false},
}

var fields_list_dead_letter_source_queues = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_list_message_move_tasks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
}

var fields_list_queue_tags = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_list_queues = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueNamePrefix", Flag: "queue-name-prefix", Type: "*string", Required: false},
}

var fields_purge_queue = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_receive_message = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]types.QueueAttributeName", Required: false},
	{Name: "MaxNumberOfMessages", Flag: "max-number-of-messages", Type: "int32", Required: false},
	{Name: "MessageAttributeNames", Flag: "message-attribute-names", Type: "[]string", Required: false},
	{Name: "MessageSystemAttributeNames", Flag: "message-system-attribute-names", Type: "[]types.MessageSystemAttributeName", Required: false},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
	{Name: "ReceiveRequestAttemptId", Flag: "receive-request-attempt-id", Type: "*string", Required: false},
	{Name: "VisibilityTimeout", Flag: "visibility-timeout", Type: "int32", Required: false},
	{Name: "WaitTimeSeconds", Flag: "wait-time-seconds", Type: "int32", Required: false},
}

var fields_remove_permission = []leanruntime.Field{
	{Name: "Label", Flag: "label", Type: "*string", Required: true},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_send_message = []leanruntime.Field{
	{Name: "DelaySeconds", Flag: "delay-seconds", Type: "int32", Required: false},
	{Name: "MessageAttributes", Flag: "message-attributes", Type: "map[string]types.MessageAttributeValue", Required: false},
	{Name: "MessageBody", Flag: "message-body", Type: "*string", Required: true},
	{Name: "MessageDeduplicationId", Flag: "message-deduplication-id", Type: "*string", Required: false},
	{Name: "MessageGroupId", Flag: "message-group-id", Type: "*string", Required: false},
	{Name: "MessageSystemAttributes", Flag: "message-system-attributes", Type: "map[string]types.MessageSystemAttributeValue", Required: false},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_send_message_batch = []leanruntime.Field{
	{Name: "Entries", Flag: "entries", Type: "[]types.SendMessageBatchRequestEntry", Required: true},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_set_queue_attributes = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: true},
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
}

var fields_start_message_move_task = []leanruntime.Field{
	{Name: "DestinationArn", Flag: "destination-arn", Type: "*string", Required: false},
	{Name: "MaxNumberOfMessagesPerSecond", Flag: "max-number-of-messages-per-second", Type: "*int32", Required: false},
	{Name: "SourceArn", Flag: "source-arn", Type: "*string", Required: true},
}

var fields_tag_queue = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_queue = []leanruntime.Field{
	{Name: "QueueUrl", Flag: "queue-url", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-permission": {
			Name:   "add-permission",
			Fields: fields_add_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddPermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddPermission(ctx, input)
			},
		},
		"cancel-message-move-task": {
			Name:   "cancel-message-move-task",
			Fields: fields_cancel_message_move_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMessageMoveTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_message_move_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMessageMoveTask(ctx, input)
			},
		},
		"change-message-visibility": {
			Name:   "change-message-visibility",
			Fields: fields_change_message_visibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangeMessageVisibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_message_visibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangeMessageVisibility(ctx, input)
			},
		},
		"change-message-visibility-batch": {
			Name:   "change-message-visibility-batch",
			Fields: fields_change_message_visibility_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangeMessageVisibilityBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_message_visibility_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangeMessageVisibilityBatch(ctx, input)
			},
		},
		"create-queue": {
			Name:   "create-queue",
			Fields: fields_create_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueue(ctx, input)
			},
		},
		"delete-message": {
			Name:   "delete-message",
			Fields: fields_delete_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMessage(ctx, input)
			},
		},
		"delete-message-batch": {
			Name:   "delete-message-batch",
			Fields: fields_delete_message_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMessageBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_message_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMessageBatch(ctx, input)
			},
		},
		"delete-queue": {
			Name:   "delete-queue",
			Fields: fields_delete_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueue(ctx, input)
			},
		},
		"get-queue-attributes": {
			Name:   "get-queue-attributes",
			Fields: fields_get_queue_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueueAttributes(ctx, input)
			},
		},
		"get-queue-url": {
			Name:   "get-queue-url",
			Fields: fields_get_queue_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueueUrl(ctx, input)
			},
		},
		"list-dead-letter-source-queues": {
			Name:   "list-dead-letter-source-queues",
			Fields: fields_list_dead_letter_source_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDeadLetterSourceQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_dead_letter_source_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDeadLetterSourceQueues(ctx, input)
				}
				var results []*svc.ListDeadLetterSourceQueuesOutput
				p := svc.NewListDeadLetterSourceQueuesPaginator(client, input)
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
		"list-message-move-tasks": {
			Name:   "list-message-move-tasks",
			Fields: fields_list_message_move_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMessageMoveTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_message_move_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMessageMoveTasks(ctx, input)
			},
		},
		"list-queue-tags": {
			Name:   "list-queue-tags",
			Fields: fields_list_queue_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueueTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_queue_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListQueueTags(ctx, input)
			},
		},
		"list-queues": {
			Name:   "list-queues",
			Fields: fields_list_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueues(ctx, input)
				}
				var results []*svc.ListQueuesOutput
				p := svc.NewListQueuesPaginator(client, input)
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
		"purge-queue": {
			Name:   "purge-queue",
			Fields: fields_purge_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurgeQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purge_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurgeQueue(ctx, input)
			},
		},
		"receive-message": {
			Name:   "receive-message",
			Fields: fields_receive_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ReceiveMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_receive_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ReceiveMessage(ctx, input)
			},
		},
		"remove-permission": {
			Name:   "remove-permission",
			Fields: fields_remove_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemovePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemovePermission(ctx, input)
			},
		},
		"send-message": {
			Name:   "send-message",
			Fields: fields_send_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendMessage(ctx, input)
			},
		},
		"send-message-batch": {
			Name:   "send-message-batch",
			Fields: fields_send_message_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendMessageBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_message_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendMessageBatch(ctx, input)
			},
		},
		"set-queue-attributes": {
			Name:   "set-queue-attributes",
			Fields: fields_set_queue_attributes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SetQueueAttributesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_set_queue_attributes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SetQueueAttributes(ctx, input)
			},
		},
		"start-message-move-task": {
			Name:   "start-message-move-task",
			Fields: fields_start_message_move_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMessageMoveTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_message_move_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMessageMoveTask(ctx, input)
			},
		},
		"tag-queue": {
			Name:   "tag-queue",
			Fields: fields_tag_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagQueue(ctx, input)
			},
		},
		"untag-queue": {
			Name:   "untag-queue",
			Fields: fields_untag_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagQueue(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("sqs", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
