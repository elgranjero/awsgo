package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53Cmd represents the route53 command
var _route53Cmd = &cobra.Command{
	Use:   "route53",
	Short: "AWS route53 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := route53.NewFromConfig(cfg)
		if _route53ActivateKeySigningKey {
			route53_ActivateKeySigningKey(cfg, client)
			return
		}
		if _route53AssociateVPCWithHostedZone {
			route53_AssociateVPCWithHostedZone(cfg, client)
			return
		}
		if _route53ChangeCidrCollection {
			route53_ChangeCidrCollection(cfg, client)
			return
		}
		if _route53ChangeResourceRecordSets {
			route53_ChangeResourceRecordSets(cfg, client)
			return
		}
		if _route53ChangeTagsForResource {
			route53_ChangeTagsForResource(cfg, client)
			return
		}
		if _route53CreateCidrCollection {
			route53_CreateCidrCollection(cfg, client)
			return
		}
		if _route53CreateHealthCheck {
			route53_CreateHealthCheck(cfg, client)
			return
		}
		if _route53CreateHostedZone {
			route53_CreateHostedZone(cfg, client)
			return
		}
		if _route53CreateKeySigningKey {
			route53_CreateKeySigningKey(cfg, client)
			return
		}
		if _route53CreateQueryLoggingConfig {
			route53_CreateQueryLoggingConfig(cfg, client)
			return
		}
		if _route53CreateReusableDelegationSet {
			route53_CreateReusableDelegationSet(cfg, client)
			return
		}
		if _route53CreateTrafficPolicy {
			route53_CreateTrafficPolicy(cfg, client)
			return
		}
		if _route53CreateTrafficPolicyInstance {
			route53_CreateTrafficPolicyInstance(cfg, client)
			return
		}
		if _route53CreateTrafficPolicyVersion {
			route53_CreateTrafficPolicyVersion(cfg, client)
			return
		}
		if _route53CreateVPCAssociationAuthorization {
			route53_CreateVPCAssociationAuthorization(cfg, client)
			return
		}
		if _route53DeactivateKeySigningKey {
			route53_DeactivateKeySigningKey(cfg, client)
			return
		}
		if _route53DeleteCidrCollection {
			route53_DeleteCidrCollection(cfg, client)
			return
		}
		if _route53DeleteHealthCheck {
			route53_DeleteHealthCheck(cfg, client)
			return
		}
		if _route53DeleteHostedZone {
			route53_DeleteHostedZone(cfg, client)
			return
		}
		if _route53DeleteKeySigningKey {
			route53_DeleteKeySigningKey(cfg, client)
			return
		}
		if _route53DeleteQueryLoggingConfig {
			route53_DeleteQueryLoggingConfig(cfg, client)
			return
		}
		if _route53DeleteReusableDelegationSet {
			route53_DeleteReusableDelegationSet(cfg, client)
			return
		}
		if _route53DeleteTrafficPolicy {
			route53_DeleteTrafficPolicy(cfg, client)
			return
		}
		if _route53DeleteTrafficPolicyInstance {
			route53_DeleteTrafficPolicyInstance(cfg, client)
			return
		}
		if _route53DeleteVPCAssociationAuthorization {
			route53_DeleteVPCAssociationAuthorization(cfg, client)
			return
		}
		if _route53DisableHostedZoneDNSSEC {
			route53_DisableHostedZoneDNSSEC(cfg, client)
			return
		}
		if _route53DisassociateVPCFromHostedZone {
			route53_DisassociateVPCFromHostedZone(cfg, client)
			return
		}
		if _route53EnableHostedZoneDNSSEC {
			route53_EnableHostedZoneDNSSEC(cfg, client)
			return
		}
		if _route53GetAccountLimit {
			route53_GetAccountLimit(cfg, client)
			return
		}
		if _route53GetChange {
			route53_GetChange(cfg, client)
			return
		}
		if _route53GetCheckerIpRanges {
			route53_GetCheckerIpRanges(cfg, client)
			return
		}
		if _route53GetDNSSEC {
			route53_GetDNSSEC(cfg, client)
			return
		}
		if _route53GetGeoLocation {
			route53_GetGeoLocation(cfg, client)
			return
		}
		if _route53GetHealthCheck {
			route53_GetHealthCheck(cfg, client)
			return
		}
		if _route53GetHealthCheckCount {
			route53_GetHealthCheckCount(cfg, client)
			return
		}
		if _route53GetHealthCheckLastFailureReason {
			route53_GetHealthCheckLastFailureReason(cfg, client)
			return
		}
		if _route53GetHealthCheckStatus {
			route53_GetHealthCheckStatus(cfg, client)
			return
		}
		if _route53GetHostedZone {
			route53_GetHostedZone(cfg, client)
			return
		}
		if _route53GetHostedZoneCount {
			route53_GetHostedZoneCount(cfg, client)
			return
		}
		if _route53GetHostedZoneLimit {
			route53_GetHostedZoneLimit(cfg, client)
			return
		}
		if _route53GetQueryLoggingConfig {
			route53_GetQueryLoggingConfig(cfg, client)
			return
		}
		if _route53GetReusableDelegationSet {
			route53_GetReusableDelegationSet(cfg, client)
			return
		}
		if _route53GetReusableDelegationSetLimit {
			route53_GetReusableDelegationSetLimit(cfg, client)
			return
		}
		if _route53GetTrafficPolicy {
			route53_GetTrafficPolicy(cfg, client)
			return
		}
		if _route53GetTrafficPolicyInstance {
			route53_GetTrafficPolicyInstance(cfg, client)
			return
		}
		if _route53GetTrafficPolicyInstanceCount {
			route53_GetTrafficPolicyInstanceCount(cfg, client)
			return
		}
		if _route53ListCidrBlocks {
			route53_ListCidrBlocks(cfg, client)
			return
		}
		if _route53ListCidrCollections {
			route53_ListCidrCollections(cfg, client)
			return
		}
		if _route53ListCidrLocations {
			route53_ListCidrLocations(cfg, client)
			return
		}
		if _route53ListGeoLocations {
			route53_ListGeoLocations(cfg, client)
			return
		}
		if _route53ListHealthChecks {
			route53_ListHealthChecks(cfg, client)
			return
		}
		if _route53ListHostedZones {
			route53_ListHostedZones(cfg, client)
			return
		}
		if _route53ListHostedZonesByName {
			route53_ListHostedZonesByName(cfg, client)
			return
		}
		if _route53ListHostedZonesByVPC {
			route53_ListHostedZonesByVPC(cfg, client)
			return
		}
		if _route53ListQueryLoggingConfigs {
			route53_ListQueryLoggingConfigs(cfg, client)
			return
		}
		if _route53ListResourceRecordSets {
			route53_ListResourceRecordSets(cfg, client)
			return
		}
		if _route53ListReusableDelegationSets {
			route53_ListReusableDelegationSets(cfg, client)
			return
		}
		if _route53ListTagsForResource {
			route53_ListTagsForResource(cfg, client)
			return
		}
		if _route53ListTagsForResources {
			route53_ListTagsForResources(cfg, client)
			return
		}
		if _route53ListTrafficPolicies {
			route53_ListTrafficPolicies(cfg, client)
			return
		}
		if _route53ListTrafficPolicyInstances {
			route53_ListTrafficPolicyInstances(cfg, client)
			return
		}
		if _route53ListTrafficPolicyInstancesByHostedZone {
			route53_ListTrafficPolicyInstancesByHostedZone(cfg, client)
			return
		}
		if _route53ListTrafficPolicyInstancesByPolicy {
			route53_ListTrafficPolicyInstancesByPolicy(cfg, client)
			return
		}
		if _route53ListTrafficPolicyVersions {
			route53_ListTrafficPolicyVersions(cfg, client)
			return
		}
		if _route53ListVPCAssociationAuthorizations {
			route53_ListVPCAssociationAuthorizations(cfg, client)
			return
		}
		if _route53TestDNSAnswer {
			route53_TestDNSAnswer(cfg, client)
			return
		}
		if _route53UpdateHealthCheck {
			route53_UpdateHealthCheck(cfg, client)
			return
		}
		if _route53UpdateHostedZoneComment {
			route53_UpdateHostedZoneComment(cfg, client)
			return
		}
		if _route53UpdateHostedZoneFeatures {
			route53_UpdateHostedZoneFeatures(cfg, client)
			return
		}
		if _route53UpdateTrafficPolicyComment {
			route53_UpdateTrafficPolicyComment(cfg, client)
			return
		}
		if _route53UpdateTrafficPolicyInstance {
			route53_UpdateTrafficPolicyInstance(cfg, client)
			return
		}

	},
}

var (
	_route53ActivateKeySigningKey                  bool
	_route53AssociateVPCWithHostedZone             bool
	_route53ChangeCidrCollection                   bool
	_route53ChangeResourceRecordSets               bool
	_route53ChangeTagsForResource                  bool
	_route53CreateCidrCollection                   bool
	_route53CreateHealthCheck                      bool
	_route53CreateHostedZone                       bool
	_route53CreateKeySigningKey                    bool
	_route53CreateQueryLoggingConfig               bool
	_route53CreateReusableDelegationSet            bool
	_route53CreateTrafficPolicy                    bool
	_route53CreateTrafficPolicyInstance            bool
	_route53CreateTrafficPolicyVersion             bool
	_route53CreateVPCAssociationAuthorization      bool
	_route53DeactivateKeySigningKey                bool
	_route53DeleteCidrCollection                   bool
	_route53DeleteHealthCheck                      bool
	_route53DeleteHostedZone                       bool
	_route53DeleteKeySigningKey                    bool
	_route53DeleteQueryLoggingConfig               bool
	_route53DeleteReusableDelegationSet            bool
	_route53DeleteTrafficPolicy                    bool
	_route53DeleteTrafficPolicyInstance            bool
	_route53DeleteVPCAssociationAuthorization      bool
	_route53DisableHostedZoneDNSSEC                bool
	_route53DisassociateVPCFromHostedZone          bool
	_route53EnableHostedZoneDNSSEC                 bool
	_route53GetAccountLimit                        bool
	_route53GetChange                              bool
	_route53GetCheckerIpRanges                     bool
	_route53GetDNSSEC                              bool
	_route53GetGeoLocation                         bool
	_route53GetHealthCheck                         bool
	_route53GetHealthCheckCount                    bool
	_route53GetHealthCheckLastFailureReason        bool
	_route53GetHealthCheckStatus                   bool
	_route53GetHostedZone                          bool
	_route53GetHostedZoneCount                     bool
	_route53GetHostedZoneLimit                     bool
	_route53GetQueryLoggingConfig                  bool
	_route53GetReusableDelegationSet               bool
	_route53GetReusableDelegationSetLimit          bool
	_route53GetTrafficPolicy                       bool
	_route53GetTrafficPolicyInstance               bool
	_route53GetTrafficPolicyInstanceCount          bool
	_route53ListCidrBlocks                         bool
	_route53ListCidrCollections                    bool
	_route53ListCidrLocations                      bool
	_route53ListGeoLocations                       bool
	_route53ListHealthChecks                       bool
	_route53ListHostedZones                        bool
	_route53ListHostedZonesByName                  bool
	_route53ListHostedZonesByVPC                   bool
	_route53ListQueryLoggingConfigs                bool
	_route53ListResourceRecordSets                 bool
	_route53ListReusableDelegationSets             bool
	_route53ListTagsForResource                    bool
	_route53ListTagsForResources                   bool
	_route53ListTrafficPolicies                    bool
	_route53ListTrafficPolicyInstances             bool
	_route53ListTrafficPolicyInstancesByHostedZone bool
	_route53ListTrafficPolicyInstancesByPolicy     bool
	_route53ListTrafficPolicyVersions              bool
	_route53ListVPCAssociationAuthorizations       bool
	_route53TestDNSAnswer                          bool
	_route53UpdateHealthCheck                      bool
	_route53UpdateHostedZoneComment                bool
	_route53UpdateHostedZoneFeatures               bool
	_route53UpdateTrafficPolicyComment             bool
	_route53UpdateTrafficPolicyInstance            bool

	_route53AddTags                         string
	_route53AlarmIdentifier                 string
	_route53CallerReference                 string
	_route53ChangeBatch                     string
	_route53Changes                         string
	_route53ChildHealthChecks               []string
	_route53CloudWatchLogsLogGroupArn       string
	_route53CollectionId                    string
	_route53CollectionVersion               string
	_route53Comment                         string
	_route53ContinentCode                   string
	_route53CountryCode                     string
	_route53DelegationSetId                 string
	_route53Disabled                        string
	_route53DNSName                         string
	_route53Document                        string
	_route53EDNS0ClientSubnetIP             string
	_route53EDNS0ClientSubnetMask           string
	_route53EnableAcceleratedRecovery       string
	_route53EnableSNI                       string
	_route53FailureThreshold                string
	_route53FullyQualifiedDomainName        string
	_route53HealthCheckConfig               string
	_route53HealthCheckId                   string
	_route53HealthCheckVersion              string
	_route53HealthThreshold                 string
	_route53HostedZoneConfig                string
	_route53HostedZoneId                    string
	_route53HostedZoneIdMarker              string
	_route53HostedZoneType                  string
	_route53Id                              string
	_route53InsufficientDataHealthStatus    string
	_route53Inverted                        string
	_route53IPAddress                       string
	_route53KeyManagementServiceArn         string
	_route53LocationName                    string
	_route53Marker                          string
	_route53MaxItems                        string
	_route53MaxResults                      string
	_route53Name                            string
	_route53NextToken                       string
	_route53Port                            string
	_route53RecordName                      string
	_route53RecordType                      string
	_route53Regions                         string
	_route53RemoveTagKeys                   []string
	_route53ResetElements                   string
	_route53ResolverIP                      string
	_route53ResourceId                      string
	_route53ResourceIds                     []string
	_route53ResourcePath                    string
	_route53ResourceType                    string
	_route53SearchString                    string
	_route53StartContinentCode              string
	_route53StartCountryCode                string
	_route53StartRecordIdentifier           string
	_route53StartRecordName                 string
	_route53StartRecordType                 string
	_route53StartSubdivisionCode            string
	_route53Status                          string
	_route53SubdivisionCode                 string
	_route53TrafficPolicyId                 string
	_route53TrafficPolicyIdMarker           string
	_route53TrafficPolicyInstanceNameMarker string
	_route53TrafficPolicyInstanceTypeMarker string
	_route53TrafficPolicyVersion            string
	_route53TrafficPolicyVersionMarker      string
	_route53TTL                             string
	_route53Type                            string
	_route53Version                         string
	_route53VPC                             string
	_route53VPCRegion                       string
	_route53VPCId                           string
)

// Activates a key-signing key (KSK) so that it can be used for signing by DNSSEC.
// This operation changes the KSK status to ACTIVE .
func route53_ActivateKeySigningKey(cfg aws.Config, client *route53.Client) {
	input := &route53.ActivateKeySigningKeyInput{
		// HostedZoneId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}

	if resp, err := client.ActivateKeySigningKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an Amazon VPC with a private hosted zone.
// To perform the association, the VPC and the private hosted zone must already
// exist. You can't convert a public hosted zone into a private hosted zone.
//
// If you want to associate a VPC that was created by using one Amazon Web
// Services account with a private hosted zone that was created by using a
// different account, the Amazon Web Services account that created the private
// hosted zone must first submit a CreateVPCAssociationAuthorization request. Then
// the account that created the VPC must submit an AssociateVPCWithHostedZone
// request.
//
// When granting access, the hosted zone and the Amazon VPC must belong to the
// same partition. A partition is a group of Amazon Web Services Regions. Each
// Amazon Web Services account is scoped to one partition.
//
// The following are the supported partitions:
//
// - aws - Amazon Web Services Regions
//
// - aws-cn - China Regions
//
// - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see [Access Management] in the Amazon Web Services General Reference.
//
// [Access Management]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func route53_AssociateVPCWithHostedZone(cfg aws.Config, client *route53.Client) {
	input := &route53.AssociateVPCWithHostedZoneInput{
		// HostedZoneId: *string, // Required
		// VPC: *types.VPC, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53VPC) > 0 {
		if err := assignInputField(input, "VPC", _route53VPC); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}
	if len(_route53Comment) > 0 {
		input.Comment = aws.String(_route53Comment)
	}

	if resp, err := client.AssociateVPCWithHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates, changes, or deletes CIDR blocks within a collection. Contains
// authoritative IP information mapping blocks to one or multiple locations.
//
// A change request can update multiple locations in a collection at a time, which
// is helpful if you want to move one or more CIDR blocks from one location to
// another in one transaction, without downtime.
//
// # Limits
//
// The max number of CIDR blocks included in the request is 1000. As a result, big
// updates require multiple API calls.
//
// PUT and DELETE_IF_EXISTS
//
// Use ChangeCidrCollection to perform the following actions:
//
// - PUT : Create a CIDR block within the specified collection.
//
// - DELETE_IF_EXISTS : Delete an existing CIDR block from the collection.
func route53_ChangeCidrCollection(cfg aws.Config, client *route53.Client) {
	input := &route53.ChangeCidrCollectionInput{
		// Changes: []types.CidrCollectionChange, // Required
		// Id: *string, // Required
	}

	if len(_route53Changes) > 0 {
		if err := assignInputField(input, "Changes", _route53Changes); err != nil {
			log.Errorf("invalid --changes: %s", err.Error())
			return
		}
	}
	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53CollectionVersion) > 0 {
		if err := assignInputField(input, "CollectionVersion", _route53CollectionVersion); err != nil {
			log.Errorf("invalid --collection-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.ChangeCidrCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates, changes, or deletes a resource record set, which contains
// authoritative DNS information for a specified domain name or subdomain name. For
// example, you can use ChangeResourceRecordSets to create a resource record set
// that routes traffic for test.example.com to a web server that has an IP address
// of 192.0.2.44.
//
// # Deleting Resource Record Sets
//
// To delete a resource record set, you must specify all the same values that you
// specified when you created it.
//
// # Change Batches and Transactional Changes
//
// The request body must include a document with a ChangeResourceRecordSetsRequest
// element. The request body contains a list of change items, known as a change
// batch. Change batches are considered transactional changes. Route 53 validates
// the changes in the request and then either makes all or none of the changes in
// the change batch request. This ensures that DNS routing isn't adversely affected
// by partial changes to the resource record sets in a hosted zone.
//
// For example, suppose a change batch request contains two changes: it deletes
// the CNAME resource record set for www.example.com and creates an alias resource
// record set for www.example.com. If validation for both records succeeds, Route
// 53 deletes the first resource record set and creates the second resource record
// set in a single operation. If validation for either the DELETE or the CREATE
// action fails, then the request is canceled, and the original CNAME record
// continues to exist.
//
// If you try to delete the same resource record set more than once in a single
// change batch, Route 53 returns an InvalidChangeBatch error.
//
// # Traffic Flow
//
// To create resource record sets for complex routing configurations, use either
// the traffic flow visual editor in the Route 53 console or the API actions for
// traffic policies and traffic policy instances. Save the configuration as a
// traffic policy, then associate the traffic policy with one or more domain names
// (such as example.com) or subdomain names (such as www.example.com), in the same
// hosted zone or in multiple hosted zones. You can roll back the updates if the
// new configuration isn't performing as expected. For more information, see [Using Traffic Flow to Route DNS Traffic]in
// the Amazon Route 53 Developer Guide.
//
// # Create, Delete, and Upsert
//
// Use ChangeResourceRecordsSetsRequest to perform the following actions:
//
// - CREATE : Creates a resource record set that has the specified values.
//
// - DELETE : Deletes an existing resource record set that has the specified
// values.
//
// - UPSERT : If a resource set doesn't exist, Route 53 creates it. If a resource
// set exists Route 53 updates it with the values in the request.
//
// # Syntaxes for Creating, Updating, and Deleting Resource Record Sets
//
// The syntax for a request depends on the type of resource record set that you
// want to create, delete, or update, such as weighted, alias, or failover. The XML
// elements in your request must appear in the order listed in the syntax.
//
// For an example for each type of resource record set, see "Examples."
//
// Don't refer to the syntax in the "Parameter Syntax" section, which includes all
// of the elements for every kind of resource record set that you can create,
// delete, or update by using ChangeResourceRecordSets .
//
// # Change Propagation to Route 53 DNS Servers
//
// When you submit a ChangeResourceRecordSets request, Route 53 propagates your
// changes to all of the Route 53 authoritative DNS servers managing the hosted
// zone. While your changes are propagating, GetChange returns a status of PENDING
// . When propagation is complete, GetChange returns a status of INSYNC . Changes
// generally propagate to all Route 53 name servers managing the hosted zone within
// 60 seconds. For more information, see [GetChange].
//
// # Limits on ChangeResourceRecordSets Requests
//
// For information about the limits on a ChangeResourceRecordSets request, see [Limits] in
// the Amazon Route 53 Developer Guide.
//
// [Limits]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html
// [GetChange]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html
// [Using Traffic Flow to Route DNS Traffic]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/traffic-flow.html
func route53_ChangeResourceRecordSets(cfg aws.Config, client *route53.Client) {
	input := &route53.ChangeResourceRecordSetsInput{
		// ChangeBatch: *types.ChangeBatch, // Required
		// HostedZoneId: *string, // Required
	}

	if len(_route53ChangeBatch) > 0 {
		if err := assignInputField(input, "ChangeBatch", _route53ChangeBatch); err != nil {
			log.Errorf("invalid --change-batch: %s", err.Error())
			return
		}
	}
	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}

	if resp, err := client.ChangeResourceRecordSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds, edits, or deletes tags for a health check or a hosted zone.
// For information about using tags for cost allocation, see [Using Cost Allocation Tags] in the Billing and
// Cost Management User Guide.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func route53_ChangeTagsForResource(cfg aws.Config, client *route53.Client) {
	input := &route53.ChangeTagsForResourceInput{
		// ResourceId: *string, // Required
		// ResourceType: types.TagResourceType, // Required
	}

	if len(_route53ResourceId) > 0 {
		input.ResourceId = aws.String(_route53ResourceId)
	}
	if len(_route53ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _route53ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_route53AddTags) > 0 {
		if err := assignInputField(input, "AddTags", _route53AddTags); err != nil {
			log.Errorf("invalid --add-tags: %s", err.Error())
			return
		}
	}
	if len(_route53RemoveTagKeys) > 0 {
		input.RemoveTagKeys = append([]string(nil), _route53RemoveTagKeys...)
	}

	if resp, err := client.ChangeTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a CIDR collection in the current Amazon Web Services account.
func route53_CreateCidrCollection(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateCidrCollectionInput{
		// CallerReference: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53CallerReference) > 0 {
		input.CallerReference = aws.String(_route53CallerReference)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}

	if resp, err := client.CreateCidrCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new health check.
// For information about adding health checks to resource record sets, see [HealthCheckId] in [ChangeResourceRecordSets].
//
// # ELB Load Balancers
//
// If you're registering EC2 instances with an Elastic Load Balancing (ELB) load
// balancer, do not create Amazon Route 53 health checks for the EC2 instances.
// When you register an EC2 instance with a load balancer, you configure settings
// for an ELB health check, which performs a similar function to a Route 53 health
// check.
//
// # Private Hosted Zones
//
// You can associate health checks with failover resource record sets in a private
// hosted zone. Note the following:
//
// - Route 53 health checkers are outside the VPC. To check the health of an
// endpoint within a VPC by IP address, you must assign a public IP address to the
// instance in the VPC.
//
// - You can configure a health checker to check the health of an external
// resource that the instance relies on, such as a database server.
//
// - You can create a CloudWatch metric, associate an alarm with the metric, and
// then create a health check that is based on the state of the alarm. For example,
// you might create a CloudWatch metric that checks the status of the Amazon EC2
// StatusCheckFailed metric, add an alarm to the metric, and then create a health
// check that is based on the state of the alarm. For information about creating
// CloudWatch metrics and alarms by using the CloudWatch console, see the [Amazon CloudWatch User Guide].
//
// [HealthCheckId]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResourceRecordSet.html#Route53-Type-ResourceRecordSet-HealthCheckId
// [ChangeResourceRecordSets]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html
// [Amazon CloudWatch User Guide]: https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/WhatIsCloudWatch.html
func route53_CreateHealthCheck(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateHealthCheckInput{
		// CallerReference: *string, // Required
		// HealthCheckConfig: *types.HealthCheckConfig, // Required
	}

	if len(_route53CallerReference) > 0 {
		input.CallerReference = aws.String(_route53CallerReference)
	}
	if len(_route53HealthCheckConfig) > 0 {
		if err := assignInputField(input, "HealthCheckConfig", _route53HealthCheckConfig); err != nil {
			log.Errorf("invalid --health-check-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new public or private hosted zone. You create records in a public
// hosted zone to define how you want to route traffic on the internet for a
// domain, such as example.com, and its subdomains (apex.example.com,
// acme.example.com). You create records in a private hosted zone to define how you
// want to route traffic for a domain and its subdomains within one or more Amazon
// Virtual Private Clouds (Amazon VPCs).
//
// You can't convert a public hosted zone to a private hosted zone or vice versa.
// Instead, you must create a new hosted zone with the same name and create new
// resource record sets.
//
// For more information about charges for hosted zones, see [Amazon Route 53 Pricing].
//
// Note the following:
//
// - You can't create a hosted zone for a top-level domain (TLD) such as .com.
//
// - For public hosted zones, Route 53 automatically creates a default SOA
// record and four NS records for the zone. For more information about SOA and NS
// records, see [NS and SOA Records that Route 53 Creates for a Hosted Zone]in the Amazon Route 53 Developer Guide.
//
// # If you want to use the same name servers for multiple public hosted zones, you
//
// can optionally associate a reusable delegation set with the hosted zone. See the
// DelegationSetId element.
//
// - If your domain is registered with a registrar other than Route 53, you must
// update the name servers with your registrar to make Route 53 the DNS service for
// the domain. For more information, see [Migrating DNS Service for an Existing Domain to Amazon Route 53]in the Amazon Route 53 Developer Guide.
//
// When you submit a CreateHostedZone request, the initial status of the hosted
// zone is PENDING . For public hosted zones, this means that the NS and SOA
// records are not yet available on all Route 53 DNS servers. When the NS and SOA
// records are available, the status of the zone changes to INSYNC .
//
// The CreateHostedZone request requires the caller to have an ec2:DescribeVpcs
// permission.
//
// When creating private hosted zones, the Amazon VPC must belong to the same
// partition where the hosted zone is created. A partition is a group of Amazon Web
// Services Regions. Each Amazon Web Services account is scoped to one partition.
//
// The following are the supported partitions:
//
// - aws - Amazon Web Services Regions
//
// - aws-cn - China Regions
//
// - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see [Access Management] in the Amazon Web Services General Reference.
//
// [Access Management]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [NS and SOA Records that Route 53 Creates for a Hosted Zone]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/SOA-NSrecords.html
// [Amazon Route 53 Pricing]: http://aws.amazon.com/route53/pricing/
// [Migrating DNS Service for an Existing Domain to Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingDNS.html
func route53_CreateHostedZone(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateHostedZoneInput{
		// CallerReference: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53CallerReference) > 0 {
		input.CallerReference = aws.String(_route53CallerReference)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}
	if len(_route53DelegationSetId) > 0 {
		input.DelegationSetId = aws.String(_route53DelegationSetId)
	}
	if len(_route53HostedZoneConfig) > 0 {
		if err := assignInputField(input, "HostedZoneConfig", _route53HostedZoneConfig); err != nil {
			log.Errorf("invalid --hosted-zone-config: %s", err.Error())
			return
		}
	}
	if len(_route53VPC) > 0 {
		if err := assignInputField(input, "VPC", _route53VPC); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new key-signing key (KSK) associated with a hosted zone. You can only
// have two KSKs per hosted zone.
func route53_CreateKeySigningKey(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateKeySigningKeyInput{
		// CallerReference: *string, // Required
		// HostedZoneId: *string, // Required
		// KeyManagementServiceArn: *string, // Required
		// Name: *string, // Required
		// Status: *string, // Required
	}

	if len(_route53CallerReference) > 0 {
		input.CallerReference = aws.String(_route53CallerReference)
	}
	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53KeyManagementServiceArn) > 0 {
		input.KeyManagementServiceArn = aws.String(_route53KeyManagementServiceArn)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}
	if len(_route53Status) > 0 {
		input.Status = aws.String(_route53Status)
	}

	if resp, err := client.CreateKeySigningKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration for DNS query logging. After you create a query logging
// configuration, Amazon Route 53 begins to publish log data to an Amazon
// CloudWatch Logs log group.
//
// DNS query logs contain information about the queries that Route 53 receives for
// a specified public hosted zone, such as the following:
//
// - Route 53 edge location that responded to the DNS query
//
// - Domain or subdomain that was requested
//
// - DNS record type, such as A or AAAA
//
// - DNS response code, such as NoError or ServFail
//
// Log Group and Resource Policy Before you create a query logging configuration,
// perform the following operations.
//
// If you create a query logging configuration using the Route 53 console, Route
// 53 performs these operations automatically.
//
// - Create a CloudWatch Logs log group, and make note of the ARN, which you
// specify when you create a query logging configuration. Note the following:
//
// - You must create the log group in the us-east-1 region.
//
// - You must use the same Amazon Web Services account to create the log group
// and the hosted zone that you want to configure query logging for.
//
// - When you create log groups for query logging, we recommend that you use a
// consistent prefix, for example:
//
// /aws/route53/hosted zone name
//
// In the next step, you'll create a resource policy, which controls access to one
//
// or more log groups and the associated Amazon Web Services resources, such as
// Route 53 hosted zones. There's a limit on the number of resource policies that
// you can create, so we recommend that you use a consistent prefix so you can use
// the same resource policy for all the log groups that you create for query
// logging.
//
// - Create a CloudWatch Logs resource policy, and give it the permissions that
// Route 53 needs to create log streams and to send query logs to log streams. You
// must create the CloudWatch Logs resource policy in the us-east-1 region. For the
// value of Resource , specify the ARN for the log group that you created in the
// previous step. To use the same resource policy for all the CloudWatch Logs log
// groups that you created for query logging configurations, replace the hosted
// zone name with * , for example:
//
// arn:aws:logs:us-east-1:123412341234:log-group:/aws/route53/*
//
// # To avoid the confused deputy problem, a security issue where an entity without
//
// a permission for an action can coerce a more-privileged entity to perform it,
// you can optionally limit the permissions that a service has to a resource in a
// resource-based policy by supplying the following values:
//
// - For aws:SourceArn , supply the hosted zone ARN used in creating the query
// logging configuration. For example, aws:SourceArn:
// arn:aws:route53:::hostedzone/hosted zone ID .
//
// - For aws:SourceAccount , supply the account ID for the account that creates
// the query logging configuration. For example, aws:SourceAccount:111111111111 .
//
// For more information, see [The confused deputy problem]in the Amazon Web Services IAM User Guide.
//
// You can't use the CloudWatch console to create or edit a resource policy. You
//
// must use the CloudWatch API, one of the Amazon Web Services SDKs, or the CLI.
//
// Log Streams and Edge Locations When Route 53 finishes creating the
// configuration for DNS query logging, it does the following:
//
// - Creates a log stream for an edge location the first time that the edge
// location responds to DNS queries for the specified hosted zone. That log stream
// is used to log all queries that Route 53 responds to for that edge location.
//
// - Begins to send query logs to the applicable log stream.
//
// The name of each log stream is in the following format:
//
// hosted zone ID/edge location code
//
// The edge location code is a three-letter code and an arbitrarily assigned
// number, for example, DFW3. The three-letter code typically corresponds with the
// International Air Transport Association airport code for an airport near the
// edge location. (These abbreviations might change in the future.) For a list of
// edge locations, see "The Route 53 Global Network" on the [Route 53 Product Details]page.
//
// Queries That Are Logged Query logs contain only the queries that DNS resolvers
// forward to Route 53. If a DNS resolver has already cached the response to a
// query (such as the IP address for a load balancer for example.com), the resolver
// will continue to return the cached response. It doesn't forward another query to
// Route 53 until the TTL for the corresponding resource record set expires.
// Depending on how many DNS queries are submitted for a resource record set, and
// depending on the TTL for that resource record set, query logs might contain
// information about only one query out of every several thousand queries that are
// submitted to DNS. For more information about how DNS works, see [Routing Internet Traffic to Your Website or Web Application]in the Amazon
// Route 53 Developer Guide.
//
// Log File Format For a list of the values in each query log and the format of
// each value, see [Logging DNS Queries]in the Amazon Route 53 Developer Guide.
//
// Pricing For information about charges for query logs, see [Amazon CloudWatch Pricing].
//
// How to Stop Logging If you want Route 53 to stop sending query logs to
// CloudWatch Logs, delete the query logging configuration. For more information,
// see [DeleteQueryLoggingConfig].
//
// [The confused deputy problem]: https://docs.aws.amazon.com/IAM/latest/UserGuide/confused-deputy.html
// [DeleteQueryLoggingConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteQueryLoggingConfig.html
// [Routing Internet Traffic to Your Website or Web Application]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/welcome-dns-service.html
// [Route 53 Product Details]: http://aws.amazon.com/route53/details/
// [Amazon CloudWatch Pricing]: http://aws.amazon.com/cloudwatch/pricing/
// [Logging DNS Queries]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html
func route53_CreateQueryLoggingConfig(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateQueryLoggingConfigInput{
		// CloudWatchLogsLogGroupArn: *string, // Required
		// HostedZoneId: *string, // Required
	}

	if len(_route53CloudWatchLogsLogGroupArn) > 0 {
		input.CloudWatchLogsLogGroupArn = aws.String(_route53CloudWatchLogsLogGroupArn)
	}
	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}

	if resp, err := client.CreateQueryLoggingConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a delegation set (a group of four name servers) that can be reused by
// multiple hosted zones that were created by the same Amazon Web Services account.
//
// You can also create a reusable delegation set that uses the four name servers
// that are associated with an existing hosted zone. Specify the hosted zone ID in
// the CreateReusableDelegationSet request.
//
// You can't associate a reusable delegation set with a private hosted zone.
//
// For information about using a reusable delegation set to configure white label
// name servers, see [Configuring White Label Name Servers].
//
// The process for migrating existing hosted zones to use a reusable delegation
// set is comparable to the process for configuring white label name servers. You
// need to perform the following steps:
//
// - Create a reusable delegation set.
//
// - Recreate hosted zones, and reduce the TTL to 60 seconds or less.
//
// - Recreate resource record sets in the new hosted zones.
//
// - Change the registrar's name servers to use the name servers for the new
// hosted zones.
//
// - Monitor traffic for the website or application.
//
// - Change TTLs back to their original values.
//
// If you want to migrate existing hosted zones to use a reusable delegation set,
// the existing hosted zones can't use any of the name servers that are assigned to
// the reusable delegation set. If one or more hosted zones do use one or more name
// servers that are assigned to the reusable delegation set, you can do one of the
// following:
//
// - For small numbers of hosted zones—up to a few hundred—it's relatively easy
// to create reusable delegation sets until you get one that has four name servers
// that don't overlap with any of the name servers in your hosted zones.
//
// - For larger numbers of hosted zones, the easiest solution is to use more
// than one reusable delegation set.
//
// - For larger numbers of hosted zones, you can also migrate hosted zones that
// have overlapping name servers to hosted zones that don't have overlapping name
// servers, then migrate the hosted zones again to use the reusable delegation set.
//
// [Configuring White Label Name Servers]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/white-label-name-servers.html
func route53_CreateReusableDelegationSet(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateReusableDelegationSetInput{
		// CallerReference: *string, // Required
	}

	if len(_route53CallerReference) > 0 {
		input.CallerReference = aws.String(_route53CallerReference)
	}
	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}

	if resp, err := client.CreateReusableDelegationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a traffic policy, which you use to create multiple DNS resource record
// sets for one domain name (such as example.com) or one subdomain name (such as
// www.example.com).
func route53_CreateTrafficPolicy(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateTrafficPolicyInput{
		// Document: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53Document) > 0 {
		input.Document = aws.String(_route53Document)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}
	if len(_route53Comment) > 0 {
		input.Comment = aws.String(_route53Comment)
	}

	if resp, err := client.CreateTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates resource record sets in a specified hosted zone based on the settings
// in a specified traffic policy version. In addition, CreateTrafficPolicyInstance
// associates the resource record sets with a specified domain name (such as
// example.com) or subdomain name (such as www.example.com). Amazon Route 53
// responds to DNS queries for the domain or subdomain name by using the resource
// record sets that CreateTrafficPolicyInstance created.
//
// After you submit an CreateTrafficPolicyInstance request, there's a brief delay
// while Amazon Route 53 creates the resource record sets that are specified in the
// traffic policy definition. Use GetTrafficPolicyInstance with the id of new
// traffic policy instance to confirm that the CreateTrafficPolicyInstance request
// completed successfully. For more information, see the State response element.
func route53_CreateTrafficPolicyInstance(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateTrafficPolicyInstanceInput{
		// HostedZoneId: *string, // Required
		// Name: *string, // Required
		// TTL: *int64, // Required
		// TrafficPolicyId: *string, // Required
		// TrafficPolicyVersion: *int32, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}
	if len(_route53TTL) > 0 {
		if err := assignInputField(input, "TTL", _route53TTL); err != nil {
			log.Errorf("invalid --ttl: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_route53TrafficPolicyId)
	}
	if len(_route53TrafficPolicyVersion) > 0 {
		if err := assignInputField(input, "TrafficPolicyVersion", _route53TrafficPolicyVersion); err != nil {
			log.Errorf("invalid --traffic-policy-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrafficPolicyInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of an existing traffic policy. When you create a new
// version of a traffic policy, you specify the ID of the traffic policy that you
// want to update and a JSON-formatted document that describes the new version. You
// use traffic policies to create multiple DNS resource record sets for one domain
// name (such as example.com) or one subdomain name (such as www.example.com). You
// can create a maximum of 1000 versions of a traffic policy. If you reach the
// limit and need to create another version, you'll need to start a new traffic
// policy.
func route53_CreateTrafficPolicyVersion(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateTrafficPolicyVersionInput{
		// Document: *string, // Required
		// Id: *string, // Required
	}

	if len(_route53Document) > 0 {
		input.Document = aws.String(_route53Document)
	}
	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53Comment) > 0 {
		input.Comment = aws.String(_route53Comment)
	}

	if resp, err := client.CreateTrafficPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Authorizes the Amazon Web Services account that created a specified VPC to
// submit an AssociateVPCWithHostedZone request to associate the VPC with a
// specified hosted zone that was created by a different account. To submit a
// CreateVPCAssociationAuthorization request, you must use the account that created
// the hosted zone. After you authorize the association, use the account that
// created the VPC to submit an AssociateVPCWithHostedZone request.
//
// If you want to associate multiple VPCs that you created by using one account
// with a hosted zone that you created by using a different account, you must
// submit one authorization request for each VPC.
func route53_CreateVPCAssociationAuthorization(cfg aws.Config, client *route53.Client) {
	input := &route53.CreateVPCAssociationAuthorizationInput{
		// HostedZoneId: *string, // Required
		// VPC: *types.VPC, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53VPC) > 0 {
		if err := assignInputField(input, "VPC", _route53VPC); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVPCAssociationAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates a key-signing key (KSK) so that it will not be used for signing by
// DNSSEC. This operation changes the KSK status to INACTIVE .
func route53_DeactivateKeySigningKey(cfg aws.Config, client *route53.Client) {
	input := &route53.DeactivateKeySigningKeyInput{
		// HostedZoneId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}

	if resp, err := client.DeactivateKeySigningKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a CIDR collection in the current Amazon Web Services account. The
// collection must be empty before it can be deleted.
func route53_DeleteCidrCollection(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteCidrCollectionInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.DeleteCidrCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a health check.
// Amazon Route 53 does not prevent you from deleting a health check even if the
// health check is associated with one or more resource record sets. If you delete
// a health check and you don't update the associated resource record sets, the
// future status of the health check can't be predicted and may change. This will
// affect the routing of DNS queries for your DNS failover configuration. For more
// information, see [Replacing and Deleting Health Checks]in the Amazon Route 53 Developer Guide.
//
// If you're using Cloud Map and you configured Cloud Map to create a Route 53
// health check when you register an instance, you can't use the Route 53
// DeleteHealthCheck command to delete the health check. The health check is
// deleted automatically when you deregister the instance; there can be a delay of
// several hours before the health check is deleted from Route 53.
//
// [Replacing and Deleting Health Checks]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html#health-checks-deleting.html
func route53_DeleteHealthCheck(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteHealthCheckInput{
		// HealthCheckId: *string, // Required
	}

	if len(_route53HealthCheckId) > 0 {
		input.HealthCheckId = aws.String(_route53HealthCheckId)
	}

	if resp, err := client.DeleteHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a hosted zone.
// If the hosted zone was created by another service, such as Cloud Map, see [Deleting Public Hosted Zones That Were Created by Another Service] in
// the Amazon Route 53 Developer Guide for information about how to delete it. (The
// process is the same for public and private hosted zones that were created by
// another service.)
//
// If you want to keep your domain registration but you want to stop routing
// internet traffic to your website or web application, we recommend that you
// delete resource record sets in the hosted zone instead of deleting the hosted
// zone.
//
// If you delete a hosted zone, you can't undelete it. You must create a new
// hosted zone and update the name servers for your domain registration, which can
// require up to 48 hours to take effect. (If you delegated responsibility for a
// subdomain to a hosted zone and you delete the child hosted zone, you must update
// the name servers in the parent hosted zone.) In addition, if you delete a hosted
// zone, someone could hijack the domain and route traffic to their own resources
// using your domain name.
//
// If you want to avoid the monthly charge for the hosted zone, you can transfer
// DNS service for the domain to a free DNS service. When you transfer DNS service,
// you have to update the name servers for the domain registration. If the domain
// is registered with Route 53, see [UpdateDomainNameservers]for information about how to replace Route 53
// name servers with name servers for the new DNS service. If the domain is
// registered with another registrar, use the method provided by the registrar to
// update name servers for the domain registration. For more information, perform
// an internet search on "free DNS service."
//
// You can delete a hosted zone only if it contains only the default SOA and NS
// records and has DNSSEC signing disabled. If the hosted zone contains other
// records or has DNSSEC enabled, you must delete the records and disable DNSSEC
// before deletion. Attempting to delete a hosted zone with additional records or
// DNSSEC enabled returns a HostedZoneNotEmpty error. For information about
// deleting records, see [ChangeResourceRecordSets].
//
// To verify that the hosted zone has been deleted, do one of the following:
//
// - Use the GetHostedZone action to request information about the hosted zone.
//
// - Use the ListHostedZones action to get a list of the hosted zones associated
// with the current Amazon Web Services account.
//
// [ChangeResourceRecordSets]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html
// [Deleting Public Hosted Zones That Were Created by Another Service]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DeleteHostedZone.html#delete-public-hosted-zone-created-by-another-service
// [UpdateDomainNameservers]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainNameservers.html
func route53_DeleteHostedZone(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteHostedZoneInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.DeleteHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a key-signing key (KSK). Before you can delete a KSK, you must
// deactivate it. The KSK must be deactivated before you can delete it regardless
// of whether the hosted zone is enabled for DNSSEC signing.
//
// You can use [DeactivateKeySigningKey] to deactivate the key before you delete it.
//
// Use [GetDNSSEC] to verify that the KSK is in an INACTIVE status.
//
// [DeactivateKeySigningKey]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeactivateKeySigningKey.html
// [GetDNSSEC]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetDNSSEC.html
func route53_DeleteKeySigningKey(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteKeySigningKeyInput{
		// HostedZoneId: *string, // Required
		// Name: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53Name) > 0 {
		input.Name = aws.String(_route53Name)
	}

	if resp, err := client.DeleteKeySigningKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration for DNS query logging. If you delete a configuration,
// Amazon Route 53 stops sending query logs to CloudWatch Logs. Route 53 doesn't
// delete any logs that are already in CloudWatch Logs.
//
// For more information about DNS query logs, see [CreateQueryLoggingConfig].
//
// [CreateQueryLoggingConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html
func route53_DeleteQueryLoggingConfig(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteQueryLoggingConfigInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.DeleteQueryLoggingConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a reusable delegation set.
// You can delete a reusable delegation set only if it isn't associated with any
// hosted zones.
//
// To verify that the reusable delegation set is not associated with any hosted
// zones, submit a [GetReusableDelegationSet]request and specify the ID of the reusable delegation set that
// you want to delete.
//
// [GetReusableDelegationSet]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetReusableDelegationSet.html
func route53_DeleteReusableDelegationSet(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteReusableDelegationSetInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.DeleteReusableDelegationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a traffic policy.
// When you delete a traffic policy, Route 53 sets a flag on the policy to
// indicate that it has been deleted. However, Route 53 never fully deletes the
// traffic policy. Note the following:
//
// - Deleted traffic policies aren't listed if you run [ListTrafficPolicies].
//
// - There's no way to get a list of deleted policies.
//
// - If you retain the ID of the policy, you can get information about the
// policy, including the traffic policy document, by running [GetTrafficPolicy].
//
// [ListTrafficPolicies]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListTrafficPolicies.html
// [GetTrafficPolicy]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetTrafficPolicy.html
func route53_DeleteTrafficPolicy(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteTrafficPolicyInput{
		// Id: *string, // Required
		// Version: *int32, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53Version) > 0 {
		if err := assignInputField(input, "Version", _route53Version); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a traffic policy instance and all of the resource record sets that
// Amazon Route 53 created when you created the instance.
//
// In the Route 53 console, traffic policy instances are known as policy records.
func route53_DeleteTrafficPolicyInstance(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteTrafficPolicyInstanceInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.DeleteTrafficPolicyInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes authorization to submit an AssociateVPCWithHostedZone request to
// associate a specified VPC with a hosted zone that was created by a different
// account. You must use the account that created the hosted zone to submit a
// DeleteVPCAssociationAuthorization request.
//
// Sending this request only prevents the Amazon Web Services account that created
// the VPC from associating the VPC with the Amazon Route 53 hosted zone in the
// future. If the VPC is already associated with the hosted zone,
// DeleteVPCAssociationAuthorization won't disassociate the VPC from the hosted
// zone. If you want to delete an existing association, use
// DisassociateVPCFromHostedZone .
func route53_DeleteVPCAssociationAuthorization(cfg aws.Config, client *route53.Client) {
	input := &route53.DeleteVPCAssociationAuthorizationInput{
		// HostedZoneId: *string, // Required
		// VPC: *types.VPC, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53VPC) > 0 {
		if err := assignInputField(input, "VPC", _route53VPC); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteVPCAssociationAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables DNSSEC signing in a specific hosted zone. This action does not
// deactivate any key-signing keys (KSKs) that are active in the hosted zone.
func route53_DisableHostedZoneDNSSEC(cfg aws.Config, client *route53.Client) {
	input := &route53.DisableHostedZoneDNSSECInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}

	if resp, err := client.DisableHostedZoneDNSSEC(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Amazon Virtual Private Cloud (Amazon VPC) from an Amazon Route
// 53 private hosted zone. Note the following:
//
// - You can't disassociate the last Amazon VPC from a private hosted zone.
//
// - You can't convert a private hosted zone into a public hosted zone.
//
// - You can submit a DisassociateVPCFromHostedZone request using either the
// account that created the hosted zone or the account that created the Amazon VPC.
//
// - Some services, such as Cloud Map and Amazon Elastic File System (Amazon
// EFS) automatically create hosted zones and associate VPCs with the hosted zones.
// A service can create a hosted zone using your account or using its own account.
// You can disassociate a VPC from a hosted zone only if the service created the
// hosted zone using your account.
//
// When you run [DisassociateVPCFromHostedZone], if the hosted zone has a value for OwningAccount , you can use
//
// DisassociateVPCFromHostedZone . If the hosted zone has a value for
// OwningService , you can't use DisassociateVPCFromHostedZone .
//
// When revoking access, the hosted zone and the Amazon VPC must belong to the
// same partition. A partition is a group of Amazon Web Services Regions. Each
// Amazon Web Services account is scoped to one partition.
//
// The following are the supported partitions:
//
// - aws - Amazon Web Services Regions
//
// - aws-cn - China Regions
//
// - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see [Access Management] in the Amazon Web Services General Reference.
//
// [Access Management]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [DisassociateVPCFromHostedZone]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListHostedZonesByVPC.html
func route53_DisassociateVPCFromHostedZone(cfg aws.Config, client *route53.Client) {
	input := &route53.DisassociateVPCFromHostedZoneInput{
		// HostedZoneId: *string, // Required
		// VPC: *types.VPC, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53VPC) > 0 {
		if err := assignInputField(input, "VPC", _route53VPC); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}
	if len(_route53Comment) > 0 {
		input.Comment = aws.String(_route53Comment)
	}

	if resp, err := client.DisassociateVPCFromHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables DNSSEC signing in a specific hosted zone.
func route53_EnableHostedZoneDNSSEC(cfg aws.Config, client *route53.Client) {
	input := &route53.EnableHostedZoneDNSSECInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}

	if resp, err := client.EnableHostedZoneDNSSEC(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified limit for the current account, for example, the maximum
// number of health checks that you can create using the account.
//
// For the default limit, see [Limits] in the Amazon Route 53 Developer Guide. To request
// a higher limit, [open a case].
//
// You can also view account limits in Amazon Web Services Trusted Advisor. Sign
// in to the Amazon Web Services Management Console and open the Trusted Advisor
// console at [https://console.aws.amazon.com/trustedadvisor/]. Then choose Service limits in the navigation pane.
//
// [Limits]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html
// [https://console.aws.amazon.com/trustedadvisor/]: https://console.aws.amazon.com/trustedadvisor
// [open a case]: https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53
func route53_GetAccountLimit(cfg aws.Config, client *route53.Client) {
	input := &route53.GetAccountLimitInput{
		// Type: types.AccountLimitType, // Required
	}

	if len(_route53Type) > 0 {
		if err := assignInputField(input, "Type", _route53Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAccountLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current status of a change batch request. The status is one of the
// following values:
//
// - PENDING indicates that the changes in this request have not propagated to
// all Amazon Route 53 DNS servers managing the hosted zone. This is the initial
// status of all change batch requests.
//
// - INSYNC indicates that the changes have propagated to all Route 53 DNS
// servers managing the hosted zone.
func route53_GetChange(cfg aws.Config, client *route53.Client) {
	input := &route53.GetChangeInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.GetChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public.
//
// GetCheckerIpRanges still works, but we recommend that you download
// ip-ranges.json, which includes IP address ranges for all Amazon Web Services
// services. For more information, see [IP Address Ranges of Amazon Route 53 Servers]in the Amazon Route 53 Developer Guide.
//
// [IP Address Ranges of Amazon Route 53 Servers]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/route-53-ip-addresses.html
func route53_GetCheckerIpRanges(cfg aws.Config, client *route53.Client) {
	input := &route53.GetCheckerIpRangesInput{}

	if resp, err := client.GetCheckerIpRanges(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about DNSSEC for a specific hosted zone, including the
// key-signing keys (KSKs) in the hosted zone.
func route53_GetDNSSEC(cfg aws.Config, client *route53.Client) {
	input := &route53.GetDNSSECInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}

	if resp, err := client.GetDNSSEC(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about whether a specified geographic location is supported for
// Amazon Route 53 geolocation resource record sets.
//
// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public.
//
// Use the following syntax to determine whether a continent is supported for
// geolocation:
//
// GET /2013-04-01/geolocation?continentcode=two-letter abbreviation for a
// continent
//
// Use the following syntax to determine whether a country is supported for
// geolocation:
//
// GET /2013-04-01/geolocation?countrycode=two-character country code
//
// Use the following syntax to determine whether a subdivision of a country is
// supported for geolocation:
//
// GET /2013-04-01/geolocation?countrycode=two-character country
// code&subdivisioncode=subdivision code
func route53_GetGeoLocation(cfg aws.Config, client *route53.Client) {
	input := &route53.GetGeoLocationInput{}

	if len(_route53ContinentCode) > 0 {
		input.ContinentCode = aws.String(_route53ContinentCode)
	}
	if len(_route53CountryCode) > 0 {
		input.CountryCode = aws.String(_route53CountryCode)
	}
	if len(_route53SubdivisionCode) > 0 {
		input.SubdivisionCode = aws.String(_route53SubdivisionCode)
	}

	if resp, err := client.GetGeoLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified health check.
func route53_GetHealthCheck(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHealthCheckInput{
		// HealthCheckId: *string, // Required
	}

	if len(_route53HealthCheckId) > 0 {
		input.HealthCheckId = aws.String(_route53HealthCheckId)
	}

	if resp, err := client.GetHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the number of health checks that are associated with the current
// Amazon Web Services account.
func route53_GetHealthCheckCount(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHealthCheckCountInput{}

	if resp, err := client.GetHealthCheckCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the reason that a specified health check failed most recently.
func route53_GetHealthCheckLastFailureReason(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHealthCheckLastFailureReasonInput{
		// HealthCheckId: *string, // Required
	}

	if len(_route53HealthCheckId) > 0 {
		input.HealthCheckId = aws.String(_route53HealthCheckId)
	}

	if resp, err := client.GetHealthCheckLastFailureReason(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets status of a specified health check.
// This API is intended for use during development to diagnose behavior. It
// doesn’t support production use-cases with high query rates that require
// immediate and actionable responses.
func route53_GetHealthCheckStatus(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHealthCheckStatusInput{
		// HealthCheckId: *string, // Required
	}

	if len(_route53HealthCheckId) > 0 {
		input.HealthCheckId = aws.String(_route53HealthCheckId)
	}

	if resp, err := client.GetHealthCheckStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified hosted zone including the four name servers
// assigned to the hosted zone.
//
// returns the VPCs associated with the specified hosted zone and does not reflect
// the VPC associations by Route 53 Profiles. To get the associations to a Profile,
// call the [ListProfileAssociations]API.
//
// [ListProfileAssociations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53profiles_ListProfileAssociations.html
func route53_GetHostedZone(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHostedZoneInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.GetHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the number of hosted zones that are associated with the current
// Amazon Web Services account.
func route53_GetHostedZoneCount(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHostedZoneCountInput{}

	if resp, err := client.GetHostedZoneCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the specified limit for a specified hosted zone, for example, the maximum
// number of records that you can create in the hosted zone.
//
// For the default limit, see [Limits] in the Amazon Route 53 Developer Guide. To request
// a higher limit, [open a case].
//
// [Limits]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html
// [open a case]: https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53
func route53_GetHostedZoneLimit(cfg aws.Config, client *route53.Client) {
	input := &route53.GetHostedZoneLimitInput{
		// HostedZoneId: *string, // Required
		// Type: types.HostedZoneLimitType, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53Type) > 0 {
		if err := assignInputField(input, "Type", _route53Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetHostedZoneLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified configuration for DNS query logging.
// For more information about DNS query logs, see [CreateQueryLoggingConfig] and [Logging DNS Queries].
//
// [CreateQueryLoggingConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html
// [Logging DNS Queries]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html
func route53_GetQueryLoggingConfig(cfg aws.Config, client *route53.Client) {
	input := &route53.GetQueryLoggingConfigInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.GetQueryLoggingConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specified reusable delegation set, including the
// four name servers that are assigned to the delegation set.
func route53_GetReusableDelegationSet(cfg aws.Config, client *route53.Client) {
	input := &route53.GetReusableDelegationSetInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.GetReusableDelegationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the maximum number of hosted zones that you can associate with the
// specified reusable delegation set.
//
// For the default limit, see [Limits] in the Amazon Route 53 Developer Guide. To request
// a higher limit, [open a case].
//
// [Limits]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html
// [open a case]: https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53
func route53_GetReusableDelegationSetLimit(cfg aws.Config, client *route53.Client) {
	input := &route53.GetReusableDelegationSetLimitInput{
		// DelegationSetId: *string, // Required
		// Type: types.ReusableDelegationSetLimitType, // Required
	}

	if len(_route53DelegationSetId) > 0 {
		input.DelegationSetId = aws.String(_route53DelegationSetId)
	}
	if len(_route53Type) > 0 {
		if err := assignInputField(input, "Type", _route53Type); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReusableDelegationSetLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific traffic policy version.
// For information about how of deleting a traffic policy affects the response
// from GetTrafficPolicy , see [DeleteTrafficPolicy].
//
// [DeleteTrafficPolicy]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteTrafficPolicy.html
func route53_GetTrafficPolicy(cfg aws.Config, client *route53.Client) {
	input := &route53.GetTrafficPolicyInput{
		// Id: *string, // Required
		// Version: *int32, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53Version) > 0 {
		if err := assignInputField(input, "Version", _route53Version); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specified traffic policy instance.
// Use GetTrafficPolicyInstance with the id of new traffic policy instance to
// confirm that the CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance
// request completed successfully. For more information, see the State response
// element.
//
// In the Route 53 console, traffic policy instances are known as policy records.
func route53_GetTrafficPolicyInstance(cfg aws.Config, client *route53.Client) {
	input := &route53.GetTrafficPolicyInstanceInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}

	if resp, err := client.GetTrafficPolicyInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the number of traffic policy instances that are associated with the
// current Amazon Web Services account.
func route53_GetTrafficPolicyInstanceCount(cfg aws.Config, client *route53.Client) {
	input := &route53.GetTrafficPolicyInstanceCountInput{}

	if resp, err := client.GetTrafficPolicyInstanceCount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a paginated list of location objects and their CIDR blocks.
func route53_ListCidrBlocks(cfg aws.Config, client *route53.Client) {
	input := &route53.ListCidrBlocksInput{
		// CollectionId: *string, // Required
	}

	if len(_route53CollectionId) > 0 {
		input.CollectionId = aws.String(_route53CollectionId)
	}
	if len(_route53LocationName) > 0 {
		input.LocationName = aws.String(_route53LocationName)
	}
	if len(_route53MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53NextToken) > 0 {
		input.NextToken = aws.String(_route53NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCidrBlocks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53.ListCidrBlocksOutput
	p := route53.NewListCidrBlocksPaginator(client, input)
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

// Returns a paginated list of CIDR collections in the Amazon Web Services account
// (metadata only).
func route53_ListCidrCollections(cfg aws.Config, client *route53.Client) {
	input := &route53.ListCidrCollectionsInput{}

	if len(_route53MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53NextToken) > 0 {
		input.NextToken = aws.String(_route53NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCidrCollections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53.ListCidrCollectionsOutput
	p := route53.NewListCidrCollectionsPaginator(client, input)
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

// Returns a paginated list of CIDR locations for the given collection (metadata
// only, does not include CIDR blocks).
func route53_ListCidrLocations(cfg aws.Config, client *route53.Client) {
	input := &route53.ListCidrLocationsInput{
		// CollectionId: *string, // Required
	}

	if len(_route53CollectionId) > 0 {
		input.CollectionId = aws.String(_route53CollectionId)
	}
	if len(_route53MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53NextToken) > 0 {
		input.NextToken = aws.String(_route53NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCidrLocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53.ListCidrLocationsOutput
	p := route53.NewListCidrLocationsPaginator(client, input)
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

// Retrieves a list of supported geographic locations.
// Countries are listed first, and continents are listed last. If Amazon Route 53
// supports subdivisions for a country (for example, states or provinces), the
// subdivisions for that country are listed in alphabetical order immediately after
// the corresponding country.
//
// Route 53 does not perform authorization for this API because it retrieves
// information that is already available to the public.
//
// For a list of supported geolocation codes, see the [GeoLocation] data type.
//
// [GeoLocation]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GeoLocation.html
func route53_ListGeoLocations(cfg aws.Config, client *route53.Client) {
	input := &route53.ListGeoLocationsInput{}

	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53StartContinentCode) > 0 {
		input.StartContinentCode = aws.String(_route53StartContinentCode)
	}
	if len(_route53StartCountryCode) > 0 {
		input.StartCountryCode = aws.String(_route53StartCountryCode)
	}
	if len(_route53StartSubdivisionCode) > 0 {
		input.StartSubdivisionCode = aws.String(_route53StartSubdivisionCode)
	}

	if resp, err := client.ListGeoLocations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a list of the health checks that are associated with the current
// Amazon Web Services account.
func route53_ListHealthChecks(cfg aws.Config, client *route53.Client) {
	input := &route53.ListHealthChecksInput{}

	if len(_route53Marker) > 0 {
		input.Marker = aws.String(_route53Marker)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHealthChecks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53.ListHealthChecksOutput
	p := route53.NewListHealthChecksPaginator(client, input)
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

// Retrieves a list of the public and private hosted zones that are associated
// with the current Amazon Web Services account. The response includes a
// HostedZones child element for each hosted zone.
//
// Amazon Route 53 returns a maximum of 100 items in each response. If you have a
// lot of hosted zones, you can use the maxitems parameter to list them in groups
// of up to 100.
func route53_ListHostedZones(cfg aws.Config, client *route53.Client) {
	input := &route53.ListHostedZonesInput{}

	if len(_route53DelegationSetId) > 0 {
		input.DelegationSetId = aws.String(_route53DelegationSetId)
	}
	if len(_route53HostedZoneType) > 0 {
		if err := assignInputField(input, "HostedZoneType", _route53HostedZoneType); err != nil {
			log.Errorf("invalid --hosted-zone-type: %s", err.Error())
			return
		}
	}
	if len(_route53Marker) > 0 {
		input.Marker = aws.String(_route53Marker)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListHostedZones(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53.ListHostedZonesOutput
	p := route53.NewListHostedZonesPaginator(client, input)
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

// Retrieves a list of your hosted zones in lexicographic order. The response
// includes a HostedZones child element for each hosted zone created by the
// current Amazon Web Services account.
//
// ListHostedZonesByName sorts hosted zones by name with the labels reversed. For
// example:
//
// com.example.www.
//
// Note the trailing dot, which can change the sort order in some circumstances.
//
// If the domain name includes escape characters or Punycode, ListHostedZonesByName
// alphabetizes the domain name using the escaped or Punycoded value, which is the
// format that Amazon Route 53 saves in its database. For example, to create a
// hosted zone for exämple.com, you specify ex\344mple.com for the domain name.
// ListHostedZonesByName alphabetizes it as:
//
// com.ex\344mple.
//
// The labels are reversed and alphabetized using the escaped value. For more
// information about valid domain name formats, including internationalized domain
// names, see [DNS Domain Name Format]in the Amazon Route 53 Developer Guide.
//
// Route 53 returns up to 100 items in each response. If you have a lot of hosted
// zones, use the MaxItems parameter to list them in groups of up to 100. The
// response includes values that help navigate from one group of MaxItems hosted
// zones to the next:
//
// - The DNSName and HostedZoneId elements in the response contain the values, if
// any, specified for the dnsname and hostedzoneid parameters in the request that
// produced the current response.
//
// - The MaxItems element in the response contains the value, if any, that you
// specified for the maxitems parameter in the request that produced the current
// response.
//
// - If the value of IsTruncated in the response is true, there are more hosted
// zones associated with the current Amazon Web Services account.
//
// # If IsTruncated is false, this response includes the last hosted zone that is
//
// associated with the current account. The NextDNSName element and
// NextHostedZoneId elements are omitted from the response.
//
// - The NextDNSName and NextHostedZoneId elements in the response contain the
// domain name and the hosted zone ID of the next hosted zone that is associated
// with the current Amazon Web Services account. If you want to list more hosted
// zones, make another call to ListHostedZonesByName , and specify the value of
// NextDNSName and NextHostedZoneId in the dnsname and hostedzoneid parameters,
// respectively.
//
// [DNS Domain Name Format]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html
func route53_ListHostedZonesByName(cfg aws.Config, client *route53.Client) {
	input := &route53.ListHostedZonesByNameInput{}

	if len(_route53DNSName) > 0 {
		input.DNSName = aws.String(_route53DNSName)
	}
	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListHostedZonesByName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the private hosted zones that a specified VPC is associated with,
// regardless of which Amazon Web Services account or Amazon Web Services service
// owns the hosted zones. The HostedZoneOwner structure in the response contains
// one of the following values:
//
// - An OwningAccount element, which contains the account number of either the
// current Amazon Web Services account or another Amazon Web Services account. Some
// services, such as Cloud Map, create hosted zones using the current account.
//
// - An OwningService element, which identifies the Amazon Web Services service
// that created and owns the hosted zone. For example, if a hosted zone was created
// by Amazon Elastic File System (Amazon EFS), the value of Owner is
// efs.amazonaws.com .
//
// ListHostedZonesByVPC returns the hosted zones associated with the specified VPC
// and does not reflect the hosted zone associations to VPCs via Route 53 Profiles.
// To get the associations to a Profile, call the [ListProfileResourceAssociations]API.
//
// When listing private hosted zones, the hosted zone and the Amazon VPC must
// belong to the same partition where the hosted zones were created. A partition is
// a group of Amazon Web Services Regions. Each Amazon Web Services account is
// scoped to one partition.
//
// The following are the supported partitions:
//
// - aws - Amazon Web Services Regions
//
// - aws-cn - China Regions
//
// - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see [Access Management] in the Amazon Web Services General Reference.
//
// [Access Management]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
// [ListProfileResourceAssociations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53profiles_ListProfileResourceAssociations.html
func route53_ListHostedZonesByVPC(cfg aws.Config, client *route53.Client) {
	input := &route53.ListHostedZonesByVPCInput{
		// VPCId: *string, // Required
		// VPCRegion: types.VPCRegion, // Required
	}

	if len(_route53VPCId) > 0 {
		input.VPCId = aws.String(_route53VPCId)
	}
	if len(_route53VPCRegion) > 0 {
		if err := assignInputField(input, "VPCRegion", _route53VPCRegion); err != nil {
			log.Errorf("invalid --vpc-region: %s", err.Error())
			return
		}
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53NextToken) > 0 {
		input.NextToken = aws.String(_route53NextToken)
	}

	if resp, err := client.ListHostedZonesByVPC(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the configurations for DNS query logging that are associated with the
// current Amazon Web Services account or the configuration that is associated with
// a specified hosted zone.
//
// For more information about DNS query logs, see [CreateQueryLoggingConfig]. Additional information,
// including the format of DNS query logs, appears in [Logging DNS Queries]in the Amazon Route 53
// Developer Guide.
//
// [CreateQueryLoggingConfig]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html
// [Logging DNS Queries]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html
func route53_ListQueryLoggingConfigs(cfg aws.Config, client *route53.Client) {
	input := &route53.ListQueryLoggingConfigsInput{}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53NextToken) > 0 {
		input.NextToken = aws.String(_route53NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListQueryLoggingConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53.ListQueryLoggingConfigsOutput
	p := route53.NewListQueryLoggingConfigsPaginator(client, input)
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

// Lists the resource record sets in a specified hosted zone.
// ListResourceRecordSets returns up to 300 resource record sets at a time in
// ASCII order, beginning at a position specified by the name and type elements.
//
// # Sort order
//
// ListResourceRecordSets sorts results first by DNS name with the labels
// reversed, for example:
//
// com.example.www.
//
// Note the trailing dot, which can change the sort order when the record name
// contains characters that appear before . (decimal 46) in the ASCII table. These
// characters include the following: ! " # $ % & ' ( ) * + , -
//
// When multiple records have the same DNS name, ListResourceRecordSets sorts
// results by the record type.
//
// # Specifying where to start listing records
//
// You can use the name and type elements to specify the resource record set that
// the list begins with:
//
// If you do not specify Name or Type The results begin with the first resource
// record set that the hosted zone contains.
//
// If you specify Name but not Type The results begin with the first resource
// record set in the list whose name is greater than or equal to Name .
//
// If you specify Type but not Name Amazon Route 53 returns the InvalidInput error.
//
// If you specify both Name and Type The results begin with the first resource
// record set in the list whose name is greater than or equal to Name , and whose
// type is greater than or equal to Type .
//
// Type is only used to sort between records with the same record Name.
//
// # Resource record sets that are PENDING
//
// This action returns the most current version of the records. This includes
// records that are PENDING , and that are not yet available on all Route 53 DNS
// servers.
//
// # Changing resource record sets
//
// To ensure that you get an accurate listing of the resource record sets for a
// hosted zone at a point in time, do not submit a ChangeResourceRecordSets
// request while you're paging through the results of a ListResourceRecordSets
// request. If you do, some pages may display results without the latest changes
// while other pages display results with the latest changes.
//
// # Displaying the next page of results
//
// If a ListResourceRecordSets command returns more than one page of results, the
// value of IsTruncated is true . To display the next page of results, get the
// values of NextRecordName , NextRecordType , and NextRecordIdentifier (if any)
// from the response. Then submit another ListResourceRecordSets request, and
// specify those values for StartRecordName , StartRecordType , and
// StartRecordIdentifier .
func route53_ListResourceRecordSets(cfg aws.Config, client *route53.Client) {
	input := &route53.ListResourceRecordSetsInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53StartRecordIdentifier) > 0 {
		input.StartRecordIdentifier = aws.String(_route53StartRecordIdentifier)
	}
	if len(_route53StartRecordName) > 0 {
		input.StartRecordName = aws.String(_route53StartRecordName)
	}
	if len(_route53StartRecordType) > 0 {
		if err := assignInputField(input, "StartRecordType", _route53StartRecordType); err != nil {
			log.Errorf("invalid --start-record-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListResourceRecordSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the reusable delegation sets that are associated with the
// current Amazon Web Services account.
func route53_ListReusableDelegationSets(cfg aws.Config, client *route53.Client) {
	input := &route53.ListReusableDelegationSetsInput{}

	if len(_route53Marker) > 0 {
		input.Marker = aws.String(_route53Marker)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListReusableDelegationSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tags for one health check or hosted zone.
// For information about using tags for cost allocation, see [Using Cost Allocation Tags] in the Billing and
// Cost Management User Guide.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func route53_ListTagsForResource(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTagsForResourceInput{
		// ResourceId: *string, // Required
		// ResourceType: types.TagResourceType, // Required
	}

	if len(_route53ResourceId) > 0 {
		input.ResourceId = aws.String(_route53ResourceId)
	}
	if len(_route53ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _route53ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists tags for up to 10 health checks or hosted zones.
// For information about using tags for cost allocation, see [Using Cost Allocation Tags] in the Billing and
// Cost Management User Guide.
//
// [Using Cost Allocation Tags]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func route53_ListTagsForResources(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTagsForResourcesInput{
		// ResourceIds: []string, // Required
		// ResourceType: types.TagResourceType, // Required
	}

	if len(_route53ResourceIds) > 0 {
		input.ResourceIds = append([]string(nil), _route53ResourceIds...)
	}
	if len(_route53ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _route53ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTagsForResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the latest version for every traffic policy that is
// associated with the current Amazon Web Services account. Policies are listed in
// the order that they were created in.
//
// For information about how of deleting a traffic policy affects the response
// from ListTrafficPolicies , see [DeleteTrafficPolicy].
//
// [DeleteTrafficPolicy]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteTrafficPolicy.html
func route53_ListTrafficPolicies(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTrafficPoliciesInput{}

	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyIdMarker) > 0 {
		input.TrafficPolicyIdMarker = aws.String(_route53TrafficPolicyIdMarker)
	}

	if resp, err := client.ListTrafficPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the traffic policy instances that you created by using
// the current Amazon Web Services account.
//
// After you submit an UpdateTrafficPolicyInstance request, there's a brief delay
// while Amazon Route 53 creates the resource record sets that are specified in the
// traffic policy definition. For more information, see the State response element.
//
// Route 53 returns a maximum of 100 items in each response. If you have a lot of
// traffic policy instances, you can use the MaxItems parameter to list them in
// groups of up to 100.
func route53_ListTrafficPolicyInstances(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTrafficPolicyInstancesInput{}

	if len(_route53HostedZoneIdMarker) > 0 {
		input.HostedZoneIdMarker = aws.String(_route53HostedZoneIdMarker)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyInstanceNameMarker) > 0 {
		input.TrafficPolicyInstanceNameMarker = aws.String(_route53TrafficPolicyInstanceNameMarker)
	}
	if len(_route53TrafficPolicyInstanceTypeMarker) > 0 {
		if err := assignInputField(input, "TrafficPolicyInstanceTypeMarker", _route53TrafficPolicyInstanceTypeMarker); err != nil {
			log.Errorf("invalid --traffic-policy-instance-type-marker: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTrafficPolicyInstances(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the traffic policy instances that you created in a
// specified hosted zone.
//
// After you submit a CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance
// request, there's a brief delay while Amazon Route 53 creates the resource record
// sets that are specified in the traffic policy definition. For more information,
// see the State response element.
//
// Route 53 returns a maximum of 100 items in each response. If you have a lot of
// traffic policy instances, you can use the MaxItems parameter to list them in
// groups of up to 100.
func route53_ListTrafficPolicyInstancesByHostedZone(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTrafficPolicyInstancesByHostedZoneInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyInstanceNameMarker) > 0 {
		input.TrafficPolicyInstanceNameMarker = aws.String(_route53TrafficPolicyInstanceNameMarker)
	}
	if len(_route53TrafficPolicyInstanceTypeMarker) > 0 {
		if err := assignInputField(input, "TrafficPolicyInstanceTypeMarker", _route53TrafficPolicyInstanceTypeMarker); err != nil {
			log.Errorf("invalid --traffic-policy-instance-type-marker: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTrafficPolicyInstancesByHostedZone(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the traffic policy instances that you created by using a
// specify traffic policy version.
//
// After you submit a CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance
// request, there's a brief delay while Amazon Route 53 creates the resource record
// sets that are specified in the traffic policy definition. For more information,
// see the State response element.
//
// Route 53 returns a maximum of 100 items in each response. If you have a lot of
// traffic policy instances, you can use the MaxItems parameter to list them in
// groups of up to 100.
func route53_ListTrafficPolicyInstancesByPolicy(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTrafficPolicyInstancesByPolicyInput{
		// TrafficPolicyId: *string, // Required
		// TrafficPolicyVersion: *int32, // Required
	}

	if len(_route53TrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_route53TrafficPolicyId)
	}
	if len(_route53TrafficPolicyVersion) > 0 {
		if err := assignInputField(input, "TrafficPolicyVersion", _route53TrafficPolicyVersion); err != nil {
			log.Errorf("invalid --traffic-policy-version: %s", err.Error())
			return
		}
	}
	if len(_route53HostedZoneIdMarker) > 0 {
		input.HostedZoneIdMarker = aws.String(_route53HostedZoneIdMarker)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyInstanceNameMarker) > 0 {
		input.TrafficPolicyInstanceNameMarker = aws.String(_route53TrafficPolicyInstanceNameMarker)
	}
	if len(_route53TrafficPolicyInstanceTypeMarker) > 0 {
		if err := assignInputField(input, "TrafficPolicyInstanceTypeMarker", _route53TrafficPolicyInstanceTypeMarker); err != nil {
			log.Errorf("invalid --traffic-policy-instance-type-marker: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTrafficPolicyInstancesByPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about all of the versions for a specified traffic policy.
// Traffic policy versions are listed in numerical order by VersionNumber .
func route53_ListTrafficPolicyVersions(cfg aws.Config, client *route53.Client) {
	input := &route53.ListTrafficPolicyVersionsInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyVersionMarker) > 0 {
		input.TrafficPolicyVersionMarker = aws.String(_route53TrafficPolicyVersionMarker)
	}

	if resp, err := client.ListTrafficPolicyVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of the VPCs that were created by other accounts and that can be
// associated with a specified hosted zone because you've submitted one or more
// CreateVPCAssociationAuthorization requests.
//
// The response includes a VPCs element with a VPC child element for each VPC that
// can be associated with the hosted zone.
func route53_ListVPCAssociationAuthorizations(cfg aws.Config, client *route53.Client) {
	input := &route53.ListVPCAssociationAuthorizationsInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _route53MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_route53NextToken) > 0 {
		input.NextToken = aws.String(_route53NextToken)
	}

	if resp, err := client.ListVPCAssociationAuthorizations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the value that Amazon Route 53 returns in response to a DNS request for a
// specified record name and type. You can optionally specify the IP address of a
// DNS resolver, an EDNS0 client subnet IP address, and a subnet mask.
//
// This call only supports querying public hosted zones.
//
// The TestDnsAnswer  returns information similar to what you would expect from
// the answer section of the dig command. Therefore, if you query for the name
// servers of a subdomain that point to the parent name servers, those will not be
// returned.
func route53_TestDNSAnswer(cfg aws.Config, client *route53.Client) {
	input := &route53.TestDNSAnswerInput{
		// HostedZoneId: *string, // Required
		// RecordName: *string, // Required
		// RecordType: types.RRType, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53RecordName) > 0 {
		input.RecordName = aws.String(_route53RecordName)
	}
	if len(_route53RecordType) > 0 {
		if err := assignInputField(input, "RecordType", _route53RecordType); err != nil {
			log.Errorf("invalid --record-type: %s", err.Error())
			return
		}
	}
	if len(_route53EDNS0ClientSubnetIP) > 0 {
		input.EDNS0ClientSubnetIP = aws.String(_route53EDNS0ClientSubnetIP)
	}
	if len(_route53EDNS0ClientSubnetMask) > 0 {
		input.EDNS0ClientSubnetMask = aws.String(_route53EDNS0ClientSubnetMask)
	}
	if len(_route53ResolverIP) > 0 {
		input.ResolverIP = aws.String(_route53ResolverIP)
	}

	if resp, err := client.TestDNSAnswer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing health check. Note that some values can't be updated.
// For more information about updating health checks, see [Creating, Updating, and Deleting Health Checks] in the Amazon Route 53
// Developer Guide.
//
// [Creating, Updating, and Deleting Health Checks]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html
func route53_UpdateHealthCheck(cfg aws.Config, client *route53.Client) {
	input := &route53.UpdateHealthCheckInput{
		// HealthCheckId: *string, // Required
	}

	if len(_route53HealthCheckId) > 0 {
		input.HealthCheckId = aws.String(_route53HealthCheckId)
	}
	if len(_route53AlarmIdentifier) > 0 {
		if err := assignInputField(input, "AlarmIdentifier", _route53AlarmIdentifier); err != nil {
			log.Errorf("invalid --alarm-identifier: %s", err.Error())
			return
		}
	}
	if len(_route53ChildHealthChecks) > 0 {
		input.ChildHealthChecks = append([]string(nil), _route53ChildHealthChecks...)
	}
	if len(_route53Disabled) > 0 {
		if err := assignInputField(input, "Disabled", _route53Disabled); err != nil {
			log.Errorf("invalid --disabled: %s", err.Error())
			return
		}
	}
	if len(_route53EnableSNI) > 0 {
		if err := assignInputField(input, "EnableSNI", _route53EnableSNI); err != nil {
			log.Errorf("invalid --enable-sni: %s", err.Error())
			return
		}
	}
	if len(_route53FailureThreshold) > 0 {
		if err := assignInputField(input, "FailureThreshold", _route53FailureThreshold); err != nil {
			log.Errorf("invalid --failure-threshold: %s", err.Error())
			return
		}
	}
	if len(_route53FullyQualifiedDomainName) > 0 {
		input.FullyQualifiedDomainName = aws.String(_route53FullyQualifiedDomainName)
	}
	if len(_route53HealthCheckVersion) > 0 {
		if err := assignInputField(input, "HealthCheckVersion", _route53HealthCheckVersion); err != nil {
			log.Errorf("invalid --health-check-version: %s", err.Error())
			return
		}
	}
	if len(_route53HealthThreshold) > 0 {
		if err := assignInputField(input, "HealthThreshold", _route53HealthThreshold); err != nil {
			log.Errorf("invalid --health-threshold: %s", err.Error())
			return
		}
	}
	if len(_route53IPAddress) > 0 {
		input.IPAddress = aws.String(_route53IPAddress)
	}
	if len(_route53InsufficientDataHealthStatus) > 0 {
		if err := assignInputField(input, "InsufficientDataHealthStatus", _route53InsufficientDataHealthStatus); err != nil {
			log.Errorf("invalid --insufficient-data-health-status: %s", err.Error())
			return
		}
	}
	if len(_route53Inverted) > 0 {
		if err := assignInputField(input, "Inverted", _route53Inverted); err != nil {
			log.Errorf("invalid --inverted: %s", err.Error())
			return
		}
	}
	if len(_route53Port) > 0 {
		if err := assignInputField(input, "Port", _route53Port); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_route53Regions) > 0 {
		if err := assignInputField(input, "Regions", _route53Regions); err != nil {
			log.Errorf("invalid --regions: %s", err.Error())
			return
		}
	}
	if len(_route53ResetElements) > 0 {
		if err := assignInputField(input, "ResetElements", _route53ResetElements); err != nil {
			log.Errorf("invalid --reset-elements: %s", err.Error())
			return
		}
	}
	if len(_route53ResourcePath) > 0 {
		input.ResourcePath = aws.String(_route53ResourcePath)
	}
	if len(_route53SearchString) > 0 {
		input.SearchString = aws.String(_route53SearchString)
	}

	if resp, err := client.UpdateHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the comment for a specified hosted zone.
func route53_UpdateHostedZoneComment(cfg aws.Config, client *route53.Client) {
	input := &route53.UpdateHostedZoneCommentInput{
		// Id: *string, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53Comment) > 0 {
		input.Comment = aws.String(_route53Comment)
	}

	if resp, err := client.UpdateHostedZoneComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the features configuration for a hosted zone. This operation allows you
// to enable or disable specific features for your hosted zone, such as accelerated
// recovery.
//
// Accelerated recovery enables you to update DNS records in your public hosted
// zone even when the us-east-1 region is unavailable.
func route53_UpdateHostedZoneFeatures(cfg aws.Config, client *route53.Client) {
	input := &route53.UpdateHostedZoneFeaturesInput{
		// HostedZoneId: *string, // Required
	}

	if len(_route53HostedZoneId) > 0 {
		input.HostedZoneId = aws.String(_route53HostedZoneId)
	}
	if len(_route53EnableAcceleratedRecovery) > 0 {
		if err := assignInputField(input, "EnableAcceleratedRecovery", _route53EnableAcceleratedRecovery); err != nil {
			log.Errorf("invalid --enable-accelerated-recovery: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateHostedZoneFeatures(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the comment for a specified traffic policy version.
func route53_UpdateTrafficPolicyComment(cfg aws.Config, client *route53.Client) {
	input := &route53.UpdateTrafficPolicyCommentInput{
		// Comment: *string, // Required
		// Id: *string, // Required
		// Version: *int32, // Required
	}

	if len(_route53Comment) > 0 {
		input.Comment = aws.String(_route53Comment)
	}
	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53Version) > 0 {
		if err := assignInputField(input, "Version", _route53Version); err != nil {
			log.Errorf("invalid --version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrafficPolicyComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// After you submit a UpdateTrafficPolicyInstance request, there's a brief delay
// while Route 53 creates the resource record sets that are specified in the
// traffic policy definition. Use GetTrafficPolicyInstance with the id of updated
// traffic policy instance confirm that the UpdateTrafficPolicyInstance request
// completed successfully. For more information, see the State response element.
//
// Updates the resource record sets in a specified hosted zone that were created
// based on the settings in a specified traffic policy version.
//
// When you update a traffic policy instance, Amazon Route 53 continues to respond
// to DNS queries for the root resource record set name (such as example.com) while
// it replaces one group of resource record sets with another. Route 53 performs
// the following operations:
//
// - Route 53 creates a new group of resource record sets based on the specified
// traffic policy. This is true regardless of how significant the differences are
// between the existing resource record sets and the new resource record sets.
//
// - When all of the new resource record sets have been created, Route 53 starts
// to respond to DNS queries for the root resource record set name (such as
// example.com) by using the new resource record sets.
//
// - Route 53 deletes the old group of resource record sets that are associated
// with the root resource record set name.
func route53_UpdateTrafficPolicyInstance(cfg aws.Config, client *route53.Client) {
	input := &route53.UpdateTrafficPolicyInstanceInput{
		// Id: *string, // Required
		// TTL: *int64, // Required
		// TrafficPolicyId: *string, // Required
		// TrafficPolicyVersion: *int32, // Required
	}

	if len(_route53Id) > 0 {
		input.Id = aws.String(_route53Id)
	}
	if len(_route53TTL) > 0 {
		if err := assignInputField(input, "TTL", _route53TTL); err != nil {
			log.Errorf("invalid --ttl: %s", err.Error())
			return
		}
	}
	if len(_route53TrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_route53TrafficPolicyId)
	}
	if len(_route53TrafficPolicyVersion) > 0 {
		if err := assignInputField(input, "TrafficPolicyVersion", _route53TrafficPolicyVersion); err != nil {
			log.Errorf("invalid --traffic-policy-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTrafficPolicyInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_route53Cmd)
	_route53Cmd.Flags().SortFlags = false

	_route53Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_route53Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53Cmd.Flags().StringVarP(&_route53AddTags, "add-tags", "", "", "Add Tags")
	_route53Cmd.Flags().StringVarP(&_route53AlarmIdentifier, "alarm-identifier", "", "", "Alarm Identifier")
	_route53Cmd.Flags().StringVarP(&_route53CallerReference, "caller-reference", "", "", "Caller Reference")
	_route53Cmd.Flags().StringVarP(&_route53ChangeBatch, "change-batch", "", "", "Change Batch")
	_route53Cmd.Flags().StringVarP(&_route53Changes, "changes", "", "", "Changes")
	_route53Cmd.Flags().StringSliceVarP(&_route53ChildHealthChecks, "child-health-checks", "", nil, "Child Health Checks")
	_route53Cmd.Flags().StringVarP(&_route53CloudWatchLogsLogGroupArn, "cloud-watch-logs-log-group-arn", "", "", "Cloud Watch Logs Log Group ARN")
	_route53Cmd.Flags().StringVarP(&_route53CollectionId, "collection-id", "", "", "Collection ID")
	_route53Cmd.Flags().StringVarP(&_route53CollectionVersion, "collection-version", "", "", "Collection Version")
	_route53Cmd.Flags().StringVarP(&_route53Comment, "comment", "", "", "Comment")
	_route53Cmd.Flags().StringVarP(&_route53ContinentCode, "continent-code", "", "", "Continent Code")
	_route53Cmd.Flags().StringVarP(&_route53CountryCode, "country-code", "", "", "Country Code")
	_route53Cmd.Flags().StringVarP(&_route53DelegationSetId, "delegation-set-id", "", "", "Delegation Set ID")
	_route53Cmd.Flags().StringVarP(&_route53Disabled, "disabled", "", "", "Disabled")
	_route53Cmd.Flags().StringVarP(&_route53DNSName, "dns-name", "", "", "DNS Name")
	_route53Cmd.Flags().StringVarP(&_route53Document, "document", "", "", "Document")
	_route53Cmd.Flags().StringVarP(&_route53EDNS0ClientSubnetIP, "edns0-client-subnet-ip", "", "", "Edns0 Client Subnet IP")
	_route53Cmd.Flags().StringVarP(&_route53EDNS0ClientSubnetMask, "edns0-client-subnet-mask", "", "", "Edns0 Client Subnet Mask")
	_route53Cmd.Flags().StringVarP(&_route53EnableAcceleratedRecovery, "enable-accelerated-recovery", "", "", "Enable Accelerated Recovery")
	_route53Cmd.Flags().StringVarP(&_route53EnableSNI, "enable-sni", "", "", "Enable Sni")
	_route53Cmd.Flags().StringVarP(&_route53FailureThreshold, "failure-threshold", "", "", "Failure Threshold")
	_route53Cmd.Flags().StringVarP(&_route53FullyQualifiedDomainName, "fully-qualified-domain-name", "", "", "Fully Qualified Domain Name")
	_route53Cmd.Flags().StringVarP(&_route53HealthCheckConfig, "health-check-config", "", "", "Health Check Config")
	_route53Cmd.Flags().StringVarP(&_route53HealthCheckId, "health-check-id", "", "", "Health Check ID")
	_route53Cmd.Flags().StringVarP(&_route53HealthCheckVersion, "health-check-version", "", "", "Health Check Version")
	_route53Cmd.Flags().StringVarP(&_route53HealthThreshold, "health-threshold", "", "", "Health Threshold")
	_route53Cmd.Flags().StringVarP(&_route53HostedZoneConfig, "hosted-zone-config", "", "", "Hosted Zone Config")
	_route53Cmd.Flags().StringVarP(&_route53HostedZoneId, "hosted-zone-id", "", "", "Hosted Zone ID")
	_route53Cmd.Flags().StringVarP(&_route53HostedZoneIdMarker, "hosted-zone-id-marker", "", "", "Hosted Zone ID Marker")
	_route53Cmd.Flags().StringVarP(&_route53HostedZoneType, "hosted-zone-type", "", "", "Hosted Zone Type")
	_route53Cmd.Flags().StringVarP(&_route53Id, "id", "", "", "ID")
	_route53Cmd.Flags().StringVarP(&_route53InsufficientDataHealthStatus, "insufficient-data-health-status", "", "", "Insufficient Data Health Status")
	_route53Cmd.Flags().StringVarP(&_route53Inverted, "inverted", "", "", "Inverted")
	_route53Cmd.Flags().StringVarP(&_route53IPAddress, "ip-address", "", "", "IP Address")
	_route53Cmd.Flags().StringVarP(&_route53KeyManagementServiceArn, "key-management-service-arn", "", "", "Key Management Service ARN")
	_route53Cmd.Flags().StringVarP(&_route53LocationName, "location-name", "", "", "Location Name")
	_route53Cmd.Flags().StringVarP(&_route53Marker, "marker", "", "", "Marker")
	_route53Cmd.Flags().StringVarP(&_route53MaxItems, "max-items", "", "", "Max Items")
	_route53Cmd.Flags().StringVarP(&_route53MaxResults, "max-results", "", "", "Max Results")
	_route53Cmd.Flags().StringVarP(&_route53Name, "name", "", "", "Name")
	_route53Cmd.Flags().StringVarP(&_route53NextToken, "next-token", "", "", "Next Token")
	_route53Cmd.Flags().StringVarP(&_route53Port, "port", "", "", "Port")
	_route53Cmd.Flags().StringVarP(&_route53RecordName, "record-name", "", "", "Record Name")
	_route53Cmd.Flags().StringVarP(&_route53RecordType, "record-type", "", "", "Record Type")
	_route53Cmd.Flags().StringVarP(&_route53Regions, "regions", "", "", "Regions")
	_route53Cmd.Flags().StringSliceVarP(&_route53RemoveTagKeys, "remove-tag-keys", "", nil, "Remove Tag Keys")
	_route53Cmd.Flags().StringVarP(&_route53ResetElements, "reset-elements", "", "", "Reset Elements")
	_route53Cmd.Flags().StringVarP(&_route53ResolverIP, "resolver-ip", "", "", "Resolver IP")
	_route53Cmd.Flags().StringVarP(&_route53ResourceId, "resource-id", "", "", "Resource ID")
	_route53Cmd.Flags().StringSliceVarP(&_route53ResourceIds, "resource-ids", "", nil, "Resource Ids")
	_route53Cmd.Flags().StringVarP(&_route53ResourcePath, "resource-path", "", "", "Resource Path")
	_route53Cmd.Flags().StringVarP(&_route53ResourceType, "resource-type", "", "", "Resource Type")
	_route53Cmd.Flags().StringVarP(&_route53SearchString, "search-string", "", "", "Search String")
	_route53Cmd.Flags().StringVarP(&_route53StartContinentCode, "start-continent-code", "", "", "Start Continent Code")
	_route53Cmd.Flags().StringVarP(&_route53StartCountryCode, "start-country-code", "", "", "Start Country Code")
	_route53Cmd.Flags().StringVarP(&_route53StartRecordIdentifier, "start-record-identifier", "", "", "Start Record Identifier")
	_route53Cmd.Flags().StringVarP(&_route53StartRecordName, "start-record-name", "", "", "Start Record Name")
	_route53Cmd.Flags().StringVarP(&_route53StartRecordType, "start-record-type", "", "", "Start Record Type")
	_route53Cmd.Flags().StringVarP(&_route53StartSubdivisionCode, "start-subdivision-code", "", "", "Start Subdivision Code")
	_route53Cmd.Flags().StringVarP(&_route53Status, "status", "", "", "Status")
	_route53Cmd.Flags().StringVarP(&_route53SubdivisionCode, "subdivision-code", "", "", "Subdivision Code")
	_route53Cmd.Flags().StringVarP(&_route53TrafficPolicyId, "traffic-policy-id", "", "", "Traffic Policy ID")
	_route53Cmd.Flags().StringVarP(&_route53TrafficPolicyIdMarker, "traffic-policy-id-marker", "", "", "Traffic Policy ID Marker")
	_route53Cmd.Flags().StringVarP(&_route53TrafficPolicyInstanceNameMarker, "traffic-policy-instance-name-marker", "", "", "Traffic Policy Instance Name Marker")
	_route53Cmd.Flags().StringVarP(&_route53TrafficPolicyInstanceTypeMarker, "traffic-policy-instance-type-marker", "", "", "Traffic Policy Instance Type Marker")
	_route53Cmd.Flags().StringVarP(&_route53TrafficPolicyVersion, "traffic-policy-version", "", "", "Traffic Policy Version")
	_route53Cmd.Flags().StringVarP(&_route53TrafficPolicyVersionMarker, "traffic-policy-version-marker", "", "", "Traffic Policy Version Marker")
	_route53Cmd.Flags().StringVarP(&_route53TTL, "ttl", "", "", "TTL")
	_route53Cmd.Flags().StringVarP(&_route53Type, "type", "", "", "Type")
	_route53Cmd.Flags().StringVarP(&_route53Version, "version", "", "", "Version")
	_route53Cmd.Flags().StringVarP(&_route53VPC, "vpc", "", "", "VPC")
	_route53Cmd.Flags().StringVarP(&_route53VPCRegion, "vpc-region", "", "", "VPC Region")
	_route53Cmd.Flags().StringVarP(&_route53VPCId, "vpcid", "", "", "Vpcid")

	_route53Cmd.Flags().BoolVarP(&_route53ActivateKeySigningKey, "activate-key-signing-key", "", false, "Activate Key Signing Key")
	_route53Cmd.Flags().BoolVarP(&_route53AssociateVPCWithHostedZone, "associate-vpc-with-hosted-zone", "", false, "Associate VPC With Hosted Zone")
	_route53Cmd.Flags().BoolVarP(&_route53ChangeCidrCollection, "change-cidr-collection", "", false, "Change CIDR Collection")
	_route53Cmd.Flags().BoolVarP(&_route53ChangeResourceRecordSets, "change-resource-record-sets", "", false, "Change Resource Record Sets")
	_route53Cmd.Flags().BoolVarP(&_route53ChangeTagsForResource, "change-tags-for-resource", "", false, "Change Tags For Resource")
	_route53Cmd.Flags().BoolVarP(&_route53CreateCidrCollection, "create-cidr-collection", "", false, "Create CIDR Collection")
	_route53Cmd.Flags().BoolVarP(&_route53CreateHealthCheck, "create-health-check", "", false, "Create Health Check")
	_route53Cmd.Flags().BoolVarP(&_route53CreateHostedZone, "create-hosted-zone", "", false, "Create Hosted Zone")
	_route53Cmd.Flags().BoolVarP(&_route53CreateKeySigningKey, "create-key-signing-key", "", false, "Create Key Signing Key")
	_route53Cmd.Flags().BoolVarP(&_route53CreateQueryLoggingConfig, "create-query-logging-config", "", false, "Create Query Logging Config")
	_route53Cmd.Flags().BoolVarP(&_route53CreateReusableDelegationSet, "create-reusable-delegation-set", "", false, "Create Reusable Delegation Set")
	_route53Cmd.Flags().BoolVarP(&_route53CreateTrafficPolicy, "create-traffic-policy", "", false, "Create Traffic Policy")
	_route53Cmd.Flags().BoolVarP(&_route53CreateTrafficPolicyInstance, "create-traffic-policy-instance", "", false, "Create Traffic Policy Instance")
	_route53Cmd.Flags().BoolVarP(&_route53CreateTrafficPolicyVersion, "create-traffic-policy-version", "", false, "Create Traffic Policy Version")
	_route53Cmd.Flags().BoolVarP(&_route53CreateVPCAssociationAuthorization, "create-vpc-association-authorization", "", false, "Create VPC Association Authorization")
	_route53Cmd.Flags().BoolVarP(&_route53DeactivateKeySigningKey, "deactivate-key-signing-key", "", false, "Deactivate Key Signing Key")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteCidrCollection, "delete-cidr-collection", "", false, "Delete CIDR Collection")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteHealthCheck, "delete-health-check", "", false, "Delete Health Check")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteHostedZone, "delete-hosted-zone", "", false, "Delete Hosted Zone")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteKeySigningKey, "delete-key-signing-key", "", false, "Delete Key Signing Key")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteQueryLoggingConfig, "delete-query-logging-config", "", false, "Delete Query Logging Config")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteReusableDelegationSet, "delete-reusable-delegation-set", "", false, "Delete Reusable Delegation Set")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteTrafficPolicy, "delete-traffic-policy", "", false, "Delete Traffic Policy")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteTrafficPolicyInstance, "delete-traffic-policy-instance", "", false, "Delete Traffic Policy Instance")
	_route53Cmd.Flags().BoolVarP(&_route53DeleteVPCAssociationAuthorization, "delete-vpc-association-authorization", "", false, "Delete VPC Association Authorization")
	_route53Cmd.Flags().BoolVarP(&_route53DisableHostedZoneDNSSEC, "disable-hosted-zone-dnssec", "", false, "Disable Hosted Zone Dnssec")
	_route53Cmd.Flags().BoolVarP(&_route53DisassociateVPCFromHostedZone, "disassociate-vpc-from-hosted-zone", "", false, "Disassociate VPC From Hosted Zone")
	_route53Cmd.Flags().BoolVarP(&_route53EnableHostedZoneDNSSEC, "enable-hosted-zone-dnssec", "", false, "Enable Hosted Zone Dnssec")
	_route53Cmd.Flags().BoolVarP(&_route53GetAccountLimit, "get-account-limit", "", false, "Get Account Limit")
	_route53Cmd.Flags().BoolVarP(&_route53GetChange, "get-change", "", false, "Get Change")
	_route53Cmd.Flags().BoolVarP(&_route53GetCheckerIpRanges, "get-checker-ip-ranges", "", false, "Get Checker IP Ranges")
	_route53Cmd.Flags().BoolVarP(&_route53GetDNSSEC, "get-dnssec", "", false, "Get Dnssec")
	_route53Cmd.Flags().BoolVarP(&_route53GetGeoLocation, "get-geo-location", "", false, "Get Geo Location")
	_route53Cmd.Flags().BoolVarP(&_route53GetHealthCheck, "get-health-check", "", false, "Get Health Check")
	_route53Cmd.Flags().BoolVarP(&_route53GetHealthCheckCount, "get-health-check-count", "", false, "Get Health Check Count")
	_route53Cmd.Flags().BoolVarP(&_route53GetHealthCheckLastFailureReason, "get-health-check-last-failure-reason", "", false, "Get Health Check Last Failure Reason")
	_route53Cmd.Flags().BoolVarP(&_route53GetHealthCheckStatus, "get-health-check-status", "", false, "Get Health Check Status")
	_route53Cmd.Flags().BoolVarP(&_route53GetHostedZone, "get-hosted-zone", "", false, "Get Hosted Zone")
	_route53Cmd.Flags().BoolVarP(&_route53GetHostedZoneCount, "get-hosted-zone-count", "", false, "Get Hosted Zone Count")
	_route53Cmd.Flags().BoolVarP(&_route53GetHostedZoneLimit, "get-hosted-zone-limit", "", false, "Get Hosted Zone Limit")
	_route53Cmd.Flags().BoolVarP(&_route53GetQueryLoggingConfig, "get-query-logging-config", "", false, "Get Query Logging Config")
	_route53Cmd.Flags().BoolVarP(&_route53GetReusableDelegationSet, "get-reusable-delegation-set", "", false, "Get Reusable Delegation Set")
	_route53Cmd.Flags().BoolVarP(&_route53GetReusableDelegationSetLimit, "get-reusable-delegation-set-limit", "", false, "Get Reusable Delegation Set Limit")
	_route53Cmd.Flags().BoolVarP(&_route53GetTrafficPolicy, "get-traffic-policy", "", false, "Get Traffic Policy")
	_route53Cmd.Flags().BoolVarP(&_route53GetTrafficPolicyInstance, "get-traffic-policy-instance", "", false, "Get Traffic Policy Instance")
	_route53Cmd.Flags().BoolVarP(&_route53GetTrafficPolicyInstanceCount, "get-traffic-policy-instance-count", "", false, "Get Traffic Policy Instance Count")
	_route53Cmd.Flags().BoolVarP(&_route53ListCidrBlocks, "list-cidr-blocks", "", false, "List CIDR Blocks")
	_route53Cmd.Flags().BoolVarP(&_route53ListCidrCollections, "list-cidr-collections", "", false, "List CIDR Collections")
	_route53Cmd.Flags().BoolVarP(&_route53ListCidrLocations, "list-cidr-locations", "", false, "List CIDR Locations")
	_route53Cmd.Flags().BoolVarP(&_route53ListGeoLocations, "list-geo-locations", "", false, "List Geo Locations")
	_route53Cmd.Flags().BoolVarP(&_route53ListHealthChecks, "list-health-checks", "", false, "List Health Checks")
	_route53Cmd.Flags().BoolVarP(&_route53ListHostedZones, "list-hosted-zones", "", false, "List Hosted Zones")
	_route53Cmd.Flags().BoolVarP(&_route53ListHostedZonesByName, "list-hosted-zones-by-name", "", false, "List Hosted Zones By Name")
	_route53Cmd.Flags().BoolVarP(&_route53ListHostedZonesByVPC, "list-hosted-zones-by-vpc", "", false, "List Hosted Zones By VPC")
	_route53Cmd.Flags().BoolVarP(&_route53ListQueryLoggingConfigs, "list-query-logging-configs", "", false, "List Query Logging Configs")
	_route53Cmd.Flags().BoolVarP(&_route53ListResourceRecordSets, "list-resource-record-sets", "", false, "List Resource Record Sets")
	_route53Cmd.Flags().BoolVarP(&_route53ListReusableDelegationSets, "list-reusable-delegation-sets", "", false, "List Reusable Delegation Sets")
	_route53Cmd.Flags().BoolVarP(&_route53ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_route53Cmd.Flags().BoolVarP(&_route53ListTagsForResources, "list-tags-for-resources", "", false, "List Tags For Resources")
	_route53Cmd.Flags().BoolVarP(&_route53ListTrafficPolicies, "list-traffic-policies", "", false, "List Traffic Policies")
	_route53Cmd.Flags().BoolVarP(&_route53ListTrafficPolicyInstances, "list-traffic-policy-instances", "", false, "List Traffic Policy Instances")
	_route53Cmd.Flags().BoolVarP(&_route53ListTrafficPolicyInstancesByHostedZone, "list-traffic-policy-instances-by-hosted-zone", "", false, "List Traffic Policy Instances By Hosted Zone")
	_route53Cmd.Flags().BoolVarP(&_route53ListTrafficPolicyInstancesByPolicy, "list-traffic-policy-instances-by-policy", "", false, "List Traffic Policy Instances By Policy")
	_route53Cmd.Flags().BoolVarP(&_route53ListTrafficPolicyVersions, "list-traffic-policy-versions", "", false, "List Traffic Policy Versions")
	_route53Cmd.Flags().BoolVarP(&_route53ListVPCAssociationAuthorizations, "list-vpc-association-authorizations", "", false, "List VPC Association Authorizations")
	_route53Cmd.Flags().BoolVarP(&_route53TestDNSAnswer, "test-dns-answer", "", false, "Test DNS Answer")
	_route53Cmd.Flags().BoolVarP(&_route53UpdateHealthCheck, "update-health-check", "", false, "Update Health Check")
	_route53Cmd.Flags().BoolVarP(&_route53UpdateHostedZoneComment, "update-hosted-zone-comment", "", false, "Update Hosted Zone Comment")
	_route53Cmd.Flags().BoolVarP(&_route53UpdateHostedZoneFeatures, "update-hosted-zone-features", "", false, "Update Hosted Zone Features")
	_route53Cmd.Flags().BoolVarP(&_route53UpdateTrafficPolicyComment, "update-traffic-policy-comment", "", false, "Update Traffic Policy Comment")
	_route53Cmd.Flags().BoolVarP(&_route53UpdateTrafficPolicyInstance, "update-traffic-policy-instance", "", false, "Update Traffic Policy Instance")

}
