package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ioteventsdataCmd represents the ioteventsdata command
var _ioteventsdataCmd = &cobra.Command{
	Use:   "ioteventsdata",
	Short: "AWS ioteventsdata CLI",
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
		client := ioteventsdata.NewFromConfig(cfg)
		if _ioteventsdataBatchAcknowledgeAlarm {
			ioteventsdata_BatchAcknowledgeAlarm(cfg, client)
			return
		}
		if _ioteventsdataBatchDeleteDetector {
			ioteventsdata_BatchDeleteDetector(cfg, client)
			return
		}
		if _ioteventsdataBatchDisableAlarm {
			ioteventsdata_BatchDisableAlarm(cfg, client)
			return
		}
		if _ioteventsdataBatchEnableAlarm {
			ioteventsdata_BatchEnableAlarm(cfg, client)
			return
		}
		if _ioteventsdataBatchPutMessage {
			ioteventsdata_BatchPutMessage(cfg, client)
			return
		}
		if _ioteventsdataBatchResetAlarm {
			ioteventsdata_BatchResetAlarm(cfg, client)
			return
		}
		if _ioteventsdataBatchSnoozeAlarm {
			ioteventsdata_BatchSnoozeAlarm(cfg, client)
			return
		}
		if _ioteventsdataBatchUpdateDetector {
			ioteventsdata_BatchUpdateDetector(cfg, client)
			return
		}
		if _ioteventsdataDescribeAlarm {
			ioteventsdata_DescribeAlarm(cfg, client)
			return
		}
		if _ioteventsdataDescribeDetector {
			ioteventsdata_DescribeDetector(cfg, client)
			return
		}
		if _ioteventsdataListAlarms {
			ioteventsdata_ListAlarms(cfg, client)
			return
		}
		if _ioteventsdataListDetectors {
			ioteventsdata_ListDetectors(cfg, client)
			return
		}

	},
}

var (
	_ioteventsdataBatchAcknowledgeAlarm bool
	_ioteventsdataBatchDeleteDetector   bool
	_ioteventsdataBatchDisableAlarm     bool
	_ioteventsdataBatchEnableAlarm      bool
	_ioteventsdataBatchPutMessage       bool
	_ioteventsdataBatchResetAlarm       bool
	_ioteventsdataBatchSnoozeAlarm      bool
	_ioteventsdataBatchUpdateDetector   bool
	_ioteventsdataDescribeAlarm         bool
	_ioteventsdataDescribeDetector      bool
	_ioteventsdataListAlarms            bool
	_ioteventsdataListDetectors         bool

	_ioteventsdataAcknowledgeActionRequests string
	_ioteventsdataAlarmModelName            string
	_ioteventsdataDetectorModelName         string
	_ioteventsdataDetectors                 string
	_ioteventsdataDisableActionRequests     string
	_ioteventsdataEnableActionRequests      string
	_ioteventsdataKeyValue                  string
	_ioteventsdataMaxResults                string
	_ioteventsdataMessages                  string
	_ioteventsdataNextToken                 string
	_ioteventsdataResetActionRequests       string
	_ioteventsdataSnoozeActionRequests      string
	_ioteventsdataStateName                 string
)

// Acknowledges one or more alarms. The alarms change to the ACKNOWLEDGED state
// after you acknowledge them.
func ioteventsdata_BatchAcknowledgeAlarm(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchAcknowledgeAlarmInput{
		// AcknowledgeActionRequests: []types.AcknowledgeAlarmActionRequest, // Required
	}

	if len(_ioteventsdataAcknowledgeActionRequests) > 0 {
		if err := assignInputField(input, "AcknowledgeActionRequests", _ioteventsdataAcknowledgeActionRequests); err != nil {
			log.Errorf("invalid --acknowledge-action-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchAcknowledgeAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more detectors that were created. When a detector is deleted,
// its state will be cleared and the detector will be removed from the list of
// detectors. The deleted detector will no longer appear if referenced in the [ListDetectors]API
// call.
//
// [ListDetectors]: https://docs.aws.amazon.com/iotevents/latest/apireference/API_iotevents-data_ListDetectors.html
func ioteventsdata_BatchDeleteDetector(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchDeleteDetectorInput{
		// Detectors: []types.DeleteDetectorRequest, // Required
	}

	if len(_ioteventsdataDetectors) > 0 {
		if err := assignInputField(input, "Detectors", _ioteventsdataDetectors); err != nil {
			log.Errorf("invalid --detectors: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDeleteDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables one or more alarms. The alarms change to the DISABLED state after you
// disable them.
func ioteventsdata_BatchDisableAlarm(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchDisableAlarmInput{
		// DisableActionRequests: []types.DisableAlarmActionRequest, // Required
	}

	if len(_ioteventsdataDisableActionRequests) > 0 {
		if err := assignInputField(input, "DisableActionRequests", _ioteventsdataDisableActionRequests); err != nil {
			log.Errorf("invalid --disable-action-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchDisableAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables one or more alarms. The alarms change to the NORMAL state after you
// enable them.
func ioteventsdata_BatchEnableAlarm(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchEnableAlarmInput{
		// EnableActionRequests: []types.EnableAlarmActionRequest, // Required
	}

	if len(_ioteventsdataEnableActionRequests) > 0 {
		if err := assignInputField(input, "EnableActionRequests", _ioteventsdataEnableActionRequests); err != nil {
			log.Errorf("invalid --enable-action-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchEnableAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a set of messages to the IoT Events system. Each message payload is
// transformed into the input you specify ( "inputName" ) and ingested into any
// detectors that monitor that input. If multiple messages are sent, the order in
// which the messages are processed isn't guaranteed. To guarantee ordering, you
// must send messages one at a time and wait for a successful response.
func ioteventsdata_BatchPutMessage(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchPutMessageInput{
		// Messages: []types.Message, // Required
	}

	if len(_ioteventsdataMessages) > 0 {
		if err := assignInputField(input, "Messages", _ioteventsdataMessages); err != nil {
			log.Errorf("invalid --messages: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchPutMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resets one or more alarms. The alarms return to the NORMAL state after you
// reset them.
func ioteventsdata_BatchResetAlarm(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchResetAlarmInput{
		// ResetActionRequests: []types.ResetAlarmActionRequest, // Required
	}

	if len(_ioteventsdataResetActionRequests) > 0 {
		if err := assignInputField(input, "ResetActionRequests", _ioteventsdataResetActionRequests); err != nil {
			log.Errorf("invalid --reset-action-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchResetAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Changes one or more alarms to the snooze mode. The alarms change to the
// SNOOZE_DISABLED state after you set them to the snooze mode.
func ioteventsdata_BatchSnoozeAlarm(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchSnoozeAlarmInput{
		// SnoozeActionRequests: []types.SnoozeAlarmActionRequest, // Required
	}

	if len(_ioteventsdataSnoozeActionRequests) > 0 {
		if err := assignInputField(input, "SnoozeActionRequests", _ioteventsdataSnoozeActionRequests); err != nil {
			log.Errorf("invalid --snooze-action-requests: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchSnoozeAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the state, variable values, and timer settings of one or more detectors
// (instances) of a specified detector model.
func ioteventsdata_BatchUpdateDetector(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.BatchUpdateDetectorInput{
		// Detectors: []types.UpdateDetectorRequest, // Required
	}

	if len(_ioteventsdataDetectors) > 0 {
		if err := assignInputField(input, "Detectors", _ioteventsdataDetectors); err != nil {
			log.Errorf("invalid --detectors: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an alarm.
func ioteventsdata_DescribeAlarm(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.DescribeAlarmInput{
		// AlarmModelName: *string, // Required
	}

	if len(_ioteventsdataAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsdataAlarmModelName)
	}
	if len(_ioteventsdataKeyValue) > 0 {
		input.KeyValue = aws.String(_ioteventsdataKeyValue)
	}

	if resp, err := client.DescribeAlarm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the specified detector (instance).
func ioteventsdata_DescribeDetector(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.DescribeDetectorInput{
		// DetectorModelName: *string, // Required
	}

	if len(_ioteventsdataDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsdataDetectorModelName)
	}
	if len(_ioteventsdataKeyValue) > 0 {
		input.KeyValue = aws.String(_ioteventsdataKeyValue)
	}

	if resp, err := client.DescribeDetector(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists one or more alarms. The operation returns only the metadata associated
// with each alarm.
func ioteventsdata_ListAlarms(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.ListAlarmsInput{
		// AlarmModelName: *string, // Required
	}

	if len(_ioteventsdataAlarmModelName) > 0 {
		input.AlarmModelName = aws.String(_ioteventsdataAlarmModelName)
	}
	if len(_ioteventsdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsdataNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsdataNextToken)
	}

	if resp, err := client.ListAlarms(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists detectors (the instances of a detector model).
func ioteventsdata_ListDetectors(cfg aws.Config, client *ioteventsdata.Client) {
	input := &ioteventsdata.ListDetectorsInput{
		// DetectorModelName: *string, // Required
	}

	if len(_ioteventsdataDetectorModelName) > 0 {
		input.DetectorModelName = aws.String(_ioteventsdataDetectorModelName)
	}
	if len(_ioteventsdataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ioteventsdataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ioteventsdataNextToken) > 0 {
		input.NextToken = aws.String(_ioteventsdataNextToken)
	}
	if len(_ioteventsdataStateName) > 0 {
		input.StateName = aws.String(_ioteventsdataStateName)
	}

	if resp, err := client.ListDetectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ioteventsdataCmd)
	_ioteventsdataCmd.Flags().SortFlags = false

	_ioteventsdataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_ioteventsdataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ioteventsdataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataAcknowledgeActionRequests, "acknowledge-action-requests", "", "", "Acknowledge Action Requests")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataAlarmModelName, "alarm-model-name", "", "", "Alarm Model Name")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataDetectorModelName, "detector-model-name", "", "", "Detector Model Name")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataDetectors, "detectors", "", "", "Detectors")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataDisableActionRequests, "disable-action-requests", "", "", "Disable Action Requests")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataEnableActionRequests, "enable-action-requests", "", "", "Enable Action Requests")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataKeyValue, "key-value", "", "", "Key Value")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataMaxResults, "max-results", "", "", "Max Results")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataMessages, "messages", "", "", "Messages")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataNextToken, "next-token", "", "", "Next Token")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataResetActionRequests, "reset-action-requests", "", "", "Reset Action Requests")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataSnoozeActionRequests, "snooze-action-requests", "", "", "Snooze Action Requests")
	_ioteventsdataCmd.Flags().StringVarP(&_ioteventsdataStateName, "state-name", "", "", "State Name")

	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchAcknowledgeAlarm, "batch-acknowledge-alarm", "", false, "Batch Acknowledge Alarm")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchDeleteDetector, "batch-delete-detector", "", false, "Batch Delete Detector")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchDisableAlarm, "batch-disable-alarm", "", false, "Batch Disable Alarm")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchEnableAlarm, "batch-enable-alarm", "", false, "Batch Enable Alarm")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchPutMessage, "batch-put-message", "", false, "Batch Put Message")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchResetAlarm, "batch-reset-alarm", "", false, "Batch Reset Alarm")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchSnoozeAlarm, "batch-snooze-alarm", "", false, "Batch Snooze Alarm")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataBatchUpdateDetector, "batch-update-detector", "", false, "Batch Update Detector")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataDescribeAlarm, "describe-alarm", "", false, "Describe Alarm")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataDescribeDetector, "describe-detector", "", false, "Describe Detector")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataListAlarms, "list-alarms", "", false, "List Alarms")
	_ioteventsdataCmd.Flags().BoolVarP(&_ioteventsdataListDetectors, "list-detectors", "", false, "List Detectors")

}
