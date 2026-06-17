package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/socialmessaging/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"associate-whats-app-business-account", "create-whats-app-message-template", "create-whats-app-message-template-from-library", "create-whats-app-message-template-media", "delete-whats-app-message-media", "delete-whats-app-message-template", "disassociate-whats-app-business-account", "get-linked-whats-app-business-account", "get-linked-whats-app-business-account-phone-number", "get-whats-app-message-media", "get-whats-app-message-template", "list-linked-whats-app-business-accounts", "list-tags-for-resource", "list-whats-app-message-templates", "list-whats-app-template-library", "post-whats-app-message-media", "put-whats-app-business-account-event-destinations", "send-whats-app-message", "tag-resource", "untag-resource", "update-whats-app-message-template"},
		OperationSet: map[string]bool{"associate-whats-app-business-account": true, "create-whats-app-message-template": true, "create-whats-app-message-template-from-library": true, "create-whats-app-message-template-media": true, "delete-whats-app-message-media": true, "delete-whats-app-message-template": true, "disassociate-whats-app-business-account": true, "get-linked-whats-app-business-account": true, "get-linked-whats-app-business-account-phone-number": true, "get-whats-app-message-media": true, "get-whats-app-message-template": true, "list-linked-whats-app-business-accounts": true, "list-tags-for-resource": true, "list-whats-app-message-templates": true, "list-whats-app-template-library": true, "post-whats-app-message-media": true, "put-whats-app-business-account-event-destinations": true, "send-whats-app-message": true, "tag-resource": true, "untag-resource": true, "update-whats-app-message-template": true},
		OperationInputs: map[string][]string{
			"associate-whats-app-business-account":               {"SetupFinalization", "SignupCallback"},
			"create-whats-app-message-template":                  {"Id", "TemplateDefinition"},
			"create-whats-app-message-template-from-library":     {"Id", "MetaLibraryTemplate"},
			"create-whats-app-message-template-media":            {"Id", "SourceS3File"},
			"delete-whats-app-message-media":                     {"MediaId", "OriginationPhoneNumberId"},
			"delete-whats-app-message-template":                  {"DeleteAllLanguages", "Id", "MetaTemplateId", "TemplateName"},
			"disassociate-whats-app-business-account":            {"Id"},
			"get-linked-whats-app-business-account":              {"Id"},
			"get-linked-whats-app-business-account-phone-number": {"Id"},
			"get-whats-app-message-media":                        {"DestinationS3File", "DestinationS3PresignedUrl", "MediaId", "MetadataOnly", "OriginationPhoneNumberId"},
			"get-whats-app-message-template":                     {"Id", "MetaTemplateId"},
			"list-linked-whats-app-business-accounts":            {"MaxResults", "NextToken"},
			"list-tags-for-resource":                             {"ResourceArn"},
			"list-whats-app-message-templates":                   {"Id", "MaxResults", "NextToken"},
			"list-whats-app-template-library":                    {"Filters", "Id", "MaxResults", "NextToken"},
			"post-whats-app-message-media":                       {"OriginationPhoneNumberId", "SourceS3File", "SourceS3PresignedUrl"},
			"put-whats-app-business-account-event-destinations":  {"EventDestinations", "Id"},
			"send-whats-app-message":                             {"Message", "MetaApiVersion", "OriginationPhoneNumberId"},
			"tag-resource":                                       {"ResourceArn", "Tags"},
			"untag-resource":                                     {"ResourceArn", "TagKeys"},
			"update-whats-app-message-template":                  {"CtaUrlLinkTrackingOptedOut", "Id", "MetaTemplateId", "ParameterFormat", "TemplateCategory", "TemplateComponents"},
		},
		OperationInputTypes: map[string]map[string]string{
			"associate-whats-app-business-account":               {"SetupFinalization": "*types.WhatsAppSetupFinalization", "SignupCallback": "*types.WhatsAppSignupCallback"},
			"create-whats-app-message-template":                  {"Id": "*string", "TemplateDefinition": "[]byte"},
			"create-whats-app-message-template-from-library":     {"Id": "*string", "MetaLibraryTemplate": "*types.MetaLibraryTemplate"},
			"create-whats-app-message-template-media":            {"Id": "*string", "SourceS3File": "*types.S3File"},
			"delete-whats-app-message-media":                     {"MediaId": "*string", "OriginationPhoneNumberId": "*string"},
			"delete-whats-app-message-template":                  {"DeleteAllLanguages": "*bool", "Id": "*string", "MetaTemplateId": "*string", "TemplateName": "*string"},
			"disassociate-whats-app-business-account":            {"Id": "*string"},
			"get-linked-whats-app-business-account":              {"Id": "*string"},
			"get-linked-whats-app-business-account-phone-number": {"Id": "*string"},
			"get-whats-app-message-media":                        {"DestinationS3File": "*types.S3File", "DestinationS3PresignedUrl": "*types.S3PresignedUrl", "MediaId": "*string", "MetadataOnly": "*bool", "OriginationPhoneNumberId": "*string"},
			"get-whats-app-message-template":                     {"Id": "*string", "MetaTemplateId": "*string"},
			"list-linked-whats-app-business-accounts":            {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                             {"ResourceArn": "*string"},
			"list-whats-app-message-templates":                   {"Id": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-whats-app-template-library":                    {"Filters": "map[string]string", "Id": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"post-whats-app-message-media":                       {"OriginationPhoneNumberId": "*string", "SourceS3File": "*types.S3File", "SourceS3PresignedUrl": "*types.S3PresignedUrl"},
			"put-whats-app-business-account-event-destinations":  {"EventDestinations": "[]types.WhatsAppBusinessAccountEventDestination", "Id": "*string"},
			"send-whats-app-message":                             {"Message": "[]byte", "MetaApiVersion": "*string", "OriginationPhoneNumberId": "*string"},
			"tag-resource":                                       {"ResourceArn": "*string", "Tags": "[]types.Tag"},
			"untag-resource":                                     {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-whats-app-message-template":                  {"CtaUrlLinkTrackingOptedOut": "*bool", "Id": "*string", "MetaTemplateId": "*string", "ParameterFormat": "*string", "TemplateCategory": "*string", "TemplateComponents": "[]byte"},
		},
		OperationInputRequired: map[string][]string{
			"associate-whats-app-business-account":               {},
			"create-whats-app-message-template":                  {"Id", "TemplateDefinition"},
			"create-whats-app-message-template-from-library":     {"Id", "MetaLibraryTemplate"},
			"create-whats-app-message-template-media":            {"Id"},
			"delete-whats-app-message-media":                     {"MediaId", "OriginationPhoneNumberId"},
			"delete-whats-app-message-template":                  {"Id", "TemplateName"},
			"disassociate-whats-app-business-account":            {"Id"},
			"get-linked-whats-app-business-account":              {"Id"},
			"get-linked-whats-app-business-account-phone-number": {"Id"},
			"get-whats-app-message-media":                        {"MediaId", "OriginationPhoneNumberId"},
			"get-whats-app-message-template":                     {"Id", "MetaTemplateId"},
			"list-linked-whats-app-business-accounts":            {},
			"list-tags-for-resource":                             {"ResourceArn"},
			"list-whats-app-message-templates":                   {"Id"},
			"list-whats-app-template-library":                    {"Id"},
			"post-whats-app-message-media":                       {"OriginationPhoneNumberId"},
			"put-whats-app-business-account-event-destinations":  {"EventDestinations", "Id"},
			"send-whats-app-message":                             {"Message", "MetaApiVersion", "OriginationPhoneNumberId"},
			"tag-resource":                                       {"ResourceArn", "Tags"},
			"untag-resource":                                     {"ResourceArn", "TagKeys"},
			"update-whats-app-message-template":                  {"Id", "MetaTemplateId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("socialmessaging", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
