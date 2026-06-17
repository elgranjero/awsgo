package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// chimeCmd represents the chime command
var _chimeCmd = &cobra.Command{
	Use:   "chime",
	Short: "AWS chime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := chime.NewFromConfig(cfg)
		if _chimeAssociatePhoneNumberWithUser {
			chime_AssociatePhoneNumberWithUser(cfg, client)
			return
		}
		if _chimeAssociateSigninDelegateGroupsWithAccount {
			chime_AssociateSigninDelegateGroupsWithAccount(cfg, client)
			return
		}
		if _chimeBatchCreateRoomMembership {
			chime_BatchCreateRoomMembership(cfg, client)
			return
		}
		if _chimeBatchDeletePhoneNumber {
			chime_BatchDeletePhoneNumber(cfg, client)
			return
		}
		if _chimeBatchSuspendUser {
			chime_BatchSuspendUser(cfg, client)
			return
		}
		if _chimeBatchUnsuspendUser {
			chime_BatchUnsuspendUser(cfg, client)
			return
		}
		if _chimeBatchUpdatePhoneNumber {
			chime_BatchUpdatePhoneNumber(cfg, client)
			return
		}
		if _chimeBatchUpdateUser {
			chime_BatchUpdateUser(cfg, client)
			return
		}
		if _chimeCreateAccount {
			chime_CreateAccount(cfg, client)
			return
		}
		if _chimeCreateBot {
			chime_CreateBot(cfg, client)
			return
		}
		if _chimeCreateMeetingDialOut {
			chime_CreateMeetingDialOut(cfg, client)
			return
		}
		if _chimeCreatePhoneNumberOrder {
			chime_CreatePhoneNumberOrder(cfg, client)
			return
		}
		if _chimeCreateRoom {
			chime_CreateRoom(cfg, client)
			return
		}
		if _chimeCreateRoomMembership {
			chime_CreateRoomMembership(cfg, client)
			return
		}
		if _chimeCreateUser {
			chime_CreateUser(cfg, client)
			return
		}
		if _chimeDeleteAccount {
			chime_DeleteAccount(cfg, client)
			return
		}
		if _chimeDeleteEventsConfiguration {
			chime_DeleteEventsConfiguration(cfg, client)
			return
		}
		if _chimeDeletePhoneNumber {
			chime_DeletePhoneNumber(cfg, client)
			return
		}
		if _chimeDeleteRoom {
			chime_DeleteRoom(cfg, client)
			return
		}
		if _chimeDeleteRoomMembership {
			chime_DeleteRoomMembership(cfg, client)
			return
		}
		if _chimeDisassociatePhoneNumberFromUser {
			chime_DisassociatePhoneNumberFromUser(cfg, client)
			return
		}
		if _chimeDisassociateSigninDelegateGroupsFromAccount {
			chime_DisassociateSigninDelegateGroupsFromAccount(cfg, client)
			return
		}
		if _chimeGetAccount {
			chime_GetAccount(cfg, client)
			return
		}
		if _chimeGetAccountSettings {
			chime_GetAccountSettings(cfg, client)
			return
		}
		if _chimeGetBot {
			chime_GetBot(cfg, client)
			return
		}
		if _chimeGetEventsConfiguration {
			chime_GetEventsConfiguration(cfg, client)
			return
		}
		if _chimeGetGlobalSettings {
			chime_GetGlobalSettings(cfg, client)
			return
		}
		if _chimeGetPhoneNumber {
			chime_GetPhoneNumber(cfg, client)
			return
		}
		if _chimeGetPhoneNumberOrder {
			chime_GetPhoneNumberOrder(cfg, client)
			return
		}
		if _chimeGetPhoneNumberSettings {
			chime_GetPhoneNumberSettings(cfg, client)
			return
		}
		if _chimeGetRetentionSettings {
			chime_GetRetentionSettings(cfg, client)
			return
		}
		if _chimeGetRoom {
			chime_GetRoom(cfg, client)
			return
		}
		if _chimeGetUser {
			chime_GetUser(cfg, client)
			return
		}
		if _chimeGetUserSettings {
			chime_GetUserSettings(cfg, client)
			return
		}
		if _chimeInviteUsers {
			chime_InviteUsers(cfg, client)
			return
		}
		if _chimeListAccounts {
			chime_ListAccounts(cfg, client)
			return
		}
		if _chimeListBots {
			chime_ListBots(cfg, client)
			return
		}
		if _chimeListPhoneNumberOrders {
			chime_ListPhoneNumberOrders(cfg, client)
			return
		}
		if _chimeListPhoneNumbers {
			chime_ListPhoneNumbers(cfg, client)
			return
		}
		if _chimeListRoomMemberships {
			chime_ListRoomMemberships(cfg, client)
			return
		}
		if _chimeListRooms {
			chime_ListRooms(cfg, client)
			return
		}
		if _chimeListSupportedPhoneNumberCountries {
			chime_ListSupportedPhoneNumberCountries(cfg, client)
			return
		}
		if _chimeListUsers {
			chime_ListUsers(cfg, client)
			return
		}
		if _chimeLogoutUser {
			chime_LogoutUser(cfg, client)
			return
		}
		if _chimePutEventsConfiguration {
			chime_PutEventsConfiguration(cfg, client)
			return
		}
		if _chimePutRetentionSettings {
			chime_PutRetentionSettings(cfg, client)
			return
		}
		if _chimeRedactConversationMessage {
			chime_RedactConversationMessage(cfg, client)
			return
		}
		if _chimeRedactRoomMessage {
			chime_RedactRoomMessage(cfg, client)
			return
		}
		if _chimeRegenerateSecurityToken {
			chime_RegenerateSecurityToken(cfg, client)
			return
		}
		if _chimeResetPersonalPIN {
			chime_ResetPersonalPIN(cfg, client)
			return
		}
		if _chimeRestorePhoneNumber {
			chime_RestorePhoneNumber(cfg, client)
			return
		}
		if _chimeSearchAvailablePhoneNumbers {
			chime_SearchAvailablePhoneNumbers(cfg, client)
			return
		}
		if _chimeUpdateAccount {
			chime_UpdateAccount(cfg, client)
			return
		}
		if _chimeUpdateAccountSettings {
			chime_UpdateAccountSettings(cfg, client)
			return
		}
		if _chimeUpdateBot {
			chime_UpdateBot(cfg, client)
			return
		}
		if _chimeUpdateGlobalSettings {
			chime_UpdateGlobalSettings(cfg, client)
			return
		}
		if _chimeUpdatePhoneNumber {
			chime_UpdatePhoneNumber(cfg, client)
			return
		}
		if _chimeUpdatePhoneNumberSettings {
			chime_UpdatePhoneNumberSettings(cfg, client)
			return
		}
		if _chimeUpdateRoom {
			chime_UpdateRoom(cfg, client)
			return
		}
		if _chimeUpdateRoomMembership {
			chime_UpdateRoomMembership(cfg, client)
			return
		}
		if _chimeUpdateUser {
			chime_UpdateUser(cfg, client)
			return
		}
		if _chimeUpdateUserSettings {
			chime_UpdateUserSettings(cfg, client)
			return
		}

	},
}

var (
	_chimeAssociatePhoneNumberWithUser                bool
	_chimeAssociateSigninDelegateGroupsWithAccount    bool
	_chimeBatchCreateRoomMembership                   bool
	_chimeBatchDeletePhoneNumber                      bool
	_chimeBatchSuspendUser                            bool
	_chimeBatchUnsuspendUser                          bool
	_chimeBatchUpdatePhoneNumber                      bool
	_chimeBatchUpdateUser                             bool
	_chimeCreateAccount                               bool
	_chimeCreateBot                                   bool
	_chimeCreateMeetingDialOut                        bool
	_chimeCreatePhoneNumberOrder                      bool
	_chimeCreateRoom                                  bool
	_chimeCreateRoomMembership                        bool
	_chimeCreateUser                                  bool
	_chimeDeleteAccount                               bool
	_chimeDeleteEventsConfiguration                   bool
	_chimeDeletePhoneNumber                           bool
	_chimeDeleteRoom                                  bool
	_chimeDeleteRoomMembership                        bool
	_chimeDisassociatePhoneNumberFromUser             bool
	_chimeDisassociateSigninDelegateGroupsFromAccount bool
	_chimeGetAccount                                  bool
	_chimeGetAccountSettings                          bool
	_chimeGetBot                                      bool
	_chimeGetEventsConfiguration                      bool
	_chimeGetGlobalSettings                           bool
	_chimeGetPhoneNumber                              bool
	_chimeGetPhoneNumberOrder                         bool
	_chimeGetPhoneNumberSettings                      bool
	_chimeGetRetentionSettings                        bool
	_chimeGetRoom                                     bool
	_chimeGetUser                                     bool
	_chimeGetUserSettings                             bool
	_chimeInviteUsers                                 bool
	_chimeListAccounts                                bool
	_chimeListBots                                    bool
	_chimeListPhoneNumberOrders                       bool
	_chimeListPhoneNumbers                            bool
	_chimeListRoomMemberships                         bool
	_chimeListRooms                                   bool
	_chimeListSupportedPhoneNumberCountries           bool
	_chimeListUsers                                   bool
	_chimeLogoutUser                                  bool
	_chimePutEventsConfiguration                      bool
	_chimePutRetentionSettings                        bool
	_chimeRedactConversationMessage                   bool
	_chimeRedactRoomMessage                           bool
	_chimeRegenerateSecurityToken                     bool
	_chimeResetPersonalPIN                            bool
	_chimeRestorePhoneNumber                          bool
	_chimeSearchAvailablePhoneNumbers                 bool
	_chimeUpdateAccount                               bool
	_chimeUpdateAccountSettings                       bool
	_chimeUpdateBot                                   bool
	_chimeUpdateGlobalSettings                        bool
	_chimeUpdatePhoneNumber                           bool
	_chimeUpdatePhoneNumberSettings                   bool
	_chimeUpdateRoom                                  bool
	_chimeUpdateRoomMembership                        bool
	_chimeUpdateUser                                  bool
	_chimeUpdateUserSettings                          bool

	_chimeAccountId                     string
	_chimeAccountSettings               string
	_chimeAlexaForBusinessMetadata      string
	_chimeAreaCode                      string
	_chimeBotId                         string
	_chimeBusinessCalling               string
	_chimeCallingName                   string
	_chimeCity                          string
	_chimeClientRequestToken            string
	_chimeConversationId                string
	_chimeCountry                       string
	_chimeDefaultLicense                string
	_chimeDisabled                      string
	_chimeDisplayName                   string
	_chimeDomain                        string
	_chimeE164PhoneNumber               string
	_chimeE164PhoneNumbers              []string
	_chimeEmail                         string
	_chimeFilterName                    string
	_chimeFilterValue                   string
	_chimeFromPhoneNumber               string
	_chimeGroupNames                    []string
	_chimeJoinToken                     string
	_chimeLambdaFunctionArn             string
	_chimeLicenseType                   string
	_chimeMaxResults                    string
	_chimeMeetingId                     string
	_chimeMemberId                      string
	_chimeMembershipItemList            string
	_chimeMessageId                     string
	_chimeName                          string
	_chimeNextToken                     string
	_chimeOutboundEventsHTTPSEndpoint   string
	_chimePhoneNumberId                 string
	_chimePhoneNumberIds                []string
	_chimePhoneNumberOrderId            string
	_chimePhoneNumberType               string
	_chimeProductType                   string
	_chimeRetentionSettings             string
	_chimeRole                          string
	_chimeRoomId                        string
	_chimeSigninDelegateGroups          string
	_chimeState                         string
	_chimeStatus                        string
	_chimeToPhoneNumber                 string
	_chimeTollFreePrefix                string
	_chimeUpdatePhoneNumberRequestItems string
	_chimeUpdateUserRequestItems        string
	_chimeUserEmail                     string
	_chimeUserEmailList                 []string
	_chimeUserId                        string
	_chimeUserIdList                    []string
	_chimeUserSettings                  string
	_chimeUserType                      string
	_chimeUsername                      string
	_chimeVoiceConnector                string
)

// Associates a phone number with the specified Amazon Chime user.
func chime_AssociatePhoneNumberWithUser(cfg aws.Config, client *chime.Client) {
	input := &chime.AssociatePhoneNumberWithUserInput{
		// AccountId: *string, // Required
		// E164PhoneNumber: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeE164PhoneNumber) > 0 {
		input.E164PhoneNumber = aws.String(_chimeE164PhoneNumber)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}

	if resp, err := client.AssociatePhoneNumberWithUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified sign-in delegate groups with the specified Amazon
// Chime account.
func chime_AssociateSigninDelegateGroupsWithAccount(cfg aws.Config, client *chime.Client) {
	input := &chime.AssociateSigninDelegateGroupsWithAccountInput{
		// AccountId: *string, // Required
		// SigninDelegateGroups: []types.SigninDelegateGroup, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeSigninDelegateGroups) > 0 {
		if err := assignInputField(input, "SigninDelegateGroups", _chimeSigninDelegateGroups); err != nil {
			log.Errorf("invalid --signin-delegate-groups: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateSigninDelegateGroupsWithAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds up to 50 members to a chat room in an Amazon Chime Enterprise account.
// Members can be users or bots. The member role designates whether the member is a
// chat room administrator or a general chat room member.
func chime_BatchCreateRoomMembership(cfg aws.Config, client *chime.Client) {
	input := &chime.BatchCreateRoomMembershipInput{
		// AccountId: *string, // Required
		// MembershipItemList: []types.MembershipItem, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMembershipItemList) > 0 {
		if err := assignInputField(input, "MembershipItemList", _chimeMembershipItemList); err != nil {
			log.Errorf("invalid --membership-item-list: %s", err.Error())
			return
		}
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}

	if resp, err := client.BatchCreateRoomMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves phone numbers into the Deletion queue. Phone numbers must be
// disassociated from any users or Amazon Chime Voice Connectors before they can be
// deleted.
//
// Phone numbers remain in the Deletion queue for 7 days before they are deleted
// permanently.
func chime_BatchDeletePhoneNumber(cfg aws.Config, client *chime.Client) {
	input := &chime.BatchDeletePhoneNumberInput{
		// PhoneNumberIds: []string, // Required
	}

	if len(_chimePhoneNumberIds) > 0 {
		input.PhoneNumberIds = append([]string(nil), _chimePhoneNumberIds...)
	}

	if resp, err := client.BatchDeletePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Suspends up to 50 users from a Team or EnterpriseLWA Amazon Chime account. For
// more information about different account types, see [Managing Your Amazon Chime Accounts]in the Amazon Chime
// Administration Guide.
//
// Users suspended from a Team account are disassociated from the account,but they
// can continue to use Amazon Chime as free users. To remove the suspension from
// suspended Team account users, invite them to the Team account again. You can
// use the InviteUsersaction to do so.
//
// Users suspended from an EnterpriseLWA account are immediately signed out of
// Amazon Chime and can no longer sign in. To remove the suspension from suspended
// EnterpriseLWA account users, use the BatchUnsuspendUser action.
//
// To sign out users without suspending them, use the LogoutUser action.
//
// [Managing Your Amazon Chime Accounts]: https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html
func chime_BatchSuspendUser(cfg aws.Config, client *chime.Client) {
	input := &chime.BatchSuspendUserInput{
		// AccountId: *string, // Required
		// UserIdList: []string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserIdList) > 0 {
		input.UserIdList = append([]string(nil), _chimeUserIdList...)
	}

	if resp, err := client.BatchSuspendUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the suspension from up to 50 previously suspended users for the
// specified Amazon Chime EnterpriseLWA account. Only users on EnterpriseLWA
// accounts can be unsuspended using this action. For more information about
// different account types, see [Managing Your Amazon Chime Accounts]in the account types, in the Amazon Chime
// Administration Guide.
//
// Previously suspended users who are unsuspended using this action are returned
// to Registered status. Users who are not previously suspended are ignored.
//
// [Managing Your Amazon Chime Accounts]: https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html
func chime_BatchUnsuspendUser(cfg aws.Config, client *chime.Client) {
	input := &chime.BatchUnsuspendUserInput{
		// AccountId: *string, // Required
		// UserIdList: []string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserIdList) > 0 {
		input.UserIdList = append([]string(nil), _chimeUserIdList...)
	}

	if resp, err := client.BatchUnsuspendUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates phone number product types or calling names. You can update one
// attribute at a time for each UpdatePhoneNumberRequestItem . For example, you can
// update the product type or the calling name.
//
// For toll-free numbers, you cannot use the Amazon Chime Business Calling product
// type. For numbers outside the U.S., you must use the Amazon Chime SIP Media
// Application Dial-In product type.
//
// Updates to outbound calling names can take up to 72 hours to complete. Pending
// updates to outbound calling names must be complete before you can request
// another update.
func chime_BatchUpdatePhoneNumber(cfg aws.Config, client *chime.Client) {
	input := &chime.BatchUpdatePhoneNumberInput{
		// UpdatePhoneNumberRequestItems: []types.UpdatePhoneNumberRequestItem, // Required
	}

	if len(_chimeUpdatePhoneNumberRequestItems) > 0 {
		if err := assignInputField(input, "UpdatePhoneNumberRequestItems", _chimeUpdatePhoneNumberRequestItems); err != nil {
			log.Errorf("invalid --update-phone-number-request-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdatePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates user details within the UpdateUserRequestItem object for up to 20 users for the specified
// Amazon Chime account. Currently, only LicenseType updates are supported for
// this action.
func chime_BatchUpdateUser(cfg aws.Config, client *chime.Client) {
	input := &chime.BatchUpdateUserInput{
		// AccountId: *string, // Required
		// UpdateUserRequestItems: []types.UpdateUserRequestItem, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUpdateUserRequestItems) > 0 {
		if err := assignInputField(input, "UpdateUserRequestItems", _chimeUpdateUserRequestItems); err != nil {
			log.Errorf("invalid --update-user-request-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Chime account under the administrator's AWS account. Only Team
// account types are currently supported for this action. For more information
// about different account types, see [Managing Your Amazon Chime Accounts]in the Amazon Chime Administration Guide.
//
// [Managing Your Amazon Chime Accounts]: https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html
func chime_CreateAccount(cfg aws.Config, client *chime.Client) {
	input := &chime.CreateAccountInput{
		// Name: *string, // Required
	}

	if len(_chimeName) > 0 {
		input.Name = aws.String(_chimeName)
	}

	if resp, err := client.CreateAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a bot for an Amazon Chime Enterprise account.
func chime_CreateBot(cfg aws.Config, client *chime.Client) {
	input := &chime.CreateBotInput{
		// AccountId: *string, // Required
		// DisplayName: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeDisplayName) > 0 {
		input.DisplayName = aws.String(_chimeDisplayName)
	}
	if len(_chimeDomain) > 0 {
		input.Domain = aws.String(_chimeDomain)
	}

	if resp, err := client.CreateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uses the join token and call metadata in a meeting request (From number, To
// number, and so forth) to initiate an outbound call to a public switched
// telephone network (PSTN) and join them into a Chime meeting. Also ensures that
// the From number belongs to the customer.
//
// To play welcome audio or implement an interactive voice response (IVR), use the
// CreateSipMediaApplicationCall action with the corresponding SIP media
// application ID.
//
// This API is not available in a dedicated namespace.
func chime_CreateMeetingDialOut(cfg aws.Config, client *chime.Client) {
	input := &chime.CreateMeetingDialOutInput{
		// FromPhoneNumber: *string, // Required
		// JoinToken: *string, // Required
		// MeetingId: *string, // Required
		// ToPhoneNumber: *string, // Required
	}

	if len(_chimeFromPhoneNumber) > 0 {
		input.FromPhoneNumber = aws.String(_chimeFromPhoneNumber)
	}
	if len(_chimeJoinToken) > 0 {
		input.JoinToken = aws.String(_chimeJoinToken)
	}
	if len(_chimeMeetingId) > 0 {
		input.MeetingId = aws.String(_chimeMeetingId)
	}
	if len(_chimeToPhoneNumber) > 0 {
		input.ToPhoneNumber = aws.String(_chimeToPhoneNumber)
	}

	if resp, err := client.CreateMeetingDialOut(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an order for phone numbers to be provisioned. For toll-free numbers,
// you cannot use the Amazon Chime Business Calling product type. For numbers
// outside the U.S., you must use the Amazon Chime SIP Media Application Dial-In
// product type.
func chime_CreatePhoneNumberOrder(cfg aws.Config, client *chime.Client) {
	input := &chime.CreatePhoneNumberOrderInput{
		// E164PhoneNumbers: []string, // Required
		// ProductType: types.PhoneNumberProductType, // Required
	}

	if len(_chimeE164PhoneNumbers) > 0 {
		input.E164PhoneNumbers = append([]string(nil), _chimeE164PhoneNumbers...)
	}
	if len(_chimeProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimeProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePhoneNumberOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a chat room for the specified Amazon Chime Enterprise account.
func chime_CreateRoom(cfg aws.Config, client *chime.Client) {
	input := &chime.CreateRoomInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeName) > 0 {
		input.Name = aws.String(_chimeName)
	}
	if len(_chimeClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_chimeClientRequestToken)
	}

	if resp, err := client.CreateRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a member to a chat room in an Amazon Chime Enterprise account. A member
// can be either a user or a bot. The member role designates whether the member is
// a chat room administrator or a general chat room member.
func chime_CreateRoomMembership(cfg aws.Config, client *chime.Client) {
	input := &chime.CreateRoomMembershipInput{
		// AccountId: *string, // Required
		// MemberId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMemberId) > 0 {
		input.MemberId = aws.String(_chimeMemberId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}
	if len(_chimeRole) > 0 {
		if err := assignInputField(input, "Role", _chimeRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRoomMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user under the specified Amazon Chime account.
func chime_CreateUser(cfg aws.Config, client *chime.Client) {
	input := &chime.CreateUserInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeEmail) > 0 {
		input.Email = aws.String(_chimeEmail)
	}
	if len(_chimeUserType) > 0 {
		if err := assignInputField(input, "UserType", _chimeUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}
	if len(_chimeUsername) > 0 {
		input.Username = aws.String(_chimeUsername)
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified Amazon Chime account. You must suspend all users before
// deleting Team account. You can use the BatchSuspendUser action to dodo.
//
// For EnterpriseLWA and EnterpriseAD accounts, you must release the claimed
// domains for your Amazon Chime account before deletion. As soon as you release
// the domain, all users under that account are suspended.
//
// Deleted accounts appear in your Disabled accounts list for 90 days. To restore
// deleted account from your Disabled accounts list, you must contact AWS Support.
//
// After 90 days, deleted accounts are permanently removed from your Disabled
// accounts list.
func chime_DeleteAccount(cfg aws.Config, client *chime.Client) {
	input := &chime.DeleteAccountInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}

	if resp, err := client.DeleteAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the events configuration that allows a bot to receive outgoing events.
func chime_DeleteEventsConfiguration(cfg aws.Config, client *chime.Client) {
	input := &chime.DeleteEventsConfigurationInput{
		// AccountId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeBotId) > 0 {
		input.BotId = aws.String(_chimeBotId)
	}

	if resp, err := client.DeleteEventsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves the specified phone number into the Deletion queue. A phone number must
// be disassociated from any users or Amazon Chime Voice Connectors before it can
// be deleted.
//
// Deleted phone numbers remain in the Deletion queue for 7 days before they are
// deleted permanently.
func chime_DeletePhoneNumber(cfg aws.Config, client *chime.Client) {
	input := &chime.DeletePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimePhoneNumberId)
	}

	if resp, err := client.DeletePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a chat room in an Amazon Chime Enterprise account.
func chime_DeleteRoom(cfg aws.Config, client *chime.Client) {
	input := &chime.DeleteRoomInput{
		// AccountId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}

	if resp, err := client.DeleteRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member from a chat room in an Amazon Chime Enterprise account.
func chime_DeleteRoomMembership(cfg aws.Config, client *chime.Client) {
	input := &chime.DeleteRoomMembershipInput{
		// AccountId: *string, // Required
		// MemberId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMemberId) > 0 {
		input.MemberId = aws.String(_chimeMemberId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}

	if resp, err := client.DeleteRoomMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the primary provisioned phone number from the specified Amazon
// Chime user.
func chime_DisassociatePhoneNumberFromUser(cfg aws.Config, client *chime.Client) {
	input := &chime.DisassociatePhoneNumberFromUserInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}

	if resp, err := client.DisassociatePhoneNumberFromUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified sign-in delegate groups from the specified Amazon
// Chime account.
func chime_DisassociateSigninDelegateGroupsFromAccount(cfg aws.Config, client *chime.Client) {
	input := &chime.DisassociateSigninDelegateGroupsFromAccountInput{
		// AccountId: *string, // Required
		// GroupNames: []string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeGroupNames) > 0 {
		input.GroupNames = append([]string(nil), _chimeGroupNames...)
	}

	if resp, err := client.DisassociateSigninDelegateGroupsFromAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified Amazon Chime account, such as account type
// and supported licenses.
func chime_GetAccount(cfg aws.Config, client *chime.Client) {
	input := &chime.GetAccountInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}

	if resp, err := client.GetAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves account settings for the specified Amazon Chime account ID, such as
// remote control and dialout settings. For more information about these settings,
// see [Use the Policies Page]in the Amazon Chime Administration Guide.
//
// [Use the Policies Page]: https://docs.aws.amazon.com/chime/latest/ag/policies.html
func chime_GetAccountSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.GetAccountSettingsInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified bot, such as bot email address, bot type,
// status, and display name.
func chime_GetBot(cfg aws.Config, client *chime.Client) {
	input := &chime.GetBotInput{
		// AccountId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeBotId) > 0 {
		input.BotId = aws.String(_chimeBotId)
	}

	if resp, err := client.GetBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details for an events configuration that allows a bot to receive outgoing
// events, such as an HTTPS endpoint or Lambda function ARN.
func chime_GetEventsConfiguration(cfg aws.Config, client *chime.Client) {
	input := &chime.GetEventsConfigurationInput{
		// AccountId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeBotId) > 0 {
		input.BotId = aws.String(_chimeBotId)
	}

	if resp, err := client.GetEventsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves global settings for the administrator's AWS account, such as Amazon
// Chime Business Calling and Amazon Chime Voice Connector settings.
func chime_GetGlobalSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.GetGlobalSettingsInput{}

	if resp, err := client.GetGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified phone number ID, such as associations,
// capabilities, and product type.
func chime_GetPhoneNumber(cfg aws.Config, client *chime.Client) {
	input := &chime.GetPhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimePhoneNumberId)
	}

	if resp, err := client.GetPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified phone number order, such as the order
// creation timestamp, phone numbers in E.164 format, product type, and order
// status.
func chime_GetPhoneNumberOrder(cfg aws.Config, client *chime.Client) {
	input := &chime.GetPhoneNumberOrderInput{
		// PhoneNumberOrderId: *string, // Required
	}

	if len(_chimePhoneNumberOrderId) > 0 {
		input.PhoneNumberOrderId = aws.String(_chimePhoneNumberOrderId)
	}

	if resp, err := client.GetPhoneNumberOrder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the phone number settings for the administrator's AWS account, such
// as the default outbound calling name.
func chime_GetPhoneNumberSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.GetPhoneNumberSettingsInput{}

	if resp, err := client.GetPhoneNumberSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the retention settings for the specified Amazon Chime Enterprise account.
// For more information about retention settings, see [Managing Chat Retention Policies]in the Amazon Chime
// Administration Guide.
//
// [Managing Chat Retention Policies]: https://docs.aws.amazon.com/chime/latest/ag/chat-retention.html
func chime_GetRetentionSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.GetRetentionSettingsInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}

	if resp, err := client.GetRetentionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves room details, such as the room name, for a room in an Amazon Chime
// Enterprise account.
func chime_GetRoom(cfg aws.Config, client *chime.Client) {
	input := &chime.GetRoomInput{
		// AccountId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}

	if resp, err := client.GetRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details for the specified user ID, such as primary email address,
// license type,and personal meeting PIN.
//
// To retrieve user details with an email address instead of a user ID, use the ListUsers
// action, and then filter by email address.
func chime_GetUser(cfg aws.Config, client *chime.Client) {
	input := &chime.GetUserInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}

	if resp, err := client.GetUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves settings for the specified user ID, such as any associated phone
// number settings.
func chime_GetUserSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.GetUserSettingsInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}

	if resp, err := client.GetUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends email to a maximum of 50 users, inviting them to the specified Amazon
// Chime Team account. Only Team account types are currently supported for this
// action.
func chime_InviteUsers(cfg aws.Config, client *chime.Client) {
	input := &chime.InviteUsersInput{
		// AccountId: *string, // Required
		// UserEmailList: []string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserEmailList) > 0 {
		input.UserEmailList = append([]string(nil), _chimeUserEmailList...)
	}
	if len(_chimeUserType) > 0 {
		if err := assignInputField(input, "UserType", _chimeUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.InviteUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Chime accounts under the administrator's AWS account. You can
// filter accounts by account name prefix. To find out which Amazon Chime account a
// user belongs to, you can filter by the user's email address, which returns one
// account result.
func chime_ListAccounts(cfg aws.Config, client *chime.Client) {
	input := &chime.ListAccountsInput{}

	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeName) > 0 {
		input.Name = aws.String(_chimeName)
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}
	if len(_chimeUserEmail) > 0 {
		input.UserEmail = aws.String(_chimeUserEmail)
	}

	if disablePaginator() {
		if resp, err := client.ListAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListAccountsOutput
	p := chime.NewListAccountsPaginator(client, input)
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

// Lists the bots associated with the administrator's Amazon Chime Enterprise
// account ID.
func chime_ListBots(cfg aws.Config, client *chime.Client) {
	input := &chime.ListBotsInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListBotsOutput
	p := chime.NewListBotsPaginator(client, input)
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

// Lists the phone number orders for the administrator's Amazon Chime account.
func chime_ListPhoneNumberOrders(cfg aws.Config, client *chime.Client) {
	input := &chime.ListPhoneNumberOrdersInput{}

	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumberOrders(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListPhoneNumberOrdersOutput
	p := chime.NewListPhoneNumberOrdersPaginator(client, input)
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

// Lists the phone numbers for the specified Amazon Chime account, Amazon Chime
// user, Amazon Chime Voice Connector, or Amazon Chime Voice Connector group.
func chime_ListPhoneNumbers(cfg aws.Config, client *chime.Client) {
	input := &chime.ListPhoneNumbersInput{}

	if len(_chimeFilterName) > 0 {
		if err := assignInputField(input, "FilterName", _chimeFilterName); err != nil {
			log.Errorf("invalid --filter-name: %s", err.Error())
			return
		}
	}
	if len(_chimeFilterValue) > 0 {
		input.FilterValue = aws.String(_chimeFilterValue)
	}
	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}
	if len(_chimeProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimeProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}
	if len(_chimeStatus) > 0 {
		if err := assignInputField(input, "Status", _chimeStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListPhoneNumbersOutput
	p := chime.NewListPhoneNumbersPaginator(client, input)
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

// Lists the membership details for the specified room in an Amazon Chime
// Enterprise account, such as the members' IDs, email addresses, and names.
func chime_ListRoomMemberships(cfg aws.Config, client *chime.Client) {
	input := &chime.ListRoomMembershipsInput{
		// AccountId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}
	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRoomMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListRoomMembershipsOutput
	p := chime.NewListRoomMembershipsPaginator(client, input)
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

// Lists the room details for the specified Amazon Chime Enterprise account.
// Optionally, filter the results by a member ID (user ID or bot ID) to see a list
// of rooms that the member belongs to.
func chime_ListRooms(cfg aws.Config, client *chime.Client) {
	input := &chime.ListRoomsInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeMemberId) > 0 {
		input.MemberId = aws.String(_chimeMemberId)
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRooms(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListRoomsOutput
	p := chime.NewListRoomsPaginator(client, input)
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

// Lists supported phone number countries.
func chime_ListSupportedPhoneNumberCountries(cfg aws.Config, client *chime.Client) {
	input := &chime.ListSupportedPhoneNumberCountriesInput{
		// ProductType: types.PhoneNumberProductType, // Required
	}

	if len(_chimeProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimeProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListSupportedPhoneNumberCountries(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the users that belong to the specified Amazon Chime account. You can
// specify an email address to list only the user that the email address belongs
// to.
func chime_ListUsers(cfg aws.Config, client *chime.Client) {
	input := &chime.ListUsersInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}
	if len(_chimeUserEmail) > 0 {
		input.UserEmail = aws.String(_chimeUserEmail)
	}
	if len(_chimeUserType) > 0 {
		if err := assignInputField(input, "UserType", _chimeUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.ListUsersOutput
	p := chime.NewListUsersPaginator(client, input)
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

// Logs out the specified user from all of the devices they are currently logged
// into.
func chime_LogoutUser(cfg aws.Config, client *chime.Client) {
	input := &chime.LogoutUserInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}

	if resp, err := client.LogoutUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an events configuration that allows a bot to receive outgoing events
// sent by Amazon Chime. Choose either an HTTPS endpoint or a Lambda function ARN.
// For more information, see Bot.
func chime_PutEventsConfiguration(cfg aws.Config, client *chime.Client) {
	input := &chime.PutEventsConfigurationInput{
		// AccountId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeBotId) > 0 {
		input.BotId = aws.String(_chimeBotId)
	}
	if len(_chimeLambdaFunctionArn) > 0 {
		input.LambdaFunctionArn = aws.String(_chimeLambdaFunctionArn)
	}
	if len(_chimeOutboundEventsHTTPSEndpoint) > 0 {
		input.OutboundEventsHTTPSEndpoint = aws.String(_chimeOutboundEventsHTTPSEndpoint)
	}

	if resp, err := client.PutEventsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts retention settings for the specified Amazon Chime Enterprise account. We
// recommend using AWS CloudTrail to monitor usage of this API for your account.
// For more information, see [Logging Amazon Chime API Calls with AWS CloudTrail]in the Amazon Chime Administration Guide.
//
// To turn off existing retention settings, remove the number of days from the
// corresponding RetentionDays field in the RetentionSettings object. For more
// information about retention settings, see [Managing Chat Retention Policies]in the Amazon Chime Administration
// Guide.
//
// [Managing Chat Retention Policies]: https://docs.aws.amazon.com/chime/latest/ag/chat-retention.html
// [Logging Amazon Chime API Calls with AWS CloudTrail]: https://docs.aws.amazon.com/chime/latest/ag/cloudtrail.html
func chime_PutRetentionSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.PutRetentionSettingsInput{
		// AccountId: *string, // Required
		// RetentionSettings: *types.RetentionSettings, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeRetentionSettings) > 0 {
		if err := assignInputField(input, "RetentionSettings", _chimeRetentionSettings); err != nil {
			log.Errorf("invalid --retention-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRetentionSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Redacts the specified message from the specified Amazon Chime conversation.
func chime_RedactConversationMessage(cfg aws.Config, client *chime.Client) {
	input := &chime.RedactConversationMessageInput{
		// AccountId: *string, // Required
		// ConversationId: *string, // Required
		// MessageId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeConversationId) > 0 {
		input.ConversationId = aws.String(_chimeConversationId)
	}
	if len(_chimeMessageId) > 0 {
		input.MessageId = aws.String(_chimeMessageId)
	}

	if resp, err := client.RedactConversationMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Redacts the specified message from the specified Amazon Chime channel.
func chime_RedactRoomMessage(cfg aws.Config, client *chime.Client) {
	input := &chime.RedactRoomMessageInput{
		// AccountId: *string, // Required
		// MessageId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMessageId) > 0 {
		input.MessageId = aws.String(_chimeMessageId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}

	if resp, err := client.RedactRoomMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Regenerates the security token for a bot.
func chime_RegenerateSecurityToken(cfg aws.Config, client *chime.Client) {
	input := &chime.RegenerateSecurityTokenInput{
		// AccountId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeBotId) > 0 {
		input.BotId = aws.String(_chimeBotId)
	}

	if resp, err := client.RegenerateSecurityToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets the personal meeting PIN for the specified user on an Amazon Chime
// account. Returns the Userobject with the updated personal meeting PIN.
func chime_ResetPersonalPIN(cfg aws.Config, client *chime.Client) {
	input := &chime.ResetPersonalPINInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}

	if resp, err := client.ResetPersonalPIN(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves a phone number from the Deletion queue back into the phone number
// Inventory.
func chime_RestorePhoneNumber(cfg aws.Config, client *chime.Client) {
	input := &chime.RestorePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimePhoneNumberId)
	}

	if resp, err := client.RestorePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches for phone numbers that can be ordered. For US numbers, provide at
// least one of the following search filters: AreaCode , City , State , or
// TollFreePrefix . If you provide City , you must also provide State . Numbers
// outside the US only support the PhoneNumberType filter, which you must use.
func chime_SearchAvailablePhoneNumbers(cfg aws.Config, client *chime.Client) {
	input := &chime.SearchAvailablePhoneNumbersInput{}

	if len(_chimeAreaCode) > 0 {
		input.AreaCode = aws.String(_chimeAreaCode)
	}
	if len(_chimeCity) > 0 {
		input.City = aws.String(_chimeCity)
	}
	if len(_chimeCountry) > 0 {
		input.Country = aws.String(_chimeCountry)
	}
	if len(_chimeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _chimeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_chimeNextToken) > 0 {
		input.NextToken = aws.String(_chimeNextToken)
	}
	if len(_chimePhoneNumberType) > 0 {
		if err := assignInputField(input, "PhoneNumberType", _chimePhoneNumberType); err != nil {
			log.Errorf("invalid --phone-number-type: %s", err.Error())
			return
		}
	}
	if len(_chimeState) > 0 {
		input.State = aws.String(_chimeState)
	}
	if len(_chimeTollFreePrefix) > 0 {
		input.TollFreePrefix = aws.String(_chimeTollFreePrefix)
	}

	if disablePaginator() {
		if resp, err := client.SearchAvailablePhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*chime.SearchAvailablePhoneNumbersOutput
	p := chime.NewSearchAvailablePhoneNumbersPaginator(client, input)
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

// Updates account details for the specified Amazon Chime account. Currently, only
// account name and default license updates are supported for this action.
func chime_UpdateAccount(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateAccountInput{
		// AccountId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeDefaultLicense) > 0 {
		if err := assignInputField(input, "DefaultLicense", _chimeDefaultLicense); err != nil {
			log.Errorf("invalid --default-license: %s", err.Error())
			return
		}
	}
	if len(_chimeName) > 0 {
		input.Name = aws.String(_chimeName)
	}

	if resp, err := client.UpdateAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for the specified Amazon Chime account. You can update
// settings for remote control of shared screens, or for the dial-out option. For
// more information about these settings, see [Use the Policies Page]in the Amazon Chime Administration
// Guide.
//
// [Use the Policies Page]: https://docs.aws.amazon.com/chime/latest/ag/policies.html
func chime_UpdateAccountSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateAccountSettingsInput{
		// AccountId: *string, // Required
		// AccountSettings: *types.AccountSettings, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeAccountSettings) > 0 {
		if err := assignInputField(input, "AccountSettings", _chimeAccountSettings); err != nil {
			log.Errorf("invalid --account-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of the specified bot, such as starting or stopping the bot
// from running in your Amazon Chime Enterprise account.
func chime_UpdateBot(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateBotInput{
		// AccountId: *string, // Required
		// BotId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeBotId) > 0 {
		input.BotId = aws.String(_chimeBotId)
	}
	if len(_chimeDisabled) > 0 {
		if err := assignInputField(input, "Disabled", _chimeDisabled); err != nil {
			log.Errorf("invalid --disabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates global settings for the administrator's AWS account, such as Amazon
// Chime Business Calling and Amazon Chime Voice Connector settings.
func chime_UpdateGlobalSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateGlobalSettingsInput{}

	if len(_chimeBusinessCalling) > 0 {
		if err := assignInputField(input, "BusinessCalling", _chimeBusinessCalling); err != nil {
			log.Errorf("invalid --business-calling: %s", err.Error())
			return
		}
	}
	if len(_chimeVoiceConnector) > 0 {
		if err := assignInputField(input, "VoiceConnector", _chimeVoiceConnector); err != nil {
			log.Errorf("invalid --voice-connector: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGlobalSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates phone number details, such as product type or calling name, for the
// specified phone number ID. You can update one phone number detail at a time. For
// example, you can update either the product type or the calling name in one
// action.
//
// For toll-free numbers, you cannot use the Amazon Chime Business Calling product
// type. For numbers outside the U.S., you must use the Amazon Chime SIP Media
// Application Dial-In product type.
//
// Updates to outbound calling names can take 72 hours to complete. Pending
// updates to outbound calling names must be complete before you can request
// another update.
func chime_UpdatePhoneNumber(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdatePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_chimePhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_chimePhoneNumberId)
	}
	if len(_chimeCallingName) > 0 {
		input.CallingName = aws.String(_chimeCallingName)
	}
	if len(_chimeProductType) > 0 {
		if err := assignInputField(input, "ProductType", _chimeProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the phone number settings for the administrator's AWS account, such as
// the default outbound calling name. You can update the default outbound calling
// name once every seven days. Outbound calling names can take up to 72 hours to
// update.
func chime_UpdatePhoneNumberSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdatePhoneNumberSettingsInput{
		// CallingName: *string, // Required
	}

	if len(_chimeCallingName) > 0 {
		input.CallingName = aws.String(_chimeCallingName)
	}

	if resp, err := client.UpdatePhoneNumberSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates room details, such as the room name, for a room in an Amazon Chime
// Enterprise account.
func chime_UpdateRoom(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateRoomInput{
		// AccountId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}
	if len(_chimeName) > 0 {
		input.Name = aws.String(_chimeName)
	}

	if resp, err := client.UpdateRoom(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates room membership details, such as the member role, for a room in an
// Amazon Chime Enterprise account. The member role designates whether the member
// is a chat room administrator or a general chat room member. The member role can
// be updated only for user IDs.
func chime_UpdateRoomMembership(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateRoomMembershipInput{
		// AccountId: *string, // Required
		// MemberId: *string, // Required
		// RoomId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeMemberId) > 0 {
		input.MemberId = aws.String(_chimeMemberId)
	}
	if len(_chimeRoomId) > 0 {
		input.RoomId = aws.String(_chimeRoomId)
	}
	if len(_chimeRole) > 0 {
		if err := assignInputField(input, "Role", _chimeRole); err != nil {
			log.Errorf("invalid --role: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRoomMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates user details for a specified user ID. Currently, only LicenseType
// updates are supported for this action.
func chime_UpdateUser(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateUserInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}
	if len(_chimeAlexaForBusinessMetadata) > 0 {
		if err := assignInputField(input, "AlexaForBusinessMetadata", _chimeAlexaForBusinessMetadata); err != nil {
			log.Errorf("invalid --alexa-for-business-metadata: %s", err.Error())
			return
		}
	}
	if len(_chimeLicenseType) > 0 {
		if err := assignInputField(input, "LicenseType", _chimeLicenseType); err != nil {
			log.Errorf("invalid --license-type: %s", err.Error())
			return
		}
	}
	if len(_chimeUserType) > 0 {
		if err := assignInputField(input, "UserType", _chimeUserType); err != nil {
			log.Errorf("invalid --user-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the settings for the specified user, such as phone number settings.
func chime_UpdateUserSettings(cfg aws.Config, client *chime.Client) {
	input := &chime.UpdateUserSettingsInput{
		// AccountId: *string, // Required
		// UserId: *string, // Required
		// UserSettings: *types.UserSettings, // Required
	}

	if len(_chimeAccountId) > 0 {
		input.AccountId = aws.String(_chimeAccountId)
	}
	if len(_chimeUserId) > 0 {
		input.UserId = aws.String(_chimeUserId)
	}
	if len(_chimeUserSettings) > 0 {
		if err := assignInputField(input, "UserSettings", _chimeUserSettings); err != nil {
			log.Errorf("invalid --user-settings: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUserSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_chimeCmd)
	_chimeCmd.Flags().SortFlags = false

	_chimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_chimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_chimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_chimeCmd.Flags().StringVarP(&_chimeAccountId, "account-id", "", "", "Account ID")
	_chimeCmd.Flags().StringVarP(&_chimeAccountSettings, "account-settings", "", "", "Account Settings")
	_chimeCmd.Flags().StringVarP(&_chimeAlexaForBusinessMetadata, "alexa-for-business-metadata", "", "", "Alexa For Business Metadata")
	_chimeCmd.Flags().StringVarP(&_chimeAreaCode, "area-code", "", "", "Area Code")
	_chimeCmd.Flags().StringVarP(&_chimeBotId, "bot-id", "", "", "Bot ID")
	_chimeCmd.Flags().StringVarP(&_chimeBusinessCalling, "business-calling", "", "", "Business Calling")
	_chimeCmd.Flags().StringVarP(&_chimeCallingName, "calling-name", "", "", "Calling Name")
	_chimeCmd.Flags().StringVarP(&_chimeCity, "city", "", "", "City")
	_chimeCmd.Flags().StringVarP(&_chimeClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_chimeCmd.Flags().StringVarP(&_chimeConversationId, "conversation-id", "", "", "Conversation ID")
	_chimeCmd.Flags().StringVarP(&_chimeCountry, "country", "", "", "Country")
	_chimeCmd.Flags().StringVarP(&_chimeDefaultLicense, "default-license", "", "", "Default License")
	_chimeCmd.Flags().StringVarP(&_chimeDisabled, "disabled", "", "", "Disabled")
	_chimeCmd.Flags().StringVarP(&_chimeDisplayName, "display-name", "", "", "Display Name")
	_chimeCmd.Flags().StringVarP(&_chimeDomain, "domain", "", "", "Domain")
	_chimeCmd.Flags().StringVarP(&_chimeE164PhoneNumber, "e164-phone-number", "", "", "E164 Phone Number")
	_chimeCmd.Flags().StringSliceVarP(&_chimeE164PhoneNumbers, "e164-phone-numbers", "", nil, "E164 Phone Numbers")
	_chimeCmd.Flags().StringVarP(&_chimeEmail, "email", "", "", "Email")
	_chimeCmd.Flags().StringVarP(&_chimeFilterName, "filter-name", "", "", "Filter Name")
	_chimeCmd.Flags().StringVarP(&_chimeFilterValue, "filter-value", "", "", "Filter Value")
	_chimeCmd.Flags().StringVarP(&_chimeFromPhoneNumber, "from-phone-number", "", "", "From Phone Number")
	_chimeCmd.Flags().StringSliceVarP(&_chimeGroupNames, "group-names", "", nil, "Group Names")
	_chimeCmd.Flags().StringVarP(&_chimeJoinToken, "join-token", "", "", "Join Token")
	_chimeCmd.Flags().StringVarP(&_chimeLambdaFunctionArn, "lambda-function-arn", "", "", "Lambda Function ARN")
	_chimeCmd.Flags().StringVarP(&_chimeLicenseType, "license-type", "", "", "License Type")
	_chimeCmd.Flags().StringVarP(&_chimeMaxResults, "max-results", "", "", "Max Results")
	_chimeCmd.Flags().StringVarP(&_chimeMeetingId, "meeting-id", "", "", "Meeting ID")
	_chimeCmd.Flags().StringVarP(&_chimeMemberId, "member-id", "", "", "Member ID")
	_chimeCmd.Flags().StringVarP(&_chimeMembershipItemList, "membership-item-list", "", "", "Membership Item List")
	_chimeCmd.Flags().StringVarP(&_chimeMessageId, "message-id", "", "", "Message ID")
	_chimeCmd.Flags().StringVarP(&_chimeName, "name", "", "", "Name")
	_chimeCmd.Flags().StringVarP(&_chimeNextToken, "next-token", "", "", "Next Token")
	_chimeCmd.Flags().StringVarP(&_chimeOutboundEventsHTTPSEndpoint, "outbound-events-https-endpoint", "", "", "Outbound Events HTTPS Endpoint")
	_chimeCmd.Flags().StringVarP(&_chimePhoneNumberId, "phone-number-id", "", "", "Phone Number ID")
	_chimeCmd.Flags().StringSliceVarP(&_chimePhoneNumberIds, "phone-number-ids", "", nil, "Phone Number Ids")
	_chimeCmd.Flags().StringVarP(&_chimePhoneNumberOrderId, "phone-number-order-id", "", "", "Phone Number Order ID")
	_chimeCmd.Flags().StringVarP(&_chimePhoneNumberType, "phone-number-type", "", "", "Phone Number Type")
	_chimeCmd.Flags().StringVarP(&_chimeProductType, "product-type", "", "", "Product Type")
	_chimeCmd.Flags().StringVarP(&_chimeRetentionSettings, "retention-settings", "", "", "Retention Settings")
	_chimeCmd.Flags().StringVarP(&_chimeRole, "role", "", "", "Role")
	_chimeCmd.Flags().StringVarP(&_chimeRoomId, "room-id", "", "", "Room ID")
	_chimeCmd.Flags().StringVarP(&_chimeSigninDelegateGroups, "signin-delegate-groups", "", "", "Signin Delegate Groups")
	_chimeCmd.Flags().StringVarP(&_chimeState, "state", "", "", "State")
	_chimeCmd.Flags().StringVarP(&_chimeStatus, "status", "", "", "Status")
	_chimeCmd.Flags().StringVarP(&_chimeToPhoneNumber, "to-phone-number", "", "", "To Phone Number")
	_chimeCmd.Flags().StringVarP(&_chimeTollFreePrefix, "toll-free-prefix", "", "", "Toll Free Prefix")
	_chimeCmd.Flags().StringVarP(&_chimeUpdatePhoneNumberRequestItems, "update-phone-number-request-items", "", "", "Update Phone Number Request Items")
	_chimeCmd.Flags().StringVarP(&_chimeUpdateUserRequestItems, "update-user-request-items", "", "", "Update User Request Items")
	_chimeCmd.Flags().StringVarP(&_chimeUserEmail, "user-email", "", "", "User Email")
	_chimeCmd.Flags().StringSliceVarP(&_chimeUserEmailList, "user-email-list", "", nil, "User Email List")
	_chimeCmd.Flags().StringVarP(&_chimeUserId, "user-id", "", "", "User ID")
	_chimeCmd.Flags().StringSliceVarP(&_chimeUserIdList, "user-id-list", "", nil, "User ID List")
	_chimeCmd.Flags().StringVarP(&_chimeUserSettings, "user-settings", "", "", "User Settings")
	_chimeCmd.Flags().StringVarP(&_chimeUserType, "user-type", "", "", "User Type")
	_chimeCmd.Flags().StringVarP(&_chimeUsername, "username", "", "", "Username")
	_chimeCmd.Flags().StringVarP(&_chimeVoiceConnector, "voice-connector", "", "", "Voice Connector")

	_chimeCmd.Flags().BoolVarP(&_chimeAssociatePhoneNumberWithUser, "associate-phone-number-with-user", "", false, "Associate Phone Number With User")
	_chimeCmd.Flags().BoolVarP(&_chimeAssociateSigninDelegateGroupsWithAccount, "associate-signin-delegate-groups-with-account", "", false, "Associate Signin Delegate Groups With Account")
	_chimeCmd.Flags().BoolVarP(&_chimeBatchCreateRoomMembership, "batch-create-room-membership", "", false, "Batch Create Room Membership")
	_chimeCmd.Flags().BoolVarP(&_chimeBatchDeletePhoneNumber, "batch-delete-phone-number", "", false, "Batch Delete Phone Number")
	_chimeCmd.Flags().BoolVarP(&_chimeBatchSuspendUser, "batch-suspend-user", "", false, "Batch Suspend User")
	_chimeCmd.Flags().BoolVarP(&_chimeBatchUnsuspendUser, "batch-unsuspend-user", "", false, "Batch Unsuspend User")
	_chimeCmd.Flags().BoolVarP(&_chimeBatchUpdatePhoneNumber, "batch-update-phone-number", "", false, "Batch Update Phone Number")
	_chimeCmd.Flags().BoolVarP(&_chimeBatchUpdateUser, "batch-update-user", "", false, "Batch Update User")
	_chimeCmd.Flags().BoolVarP(&_chimeCreateAccount, "create-account", "", false, "Create Account")
	_chimeCmd.Flags().BoolVarP(&_chimeCreateBot, "create-bot", "", false, "Create Bot")
	_chimeCmd.Flags().BoolVarP(&_chimeCreateMeetingDialOut, "create-meeting-dial-out", "", false, "Create Meeting Dial Out")
	_chimeCmd.Flags().BoolVarP(&_chimeCreatePhoneNumberOrder, "create-phone-number-order", "", false, "Create Phone Number Order")
	_chimeCmd.Flags().BoolVarP(&_chimeCreateRoom, "create-room", "", false, "Create Room")
	_chimeCmd.Flags().BoolVarP(&_chimeCreateRoomMembership, "create-room-membership", "", false, "Create Room Membership")
	_chimeCmd.Flags().BoolVarP(&_chimeCreateUser, "create-user", "", false, "Create User")
	_chimeCmd.Flags().BoolVarP(&_chimeDeleteAccount, "delete-account", "", false, "Delete Account")
	_chimeCmd.Flags().BoolVarP(&_chimeDeleteEventsConfiguration, "delete-events-configuration", "", false, "Delete Events Configuration")
	_chimeCmd.Flags().BoolVarP(&_chimeDeletePhoneNumber, "delete-phone-number", "", false, "Delete Phone Number")
	_chimeCmd.Flags().BoolVarP(&_chimeDeleteRoom, "delete-room", "", false, "Delete Room")
	_chimeCmd.Flags().BoolVarP(&_chimeDeleteRoomMembership, "delete-room-membership", "", false, "Delete Room Membership")
	_chimeCmd.Flags().BoolVarP(&_chimeDisassociatePhoneNumberFromUser, "disassociate-phone-number-from-user", "", false, "Disassociate Phone Number From User")
	_chimeCmd.Flags().BoolVarP(&_chimeDisassociateSigninDelegateGroupsFromAccount, "disassociate-signin-delegate-groups-from-account", "", false, "Disassociate Signin Delegate Groups From Account")
	_chimeCmd.Flags().BoolVarP(&_chimeGetAccount, "get-account", "", false, "Get Account")
	_chimeCmd.Flags().BoolVarP(&_chimeGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeGetBot, "get-bot", "", false, "Get Bot")
	_chimeCmd.Flags().BoolVarP(&_chimeGetEventsConfiguration, "get-events-configuration", "", false, "Get Events Configuration")
	_chimeCmd.Flags().BoolVarP(&_chimeGetGlobalSettings, "get-global-settings", "", false, "Get Global Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeGetPhoneNumber, "get-phone-number", "", false, "Get Phone Number")
	_chimeCmd.Flags().BoolVarP(&_chimeGetPhoneNumberOrder, "get-phone-number-order", "", false, "Get Phone Number Order")
	_chimeCmd.Flags().BoolVarP(&_chimeGetPhoneNumberSettings, "get-phone-number-settings", "", false, "Get Phone Number Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeGetRetentionSettings, "get-retention-settings", "", false, "Get Retention Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeGetRoom, "get-room", "", false, "Get Room")
	_chimeCmd.Flags().BoolVarP(&_chimeGetUser, "get-user", "", false, "Get User")
	_chimeCmd.Flags().BoolVarP(&_chimeGetUserSettings, "get-user-settings", "", false, "Get User Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeInviteUsers, "invite-users", "", false, "Invite Users")
	_chimeCmd.Flags().BoolVarP(&_chimeListAccounts, "list-accounts", "", false, "List Accounts")
	_chimeCmd.Flags().BoolVarP(&_chimeListBots, "list-bots", "", false, "List Bots")
	_chimeCmd.Flags().BoolVarP(&_chimeListPhoneNumberOrders, "list-phone-number-orders", "", false, "List Phone Number Orders")
	_chimeCmd.Flags().BoolVarP(&_chimeListPhoneNumbers, "list-phone-numbers", "", false, "List Phone Numbers")
	_chimeCmd.Flags().BoolVarP(&_chimeListRoomMemberships, "list-room-memberships", "", false, "List Room Memberships")
	_chimeCmd.Flags().BoolVarP(&_chimeListRooms, "list-rooms", "", false, "List Rooms")
	_chimeCmd.Flags().BoolVarP(&_chimeListSupportedPhoneNumberCountries, "list-supported-phone-number-countries", "", false, "List Supported Phone Number Countries")
	_chimeCmd.Flags().BoolVarP(&_chimeListUsers, "list-users", "", false, "List Users")
	_chimeCmd.Flags().BoolVarP(&_chimeLogoutUser, "logout-user", "", false, "Logout User")
	_chimeCmd.Flags().BoolVarP(&_chimePutEventsConfiguration, "put-events-configuration", "", false, "Put Events Configuration")
	_chimeCmd.Flags().BoolVarP(&_chimePutRetentionSettings, "put-retention-settings", "", false, "Put Retention Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeRedactConversationMessage, "redact-conversation-message", "", false, "Redact Conversation Message")
	_chimeCmd.Flags().BoolVarP(&_chimeRedactRoomMessage, "redact-room-message", "", false, "Redact Room Message")
	_chimeCmd.Flags().BoolVarP(&_chimeRegenerateSecurityToken, "regenerate-security-token", "", false, "Regenerate Security Token")
	_chimeCmd.Flags().BoolVarP(&_chimeResetPersonalPIN, "reset-personal-pin", "", false, "Reset Personal Pin")
	_chimeCmd.Flags().BoolVarP(&_chimeRestorePhoneNumber, "restore-phone-number", "", false, "Restore Phone Number")
	_chimeCmd.Flags().BoolVarP(&_chimeSearchAvailablePhoneNumbers, "search-available-phone-numbers", "", false, "Search Available Phone Numbers")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateAccount, "update-account", "", false, "Update Account")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateBot, "update-bot", "", false, "Update Bot")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateGlobalSettings, "update-global-settings", "", false, "Update Global Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdatePhoneNumber, "update-phone-number", "", false, "Update Phone Number")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdatePhoneNumberSettings, "update-phone-number-settings", "", false, "Update Phone Number Settings")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateRoom, "update-room", "", false, "Update Room")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateRoomMembership, "update-room-membership", "", false, "Update Room Membership")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateUser, "update-user", "", false, "Update User")
	_chimeCmd.Flags().BoolVarP(&_chimeUpdateUserSettings, "update-user-settings", "", false, "Update User Settings")

}
