package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/chime"
)

var fields_associate_phone_number_with_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "E164PhoneNumber", Flag: "e164-phone-number", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_associate_signin_delegate_groups_with_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "SigninDelegateGroups", Flag: "signin-delegate-groups", Type: "[]types.SigninDelegateGroup", Required: true},
}

var fields_batch_create_room_membership = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MembershipItemList", Flag: "membership-item-list", Type: "[]types.MembershipItem", Required: true},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_batch_delete_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberIds", Flag: "phone-number-ids", Type: "[]string", Required: true},
}

var fields_batch_suspend_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserIdList", Flag: "user-id-list", Type: "[]string", Required: true},
}

var fields_batch_unsuspend_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserIdList", Flag: "user-id-list", Type: "[]string", Required: true},
}

var fields_batch_update_phone_number = []leanruntime.Field{
	{Name: "UpdatePhoneNumberRequestItems", Flag: "update-phone-number-request-items", Type: "[]types.UpdatePhoneNumberRequestItem", Required: true},
}

var fields_batch_update_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UpdateUserRequestItems", Flag: "update-user-request-items", Type: "[]types.UpdateUserRequestItem", Required: true},
}

var fields_create_account = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_bot = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "Domain", Flag: "domain", Type: "*string", Required: false},
}

var fields_create_meeting_dial_out = []leanruntime.Field{
	{Name: "FromPhoneNumber", Flag: "from-phone-number", Type: "*string", Required: true},
	{Name: "JoinToken", Flag: "join-token", Type: "*string", Required: true},
	{Name: "MeetingId", Flag: "meeting-id", Type: "*string", Required: true},
	{Name: "ToPhoneNumber", Flag: "to-phone-number", Type: "*string", Required: true},
}

var fields_create_phone_number_order = []leanruntime.Field{
	{Name: "E164PhoneNumbers", Flag: "e164-phone-numbers", Type: "[]string", Required: true},
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: true},
}

var fields_create_room = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_room_membership = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.RoomMembershipRole", Required: false},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_create_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "UserType", Flag: "user-type", Type: "types.UserType", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: false},
}

var fields_delete_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_delete_events_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_delete_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_delete_room = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_delete_room_membership = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_disassociate_phone_number_from_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_disassociate_signin_delegate_groups_from_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "GroupNames", Flag: "group-names", Type: "[]string", Required: true},
}

var fields_get_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_account_settings = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_bot = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_get_events_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_get_global_settings = []leanruntime.Field{}

var fields_get_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_get_phone_number_order = []leanruntime.Field{
	{Name: "PhoneNumberOrderId", Flag: "phone-number-order-id", Type: "*string", Required: true},
}

var fields_get_phone_number_settings = []leanruntime.Field{}

var fields_get_retention_settings = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_get_room = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_get_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_user_settings = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_invite_users = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserEmailList", Flag: "user-email-list", Type: "[]string", Required: true},
	{Name: "UserType", Flag: "user-type", Type: "types.UserType", Required: false},
}

var fields_list_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserEmail", Flag: "user-email", Type: "*string", Required: false},
}

var fields_list_bots = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_phone_number_orders = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_phone_numbers = []leanruntime.Field{
	{Name: "FilterName", Flag: "filter-name", Type: "types.PhoneNumberAssociationName", Required: false},
	{Name: "FilterValue", Flag: "filter-value", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: false},
	{Name: "Status", Flag: "status", Type: "types.PhoneNumberStatus", Required: false},
}

var fields_list_room_memberships = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_list_rooms = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_supported_phone_number_countries = []leanruntime.Field{
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserEmail", Flag: "user-email", Type: "*string", Required: false},
	{Name: "UserType", Flag: "user-type", Type: "types.UserType", Required: false},
}

var fields_logout_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_put_events_configuration = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "LambdaFunctionArn", Flag: "lambda-function-arn", Type: "*string", Required: false},
	{Name: "OutboundEventsHTTPSEndpoint", Flag: "outbound-events-https-endpoint", Type: "*string", Required: false},
}

var fields_put_retention_settings = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "RetentionSettings", Flag: "retention-settings", Type: "*types.RetentionSettings", Required: true},
}

var fields_redact_conversation_message = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ConversationId", Flag: "conversation-id", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
}

var fields_redact_room_message = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MessageId", Flag: "message-id", Type: "*string", Required: true},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_regenerate_security_token = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
}

var fields_reset_personal_pin = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_restore_phone_number = []leanruntime.Field{
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
}

var fields_search_available_phone_numbers = []leanruntime.Field{
	{Name: "AreaCode", Flag: "area-code", Type: "*string", Required: false},
	{Name: "City", Flag: "city", Type: "*string", Required: false},
	{Name: "Country", Flag: "country", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PhoneNumberType", Flag: "phone-number-type", Type: "types.PhoneNumberType", Required: false},
	{Name: "State", Flag: "state", Type: "*string", Required: false},
	{Name: "TollFreePrefix", Flag: "toll-free-prefix", Type: "*string", Required: false},
}

var fields_update_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "DefaultLicense", Flag: "default-license", Type: "types.License", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_account_settings = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "AccountSettings", Flag: "account-settings", Type: "*types.AccountSettings", Required: true},
}

var fields_update_bot = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "BotId", Flag: "bot-id", Type: "*string", Required: true},
	{Name: "Disabled", Flag: "disabled", Type: "*bool", Required: false},
}

var fields_update_global_settings = []leanruntime.Field{
	{Name: "BusinessCalling", Flag: "business-calling", Type: "*types.BusinessCallingSettings", Required: false},
	{Name: "VoiceConnector", Flag: "voice-connector", Type: "*types.VoiceConnectorSettings", Required: false},
}

var fields_update_phone_number = []leanruntime.Field{
	{Name: "CallingName", Flag: "calling-name", Type: "*string", Required: false},
	{Name: "PhoneNumberId", Flag: "phone-number-id", Type: "*string", Required: true},
	{Name: "ProductType", Flag: "product-type", Type: "types.PhoneNumberProductType", Required: false},
}

var fields_update_phone_number_settings = []leanruntime.Field{
	{Name: "CallingName", Flag: "calling-name", Type: "*string", Required: true},
}

var fields_update_room = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_update_room_membership = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "Role", Flag: "role", Type: "types.RoomMembershipRole", Required: false},
	{Name: "RoomId", Flag: "room-id", Type: "*string", Required: true},
}

var fields_update_user = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "AlexaForBusinessMetadata", Flag: "alexa-for-business-metadata", Type: "*types.AlexaForBusinessMetadata", Required: false},
	{Name: "LicenseType", Flag: "license-type", Type: "types.License", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserType", Flag: "user-type", Type: "types.UserType", Required: false},
}

var fields_update_user_settings = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
	{Name: "UserSettings", Flag: "user-settings", Type: "*types.UserSettings", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-phone-number-with-user": {
			Name:   "associate-phone-number-with-user",
			Fields: fields_associate_phone_number_with_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePhoneNumberWithUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_phone_number_with_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePhoneNumberWithUser(ctx, input)
			},
		},
		"associate-signin-delegate-groups-with-account": {
			Name:   "associate-signin-delegate-groups-with-account",
			Fields: fields_associate_signin_delegate_groups_with_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateSigninDelegateGroupsWithAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_signin_delegate_groups_with_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateSigninDelegateGroupsWithAccount(ctx, input)
			},
		},
		"batch-create-room-membership": {
			Name:   "batch-create-room-membership",
			Fields: fields_batch_create_room_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateRoomMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_room_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateRoomMembership(ctx, input)
			},
		},
		"batch-delete-phone-number": {
			Name:   "batch-delete-phone-number",
			Fields: fields_batch_delete_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeletePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeletePhoneNumber(ctx, input)
			},
		},
		"batch-suspend-user": {
			Name:   "batch-suspend-user",
			Fields: fields_batch_suspend_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchSuspendUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_suspend_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchSuspendUser(ctx, input)
			},
		},
		"batch-unsuspend-user": {
			Name:   "batch-unsuspend-user",
			Fields: fields_batch_unsuspend_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUnsuspendUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_unsuspend_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUnsuspendUser(ctx, input)
			},
		},
		"batch-update-phone-number": {
			Name:   "batch-update-phone-number",
			Fields: fields_batch_update_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdatePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdatePhoneNumber(ctx, input)
			},
		},
		"batch-update-user": {
			Name:   "batch-update-user",
			Fields: fields_batch_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateUser(ctx, input)
			},
		},
		"create-account": {
			Name:   "create-account",
			Fields: fields_create_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccount(ctx, input)
			},
		},
		"create-bot": {
			Name:   "create-bot",
			Fields: fields_create_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBot(ctx, input)
			},
		},
		"create-meeting-dial-out": {
			Name:   "create-meeting-dial-out",
			Fields: fields_create_meeting_dial_out,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMeetingDialOutInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_meeting_dial_out, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMeetingDialOut(ctx, input)
			},
		},
		"create-phone-number-order": {
			Name:   "create-phone-number-order",
			Fields: fields_create_phone_number_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePhoneNumberOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_phone_number_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePhoneNumberOrder(ctx, input)
			},
		},
		"create-room": {
			Name:   "create-room",
			Fields: fields_create_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoom(ctx, input)
			},
		},
		"create-room-membership": {
			Name:   "create-room-membership",
			Fields: fields_create_room_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRoomMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_room_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRoomMembership(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"delete-account": {
			Name:   "delete-account",
			Fields: fields_delete_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccount(ctx, input)
			},
		},
		"delete-events-configuration": {
			Name:   "delete-events-configuration",
			Fields: fields_delete_events_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEventsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_events_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEventsConfiguration(ctx, input)
			},
		},
		"delete-phone-number": {
			Name:   "delete-phone-number",
			Fields: fields_delete_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePhoneNumber(ctx, input)
			},
		},
		"delete-room": {
			Name:   "delete-room",
			Fields: fields_delete_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoom(ctx, input)
			},
		},
		"delete-room-membership": {
			Name:   "delete-room-membership",
			Fields: fields_delete_room_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRoomMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_room_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRoomMembership(ctx, input)
			},
		},
		"disassociate-phone-number-from-user": {
			Name:   "disassociate-phone-number-from-user",
			Fields: fields_disassociate_phone_number_from_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePhoneNumberFromUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_phone_number_from_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePhoneNumberFromUser(ctx, input)
			},
		},
		"disassociate-signin-delegate-groups-from-account": {
			Name:   "disassociate-signin-delegate-groups-from-account",
			Fields: fields_disassociate_signin_delegate_groups_from_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateSigninDelegateGroupsFromAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_signin_delegate_groups_from_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateSigninDelegateGroupsFromAccount(ctx, input)
			},
		},
		"get-account": {
			Name:   "get-account",
			Fields: fields_get_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccount(ctx, input)
			},
		},
		"get-account-settings": {
			Name:   "get-account-settings",
			Fields: fields_get_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountSettings(ctx, input)
			},
		},
		"get-bot": {
			Name:   "get-bot",
			Fields: fields_get_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBot(ctx, input)
			},
		},
		"get-events-configuration": {
			Name:   "get-events-configuration",
			Fields: fields_get_events_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEventsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_events_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEventsConfiguration(ctx, input)
			},
		},
		"get-global-settings": {
			Name:   "get-global-settings",
			Fields: fields_get_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlobalSettings(ctx, input)
			},
		},
		"get-phone-number": {
			Name:   "get-phone-number",
			Fields: fields_get_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPhoneNumber(ctx, input)
			},
		},
		"get-phone-number-order": {
			Name:   "get-phone-number-order",
			Fields: fields_get_phone_number_order,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPhoneNumberOrderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_phone_number_order, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPhoneNumberOrder(ctx, input)
			},
		},
		"get-phone-number-settings": {
			Name:   "get-phone-number-settings",
			Fields: fields_get_phone_number_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPhoneNumberSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_phone_number_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPhoneNumberSettings(ctx, input)
			},
		},
		"get-retention-settings": {
			Name:   "get-retention-settings",
			Fields: fields_get_retention_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRetentionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_retention_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRetentionSettings(ctx, input)
			},
		},
		"get-room": {
			Name:   "get-room",
			Fields: fields_get_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRoom(ctx, input)
			},
		},
		"get-user": {
			Name:   "get-user",
			Fields: fields_get_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUser(ctx, input)
			},
		},
		"get-user-settings": {
			Name:   "get-user-settings",
			Fields: fields_get_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserSettings(ctx, input)
			},
		},
		"invite-users": {
			Name:   "invite-users",
			Fields: fields_invite_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InviteUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invite_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InviteUsers(ctx, input)
			},
		},
		"list-accounts": {
			Name:   "list-accounts",
			Fields: fields_list_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccounts(ctx, input)
				}
				var results []*svc.ListAccountsOutput
				p := svc.NewListAccountsPaginator(client, input)
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
		"list-bots": {
			Name:   "list-bots",
			Fields: fields_list_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBots(ctx, input)
				}
				var results []*svc.ListBotsOutput
				p := svc.NewListBotsPaginator(client, input)
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
		"list-phone-number-orders": {
			Name:   "list-phone-number-orders",
			Fields: fields_list_phone_number_orders,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPhoneNumberOrdersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_phone_number_orders, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPhoneNumberOrders(ctx, input)
				}
				var results []*svc.ListPhoneNumberOrdersOutput
				p := svc.NewListPhoneNumberOrdersPaginator(client, input)
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
		"list-phone-numbers": {
			Name:   "list-phone-numbers",
			Fields: fields_list_phone_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPhoneNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_phone_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPhoneNumbers(ctx, input)
				}
				var results []*svc.ListPhoneNumbersOutput
				p := svc.NewListPhoneNumbersPaginator(client, input)
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
		"list-room-memberships": {
			Name:   "list-room-memberships",
			Fields: fields_list_room_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoomMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_room_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoomMemberships(ctx, input)
				}
				var results []*svc.ListRoomMembershipsOutput
				p := svc.NewListRoomMembershipsPaginator(client, input)
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
		"list-rooms": {
			Name:   "list-rooms",
			Fields: fields_list_rooms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRoomsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rooms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRooms(ctx, input)
				}
				var results []*svc.ListRoomsOutput
				p := svc.NewListRoomsPaginator(client, input)
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
		"list-supported-phone-number-countries": {
			Name:   "list-supported-phone-number-countries",
			Fields: fields_list_supported_phone_number_countries,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSupportedPhoneNumberCountriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_supported_phone_number_countries, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSupportedPhoneNumberCountries(ctx, input)
			},
		},
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_users, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListUsers(ctx, input)
				}
				var results []*svc.ListUsersOutput
				p := svc.NewListUsersPaginator(client, input)
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
		"logout-user": {
			Name:   "logout-user",
			Fields: fields_logout_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LogoutUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_logout_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.LogoutUser(ctx, input)
			},
		},
		"put-events-configuration": {
			Name:   "put-events-configuration",
			Fields: fields_put_events_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEventsConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_events_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEventsConfiguration(ctx, input)
			},
		},
		"put-retention-settings": {
			Name:   "put-retention-settings",
			Fields: fields_put_retention_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRetentionSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_retention_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRetentionSettings(ctx, input)
			},
		},
		"redact-conversation-message": {
			Name:   "redact-conversation-message",
			Fields: fields_redact_conversation_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RedactConversationMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_redact_conversation_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RedactConversationMessage(ctx, input)
			},
		},
		"redact-room-message": {
			Name:   "redact-room-message",
			Fields: fields_redact_room_message,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RedactRoomMessageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_redact_room_message, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RedactRoomMessage(ctx, input)
			},
		},
		"regenerate-security-token": {
			Name:   "regenerate-security-token",
			Fields: fields_regenerate_security_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegenerateSecurityTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_regenerate_security_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegenerateSecurityToken(ctx, input)
			},
		},
		"reset-personal-pin": {
			Name:   "reset-personal-pin",
			Fields: fields_reset_personal_pin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetPersonalPINInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_personal_pin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetPersonalPIN(ctx, input)
			},
		},
		"restore-phone-number": {
			Name:   "restore-phone-number",
			Fields: fields_restore_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RestorePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_restore_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RestorePhoneNumber(ctx, input)
			},
		},
		"search-available-phone-numbers": {
			Name:   "search-available-phone-numbers",
			Fields: fields_search_available_phone_numbers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchAvailablePhoneNumbersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_available_phone_numbers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchAvailablePhoneNumbers(ctx, input)
				}
				var results []*svc.SearchAvailablePhoneNumbersOutput
				p := svc.NewSearchAvailablePhoneNumbersPaginator(client, input)
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
		"update-account": {
			Name:   "update-account",
			Fields: fields_update_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccount(ctx, input)
			},
		},
		"update-account-settings": {
			Name:   "update-account-settings",
			Fields: fields_update_account_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountSettings(ctx, input)
			},
		},
		"update-bot": {
			Name:   "update-bot",
			Fields: fields_update_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBot(ctx, input)
			},
		},
		"update-global-settings": {
			Name:   "update-global-settings",
			Fields: fields_update_global_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlobalSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_global_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlobalSettings(ctx, input)
			},
		},
		"update-phone-number": {
			Name:   "update-phone-number",
			Fields: fields_update_phone_number,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePhoneNumberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_phone_number, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePhoneNumber(ctx, input)
			},
		},
		"update-phone-number-settings": {
			Name:   "update-phone-number-settings",
			Fields: fields_update_phone_number_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePhoneNumberSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_phone_number_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePhoneNumberSettings(ctx, input)
			},
		},
		"update-room": {
			Name:   "update-room",
			Fields: fields_update_room,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoomInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_room, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoom(ctx, input)
			},
		},
		"update-room-membership": {
			Name:   "update-room-membership",
			Fields: fields_update_room_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRoomMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_room_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRoomMembership(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
		"update-user-settings": {
			Name:   "update-user-settings",
			Fields: fields_update_user_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("chime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
