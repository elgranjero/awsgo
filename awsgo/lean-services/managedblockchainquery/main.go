package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/managedblockchainquery"
)

var fields_batch_get_token_balance = []leanruntime.Field{
	{Name: "GetTokenBalanceInputs", Flag: "get-token-balance-inputs", Type: "[]types.BatchGetTokenBalanceInputItem", Required: false},
}

var fields_get_asset_contract = []leanruntime.Field{
	{Name: "ContractIdentifier", Flag: "contract-identifier", Type: "*types.ContractIdentifier", Required: true},
}

var fields_get_token_balance = []leanruntime.Field{
	{Name: "AtBlockchainInstant", Flag: "at-blockchain-instant", Type: "*types.BlockchainInstant", Required: false},
	{Name: "OwnerIdentifier", Flag: "owner-identifier", Type: "*types.OwnerIdentifier", Required: true},
	{Name: "TokenIdentifier", Flag: "token-identifier", Type: "*types.TokenIdentifier", Required: true},
}

var fields_get_transaction = []leanruntime.Field{
	{Name: "Network", Flag: "network", Type: "types.QueryNetwork", Required: true},
	{Name: "TransactionHash", Flag: "transaction-hash", Type: "*string", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_list_asset_contracts = []leanruntime.Field{
	{Name: "ContractFilter", Flag: "contract-filter", Type: "*types.ContractFilter", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_filtered_transaction_events = []leanruntime.Field{
	{Name: "AddressIdentifierFilter", Flag: "address-identifier-filter", Type: "*types.AddressIdentifierFilter", Required: true},
	{Name: "ConfirmationStatusFilter", Flag: "confirmation-status-filter", Type: "*types.ConfirmationStatusFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Network", Flag: "network", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListFilteredTransactionEventsSort", Required: false},
	{Name: "TimeFilter", Flag: "time-filter", Type: "*types.TimeFilter", Required: false},
	{Name: "VoutFilter", Flag: "vout-filter", Type: "*types.VoutFilter", Required: false},
}

var fields_list_token_balances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwnerFilter", Flag: "owner-filter", Type: "*types.OwnerFilter", Required: false},
	{Name: "TokenFilter", Flag: "token-filter", Type: "*types.TokenFilter", Required: true},
}

var fields_list_transaction_events = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Network", Flag: "network", Type: "types.QueryNetwork", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TransactionHash", Flag: "transaction-hash", Type: "*string", Required: false},
	{Name: "TransactionId", Flag: "transaction-id", Type: "*string", Required: false},
}

var fields_list_transactions = []leanruntime.Field{
	{Name: "Address", Flag: "address", Type: "*string", Required: true},
	{Name: "ConfirmationStatusFilter", Flag: "confirmation-status-filter", Type: "*types.ConfirmationStatusFilter", Required: false},
	{Name: "FromBlockchainInstant", Flag: "from-blockchain-instant", Type: "*types.BlockchainInstant", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Network", Flag: "network", Type: "types.QueryNetwork", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.ListTransactionsSort", Required: false},
	{Name: "ToBlockchainInstant", Flag: "to-blockchain-instant", Type: "*types.BlockchainInstant", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"batch-get-token-balance": {
			Name:   "batch-get-token-balance",
			Fields: fields_batch_get_token_balance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetTokenBalanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_token_balance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetTokenBalance(ctx, input)
			},
		},
		"get-asset-contract": {
			Name:   "get-asset-contract",
			Fields: fields_get_asset_contract,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetContractInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_asset_contract, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssetContract(ctx, input)
			},
		},
		"get-token-balance": {
			Name:   "get-token-balance",
			Fields: fields_get_token_balance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTokenBalanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_token_balance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTokenBalance(ctx, input)
			},
		},
		"get-transaction": {
			Name:   "get-transaction",
			Fields: fields_get_transaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransaction(ctx, input)
			},
		},
		"list-asset-contracts": {
			Name:   "list-asset-contracts",
			Fields: fields_list_asset_contracts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetContractsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_contracts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetContracts(ctx, input)
				}
				var results []*svc.ListAssetContractsOutput
				p := svc.NewListAssetContractsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-filtered-transaction-events": {
			Name:   "list-filtered-transaction-events",
			Fields: fields_list_filtered_transaction_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFilteredTransactionEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_filtered_transaction_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFilteredTransactionEvents(ctx, input)
				}
				var results []*svc.ListFilteredTransactionEventsOutput
				p := svc.NewListFilteredTransactionEventsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-token-balances": {
			Name:   "list-token-balances",
			Fields: fields_list_token_balances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTokenBalancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_token_balances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTokenBalances(ctx, input)
				}
				var results []*svc.ListTokenBalancesOutput
				p := svc.NewListTokenBalancesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-transaction-events": {
			Name:   "list-transaction-events",
			Fields: fields_list_transaction_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTransactionEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_transaction_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTransactionEvents(ctx, input)
				}
				var results []*svc.ListTransactionEventsOutput
				p := svc.NewListTransactionEventsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-transactions": {
			Name:   "list-transactions",
			Fields: fields_list_transactions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTransactionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_transactions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTransactions(ctx, input)
				}
				var results []*svc.ListTransactionsOutput
				p := svc.NewListTransactionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
	}
	if err := leanruntime.Execute("managedblockchainquery", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
