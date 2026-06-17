package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mturk"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mturkCmd represents the mturk command
var _mturkCmd = &cobra.Command{
	Use:   "mturk",
	Short: "AWS mturk CLI",
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
		client := mturk.NewFromConfig(cfg)
		if _mturkAcceptQualificationRequest {
			mturk_AcceptQualificationRequest(cfg, client)
			return
		}
		if _mturkApproveAssignment {
			mturk_ApproveAssignment(cfg, client)
			return
		}
		if _mturkAssociateQualificationWithWorker {
			mturk_AssociateQualificationWithWorker(cfg, client)
			return
		}
		if _mturkCreateAdditionalAssignmentsForHIT {
			mturk_CreateAdditionalAssignmentsForHIT(cfg, client)
			return
		}
		if _mturkCreateHIT {
			mturk_CreateHIT(cfg, client)
			return
		}
		if _mturkCreateHITType {
			mturk_CreateHITType(cfg, client)
			return
		}
		if _mturkCreateHITWithHITType {
			mturk_CreateHITWithHITType(cfg, client)
			return
		}
		if _mturkCreateQualificationType {
			mturk_CreateQualificationType(cfg, client)
			return
		}
		if _mturkCreateWorkerBlock {
			mturk_CreateWorkerBlock(cfg, client)
			return
		}
		if _mturkDeleteHIT {
			mturk_DeleteHIT(cfg, client)
			return
		}
		if _mturkDeleteQualificationType {
			mturk_DeleteQualificationType(cfg, client)
			return
		}
		if _mturkDeleteWorkerBlock {
			mturk_DeleteWorkerBlock(cfg, client)
			return
		}
		if _mturkDisassociateQualificationFromWorker {
			mturk_DisassociateQualificationFromWorker(cfg, client)
			return
		}
		if _mturkGetAccountBalance {
			mturk_GetAccountBalance(cfg, client)
			return
		}
		if _mturkGetAssignment {
			mturk_GetAssignment(cfg, client)
			return
		}
		if _mturkGetFileUploadURL {
			mturk_GetFileUploadURL(cfg, client)
			return
		}
		if _mturkGetHIT {
			mturk_GetHIT(cfg, client)
			return
		}
		if _mturkGetQualificationScore {
			mturk_GetQualificationScore(cfg, client)
			return
		}
		if _mturkGetQualificationType {
			mturk_GetQualificationType(cfg, client)
			return
		}
		if _mturkListAssignmentsForHIT {
			mturk_ListAssignmentsForHIT(cfg, client)
			return
		}
		if _mturkListBonusPayments {
			mturk_ListBonusPayments(cfg, client)
			return
		}
		if _mturkListHITs {
			mturk_ListHITs(cfg, client)
			return
		}
		if _mturkListHITsForQualificationType {
			mturk_ListHITsForQualificationType(cfg, client)
			return
		}
		if _mturkListQualificationRequests {
			mturk_ListQualificationRequests(cfg, client)
			return
		}
		if _mturkListQualificationTypes {
			mturk_ListQualificationTypes(cfg, client)
			return
		}
		if _mturkListReviewPolicyResultsForHIT {
			mturk_ListReviewPolicyResultsForHIT(cfg, client)
			return
		}
		if _mturkListReviewableHITs {
			mturk_ListReviewableHITs(cfg, client)
			return
		}
		if _mturkListWorkerBlocks {
			mturk_ListWorkerBlocks(cfg, client)
			return
		}
		if _mturkListWorkersWithQualificationType {
			mturk_ListWorkersWithQualificationType(cfg, client)
			return
		}
		if _mturkNotifyWorkers {
			mturk_NotifyWorkers(cfg, client)
			return
		}
		if _mturkRejectAssignment {
			mturk_RejectAssignment(cfg, client)
			return
		}
		if _mturkRejectQualificationRequest {
			mturk_RejectQualificationRequest(cfg, client)
			return
		}
		if _mturkSendBonus {
			mturk_SendBonus(cfg, client)
			return
		}
		if _mturkSendTestEventNotification {
			mturk_SendTestEventNotification(cfg, client)
			return
		}
		if _mturkUpdateExpirationForHIT {
			mturk_UpdateExpirationForHIT(cfg, client)
			return
		}
		if _mturkUpdateHITReviewStatus {
			mturk_UpdateHITReviewStatus(cfg, client)
			return
		}
		if _mturkUpdateHITTypeOfHIT {
			mturk_UpdateHITTypeOfHIT(cfg, client)
			return
		}
		if _mturkUpdateNotificationSettings {
			mturk_UpdateNotificationSettings(cfg, client)
			return
		}
		if _mturkUpdateQualificationType {
			mturk_UpdateQualificationType(cfg, client)
			return
		}

	},
}

var (
	_mturkAcceptQualificationRequest          bool
	_mturkApproveAssignment                   bool
	_mturkAssociateQualificationWithWorker    bool
	_mturkCreateAdditionalAssignmentsForHIT   bool
	_mturkCreateHIT                           bool
	_mturkCreateHITType                       bool
	_mturkCreateHITWithHITType                bool
	_mturkCreateQualificationType             bool
	_mturkCreateWorkerBlock                   bool
	_mturkDeleteHIT                           bool
	_mturkDeleteQualificationType             bool
	_mturkDeleteWorkerBlock                   bool
	_mturkDisassociateQualificationFromWorker bool
	_mturkGetAccountBalance                   bool
	_mturkGetAssignment                       bool
	_mturkGetFileUploadURL                    bool
	_mturkGetHIT                              bool
	_mturkGetQualificationScore               bool
	_mturkGetQualificationType                bool
	_mturkListAssignmentsForHIT               bool
	_mturkListBonusPayments                   bool
	_mturkListHITs                            bool
	_mturkListHITsForQualificationType        bool
	_mturkListQualificationRequests           bool
	_mturkListQualificationTypes              bool
	_mturkListReviewPolicyResultsForHIT       bool
	_mturkListReviewableHITs                  bool
	_mturkListWorkerBlocks                    bool
	_mturkListWorkersWithQualificationType    bool
	_mturkNotifyWorkers                       bool
	_mturkRejectAssignment                    bool
	_mturkRejectQualificationRequest          bool
	_mturkSendBonus                           bool
	_mturkSendTestEventNotification           bool
	_mturkUpdateExpirationForHIT              bool
	_mturkUpdateHITReviewStatus               bool
	_mturkUpdateHITTypeOfHIT                  bool
	_mturkUpdateNotificationSettings          bool
	_mturkUpdateQualificationType             bool

	_mturkActive                        string
	_mturkAnswerKey                     string
	_mturkAssignmentDurationInSeconds   string
	_mturkAssignmentId                  string
	_mturkAssignmentReviewPolicy        string
	_mturkAssignmentStatuses            string
	_mturkAutoApprovalDelayInSeconds    string
	_mturkAutoGranted                   string
	_mturkAutoGrantedValue              string
	_mturkBonusAmount                   string
	_mturkDescription                   string
	_mturkExpireAt                      string
	_mturkHITLayoutId                   string
	_mturkHITLayoutParameters           string
	_mturkHITReviewPolicy               string
	_mturkHITTypeId                     string
	_mturkHITId                         string
	_mturkIntegerValue                  string
	_mturkKeywords                      string
	_mturkLifetimeInSeconds             string
	_mturkMaxAssignments                string
	_mturkMaxResults                    string
	_mturkMessageText                   string
	_mturkMustBeOwnedByCaller           string
	_mturkMustBeRequestable             string
	_mturkName                          string
	_mturkNextToken                     string
	_mturkNotification                  string
	_mturkNumberOfAdditionalAssignments string
	_mturkOverrideRejection             string
	_mturkPolicyLevels                  string
	_mturkQualificationRequestId        string
	_mturkQualificationRequirements     string
	_mturkQualificationTypeId           string
	_mturkQualificationTypeStatus       string
	_mturkQuery                         string
	_mturkQuestion                      string
	_mturkQuestionIdentifier            string
	_mturkReason                        string
	_mturkRequesterAnnotation           string
	_mturkRequesterFeedback             string
	_mturkRetrieveActions               string
	_mturkRetrieveResults               string
	_mturkRetryDelayInSeconds           string
	_mturkRevert                        string
	_mturkReward                        string
	_mturkSendNotification              string
	_mturkStatus                        string
	_mturkSubject                       string
	_mturkTest                          string
	_mturkTestDurationInSeconds         string
	_mturkTestEventType                 string
	_mturkTitle                         string
	_mturkUniqueRequestToken            string
	_mturkWorkerId                      string
	_mturkWorkerIds                     []string
)

// The AcceptQualificationRequest operation approves a Worker's request for a
// Qualification.
//
// Only the owner of the Qualification type can grant a Qualification request for
// that type.
//
// A successful request for the AcceptQualificationRequest operation returns with
// no errors and an empty body.
func mturk_AcceptQualificationRequest(cfg aws.Config, client *mturk.Client) {
	input := &mturk.AcceptQualificationRequestInput{
		// QualificationRequestId: *string, // Required
	}

	if len(_mturkQualificationRequestId) > 0 {
		input.QualificationRequestId = aws.String(_mturkQualificationRequestId)
	}
	if len(_mturkIntegerValue) > 0 {
		if err := assignInputField(input, "IntegerValue", _mturkIntegerValue); err != nil {
			log.Errorf("invalid --integer-value: %s", err.Error())
			return
		}
	}

	if resp, err := client.AcceptQualificationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The ApproveAssignment operation approves the results of a completed
// assignment.
//
// Approving an assignment initiates two payments from the Requester's Amazon.com
// account
//
// - The Worker who submitted the results is paid the reward specified in the
// HIT.
//
// - Amazon Mechanical Turk fees are debited.
//
// If the Requester's account does not have adequate funds for these payments, the
// call to ApproveAssignment returns an exception, and the approval is not
// processed. You can include an optional feedback message with the approval, which
// the Worker can see in the Status section of the web site.
//
// You can also call this operation for assignments that were previous rejected
// and approve them by explicitly overriding the previous rejection. This only
// works on rejected assignments that were submitted within the previous 30 days
// and only if the assignment's related HIT has not been deleted.
func mturk_ApproveAssignment(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ApproveAssignmentInput{
		// AssignmentId: *string, // Required
	}

	if len(_mturkAssignmentId) > 0 {
		input.AssignmentId = aws.String(_mturkAssignmentId)
	}
	if len(_mturkOverrideRejection) > 0 {
		if err := assignInputField(input, "OverrideRejection", _mturkOverrideRejection); err != nil {
			log.Errorf("invalid --override-rejection: %s", err.Error())
			return
		}
	}
	if len(_mturkRequesterFeedback) > 0 {
		input.RequesterFeedback = aws.String(_mturkRequesterFeedback)
	}

	if resp, err := client.ApproveAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The AssociateQualificationWithWorker operation gives a Worker a Qualification.
// AssociateQualificationWithWorker does not require that the Worker submit a
// Qualification request. It gives the Qualification directly to the Worker.
//
// You can only assign a Qualification of a Qualification type that you created
// (using the CreateQualificationType operation).
//
// Note: AssociateQualificationWithWorker does not affect any pending
// Qualification requests for the Qualification by the Worker. If you assign a
// Qualification to a Worker, then later grant a Qualification request made by the
// Worker, the granting of the request may modify the Qualification score. To
// resolve a pending Qualification request without affecting the Qualification the
// Worker already has, reject the request with the RejectQualificationRequest
// operation.
func mturk_AssociateQualificationWithWorker(cfg aws.Config, client *mturk.Client) {
	input := &mturk.AssociateQualificationWithWorkerInput{
		// QualificationTypeId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}
	if len(_mturkWorkerId) > 0 {
		input.WorkerId = aws.String(_mturkWorkerId)
	}
	if len(_mturkIntegerValue) > 0 {
		if err := assignInputField(input, "IntegerValue", _mturkIntegerValue); err != nil {
			log.Errorf("invalid --integer-value: %s", err.Error())
			return
		}
	}
	if len(_mturkSendNotification) > 0 {
		if err := assignInputField(input, "SendNotification", _mturkSendNotification); err != nil {
			log.Errorf("invalid --send-notification: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateQualificationWithWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateAdditionalAssignmentsForHIT operation increases the maximum number
// of assignments of an existing HIT.
//
// To extend the maximum number of assignments, specify the number of additional
// assignments.
//
// - HITs created with fewer than 10 assignments cannot be extended to have 10
// or more assignments. Attempting to add assignments in a way that brings the
// total number of assignments for a HIT from fewer than 10 assignments to 10 or
// more assignments will result in an
// AWS.MechanicalTurk.InvalidMaximumAssignmentsIncrease exception.
//
// - HITs that were created before July 22, 2015 cannot be extended. Attempting
// to extend HITs that were created before July 22, 2015 will result in an
// AWS.MechanicalTurk.HITTooOldForExtension exception.
func mturk_CreateAdditionalAssignmentsForHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.CreateAdditionalAssignmentsForHITInput{
		// HITId: *string, // Required
		// NumberOfAdditionalAssignments: *int32, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}
	if len(_mturkNumberOfAdditionalAssignments) > 0 {
		if err := assignInputField(input, "NumberOfAdditionalAssignments", _mturkNumberOfAdditionalAssignments); err != nil {
			log.Errorf("invalid --number-of-additional-assignments: %s", err.Error())
			return
		}
	}
	if len(_mturkUniqueRequestToken) > 0 {
		input.UniqueRequestToken = aws.String(_mturkUniqueRequestToken)
	}

	if resp, err := client.CreateAdditionalAssignmentsForHIT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateHIT operation creates a new Human Intelligence Task (HIT). The new
// HIT is made available for Workers to find and accept on the Amazon Mechanical
// Turk website.
//
// This operation allows you to specify a new HIT by passing in values for the
// properties of the HIT, such as its title, reward amount and number of
// assignments. When you pass these values to CreateHIT , a new HIT is created for
// you, with a new HITTypeID . The HITTypeID can be used to create additional HITs
// in the future without needing to specify common parameters such as the title,
// description and reward amount each time.
//
// An alternative way to create HITs is to first generate a HITTypeID using the
// CreateHITType operation and then call the CreateHITWithHITType operation. This
// is the recommended best practice for Requesters who are creating large numbers
// of HITs.
//
// CreateHIT also supports several ways to provide question data: by providing a
// value for the Question parameter that fully specifies the contents of the HIT,
// or by providing a HitLayoutId and associated HitLayoutParameters .
//
// If a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see [Amazon Mechanical Turk Pricing].
//
// [Amazon Mechanical Turk Pricing]: https://requester.mturk.com/pricing
func mturk_CreateHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.CreateHITInput{
		// AssignmentDurationInSeconds: *int64, // Required
		// Description: *string, // Required
		// LifetimeInSeconds: *int64, // Required
		// Reward: *string, // Required
		// Title: *string, // Required
	}

	if len(_mturkAssignmentDurationInSeconds) > 0 {
		if err := assignInputField(input, "AssignmentDurationInSeconds", _mturkAssignmentDurationInSeconds); err != nil {
			log.Errorf("invalid --assignment-duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkDescription) > 0 {
		input.Description = aws.String(_mturkDescription)
	}
	if len(_mturkLifetimeInSeconds) > 0 {
		if err := assignInputField(input, "LifetimeInSeconds", _mturkLifetimeInSeconds); err != nil {
			log.Errorf("invalid --lifetime-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkReward) > 0 {
		input.Reward = aws.String(_mturkReward)
	}
	if len(_mturkTitle) > 0 {
		input.Title = aws.String(_mturkTitle)
	}
	if len(_mturkAssignmentReviewPolicy) > 0 {
		if err := assignInputField(input, "AssignmentReviewPolicy", _mturkAssignmentReviewPolicy); err != nil {
			log.Errorf("invalid --assignment-review-policy: %s", err.Error())
			return
		}
	}
	if len(_mturkAutoApprovalDelayInSeconds) > 0 {
		if err := assignInputField(input, "AutoApprovalDelayInSeconds", _mturkAutoApprovalDelayInSeconds); err != nil {
			log.Errorf("invalid --auto-approval-delay-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkHITLayoutId) > 0 {
		input.HITLayoutId = aws.String(_mturkHITLayoutId)
	}
	if len(_mturkHITLayoutParameters) > 0 {
		if err := assignInputField(input, "HITLayoutParameters", _mturkHITLayoutParameters); err != nil {
			log.Errorf("invalid --hit-layout-parameters: %s", err.Error())
			return
		}
	}
	if len(_mturkHITReviewPolicy) > 0 {
		if err := assignInputField(input, "HITReviewPolicy", _mturkHITReviewPolicy); err != nil {
			log.Errorf("invalid --hit-review-policy: %s", err.Error())
			return
		}
	}
	if len(_mturkKeywords) > 0 {
		input.Keywords = aws.String(_mturkKeywords)
	}
	if len(_mturkMaxAssignments) > 0 {
		if err := assignInputField(input, "MaxAssignments", _mturkMaxAssignments); err != nil {
			log.Errorf("invalid --max-assignments: %s", err.Error())
			return
		}
	}
	if len(_mturkQualificationRequirements) > 0 {
		if err := assignInputField(input, "QualificationRequirements", _mturkQualificationRequirements); err != nil {
			log.Errorf("invalid --qualification-requirements: %s", err.Error())
			return
		}
	}
	if len(_mturkQuestion) > 0 {
		input.Question = aws.String(_mturkQuestion)
	}
	if len(_mturkRequesterAnnotation) > 0 {
		input.RequesterAnnotation = aws.String(_mturkRequesterAnnotation)
	}
	if len(_mturkUniqueRequestToken) > 0 {
		input.UniqueRequestToken = aws.String(_mturkUniqueRequestToken)
	}

	if resp, err := client.CreateHIT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateHITType operation creates a new HIT type. This operation allows you
// to define a standard set of HIT properties to use when creating HITs. If you
// register a HIT type with values that match an existing HIT type, the HIT type ID
// of the existing type will be returned.
func mturk_CreateHITType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.CreateHITTypeInput{
		// AssignmentDurationInSeconds: *int64, // Required
		// Description: *string, // Required
		// Reward: *string, // Required
		// Title: *string, // Required
	}

	if len(_mturkAssignmentDurationInSeconds) > 0 {
		if err := assignInputField(input, "AssignmentDurationInSeconds", _mturkAssignmentDurationInSeconds); err != nil {
			log.Errorf("invalid --assignment-duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkDescription) > 0 {
		input.Description = aws.String(_mturkDescription)
	}
	if len(_mturkReward) > 0 {
		input.Reward = aws.String(_mturkReward)
	}
	if len(_mturkTitle) > 0 {
		input.Title = aws.String(_mturkTitle)
	}
	if len(_mturkAutoApprovalDelayInSeconds) > 0 {
		if err := assignInputField(input, "AutoApprovalDelayInSeconds", _mturkAutoApprovalDelayInSeconds); err != nil {
			log.Errorf("invalid --auto-approval-delay-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkKeywords) > 0 {
		input.Keywords = aws.String(_mturkKeywords)
	}
	if len(_mturkQualificationRequirements) > 0 {
		if err := assignInputField(input, "QualificationRequirements", _mturkQualificationRequirements); err != nil {
			log.Errorf("invalid --qualification-requirements: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHITType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateHITWithHITType operation creates a new Human Intelligence Task (HIT)
// using an existing HITTypeID generated by the CreateHITType operation.
//
// This is an alternative way to create HITs from the CreateHIT operation. This is
// the recommended best practice for Requesters who are creating large numbers of
// HITs.
//
// CreateHITWithHITType also supports several ways to provide question data: by
// providing a value for the Question parameter that fully specifies the contents
// of the HIT, or by providing a HitLayoutId and associated HitLayoutParameters .
//
// If a HIT is created with 10 or more maximum assignments, there is an additional
// fee. For more information, see [Amazon Mechanical Turk Pricing].
//
// [Amazon Mechanical Turk Pricing]: https://requester.mturk.com/pricing
func mturk_CreateHITWithHITType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.CreateHITWithHITTypeInput{
		// HITTypeId: *string, // Required
		// LifetimeInSeconds: *int64, // Required
	}

	if len(_mturkHITTypeId) > 0 {
		input.HITTypeId = aws.String(_mturkHITTypeId)
	}
	if len(_mturkLifetimeInSeconds) > 0 {
		if err := assignInputField(input, "LifetimeInSeconds", _mturkLifetimeInSeconds); err != nil {
			log.Errorf("invalid --lifetime-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkAssignmentReviewPolicy) > 0 {
		if err := assignInputField(input, "AssignmentReviewPolicy", _mturkAssignmentReviewPolicy); err != nil {
			log.Errorf("invalid --assignment-review-policy: %s", err.Error())
			return
		}
	}
	if len(_mturkHITLayoutId) > 0 {
		input.HITLayoutId = aws.String(_mturkHITLayoutId)
	}
	if len(_mturkHITLayoutParameters) > 0 {
		if err := assignInputField(input, "HITLayoutParameters", _mturkHITLayoutParameters); err != nil {
			log.Errorf("invalid --hit-layout-parameters: %s", err.Error())
			return
		}
	}
	if len(_mturkHITReviewPolicy) > 0 {
		if err := assignInputField(input, "HITReviewPolicy", _mturkHITReviewPolicy); err != nil {
			log.Errorf("invalid --hit-review-policy: %s", err.Error())
			return
		}
	}
	if len(_mturkMaxAssignments) > 0 {
		if err := assignInputField(input, "MaxAssignments", _mturkMaxAssignments); err != nil {
			log.Errorf("invalid --max-assignments: %s", err.Error())
			return
		}
	}
	if len(_mturkQuestion) > 0 {
		input.Question = aws.String(_mturkQuestion)
	}
	if len(_mturkRequesterAnnotation) > 0 {
		input.RequesterAnnotation = aws.String(_mturkRequesterAnnotation)
	}
	if len(_mturkUniqueRequestToken) > 0 {
		input.UniqueRequestToken = aws.String(_mturkUniqueRequestToken)
	}

	if resp, err := client.CreateHITWithHITType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateQualificationType operation creates a new Qualification type, which
// is represented by a QualificationType data structure.
func mturk_CreateQualificationType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.CreateQualificationTypeInput{
		// Description: *string, // Required
		// Name: *string, // Required
		// QualificationTypeStatus: types.QualificationTypeStatus, // Required
	}

	if len(_mturkDescription) > 0 {
		input.Description = aws.String(_mturkDescription)
	}
	if len(_mturkName) > 0 {
		input.Name = aws.String(_mturkName)
	}
	if len(_mturkQualificationTypeStatus) > 0 {
		if err := assignInputField(input, "QualificationTypeStatus", _mturkQualificationTypeStatus); err != nil {
			log.Errorf("invalid --qualification-type-status: %s", err.Error())
			return
		}
	}
	if len(_mturkAnswerKey) > 0 {
		input.AnswerKey = aws.String(_mturkAnswerKey)
	}
	if len(_mturkAutoGranted) > 0 {
		if err := assignInputField(input, "AutoGranted", _mturkAutoGranted); err != nil {
			log.Errorf("invalid --auto-granted: %s", err.Error())
			return
		}
	}
	if len(_mturkAutoGrantedValue) > 0 {
		if err := assignInputField(input, "AutoGrantedValue", _mturkAutoGrantedValue); err != nil {
			log.Errorf("invalid --auto-granted-value: %s", err.Error())
			return
		}
	}
	if len(_mturkKeywords) > 0 {
		input.Keywords = aws.String(_mturkKeywords)
	}
	if len(_mturkRetryDelayInSeconds) > 0 {
		if err := assignInputField(input, "RetryDelayInSeconds", _mturkRetryDelayInSeconds); err != nil {
			log.Errorf("invalid --retry-delay-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkTest) > 0 {
		input.Test = aws.String(_mturkTest)
	}
	if len(_mturkTestDurationInSeconds) > 0 {
		if err := assignInputField(input, "TestDurationInSeconds", _mturkTestDurationInSeconds); err != nil {
			log.Errorf("invalid --test-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateQualificationType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateWorkerBlock operation allows you to prevent a Worker from working on
// your HITs. For example, you can block a Worker who is producing poor quality
// work. You can block up to 100,000 Workers.
func mturk_CreateWorkerBlock(cfg aws.Config, client *mturk.Client) {
	input := &mturk.CreateWorkerBlockInput{
		// Reason: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_mturkReason) > 0 {
		input.Reason = aws.String(_mturkReason)
	}
	if len(_mturkWorkerId) > 0 {
		input.WorkerId = aws.String(_mturkWorkerId)
	}

	if resp, err := client.CreateWorkerBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteHIT operation is used to delete HIT that is no longer needed. Only
// the Requester who created the HIT can delete it.
//
// You can only dispose of HITs that are in the Reviewable state, with all of
// their submitted assignments already either approved or rejected. If you call the
// DeleteHIT operation on a HIT that is not in the Reviewable state (for example,
// that has not expired, or still has active assignments), or on a HIT that is
// Reviewable but without all of its submitted assignments already approved or
// rejected, the service will return an error.
//
// - HITs are automatically disposed of after 120 days.
//
// - After you dispose of a HIT, you can no longer approve the HIT's rejected
// assignments.
//
// - Disposed HITs are not returned in results for the ListHITs operation.
//
// - Disposing HITs can improve the performance of operations such as
// ListReviewableHITs and ListHITs.
func mturk_DeleteHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.DeleteHITInput{
		// HITId: *string, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}

	if resp, err := client.DeleteHIT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteQualificationType deletes a Qualification type and deletes any HIT
// types that are associated with the Qualification type.
//
// This operation does not revoke Qualifications already assigned to Workers
// because the Qualifications might be needed for active HITs. If there are any
// pending requests for the Qualification type, Amazon Mechanical Turk rejects
// those requests. After you delete a Qualification type, you can no longer use it
// to create HITs or HIT types.
//
// DeleteQualificationType must wait for all the HITs that use the deleted
// Qualification type to be deleted before completing. It may take up to 48 hours
// before DeleteQualificationType completes and the unique name of the
// Qualification type is available for reuse with CreateQualificationType.
func mturk_DeleteQualificationType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.DeleteQualificationTypeInput{
		// QualificationTypeId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}

	if resp, err := client.DeleteQualificationType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteWorkerBlock operation allows you to reinstate a blocked Worker to
// work on your HITs. This operation reverses the effects of the CreateWorkerBlock
// operation. You need the Worker ID to use this operation. If the Worker ID is
// missing or invalid, this operation fails and returns the message “WorkerId is
// invalid.” If the specified Worker is not blocked, this operation returns
// successfully.
func mturk_DeleteWorkerBlock(cfg aws.Config, client *mturk.Client) {
	input := &mturk.DeleteWorkerBlockInput{
		// WorkerId: *string, // Required
	}

	if len(_mturkWorkerId) > 0 {
		input.WorkerId = aws.String(_mturkWorkerId)
	}
	if len(_mturkReason) > 0 {
		input.Reason = aws.String(_mturkReason)
	}

	if resp, err := client.DeleteWorkerBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DisassociateQualificationFromWorker revokes a previously granted
// Qualification from a user.
//
// You can provide a text message explaining why the Qualification was revoked.
// The user who had the Qualification can see this message.
func mturk_DisassociateQualificationFromWorker(cfg aws.Config, client *mturk.Client) {
	input := &mturk.DisassociateQualificationFromWorkerInput{
		// QualificationTypeId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}
	if len(_mturkWorkerId) > 0 {
		input.WorkerId = aws.String(_mturkWorkerId)
	}
	if len(_mturkReason) > 0 {
		input.Reason = aws.String(_mturkReason)
	}

	if resp, err := client.DisassociateQualificationFromWorker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetAccountBalance operation retrieves the Prepaid HITs balance in your
// Amazon Mechanical Turk account if you are a Prepaid Requester. Alternatively,
// this operation will retrieve the remaining available AWS Billing usage if you
// have enabled AWS Billing. Note: If you have enabled AWS Billing and still have a
// remaining Prepaid HITs balance, this balance can be viewed on the My Account
// page in the Requester console.
func mturk_GetAccountBalance(cfg aws.Config, client *mturk.Client) {
	input := &mturk.GetAccountBalanceInput{}

	if resp, err := client.GetAccountBalance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetAssignment operation retrieves the details of the specified Assignment.
func mturk_GetAssignment(cfg aws.Config, client *mturk.Client) {
	input := &mturk.GetAssignmentInput{
		// AssignmentId: *string, // Required
	}

	if len(_mturkAssignmentId) > 0 {
		input.AssignmentId = aws.String(_mturkAssignmentId)
	}

	if resp, err := client.GetAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetFileUploadURL operation generates and returns a temporary URL. You use
// the temporary URL to retrieve a file uploaded by a Worker as an answer to a
// FileUploadAnswer question for a HIT. The temporary URL is generated the instant
// the GetFileUploadURL operation is called, and is valid for 60 seconds. You can
// get a temporary file upload URL any time until the HIT is disposed. After the
// HIT is disposed, any uploaded files are deleted, and cannot be retrieved.
//
// Pending Deprecation on December 12, 2017. The Answer Specification
//
// structure will no longer support the FileUploadAnswer element to be used for
// the QuestionForm data structure. Instead, we recommend that Requesters who want
// to create HITs asking Workers to upload files to use Amazon S3.
func mturk_GetFileUploadURL(cfg aws.Config, client *mturk.Client) {
	input := &mturk.GetFileUploadURLInput{
		// AssignmentId: *string, // Required
		// QuestionIdentifier: *string, // Required
	}

	if len(_mturkAssignmentId) > 0 {
		input.AssignmentId = aws.String(_mturkAssignmentId)
	}
	if len(_mturkQuestionIdentifier) > 0 {
		input.QuestionIdentifier = aws.String(_mturkQuestionIdentifier)
	}

	if resp, err := client.GetFileUploadURL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetHIT operation retrieves the details of the specified HIT.
func mturk_GetHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.GetHITInput{
		// HITId: *string, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}

	if resp, err := client.GetHIT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetQualificationScore operation returns the value of a Worker's
// Qualification for a given Qualification type.
//
// To get a Worker's Qualification, you must know the Worker's ID. The Worker's ID
// is included in the assignment data returned by the ListAssignmentsForHIT
// operation.
//
// Only the owner of a Qualification type can query the value of a Worker's
// Qualification of that type.
func mturk_GetQualificationScore(cfg aws.Config, client *mturk.Client) {
	input := &mturk.GetQualificationScoreInput{
		// QualificationTypeId: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}
	if len(_mturkWorkerId) > 0 {
		input.WorkerId = aws.String(_mturkWorkerId)
	}

	if resp, err := client.GetQualificationScore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetQualificationType operation retrieves information about a Qualification
// type using its ID.
func mturk_GetQualificationType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.GetQualificationTypeInput{
		// QualificationTypeId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}

	if resp, err := client.GetQualificationType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The ListAssignmentsForHIT operation retrieves completed assignments for a HIT.
// You can use this operation to retrieve the results for a HIT.
//
// You can get assignments for a HIT at any time, even if the HIT is not yet
// Reviewable. If a HIT requested multiple assignments, and has received some
// results but has not yet become Reviewable, you can still retrieve the partial
// results with this operation.
//
// Use the AssignmentStatus parameter to control which set of assignments for a
// HIT are returned. The ListAssignmentsForHIT operation can return submitted
// assignments awaiting approval, or it can return assignments that have already
// been approved or rejected. You can set AssignmentStatus=Approved,Rejected to get
// assignments that have already been approved and rejected together in one result
// set.
//
// Only the Requester who created the HIT can retrieve the assignments for that
// HIT.
//
// Results are sorted and divided into numbered pages and the operation returns a
// single page of results. You can use the parameters of the operation to control
// sorting and pagination.
func mturk_ListAssignmentsForHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListAssignmentsForHITInput{
		// HITId: *string, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}
	if len(_mturkAssignmentStatuses) > 0 {
		if err := assignInputField(input, "AssignmentStatuses", _mturkAssignmentStatuses); err != nil {
			log.Errorf("invalid --assignment-statuses: %s", err.Error())
			return
		}
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssignmentsForHIT(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListAssignmentsForHITOutput
	p := mturk.NewListAssignmentsForHITPaginator(client, input)
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

// The ListBonusPayments operation retrieves the amounts of bonuses you have paid
// to Workers for a given HIT or assignment.
func mturk_ListBonusPayments(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListBonusPaymentsInput{}

	if len(_mturkAssignmentId) > 0 {
		input.AssignmentId = aws.String(_mturkAssignmentId)
	}
	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBonusPayments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListBonusPaymentsOutput
	p := mturk.NewListBonusPaymentsPaginator(client, input)
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

// The ListHITs operation returns all of a Requester's HITs. The operation
// returns HITs of any status, except for HITs that have been deleted of with the
// DeleteHIT operation or that have been auto-deleted.
func mturk_ListHITs(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListHITsInput{}

	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHITs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListHITsOutput
	p := mturk.NewListHITsPaginator(client, input)
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

// The ListHITsForQualificationType operation returns the HITs that use the given
// Qualification type for a Qualification requirement. The operation returns HITs
// of any status, except for HITs that have been deleted with the DeleteHIT
// operation or that have been auto-deleted.
func mturk_ListHITsForQualificationType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListHITsForQualificationTypeInput{
		// QualificationTypeId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHITsForQualificationType(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListHITsForQualificationTypeOutput
	p := mturk.NewListHITsForQualificationTypePaginator(client, input)
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

// The ListQualificationRequests operation retrieves requests for Qualifications
// of a particular Qualification type. The owner of the Qualification type calls
// this operation to poll for pending requests, and accepts them using the
// AcceptQualification operation.
func mturk_ListQualificationRequests(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListQualificationRequestsInput{}

	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}
	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}

	if disablePaginator() {
		if resp, err := client.ListQualificationRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListQualificationRequestsOutput
	p := mturk.NewListQualificationRequestsPaginator(client, input)
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

// The ListQualificationTypes operation returns a list of Qualification types,
// filtered by an optional search term.
func mturk_ListQualificationTypes(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListQualificationTypesInput{
		// MustBeRequestable: *bool, // Required
	}

	if len(_mturkMustBeRequestable) > 0 {
		if err := assignInputField(input, "MustBeRequestable", _mturkMustBeRequestable); err != nil {
			log.Errorf("invalid --must-be-requestable: %s", err.Error())
			return
		}
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkMustBeOwnedByCaller) > 0 {
		if err := assignInputField(input, "MustBeOwnedByCaller", _mturkMustBeOwnedByCaller); err != nil {
			log.Errorf("invalid --must-be-owned-by-caller: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}
	if len(_mturkQuery) > 0 {
		input.Query = aws.String(_mturkQuery)
	}

	if disablePaginator() {
		if resp, err := client.ListQualificationTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListQualificationTypesOutput
	p := mturk.NewListQualificationTypesPaginator(client, input)
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

// The ListReviewPolicyResultsForHIT operation retrieves the computed results and
// the actions taken in the course of executing your Review Policies for a given
// HIT. For information about how to specify Review Policies when you call
// CreateHIT, see Review Policies. The ListReviewPolicyResultsForHIT operation can
// return results for both Assignment-level and HIT-level review results.
func mturk_ListReviewPolicyResultsForHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListReviewPolicyResultsForHITInput{
		// HITId: *string, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}
	if len(_mturkPolicyLevels) > 0 {
		if err := assignInputField(input, "PolicyLevels", _mturkPolicyLevels); err != nil {
			log.Errorf("invalid --policy-levels: %s", err.Error())
			return
		}
	}
	if len(_mturkRetrieveActions) > 0 {
		if err := assignInputField(input, "RetrieveActions", _mturkRetrieveActions); err != nil {
			log.Errorf("invalid --retrieve-actions: %s", err.Error())
			return
		}
	}
	if len(_mturkRetrieveResults) > 0 {
		if err := assignInputField(input, "RetrieveResults", _mturkRetrieveResults); err != nil {
			log.Errorf("invalid --retrieve-results: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListReviewPolicyResultsForHIT(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListReviewPolicyResultsForHITOutput
	p := mturk.NewListReviewPolicyResultsForHITPaginator(client, input)
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

// The ListReviewableHITs operation retrieves the HITs with Status equal to
// Reviewable or Status equal to Reviewing that belong to the Requester calling the
// operation.
func mturk_ListReviewableHITs(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListReviewableHITsInput{}

	if len(_mturkHITTypeId) > 0 {
		input.HITTypeId = aws.String(_mturkHITTypeId)
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}
	if len(_mturkStatus) > 0 {
		if err := assignInputField(input, "Status", _mturkStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListReviewableHITs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListReviewableHITsOutput
	p := mturk.NewListReviewableHITsPaginator(client, input)
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

// The ListWorkersBlocks operation retrieves a list of Workers who are blocked
// from working on your HITs.
func mturk_ListWorkerBlocks(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListWorkerBlocksInput{}

	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkerBlocks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListWorkerBlocksOutput
	p := mturk.NewListWorkerBlocksPaginator(client, input)
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

// The ListWorkersWithQualificationType operation returns all of the Workers that
// have been associated with a given Qualification type.
func mturk_ListWorkersWithQualificationType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.ListWorkersWithQualificationTypeInput{
		// QualificationTypeId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}
	if len(_mturkMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mturkMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mturkNextToken) > 0 {
		input.NextToken = aws.String(_mturkNextToken)
	}
	if len(_mturkStatus) > 0 {
		if err := assignInputField(input, "Status", _mturkStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListWorkersWithQualificationType(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mturk.ListWorkersWithQualificationTypeOutput
	p := mturk.NewListWorkersWithQualificationTypePaginator(client, input)
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

// The NotifyWorkers operation sends an email to one or more Workers that you
// specify with the Worker ID. You can specify up to 100 Worker IDs to send the
// same message with a single call to the NotifyWorkers operation. The
// NotifyWorkers operation will send a notification email to a Worker only if you
// have previously approved or rejected work from the Worker.
func mturk_NotifyWorkers(cfg aws.Config, client *mturk.Client) {
	input := &mturk.NotifyWorkersInput{
		// MessageText: *string, // Required
		// Subject: *string, // Required
		// WorkerIds: []string, // Required
	}

	if len(_mturkMessageText) > 0 {
		input.MessageText = aws.String(_mturkMessageText)
	}
	if len(_mturkSubject) > 0 {
		input.Subject = aws.String(_mturkSubject)
	}
	if len(_mturkWorkerIds) > 0 {
		input.WorkerIds = append([]string(nil), _mturkWorkerIds...)
	}

	if resp, err := client.NotifyWorkers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The RejectAssignment operation rejects the results of a completed assignment.
// You can include an optional feedback message with the rejection, which the
// Worker can see in the Status section of the web site. When you include a
// feedback message with the rejection, it helps the Worker understand why the
// assignment was rejected, and can improve the quality of the results the Worker
// submits in the future.
//
// Only the Requester who created the HIT can reject an assignment for the HIT.
func mturk_RejectAssignment(cfg aws.Config, client *mturk.Client) {
	input := &mturk.RejectAssignmentInput{
		// AssignmentId: *string, // Required
		// RequesterFeedback: *string, // Required
	}

	if len(_mturkAssignmentId) > 0 {
		input.AssignmentId = aws.String(_mturkAssignmentId)
	}
	if len(_mturkRequesterFeedback) > 0 {
		input.RequesterFeedback = aws.String(_mturkRequesterFeedback)
	}

	if resp, err := client.RejectAssignment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The RejectQualificationRequest operation rejects a user's request for a
// Qualification.
//
// You can provide a text message explaining why the request was rejected. The
// Worker who made the request can see this message.
func mturk_RejectQualificationRequest(cfg aws.Config, client *mturk.Client) {
	input := &mturk.RejectQualificationRequestInput{
		// QualificationRequestId: *string, // Required
	}

	if len(_mturkQualificationRequestId) > 0 {
		input.QualificationRequestId = aws.String(_mturkQualificationRequestId)
	}
	if len(_mturkReason) > 0 {
		input.Reason = aws.String(_mturkReason)
	}

	if resp, err := client.RejectQualificationRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The SendBonus operation issues a payment of money from your account to a
// Worker. This payment happens separately from the reward you pay to the Worker
// when you approve the Worker's assignment. The SendBonus operation requires the
// Worker's ID and the assignment ID as parameters to initiate payment of the
// bonus. You must include a message that explains the reason for the bonus
// payment, as the Worker may not be expecting the payment. Amazon Mechanical Turk
// collects a fee for bonus payments, similar to the HIT listing fee. This
// operation fails if your account does not have enough funds to pay for both the
// bonus and the fees.
func mturk_SendBonus(cfg aws.Config, client *mturk.Client) {
	input := &mturk.SendBonusInput{
		// AssignmentId: *string, // Required
		// BonusAmount: *string, // Required
		// Reason: *string, // Required
		// WorkerId: *string, // Required
	}

	if len(_mturkAssignmentId) > 0 {
		input.AssignmentId = aws.String(_mturkAssignmentId)
	}
	if len(_mturkBonusAmount) > 0 {
		input.BonusAmount = aws.String(_mturkBonusAmount)
	}
	if len(_mturkReason) > 0 {
		input.Reason = aws.String(_mturkReason)
	}
	if len(_mturkWorkerId) > 0 {
		input.WorkerId = aws.String(_mturkWorkerId)
	}
	if len(_mturkUniqueRequestToken) > 0 {
		input.UniqueRequestToken = aws.String(_mturkUniqueRequestToken)
	}

	if resp, err := client.SendBonus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The SendTestEventNotification operation causes Amazon Mechanical Turk to send
// a notification message as if a HIT event occurred, according to the provided
// notification specification. This allows you to test notifications without
// setting up notifications for a real HIT type and trying to trigger them using
// the website. When you call this operation, the service attempts to send the test
// notification immediately.
func mturk_SendTestEventNotification(cfg aws.Config, client *mturk.Client) {
	input := &mturk.SendTestEventNotificationInput{
		// Notification: *types.NotificationSpecification, // Required
		// TestEventType: types.EventType, // Required
	}

	if len(_mturkNotification) > 0 {
		if err := assignInputField(input, "Notification", _mturkNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_mturkTestEventType) > 0 {
		if err := assignInputField(input, "TestEventType", _mturkTestEventType); err != nil {
			log.Errorf("invalid --test-event-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendTestEventNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateExpirationForHIT operation allows you update the expiration time of
// a HIT. If you update it to a time in the past, the HIT will be immediately
// expired.
func mturk_UpdateExpirationForHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.UpdateExpirationForHITInput{
		// ExpireAt: *time.Time, // Required
		// HITId: *string, // Required
	}

	if len(_mturkExpireAt) > 0 {
		if err := assignInputField(input, "ExpireAt", _mturkExpireAt); err != nil {
			log.Errorf("invalid --expire-at: %s", err.Error())
			return
		}
	}
	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}

	if resp, err := client.UpdateExpirationForHIT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateHITReviewStatus operation updates the status of a HIT. If the status
// is Reviewable, this operation can update the status to Reviewing, or it can
// revert a Reviewing HIT back to the Reviewable status.
func mturk_UpdateHITReviewStatus(cfg aws.Config, client *mturk.Client) {
	input := &mturk.UpdateHITReviewStatusInput{
		// HITId: *string, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}
	if len(_mturkRevert) > 0 {
		if err := assignInputField(input, "Revert", _mturkRevert); err != nil {
			log.Errorf("invalid --revert: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateHITReviewStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateHITTypeOfHIT operation allows you to change the HITType properties
// of a HIT. This operation disassociates the HIT from its old HITType properties
// and associates it with the new HITType properties. The HIT takes on the
// properties of the new HITType in place of the old ones.
func mturk_UpdateHITTypeOfHIT(cfg aws.Config, client *mturk.Client) {
	input := &mturk.UpdateHITTypeOfHITInput{
		// HITId: *string, // Required
		// HITTypeId: *string, // Required
	}

	if len(_mturkHITId) > 0 {
		input.HITId = aws.String(_mturkHITId)
	}
	if len(_mturkHITTypeId) > 0 {
		input.HITTypeId = aws.String(_mturkHITTypeId)
	}

	if resp, err := client.UpdateHITTypeOfHIT(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateNotificationSettings operation creates, updates, disables or
// re-enables notifications for a HIT type. If you call the
// UpdateNotificationSettings operation for a HIT type that already has a
// notification specification, the operation replaces the old specification with a
// new one. You can call the UpdateNotificationSettings operation to enable or
// disable notifications for the HIT type, without having to modify the
// notification specification itself by providing updates to the Active status
// without specifying a new notification specification. To change the Active status
// of a HIT type's notifications, the HIT type must already have a notification
// specification, or one must be provided in the same call to
// UpdateNotificationSettings .
func mturk_UpdateNotificationSettings(cfg aws.Config, client *mturk.Client) {
	input := &mturk.UpdateNotificationSettingsInput{
		// HITTypeId: *string, // Required
	}

	if len(_mturkHITTypeId) > 0 {
		input.HITTypeId = aws.String(_mturkHITTypeId)
	}
	if len(_mturkActive) > 0 {
		if err := assignInputField(input, "Active", _mturkActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_mturkNotification) > 0 {
		if err := assignInputField(input, "Notification", _mturkNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNotificationSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateQualificationType operation modifies the attributes of an existing
// Qualification type, which is represented by a QualificationType data structure.
// Only the owner of a Qualification type can modify its attributes.
//
// Most attributes of a Qualification type can be changed after the type has been
// created. However, the Name and Keywords fields cannot be modified. The
// RetryDelayInSeconds parameter can be modified or added to change the delay or to
// enable retries, but RetryDelayInSeconds cannot be used to disable retries.
//
// You can use this operation to update the test for a Qualification type. The
// test is updated based on the values specified for the Test,
// TestDurationInSeconds and AnswerKey parameters. All three parameters specify the
// updated test. If you are updating the test for a type, you must specify the Test
// and TestDurationInSeconds parameters. The AnswerKey parameter is optional;
// omitting it specifies that the updated test does not have an answer key.
//
// If you omit the Test parameter, the test for the Qualification type is
// unchanged. There is no way to remove a test from a Qualification type that has
// one. If the type already has a test, you cannot update it to be AutoGranted. If
// the Qualification type does not have a test and one is provided by an update,
// the type will henceforth have a test.
//
// If you want to update the test duration or answer key for an existing test
// without changing the questions, you must specify a Test parameter with the
// original questions, along with the updated values.
//
// If you provide an updated Test but no AnswerKey, the new test will not have an
// answer key. Requests for such Qualifications must be granted manually.
//
// You can also update the AutoGranted and AutoGrantedValue attributes of the
// Qualification type.
func mturk_UpdateQualificationType(cfg aws.Config, client *mturk.Client) {
	input := &mturk.UpdateQualificationTypeInput{
		// QualificationTypeId: *string, // Required
	}

	if len(_mturkQualificationTypeId) > 0 {
		input.QualificationTypeId = aws.String(_mturkQualificationTypeId)
	}
	if len(_mturkAnswerKey) > 0 {
		input.AnswerKey = aws.String(_mturkAnswerKey)
	}
	if len(_mturkAutoGranted) > 0 {
		if err := assignInputField(input, "AutoGranted", _mturkAutoGranted); err != nil {
			log.Errorf("invalid --auto-granted: %s", err.Error())
			return
		}
	}
	if len(_mturkAutoGrantedValue) > 0 {
		if err := assignInputField(input, "AutoGrantedValue", _mturkAutoGrantedValue); err != nil {
			log.Errorf("invalid --auto-granted-value: %s", err.Error())
			return
		}
	}
	if len(_mturkDescription) > 0 {
		input.Description = aws.String(_mturkDescription)
	}
	if len(_mturkQualificationTypeStatus) > 0 {
		if err := assignInputField(input, "QualificationTypeStatus", _mturkQualificationTypeStatus); err != nil {
			log.Errorf("invalid --qualification-type-status: %s", err.Error())
			return
		}
	}
	if len(_mturkRetryDelayInSeconds) > 0 {
		if err := assignInputField(input, "RetryDelayInSeconds", _mturkRetryDelayInSeconds); err != nil {
			log.Errorf("invalid --retry-delay-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_mturkTest) > 0 {
		input.Test = aws.String(_mturkTest)
	}
	if len(_mturkTestDurationInSeconds) > 0 {
		if err := assignInputField(input, "TestDurationInSeconds", _mturkTestDurationInSeconds); err != nil {
			log.Errorf("invalid --test-duration-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateQualificationType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mturkCmd)
	_mturkCmd.Flags().SortFlags = false

	_mturkCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_mturkCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mturkCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_mturkCmd.Flags().StringVarP(&_mturkActive, "active", "", "", "Active")
	_mturkCmd.Flags().StringVarP(&_mturkAnswerKey, "answer-key", "", "", "Answer Key")
	_mturkCmd.Flags().StringVarP(&_mturkAssignmentDurationInSeconds, "assignment-duration-in-seconds", "", "", "Assignment Duration In Seconds")
	_mturkCmd.Flags().StringVarP(&_mturkAssignmentId, "assignment-id", "", "", "Assignment ID")
	_mturkCmd.Flags().StringVarP(&_mturkAssignmentReviewPolicy, "assignment-review-policy", "", "", "Assignment Review Policy")
	_mturkCmd.Flags().StringVarP(&_mturkAssignmentStatuses, "assignment-statuses", "", "", "Assignment Statuses")
	_mturkCmd.Flags().StringVarP(&_mturkAutoApprovalDelayInSeconds, "auto-approval-delay-in-seconds", "", "", "Auto Approval Delay In Seconds")
	_mturkCmd.Flags().StringVarP(&_mturkAutoGranted, "auto-granted", "", "", "Auto Granted")
	_mturkCmd.Flags().StringVarP(&_mturkAutoGrantedValue, "auto-granted-value", "", "", "Auto Granted Value")
	_mturkCmd.Flags().StringVarP(&_mturkBonusAmount, "bonus-amount", "", "", "Bonus Amount")
	_mturkCmd.Flags().StringVarP(&_mturkDescription, "description", "", "", "Description")
	_mturkCmd.Flags().StringVarP(&_mturkExpireAt, "expire-at", "", "", "Expire At")
	_mturkCmd.Flags().StringVarP(&_mturkHITLayoutId, "hit-layout-id", "", "", "Hit Layout ID")
	_mturkCmd.Flags().StringVarP(&_mturkHITLayoutParameters, "hit-layout-parameters", "", "", "Hit Layout Parameters")
	_mturkCmd.Flags().StringVarP(&_mturkHITReviewPolicy, "hit-review-policy", "", "", "Hit Review Policy")
	_mturkCmd.Flags().StringVarP(&_mturkHITTypeId, "hit-type-id", "", "", "Hit Type ID")
	_mturkCmd.Flags().StringVarP(&_mturkHITId, "hitid", "", "", "Hitid")
	_mturkCmd.Flags().StringVarP(&_mturkIntegerValue, "integer-value", "", "", "Integer Value")
	_mturkCmd.Flags().StringVarP(&_mturkKeywords, "keywords", "", "", "Keywords")
	_mturkCmd.Flags().StringVarP(&_mturkLifetimeInSeconds, "lifetime-in-seconds", "", "", "Lifetime In Seconds")
	_mturkCmd.Flags().StringVarP(&_mturkMaxAssignments, "max-assignments", "", "", "Max Assignments")
	_mturkCmd.Flags().StringVarP(&_mturkMaxResults, "max-results", "", "", "Max Results")
	_mturkCmd.Flags().StringVarP(&_mturkMessageText, "message-text", "", "", "Message Text")
	_mturkCmd.Flags().StringVarP(&_mturkMustBeOwnedByCaller, "must-be-owned-by-caller", "", "", "Must Be Owned By Caller")
	_mturkCmd.Flags().StringVarP(&_mturkMustBeRequestable, "must-be-requestable", "", "", "Must Be Requestable")
	_mturkCmd.Flags().StringVarP(&_mturkName, "name", "", "", "Name")
	_mturkCmd.Flags().StringVarP(&_mturkNextToken, "next-token", "", "", "Next Token")
	_mturkCmd.Flags().StringVarP(&_mturkNotification, "notification", "", "", "Notification")
	_mturkCmd.Flags().StringVarP(&_mturkNumberOfAdditionalAssignments, "number-of-additional-assignments", "", "", "Number Of Additional Assignments")
	_mturkCmd.Flags().StringVarP(&_mturkOverrideRejection, "override-rejection", "", "", "Override Rejection")
	_mturkCmd.Flags().StringVarP(&_mturkPolicyLevels, "policy-levels", "", "", "Policy Levels")
	_mturkCmd.Flags().StringVarP(&_mturkQualificationRequestId, "qualification-request-id", "", "", "Qualification Request ID")
	_mturkCmd.Flags().StringVarP(&_mturkQualificationRequirements, "qualification-requirements", "", "", "Qualification Requirements")
	_mturkCmd.Flags().StringVarP(&_mturkQualificationTypeId, "qualification-type-id", "", "", "Qualification Type ID")
	_mturkCmd.Flags().StringVarP(&_mturkQualificationTypeStatus, "qualification-type-status", "", "", "Qualification Type Status")
	_mturkCmd.Flags().StringVarP(&_mturkQuery, "query", "", "", "Query")
	_mturkCmd.Flags().StringVarP(&_mturkQuestion, "question", "", "", "Question")
	_mturkCmd.Flags().StringVarP(&_mturkQuestionIdentifier, "question-identifier", "", "", "Question Identifier")
	_mturkCmd.Flags().StringVarP(&_mturkReason, "reason", "", "", "Reason")
	_mturkCmd.Flags().StringVarP(&_mturkRequesterAnnotation, "requester-annotation", "", "", "Requester Annotation")
	_mturkCmd.Flags().StringVarP(&_mturkRequesterFeedback, "requester-feedback", "", "", "Requester Feedback")
	_mturkCmd.Flags().StringVarP(&_mturkRetrieveActions, "retrieve-actions", "", "", "Retrieve Actions")
	_mturkCmd.Flags().StringVarP(&_mturkRetrieveResults, "retrieve-results", "", "", "Retrieve Results")
	_mturkCmd.Flags().StringVarP(&_mturkRetryDelayInSeconds, "retry-delay-in-seconds", "", "", "Retry Delay In Seconds")
	_mturkCmd.Flags().StringVarP(&_mturkRevert, "revert", "", "", "Revert")
	_mturkCmd.Flags().StringVarP(&_mturkReward, "reward", "", "", "Reward")
	_mturkCmd.Flags().StringVarP(&_mturkSendNotification, "send-notification", "", "", "Send Notification")
	_mturkCmd.Flags().StringVarP(&_mturkStatus, "status", "", "", "Status")
	_mturkCmd.Flags().StringVarP(&_mturkSubject, "subject", "", "", "Subject")
	_mturkCmd.Flags().StringVarP(&_mturkTest, "test", "", "", "Test")
	_mturkCmd.Flags().StringVarP(&_mturkTestDurationInSeconds, "test-duration-in-seconds", "", "", "Test Duration In Seconds")
	_mturkCmd.Flags().StringVarP(&_mturkTestEventType, "test-event-type", "", "", "Test Event Type")
	_mturkCmd.Flags().StringVarP(&_mturkTitle, "title", "", "", "Title")
	_mturkCmd.Flags().StringVarP(&_mturkUniqueRequestToken, "unique-request-token", "", "", "Unique Request Token")
	_mturkCmd.Flags().StringVarP(&_mturkWorkerId, "worker-id", "", "", "Worker ID")
	_mturkCmd.Flags().StringSliceVarP(&_mturkWorkerIds, "worker-ids", "", nil, "Worker Ids")

	_mturkCmd.Flags().BoolVarP(&_mturkAcceptQualificationRequest, "accept-qualification-request", "", false, "Accept Qualification Request")
	_mturkCmd.Flags().BoolVarP(&_mturkApproveAssignment, "approve-assignment", "", false, "Approve Assignment")
	_mturkCmd.Flags().BoolVarP(&_mturkAssociateQualificationWithWorker, "associate-qualification-with-worker", "", false, "Associate Qualification With Worker")
	_mturkCmd.Flags().BoolVarP(&_mturkCreateAdditionalAssignmentsForHIT, "create-additional-assignments-for-hit", "", false, "Create Additional Assignments For Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkCreateHIT, "create-hit", "", false, "Create Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkCreateHITType, "create-hit-type", "", false, "Create Hit Type")
	_mturkCmd.Flags().BoolVarP(&_mturkCreateHITWithHITType, "create-hit-with-hit-type", "", false, "Create Hit With Hit Type")
	_mturkCmd.Flags().BoolVarP(&_mturkCreateQualificationType, "create-qualification-type", "", false, "Create Qualification Type")
	_mturkCmd.Flags().BoolVarP(&_mturkCreateWorkerBlock, "create-worker-block", "", false, "Create Worker Block")
	_mturkCmd.Flags().BoolVarP(&_mturkDeleteHIT, "delete-hit", "", false, "Delete Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkDeleteQualificationType, "delete-qualification-type", "", false, "Delete Qualification Type")
	_mturkCmd.Flags().BoolVarP(&_mturkDeleteWorkerBlock, "delete-worker-block", "", false, "Delete Worker Block")
	_mturkCmd.Flags().BoolVarP(&_mturkDisassociateQualificationFromWorker, "disassociate-qualification-from-worker", "", false, "Disassociate Qualification From Worker")
	_mturkCmd.Flags().BoolVarP(&_mturkGetAccountBalance, "get-account-balance", "", false, "Get Account Balance")
	_mturkCmd.Flags().BoolVarP(&_mturkGetAssignment, "get-assignment", "", false, "Get Assignment")
	_mturkCmd.Flags().BoolVarP(&_mturkGetFileUploadURL, "get-file-upload-url", "", false, "Get File Upload URL")
	_mturkCmd.Flags().BoolVarP(&_mturkGetHIT, "get-hit", "", false, "Get Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkGetQualificationScore, "get-qualification-score", "", false, "Get Qualification Score")
	_mturkCmd.Flags().BoolVarP(&_mturkGetQualificationType, "get-qualification-type", "", false, "Get Qualification Type")
	_mturkCmd.Flags().BoolVarP(&_mturkListAssignmentsForHIT, "list-assignments-for-hit", "", false, "List Assignments For Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkListBonusPayments, "list-bonus-payments", "", false, "List Bonus Payments")
	_mturkCmd.Flags().BoolVarP(&_mturkListHITs, "list-hits", "", false, "List Hits")
	_mturkCmd.Flags().BoolVarP(&_mturkListHITsForQualificationType, "list-hits-for-qualification-type", "", false, "List Hits For Qualification Type")
	_mturkCmd.Flags().BoolVarP(&_mturkListQualificationRequests, "list-qualification-requests", "", false, "List Qualification Requests")
	_mturkCmd.Flags().BoolVarP(&_mturkListQualificationTypes, "list-qualification-types", "", false, "List Qualification Types")
	_mturkCmd.Flags().BoolVarP(&_mturkListReviewPolicyResultsForHIT, "list-review-policy-results-for-hit", "", false, "List Review Policy Results For Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkListReviewableHITs, "list-reviewable-hits", "", false, "List Reviewable Hits")
	_mturkCmd.Flags().BoolVarP(&_mturkListWorkerBlocks, "list-worker-blocks", "", false, "List Worker Blocks")
	_mturkCmd.Flags().BoolVarP(&_mturkListWorkersWithQualificationType, "list-workers-with-qualification-type", "", false, "List Workers With Qualification Type")
	_mturkCmd.Flags().BoolVarP(&_mturkNotifyWorkers, "notify-workers", "", false, "Notify Workers")
	_mturkCmd.Flags().BoolVarP(&_mturkRejectAssignment, "reject-assignment", "", false, "Reject Assignment")
	_mturkCmd.Flags().BoolVarP(&_mturkRejectQualificationRequest, "reject-qualification-request", "", false, "Reject Qualification Request")
	_mturkCmd.Flags().BoolVarP(&_mturkSendBonus, "send-bonus", "", false, "Send Bonus")
	_mturkCmd.Flags().BoolVarP(&_mturkSendTestEventNotification, "send-test-event-notification", "", false, "Send Test Event Notification")
	_mturkCmd.Flags().BoolVarP(&_mturkUpdateExpirationForHIT, "update-expiration-for-hit", "", false, "Update Expiration For Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkUpdateHITReviewStatus, "update-hit-review-status", "", false, "Update Hit Review Status")
	_mturkCmd.Flags().BoolVarP(&_mturkUpdateHITTypeOfHIT, "update-hit-type-of-hit", "", false, "Update Hit Type Of Hit")
	_mturkCmd.Flags().BoolVarP(&_mturkUpdateNotificationSettings, "update-notification-settings", "", false, "Update Notification Settings")
	_mturkCmd.Flags().BoolVarP(&_mturkUpdateQualificationType, "update-qualification-type", "", false, "Update Qualification Type")

}
