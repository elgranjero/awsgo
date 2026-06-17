package wafv2

// PutManagedRuleSetVersions is generated as a reference stub.
// Executable command wiring lives under cmd/wafv2.go.
//
// Defines the versions of your managed rule set that you are offering to the
// customers. Customers see your offerings as managed rule groups with versioning.
//
// This is intended for use only by vendors of managed rule sets. Vendors are
// Amazon Web Services and Amazon Web Services Marketplace sellers.
//
// Vendors, you can use the managed rule set APIs to provide controlled rollout of
// your versioned managed rule group offerings for your customers. The APIs are
// ListManagedRuleSets , GetManagedRuleSet , PutManagedRuleSetVersions , and
// UpdateManagedRuleSetVersionExpiryDate .
//
// Customers retrieve their managed rule group list by calling ListAvailableManagedRuleGroups. The name that you
// provide here for your managed rule set is the name the customer sees for the
// corresponding managed rule group. Customers can retrieve the available versions
// for a managed rule group by calling ListAvailableManagedRuleGroupVersions. You provide a rule group specification
// for each version. For each managed rule set, you must specify a version that you
// recommend using.
//
// To initiate the expiration of a managed rule group version, use UpdateManagedRuleSetVersionExpiryDate.
