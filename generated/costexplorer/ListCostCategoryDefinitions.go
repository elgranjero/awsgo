package costexplorer

// ListCostCategoryDefinitions is generated as a reference stub.
// Executable command wiring lives under cmd/costexplorer.go.
//
// Returns the name, Amazon Resource Name (ARN), NumberOfRules and effective dates
// of all cost categories defined in the account. You have the option to use
// EffectiveOn and SupportedResourceTypes to return a list of cost categories that
// were active on a specific date. If there is no EffectiveOn specified, you’ll
// see cost categories that are effective on the current date. If cost category is
// still effective, EffectiveEnd is omitted in the response.
// ListCostCategoryDefinitions supports pagination. The request can have a
// MaxResults range up to 100.
