package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/snowball/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-cluster", "cancel-job", "create-address", "create-cluster", "create-job", "create-long-term-pricing", "create-return-shipping-label", "describe-address", "describe-addresses", "describe-cluster", "describe-job", "describe-return-shipping-label", "get-job-manifest", "get-job-unlock-code", "get-snowball-usage", "get-software-updates", "list-cluster-jobs", "list-clusters", "list-compatible-images", "list-jobs", "list-long-term-pricing", "list-pickup-locations", "list-service-versions", "update-cluster", "update-job", "update-job-shipment-state", "update-long-term-pricing"},
		OperationSet: map[string]bool{"cancel-cluster": true, "cancel-job": true, "create-address": true, "create-cluster": true, "create-job": true, "create-long-term-pricing": true, "create-return-shipping-label": true, "describe-address": true, "describe-addresses": true, "describe-cluster": true, "describe-job": true, "describe-return-shipping-label": true, "get-job-manifest": true, "get-job-unlock-code": true, "get-snowball-usage": true, "get-software-updates": true, "list-cluster-jobs": true, "list-clusters": true, "list-compatible-images": true, "list-jobs": true, "list-long-term-pricing": true, "list-pickup-locations": true, "list-service-versions": true, "update-cluster": true, "update-job": true, "update-job-shipment-state": true, "update-long-term-pricing": true},
		OperationInputs: map[string][]string{
			"cancel-cluster":                 {"ClusterId"},
			"cancel-job":                     {"JobId"},
			"create-address":                 {"Address"},
			"create-cluster":                 {"AddressId", "Description", "ForceCreateJobs", "ForwardingAddressId", "InitialClusterSize", "JobType", "KmsKeyARN", "LongTermPricingIds", "Notification", "OnDeviceServiceConfiguration", "RemoteManagement", "Resources", "RoleARN", "ShippingOption", "SnowballCapacityPreference", "SnowballType", "TaxDocuments"},
			"create-job":                     {"AddressId", "ClusterId", "Description", "DeviceConfiguration", "ForwardingAddressId", "ImpactLevel", "JobType", "KmsKeyARN", "LongTermPricingId", "Notification", "OnDeviceServiceConfiguration", "PickupDetails", "RemoteManagement", "Resources", "RoleARN", "ShippingOption", "SnowballCapacityPreference", "SnowballType", "TaxDocuments"},
			"create-long-term-pricing":       {"IsLongTermPricingAutoRenew", "LongTermPricingType", "SnowballType"},
			"create-return-shipping-label":   {"JobId", "ShippingOption"},
			"describe-address":               {"AddressId"},
			"describe-addresses":             {"MaxResults", "NextToken"},
			"describe-cluster":               {"ClusterId"},
			"describe-job":                   {"JobId"},
			"describe-return-shipping-label": {"JobId"},
			"get-job-manifest":               {"JobId"},
			"get-job-unlock-code":            {"JobId"},
			"get-snowball-usage":             {},
			"get-software-updates":           {"JobId"},
			"list-cluster-jobs":              {"ClusterId", "MaxResults", "NextToken"},
			"list-clusters":                  {"MaxResults", "NextToken"},
			"list-compatible-images":         {"MaxResults", "NextToken"},
			"list-jobs":                      {"MaxResults", "NextToken"},
			"list-long-term-pricing":         {"MaxResults", "NextToken"},
			"list-pickup-locations":          {"MaxResults", "NextToken"},
			"list-service-versions":          {"DependentServices", "MaxResults", "NextToken", "ServiceName"},
			"update-cluster":                 {"AddressId", "ClusterId", "Description", "ForwardingAddressId", "Notification", "OnDeviceServiceConfiguration", "Resources", "RoleARN", "ShippingOption"},
			"update-job":                     {"AddressId", "Description", "ForwardingAddressId", "JobId", "Notification", "OnDeviceServiceConfiguration", "PickupDetails", "Resources", "RoleARN", "ShippingOption", "SnowballCapacityPreference"},
			"update-job-shipment-state":      {"JobId", "ShipmentState"},
			"update-long-term-pricing":       {"IsLongTermPricingAutoRenew", "LongTermPricingId", "ReplacementJob"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-cluster":                 {"ClusterId": "*string"},
			"cancel-job":                     {"JobId": "*string"},
			"create-address":                 {"Address": "*types.Address"},
			"create-cluster":                 {"AddressId": "*string", "Description": "*string", "ForceCreateJobs": "bool", "ForwardingAddressId": "*string", "InitialClusterSize": "*int32", "JobType": "types.JobType", "KmsKeyARN": "*string", "LongTermPricingIds": "[]string", "Notification": "*types.Notification", "OnDeviceServiceConfiguration": "*types.OnDeviceServiceConfiguration", "RemoteManagement": "types.RemoteManagement", "Resources": "*types.JobResource", "RoleARN": "*string", "ShippingOption": "types.ShippingOption", "SnowballCapacityPreference": "types.SnowballCapacity", "SnowballType": "types.SnowballType", "TaxDocuments": "*types.TaxDocuments"},
			"create-job":                     {"AddressId": "*string", "ClusterId": "*string", "Description": "*string", "DeviceConfiguration": "*types.DeviceConfiguration", "ForwardingAddressId": "*string", "ImpactLevel": "types.ImpactLevel", "JobType": "types.JobType", "KmsKeyARN": "*string", "LongTermPricingId": "*string", "Notification": "*types.Notification", "OnDeviceServiceConfiguration": "*types.OnDeviceServiceConfiguration", "PickupDetails": "*types.PickupDetails", "RemoteManagement": "types.RemoteManagement", "Resources": "*types.JobResource", "RoleARN": "*string", "ShippingOption": "types.ShippingOption", "SnowballCapacityPreference": "types.SnowballCapacity", "SnowballType": "types.SnowballType", "TaxDocuments": "*types.TaxDocuments"},
			"create-long-term-pricing":       {"IsLongTermPricingAutoRenew": "*bool", "LongTermPricingType": "types.LongTermPricingType", "SnowballType": "types.SnowballType"},
			"create-return-shipping-label":   {"JobId": "*string", "ShippingOption": "types.ShippingOption"},
			"describe-address":               {"AddressId": "*string"},
			"describe-addresses":             {"MaxResults": "*int32", "NextToken": "*string"},
			"describe-cluster":               {"ClusterId": "*string"},
			"describe-job":                   {"JobId": "*string"},
			"describe-return-shipping-label": {"JobId": "*string"},
			"get-job-manifest":               {"JobId": "*string"},
			"get-job-unlock-code":            {"JobId": "*string"},
			"get-snowball-usage":             {},
			"get-software-updates":           {"JobId": "*string"},
			"list-cluster-jobs":              {"ClusterId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-clusters":                  {"MaxResults": "*int32", "NextToken": "*string"},
			"list-compatible-images":         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-jobs":                      {"MaxResults": "*int32", "NextToken": "*string"},
			"list-long-term-pricing":         {"MaxResults": "*int32", "NextToken": "*string"},
			"list-pickup-locations":          {"MaxResults": "*int32", "NextToken": "*string"},
			"list-service-versions":          {"DependentServices": "[]types.DependentService", "MaxResults": "*int32", "NextToken": "*string", "ServiceName": "types.ServiceName"},
			"update-cluster":                 {"AddressId": "*string", "ClusterId": "*string", "Description": "*string", "ForwardingAddressId": "*string", "Notification": "*types.Notification", "OnDeviceServiceConfiguration": "*types.OnDeviceServiceConfiguration", "Resources": "*types.JobResource", "RoleARN": "*string", "ShippingOption": "types.ShippingOption"},
			"update-job":                     {"AddressId": "*string", "Description": "*string", "ForwardingAddressId": "*string", "JobId": "*string", "Notification": "*types.Notification", "OnDeviceServiceConfiguration": "*types.OnDeviceServiceConfiguration", "PickupDetails": "*types.PickupDetails", "Resources": "*types.JobResource", "RoleARN": "*string", "ShippingOption": "types.ShippingOption", "SnowballCapacityPreference": "types.SnowballCapacity"},
			"update-job-shipment-state":      {"JobId": "*string", "ShipmentState": "types.ShipmentState"},
			"update-long-term-pricing":       {"IsLongTermPricingAutoRenew": "*bool", "LongTermPricingId": "*string", "ReplacementJob": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-cluster":                 {"ClusterId"},
			"cancel-job":                     {"JobId"},
			"create-address":                 {"Address"},
			"create-cluster":                 {"AddressId", "JobType", "ShippingOption", "SnowballType"},
			"create-job":                     {},
			"create-long-term-pricing":       {"LongTermPricingType", "SnowballType"},
			"create-return-shipping-label":   {"JobId"},
			"describe-address":               {"AddressId"},
			"describe-addresses":             {},
			"describe-cluster":               {"ClusterId"},
			"describe-job":                   {"JobId"},
			"describe-return-shipping-label": {"JobId"},
			"get-job-manifest":               {"JobId"},
			"get-job-unlock-code":            {"JobId"},
			"get-snowball-usage":             {},
			"get-software-updates":           {"JobId"},
			"list-cluster-jobs":              {"ClusterId"},
			"list-clusters":                  {},
			"list-compatible-images":         {},
			"list-jobs":                      {},
			"list-long-term-pricing":         {},
			"list-pickup-locations":          {},
			"list-service-versions":          {"ServiceName"},
			"update-cluster":                 {"ClusterId"},
			"update-job":                     {"JobId"},
			"update-job-shipment-state":      {"JobId", "ShipmentState"},
			"update-long-term-pricing":       {"LongTermPricingId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("snowball", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
