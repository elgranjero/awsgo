package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// managedblockchainqueryCmd represents the managedblockchainquery command
var _managedblockchainqueryCmd = &cobra.Command{
	Use:   "managedblockchainquery",
	Short: "AWS managedblockchainquery CLI",
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
		client := managedblockchainquery.NewFromConfig(cfg)
		if _managedblockchainqueryBatchGetTokenBalance {
			managedblockchainquery_BatchGetTokenBalance(cfg, client)
			return
		}
		if _managedblockchainqueryGetAssetContract {
			managedblockchainquery_GetAssetContract(cfg, client)
			return
		}
		if _managedblockchainqueryGetTokenBalance {
			managedblockchainquery_GetTokenBalance(cfg, client)
			return
		}
		if _managedblockchainqueryGetTransaction {
			managedblockchainquery_GetTransaction(cfg, client)
			return
		}
		if _managedblockchainqueryListAssetContracts {
			managedblockchainquery_ListAssetContracts(cfg, client)
			return
		}
		if _managedblockchainqueryListFilteredTransactionEvents {
			managedblockchainquery_ListFilteredTransactionEvents(cfg, client)
			return
		}
		if _managedblockchainqueryListTokenBalances {
			managedblockchainquery_ListTokenBalances(cfg, client)
			return
		}
		if _managedblockchainqueryListTransactionEvents {
			managedblockchainquery_ListTransactionEvents(cfg, client)
			return
		}
		if _managedblockchainqueryListTransactions {
			managedblockchainquery_ListTransactions(cfg, client)
			return
		}

	},
}

var (
	_managedblockchainqueryBatchGetTokenBalance          bool
	_managedblockchainqueryGetAssetContract              bool
	_managedblockchainqueryGetTokenBalance               bool
	_managedblockchainqueryGetTransaction                bool
	_managedblockchainqueryListAssetContracts            bool
	_managedblockchainqueryListFilteredTransactionEvents bool
	_managedblockchainqueryListTokenBalances             bool
	_managedblockchainqueryListTransactionEvents         bool
	_managedblockchainqueryListTransactions              bool

	_managedblockchainqueryAddress                  string
	_managedblockchainqueryAddressIdentifierFilter  string
	_managedblockchainqueryAtBlockchainInstant      string
	_managedblockchainqueryConfirmationStatusFilter string
	_managedblockchainqueryContractFilter           string
	_managedblockchainqueryContractIdentifier       string
	_managedblockchainqueryFromBlockchainInstant    string
	_managedblockchainqueryGetTokenBalanceInputs    string
	_managedblockchainqueryMaxResults               string
	_managedblockchainqueryNetwork                  string
	_managedblockchainqueryNextToken                string
	_managedblockchainqueryOwnerFilter              string
	_managedblockchainqueryOwnerIdentifier          string
	_managedblockchainquerySort                     string
	_managedblockchainqueryTimeFilter               string
	_managedblockchainqueryToBlockchainInstant      string
	_managedblockchainqueryTokenFilter              string
	_managedblockchainqueryTokenIdentifier          string
	_managedblockchainqueryTransactionHash          string
	_managedblockchainqueryTransactionId            string
	_managedblockchainqueryVoutFilter               string
)

// Gets the token balance for a batch of tokens by using the BatchGetTokenBalance
// action for every token in the request.
//
// Only the native tokens BTC and ETH, and the ERC-20, ERC-721, and ERC 1155 token
// standards are supported.
func managedblockchainquery_BatchGetTokenBalance(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.BatchGetTokenBalanceInput{}

	if len(_managedblockchainqueryGetTokenBalanceInputs) > 0 {
		if err := assignInputField(input, "GetTokenBalanceInputs", _managedblockchainqueryGetTokenBalanceInputs); err != nil {
			log.Errorf("invalid --get-token-balance-inputs: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetTokenBalance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the information about a specific contract deployed on the blockchain.
// - The Bitcoin blockchain networks do not support this operation.
//
// - Metadata is currently only available for some ERC-20 contracts. Metadata
// will be available for additional contracts in the future.
func managedblockchainquery_GetAssetContract(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.GetAssetContractInput{
		// ContractIdentifier: *types.ContractIdentifier, // Required
	}

	if len(_managedblockchainqueryContractIdentifier) > 0 {
		if err := assignInputField(input, "ContractIdentifier", _managedblockchainqueryContractIdentifier); err != nil {
			log.Errorf("invalid --contract-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAssetContract(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the balance of a specific token, including native tokens, for a given
// address (wallet or contract) on the blockchain.
//
// Only the native tokens BTC and ETH, and the ERC-20, ERC-721, and ERC 1155 token
// standards are supported.
func managedblockchainquery_GetTokenBalance(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.GetTokenBalanceInput{
		// OwnerIdentifier: *types.OwnerIdentifier, // Required
		// TokenIdentifier: *types.TokenIdentifier, // Required
	}

	if len(_managedblockchainqueryOwnerIdentifier) > 0 {
		if err := assignInputField(input, "OwnerIdentifier", _managedblockchainqueryOwnerIdentifier); err != nil {
			log.Errorf("invalid --owner-identifier: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryTokenIdentifier) > 0 {
		if err := assignInputField(input, "TokenIdentifier", _managedblockchainqueryTokenIdentifier); err != nil {
			log.Errorf("invalid --token-identifier: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryAtBlockchainInstant) > 0 {
		if err := assignInputField(input, "AtBlockchainInstant", _managedblockchainqueryAtBlockchainInstant); err != nil {
			log.Errorf("invalid --at-blockchain-instant: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTokenBalance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the details of a transaction.
// This action will return transaction details for all transactions that are
// confirmed on the blockchain, even if they have not reached [finality].
//
// [finality]: https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality
func managedblockchainquery_GetTransaction(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.GetTransactionInput{
		// Network: types.QueryNetwork, // Required
	}

	if len(_managedblockchainqueryNetwork) > 0 {
		if err := assignInputField(input, "Network", _managedblockchainqueryNetwork); err != nil {
			log.Errorf("invalid --network: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryTransactionHash) > 0 {
		input.TransactionHash = aws.String(_managedblockchainqueryTransactionHash)
	}
	if len(_managedblockchainqueryTransactionId) > 0 {
		input.TransactionId = aws.String(_managedblockchainqueryTransactionId)
	}

	if resp, err := client.GetTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the contracts for a given contract type deployed by an address
// (either a contract address or a wallet address).
//
// The Bitcoin blockchain networks do not support this operation.
func managedblockchainquery_ListAssetContracts(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.ListAssetContractsInput{
		// ContractFilter: *types.ContractFilter, // Required
	}

	if len(_managedblockchainqueryContractFilter) > 0 {
		if err := assignInputField(input, "ContractFilter", _managedblockchainqueryContractFilter); err != nil {
			log.Errorf("invalid --contract-filter: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainqueryNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssetContracts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchainquery.ListAssetContractsOutput
	p := managedblockchainquery.NewListAssetContractsPaginator(client, input)
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

// Lists all the transaction events for an address on the blockchain.
// This operation is only supported on the Bitcoin networks.
func managedblockchainquery_ListFilteredTransactionEvents(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.ListFilteredTransactionEventsInput{
		// AddressIdentifierFilter: *types.AddressIdentifierFilter, // Required
		// Network: *string, // Required
	}

	if len(_managedblockchainqueryAddressIdentifierFilter) > 0 {
		if err := assignInputField(input, "AddressIdentifierFilter", _managedblockchainqueryAddressIdentifierFilter); err != nil {
			log.Errorf("invalid --address-identifier-filter: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryNetwork) > 0 {
		input.Network = aws.String(_managedblockchainqueryNetwork)
	}
	if len(_managedblockchainqueryConfirmationStatusFilter) > 0 {
		if err := assignInputField(input, "ConfirmationStatusFilter", _managedblockchainqueryConfirmationStatusFilter); err != nil {
			log.Errorf("invalid --confirmation-status-filter: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainqueryNextToken)
	}
	if len(_managedblockchainquerySort) > 0 {
		if err := assignInputField(input, "Sort", _managedblockchainquerySort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryTimeFilter) > 0 {
		if err := assignInputField(input, "TimeFilter", _managedblockchainqueryTimeFilter); err != nil {
			log.Errorf("invalid --time-filter: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryVoutFilter) > 0 {
		if err := assignInputField(input, "VoutFilter", _managedblockchainqueryVoutFilter); err != nil {
			log.Errorf("invalid --vout-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFilteredTransactionEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchainquery.ListFilteredTransactionEventsOutput
	p := managedblockchainquery.NewListFilteredTransactionEventsPaginator(client, input)
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

// This action returns the following for a given blockchain network:
// - Lists all token balances owned by an address (either a contract address or
// a wallet address).
//
// - Lists all token balances for all tokens created by a contract.
//
// - Lists all token balances for a given token.
//
// You must always specify the network property of the tokenFilter when using this
// operation.
func managedblockchainquery_ListTokenBalances(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.ListTokenBalancesInput{
		// TokenFilter: *types.TokenFilter, // Required
	}

	if len(_managedblockchainqueryTokenFilter) > 0 {
		if err := assignInputField(input, "TokenFilter", _managedblockchainqueryTokenFilter); err != nil {
			log.Errorf("invalid --token-filter: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainqueryNextToken)
	}
	if len(_managedblockchainqueryOwnerFilter) > 0 {
		if err := assignInputField(input, "OwnerFilter", _managedblockchainqueryOwnerFilter); err != nil {
			log.Errorf("invalid --owner-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTokenBalances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchainquery.ListTokenBalancesOutput
	p := managedblockchainquery.NewListTokenBalancesPaginator(client, input)
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

// Lists all the transaction events for a transaction
// This action will return transaction details for all transactions that are
// confirmed on the blockchain, even if they have not reached [finality].
//
// [finality]: https://docs.aws.amazon.com/managed-blockchain/latest/ambq-dg/key-concepts.html#finality
func managedblockchainquery_ListTransactionEvents(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.ListTransactionEventsInput{
		// Network: types.QueryNetwork, // Required
	}

	if len(_managedblockchainqueryNetwork) > 0 {
		if err := assignInputField(input, "Network", _managedblockchainqueryNetwork); err != nil {
			log.Errorf("invalid --network: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainqueryNextToken)
	}
	if len(_managedblockchainqueryTransactionHash) > 0 {
		input.TransactionHash = aws.String(_managedblockchainqueryTransactionHash)
	}
	if len(_managedblockchainqueryTransactionId) > 0 {
		input.TransactionId = aws.String(_managedblockchainqueryTransactionId)
	}

	if disablePaginator() {
		if resp, err := client.ListTransactionEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchainquery.ListTransactionEventsOutput
	p := managedblockchainquery.NewListTransactionEventsPaginator(client, input)
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

// Lists all the transaction events for a transaction.
func managedblockchainquery_ListTransactions(cfg aws.Config, client *managedblockchainquery.Client) {
	input := &managedblockchainquery.ListTransactionsInput{
		// Address: *string, // Required
		// Network: types.QueryNetwork, // Required
	}

	if len(_managedblockchainqueryAddress) > 0 {
		input.Address = aws.String(_managedblockchainqueryAddress)
	}
	if len(_managedblockchainqueryNetwork) > 0 {
		if err := assignInputField(input, "Network", _managedblockchainqueryNetwork); err != nil {
			log.Errorf("invalid --network: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryConfirmationStatusFilter) > 0 {
		if err := assignInputField(input, "ConfirmationStatusFilter", _managedblockchainqueryConfirmationStatusFilter); err != nil {
			log.Errorf("invalid --confirmation-status-filter: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryFromBlockchainInstant) > 0 {
		if err := assignInputField(input, "FromBlockchainInstant", _managedblockchainqueryFromBlockchainInstant); err != nil {
			log.Errorf("invalid --from-blockchain-instant: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainqueryMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainqueryNextToken)
	}
	if len(_managedblockchainquerySort) > 0 {
		if err := assignInputField(input, "Sort", _managedblockchainquerySort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainqueryToBlockchainInstant) > 0 {
		if err := assignInputField(input, "ToBlockchainInstant", _managedblockchainqueryToBlockchainInstant); err != nil {
			log.Errorf("invalid --to-blockchain-instant: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTransactions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchainquery.ListTransactionsOutput
	p := managedblockchainquery.NewListTransactionsPaginator(client, input)
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

func init() {
	_rootCmd.AddCommand(_managedblockchainqueryCmd)
	_managedblockchainqueryCmd.Flags().SortFlags = false

	_managedblockchainqueryCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_managedblockchainqueryCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_managedblockchainqueryCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryAddress, "address", "", "", "Address")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryAddressIdentifierFilter, "address-identifier-filter", "", "", "Address Identifier Filter")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryAtBlockchainInstant, "at-blockchain-instant", "", "", "At Blockchain Instant")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryConfirmationStatusFilter, "confirmation-status-filter", "", "", "Confirmation Status Filter")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryContractFilter, "contract-filter", "", "", "Contract Filter")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryContractIdentifier, "contract-identifier", "", "", "Contract Identifier")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryFromBlockchainInstant, "from-blockchain-instant", "", "", "From Blockchain Instant")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryGetTokenBalanceInputs, "get-token-balance-inputs", "", "", "Get Token Balance Inputs")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryMaxResults, "max-results", "", "", "Max Results")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryNetwork, "network", "", "", "Network")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryNextToken, "next-token", "", "", "Next Token")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryOwnerFilter, "owner-filter", "", "", "Owner Filter")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryOwnerIdentifier, "owner-identifier", "", "", "Owner Identifier")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainquerySort, "sort", "", "", "Sort")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryTimeFilter, "time-filter", "", "", "Time Filter")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryToBlockchainInstant, "to-blockchain-instant", "", "", "To Blockchain Instant")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryTokenFilter, "token-filter", "", "", "Token Filter")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryTokenIdentifier, "token-identifier", "", "", "Token Identifier")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryTransactionHash, "transaction-hash", "", "", "Transaction Hash")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryTransactionId, "transaction-id", "", "", "Transaction ID")
	_managedblockchainqueryCmd.Flags().StringVarP(&_managedblockchainqueryVoutFilter, "vout-filter", "", "", "Vout Filter")

	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryBatchGetTokenBalance, "batch-get-token-balance", "", false, "Batch Get Token Balance")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryGetAssetContract, "get-asset-contract", "", false, "Get Asset Contract")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryGetTokenBalance, "get-token-balance", "", false, "Get Token Balance")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryGetTransaction, "get-transaction", "", false, "Get Transaction")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryListAssetContracts, "list-asset-contracts", "", false, "List Asset Contracts")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryListFilteredTransactionEvents, "list-filtered-transaction-events", "", false, "List Filtered Transaction Events")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryListTokenBalances, "list-token-balances", "", false, "List Token Balances")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryListTransactionEvents, "list-transaction-events", "", false, "List Transaction Events")
	_managedblockchainqueryCmd.Flags().BoolVarP(&_managedblockchainqueryListTransactions, "list-transactions", "", false, "List Transactions")

}
