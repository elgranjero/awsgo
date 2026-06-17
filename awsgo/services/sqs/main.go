package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/sqs/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"add-permission", "cancel-message-move-task", "change-message-visibility", "change-message-visibility-batch", "create-queue", "delete-message", "delete-message-batch", "delete-queue", "get-queue-attributes", "get-queue-url", "list-dead-letter-source-queues", "list-message-move-tasks", "list-queue-tags", "list-queues", "purge-queue", "receive-message", "remove-permission", "send-message", "send-message-batch", "set-queue-attributes", "start-message-move-task", "tag-queue", "untag-queue"},
		OperationSet: map[string]bool{"add-permission": true, "cancel-message-move-task": true, "change-message-visibility": true, "change-message-visibility-batch": true, "create-queue": true, "delete-message": true, "delete-message-batch": true, "delete-queue": true, "get-queue-attributes": true, "get-queue-url": true, "list-dead-letter-source-queues": true, "list-message-move-tasks": true, "list-queue-tags": true, "list-queues": true, "purge-queue": true, "receive-message": true, "remove-permission": true, "send-message": true, "send-message-batch": true, "set-queue-attributes": true, "start-message-move-task": true, "tag-queue": true, "untag-queue": true},
		OperationInputs: map[string][]string{
			"add-permission":                  {"AWSAccountIds", "Actions", "Label", "QueueUrl"},
			"cancel-message-move-task":        {"TaskHandle"},
			"change-message-visibility":       {"QueueUrl", "ReceiptHandle", "VisibilityTimeout"},
			"change-message-visibility-batch": {"Entries", "QueueUrl"},
			"create-queue":                    {"Attributes", "QueueName", "Tags"},
			"delete-message":                  {"QueueUrl", "ReceiptHandle"},
			"delete-message-batch":            {"Entries", "QueueUrl"},
			"delete-queue":                    {"QueueUrl"},
			"get-queue-attributes":            {"AttributeNames", "QueueUrl"},
			"get-queue-url":                   {"QueueName", "QueueOwnerAWSAccountId"},
			"list-dead-letter-source-queues":  {"MaxResults", "NextToken", "QueueUrl"},
			"list-message-move-tasks":         {"MaxResults", "SourceArn"},
			"list-queue-tags":                 {"QueueUrl"},
			"list-queues":                     {"MaxResults", "NextToken", "QueueNamePrefix"},
			"purge-queue":                     {"QueueUrl"},
			"receive-message":                 {"AttributeNames", "MaxNumberOfMessages", "MessageAttributeNames", "MessageSystemAttributeNames", "QueueUrl", "ReceiveRequestAttemptId", "VisibilityTimeout", "WaitTimeSeconds"},
			"remove-permission":               {"Label", "QueueUrl"},
			"send-message":                    {"DelaySeconds", "MessageAttributes", "MessageBody", "MessageDeduplicationId", "MessageGroupId", "MessageSystemAttributes", "QueueUrl"},
			"send-message-batch":              {"Entries", "QueueUrl"},
			"set-queue-attributes":            {"Attributes", "QueueUrl"},
			"start-message-move-task":         {"DestinationArn", "MaxNumberOfMessagesPerSecond", "SourceArn"},
			"tag-queue":                       {"QueueUrl", "Tags"},
			"untag-queue":                     {"QueueUrl", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"add-permission":                  {"AWSAccountIds": "[]string", "Actions": "[]string", "Label": "*string", "QueueUrl": "*string"},
			"cancel-message-move-task":        {"TaskHandle": "*string"},
			"change-message-visibility":       {"QueueUrl": "*string", "ReceiptHandle": "*string", "VisibilityTimeout": "int32"},
			"change-message-visibility-batch": {"Entries": "[]types.ChangeMessageVisibilityBatchRequestEntry", "QueueUrl": "*string"},
			"create-queue":                    {"Attributes": "map[string]string", "QueueName": "*string", "Tags": "map[string]string"},
			"delete-message":                  {"QueueUrl": "*string", "ReceiptHandle": "*string"},
			"delete-message-batch":            {"Entries": "[]types.DeleteMessageBatchRequestEntry", "QueueUrl": "*string"},
			"delete-queue":                    {"QueueUrl": "*string"},
			"get-queue-attributes":            {"AttributeNames": "[]types.QueueAttributeName", "QueueUrl": "*string"},
			"get-queue-url":                   {"QueueName": "*string", "QueueOwnerAWSAccountId": "*string"},
			"list-dead-letter-source-queues":  {"MaxResults": "*int32", "NextToken": "*string", "QueueUrl": "*string"},
			"list-message-move-tasks":         {"MaxResults": "*int32", "SourceArn": "*string"},
			"list-queue-tags":                 {"QueueUrl": "*string"},
			"list-queues":                     {"MaxResults": "*int32", "NextToken": "*string", "QueueNamePrefix": "*string"},
			"purge-queue":                     {"QueueUrl": "*string"},
			"receive-message":                 {"AttributeNames": "[]types.QueueAttributeName", "MaxNumberOfMessages": "int32", "MessageAttributeNames": "[]string", "MessageSystemAttributeNames": "[]types.MessageSystemAttributeName", "QueueUrl": "*string", "ReceiveRequestAttemptId": "*string", "VisibilityTimeout": "int32", "WaitTimeSeconds": "int32"},
			"remove-permission":               {"Label": "*string", "QueueUrl": "*string"},
			"send-message":                    {"DelaySeconds": "int32", "MessageAttributes": "map[string]types.MessageAttributeValue", "MessageBody": "*string", "MessageDeduplicationId": "*string", "MessageGroupId": "*string", "MessageSystemAttributes": "map[string]types.MessageSystemAttributeValue", "QueueUrl": "*string"},
			"send-message-batch":              {"Entries": "[]types.SendMessageBatchRequestEntry", "QueueUrl": "*string"},
			"set-queue-attributes":            {"Attributes": "map[string]string", "QueueUrl": "*string"},
			"start-message-move-task":         {"DestinationArn": "*string", "MaxNumberOfMessagesPerSecond": "*int32", "SourceArn": "*string"},
			"tag-queue":                       {"QueueUrl": "*string", "Tags": "map[string]string"},
			"untag-queue":                     {"QueueUrl": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"add-permission":                  {"AWSAccountIds", "Actions", "Label", "QueueUrl"},
			"cancel-message-move-task":        {"TaskHandle"},
			"change-message-visibility":       {"QueueUrl", "ReceiptHandle", "VisibilityTimeout"},
			"change-message-visibility-batch": {"Entries", "QueueUrl"},
			"create-queue":                    {"QueueName"},
			"delete-message":                  {"QueueUrl", "ReceiptHandle"},
			"delete-message-batch":            {"Entries", "QueueUrl"},
			"delete-queue":                    {"QueueUrl"},
			"get-queue-attributes":            {"QueueUrl"},
			"get-queue-url":                   {"QueueName"},
			"list-dead-letter-source-queues":  {"QueueUrl"},
			"list-message-move-tasks":         {"SourceArn"},
			"list-queue-tags":                 {"QueueUrl"},
			"list-queues":                     {},
			"purge-queue":                     {"QueueUrl"},
			"receive-message":                 {"QueueUrl"},
			"remove-permission":               {"Label", "QueueUrl"},
			"send-message":                    {"MessageBody", "QueueUrl"},
			"send-message-batch":              {"Entries", "QueueUrl"},
			"set-queue-attributes":            {"Attributes", "QueueUrl"},
			"start-message-move-task":         {"SourceArn"},
			"tag-queue":                       {"QueueUrl", "Tags"},
			"untag-queue":                     {"QueueUrl", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("sqs", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
