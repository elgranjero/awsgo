package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// budgetsCmd represents the budgets command
var _budgetsCmd = &cobra.Command{
	Use:   "budgets",
	Short: "AWS budgets CLI",
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
		client := budgets.NewFromConfig(cfg)
		if _budgetsCreateBudget {
			budgets_CreateBudget(cfg, client)
			return
		}
		if _budgetsCreateBudgetAction {
			budgets_CreateBudgetAction(cfg, client)
			return
		}
		if _budgetsCreateNotification {
			budgets_CreateNotification(cfg, client)
			return
		}
		if _budgetsCreateSubscriber {
			budgets_CreateSubscriber(cfg, client)
			return
		}
		if _budgetsDeleteBudget {
			budgets_DeleteBudget(cfg, client)
			return
		}
		if _budgetsDeleteBudgetAction {
			budgets_DeleteBudgetAction(cfg, client)
			return
		}
		if _budgetsDeleteNotification {
			budgets_DeleteNotification(cfg, client)
			return
		}
		if _budgetsDeleteSubscriber {
			budgets_DeleteSubscriber(cfg, client)
			return
		}
		if _budgetsDescribeBudget {
			budgets_DescribeBudget(cfg, client)
			return
		}
		if _budgetsDescribeBudgetAction {
			budgets_DescribeBudgetAction(cfg, client)
			return
		}
		if _budgetsDescribeBudgetActionHistories {
			budgets_DescribeBudgetActionHistories(cfg, client)
			return
		}
		if _budgetsDescribeBudgetActionsForAccount {
			budgets_DescribeBudgetActionsForAccount(cfg, client)
			return
		}
		if _budgetsDescribeBudgetActionsForBudget {
			budgets_DescribeBudgetActionsForBudget(cfg, client)
			return
		}
		if _budgetsDescribeBudgetNotificationsForAccount {
			budgets_DescribeBudgetNotificationsForAccount(cfg, client)
			return
		}
		if _budgetsDescribeBudgetPerformanceHistory {
			budgets_DescribeBudgetPerformanceHistory(cfg, client)
			return
		}
		if _budgetsDescribeBudgets {
			budgets_DescribeBudgets(cfg, client)
			return
		}
		if _budgetsDescribeNotificationsForBudget {
			budgets_DescribeNotificationsForBudget(cfg, client)
			return
		}
		if _budgetsDescribeSubscribersForNotification {
			budgets_DescribeSubscribersForNotification(cfg, client)
			return
		}
		if _budgetsExecuteBudgetAction {
			budgets_ExecuteBudgetAction(cfg, client)
			return
		}
		if _budgetsListTagsForResource {
			budgets_ListTagsForResource(cfg, client)
			return
		}
		if _budgetsTagResource {
			budgets_TagResource(cfg, client)
			return
		}
		if _budgetsUntagResource {
			budgets_UntagResource(cfg, client)
			return
		}
		if _budgetsUpdateBudget {
			budgets_UpdateBudget(cfg, client)
			return
		}
		if _budgetsUpdateBudgetAction {
			budgets_UpdateBudgetAction(cfg, client)
			return
		}
		if _budgetsUpdateNotification {
			budgets_UpdateNotification(cfg, client)
			return
		}
		if _budgetsUpdateSubscriber {
			budgets_UpdateSubscriber(cfg, client)
			return
		}

	},
}

var (
	_budgetsCreateBudget                          bool
	_budgetsCreateBudgetAction                    bool
	_budgetsCreateNotification                    bool
	_budgetsCreateSubscriber                      bool
	_budgetsDeleteBudget                          bool
	_budgetsDeleteBudgetAction                    bool
	_budgetsDeleteNotification                    bool
	_budgetsDeleteSubscriber                      bool
	_budgetsDescribeBudget                        bool
	_budgetsDescribeBudgetAction                  bool
	_budgetsDescribeBudgetActionHistories         bool
	_budgetsDescribeBudgetActionsForAccount       bool
	_budgetsDescribeBudgetActionsForBudget        bool
	_budgetsDescribeBudgetNotificationsForAccount bool
	_budgetsDescribeBudgetPerformanceHistory      bool
	_budgetsDescribeBudgets                       bool
	_budgetsDescribeNotificationsForBudget        bool
	_budgetsDescribeSubscribersForNotification    bool
	_budgetsExecuteBudgetAction                   bool
	_budgetsListTagsForResource                   bool
	_budgetsTagResource                           bool
	_budgetsUntagResource                         bool
	_budgetsUpdateBudget                          bool
	_budgetsUpdateBudgetAction                    bool
	_budgetsUpdateNotification                    bool
	_budgetsUpdateSubscriber                      bool

	_budgetsAccountId                    string
	_budgetsActionId                     string
	_budgetsActionThreshold              string
	_budgetsActionType                   string
	_budgetsApprovalModel                string
	_budgetsBudget                       string
	_budgetsBudgetName                   string
	_budgetsDefinition                   string
	_budgetsExecutionRoleArn             string
	_budgetsExecutionType                string
	_budgetsMaxResults                   string
	_budgetsNewBudget                    string
	_budgetsNewNotification              string
	_budgetsNewSubscriber                string
	_budgetsNextToken                    string
	_budgetsNotification                 string
	_budgetsNotificationType             string
	_budgetsNotificationsWithSubscribers string
	_budgetsOldNotification              string
	_budgetsOldSubscriber                string
	_budgetsResourceARN                  string
	_budgetsResourceTagKeys              []string
	_budgetsResourceTags                 string
	_budgetsShowFilterExpression         string
	_budgetsSubscriber                   string
	_budgetsSubscribers                  string
	_budgetsTimePeriod                   string
)

// Creates a budget and, if included, notifications and subscribers.
// Only one of BudgetLimit or PlannedBudgetLimits can be present in the syntax at
// one time. Use the syntax that matches your use case. The Request Syntax section
// shows the BudgetLimit syntax. For PlannedBudgetLimits , see the [Examples] section.
//
// Similarly, only one set of filter and metric selections can be present in the
// syntax at one time. Either FilterExpression and Metrics or CostFilters and
// CostTypes , not both or a different combination. We recommend using
// FilterExpression and Metrics as they provide more flexible and powerful
// filtering capabilities. The Request Syntax section shows the FilterExpression /
// Metrics syntax.
//
// [Examples]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_CreateBudget.html#API_CreateBudget_Examples
func budgets_CreateBudget(cfg aws.Config, client *budgets.Client) {
	input := &budgets.CreateBudgetInput{
		// AccountId: *string, // Required
		// Budget: *types.Budget, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudget) > 0 {
		if err := assignInputField(input, "Budget", _budgetsBudget); err != nil {
			log.Errorf("invalid --budget: %s", err.Error())
			return
		}
	}
	if len(_budgetsNotificationsWithSubscribers) > 0 {
		if err := assignInputField(input, "NotificationsWithSubscribers", _budgetsNotificationsWithSubscribers); err != nil {
			log.Errorf("invalid --notifications-with-subscribers: %s", err.Error())
			return
		}
	}
	if len(_budgetsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _budgetsResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a budget action.
func budgets_CreateBudgetAction(cfg aws.Config, client *budgets.Client) {
	input := &budgets.CreateBudgetActionInput{
		// AccountId: *string, // Required
		// ActionThreshold: *types.ActionThreshold, // Required
		// ActionType: types.ActionType, // Required
		// ApprovalModel: types.ApprovalModel, // Required
		// BudgetName: *string, // Required
		// Definition: *types.Definition, // Required
		// ExecutionRoleArn: *string, // Required
		// NotificationType: types.NotificationType, // Required
		// Subscribers: []types.Subscriber, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsActionThreshold) > 0 {
		if err := assignInputField(input, "ActionThreshold", _budgetsActionThreshold); err != nil {
			log.Errorf("invalid --action-threshold: %s", err.Error())
			return
		}
	}
	if len(_budgetsActionType) > 0 {
		if err := assignInputField(input, "ActionType", _budgetsActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_budgetsApprovalModel) > 0 {
		if err := assignInputField(input, "ApprovalModel", _budgetsApprovalModel); err != nil {
			log.Errorf("invalid --approval-model: %s", err.Error())
			return
		}
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsDefinition) > 0 {
		if err := assignInputField(input, "Definition", _budgetsDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_budgetsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_budgetsExecutionRoleArn)
	}
	if len(_budgetsNotificationType) > 0 {
		if err := assignInputField(input, "NotificationType", _budgetsNotificationType); err != nil {
			log.Errorf("invalid --notification-type: %s", err.Error())
			return
		}
	}
	if len(_budgetsSubscribers) > 0 {
		if err := assignInputField(input, "Subscribers", _budgetsSubscribers); err != nil {
			log.Errorf("invalid --subscribers: %s", err.Error())
			return
		}
	}
	if len(_budgetsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _budgetsResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBudgetAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a notification. You must create the budget before you create the
// associated notification.
func budgets_CreateNotification(cfg aws.Config, client *budgets.Client) {
	input := &budgets.CreateNotificationInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// Notification: *types.Notification, // Required
		// Subscribers: []types.Subscriber, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNotification) > 0 {
		if err := assignInputField(input, "Notification", _budgetsNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_budgetsSubscribers) > 0 {
		if err := assignInputField(input, "Subscribers", _budgetsSubscribers); err != nil {
			log.Errorf("invalid --subscribers: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscriber. You must create the associated budget and notification
// before you create the subscriber.
func budgets_CreateSubscriber(cfg aws.Config, client *budgets.Client) {
	input := &budgets.CreateSubscriberInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// Notification: *types.Notification, // Required
		// Subscriber: *types.Subscriber, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNotification) > 0 {
		if err := assignInputField(input, "Notification", _budgetsNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_budgetsSubscriber) > 0 {
		if err := assignInputField(input, "Subscriber", _budgetsSubscriber); err != nil {
			log.Errorf("invalid --subscriber: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a budget. You can delete your budget at any time.
// Deleting a budget also deletes the notifications and subscribers that are
// associated with that budget.
func budgets_DeleteBudget(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DeleteBudgetInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}

	if resp, err := client.DeleteBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a budget action.
func budgets_DeleteBudgetAction(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DeleteBudgetActionInput{
		// AccountId: *string, // Required
		// ActionId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsActionId) > 0 {
		input.ActionId = aws.String(_budgetsActionId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}

	if resp, err := client.DeleteBudgetAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a notification.
// Deleting a notification also deletes the subscribers that are associated with
// the notification.
func budgets_DeleteNotification(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DeleteNotificationInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// Notification: *types.Notification, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNotification) > 0 {
		if err := assignInputField(input, "Notification", _budgetsNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subscriber.
// Deleting the last subscriber to a notification also deletes the notification.
func budgets_DeleteSubscriber(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DeleteSubscriberInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// Notification: *types.Notification, // Required
		// Subscriber: *types.Subscriber, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNotification) > 0 {
		if err := assignInputField(input, "Notification", _budgetsNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_budgetsSubscriber) > 0 {
		if err := assignInputField(input, "Subscriber", _budgetsSubscriber); err != nil {
			log.Errorf("invalid --subscriber: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a budget.
// The Request Syntax section shows the BudgetLimit syntax. For PlannedBudgetLimits
// , see the [Examples]section.
//
// [Examples]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_DescribeBudget.html#API_DescribeBudget_Examples
func budgets_DescribeBudget(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsShowFilterExpression) > 0 {
		if err := assignInputField(input, "ShowFilterExpression", _budgetsShowFilterExpression); err != nil {
			log.Errorf("invalid --show-filter-expression: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a budget action detail.
func budgets_DescribeBudgetAction(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetActionInput{
		// AccountId: *string, // Required
		// ActionId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsActionId) > 0 {
		input.ActionId = aws.String(_budgetsActionId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}

	if resp, err := client.DescribeBudgetAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a budget action history detail.
func budgets_DescribeBudgetActionHistories(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetActionHistoriesInput{
		// AccountId: *string, // Required
		// ActionId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsActionId) > 0 {
		input.ActionId = aws.String(_budgetsActionId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}
	if len(_budgetsTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _budgetsTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBudgetActionHistories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeBudgetActionHistoriesOutput
	p := budgets.NewDescribeBudgetActionHistoriesPaginator(client, input)
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

// Describes all of the budget actions for an account.
func budgets_DescribeBudgetActionsForAccount(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetActionsForAccountInput{
		// AccountId: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeBudgetActionsForAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeBudgetActionsForAccountOutput
	p := budgets.NewDescribeBudgetActionsForAccountPaginator(client, input)
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

// Describes all of the budget actions for a budget.
func budgets_DescribeBudgetActionsForBudget(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetActionsForBudgetInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeBudgetActionsForBudget(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeBudgetActionsForBudgetOutput
	p := budgets.NewDescribeBudgetActionsForBudgetPaginator(client, input)
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

// Lists the budget names and notifications that are associated with an account.
func budgets_DescribeBudgetNotificationsForAccount(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetNotificationsForAccountInput{
		// AccountId: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeBudgetNotificationsForAccount(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeBudgetNotificationsForAccountOutput
	p := budgets.NewDescribeBudgetNotificationsForAccountPaginator(client, input)
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

// Describes the history for DAILY , MONTHLY , and QUARTERLY budgets. Budget
// history isn't available for ANNUAL budgets.
func budgets_DescribeBudgetPerformanceHistory(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetPerformanceHistoryInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}
	if len(_budgetsTimePeriod) > 0 {
		if err := assignInputField(input, "TimePeriod", _budgetsTimePeriod); err != nil {
			log.Errorf("invalid --time-period: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBudgetPerformanceHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeBudgetPerformanceHistoryOutput
	p := budgets.NewDescribeBudgetPerformanceHistoryPaginator(client, input)
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

// Lists the budgets that are associated with an account.
// The Request Syntax section shows the BudgetLimit syntax. For PlannedBudgetLimits
// , see the [Examples]section.
//
// [Examples]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_DescribeBudgets.html#API_DescribeBudgets_Examples
func budgets_DescribeBudgets(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeBudgetsInput{
		// AccountId: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}
	if len(_budgetsShowFilterExpression) > 0 {
		if err := assignInputField(input, "ShowFilterExpression", _budgetsShowFilterExpression); err != nil {
			log.Errorf("invalid --show-filter-expression: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBudgets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeBudgetsOutput
	p := budgets.NewDescribeBudgetsPaginator(client, input)
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

// Lists the notifications that are associated with a budget.
func budgets_DescribeNotificationsForBudget(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeNotificationsForBudgetInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeNotificationsForBudget(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeNotificationsForBudgetOutput
	p := budgets.NewDescribeNotificationsForBudgetPaginator(client, input)
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

// Lists the subscribers that are associated with a notification.
func budgets_DescribeSubscribersForNotification(cfg aws.Config, client *budgets.Client) {
	input := &budgets.DescribeSubscribersForNotificationInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// Notification: *types.Notification, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNotification) > 0 {
		if err := assignInputField(input, "Notification", _budgetsNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_budgetsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _budgetsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_budgetsNextToken) > 0 {
		input.NextToken = aws.String(_budgetsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSubscribersForNotification(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*budgets.DescribeSubscribersForNotificationOutput
	p := budgets.NewDescribeSubscribersForNotificationPaginator(client, input)
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

// Executes a budget action.
func budgets_ExecuteBudgetAction(cfg aws.Config, client *budgets.Client) {
	input := &budgets.ExecuteBudgetActionInput{
		// AccountId: *string, // Required
		// ActionId: *string, // Required
		// BudgetName: *string, // Required
		// ExecutionType: types.ExecutionType, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsActionId) > 0 {
		input.ActionId = aws.String(_budgetsActionId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsExecutionType) > 0 {
		if err := assignInputField(input, "ExecutionType", _budgetsExecutionType); err != nil {
			log.Errorf("invalid --execution-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteBudgetAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tags associated with a budget or budget action resource.
func budgets_ListTagsForResource(cfg aws.Config, client *budgets.Client) {
	input := &budgets.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_budgetsResourceARN) > 0 {
		input.ResourceARN = aws.String(_budgetsResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates tags for a budget or budget action resource.
func budgets_TagResource(cfg aws.Config, client *budgets.Client) {
	input := &budgets.TagResourceInput{
		// ResourceARN: *string, // Required
		// ResourceTags: []types.ResourceTag, // Required
	}

	if len(_budgetsResourceARN) > 0 {
		input.ResourceARN = aws.String(_budgetsResourceARN)
	}
	if len(_budgetsResourceTags) > 0 {
		if err := assignInputField(input, "ResourceTags", _budgetsResourceTags); err != nil {
			log.Errorf("invalid --resource-tags: %s", err.Error())
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

// Deletes tags associated with a budget or budget action resource.
func budgets_UntagResource(cfg aws.Config, client *budgets.Client) {
	input := &budgets.UntagResourceInput{
		// ResourceARN: *string, // Required
		// ResourceTagKeys: []string, // Required
	}

	if len(_budgetsResourceARN) > 0 {
		input.ResourceARN = aws.String(_budgetsResourceARN)
	}
	if len(_budgetsResourceTagKeys) > 0 {
		input.ResourceTagKeys = append([]string(nil), _budgetsResourceTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a budget. You can change every part of a budget except for the
// budgetName and the calculatedSpend . When you modify a budget, the
// calculatedSpend drops to zero until Amazon Web Services has new usage data to
// use for forecasting.
//
// Only one of BudgetLimit or PlannedBudgetLimits can be present in the syntax at
// one time. Use the syntax that matches your case. The Request Syntax section
// shows the BudgetLimit syntax. For PlannedBudgetLimits , see the [Examples] section.
//
// Similarly, only one set of filter and metric selections can be present in the
// syntax at one time. Either FilterExpression and Metrics or CostFilters and
// CostTypes , not both or a different combination. We recommend using
// FilterExpression and Metrics as they provide more flexible and powerful
// filtering capabilities. The Request Syntax section shows the FilterExpression /
// Metrics syntax.
//
// [Examples]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_UpdateBudget.html#API_UpdateBudget_Examples
func budgets_UpdateBudget(cfg aws.Config, client *budgets.Client) {
	input := &budgets.UpdateBudgetInput{
		// AccountId: *string, // Required
		// NewBudget: *types.Budget, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsNewBudget) > 0 {
		if err := assignInputField(input, "NewBudget", _budgetsNewBudget); err != nil {
			log.Errorf("invalid --new-budget: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBudget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a budget action.
func budgets_UpdateBudgetAction(cfg aws.Config, client *budgets.Client) {
	input := &budgets.UpdateBudgetActionInput{
		// AccountId: *string, // Required
		// ActionId: *string, // Required
		// BudgetName: *string, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsActionId) > 0 {
		input.ActionId = aws.String(_budgetsActionId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsActionThreshold) > 0 {
		if err := assignInputField(input, "ActionThreshold", _budgetsActionThreshold); err != nil {
			log.Errorf("invalid --action-threshold: %s", err.Error())
			return
		}
	}
	if len(_budgetsApprovalModel) > 0 {
		if err := assignInputField(input, "ApprovalModel", _budgetsApprovalModel); err != nil {
			log.Errorf("invalid --approval-model: %s", err.Error())
			return
		}
	}
	if len(_budgetsDefinition) > 0 {
		if err := assignInputField(input, "Definition", _budgetsDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_budgetsExecutionRoleArn) > 0 {
		input.ExecutionRoleArn = aws.String(_budgetsExecutionRoleArn)
	}
	if len(_budgetsNotificationType) > 0 {
		if err := assignInputField(input, "NotificationType", _budgetsNotificationType); err != nil {
			log.Errorf("invalid --notification-type: %s", err.Error())
			return
		}
	}
	if len(_budgetsSubscribers) > 0 {
		if err := assignInputField(input, "Subscribers", _budgetsSubscribers); err != nil {
			log.Errorf("invalid --subscribers: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateBudgetAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a notification.
func budgets_UpdateNotification(cfg aws.Config, client *budgets.Client) {
	input := &budgets.UpdateNotificationInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// NewNotification: *types.Notification, // Required
		// OldNotification: *types.Notification, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNewNotification) > 0 {
		if err := assignInputField(input, "NewNotification", _budgetsNewNotification); err != nil {
			log.Errorf("invalid --new-notification: %s", err.Error())
			return
		}
	}
	if len(_budgetsOldNotification) > 0 {
		if err := assignInputField(input, "OldNotification", _budgetsOldNotification); err != nil {
			log.Errorf("invalid --old-notification: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a subscriber.
func budgets_UpdateSubscriber(cfg aws.Config, client *budgets.Client) {
	input := &budgets.UpdateSubscriberInput{
		// AccountId: *string, // Required
		// BudgetName: *string, // Required
		// NewSubscriber: *types.Subscriber, // Required
		// Notification: *types.Notification, // Required
		// OldSubscriber: *types.Subscriber, // Required
	}

	if len(_budgetsAccountId) > 0 {
		input.AccountId = aws.String(_budgetsAccountId)
	}
	if len(_budgetsBudgetName) > 0 {
		input.BudgetName = aws.String(_budgetsBudgetName)
	}
	if len(_budgetsNewSubscriber) > 0 {
		if err := assignInputField(input, "NewSubscriber", _budgetsNewSubscriber); err != nil {
			log.Errorf("invalid --new-subscriber: %s", err.Error())
			return
		}
	}
	if len(_budgetsNotification) > 0 {
		if err := assignInputField(input, "Notification", _budgetsNotification); err != nil {
			log.Errorf("invalid --notification: %s", err.Error())
			return
		}
	}
	if len(_budgetsOldSubscriber) > 0 {
		if err := assignInputField(input, "OldSubscriber", _budgetsOldSubscriber); err != nil {
			log.Errorf("invalid --old-subscriber: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_budgetsCmd)
	_budgetsCmd.Flags().SortFlags = false

	_budgetsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_budgetsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_budgetsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_budgetsCmd.Flags().StringVarP(&_budgetsAccountId, "account-id", "", "", "Account ID")
	_budgetsCmd.Flags().StringVarP(&_budgetsActionId, "action-id", "", "", "Action ID")
	_budgetsCmd.Flags().StringVarP(&_budgetsActionThreshold, "action-threshold", "", "", "Action Threshold")
	_budgetsCmd.Flags().StringVarP(&_budgetsActionType, "action-type", "", "", "Action Type")
	_budgetsCmd.Flags().StringVarP(&_budgetsApprovalModel, "approval-model", "", "", "Approval Model")
	_budgetsCmd.Flags().StringVarP(&_budgetsBudget, "budget", "", "", "Budget")
	_budgetsCmd.Flags().StringVarP(&_budgetsBudgetName, "budget-name", "", "", "Budget Name")
	_budgetsCmd.Flags().StringVarP(&_budgetsDefinition, "definition", "", "", "Definition")
	_budgetsCmd.Flags().StringVarP(&_budgetsExecutionRoleArn, "execution-role-arn", "", "", "Execution Role ARN")
	_budgetsCmd.Flags().StringVarP(&_budgetsExecutionType, "execution-type", "", "", "Execution Type")
	_budgetsCmd.Flags().StringVarP(&_budgetsMaxResults, "max-results", "", "", "Max Results")
	_budgetsCmd.Flags().StringVarP(&_budgetsNewBudget, "new-budget", "", "", "New Budget")
	_budgetsCmd.Flags().StringVarP(&_budgetsNewNotification, "new-notification", "", "", "New Notification")
	_budgetsCmd.Flags().StringVarP(&_budgetsNewSubscriber, "new-subscriber", "", "", "New Subscriber")
	_budgetsCmd.Flags().StringVarP(&_budgetsNextToken, "next-token", "", "", "Next Token")
	_budgetsCmd.Flags().StringVarP(&_budgetsNotification, "notification", "", "", "Notification")
	_budgetsCmd.Flags().StringVarP(&_budgetsNotificationType, "notification-type", "", "", "Notification Type")
	_budgetsCmd.Flags().StringVarP(&_budgetsNotificationsWithSubscribers, "notifications-with-subscribers", "", "", "Notifications With Subscribers")
	_budgetsCmd.Flags().StringVarP(&_budgetsOldNotification, "old-notification", "", "", "Old Notification")
	_budgetsCmd.Flags().StringVarP(&_budgetsOldSubscriber, "old-subscriber", "", "", "Old Subscriber")
	_budgetsCmd.Flags().StringVarP(&_budgetsResourceARN, "resource-arn", "", "", "Resource ARN")
	_budgetsCmd.Flags().StringSliceVarP(&_budgetsResourceTagKeys, "resource-tag-keys", "", nil, "Resource Tag Keys")
	_budgetsCmd.Flags().StringVarP(&_budgetsResourceTags, "resource-tags", "", "", "Resource Tags")
	_budgetsCmd.Flags().StringVarP(&_budgetsShowFilterExpression, "show-filter-expression", "", "", "Show Filter Expression")
	_budgetsCmd.Flags().StringVarP(&_budgetsSubscriber, "subscriber", "", "", "Subscriber")
	_budgetsCmd.Flags().StringVarP(&_budgetsSubscribers, "subscribers", "", "", "Subscribers")
	_budgetsCmd.Flags().StringVarP(&_budgetsTimePeriod, "time-period", "", "", "Time Period")

	_budgetsCmd.Flags().BoolVarP(&_budgetsCreateBudget, "create-budget", "", false, "Create Budget")
	_budgetsCmd.Flags().BoolVarP(&_budgetsCreateBudgetAction, "create-budget-action", "", false, "Create Budget Action")
	_budgetsCmd.Flags().BoolVarP(&_budgetsCreateNotification, "create-notification", "", false, "Create Notification")
	_budgetsCmd.Flags().BoolVarP(&_budgetsCreateSubscriber, "create-subscriber", "", false, "Create Subscriber")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDeleteBudget, "delete-budget", "", false, "Delete Budget")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDeleteBudgetAction, "delete-budget-action", "", false, "Delete Budget Action")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDeleteNotification, "delete-notification", "", false, "Delete Notification")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDeleteSubscriber, "delete-subscriber", "", false, "Delete Subscriber")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudget, "describe-budget", "", false, "Describe Budget")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgetAction, "describe-budget-action", "", false, "Describe Budget Action")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgetActionHistories, "describe-budget-action-histories", "", false, "Describe Budget Action Histories")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgetActionsForAccount, "describe-budget-actions-for-account", "", false, "Describe Budget Actions For Account")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgetActionsForBudget, "describe-budget-actions-for-budget", "", false, "Describe Budget Actions For Budget")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgetNotificationsForAccount, "describe-budget-notifications-for-account", "", false, "Describe Budget Notifications For Account")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgetPerformanceHistory, "describe-budget-performance-history", "", false, "Describe Budget Performance History")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeBudgets, "describe-budgets", "", false, "Describe Budgets")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeNotificationsForBudget, "describe-notifications-for-budget", "", false, "Describe Notifications For Budget")
	_budgetsCmd.Flags().BoolVarP(&_budgetsDescribeSubscribersForNotification, "describe-subscribers-for-notification", "", false, "Describe Subscribers For Notification")
	_budgetsCmd.Flags().BoolVarP(&_budgetsExecuteBudgetAction, "execute-budget-action", "", false, "Execute Budget Action")
	_budgetsCmd.Flags().BoolVarP(&_budgetsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_budgetsCmd.Flags().BoolVarP(&_budgetsTagResource, "tag-resource", "", false, "Tag Resource")
	_budgetsCmd.Flags().BoolVarP(&_budgetsUntagResource, "untag-resource", "", false, "Untag Resource")
	_budgetsCmd.Flags().BoolVarP(&_budgetsUpdateBudget, "update-budget", "", false, "Update Budget")
	_budgetsCmd.Flags().BoolVarP(&_budgetsUpdateBudgetAction, "update-budget-action", "", false, "Update Budget Action")
	_budgetsCmd.Flags().BoolVarP(&_budgetsUpdateNotification, "update-notification", "", false, "Update Notification")
	_budgetsCmd.Flags().BoolVarP(&_budgetsUpdateSubscriber, "update-subscriber", "", false, "Update Subscriber")

}
