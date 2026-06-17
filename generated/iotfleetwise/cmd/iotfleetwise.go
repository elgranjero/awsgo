package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotfleetwise"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotfleetwiseCmd represents the iotfleetwise command
var _iotfleetwiseCmd = &cobra.Command{
	Use:   "iotfleetwise",
	Short: "AWS iotfleetwise CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iotfleetwise.NewFromConfig(cfg)
		if _iotfleetwiseAssociateVehicleFleet {
			iotfleetwise_AssociateVehicleFleet(cfg, client)
			return
		}
		if _iotfleetwiseBatchCreateVehicle {
			iotfleetwise_BatchCreateVehicle(cfg, client)
			return
		}
		if _iotfleetwiseBatchUpdateVehicle {
			iotfleetwise_BatchUpdateVehicle(cfg, client)
			return
		}
		if _iotfleetwiseCreateCampaign {
			iotfleetwise_CreateCampaign(cfg, client)
			return
		}
		if _iotfleetwiseCreateDecoderManifest {
			iotfleetwise_CreateDecoderManifest(cfg, client)
			return
		}
		if _iotfleetwiseCreateFleet {
			iotfleetwise_CreateFleet(cfg, client)
			return
		}
		if _iotfleetwiseCreateModelManifest {
			iotfleetwise_CreateModelManifest(cfg, client)
			return
		}
		if _iotfleetwiseCreateSignalCatalog {
			iotfleetwise_CreateSignalCatalog(cfg, client)
			return
		}
		if _iotfleetwiseCreateStateTemplate {
			iotfleetwise_CreateStateTemplate(cfg, client)
			return
		}
		if _iotfleetwiseCreateVehicle {
			iotfleetwise_CreateVehicle(cfg, client)
			return
		}
		if _iotfleetwiseDeleteCampaign {
			iotfleetwise_DeleteCampaign(cfg, client)
			return
		}
		if _iotfleetwiseDeleteDecoderManifest {
			iotfleetwise_DeleteDecoderManifest(cfg, client)
			return
		}
		if _iotfleetwiseDeleteFleet {
			iotfleetwise_DeleteFleet(cfg, client)
			return
		}
		if _iotfleetwiseDeleteModelManifest {
			iotfleetwise_DeleteModelManifest(cfg, client)
			return
		}
		if _iotfleetwiseDeleteSignalCatalog {
			iotfleetwise_DeleteSignalCatalog(cfg, client)
			return
		}
		if _iotfleetwiseDeleteStateTemplate {
			iotfleetwise_DeleteStateTemplate(cfg, client)
			return
		}
		if _iotfleetwiseDeleteVehicle {
			iotfleetwise_DeleteVehicle(cfg, client)
			return
		}
		if _iotfleetwiseDisassociateVehicleFleet {
			iotfleetwise_DisassociateVehicleFleet(cfg, client)
			return
		}
		if _iotfleetwiseGetCampaign {
			iotfleetwise_GetCampaign(cfg, client)
			return
		}
		if _iotfleetwiseGetDecoderManifest {
			iotfleetwise_GetDecoderManifest(cfg, client)
			return
		}
		if _iotfleetwiseGetEncryptionConfiguration {
			iotfleetwise_GetEncryptionConfiguration(cfg, client)
			return
		}
		if _iotfleetwiseGetFleet {
			iotfleetwise_GetFleet(cfg, client)
			return
		}
		if _iotfleetwiseGetLoggingOptions {
			iotfleetwise_GetLoggingOptions(cfg, client)
			return
		}
		if _iotfleetwiseGetModelManifest {
			iotfleetwise_GetModelManifest(cfg, client)
			return
		}
		if _iotfleetwiseGetRegisterAccountStatus {
			iotfleetwise_GetRegisterAccountStatus(cfg, client)
			return
		}
		if _iotfleetwiseGetSignalCatalog {
			iotfleetwise_GetSignalCatalog(cfg, client)
			return
		}
		if _iotfleetwiseGetStateTemplate {
			iotfleetwise_GetStateTemplate(cfg, client)
			return
		}
		if _iotfleetwiseGetVehicle {
			iotfleetwise_GetVehicle(cfg, client)
			return
		}
		if _iotfleetwiseGetVehicleStatus {
			iotfleetwise_GetVehicleStatus(cfg, client)
			return
		}
		if _iotfleetwiseImportDecoderManifest {
			iotfleetwise_ImportDecoderManifest(cfg, client)
			return
		}
		if _iotfleetwiseImportSignalCatalog {
			iotfleetwise_ImportSignalCatalog(cfg, client)
			return
		}
		if _iotfleetwiseListCampaigns {
			iotfleetwise_ListCampaigns(cfg, client)
			return
		}
		if _iotfleetwiseListDecoderManifestNetworkInterfaces {
			iotfleetwise_ListDecoderManifestNetworkInterfaces(cfg, client)
			return
		}
		if _iotfleetwiseListDecoderManifestSignals {
			iotfleetwise_ListDecoderManifestSignals(cfg, client)
			return
		}
		if _iotfleetwiseListDecoderManifests {
			iotfleetwise_ListDecoderManifests(cfg, client)
			return
		}
		if _iotfleetwiseListFleets {
			iotfleetwise_ListFleets(cfg, client)
			return
		}
		if _iotfleetwiseListFleetsForVehicle {
			iotfleetwise_ListFleetsForVehicle(cfg, client)
			return
		}
		if _iotfleetwiseListModelManifestNodes {
			iotfleetwise_ListModelManifestNodes(cfg, client)
			return
		}
		if _iotfleetwiseListModelManifests {
			iotfleetwise_ListModelManifests(cfg, client)
			return
		}
		if _iotfleetwiseListSignalCatalogNodes {
			iotfleetwise_ListSignalCatalogNodes(cfg, client)
			return
		}
		if _iotfleetwiseListSignalCatalogs {
			iotfleetwise_ListSignalCatalogs(cfg, client)
			return
		}
		if _iotfleetwiseListStateTemplates {
			iotfleetwise_ListStateTemplates(cfg, client)
			return
		}
		if _iotfleetwiseListTagsForResource {
			iotfleetwise_ListTagsForResource(cfg, client)
			return
		}
		if _iotfleetwiseListVehicles {
			iotfleetwise_ListVehicles(cfg, client)
			return
		}
		if _iotfleetwiseListVehiclesInFleet {
			iotfleetwise_ListVehiclesInFleet(cfg, client)
			return
		}
		if _iotfleetwisePutEncryptionConfiguration {
			iotfleetwise_PutEncryptionConfiguration(cfg, client)
			return
		}
		if _iotfleetwisePutLoggingOptions {
			iotfleetwise_PutLoggingOptions(cfg, client)
			return
		}
		if _iotfleetwiseRegisterAccount {
			iotfleetwise_RegisterAccount(cfg, client)
			return
		}
		if _iotfleetwiseTagResource {
			iotfleetwise_TagResource(cfg, client)
			return
		}
		if _iotfleetwiseUntagResource {
			iotfleetwise_UntagResource(cfg, client)
			return
		}
		if _iotfleetwiseUpdateCampaign {
			iotfleetwise_UpdateCampaign(cfg, client)
			return
		}
		if _iotfleetwiseUpdateDecoderManifest {
			iotfleetwise_UpdateDecoderManifest(cfg, client)
			return
		}
		if _iotfleetwiseUpdateFleet {
			iotfleetwise_UpdateFleet(cfg, client)
			return
		}
		if _iotfleetwiseUpdateModelManifest {
			iotfleetwise_UpdateModelManifest(cfg, client)
			return
		}
		if _iotfleetwiseUpdateSignalCatalog {
			iotfleetwise_UpdateSignalCatalog(cfg, client)
			return
		}
		if _iotfleetwiseUpdateStateTemplate {
			iotfleetwise_UpdateStateTemplate(cfg, client)
			return
		}
		if _iotfleetwiseUpdateVehicle {
			iotfleetwise_UpdateVehicle(cfg, client)
			return
		}

	},
}

var (
	_iotfleetwiseAssociateVehicleFleet                bool
	_iotfleetwiseBatchCreateVehicle                   bool
	_iotfleetwiseBatchUpdateVehicle                   bool
	_iotfleetwiseCreateCampaign                       bool
	_iotfleetwiseCreateDecoderManifest                bool
	_iotfleetwiseCreateFleet                          bool
	_iotfleetwiseCreateModelManifest                  bool
	_iotfleetwiseCreateSignalCatalog                  bool
	_iotfleetwiseCreateStateTemplate                  bool
	_iotfleetwiseCreateVehicle                        bool
	_iotfleetwiseDeleteCampaign                       bool
	_iotfleetwiseDeleteDecoderManifest                bool
	_iotfleetwiseDeleteFleet                          bool
	_iotfleetwiseDeleteModelManifest                  bool
	_iotfleetwiseDeleteSignalCatalog                  bool
	_iotfleetwiseDeleteStateTemplate                  bool
	_iotfleetwiseDeleteVehicle                        bool
	_iotfleetwiseDisassociateVehicleFleet             bool
	_iotfleetwiseGetCampaign                          bool
	_iotfleetwiseGetDecoderManifest                   bool
	_iotfleetwiseGetEncryptionConfiguration           bool
	_iotfleetwiseGetFleet                             bool
	_iotfleetwiseGetLoggingOptions                    bool
	_iotfleetwiseGetModelManifest                     bool
	_iotfleetwiseGetRegisterAccountStatus             bool
	_iotfleetwiseGetSignalCatalog                     bool
	_iotfleetwiseGetStateTemplate                     bool
	_iotfleetwiseGetVehicle                           bool
	_iotfleetwiseGetVehicleStatus                     bool
	_iotfleetwiseImportDecoderManifest                bool
	_iotfleetwiseImportSignalCatalog                  bool
	_iotfleetwiseListCampaigns                        bool
	_iotfleetwiseListDecoderManifestNetworkInterfaces bool
	_iotfleetwiseListDecoderManifestSignals           bool
	_iotfleetwiseListDecoderManifests                 bool
	_iotfleetwiseListFleets                           bool
	_iotfleetwiseListFleetsForVehicle                 bool
	_iotfleetwiseListModelManifestNodes               bool
	_iotfleetwiseListModelManifests                   bool
	_iotfleetwiseListSignalCatalogNodes               bool
	_iotfleetwiseListSignalCatalogs                   bool
	_iotfleetwiseListStateTemplates                   bool
	_iotfleetwiseListTagsForResource                  bool
	_iotfleetwiseListVehicles                         bool
	_iotfleetwiseListVehiclesInFleet                  bool
	_iotfleetwisePutEncryptionConfiguration           bool
	_iotfleetwisePutLoggingOptions                    bool
	_iotfleetwiseRegisterAccount                      bool
	_iotfleetwiseTagResource                          bool
	_iotfleetwiseUntagResource                        bool
	_iotfleetwiseUpdateCampaign                       bool
	_iotfleetwiseUpdateDecoderManifest                bool
	_iotfleetwiseUpdateFleet                          bool
	_iotfleetwiseUpdateModelManifest                  bool
	_iotfleetwiseUpdateSignalCatalog                  bool
	_iotfleetwiseUpdateStateTemplate                  bool
	_iotfleetwiseUpdateVehicle                        bool

	_iotfleetwiseAction                          string
	_iotfleetwiseAssociationBehavior             string
	_iotfleetwiseAttributeNames                  []string
	_iotfleetwiseAttributeUpdateMode             string
	_iotfleetwiseAttributeValues                 []string
	_iotfleetwiseAttributes                      string
	_iotfleetwiseCloudWatchLogDelivery           string
	_iotfleetwiseCollectionScheme                string
	_iotfleetwiseCompression                     string
	_iotfleetwiseDataDestinationConfigs          string
	_iotfleetwiseDataExtraDimensions             []string
	_iotfleetwiseDataPartitions                  string
	_iotfleetwiseDecoderManifestArn              string
	_iotfleetwiseDefaultForUnmappedSignals       string
	_iotfleetwiseDescription                     string
	_iotfleetwiseDiagnosticsMode                 string
	_iotfleetwiseEncryptionType                  string
	_iotfleetwiseExpiryTime                      string
	_iotfleetwiseFleetId                         string
	_iotfleetwiseIamResources                    string
	_iotfleetwiseIdentifier                      string
	_iotfleetwiseKmsKeyId                        string
	_iotfleetwiseListResponseScope               string
	_iotfleetwiseMaxResults                      string
	_iotfleetwiseMetadataExtraDimensions         []string
	_iotfleetwiseModelManifestArn                string
	_iotfleetwiseName                            string
	_iotfleetwiseNetworkFileDefinitions          string
	_iotfleetwiseNetworkInterfaces               string
	_iotfleetwiseNetworkInterfacesToAdd          string
	_iotfleetwiseNetworkInterfacesToRemove       []string
	_iotfleetwiseNetworkInterfacesToUpdate       string
	_iotfleetwiseNextToken                       string
	_iotfleetwiseNodes                           string
	_iotfleetwiseNodesToAdd                      string
	_iotfleetwiseNodesToRemove                   []string
	_iotfleetwiseNodesToUpdate                   string
	_iotfleetwisePostTriggerCollectionDuration   string
	_iotfleetwisePriority                        string
	_iotfleetwiseResourceARN                     string
	_iotfleetwiseSignalCatalogArn                string
	_iotfleetwiseSignalDecoders                  string
	_iotfleetwiseSignalDecodersToAdd             string
	_iotfleetwiseSignalDecodersToRemove          []string
	_iotfleetwiseSignalDecodersToUpdate          string
	_iotfleetwiseSignalNodeType                  string
	_iotfleetwiseSignalsToCollect                string
	_iotfleetwiseSignalsToFetch                  string
	_iotfleetwiseSpoolingMode                    string
	_iotfleetwiseStartTime                       string
	_iotfleetwiseStateTemplateProperties         []string
	_iotfleetwiseStateTemplatePropertiesToAdd    []string
	_iotfleetwiseStateTemplatePropertiesToRemove []string
	_iotfleetwiseStateTemplates                  string
	_iotfleetwiseStateTemplatesToAdd             string
	_iotfleetwiseStateTemplatesToRemove          []string
	_iotfleetwiseStateTemplatesToUpdate          string
	_iotfleetwiseStatus                          string
	_iotfleetwiseTagKeys                         []string
	_iotfleetwiseTags                            string
	_iotfleetwiseTargetArn                       string
	_iotfleetwiseTimestreamResources             string
	_iotfleetwiseVehicleName                     string
	_iotfleetwiseVehicles                        string
	_iotfleetwiseVss                             string
)

// Adds, or associates, a vehicle with a fleet.
func iotfleetwise_AssociateVehicleFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.AssociateVehicleFleetInput{
		// FleetId: *string, // Required
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}
	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}

	if resp, err := client.AssociateVehicleFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a group, or batch, of vehicles.
// You must specify a decoder manifest and a vehicle model (model manifest) for
// each vehicle.
//
// For more information, see [Create multiple vehicles (AWS CLI)] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// [Create multiple vehicles (AWS CLI)]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/create-vehicles-cli.html
func iotfleetwise_BatchCreateVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.BatchCreateVehicleInput{
		// Vehicles: []types.CreateVehicleRequestItem, // Required
	}

	if len(_iotfleetwiseVehicles) > 0 {
		if err := assignInputField(input, "Vehicles", _iotfleetwiseVehicles); err != nil {
			log.Errorf("invalid --vehicles: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchCreateVehicle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a group, or batch, of vehicles.
// You must specify a decoder manifest and a vehicle model (model manifest) for
// each vehicle.
//
// For more information, see [Update multiple vehicles (AWS CLI)] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// [Update multiple vehicles (AWS CLI)]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/update-vehicles-cli.html
func iotfleetwise_BatchUpdateVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.BatchUpdateVehicleInput{
		// Vehicles: []types.UpdateVehicleRequestItem, // Required
	}

	if len(_iotfleetwiseVehicles) > 0 {
		if err := assignInputField(input, "Vehicles", _iotfleetwiseVehicles); err != nil {
			log.Errorf("invalid --vehicles: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateVehicle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an orchestration of data collection rules. The Amazon Web Services IoT
// FleetWise Edge Agent software running in vehicles uses campaigns to decide how
// to collect and transfer data to the cloud. You create campaigns in the cloud.
// After you or your team approve campaigns, Amazon Web Services IoT FleetWise
// automatically deploys them to vehicles.
//
// For more information, see [Collect and transfer data with campaigns] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
// [Collect and transfer data with campaigns]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/campaigns.html
func iotfleetwise_CreateCampaign(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateCampaignInput{
		// CollectionScheme: types.CollectionScheme, // Required
		// Name: *string, // Required
		// SignalCatalogArn: *string, // Required
		// TargetArn: *string, // Required
	}

	if len(_iotfleetwiseCollectionScheme) > 0 {
		if err := assignInputField(input, "CollectionScheme", _iotfleetwiseCollectionScheme); err != nil {
			log.Errorf("invalid --collection-scheme: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseSignalCatalogArn) > 0 {
		input.SignalCatalogArn = aws.String(_iotfleetwiseSignalCatalogArn)
	}
	if len(_iotfleetwiseTargetArn) > 0 {
		input.TargetArn = aws.String(_iotfleetwiseTargetArn)
	}
	if len(_iotfleetwiseCompression) > 0 {
		if err := assignInputField(input, "Compression", _iotfleetwiseCompression); err != nil {
			log.Errorf("invalid --compression: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseDataDestinationConfigs) > 0 {
		if err := assignInputField(input, "DataDestinationConfigs", _iotfleetwiseDataDestinationConfigs); err != nil {
			log.Errorf("invalid --data-destination-configs: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseDataExtraDimensions) > 0 {
		input.DataExtraDimensions = append([]string(nil), _iotfleetwiseDataExtraDimensions...)
	}
	if len(_iotfleetwiseDataPartitions) > 0 {
		if err := assignInputField(input, "DataPartitions", _iotfleetwiseDataPartitions); err != nil {
			log.Errorf("invalid --data-partitions: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseDiagnosticsMode) > 0 {
		if err := assignInputField(input, "DiagnosticsMode", _iotfleetwiseDiagnosticsMode); err != nil {
			log.Errorf("invalid --diagnostics-mode: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseExpiryTime) > 0 {
		if err := assignInputField(input, "ExpiryTime", _iotfleetwiseExpiryTime); err != nil {
			log.Errorf("invalid --expiry-time: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwisePostTriggerCollectionDuration) > 0 {
		if err := assignInputField(input, "PostTriggerCollectionDuration", _iotfleetwisePostTriggerCollectionDuration); err != nil {
			log.Errorf("invalid --post-trigger-collection-duration: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwisePriority) > 0 {
		if err := assignInputField(input, "Priority", _iotfleetwisePriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseSignalsToCollect) > 0 {
		if err := assignInputField(input, "SignalsToCollect", _iotfleetwiseSignalsToCollect); err != nil {
			log.Errorf("invalid --signals-to-collect: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseSignalsToFetch) > 0 {
		if err := assignInputField(input, "SignalsToFetch", _iotfleetwiseSignalsToFetch); err != nil {
			log.Errorf("invalid --signals-to-fetch: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseSpoolingMode) > 0 {
		if err := assignInputField(input, "SpoolingMode", _iotfleetwiseSpoolingMode); err != nil {
			log.Errorf("invalid --spooling-mode: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotfleetwiseStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the decoder manifest associated with a model manifest. To create a
// decoder manifest, the following must be true:
//
// - Every signal decoder has a unique name.
//
// - Each signal decoder is associated with a network interface.
//
// - Each network interface has a unique ID.
//
// - The signal decoders are specified in the model manifest.
func iotfleetwise_CreateDecoderManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateDecoderManifestInput{
		// ModelManifestArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_iotfleetwiseModelManifestArn) > 0 {
		input.ModelManifestArn = aws.String(_iotfleetwiseModelManifestArn)
	}
	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDefaultForUnmappedSignals) > 0 {
		if err := assignInputField(input, "DefaultForUnmappedSignals", _iotfleetwiseDefaultForUnmappedSignals); err != nil {
			log.Errorf("invalid --default-for-unmapped-signals: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseNetworkInterfaces) > 0 {
		if err := assignInputField(input, "NetworkInterfaces", _iotfleetwiseNetworkInterfaces); err != nil {
			log.Errorf("invalid --network-interfaces: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseSignalDecoders) > 0 {
		if err := assignInputField(input, "SignalDecoders", _iotfleetwiseSignalDecoders); err != nil {
			log.Errorf("invalid --signal-decoders: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDecoderManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a fleet that represents a group of vehicles.
// You must create both a signal catalog and vehicles before you can create a
// fleet.
//
// For more information, see [Fleets] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// [Fleets]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleets.html
func iotfleetwise_CreateFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateFleetInput{
		// FleetId: *string, // Required
		// SignalCatalogArn: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}
	if len(_iotfleetwiseSignalCatalogArn) > 0 {
		input.SignalCatalogArn = aws.String(_iotfleetwiseSignalCatalogArn)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a vehicle model (model manifest) that specifies signals (attributes,
// branches, sensors, and actuators).
//
// For more information, see [Vehicle models] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// [Vehicle models]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/vehicle-models.html
func iotfleetwise_CreateModelManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateModelManifestInput{
		// Name: *string, // Required
		// Nodes: []string, // Required
		// SignalCatalogArn: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseNodes) > 0 {
		input.Nodes = []string{_iotfleetwiseNodes}
	}
	if len(_iotfleetwiseSignalCatalogArn) > 0 {
		input.SignalCatalogArn = aws.String(_iotfleetwiseSignalCatalogArn)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateModelManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a collection of standardized signals that can be reused to create
// vehicle models.
func iotfleetwise_CreateSignalCatalog(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateSignalCatalogInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseNodes) > 0 {
		if err := assignInputField(input, "Nodes", _iotfleetwiseNodes); err != nil {
			log.Errorf("invalid --nodes: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSignalCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a state template. State templates contain state properties, which are
// signals that belong to a signal catalog that is synchronized between the Amazon
// Web Services IoT FleetWise Edge and the Amazon Web Services Cloud.
//
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
func iotfleetwise_CreateStateTemplate(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateStateTemplateInput{
		// Name: *string, // Required
		// SignalCatalogArn: *string, // Required
		// StateTemplateProperties: []string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseSignalCatalogArn) > 0 {
		input.SignalCatalogArn = aws.String(_iotfleetwiseSignalCatalogArn)
	}
	if len(_iotfleetwiseStateTemplateProperties) > 0 {
		input.StateTemplateProperties = append([]string(nil), _iotfleetwiseStateTemplateProperties...)
	}
	if len(_iotfleetwiseDataExtraDimensions) > 0 {
		input.DataExtraDimensions = append([]string(nil), _iotfleetwiseDataExtraDimensions...)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseMetadataExtraDimensions) > 0 {
		input.MetadataExtraDimensions = append([]string(nil), _iotfleetwiseMetadataExtraDimensions...)
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a vehicle, which is an instance of a vehicle model (model manifest).
// Vehicles created from the same vehicle model consist of the same signals
// inherited from the vehicle model.
//
// If you have an existing Amazon Web Services IoT thing, you can use Amazon Web
// Services IoT FleetWise to create a vehicle and collect data from your thing.
//
// For more information, see [Create a vehicle (AWS CLI)] in the Amazon Web Services IoT FleetWise Developer
// Guide.
//
// [Create a vehicle (AWS CLI)]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/create-vehicle-cli.html
func iotfleetwise_CreateVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.CreateVehicleInput{
		// DecoderManifestArn: *string, // Required
		// ModelManifestArn: *string, // Required
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseDecoderManifestArn) > 0 {
		input.DecoderManifestArn = aws.String(_iotfleetwiseDecoderManifestArn)
	}
	if len(_iotfleetwiseModelManifestArn) > 0 {
		input.ModelManifestArn = aws.String(_iotfleetwiseModelManifestArn)
	}
	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}
	if len(_iotfleetwiseAssociationBehavior) > 0 {
		if err := assignInputField(input, "AssociationBehavior", _iotfleetwiseAssociationBehavior); err != nil {
			log.Errorf("invalid --association-behavior: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _iotfleetwiseAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseStateTemplates) > 0 {
		if err := assignInputField(input, "StateTemplates", _iotfleetwiseStateTemplates); err != nil {
			log.Errorf("invalid --state-templates: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVehicle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data collection campaign. Deleting a campaign suspends all data
// collection and removes it from any vehicles.
func iotfleetwise_DeleteCampaign(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteCampaignInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.DeleteCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a decoder manifest. You can't delete a decoder manifest if it has
// vehicles associated with it.
func iotfleetwise_DeleteDecoderManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteDecoderManifestInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.DeleteDecoderManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a fleet. Before you delete a fleet, all vehicles must be dissociated
// from the fleet. For more information, see [Delete a fleet (AWS CLI)]in the Amazon Web Services IoT
// FleetWise Developer Guide.
//
// [Delete a fleet (AWS CLI)]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/delete-fleet-cli.html
func iotfleetwise_DeleteFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteFleetInput{
		// FleetId: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}

	if resp, err := client.DeleteFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a vehicle model (model manifest).
func iotfleetwise_DeleteModelManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteModelManifestInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.DeleteModelManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a signal catalog.
func iotfleetwise_DeleteSignalCatalog(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteSignalCatalogInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.DeleteSignalCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a state template.
func iotfleetwise_DeleteStateTemplate(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteStateTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_iotfleetwiseIdentifier) > 0 {
		input.Identifier = aws.String(_iotfleetwiseIdentifier)
	}

	if resp, err := client.DeleteStateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a vehicle and removes it from any campaigns.
func iotfleetwise_DeleteVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DeleteVehicleInput{
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}

	if resp, err := client.DeleteVehicle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes, or disassociates, a vehicle from a fleet. Disassociating a vehicle
// from a fleet doesn't delete the vehicle.
func iotfleetwise_DisassociateVehicleFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.DisassociateVehicleFleetInput{
		// FleetId: *string, // Required
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}
	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}

	if resp, err := client.DisassociateVehicleFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a campaign.
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
func iotfleetwise_GetCampaign(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetCampaignInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.GetCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a created decoder manifest.
func iotfleetwise_GetDecoderManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetDecoderManifestInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.GetDecoderManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the encryption configuration for resources and data in Amazon Web
// Services IoT FleetWise.
func iotfleetwise_GetEncryptionConfiguration(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetEncryptionConfigurationInput{}

	if resp, err := client.GetEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a fleet.
func iotfleetwise_GetFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetFleetInput{
		// FleetId: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}

	if resp, err := client.GetFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the logging options.
func iotfleetwise_GetLoggingOptions(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetLoggingOptionsInput{}

	if resp, err := client.GetLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a vehicle model (model manifest).
func iotfleetwise_GetModelManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetModelManifestInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.GetModelManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status of registering your Amazon Web Services
// account, IAM, and Amazon Timestream resources so that Amazon Web Services IoT
// FleetWise can transfer your vehicle data to the Amazon Web Services Cloud.
//
// For more information, including step-by-step procedures, see [Setting up Amazon Web Services IoT FleetWise].
//
// This API operation doesn't require input parameters.
//
// [Setting up Amazon Web Services IoT FleetWise]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/setting-up.html
func iotfleetwise_GetRegisterAccountStatus(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetRegisterAccountStatusInput{}

	if resp, err := client.GetRegisterAccountStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a signal catalog.
func iotfleetwise_GetSignalCatalog(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetSignalCatalogInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}

	if resp, err := client.GetSignalCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a state template.
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
func iotfleetwise_GetStateTemplate(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetStateTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_iotfleetwiseIdentifier) > 0 {
		input.Identifier = aws.String(_iotfleetwiseIdentifier)
	}

	if resp, err := client.GetStateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a vehicle.
func iotfleetwise_GetVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetVehicleInput{
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}

	if resp, err := client.GetVehicle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the status of campaigns, decoder manifests, or
// state templates associated with a vehicle.
func iotfleetwise_GetVehicleStatus(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.GetVehicleStatusInput{
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetVehicleStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.GetVehicleStatusOutput
	p := iotfleetwise.NewGetVehicleStatusPaginator(client, input)
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

// Creates a decoder manifest using your existing CAN DBC file from your local
// device.
//
// The CAN signal name must be unique and not repeated across CAN message
// definitions in a .dbc file.
func iotfleetwise_ImportDecoderManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ImportDecoderManifestInput{
		// Name: *string, // Required
		// NetworkFileDefinitions: []types.NetworkFileDefinition, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseNetworkFileDefinitions) > 0 {
		if err := assignInputField(input, "NetworkFileDefinitions", _iotfleetwiseNetworkFileDefinitions); err != nil {
			log.Errorf("invalid --network-file-definitions: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportDecoderManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a signal catalog using your existing VSS formatted content from your
// local device.
func iotfleetwise_ImportSignalCatalog(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ImportSignalCatalogInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseVss) > 0 {
		if err := assignInputField(input, "Vss", _iotfleetwiseVss); err != nil {
			log.Errorf("invalid --vss: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportSignalCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists information about created campaigns.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListCampaigns(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListCampaignsInput{}

	if len(_iotfleetwiseListResponseScope) > 0 {
		if err := assignInputField(input, "ListResponseScope", _iotfleetwiseListResponseScope); err != nil {
			log.Errorf("invalid --list-response-scope: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}
	if len(_iotfleetwiseStatus) > 0 {
		input.Status = aws.String(_iotfleetwiseStatus)
	}

	if disablePaginator() {
		if resp, err := client.ListCampaigns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListCampaignsOutput
	p := iotfleetwise.NewListCampaignsPaginator(client, input)
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

// Lists the network interfaces specified in a decoder manifest.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListDecoderManifestNetworkInterfaces(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListDecoderManifestNetworkInterfacesInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDecoderManifestNetworkInterfaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput
	p := iotfleetwise.NewListDecoderManifestNetworkInterfacesPaginator(client, input)
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

// A list of information about signal decoders specified in a decoder manifest.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListDecoderManifestSignals(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListDecoderManifestSignalsInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDecoderManifestSignals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListDecoderManifestSignalsOutput
	p := iotfleetwise.NewListDecoderManifestSignalsPaginator(client, input)
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

// Lists decoder manifests.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListDecoderManifests(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListDecoderManifestsInput{}

	if len(_iotfleetwiseListResponseScope) > 0 {
		if err := assignInputField(input, "ListResponseScope", _iotfleetwiseListResponseScope); err != nil {
			log.Errorf("invalid --list-response-scope: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseModelManifestArn) > 0 {
		input.ModelManifestArn = aws.String(_iotfleetwiseModelManifestArn)
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDecoderManifests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListDecoderManifestsOutput
	p := iotfleetwise.NewListDecoderManifestsPaginator(client, input)
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

// Retrieves information for each created fleet in an Amazon Web Services
// account.
//
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListFleets(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListFleetsInput{}

	if len(_iotfleetwiseListResponseScope) > 0 {
		if err := assignInputField(input, "ListResponseScope", _iotfleetwiseListResponseScope); err != nil {
			log.Errorf("invalid --list-response-scope: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFleets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListFleetsOutput
	p := iotfleetwise.NewListFleetsPaginator(client, input)
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

// Retrieves a list of IDs for all fleets that the vehicle is associated with.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListFleetsForVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListFleetsForVehicleInput{
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFleetsForVehicle(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListFleetsForVehicleOutput
	p := iotfleetwise.NewListFleetsForVehiclePaginator(client, input)
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

// Lists information about nodes specified in a vehicle model (model manifest).
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListModelManifestNodes(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListModelManifestNodesInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListModelManifestNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListModelManifestNodesOutput
	p := iotfleetwise.NewListModelManifestNodesPaginator(client, input)
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

// Retrieves a list of vehicle models (model manifests).
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListModelManifests(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListModelManifestsInput{}

	if len(_iotfleetwiseListResponseScope) > 0 {
		if err := assignInputField(input, "ListResponseScope", _iotfleetwiseListResponseScope); err != nil {
			log.Errorf("invalid --list-response-scope: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}
	if len(_iotfleetwiseSignalCatalogArn) > 0 {
		input.SignalCatalogArn = aws.String(_iotfleetwiseSignalCatalogArn)
	}

	if disablePaginator() {
		if resp, err := client.ListModelManifests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListModelManifestsOutput
	p := iotfleetwise.NewListModelManifestsPaginator(client, input)
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

// Lists of information about the signals (nodes) specified in a signal catalog.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListSignalCatalogNodes(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListSignalCatalogNodesInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}
	if len(_iotfleetwiseSignalNodeType) > 0 {
		if err := assignInputField(input, "SignalNodeType", _iotfleetwiseSignalNodeType); err != nil {
			log.Errorf("invalid --signal-node-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSignalCatalogNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListSignalCatalogNodesOutput
	p := iotfleetwise.NewListSignalCatalogNodesPaginator(client, input)
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

// Lists all the created signal catalogs in an Amazon Web Services account.
// You can use to list information about each signal (node) specified in a signal
// catalog.
//
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListSignalCatalogs(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListSignalCatalogsInput{}

	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSignalCatalogs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListSignalCatalogsOutput
	p := iotfleetwise.NewListSignalCatalogsPaginator(client, input)
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

// Lists information about created state templates.
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
func iotfleetwise_ListStateTemplates(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListStateTemplatesInput{}

	if len(_iotfleetwiseListResponseScope) > 0 {
		if err := assignInputField(input, "ListResponseScope", _iotfleetwiseListResponseScope); err != nil {
			log.Errorf("invalid --list-response-scope: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStateTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListStateTemplatesOutput
	p := iotfleetwise.NewListStateTemplatesPaginator(client, input)
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

// Lists the tags (metadata) you have assigned to the resource.
func iotfleetwise_ListTagsForResource(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_iotfleetwiseResourceARN) > 0 {
		input.ResourceARN = aws.String(_iotfleetwiseResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of summaries of created vehicles.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListVehicles(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListVehiclesInput{}

	if len(_iotfleetwiseAttributeNames) > 0 {
		input.AttributeNames = append([]string(nil), _iotfleetwiseAttributeNames...)
	}
	if len(_iotfleetwiseAttributeValues) > 0 {
		input.AttributeValues = append([]string(nil), _iotfleetwiseAttributeValues...)
	}
	if len(_iotfleetwiseListResponseScope) > 0 {
		if err := assignInputField(input, "ListResponseScope", _iotfleetwiseListResponseScope); err != nil {
			log.Errorf("invalid --list-response-scope: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseModelManifestArn) > 0 {
		input.ModelManifestArn = aws.String(_iotfleetwiseModelManifestArn)
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVehicles(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListVehiclesOutput
	p := iotfleetwise.NewListVehiclesPaginator(client, input)
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

// Retrieves a list of summaries of all vehicles associated with a fleet.
// This API operation uses pagination. Specify the nextToken parameter in the
// request to return more results.
func iotfleetwise_ListVehiclesInFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.ListVehiclesInFleetInput{
		// FleetId: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}
	if len(_iotfleetwiseMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotfleetwiseMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNextToken) > 0 {
		input.NextToken = aws.String(_iotfleetwiseNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVehiclesInFleet(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotfleetwise.ListVehiclesInFleetOutput
	p := iotfleetwise.NewListVehiclesInFleetPaginator(client, input)
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

// Creates or updates the encryption configuration. Amazon Web Services IoT
// FleetWise can encrypt your data and resources using an Amazon Web Services
// managed key. Or, you can use a KMS key that you own and manage. For more
// information, see [Data encryption]in the Amazon Web Services IoT FleetWise Developer Guide.
//
// [Data encryption]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/data-encryption.html
func iotfleetwise_PutEncryptionConfiguration(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.PutEncryptionConfigurationInput{
		// EncryptionType: types.EncryptionType, // Required
	}

	if len(_iotfleetwiseEncryptionType) > 0 {
		if err := assignInputField(input, "EncryptionType", _iotfleetwiseEncryptionType); err != nil {
			log.Errorf("invalid --encryption-type: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_iotfleetwiseKmsKeyId)
	}

	if resp, err := client.PutEncryptionConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the logging option.
func iotfleetwise_PutLoggingOptions(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.PutLoggingOptionsInput{
		// CloudWatchLogDelivery: *types.CloudWatchLogDeliveryOptions, // Required
	}

	if len(_iotfleetwiseCloudWatchLogDelivery) > 0 {
		if err := assignInputField(input, "CloudWatchLogDelivery", _iotfleetwiseCloudWatchLogDelivery); err != nil {
			log.Errorf("invalid --cloud-watch-log-delivery: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLoggingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API operation contains deprecated parameters. Register your account again
// without the Timestream resources parameter so that Amazon Web Services IoT
// FleetWise can remove the Timestream metadata stored. You should then pass the
// data destination into the [CreateCampaign]API operation.
//
// You must delete any existing campaigns that include an empty data destination
// before you register your account again. For more information, see the [DeleteCampaign]API
// operation.
//
// If you want to delete the Timestream inline policy from the service-linked
// role, such as to mitigate an overly permissive policy, you must first delete any
// existing campaigns. Then delete the service-linked role and register your
// account again to enable CloudWatch metrics. For more information, see [DeleteServiceLinkedRole]in the
// Identity and Access Management API Reference.
//
// Registers your Amazon Web Services account, IAM, and Amazon Timestream
// resources so Amazon Web Services IoT FleetWise can transfer your vehicle data to
// the Amazon Web Services Cloud. For more information, including step-by-step
// procedures, see [Setting up Amazon Web Services IoT FleetWise].
//
// An Amazon Web Services account is not the same thing as a "user." An [Amazon Web Services user] is an
// identity that you create using Identity and Access Management (IAM) and takes
// the form of either an [IAM user]or an [IAM role, both with credentials]. A single Amazon Web Services account can, and
// typically does, contain many users and roles.
//
// [CreateCampaign]: https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_CreateCampaign.html
// [DeleteServiceLinkedRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceLinkedRole.html
// [Amazon Web Services user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction_identity-management.html#intro-identity-users
// [Setting up Amazon Web Services IoT FleetWise]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/setting-up.html
// [DeleteCampaign]: https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_DeleteCampaign.html
// [IAM user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html
// [IAM role, both with credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
func iotfleetwise_RegisterAccount(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.RegisterAccountInput{}

	if len(_iotfleetwiseIamResources) > 0 {
		if err := assignInputField(input, "IamResources", _iotfleetwiseIamResources); err != nil {
			log.Errorf("invalid --iam-resources: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseTimestreamResources) > 0 {
		if err := assignInputField(input, "TimestreamResources", _iotfleetwiseTimestreamResources); err != nil {
			log.Errorf("invalid --timestream-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds to or modifies the tags of the given resource. Tags are metadata which can
// be used to manage a resource.
func iotfleetwise_TagResource(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iotfleetwiseResourceARN) > 0 {
		input.ResourceARN = aws.String(_iotfleetwiseResourceARN)
	}
	if len(_iotfleetwiseTags) > 0 {
		if err := assignInputField(input, "Tags", _iotfleetwiseTags); err != nil {
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

// Removes the given tags (metadata) from the resource.
func iotfleetwise_UntagResource(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotfleetwiseResourceARN) > 0 {
		input.ResourceARN = aws.String(_iotfleetwiseResourceARN)
	}
	if len(_iotfleetwiseTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotfleetwiseTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a campaign.
func iotfleetwise_UpdateCampaign(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateCampaignInput{
		// Action: types.UpdateCampaignAction, // Required
		// Name: *string, // Required
	}

	if len(_iotfleetwiseAction) > 0 {
		if err := assignInputField(input, "Action", _iotfleetwiseAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDataExtraDimensions) > 0 {
		input.DataExtraDimensions = append([]string(nil), _iotfleetwiseDataExtraDimensions...)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}

	if resp, err := client.UpdateCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a decoder manifest.
// A decoder manifest can only be updated when the status is DRAFT . Only ACTIVE
// decoder manifests can be associated with vehicles.
func iotfleetwise_UpdateDecoderManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateDecoderManifestInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDefaultForUnmappedSignals) > 0 {
		if err := assignInputField(input, "DefaultForUnmappedSignals", _iotfleetwiseDefaultForUnmappedSignals); err != nil {
			log.Errorf("invalid --default-for-unmapped-signals: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseNetworkInterfacesToAdd) > 0 {
		if err := assignInputField(input, "NetworkInterfacesToAdd", _iotfleetwiseNetworkInterfacesToAdd); err != nil {
			log.Errorf("invalid --network-interfaces-to-add: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNetworkInterfacesToRemove) > 0 {
		input.NetworkInterfacesToRemove = append([]string(nil), _iotfleetwiseNetworkInterfacesToRemove...)
	}
	if len(_iotfleetwiseNetworkInterfacesToUpdate) > 0 {
		if err := assignInputField(input, "NetworkInterfacesToUpdate", _iotfleetwiseNetworkInterfacesToUpdate); err != nil {
			log.Errorf("invalid --network-interfaces-to-update: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseSignalDecodersToAdd) > 0 {
		if err := assignInputField(input, "SignalDecodersToAdd", _iotfleetwiseSignalDecodersToAdd); err != nil {
			log.Errorf("invalid --signal-decoders-to-add: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseSignalDecodersToRemove) > 0 {
		input.SignalDecodersToRemove = append([]string(nil), _iotfleetwiseSignalDecodersToRemove...)
	}
	if len(_iotfleetwiseSignalDecodersToUpdate) > 0 {
		if err := assignInputField(input, "SignalDecodersToUpdate", _iotfleetwiseSignalDecodersToUpdate); err != nil {
			log.Errorf("invalid --signal-decoders-to-update: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseStatus) > 0 {
		if err := assignInputField(input, "Status", _iotfleetwiseStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDecoderManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description of an existing fleet.
func iotfleetwise_UpdateFleet(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateFleetInput{
		// FleetId: *string, // Required
	}

	if len(_iotfleetwiseFleetId) > 0 {
		input.FleetId = aws.String(_iotfleetwiseFleetId)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}

	if resp, err := client.UpdateFleet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a vehicle model (model manifest). If created vehicles are associated
// with a vehicle model, it can't be updated.
func iotfleetwise_UpdateModelManifest(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateModelManifestInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseNodesToAdd) > 0 {
		input.NodesToAdd = []string{_iotfleetwiseNodesToAdd}
	}
	if len(_iotfleetwiseNodesToRemove) > 0 {
		input.NodesToRemove = append([]string(nil), _iotfleetwiseNodesToRemove...)
	}
	if len(_iotfleetwiseStatus) > 0 {
		if err := assignInputField(input, "Status", _iotfleetwiseStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateModelManifest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a signal catalog.
func iotfleetwise_UpdateSignalCatalog(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateSignalCatalogInput{
		// Name: *string, // Required
	}

	if len(_iotfleetwiseName) > 0 {
		input.Name = aws.String(_iotfleetwiseName)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseNodesToAdd) > 0 {
		if err := assignInputField(input, "NodesToAdd", _iotfleetwiseNodesToAdd); err != nil {
			log.Errorf("invalid --nodes-to-add: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseNodesToRemove) > 0 {
		input.NodesToRemove = append([]string(nil), _iotfleetwiseNodesToRemove...)
	}
	if len(_iotfleetwiseNodesToUpdate) > 0 {
		if err := assignInputField(input, "NodesToUpdate", _iotfleetwiseNodesToUpdate); err != nil {
			log.Errorf("invalid --nodes-to-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSignalCatalog(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a state template.
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
func iotfleetwise_UpdateStateTemplate(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateStateTemplateInput{
		// Identifier: *string, // Required
	}

	if len(_iotfleetwiseIdentifier) > 0 {
		input.Identifier = aws.String(_iotfleetwiseIdentifier)
	}
	if len(_iotfleetwiseDataExtraDimensions) > 0 {
		input.DataExtraDimensions = append([]string(nil), _iotfleetwiseDataExtraDimensions...)
	}
	if len(_iotfleetwiseDescription) > 0 {
		input.Description = aws.String(_iotfleetwiseDescription)
	}
	if len(_iotfleetwiseMetadataExtraDimensions) > 0 {
		input.MetadataExtraDimensions = append([]string(nil), _iotfleetwiseMetadataExtraDimensions...)
	}
	if len(_iotfleetwiseStateTemplatePropertiesToAdd) > 0 {
		input.StateTemplatePropertiesToAdd = append([]string(nil), _iotfleetwiseStateTemplatePropertiesToAdd...)
	}
	if len(_iotfleetwiseStateTemplatePropertiesToRemove) > 0 {
		input.StateTemplatePropertiesToRemove = append([]string(nil), _iotfleetwiseStateTemplatePropertiesToRemove...)
	}

	if resp, err := client.UpdateStateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a vehicle.
// Access to certain Amazon Web Services IoT FleetWise features is currently
// gated. For more information, see [Amazon Web Services Region and feature availability]in the Amazon Web Services IoT FleetWise
// Developer Guide.
//
// [Amazon Web Services Region and feature availability]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html
func iotfleetwise_UpdateVehicle(cfg aws.Config, client *iotfleetwise.Client) {
	input := &iotfleetwise.UpdateVehicleInput{
		// VehicleName: *string, // Required
	}

	if len(_iotfleetwiseVehicleName) > 0 {
		input.VehicleName = aws.String(_iotfleetwiseVehicleName)
	}
	if len(_iotfleetwiseAttributeUpdateMode) > 0 {
		if err := assignInputField(input, "AttributeUpdateMode", _iotfleetwiseAttributeUpdateMode); err != nil {
			log.Errorf("invalid --attribute-update-mode: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _iotfleetwiseAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseDecoderManifestArn) > 0 {
		input.DecoderManifestArn = aws.String(_iotfleetwiseDecoderManifestArn)
	}
	if len(_iotfleetwiseModelManifestArn) > 0 {
		input.ModelManifestArn = aws.String(_iotfleetwiseModelManifestArn)
	}
	if len(_iotfleetwiseStateTemplatesToAdd) > 0 {
		if err := assignInputField(input, "StateTemplatesToAdd", _iotfleetwiseStateTemplatesToAdd); err != nil {
			log.Errorf("invalid --state-templates-to-add: %s", err.Error())
			return
		}
	}
	if len(_iotfleetwiseStateTemplatesToRemove) > 0 {
		input.StateTemplatesToRemove = append([]string(nil), _iotfleetwiseStateTemplatesToRemove...)
	}
	if len(_iotfleetwiseStateTemplatesToUpdate) > 0 {
		if err := assignInputField(input, "StateTemplatesToUpdate", _iotfleetwiseStateTemplatesToUpdate); err != nil {
			log.Errorf("invalid --state-templates-to-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVehicle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotfleetwiseCmd)
	_iotfleetwiseCmd.Flags().SortFlags = false

	_iotfleetwiseCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_iotfleetwiseCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotfleetwiseCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseAction, "action", "", "", "Action")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseAssociationBehavior, "association-behavior", "", "", "Association Behavior")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseAttributeNames, "attribute-names", "", nil, "Attribute Names")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseAttributeUpdateMode, "attribute-update-mode", "", "", "Attribute Update Mode")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseAttributeValues, "attribute-values", "", nil, "Attribute Values")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseAttributes, "attributes", "", "", "Attributes")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseCloudWatchLogDelivery, "cloud-watch-log-delivery", "", "", "Cloud Watch Log Delivery")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseCollectionScheme, "collection-scheme", "", "", "Collection Scheme")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseCompression, "compression", "", "", "Compression")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseDataDestinationConfigs, "data-destination-configs", "", "", "Data Destination Configs")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseDataExtraDimensions, "data-extra-dimensions", "", nil, "Data Extra Dimensions")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseDataPartitions, "data-partitions", "", "", "Data Partitions")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseDecoderManifestArn, "decoder-manifest-arn", "", "", "Decoder Manifest ARN")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseDefaultForUnmappedSignals, "default-for-unmapped-signals", "", "", "Default For Unmapped Signals")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseDescription, "description", "", "", "Description")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseDiagnosticsMode, "diagnostics-mode", "", "", "Diagnostics Mode")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseEncryptionType, "encryption-type", "", "", "Encryption Type")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseExpiryTime, "expiry-time", "", "", "Expiry Time")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseFleetId, "fleet-id", "", "", "Fleet ID")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseIamResources, "iam-resources", "", "", "IAM Resources")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseIdentifier, "identifier", "", "", "Identifier")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseListResponseScope, "list-response-scope", "", "", "List Response Scope")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseMaxResults, "max-results", "", "", "Max Results")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseMetadataExtraDimensions, "metadata-extra-dimensions", "", nil, "Metadata Extra Dimensions")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseModelManifestArn, "model-manifest-arn", "", "", "Model Manifest ARN")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseName, "name", "", "", "Name")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNetworkFileDefinitions, "network-file-definitions", "", "", "Network File Definitions")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNetworkInterfaces, "network-interfaces", "", "", "Network Interfaces")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNetworkInterfacesToAdd, "network-interfaces-to-add", "", "", "Network Interfaces To Add")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseNetworkInterfacesToRemove, "network-interfaces-to-remove", "", nil, "Network Interfaces To Remove")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNetworkInterfacesToUpdate, "network-interfaces-to-update", "", "", "Network Interfaces To Update")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNextToken, "next-token", "", "", "Next Token")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNodes, "nodes", "", "", "Nodes")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNodesToAdd, "nodes-to-add", "", "", "Nodes To Add")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseNodesToRemove, "nodes-to-remove", "", nil, "Nodes To Remove")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseNodesToUpdate, "nodes-to-update", "", "", "Nodes To Update")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwisePostTriggerCollectionDuration, "post-trigger-collection-duration", "", "", "Post Trigger Collection Duration")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwisePriority, "priority", "", "", "Priority")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseResourceARN, "resource-arn", "", "", "Resource ARN")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalCatalogArn, "signal-catalog-arn", "", "", "Signal Catalog ARN")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalDecoders, "signal-decoders", "", "", "Signal Decoders")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalDecodersToAdd, "signal-decoders-to-add", "", "", "Signal Decoders To Add")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseSignalDecodersToRemove, "signal-decoders-to-remove", "", nil, "Signal Decoders To Remove")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalDecodersToUpdate, "signal-decoders-to-update", "", "", "Signal Decoders To Update")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalNodeType, "signal-node-type", "", "", "Signal Node Type")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalsToCollect, "signals-to-collect", "", "", "Signals To Collect")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSignalsToFetch, "signals-to-fetch", "", "", "Signals To Fetch")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseSpoolingMode, "spooling-mode", "", "", "Spooling Mode")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseStartTime, "start-time", "", "", "Start Time")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseStateTemplateProperties, "state-template-properties", "", nil, "State Template Properties")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseStateTemplatePropertiesToAdd, "state-template-properties-to-add", "", nil, "State Template Properties To Add")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseStateTemplatePropertiesToRemove, "state-template-properties-to-remove", "", nil, "State Template Properties To Remove")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseStateTemplates, "state-templates", "", "", "State Templates")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseStateTemplatesToAdd, "state-templates-to-add", "", "", "State Templates To Add")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseStateTemplatesToRemove, "state-templates-to-remove", "", nil, "State Templates To Remove")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseStateTemplatesToUpdate, "state-templates-to-update", "", "", "State Templates To Update")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseStatus, "status", "", "", "Status")
	_iotfleetwiseCmd.Flags().StringSliceVarP(&_iotfleetwiseTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseTags, "tags", "", "", "Tags")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseTargetArn, "target-arn", "", "", "Target ARN")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseTimestreamResources, "timestream-resources", "", "", "Timestream Resources")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseVehicleName, "vehicle-name", "", "", "Vehicle Name")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseVehicles, "vehicles", "", "", "Vehicles")
	_iotfleetwiseCmd.Flags().StringVarP(&_iotfleetwiseVss, "vss", "", "", "Vss")

	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseAssociateVehicleFleet, "associate-vehicle-fleet", "", false, "Associate Vehicle Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseBatchCreateVehicle, "batch-create-vehicle", "", false, "Batch Create Vehicle")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseBatchUpdateVehicle, "batch-update-vehicle", "", false, "Batch Update Vehicle")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateCampaign, "create-campaign", "", false, "Create Campaign")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateDecoderManifest, "create-decoder-manifest", "", false, "Create Decoder Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateFleet, "create-fleet", "", false, "Create Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateModelManifest, "create-model-manifest", "", false, "Create Model Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateSignalCatalog, "create-signal-catalog", "", false, "Create Signal Catalog")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateStateTemplate, "create-state-template", "", false, "Create State Template")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseCreateVehicle, "create-vehicle", "", false, "Create Vehicle")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteCampaign, "delete-campaign", "", false, "Delete Campaign")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteDecoderManifest, "delete-decoder-manifest", "", false, "Delete Decoder Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteFleet, "delete-fleet", "", false, "Delete Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteModelManifest, "delete-model-manifest", "", false, "Delete Model Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteSignalCatalog, "delete-signal-catalog", "", false, "Delete Signal Catalog")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteStateTemplate, "delete-state-template", "", false, "Delete State Template")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDeleteVehicle, "delete-vehicle", "", false, "Delete Vehicle")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseDisassociateVehicleFleet, "disassociate-vehicle-fleet", "", false, "Disassociate Vehicle Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetCampaign, "get-campaign", "", false, "Get Campaign")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetDecoderManifest, "get-decoder-manifest", "", false, "Get Decoder Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetEncryptionConfiguration, "get-encryption-configuration", "", false, "Get Encryption Configuration")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetFleet, "get-fleet", "", false, "Get Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetLoggingOptions, "get-logging-options", "", false, "Get Logging Options")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetModelManifest, "get-model-manifest", "", false, "Get Model Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetRegisterAccountStatus, "get-register-account-status", "", false, "Get Register Account Status")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetSignalCatalog, "get-signal-catalog", "", false, "Get Signal Catalog")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetStateTemplate, "get-state-template", "", false, "Get State Template")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetVehicle, "get-vehicle", "", false, "Get Vehicle")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseGetVehicleStatus, "get-vehicle-status", "", false, "Get Vehicle Status")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseImportDecoderManifest, "import-decoder-manifest", "", false, "Import Decoder Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseImportSignalCatalog, "import-signal-catalog", "", false, "Import Signal Catalog")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListCampaigns, "list-campaigns", "", false, "List Campaigns")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListDecoderManifestNetworkInterfaces, "list-decoder-manifest-network-interfaces", "", false, "List Decoder Manifest Network Interfaces")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListDecoderManifestSignals, "list-decoder-manifest-signals", "", false, "List Decoder Manifest Signals")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListDecoderManifests, "list-decoder-manifests", "", false, "List Decoder Manifests")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListFleets, "list-fleets", "", false, "List Fleets")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListFleetsForVehicle, "list-fleets-for-vehicle", "", false, "List Fleets For Vehicle")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListModelManifestNodes, "list-model-manifest-nodes", "", false, "List Model Manifest Nodes")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListModelManifests, "list-model-manifests", "", false, "List Model Manifests")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListSignalCatalogNodes, "list-signal-catalog-nodes", "", false, "List Signal Catalog Nodes")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListSignalCatalogs, "list-signal-catalogs", "", false, "List Signal Catalogs")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListStateTemplates, "list-state-templates", "", false, "List State Templates")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListVehicles, "list-vehicles", "", false, "List Vehicles")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseListVehiclesInFleet, "list-vehicles-in-fleet", "", false, "List Vehicles In Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwisePutEncryptionConfiguration, "put-encryption-configuration", "", false, "Put Encryption Configuration")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwisePutLoggingOptions, "put-logging-options", "", false, "Put Logging Options")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseRegisterAccount, "register-account", "", false, "Register Account")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseTagResource, "tag-resource", "", false, "Tag Resource")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateCampaign, "update-campaign", "", false, "Update Campaign")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateDecoderManifest, "update-decoder-manifest", "", false, "Update Decoder Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateFleet, "update-fleet", "", false, "Update Fleet")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateModelManifest, "update-model-manifest", "", false, "Update Model Manifest")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateSignalCatalog, "update-signal-catalog", "", false, "Update Signal Catalog")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateStateTemplate, "update-state-template", "", false, "Update State Template")
	_iotfleetwiseCmd.Flags().BoolVarP(&_iotfleetwiseUpdateVehicle, "update-vehicle", "", false, "Update Vehicle")

}
