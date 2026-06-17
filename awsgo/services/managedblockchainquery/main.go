package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/managedblockchainquery/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"batch-get-token-balance", "get-asset-contract", "get-token-balance", "get-transaction", "list-asset-contracts", "list-filtered-transaction-events", "list-token-balances", "list-transaction-events", "list-transactions"},
		OperationSet: map[string]bool{"batch-get-token-balance": true, "get-asset-contract": true, "get-token-balance": true, "get-transaction": true, "list-asset-contracts": true, "list-filtered-transaction-events": true, "list-token-balances": true, "list-transaction-events": true, "list-transactions": true},
		OperationInputs: map[string][]string{
			"batch-get-token-balance":          {"GetTokenBalanceInputs"},
			"get-asset-contract":               {"ContractIdentifier"},
			"get-token-balance":                {"AtBlockchainInstant", "OwnerIdentifier", "TokenIdentifier"},
			"get-transaction":                  {"Network", "TransactionHash", "TransactionId"},
			"list-asset-contracts":             {"ContractFilter", "MaxResults", "NextToken"},
			"list-filtered-transaction-events": {"AddressIdentifierFilter", "ConfirmationStatusFilter", "MaxResults", "Network", "NextToken", "Sort", "TimeFilter", "VoutFilter"},
			"list-token-balances":              {"MaxResults", "NextToken", "OwnerFilter", "TokenFilter"},
			"list-transaction-events":          {"MaxResults", "Network", "NextToken", "TransactionHash", "TransactionId"},
			"list-transactions":                {"Address", "ConfirmationStatusFilter", "FromBlockchainInstant", "MaxResults", "Network", "NextToken", "Sort", "ToBlockchainInstant"},
		},
		OperationInputTypes: map[string]map[string]string{
			"batch-get-token-balance":          {"GetTokenBalanceInputs": "[]types.BatchGetTokenBalanceInputItem"},
			"get-asset-contract":               {"ContractIdentifier": "*types.ContractIdentifier"},
			"get-token-balance":                {"AtBlockchainInstant": "*types.BlockchainInstant", "OwnerIdentifier": "*types.OwnerIdentifier", "TokenIdentifier": "*types.TokenIdentifier"},
			"get-transaction":                  {"Network": "types.QueryNetwork", "TransactionHash": "*string", "TransactionId": "*string"},
			"list-asset-contracts":             {"ContractFilter": "*types.ContractFilter", "MaxResults": "*int32", "NextToken": "*string"},
			"list-filtered-transaction-events": {"AddressIdentifierFilter": "*types.AddressIdentifierFilter", "ConfirmationStatusFilter": "*types.ConfirmationStatusFilter", "MaxResults": "*int32", "Network": "*string", "NextToken": "*string", "Sort": "*types.ListFilteredTransactionEventsSort", "TimeFilter": "*types.TimeFilter", "VoutFilter": "*types.VoutFilter"},
			"list-token-balances":              {"MaxResults": "*int32", "NextToken": "*string", "OwnerFilter": "*types.OwnerFilter", "TokenFilter": "*types.TokenFilter"},
			"list-transaction-events":          {"MaxResults": "*int32", "Network": "types.QueryNetwork", "NextToken": "*string", "TransactionHash": "*string", "TransactionId": "*string"},
			"list-transactions":                {"Address": "*string", "ConfirmationStatusFilter": "*types.ConfirmationStatusFilter", "FromBlockchainInstant": "*types.BlockchainInstant", "MaxResults": "*int32", "Network": "types.QueryNetwork", "NextToken": "*string", "Sort": "*types.ListTransactionsSort", "ToBlockchainInstant": "*types.BlockchainInstant"},
		},
		OperationInputRequired: map[string][]string{
			"batch-get-token-balance":          {},
			"get-asset-contract":               {"ContractIdentifier"},
			"get-token-balance":                {"OwnerIdentifier", "TokenIdentifier"},
			"get-transaction":                  {"Network"},
			"list-asset-contracts":             {"ContractFilter"},
			"list-filtered-transaction-events": {"AddressIdentifierFilter", "Network"},
			"list-token-balances":              {"TokenFilter"},
			"list-transaction-events":          {"Network"},
			"list-transactions":                {"Address", "Network"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("managedblockchainquery", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
