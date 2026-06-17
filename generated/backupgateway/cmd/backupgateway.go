package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backupgateway"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// backupgatewayCmd represents the backupgateway command
var _backupgatewayCmd = &cobra.Command{
	Use:   "backupgateway",
	Short: "AWS backupgateway CLI",
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
		client := backupgateway.NewFromConfig(cfg)
		if _backupgatewayAssociateGatewayToServer {
			backupgateway_AssociateGatewayToServer(cfg, client)
			return
		}
		if _backupgatewayCreateGateway {
			backupgateway_CreateGateway(cfg, client)
			return
		}
		if _backupgatewayDeleteGateway {
			backupgateway_DeleteGateway(cfg, client)
			return
		}
		if _backupgatewayDeleteHypervisor {
			backupgateway_DeleteHypervisor(cfg, client)
			return
		}
		if _backupgatewayDisassociateGatewayFromServer {
			backupgateway_DisassociateGatewayFromServer(cfg, client)
			return
		}
		if _backupgatewayGetBandwidthRateLimitSchedule {
			backupgateway_GetBandwidthRateLimitSchedule(cfg, client)
			return
		}
		if _backupgatewayGetGateway {
			backupgateway_GetGateway(cfg, client)
			return
		}
		if _backupgatewayGetHypervisor {
			backupgateway_GetHypervisor(cfg, client)
			return
		}
		if _backupgatewayGetHypervisorPropertyMappings {
			backupgateway_GetHypervisorPropertyMappings(cfg, client)
			return
		}
		if _backupgatewayGetVirtualMachine {
			backupgateway_GetVirtualMachine(cfg, client)
			return
		}
		if _backupgatewayImportHypervisorConfiguration {
			backupgateway_ImportHypervisorConfiguration(cfg, client)
			return
		}
		if _backupgatewayListGateways {
			backupgateway_ListGateways(cfg, client)
			return
		}
		if _backupgatewayListHypervisors {
			backupgateway_ListHypervisors(cfg, client)
			return
		}
		if _backupgatewayListTagsForResource {
			backupgateway_ListTagsForResource(cfg, client)
			return
		}
		if _backupgatewayListVirtualMachines {
			backupgateway_ListVirtualMachines(cfg, client)
			return
		}
		if _backupgatewayPutBandwidthRateLimitSchedule {
			backupgateway_PutBandwidthRateLimitSchedule(cfg, client)
			return
		}
		if _backupgatewayPutHypervisorPropertyMappings {
			backupgateway_PutHypervisorPropertyMappings(cfg, client)
			return
		}
		if _backupgatewayPutMaintenanceStartTime {
			backupgateway_PutMaintenanceStartTime(cfg, client)
			return
		}
		if _backupgatewayStartVirtualMachinesMetadataSync {
			backupgateway_StartVirtualMachinesMetadataSync(cfg, client)
			return
		}
		if _backupgatewayTagResource {
			backupgateway_TagResource(cfg, client)
			return
		}
		if _backupgatewayTestHypervisorConfiguration {
			backupgateway_TestHypervisorConfiguration(cfg, client)
			return
		}
		if _backupgatewayUntagResource {
			backupgateway_UntagResource(cfg, client)
			return
		}
		if _backupgatewayUpdateGatewayInformation {
			backupgateway_UpdateGatewayInformation(cfg, client)
			return
		}
		if _backupgatewayUpdateGatewaySoftwareNow {
			backupgateway_UpdateGatewaySoftwareNow(cfg, client)
			return
		}
		if _backupgatewayUpdateHypervisor {
			backupgateway_UpdateHypervisor(cfg, client)
			return
		}

	},
}

var (
	_backupgatewayAssociateGatewayToServer         bool
	_backupgatewayCreateGateway                    bool
	_backupgatewayDeleteGateway                    bool
	_backupgatewayDeleteHypervisor                 bool
	_backupgatewayDisassociateGatewayFromServer    bool
	_backupgatewayGetBandwidthRateLimitSchedule    bool
	_backupgatewayGetGateway                       bool
	_backupgatewayGetHypervisor                    bool
	_backupgatewayGetHypervisorPropertyMappings    bool
	_backupgatewayGetVirtualMachine                bool
	_backupgatewayImportHypervisorConfiguration    bool
	_backupgatewayListGateways                     bool
	_backupgatewayListHypervisors                  bool
	_backupgatewayListTagsForResource              bool
	_backupgatewayListVirtualMachines              bool
	_backupgatewayPutBandwidthRateLimitSchedule    bool
	_backupgatewayPutHypervisorPropertyMappings    bool
	_backupgatewayPutMaintenanceStartTime          bool
	_backupgatewayStartVirtualMachinesMetadataSync bool
	_backupgatewayTagResource                      bool
	_backupgatewayTestHypervisorConfiguration      bool
	_backupgatewayUntagResource                    bool
	_backupgatewayUpdateGatewayInformation         bool
	_backupgatewayUpdateGatewaySoftwareNow         bool
	_backupgatewayUpdateHypervisor                 bool

	_backupgatewayActivationKey               string
	_backupgatewayBandwidthRateLimitIntervals string
	_backupgatewayDayOfMonth                  string
	_backupgatewayDayOfWeek                   string
	_backupgatewayGatewayArn                  string
	_backupgatewayGatewayDisplayName          string
	_backupgatewayGatewayType                 string
	_backupgatewayHost                        string
	_backupgatewayHourOfDay                   string
	_backupgatewayHypervisorArn               string
	_backupgatewayIamRoleArn                  string
	_backupgatewayKmsKeyArn                   string
	_backupgatewayLogGroupArn                 string
	_backupgatewayMaxResults                  string
	_backupgatewayMinuteOfHour                string
	_backupgatewayName                        string
	_backupgatewayNextToken                   string
	_backupgatewayPassword                    string
	_backupgatewayResourceARN                 string
	_backupgatewayServerArn                   string
	_backupgatewayTagKeys                     []string
	_backupgatewayTags                        string
	_backupgatewayUsername                    string
	_backupgatewayVmwareToAwsTagMappings      string
)

// Associates a backup gateway with your server. After you complete the
// association process, you can back up and restore your VMs through the gateway.
func backupgateway_AssociateGatewayToServer(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.AssociateGatewayToServerInput{
		// GatewayArn: *string, // Required
		// ServerArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}
	if len(_backupgatewayServerArn) > 0 {
		input.ServerArn = aws.String(_backupgatewayServerArn)
	}

	if resp, err := client.AssociateGatewayToServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a backup gateway. After you create a gateway, you can associate it with
// a server using the AssociateGatewayToServer operation.
func backupgateway_CreateGateway(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.CreateGatewayInput{
		// ActivationKey: *string, // Required
		// GatewayDisplayName: *string, // Required
		// GatewayType: types.GatewayType, // Required
	}

	if len(_backupgatewayActivationKey) > 0 {
		input.ActivationKey = aws.String(_backupgatewayActivationKey)
	}
	if len(_backupgatewayGatewayDisplayName) > 0 {
		input.GatewayDisplayName = aws.String(_backupgatewayGatewayDisplayName)
	}
	if len(_backupgatewayGatewayType) > 0 {
		if err := assignInputField(input, "GatewayType", _backupgatewayGatewayType); err != nil {
			log.Errorf("invalid --gateway-type: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _backupgatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a backup gateway.
func backupgateway_DeleteGateway(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.DeleteGatewayInput{
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}

	if resp, err := client.DeleteGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a hypervisor.
func backupgateway_DeleteHypervisor(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.DeleteHypervisorInput{
		// HypervisorArn: *string, // Required
	}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}

	if resp, err := client.DeleteHypervisor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a backup gateway from the specified server. After the
// disassociation process finishes, the gateway can no longer access the virtual
// machines on the server.
func backupgateway_DisassociateGatewayFromServer(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.DisassociateGatewayFromServerInput{
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}

	if resp, err := client.DisassociateGatewayFromServer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the bandwidth rate limit schedule for a specified gateway. By
// default, gateways do not have bandwidth rate limit schedules, which means no
// bandwidth rate limiting is in effect. Use this to get a gateway's bandwidth rate
// limit schedule.
func backupgateway_GetBandwidthRateLimitSchedule(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.GetBandwidthRateLimitScheduleInput{
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}

	if resp, err := client.GetBandwidthRateLimitSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// By providing the ARN (Amazon Resource Name), this API returns the gateway.
func backupgateway_GetGateway(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.GetGatewayInput{
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}

	if resp, err := client.GetGateway(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action requests information about the specified hypervisor to which the
// gateway will connect. A hypervisor is hardware, software, or firmware that
// creates and manages virtual machines, and allocates resources to them.
func backupgateway_GetHypervisor(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.GetHypervisorInput{
		// HypervisorArn: *string, // Required
	}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}

	if resp, err := client.GetHypervisor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action retrieves the property mappings for the specified hypervisor. A
// hypervisor property mapping displays the relationship of entity properties
// available from the hypervisor to the properties available in Amazon Web
// Services.
func backupgateway_GetHypervisorPropertyMappings(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.GetHypervisorPropertyMappingsInput{
		// HypervisorArn: *string, // Required
	}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}

	if resp, err := client.GetHypervisorPropertyMappings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// By providing the ARN (Amazon Resource Name), this API returns the virtual
// machine.
func backupgateway_GetVirtualMachine(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.GetVirtualMachineInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupgatewayResourceARN) > 0 {
		input.ResourceArn = aws.String(_backupgatewayResourceARN)
	}

	if resp, err := client.GetVirtualMachine(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Connect to a hypervisor by importing its configuration.
func backupgateway_ImportHypervisorConfiguration(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.ImportHypervisorConfigurationInput{
		// Host: *string, // Required
		// Name: *string, // Required
	}

	if len(_backupgatewayHost) > 0 {
		input.Host = aws.String(_backupgatewayHost)
	}
	if len(_backupgatewayName) > 0 {
		input.Name = aws.String(_backupgatewayName)
	}
	if len(_backupgatewayKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_backupgatewayKmsKeyArn)
	}
	if len(_backupgatewayPassword) > 0 {
		input.Password = aws.String(_backupgatewayPassword)
	}
	if len(_backupgatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _backupgatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayUsername) > 0 {
		input.Username = aws.String(_backupgatewayUsername)
	}

	if resp, err := client.ImportHypervisorConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists backup gateways owned by an Amazon Web Services account in an Amazon Web
// Services Region. The returned list is ordered by gateway Amazon Resource Name
// (ARN).
func backupgateway_ListGateways(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.ListGatewaysInput{}

	if len(_backupgatewayMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupgatewayMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayNextToken) > 0 {
		input.NextToken = aws.String(_backupgatewayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGateways(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupgateway.ListGatewaysOutput
	p := backupgateway.NewListGatewaysPaginator(client, input)
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

// Lists your hypervisors.
func backupgateway_ListHypervisors(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.ListHypervisorsInput{}

	if len(_backupgatewayMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupgatewayMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayNextToken) > 0 {
		input.NextToken = aws.String(_backupgatewayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListHypervisors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupgateway.ListHypervisorsOutput
	p := backupgateway.NewListHypervisorsPaginator(client, input)
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

// Lists the tags applied to the resource identified by its Amazon Resource Name
// (ARN).
func backupgateway_ListTagsForResource(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_backupgatewayResourceARN) > 0 {
		input.ResourceArn = aws.String(_backupgatewayResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your virtual machines.
func backupgateway_ListVirtualMachines(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.ListVirtualMachinesInput{}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}
	if len(_backupgatewayMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _backupgatewayMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayNextToken) > 0 {
		input.NextToken = aws.String(_backupgatewayNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVirtualMachines(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*backupgateway.ListVirtualMachinesOutput
	p := backupgateway.NewListVirtualMachinesPaginator(client, input)
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

// This action sets the bandwidth rate limit schedule for a specified gateway. By
// default, gateways do not have a bandwidth rate limit schedule, which means no
// bandwidth rate limiting is in effect. Use this to initiate a gateway's bandwidth
// rate limit schedule.
func backupgateway_PutBandwidthRateLimitSchedule(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.PutBandwidthRateLimitScheduleInput{
		// BandwidthRateLimitIntervals: []types.BandwidthRateLimitInterval, // Required
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayBandwidthRateLimitIntervals) > 0 {
		if err := assignInputField(input, "BandwidthRateLimitIntervals", _backupgatewayBandwidthRateLimitIntervals); err != nil {
			log.Errorf("invalid --bandwidth-rate-limit-intervals: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}

	if resp, err := client.PutBandwidthRateLimitSchedule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action sets the property mappings for the specified hypervisor. A
// hypervisor property mapping displays the relationship of entity properties
// available from the hypervisor to the properties available in Amazon Web
// Services.
func backupgateway_PutHypervisorPropertyMappings(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.PutHypervisorPropertyMappingsInput{
		// HypervisorArn: *string, // Required
		// IamRoleArn: *string, // Required
		// VmwareToAwsTagMappings: []types.VmwareToAwsTagMapping, // Required
	}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}
	if len(_backupgatewayIamRoleArn) > 0 {
		input.IamRoleArn = aws.String(_backupgatewayIamRoleArn)
	}
	if len(_backupgatewayVmwareToAwsTagMappings) > 0 {
		if err := assignInputField(input, "VmwareToAwsTagMappings", _backupgatewayVmwareToAwsTagMappings); err != nil {
			log.Errorf("invalid --vmware-to-aws-tag-mappings: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutHypervisorPropertyMappings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the maintenance start time for a gateway.
func backupgateway_PutMaintenanceStartTime(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.PutMaintenanceStartTimeInput{
		// GatewayArn: *string, // Required
		// HourOfDay: *int32, // Required
		// MinuteOfHour: *int32, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}
	if len(_backupgatewayHourOfDay) > 0 {
		if err := assignInputField(input, "HourOfDay", _backupgatewayHourOfDay); err != nil {
			log.Errorf("invalid --hour-of-day: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayMinuteOfHour) > 0 {
		if err := assignInputField(input, "MinuteOfHour", _backupgatewayMinuteOfHour); err != nil {
			log.Errorf("invalid --minute-of-hour: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayDayOfMonth) > 0 {
		if err := assignInputField(input, "DayOfMonth", _backupgatewayDayOfMonth); err != nil {
			log.Errorf("invalid --day-of-month: %s", err.Error())
			return
		}
	}
	if len(_backupgatewayDayOfWeek) > 0 {
		if err := assignInputField(input, "DayOfWeek", _backupgatewayDayOfWeek); err != nil {
			log.Errorf("invalid --day-of-week: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMaintenanceStartTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action sends a request to sync metadata across the specified virtual
// machines.
func backupgateway_StartVirtualMachinesMetadataSync(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.StartVirtualMachinesMetadataSyncInput{
		// HypervisorArn: *string, // Required
	}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}

	if resp, err := client.StartVirtualMachinesMetadataSync(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tag the resource.
func backupgateway_TagResource(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_backupgatewayResourceARN) > 0 {
		input.ResourceARN = aws.String(_backupgatewayResourceARN)
	}
	if len(_backupgatewayTags) > 0 {
		if err := assignInputField(input, "Tags", _backupgatewayTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Tests your hypervisor configuration to validate that backup gateway can connect
// with the hypervisor and its resources.
func backupgateway_TestHypervisorConfiguration(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.TestHypervisorConfigurationInput{
		// GatewayArn: *string, // Required
		// Host: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}
	if len(_backupgatewayHost) > 0 {
		input.Host = aws.String(_backupgatewayHost)
	}
	if len(_backupgatewayPassword) > 0 {
		input.Password = aws.String(_backupgatewayPassword)
	}
	if len(_backupgatewayUsername) > 0 {
		input.Username = aws.String(_backupgatewayUsername)
	}

	if resp, err := client.TestHypervisorConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from the resource.
func backupgateway_UntagResource(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_backupgatewayResourceARN) > 0 {
		input.ResourceARN = aws.String(_backupgatewayResourceARN)
	}
	if len(_backupgatewayTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _backupgatewayTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a gateway's name. Specify which gateway to update using the Amazon
// Resource Name (ARN) of the gateway in your request.
func backupgateway_UpdateGatewayInformation(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.UpdateGatewayInformationInput{
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}
	if len(_backupgatewayGatewayDisplayName) > 0 {
		input.GatewayDisplayName = aws.String(_backupgatewayGatewayDisplayName)
	}

	if resp, err := client.UpdateGatewayInformation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the gateway virtual machine (VM) software. The request immediately
// triggers the software update.
//
// When you make this request, you get a 200 OK success response immediately.
// However, it might take some time for the update to complete.
func backupgateway_UpdateGatewaySoftwareNow(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.UpdateGatewaySoftwareNowInput{
		// GatewayArn: *string, // Required
	}

	if len(_backupgatewayGatewayArn) > 0 {
		input.GatewayArn = aws.String(_backupgatewayGatewayArn)
	}

	if resp, err := client.UpdateGatewaySoftwareNow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a hypervisor metadata, including its host, username, and password.
// Specify which hypervisor to update using the Amazon Resource Name (ARN) of the
// hypervisor in your request.
func backupgateway_UpdateHypervisor(cfg aws.Config, client *backupgateway.Client) {
	input := &backupgateway.UpdateHypervisorInput{
		// HypervisorArn: *string, // Required
	}

	if len(_backupgatewayHypervisorArn) > 0 {
		input.HypervisorArn = aws.String(_backupgatewayHypervisorArn)
	}
	if len(_backupgatewayHost) > 0 {
		input.Host = aws.String(_backupgatewayHost)
	}
	if len(_backupgatewayLogGroupArn) > 0 {
		input.LogGroupArn = aws.String(_backupgatewayLogGroupArn)
	}
	if len(_backupgatewayName) > 0 {
		input.Name = aws.String(_backupgatewayName)
	}
	if len(_backupgatewayPassword) > 0 {
		input.Password = aws.String(_backupgatewayPassword)
	}
	if len(_backupgatewayUsername) > 0 {
		input.Username = aws.String(_backupgatewayUsername)
	}

	if resp, err := client.UpdateHypervisor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_backupgatewayCmd)
	_backupgatewayCmd.Flags().SortFlags = false

	_backupgatewayCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_backupgatewayCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_backupgatewayCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayActivationKey, "activation-key", "", "", "Activation Key")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayBandwidthRateLimitIntervals, "bandwidth-rate-limit-intervals", "", "", "Bandwidth Rate Limit Intervals")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayDayOfMonth, "day-of-month", "", "", "Day Of Month")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayDayOfWeek, "day-of-week", "", "", "Day Of Week")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayGatewayArn, "gateway-arn", "", "", "Gateway ARN")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayGatewayDisplayName, "gateway-display-name", "", "", "Gateway Display Name")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayGatewayType, "gateway-type", "", "", "Gateway Type")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayHost, "host", "", "", "Host")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayHourOfDay, "hour-of-day", "", "", "Hour Of Day")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayHypervisorArn, "hypervisor-arn", "", "", "Hypervisor ARN")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayIamRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayLogGroupArn, "log-group-arn", "", "", "Log Group ARN")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayMaxResults, "max-results", "", "", "Max Results")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayMinuteOfHour, "minute-of-hour", "", "", "Minute Of Hour")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayName, "name", "", "", "Name")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayNextToken, "next-token", "", "", "Next Token")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayPassword, "password", "", "", "Password")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayResourceARN, "resource-arn", "", "", "Resource ARN")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayServerArn, "server-arn", "", "", "Server ARN")
	_backupgatewayCmd.Flags().StringSliceVarP(&_backupgatewayTagKeys, "tag-keys", "", nil, "Tag Keys")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayTags, "tags", "", "", "Tags")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayUsername, "username", "", "", "Username")
	_backupgatewayCmd.Flags().StringVarP(&_backupgatewayVmwareToAwsTagMappings, "vmware-to-aws-tag-mappings", "", "", "Vmware To AWS Tag Mappings")

	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayAssociateGatewayToServer, "associate-gateway-to-server", "", false, "Associate Gateway To Server")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayCreateGateway, "create-gateway", "", false, "Create Gateway")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayDeleteGateway, "delete-gateway", "", false, "Delete Gateway")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayDeleteHypervisor, "delete-hypervisor", "", false, "Delete Hypervisor")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayDisassociateGatewayFromServer, "disassociate-gateway-from-server", "", false, "Disassociate Gateway From Server")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayGetBandwidthRateLimitSchedule, "get-bandwidth-rate-limit-schedule", "", false, "Get Bandwidth Rate Limit Schedule")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayGetGateway, "get-gateway", "", false, "Get Gateway")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayGetHypervisor, "get-hypervisor", "", false, "Get Hypervisor")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayGetHypervisorPropertyMappings, "get-hypervisor-property-mappings", "", false, "Get Hypervisor Property Mappings")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayGetVirtualMachine, "get-virtual-machine", "", false, "Get Virtual Machine")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayImportHypervisorConfiguration, "import-hypervisor-configuration", "", false, "Import Hypervisor Configuration")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayListGateways, "list-gateways", "", false, "List Gateways")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayListHypervisors, "list-hypervisors", "", false, "List Hypervisors")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayListVirtualMachines, "list-virtual-machines", "", false, "List Virtual Machines")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayPutBandwidthRateLimitSchedule, "put-bandwidth-rate-limit-schedule", "", false, "Put Bandwidth Rate Limit Schedule")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayPutHypervisorPropertyMappings, "put-hypervisor-property-mappings", "", false, "Put Hypervisor Property Mappings")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayPutMaintenanceStartTime, "put-maintenance-start-time", "", false, "Put Maintenance Start Time")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayStartVirtualMachinesMetadataSync, "start-virtual-machines-metadata-sync", "", false, "Start Virtual Machines Metadata Sync")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayTagResource, "tag-resource", "", false, "Tag Resource")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayTestHypervisorConfiguration, "test-hypervisor-configuration", "", false, "Test Hypervisor Configuration")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayUntagResource, "untag-resource", "", false, "Untag Resource")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayUpdateGatewayInformation, "update-gateway-information", "", false, "Update Gateway Information")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayUpdateGatewaySoftwareNow, "update-gateway-software-now", "", false, "Update Gateway Software Now")
	_backupgatewayCmd.Flags().BoolVarP(&_backupgatewayUpdateHypervisor, "update-hypervisor", "", false, "Update Hypervisor")

}
