package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/arczonalshift"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// arczonalshiftCmd represents the arczonalshift command
var _arczonalshiftCmd = &cobra.Command{
	Use:   "arczonalshift",
	Short: "AWS arczonalshift CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := arczonalshift.NewFromConfig(cfg)
		if _arczonalshiftCancelPracticeRun {
			arczonalshift_CancelPracticeRun(cfg, client)
			return
		}
		if _arczonalshiftCancelZonalShift {
			arczonalshift_CancelZonalShift(cfg, client)
			return
		}
		if _arczonalshiftCreatePracticeRunConfiguration {
			arczonalshift_CreatePracticeRunConfiguration(cfg, client)
			return
		}
		if _arczonalshiftDeletePracticeRunConfiguration {
			arczonalshift_DeletePracticeRunConfiguration(cfg, client)
			return
		}
		if _arczonalshiftGetAutoshiftObserverNotificationStatus {
			arczonalshift_GetAutoshiftObserverNotificationStatus(cfg, client)
			return
		}
		if _arczonalshiftGetManagedResource {
			arczonalshift_GetManagedResource(cfg, client)
			return
		}
		if _arczonalshiftListAutoshifts {
			arczonalshift_ListAutoshifts(cfg, client)
			return
		}
		if _arczonalshiftListManagedResources {
			arczonalshift_ListManagedResources(cfg, client)
			return
		}
		if _arczonalshiftListZonalShifts {
			arczonalshift_ListZonalShifts(cfg, client)
			return
		}
		if _arczonalshiftStartPracticeRun {
			arczonalshift_StartPracticeRun(cfg, client)
			return
		}
		if _arczonalshiftStartZonalShift {
			arczonalshift_StartZonalShift(cfg, client)
			return
		}
		if _arczonalshiftUpdateAutoshiftObserverNotificationStatus {
			arczonalshift_UpdateAutoshiftObserverNotificationStatus(cfg, client)
			return
		}
		if _arczonalshiftUpdatePracticeRunConfiguration {
			arczonalshift_UpdatePracticeRunConfiguration(cfg, client)
			return
		}
		if _arczonalshiftUpdateZonalAutoshiftConfiguration {
			arczonalshift_UpdateZonalAutoshiftConfiguration(cfg, client)
			return
		}
		if _arczonalshiftUpdateZonalShift {
			arczonalshift_UpdateZonalShift(cfg, client)
			return
		}

	},
}

var (
	_arczonalshiftCancelPracticeRun                         bool
	_arczonalshiftCancelZonalShift                          bool
	_arczonalshiftCreatePracticeRunConfiguration            bool
	_arczonalshiftDeletePracticeRunConfiguration            bool
	_arczonalshiftGetAutoshiftObserverNotificationStatus    bool
	_arczonalshiftGetManagedResource                        bool
	_arczonalshiftListAutoshifts                            bool
	_arczonalshiftListManagedResources                      bool
	_arczonalshiftListZonalShifts                           bool
	_arczonalshiftStartPracticeRun                          bool
	_arczonalshiftStartZonalShift                           bool
	_arczonalshiftUpdateAutoshiftObserverNotificationStatus bool
	_arczonalshiftUpdatePracticeRunConfiguration            bool
	_arczonalshiftUpdateZonalAutoshiftConfiguration         bool
	_arczonalshiftUpdateZonalShift                          bool

	_arczonalshiftAllowedWindows       []string
	_arczonalshiftAwayFrom             string
	_arczonalshiftBlockedDates         []string
	_arczonalshiftBlockedWindows       []string
	_arczonalshiftBlockingAlarms       string
	_arczonalshiftComment              string
	_arczonalshiftExpiresIn            string
	_arczonalshiftMaxResults           string
	_arczonalshiftNextToken            string
	_arczonalshiftOutcomeAlarms        string
	_arczonalshiftResourceIdentifier   string
	_arczonalshiftStatus               string
	_arczonalshiftZonalAutoshiftStatus string
	_arczonalshiftZonalShiftId         string
)

// Cancel an in-progress practice run zonal shift in Amazon Application Recovery
// Controller.
func arczonalshift_CancelPracticeRun(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.CancelPracticeRunInput{
		// ZonalShiftId: *string, // Required
	}

	if len(_arczonalshiftZonalShiftId) > 0 {
		input.ZonalShiftId = aws.String(_arczonalshiftZonalShiftId)
	}

	if resp, err := client.CancelPracticeRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancel a zonal shift in Amazon Application Recovery Controller. To cancel the
// zonal shift, specify the zonal shift ID.
//
// A zonal shift can be one that you've started for a resource in your Amazon Web
// Services account in an Amazon Web Services Region, or it can be a zonal shift
// started by a practice run with zonal autoshift.
func arczonalshift_CancelZonalShift(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.CancelZonalShiftInput{
		// ZonalShiftId: *string, // Required
	}

	if len(_arczonalshiftZonalShiftId) > 0 {
		input.ZonalShiftId = aws.String(_arczonalshiftZonalShiftId)
	}

	if resp, err := client.CancelZonalShift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A practice run configuration for zonal autoshift is required when you enable
// zonal autoshift. A practice run configuration includes specifications for
// blocked dates and blocked time windows, and for Amazon CloudWatch alarms that
// you create to use with practice runs. The alarms that you specify are an outcome
// alarm, to monitor application health during practice runs and, optionally, a
// blocking alarm, to block practice runs from starting.
//
// When a resource has a practice run configuration, ARC starts zonal shifts for
// the resource weekly, to shift traffic for practice runs. Practice runs help you
// to ensure that shifting away traffic from an Availability Zone during an
// autoshift is safe for your application.
//
// For more information, see [Considerations when you configure zonal autoshift] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Considerations when you configure zonal autoshift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.considerations.html
func arczonalshift_CreatePracticeRunConfiguration(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.CreatePracticeRunConfigurationInput{
		// OutcomeAlarms: []types.ControlCondition, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_arczonalshiftOutcomeAlarms) > 0 {
		if err := assignInputField(input, "OutcomeAlarms", _arczonalshiftOutcomeAlarms); err != nil {
			log.Errorf("invalid --outcome-alarms: %s", err.Error())
			return
		}
	}
	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}
	if len(_arczonalshiftAllowedWindows) > 0 {
		input.AllowedWindows = append([]string(nil), _arczonalshiftAllowedWindows...)
	}
	if len(_arczonalshiftBlockedDates) > 0 {
		input.BlockedDates = append([]string(nil), _arczonalshiftBlockedDates...)
	}
	if len(_arczonalshiftBlockedWindows) > 0 {
		input.BlockedWindows = append([]string(nil), _arczonalshiftBlockedWindows...)
	}
	if len(_arczonalshiftBlockingAlarms) > 0 {
		if err := assignInputField(input, "BlockingAlarms", _arczonalshiftBlockingAlarms); err != nil {
			log.Errorf("invalid --blocking-alarms: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePracticeRunConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the practice run configuration for a resource. Before you can delete a
// practice run configuration for a resource., you must disable zonal autoshift for
// the resource. Practice runs must be configured for zonal autoshift to be
// enabled.
func arczonalshift_DeletePracticeRunConfiguration(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.DeletePracticeRunConfigurationInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}

	if resp, err := client.DeletePracticeRunConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status of the autoshift observer notification. Autoshift observer
// notifications notify you through Amazon EventBridge when there is an autoshift
// event for zonal autoshift. The status can be ENABLED or DISABLED . When ENABLED
// , a notification is sent when an autoshift is triggered. When DISABLED ,
// notifications are not sent.
func arczonalshift_GetAutoshiftObserverNotificationStatus(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.GetAutoshiftObserverNotificationStatusInput{}

	if resp, err := client.GetAutoshiftObserverNotificationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about a resource that's been registered for zonal shifts with
// Amazon Application Recovery Controller in this Amazon Web Services Region.
// Resources that are registered for zonal shifts are managed resources in ARC. You
// can start zonal shifts and configure zonal autoshift for managed resources.
func arczonalshift_GetManagedResource(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.GetManagedResourceInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}

	if resp, err := client.GetManagedResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the autoshifts for an Amazon Web Services Region. By default, the call
// returns only ACTIVE autoshifts. Optionally, you can specify the status
// parameter to return COMPLETED autoshifts.
func arczonalshift_ListAutoshifts(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.ListAutoshiftsInput{}

	if len(_arczonalshiftMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arczonalshiftMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arczonalshiftNextToken) > 0 {
		input.NextToken = aws.String(_arczonalshiftNextToken)
	}
	if len(_arczonalshiftStatus) > 0 {
		if err := assignInputField(input, "Status", _arczonalshiftStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAutoshifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arczonalshift.ListAutoshiftsOutput
	p := arczonalshift.NewListAutoshiftsPaginator(client, input)
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

// Lists all the resources in your Amazon Web Services account in this Amazon Web
// Services Region that are managed for zonal shifts in Amazon Application Recovery
// Controller, and information about them. The information includes the zonal
// autoshift status for the resource, as well as the Amazon Resource Name (ARN),
// the Availability Zones that each resource is deployed in, and the resource name.
func arczonalshift_ListManagedResources(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.ListManagedResourcesInput{}

	if len(_arczonalshiftMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arczonalshiftMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arczonalshiftNextToken) > 0 {
		input.NextToken = aws.String(_arczonalshiftNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arczonalshift.ListManagedResourcesOutput
	p := arczonalshift.NewListManagedResourcesPaginator(client, input)
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

// Lists all active and completed zonal shifts in Amazon Application Recovery
// Controller in your Amazon Web Services account in this Amazon Web Services
// Region. ListZonalShifts returns customer-initiated zonal shifts, as well as
// practice run zonal shifts that ARC started on your behalf for zonal autoshift.
//
// For more information about listing autoshifts, see [">ListAutoshifts].
//
// [">ListAutoshifts]: https://docs.aws.amazon.com/arc-zonal-shift/latest/api/API_ListAutoshifts.html
func arczonalshift_ListZonalShifts(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.ListZonalShiftsInput{}

	if len(_arczonalshiftMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _arczonalshiftMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_arczonalshiftNextToken) > 0 {
		input.NextToken = aws.String(_arczonalshiftNextToken)
	}
	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}
	if len(_arczonalshiftStatus) > 0 {
		if err := assignInputField(input, "Status", _arczonalshiftStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListZonalShifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*arczonalshift.ListZonalShiftsOutput
	p := arczonalshift.NewListZonalShiftsPaginator(client, input)
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

// Start an on-demand practice run zonal shift in Amazon Application Recovery
// Controller. With zonal autoshift enabled, you can start an on-demand practice
// run to verify preparedness at any time. Amazon Web Services also runs automated
// practice runs about weekly when you have enabled zonal autoshift.
//
// For more information, see [Considerations when you configure zonal autoshift] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Considerations when you configure zonal autoshift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.considerations.html
func arczonalshift_StartPracticeRun(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.StartPracticeRunInput{
		// AwayFrom: *string, // Required
		// Comment: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_arczonalshiftAwayFrom) > 0 {
		input.AwayFrom = aws.String(_arczonalshiftAwayFrom)
	}
	if len(_arczonalshiftComment) > 0 {
		input.Comment = aws.String(_arczonalshiftComment)
	}
	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}

	if resp, err := client.StartPracticeRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You start a zonal shift to temporarily move load balancer traffic away from an
// Availability Zone in an Amazon Web Services Region, to help your application
// recover immediately, for example, from a developer's bad code deployment or from
// an Amazon Web Services infrastructure failure in a single Availability Zone. You
// can start a zonal shift in ARC only for managed resources in your Amazon Web
// Services account in an Amazon Web Services Region. Resources are automatically
// registered with ARC by Amazon Web Services services.
//
// Amazon Application Recovery Controller currently supports enabling the
// following resources for zonal shift and zonal autoshift:
//
// [Amazon EC2 Auto Scaling groups]
//
// [Amazon Elastic Kubernetes Service]
//
// [Application Load Balancer]
//
// [Network Load Balancer]
//
// When you start a zonal shift, traffic for the resource is no longer routed to
// the Availability Zone. The zonal shift is created immediately in ARC. However,
// it can take a short time, typically up to a few minutes, for existing,
// in-progress connections in the Availability Zone to complete.
//
// For more information, see [Zonal shift] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Amazon EC2 Auto Scaling groups]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.ec2-auto-scaling-groups.html
// [Amazon Elastic Kubernetes Service]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.eks.html
// [Application Load Balancer]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.app-load-balancers.html
// [Network Load Balancer]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.resource-types.network-load-balancers.html
// [Zonal shift]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-shift.html
func arczonalshift_StartZonalShift(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.StartZonalShiftInput{
		// AwayFrom: *string, // Required
		// Comment: *string, // Required
		// ExpiresIn: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_arczonalshiftAwayFrom) > 0 {
		input.AwayFrom = aws.String(_arczonalshiftAwayFrom)
	}
	if len(_arczonalshiftComment) > 0 {
		input.Comment = aws.String(_arczonalshiftComment)
	}
	if len(_arczonalshiftExpiresIn) > 0 {
		input.ExpiresIn = aws.String(_arczonalshiftExpiresIn)
	}
	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}

	if resp, err := client.StartZonalShift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the status of autoshift observer notification. Autoshift observer
// notification enables you to be notified, through Amazon EventBridge, when there
// is an autoshift event for zonal autoshift.
//
// If the status is ENABLED , ARC includes all autoshift events when you use the
// EventBridge pattern Autoshift In Progress . When the status is DISABLED , ARC
// includes only autoshift events for autoshifts when one or more of your resources
// is included in the autoshift.
//
// For more information, see [Notifications for practice runs and autoshifts] in the Amazon Application Recovery Controller
// Developer Guide.
//
// [Notifications for practice runs and autoshifts]: https://docs.aws.amazon.com/r53recovery/latest/dg/arc-zonal-autoshift.how-it-works.html#ZAShiftNotification
func arczonalshift_UpdateAutoshiftObserverNotificationStatus(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.UpdateAutoshiftObserverNotificationStatusInput{
		// Status: types.AutoshiftObserverNotificationStatus, // Required
	}

	if len(_arczonalshiftStatus) > 0 {
		if err := assignInputField(input, "Status", _arczonalshiftStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAutoshiftObserverNotificationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a practice run configuration to change one or more of the following:
// add, change, or remove the blocking alarm; change the outcome alarm; or add,
// change, or remove blocking dates or time windows.
func arczonalshift_UpdatePracticeRunConfiguration(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.UpdatePracticeRunConfigurationInput{
		// ResourceIdentifier: *string, // Required
	}

	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}
	if len(_arczonalshiftAllowedWindows) > 0 {
		input.AllowedWindows = append([]string(nil), _arczonalshiftAllowedWindows...)
	}
	if len(_arczonalshiftBlockedDates) > 0 {
		input.BlockedDates = append([]string(nil), _arczonalshiftBlockedDates...)
	}
	if len(_arczonalshiftBlockedWindows) > 0 {
		input.BlockedWindows = append([]string(nil), _arczonalshiftBlockedWindows...)
	}
	if len(_arczonalshiftBlockingAlarms) > 0 {
		if err := assignInputField(input, "BlockingAlarms", _arczonalshiftBlockingAlarms); err != nil {
			log.Errorf("invalid --blocking-alarms: %s", err.Error())
			return
		}
	}
	if len(_arczonalshiftOutcomeAlarms) > 0 {
		if err := assignInputField(input, "OutcomeAlarms", _arczonalshiftOutcomeAlarms); err != nil {
			log.Errorf("invalid --outcome-alarms: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePracticeRunConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The zonal autoshift configuration for a resource includes the practice run
// configuration and the status for running autoshifts, zonal autoshift status.
// When a resource has a practice run configuration, ARC starts weekly zonal shifts
// for the resource, to shift traffic away from an Availability Zone. Weekly
// practice runs help you to make sure that your application can continue to
// operate normally with the loss of one Availability Zone.
//
// You can update the zonal autoshift status to enable or disable zonal autoshift.
// When zonal autoshift is ENABLED , you authorize Amazon Web Services to shift
// away resource traffic for an application from an Availability Zone during
// events, on your behalf, to help reduce time to recovery. Traffic is also shifted
// away for the required weekly practice runs.
func arczonalshift_UpdateZonalAutoshiftConfiguration(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.UpdateZonalAutoshiftConfigurationInput{
		// ResourceIdentifier: *string, // Required
		// ZonalAutoshiftStatus: types.ZonalAutoshiftStatus, // Required
	}

	if len(_arczonalshiftResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_arczonalshiftResourceIdentifier)
	}
	if len(_arczonalshiftZonalAutoshiftStatus) > 0 {
		if err := assignInputField(input, "ZonalAutoshiftStatus", _arczonalshiftZonalAutoshiftStatus); err != nil {
			log.Errorf("invalid --zonal-autoshift-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateZonalAutoshiftConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update an active zonal shift in Amazon Application Recovery Controller in your
// Amazon Web Services account. You can update a zonal shift to set a new
// expiration, or edit or replace the comment for the zonal shift.
func arczonalshift_UpdateZonalShift(cfg aws.Config, client *arczonalshift.Client) {
	input := &arczonalshift.UpdateZonalShiftInput{
		// ZonalShiftId: *string, // Required
	}

	if len(_arczonalshiftZonalShiftId) > 0 {
		input.ZonalShiftId = aws.String(_arczonalshiftZonalShiftId)
	}
	if len(_arczonalshiftComment) > 0 {
		input.Comment = aws.String(_arczonalshiftComment)
	}
	if len(_arczonalshiftExpiresIn) > 0 {
		input.ExpiresIn = aws.String(_arczonalshiftExpiresIn)
	}

	if resp, err := client.UpdateZonalShift(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_arczonalshiftCmd)
	_arczonalshiftCmd.Flags().SortFlags = false

	_arczonalshiftCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_arczonalshiftCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_arczonalshiftCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_arczonalshiftCmd.Flags().StringSliceVarP(&_arczonalshiftAllowedWindows, "allowed-windows", "", nil, "Allowed Windows")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftAwayFrom, "away-from", "", "", "Away From")
	_arczonalshiftCmd.Flags().StringSliceVarP(&_arczonalshiftBlockedDates, "blocked-dates", "", nil, "Blocked Dates")
	_arczonalshiftCmd.Flags().StringSliceVarP(&_arczonalshiftBlockedWindows, "blocked-windows", "", nil, "Blocked Windows")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftBlockingAlarms, "blocking-alarms", "", "", "Blocking Alarms")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftComment, "comment", "", "", "Comment")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftExpiresIn, "expires-in", "", "", "Expires In")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftMaxResults, "max-results", "", "", "Max Results")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftNextToken, "next-token", "", "", "Next Token")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftOutcomeAlarms, "outcome-alarms", "", "", "Outcome Alarms")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftStatus, "status", "", "", "Status")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftZonalAutoshiftStatus, "zonal-autoshift-status", "", "", "Zonal Autoshift Status")
	_arczonalshiftCmd.Flags().StringVarP(&_arczonalshiftZonalShiftId, "zonal-shift-id", "", "", "Zonal Shift ID")

	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftCancelPracticeRun, "cancel-practice-run", "", false, "Cancel Practice Run")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftCancelZonalShift, "cancel-zonal-shift", "", false, "Cancel Zonal Shift")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftCreatePracticeRunConfiguration, "create-practice-run-configuration", "", false, "Create Practice Run Configuration")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftDeletePracticeRunConfiguration, "delete-practice-run-configuration", "", false, "Delete Practice Run Configuration")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftGetAutoshiftObserverNotificationStatus, "get-autoshift-observer-notification-status", "", false, "Get Autoshift Observer Notification Status")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftGetManagedResource, "get-managed-resource", "", false, "Get Managed Resource")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftListAutoshifts, "list-autoshifts", "", false, "List Autoshifts")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftListManagedResources, "list-managed-resources", "", false, "List Managed Resources")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftListZonalShifts, "list-zonal-shifts", "", false, "List Zonal Shifts")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftStartPracticeRun, "start-practice-run", "", false, "Start Practice Run")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftStartZonalShift, "start-zonal-shift", "", false, "Start Zonal Shift")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftUpdateAutoshiftObserverNotificationStatus, "update-autoshift-observer-notification-status", "", false, "Update Autoshift Observer Notification Status")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftUpdatePracticeRunConfiguration, "update-practice-run-configuration", "", false, "Update Practice Run Configuration")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftUpdateZonalAutoshiftConfiguration, "update-zonal-autoshift-configuration", "", false, "Update Zonal Autoshift Configuration")
	_arczonalshiftCmd.Flags().BoolVarP(&_arczonalshiftUpdateZonalShift, "update-zonal-shift", "", false, "Update Zonal Shift")

}
