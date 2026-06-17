package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfrontkeyvaluestore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudfrontkeyvaluestoreCmd represents the cloudfrontkeyvaluestore command
var _cloudfrontkeyvaluestoreCmd = &cobra.Command{
	Use:   "cloudfrontkeyvaluestore",
	Short: "AWS cloudfrontkeyvaluestore CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := cloudfrontkeyvaluestore.NewFromConfig(cfg)
		if _cloudfrontkeyvaluestoreDeleteKey {
			cloudfrontkeyvaluestore_DeleteKey(cfg, client)
			return
		}
		if _cloudfrontkeyvaluestoreDescribeKeyValueStore {
			cloudfrontkeyvaluestore_DescribeKeyValueStore(cfg, client)
			return
		}
		if _cloudfrontkeyvaluestoreGetKey {
			cloudfrontkeyvaluestore_GetKey(cfg, client)
			return
		}
		if _cloudfrontkeyvaluestoreListKeys {
			cloudfrontkeyvaluestore_ListKeys(cfg, client)
			return
		}
		if _cloudfrontkeyvaluestorePutKey {
			cloudfrontkeyvaluestore_PutKey(cfg, client)
			return
		}
		if _cloudfrontkeyvaluestoreUpdateKeys {
			cloudfrontkeyvaluestore_UpdateKeys(cfg, client)
			return
		}

	},
}

var (
	_cloudfrontkeyvaluestoreDeleteKey             bool
	_cloudfrontkeyvaluestoreDescribeKeyValueStore bool
	_cloudfrontkeyvaluestoreGetKey                bool
	_cloudfrontkeyvaluestoreListKeys              bool
	_cloudfrontkeyvaluestorePutKey                bool
	_cloudfrontkeyvaluestoreUpdateKeys            bool

	_cloudfrontkeyvaluestoreDeletes    string
	_cloudfrontkeyvaluestoreIfMatch    string
	_cloudfrontkeyvaluestoreKey        string
	_cloudfrontkeyvaluestoreKvsARN     string
	_cloudfrontkeyvaluestoreMaxResults string
	_cloudfrontkeyvaluestoreNextToken  string
	_cloudfrontkeyvaluestorePuts       string
	_cloudfrontkeyvaluestoreValue      string
)

// Deletes the key value pair specified by the key.
func cloudfrontkeyvaluestore_DeleteKey(cfg aws.Config, client *cloudfrontkeyvaluestore.Client) {
	input := &cloudfrontkeyvaluestore.DeleteKeyInput{
		// IfMatch: *string, // Required
		// Key: *string, // Required
		// KvsARN: *string, // Required
	}

	if len(_cloudfrontkeyvaluestoreIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontkeyvaluestoreIfMatch)
	}
	if len(_cloudfrontkeyvaluestoreKey) > 0 {
		input.Key = aws.String(_cloudfrontkeyvaluestoreKey)
	}
	if len(_cloudfrontkeyvaluestoreKvsARN) > 0 {
		input.KvsARN = aws.String(_cloudfrontkeyvaluestoreKvsARN)
	}

	if resp, err := client.DeleteKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns metadata information about Key Value Store.
func cloudfrontkeyvaluestore_DescribeKeyValueStore(cfg aws.Config, client *cloudfrontkeyvaluestore.Client) {
	input := &cloudfrontkeyvaluestore.DescribeKeyValueStoreInput{
		// KvsARN: *string, // Required
	}

	if len(_cloudfrontkeyvaluestoreKvsARN) > 0 {
		input.KvsARN = aws.String(_cloudfrontkeyvaluestoreKvsARN)
	}

	if resp, err := client.DescribeKeyValueStore(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a key value pair.
func cloudfrontkeyvaluestore_GetKey(cfg aws.Config, client *cloudfrontkeyvaluestore.Client) {
	input := &cloudfrontkeyvaluestore.GetKeyInput{
		// Key: *string, // Required
		// KvsARN: *string, // Required
	}

	if len(_cloudfrontkeyvaluestoreKey) > 0 {
		input.Key = aws.String(_cloudfrontkeyvaluestoreKey)
	}
	if len(_cloudfrontkeyvaluestoreKvsARN) > 0 {
		input.KvsARN = aws.String(_cloudfrontkeyvaluestoreKvsARN)
	}

	if resp, err := client.GetKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of key value pairs.
func cloudfrontkeyvaluestore_ListKeys(cfg aws.Config, client *cloudfrontkeyvaluestore.Client) {
	input := &cloudfrontkeyvaluestore.ListKeysInput{
		// KvsARN: *string, // Required
	}

	if len(_cloudfrontkeyvaluestoreKvsARN) > 0 {
		input.KvsARN = aws.String(_cloudfrontkeyvaluestoreKvsARN)
	}
	if len(_cloudfrontkeyvaluestoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudfrontkeyvaluestoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontkeyvaluestoreNextToken) > 0 {
		input.NextToken = aws.String(_cloudfrontkeyvaluestoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudfrontkeyvaluestore.ListKeysOutput
	p := cloudfrontkeyvaluestore.NewListKeysPaginator(client, input)
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

// Creates a new key value pair or replaces the value of an existing key.
func cloudfrontkeyvaluestore_PutKey(cfg aws.Config, client *cloudfrontkeyvaluestore.Client) {
	input := &cloudfrontkeyvaluestore.PutKeyInput{
		// IfMatch: *string, // Required
		// Key: *string, // Required
		// KvsARN: *string, // Required
		// Value: *string, // Required
	}

	if len(_cloudfrontkeyvaluestoreIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontkeyvaluestoreIfMatch)
	}
	if len(_cloudfrontkeyvaluestoreKey) > 0 {
		input.Key = aws.String(_cloudfrontkeyvaluestoreKey)
	}
	if len(_cloudfrontkeyvaluestoreKvsARN) > 0 {
		input.KvsARN = aws.String(_cloudfrontkeyvaluestoreKvsARN)
	}
	if len(_cloudfrontkeyvaluestoreValue) > 0 {
		input.Value = aws.String(_cloudfrontkeyvaluestoreValue)
	}

	if resp, err := client.PutKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Puts or Deletes multiple key value pairs in a single, all-or-nothing operation.
func cloudfrontkeyvaluestore_UpdateKeys(cfg aws.Config, client *cloudfrontkeyvaluestore.Client) {
	input := &cloudfrontkeyvaluestore.UpdateKeysInput{
		// IfMatch: *string, // Required
		// KvsARN: *string, // Required
	}

	if len(_cloudfrontkeyvaluestoreIfMatch) > 0 {
		input.IfMatch = aws.String(_cloudfrontkeyvaluestoreIfMatch)
	}
	if len(_cloudfrontkeyvaluestoreKvsARN) > 0 {
		input.KvsARN = aws.String(_cloudfrontkeyvaluestoreKvsARN)
	}
	if len(_cloudfrontkeyvaluestoreDeletes) > 0 {
		if err := assignInputField(input, "Deletes", _cloudfrontkeyvaluestoreDeletes); err != nil {
			log.Errorf("invalid --deletes: %s", err.Error())
			return
		}
	}
	if len(_cloudfrontkeyvaluestorePuts) > 0 {
		if err := assignInputField(input, "Puts", _cloudfrontkeyvaluestorePuts); err != nil {
			log.Errorf("invalid --puts: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudfrontkeyvaluestoreCmd)
	_cloudfrontkeyvaluestoreCmd.Flags().SortFlags = false

	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreDeletes, "deletes", "", "", "Deletes")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreIfMatch, "if-match", "", "", "If Match")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreKey, "key", "", "", "Key")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreKvsARN, "kvs-arn", "", "", "Kvs ARN")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreMaxResults, "max-results", "", "", "Max Results")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreNextToken, "next-token", "", "", "Next Token")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestorePuts, "puts", "", "", "Puts")
	_cloudfrontkeyvaluestoreCmd.Flags().StringVarP(&_cloudfrontkeyvaluestoreValue, "value", "", "", "Value")

	_cloudfrontkeyvaluestoreCmd.Flags().BoolVarP(&_cloudfrontkeyvaluestoreDeleteKey, "delete-key", "", false, "Delete Key")
	_cloudfrontkeyvaluestoreCmd.Flags().BoolVarP(&_cloudfrontkeyvaluestoreDescribeKeyValueStore, "describe-key-value-store", "", false, "Describe Key Value Store")
	_cloudfrontkeyvaluestoreCmd.Flags().BoolVarP(&_cloudfrontkeyvaluestoreGetKey, "get-key", "", false, "Get Key")
	_cloudfrontkeyvaluestoreCmd.Flags().BoolVarP(&_cloudfrontkeyvaluestoreListKeys, "list-keys", "", false, "List Keys")
	_cloudfrontkeyvaluestoreCmd.Flags().BoolVarP(&_cloudfrontkeyvaluestorePutKey, "put-key", "", false, "Put Key")
	_cloudfrontkeyvaluestoreCmd.Flags().BoolVarP(&_cloudfrontkeyvaluestoreUpdateKeys, "update-keys", "", false, "Update Keys")

}
