package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// secretsmanagerCmd represents the secretsmanager command
var _secretsmanagerCmd = &cobra.Command{
	Use:   "secretsmanager",
	Short: "AWS secretsmanager CLI",
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
		client := secretsmanager.NewFromConfig(cfg)
		if _secretsmanagerBatchGetSecretValue {
			secretsmanager_BatchGetSecretValue(cfg, client)
			return
		}
		if _secretsmanagerCancelRotateSecret {
			secretsmanager_CancelRotateSecret(cfg, client)
			return
		}
		if _secretsmanagerCreateSecret {
			secretsmanager_CreateSecret(cfg, client)
			return
		}
		if _secretsmanagerDeleteResourcePolicy {
			secretsmanager_DeleteResourcePolicy(cfg, client)
			return
		}
		if _secretsmanagerDeleteSecret {
			secretsmanager_DeleteSecret(cfg, client)
			return
		}
		if _secretsmanagerDescribeSecret {
			secretsmanager_DescribeSecret(cfg, client)
			return
		}
		if _secretsmanagerGetRandomPassword {
			secretsmanager_GetRandomPassword(cfg, client)
			return
		}
		if _secretsmanagerGetResourcePolicy {
			secretsmanager_GetResourcePolicy(cfg, client)
			return
		}
		if _secretsmanagerGetSecretValue {
			secretsmanager_GetSecretValue(cfg, client)
			return
		}
		if _secretsmanagerListSecretVersionIds {
			secretsmanager_ListSecretVersionIds(cfg, client)
			return
		}
		if _secretsmanagerListSecrets {
			secretsmanager_ListSecrets(cfg, client)
			return
		}
		if _secretsmanagerPutResourcePolicy {
			secretsmanager_PutResourcePolicy(cfg, client)
			return
		}
		if _secretsmanagerPutSecretValue {
			secretsmanager_PutSecretValue(cfg, client)
			return
		}
		if _secretsmanagerRemoveRegionsFromReplication {
			secretsmanager_RemoveRegionsFromReplication(cfg, client)
			return
		}
		if _secretsmanagerReplicateSecretToRegions {
			secretsmanager_ReplicateSecretToRegions(cfg, client)
			return
		}
		if _secretsmanagerRestoreSecret {
			secretsmanager_RestoreSecret(cfg, client)
			return
		}
		if _secretsmanagerRotateSecret {
			secretsmanager_RotateSecret(cfg, client)
			return
		}
		if _secretsmanagerStopReplicationToReplica {
			secretsmanager_StopReplicationToReplica(cfg, client)
			return
		}
		if _secretsmanagerTagResource {
			secretsmanager_TagResource(cfg, client)
			return
		}
		if _secretsmanagerUntagResource {
			secretsmanager_UntagResource(cfg, client)
			return
		}
		if _secretsmanagerUpdateSecret {
			secretsmanager_UpdateSecret(cfg, client)
			return
		}
		if _secretsmanagerUpdateSecretVersionStage {
			secretsmanager_UpdateSecretVersionStage(cfg, client)
			return
		}
		if _secretsmanagerValidateResourcePolicy {
			secretsmanager_ValidateResourcePolicy(cfg, client)
			return
		}

	},
}

var (
	_secretsmanagerBatchGetSecretValue          bool
	_secretsmanagerCancelRotateSecret           bool
	_secretsmanagerCreateSecret                 bool
	_secretsmanagerDeleteResourcePolicy         bool
	_secretsmanagerDeleteSecret                 bool
	_secretsmanagerDescribeSecret               bool
	_secretsmanagerGetRandomPassword            bool
	_secretsmanagerGetResourcePolicy            bool
	_secretsmanagerGetSecretValue               bool
	_secretsmanagerListSecretVersionIds         bool
	_secretsmanagerListSecrets                  bool
	_secretsmanagerPutResourcePolicy            bool
	_secretsmanagerPutSecretValue               bool
	_secretsmanagerRemoveRegionsFromReplication bool
	_secretsmanagerReplicateSecretToRegions     bool
	_secretsmanagerRestoreSecret                bool
	_secretsmanagerRotateSecret                 bool
	_secretsmanagerStopReplicationToReplica     bool
	_secretsmanagerTagResource                  bool
	_secretsmanagerUntagResource                bool
	_secretsmanagerUpdateSecret                 bool
	_secretsmanagerUpdateSecretVersionStage     bool
	_secretsmanagerValidateResourcePolicy       bool

	_secretsmanagerAddReplicaRegions              string
	_secretsmanagerBlockPublicPolicy              string
	_secretsmanagerClientRequestToken             string
	_secretsmanagerDescription                    string
	_secretsmanagerExcludeCharacters              string
	_secretsmanagerExcludeLowercase               string
	_secretsmanagerExcludeNumbers                 string
	_secretsmanagerExcludePunctuation             string
	_secretsmanagerExcludeUppercase               string
	_secretsmanagerExternalSecretRotationMetadata string
	_secretsmanagerExternalSecretRotationRoleArn  string
	_secretsmanagerFilters                        string
	_secretsmanagerForceDeleteWithoutRecovery     string
	_secretsmanagerForceOverwriteReplicaSecret    string
	_secretsmanagerIncludeDeprecated              string
	_secretsmanagerIncludePlannedDeletion         string
	_secretsmanagerIncludeSpace                   string
	_secretsmanagerKmsKeyId                       string
	_secretsmanagerMaxResults                     string
	_secretsmanagerMoveToVersionId                string
	_secretsmanagerName                           string
	_secretsmanagerNextToken                      string
	_secretsmanagerPasswordLength                 string
	_secretsmanagerRecoveryWindowInDays           string
	_secretsmanagerRemoveFromVersionId            string
	_secretsmanagerRemoveReplicaRegions           []string
	_secretsmanagerRequireEachIncludedType        string
	_secretsmanagerResourcePolicy                 string
	_secretsmanagerRotateImmediately              string
	_secretsmanagerRotationLambdaARN              string
	_secretsmanagerRotationRules                  string
	_secretsmanagerRotationToken                  string
	_secretsmanagerSecretBinary                   string
	_secretsmanagerSecretId                       string
	_secretsmanagerSecretIdList                   []string
	_secretsmanagerSecretString                   string
	_secretsmanagerSortBy                         string
	_secretsmanagerSortOrder                      string
	_secretsmanagerTagKeys                        []string
	_secretsmanagerTags                           string
	_secretsmanagerType                           string
	_secretsmanagerVersionId                      string
	_secretsmanagerVersionStage                   string
	_secretsmanagerVersionStages                  []string
)

// Retrieves the contents of the encrypted fields SecretString or SecretBinary for
// up to 20 secrets. To retrieve a single secret, call GetSecretValue.
//
// To choose which secrets to retrieve, you can specify a list of secrets by name
// or ARN, or you can use filters. If Secrets Manager encounters errors such as
// AccessDeniedException while attempting to retrieve any of the secrets, you can
// see the errors in Errors in the response.
//
// Secrets Manager generates CloudTrail GetSecretValue log entries for each secret
// you request when you call this action. Do not include sensitive information in
// request parameters because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:BatchGetSecretValue , and you must have
// secretsmanager:GetSecretValue for each secret. If you use filters, you must also
// have secretsmanager:ListSecrets . If the secrets are encrypted using
// customer-managed keys instead of the Amazon Web Services managed key
// aws/secretsmanager , then you also need kms:Decrypt permissions for the keys.
// For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_BatchGetSecretValue(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.BatchGetSecretValueInput{}

	if len(_secretsmanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _secretsmanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _secretsmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerNextToken) > 0 {
		input.NextToken = aws.String(_secretsmanagerNextToken)
	}
	if len(_secretsmanagerSecretIdList) > 0 {
		input.SecretIdList = append([]string(nil), _secretsmanagerSecretIdList...)
	}

	if disablePaginator() {
		if resp, err := client.BatchGetSecretValue(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*secretsmanager.BatchGetSecretValueOutput
	p := secretsmanager.NewBatchGetSecretValuePaginator(client, input)
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

// Turns off automatic rotation, and if a rotation is currently in progress,
// cancels the rotation.
//
// If you cancel a rotation in progress, it can leave the VersionStage labels in
// an unexpected state. You might need to remove the staging label AWSPENDING from
// the partially created version. You also need to determine whether to roll back
// to the previous version of the secret by moving the staging label AWSCURRENT to
// the version that has AWSPENDING . To determine which version has a specific
// staging label, call ListSecretVersionIds. Then use UpdateSecretVersionStage to change staging labels. For more information,
// see [How rotation works].
//
// To turn on automatic rotation again, call RotateSecret.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:CancelRotateSecret . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [How rotation works]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_CancelRotateSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.CancelRotateSecretInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.CancelRotateSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new secret. A secret can be a password, a set of credentials such as
// a user name and password, an OAuth token, or other secret information that you
// store in an encrypted form in Secrets Manager. The secret also includes the
// connection information to access a database or other service, which Secrets
// Manager doesn't encrypt. A secret in Secrets Manager consists of both the
// protected secret data and the important information needed to manage the secret.
//
// For secrets that use managed rotation, you need to create the secret through
// the managing service. For more information, see [Secrets Manager secrets managed by other Amazon Web Services services].
//
// For information about creating a secret in the console, see [Create a secret].
//
// To create a secret, you can provide the secret value to be encrypted in either
// the SecretString parameter or the SecretBinary parameter, but not both. If you
// include SecretString or SecretBinary then Secrets Manager creates an initial
// secret version and automatically attaches the staging label AWSCURRENT to it.
//
// For database credentials you want to rotate, for Secrets Manager to be able to
// rotate the secret, you must make sure the JSON you store in the SecretString
// matches the [JSON structure of a database secret].
//
// If you don't specify an KMS encryption key, Secrets Manager uses the Amazon Web
// Services managed key aws/secretsmanager . If this key doesn't already exist in
// your account, then Secrets Manager creates it for you automatically. All users
// and roles in the Amazon Web Services account automatically have access to use
// aws/secretsmanager . Creating aws/secretsmanager can result in a one-time
// significant delay in returning the result.
//
// If the secret is in a different Amazon Web Services account from the
// credentials calling the API, then you can't use aws/secretsmanager to encrypt
// the secret, and you must create and use a customer managed KMS key.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary or
// SecretString because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:CreateSecret . If you include tags in the
// secret, you also need secretsmanager:TagResource . To add replica Regions, you
// must also have secretsmanager:ReplicateSecretToRegions . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// To encrypt the secret with a KMS key other than aws/secretsmanager , you need
// kms:GenerateDataKey and kms:Decrypt permission to the key.
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secrets Manager secrets managed by other Amazon Web Services services]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/service-linked-secrets.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [Create a secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_create-basic-secret.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [JSON structure of a database secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_secret_json_structure.html
func secretsmanager_CreateSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.CreateSecretInput{
		// Name: *string, // Required
	}

	if len(_secretsmanagerName) > 0 {
		input.Name = aws.String(_secretsmanagerName)
	}
	if len(_secretsmanagerAddReplicaRegions) > 0 {
		if err := assignInputField(input, "AddReplicaRegions", _secretsmanagerAddReplicaRegions); err != nil {
			log.Errorf("invalid --add-replica-regions: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_secretsmanagerClientRequestToken)
	}
	if len(_secretsmanagerDescription) > 0 {
		input.Description = aws.String(_secretsmanagerDescription)
	}
	if len(_secretsmanagerForceOverwriteReplicaSecret) > 0 {
		if err := assignInputField(input, "ForceOverwriteReplicaSecret", _secretsmanagerForceOverwriteReplicaSecret); err != nil {
			log.Errorf("invalid --force-overwrite-replica-secret: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_secretsmanagerKmsKeyId)
	}
	if len(_secretsmanagerSecretBinary) > 0 {
		if err := assignInputField(input, "SecretBinary", _secretsmanagerSecretBinary); err != nil {
			log.Errorf("invalid --secret-binary: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerSecretString) > 0 {
		input.SecretString = aws.String(_secretsmanagerSecretString)
	}
	if len(_secretsmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _secretsmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerType) > 0 {
		input.Type = aws.String(_secretsmanagerType)
	}

	if resp, err := client.CreateSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based permission policy attached to the secret. To attach
// a policy to a secret, use PutResourcePolicy.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:DeleteResourcePolicy . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_DeleteResourcePolicy(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.DeleteResourcePolicyInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a secret and all of its versions. You can specify a recovery window
// during which you can restore the secret. The minimum recovery window is 7 days.
// The default recovery window is 30 days. Secrets Manager attaches a DeletionDate
// stamp to the secret that specifies the end of the recovery window. At the end of
// the recovery window, Secrets Manager deletes the secret permanently.
//
// You can't delete a primary secret that is replicated to other Regions. You must
// first delete the replicas using RemoveRegionsFromReplication, and then delete the primary secret. When you
// delete a replica, it is deleted immediately.
//
// You can't directly delete a version of a secret. Instead, you remove all
// staging labels from the version using UpdateSecretVersionStage. This marks the version as deprecated,
// and then Secrets Manager can automatically delete the version in the background.
//
// To determine whether an application still uses a secret, you can create an
// Amazon CloudWatch alarm to alert you to any attempts to access a secret during
// the recovery window. For more information, see [Monitor secrets scheduled for deletion].
//
// Secrets Manager performs the permanent secret deletion at the end of the
// waiting period as a background task with low priority. There is no guarantee of
// a specific time after the recovery window for the permanent delete to occur.
//
// At any time before recovery window ends, you can use RestoreSecret to remove the DeletionDate
// and cancel the deletion of the secret.
//
// When a secret is scheduled for deletion, you cannot retrieve the secret value.
// You must first cancel the deletion with RestoreSecretand then you can retrieve the secret.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:DeleteSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Monitor secrets scheduled for deletion]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/monitoring_cloudwatch_deleted-secrets.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_DeleteSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.DeleteSecretInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerForceDeleteWithoutRecovery) > 0 {
		if err := assignInputField(input, "ForceDeleteWithoutRecovery", _secretsmanagerForceDeleteWithoutRecovery); err != nil {
			log.Errorf("invalid --force-delete-without-recovery: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerRecoveryWindowInDays) > 0 {
		if err := assignInputField(input, "RecoveryWindowInDays", _secretsmanagerRecoveryWindowInDays); err != nil {
			log.Errorf("invalid --recovery-window-in-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a secret. It does not include the encrypted secret
// value. Secrets Manager only returns fields that have a value in the response.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:DescribeSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_DescribeSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.DescribeSecretInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.DescribeSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a random password. We recommend that you specify the maximum length
// and include every character type that the system you are generating a password
// for can support. By default, Secrets Manager uses uppercase and lowercase
// letters, numbers, and the following characters in passwords:
// !\"#$%&'()*+,-./:;<=>?(at)[\\]^_`{|}~
//
// Secrets Manager generates a CloudTrail log entry when you call this action.
//
// Required permissions: secretsmanager:GetRandomPassword . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_GetRandomPassword(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.GetRandomPasswordInput{}

	if len(_secretsmanagerExcludeCharacters) > 0 {
		input.ExcludeCharacters = aws.String(_secretsmanagerExcludeCharacters)
	}
	if len(_secretsmanagerExcludeLowercase) > 0 {
		if err := assignInputField(input, "ExcludeLowercase", _secretsmanagerExcludeLowercase); err != nil {
			log.Errorf("invalid --exclude-lowercase: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerExcludeNumbers) > 0 {
		if err := assignInputField(input, "ExcludeNumbers", _secretsmanagerExcludeNumbers); err != nil {
			log.Errorf("invalid --exclude-numbers: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerExcludePunctuation) > 0 {
		if err := assignInputField(input, "ExcludePunctuation", _secretsmanagerExcludePunctuation); err != nil {
			log.Errorf("invalid --exclude-punctuation: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerExcludeUppercase) > 0 {
		if err := assignInputField(input, "ExcludeUppercase", _secretsmanagerExcludeUppercase); err != nil {
			log.Errorf("invalid --exclude-uppercase: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerIncludeSpace) > 0 {
		if err := assignInputField(input, "IncludeSpace", _secretsmanagerIncludeSpace); err != nil {
			log.Errorf("invalid --include-space: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerPasswordLength) > 0 {
		if err := assignInputField(input, "PasswordLength", _secretsmanagerPasswordLength); err != nil {
			log.Errorf("invalid --password-length: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerRequireEachIncludedType) > 0 {
		if err := assignInputField(input, "RequireEachIncludedType", _secretsmanagerRequireEachIncludedType); err != nil {
			log.Errorf("invalid --require-each-included-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRandomPassword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the JSON text of the resource-based policy document attached to the
// secret. For more information about permissions policies attached to a secret,
// see [Permissions policies attached to a secret].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:GetResourcePolicy . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Permissions policies attached to a secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-policies.html
func secretsmanager_GetResourcePolicy(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.GetResourcePolicyInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the contents of the encrypted fields SecretString or SecretBinary
// from the specified version of a secret, whichever contains content.
//
// To retrieve the values for a group of secrets, call BatchGetSecretValue.
//
// We recommend that you cache your secret values by using client-side caching.
// Caching secrets improves speed and reduces your costs. For more information, see
// [Cache secrets for your applications].
//
// To retrieve the previous version of a secret, use VersionStage and specify
// AWSPREVIOUS. To revert to the previous version of a secret, call [UpdateSecretVersionStage].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:GetSecretValue . If the secret is encrypted
// using a customer-managed key instead of the Amazon Web Services managed key
// aws/secretsmanager , then you also need kms:Decrypt permissions for that key.
// For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [UpdateSecretVersionStage]: https://docs.aws.amazon.com/cli/latest/reference/secretsmanager/update-secret-version-stage.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Cache secrets for your applications]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieving-secrets.html
func secretsmanager_GetSecretValue(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.GetSecretValueInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerVersionId) > 0 {
		input.VersionId = aws.String(_secretsmanagerVersionId)
	}
	if len(_secretsmanagerVersionStage) > 0 {
		input.VersionStage = aws.String(_secretsmanagerVersionStage)
	}

	if resp, err := client.GetSecretValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the versions of a secret. Secrets Manager uses staging labels to indicate
// the different versions of a secret. For more information, see [Secrets Manager concepts: Versions].
//
// To list the secrets in the account, use ListSecrets.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ListSecretVersionIds . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secrets Manager concepts: Versions]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_ListSecretVersionIds(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.ListSecretVersionIdsInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerIncludeDeprecated) > 0 {
		if err := assignInputField(input, "IncludeDeprecated", _secretsmanagerIncludeDeprecated); err != nil {
			log.Errorf("invalid --include-deprecated: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _secretsmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerNextToken) > 0 {
		input.NextToken = aws.String(_secretsmanagerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecretVersionIds(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*secretsmanager.ListSecretVersionIdsOutput
	p := secretsmanager.NewListSecretVersionIdsPaginator(client, input)
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

// Lists the secrets that are stored by Secrets Manager in the Amazon Web Services
// account, not including secrets that are marked for deletion. To see secrets
// marked for deletion, use the Secrets Manager console.
//
// All Secrets Manager operations are eventually consistent. ListSecrets might not
// reflect changes from the last five minutes. You can get more recent information
// for a specific secret by calling DescribeSecret.
//
// To list the versions of a secret, use ListSecretVersionIds.
//
// To retrieve the values for the secrets, call BatchGetSecretValue or GetSecretValue.
//
// For information about finding secrets in the console, see [Find secrets in Secrets Manager].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ListSecrets . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Find secrets in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/manage_search-secret.html
func secretsmanager_ListSecrets(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.ListSecretsInput{}

	if len(_secretsmanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _secretsmanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerIncludePlannedDeletion) > 0 {
		if err := assignInputField(input, "IncludePlannedDeletion", _secretsmanagerIncludePlannedDeletion); err != nil {
			log.Errorf("invalid --include-planned-deletion: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _secretsmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerNextToken) > 0 {
		input.NextToken = aws.String(_secretsmanagerNextToken)
	}
	if len(_secretsmanagerSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _secretsmanagerSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _secretsmanagerSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSecrets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*secretsmanager.ListSecretsOutput
	p := secretsmanager.NewListSecretsPaginator(client, input)
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

// Attaches a resource-based permission policy to a secret. A resource-based
// policy is optional. For more information, see [Authentication and access control for Secrets Manager]
//
// For information about attaching a policy in the console, see [Attach a permissions policy to a secret].
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:PutResourcePolicy . For more information,
// see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Attach a permissions policy to a secret]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Authentication and access control for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
func secretsmanager_PutResourcePolicy(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.PutResourcePolicyInput{
		// ResourcePolicy: *string, // Required
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_secretsmanagerResourcePolicy)
	}
	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerBlockPublicPolicy) > 0 {
		if err := assignInputField(input, "BlockPublicPolicy", _secretsmanagerBlockPublicPolicy); err != nil {
			log.Errorf("invalid --block-public-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of your secret by creating a new encrypted value and
// attaching it to the secret. version can contain a new SecretString value or a
// new SecretBinary value.
//
// Do not call PutSecretValue at a sustained rate of more than once every 10
// minutes. When you update the secret value, Secrets Manager creates a new version
// of the secret. Secrets Manager keeps 100 of the most recent versions, but it
// keeps all secret versions created in the last 24 hours. If you call
// PutSecretValue more than once every 10 minutes, you will create more versions
// than Secrets Manager removes, and you will reach the quota for secret versions.
//
// You can specify the staging labels to attach to the new version in VersionStages
// . If you don't include VersionStages , then Secrets Manager automatically moves
// the staging label AWSCURRENT to this version. If this operation creates the
// first version for the secret, then Secrets Manager automatically attaches the
// staging label AWSCURRENT to it. If this operation moves the staging label
// AWSCURRENT from another version to this version, then Secrets Manager also
// automatically moves the staging label AWSPREVIOUS to the version that AWSCURRENT
// was removed from.
//
// This operation is idempotent. If you call this operation with a
// ClientRequestToken that matches an existing version's VersionId, and you specify
// the same secret data, the operation succeeds but does nothing. However, if the
// secret data is different, then the operation fails because you can't modify an
// existing version; you can only create new ones.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary ,
// SecretString , or RotationToken because it might be logged. For more
// information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:PutSecretValue . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_PutSecretValue(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.PutSecretValueInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_secretsmanagerClientRequestToken)
	}
	if len(_secretsmanagerRotationToken) > 0 {
		input.RotationToken = aws.String(_secretsmanagerRotationToken)
	}
	if len(_secretsmanagerSecretBinary) > 0 {
		if err := assignInputField(input, "SecretBinary", _secretsmanagerSecretBinary); err != nil {
			log.Errorf("invalid --secret-binary: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerSecretString) > 0 {
		input.SecretString = aws.String(_secretsmanagerSecretString)
	}
	if len(_secretsmanagerVersionStages) > 0 {
		input.VersionStages = append([]string(nil), _secretsmanagerVersionStages...)
	}

	if resp, err := client.PutSecretValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a secret that is replicated to other Regions, deletes the secret replicas
// from the Regions you specify.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:RemoveRegionsFromReplication . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_RemoveRegionsFromReplication(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.RemoveRegionsFromReplicationInput{
		// RemoveReplicaRegions: []string, // Required
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerRemoveReplicaRegions) > 0 {
		input.RemoveReplicaRegions = append([]string(nil), _secretsmanagerRemoveReplicaRegions...)
	}
	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.RemoveRegionsFromReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replicates the secret to a new Regions. See [Multi-Region secrets].
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ReplicateSecretToRegions . If the primary
// secret is encrypted with a KMS key other than aws/secretsmanager , you also need
// kms:Decrypt permission to the key. To encrypt the replicated secret with a KMS
// key other than aws/secretsmanager , you need kms:GenerateDataKey and kms:Encrypt
// to the key. For more information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Multi-Region secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/create-manage-multi-region-secrets.html
func secretsmanager_ReplicateSecretToRegions(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.ReplicateSecretToRegionsInput{
		// AddReplicaRegions: []types.ReplicaRegionType, // Required
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerAddReplicaRegions) > 0 {
		if err := assignInputField(input, "AddReplicaRegions", _secretsmanagerAddReplicaRegions); err != nil {
			log.Errorf("invalid --add-replica-regions: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerForceOverwriteReplicaSecret) > 0 {
		if err := assignInputField(input, "ForceOverwriteReplicaSecret", _secretsmanagerForceOverwriteReplicaSecret); err != nil {
			log.Errorf("invalid --force-overwrite-replica-secret: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReplicateSecretToRegions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the scheduled deletion of a secret by removing the DeletedDate time
// stamp. You can access a secret again after it has been restored.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:RestoreSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_RestoreSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.RestoreSecretInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.RestoreSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Configures and starts the asynchronous process of rotating the secret. For
// information about rotation, see [Rotate secrets]in the Secrets Manager User Guide. If you
// include the configuration parameters, the operation sets the values for the
// secret and then immediately starts a rotation. If you don't include the
// configuration parameters, the operation starts a rotation with the values
// already stored in the secret.
//
// When rotation is successful, the AWSPENDING staging label might be attached to
// the same version as the AWSCURRENT version, or it might not be attached to any
// version. If the AWSPENDING staging label is present but not attached to the
// same version as AWSCURRENT , then any later invocation of RotateSecret assumes
// that a previous rotation request is still in progress and returns an error. When
// rotation is unsuccessful, the AWSPENDING staging label might be attached to an
// empty secret version. For more information, see [Troubleshoot rotation]in the Secrets Manager User
// Guide.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:RotateSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager]. You also need lambda:InvokeFunction permissions on the rotation function.
// For more information, see [Permissions for rotation].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Permissions for rotation]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets-required-permissions-function.html
// [Rotate secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Troubleshoot rotation]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot_rotation.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_RotateSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.RotateSecretInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_secretsmanagerClientRequestToken)
	}
	if len(_secretsmanagerExternalSecretRotationMetadata) > 0 {
		if err := assignInputField(input, "ExternalSecretRotationMetadata", _secretsmanagerExternalSecretRotationMetadata); err != nil {
			log.Errorf("invalid --external-secret-rotation-metadata: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerExternalSecretRotationRoleArn) > 0 {
		input.ExternalSecretRotationRoleArn = aws.String(_secretsmanagerExternalSecretRotationRoleArn)
	}
	if len(_secretsmanagerRotateImmediately) > 0 {
		if err := assignInputField(input, "RotateImmediately", _secretsmanagerRotateImmediately); err != nil {
			log.Errorf("invalid --rotate-immediately: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerRotationLambdaARN) > 0 {
		input.RotationLambdaARN = aws.String(_secretsmanagerRotationLambdaARN)
	}
	if len(_secretsmanagerRotationRules) > 0 {
		if err := assignInputField(input, "RotationRules", _secretsmanagerRotationRules); err != nil {
			log.Errorf("invalid --rotation-rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.RotateSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the link between the replica secret and the primary secret and promotes
// the replica to a primary secret in the replica Region.
//
// You must call this operation from the Region in which you want to promote the
// replica to a primary secret.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:StopReplicationToReplica . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_StopReplicationToReplica(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.StopReplicationToReplicaInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.StopReplicationToReplica(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches tags to a secret. Tags consist of a key name and a value. Tags are
// part of the secret's metadata. They are not associated with specific versions of
// the secret. This operation appends tags to the existing list of tags.
//
// For tag quotas and naming restrictions, see [Service quotas for Tagging] in the Amazon Web Services General
// Reference guide.
//
// If you use tags as part of your security strategy, then adding or removing a
// tag can change permissions. If successfully completing this operation would
// result in you losing your permissions for this secret, then the operation is
// blocked and returns an Access Denied error.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:TagResource . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Service quotas for Tagging]: https://docs.aws.amazon.com/general/latest/gr/arg.html#taged-reference-quotas
func secretsmanager_TagResource(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.TagResourceInput{
		// SecretId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _secretsmanagerTags); err != nil {
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

// Removes specific tags from a secret.
// This operation is idempotent. If a requested tag is not attached to the secret,
// no error is returned and the secret metadata is unchanged.
//
// If you use tags as part of your security strategy, then removing a tag can
// change permissions. If successfully completing this operation would result in
// you losing your permissions for this secret, then the operation is blocked and
// returns an Access Denied error.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:UntagResource . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_UntagResource(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.UntagResourceInput{
		// SecretId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _secretsmanagerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the details of a secret, including metadata and the secret value. To
// change the secret value, you can also use PutSecretValue.
//
// To change the rotation configuration of a secret, use RotateSecret instead.
//
// To change a secret so that it is managed by another service, you need to
// recreate the secret in that service. See [Secrets Manager secrets managed by other Amazon Web Services services].
//
// We recommend you avoid calling UpdateSecret at a sustained rate of more than
// once every 10 minutes. When you call UpdateSecret to update the secret value,
// Secrets Manager creates a new version of the secret. Secrets Manager removes
// outdated versions when there are more than 100, but it does not remove versions
// created less than 24 hours ago. If you update the secret value more than once
// every 10 minutes, you create more versions than Secrets Manager removes, and you
// will reach the quota for secret versions.
//
// If you include SecretString or SecretBinary to create a new secret version,
// Secrets Manager automatically moves the staging label AWSCURRENT to the new
// version. Then it attaches the label AWSPREVIOUS to the version that AWSCURRENT
// was removed from.
//
// If you call this operation with a ClientRequestToken that matches an existing
// version's VersionId , the operation results in an error. You can't modify an
// existing version, you can only create a new version. To remove a version, remove
// all staging labels from it. See UpdateSecretVersionStage.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters except SecretBinary or
// SecretString because it might be logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:UpdateSecret . For more information, see [IAM policy actions for Secrets Manager]
// and [Authentication and access control in Secrets Manager]. If you use a customer managed key, you must also have kms:GenerateDataKey
// , kms:Encrypt , and kms:Decrypt permissions on the key. If you change the KMS
// key and you don't have kms:Encrypt permission to the new key, Secrets Manager
// does not re-encrypt existing secret versions with the new key. For more
// information, see [Secret encryption and decryption].
//
// When you enter commands in a command shell, there is a risk of the command
// history being accessed or utilities having access to your command parameters.
// This is a concern if the command includes the value of a secret. Learn how to [Mitigate the risks of using command-line tools to store Secrets Manager secrets].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Secret encryption and decryption]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security-encryption.html
// [Secrets Manager secrets managed by other Amazon Web Services services]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/service-linked-secrets.html
// [Mitigate the risks of using command-line tools to store Secrets Manager secrets]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/security_cli-exposure-risks.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_UpdateSecret(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.UpdateSecretInput{
		// SecretId: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_secretsmanagerClientRequestToken)
	}
	if len(_secretsmanagerDescription) > 0 {
		input.Description = aws.String(_secretsmanagerDescription)
	}
	if len(_secretsmanagerKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_secretsmanagerKmsKeyId)
	}
	if len(_secretsmanagerSecretBinary) > 0 {
		if err := assignInputField(input, "SecretBinary", _secretsmanagerSecretBinary); err != nil {
			log.Errorf("invalid --secret-binary: %s", err.Error())
			return
		}
	}
	if len(_secretsmanagerSecretString) > 0 {
		input.SecretString = aws.String(_secretsmanagerSecretString)
	}
	if len(_secretsmanagerType) > 0 {
		input.Type = aws.String(_secretsmanagerType)
	}

	if resp, err := client.UpdateSecret(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the staging labels attached to a version of a secret. Secrets Manager
// uses staging labels to track a version as it progresses through the secret
// rotation process. Each staging label can be attached to only one version at a
// time. To add a staging label to a version when it is already attached to another
// version, Secrets Manager first removes it from the other version first and then
// attaches it to this one. For more information about versions and staging labels,
// see [Concepts: Version].
//
// The staging labels that you specify in the VersionStage parameter are added to
// the existing list of staging labels for the version.
//
// You can move the AWSCURRENT staging label to this version by including it in
// this call.
//
// Whenever you move AWSCURRENT , Secrets Manager automatically moves the label
// AWSPREVIOUS to the version that AWSCURRENT was removed from.
//
// If this action results in the last label being removed from a version, then the
// version is considered to be 'deprecated' and can be deleted by Secrets Manager.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:UpdateSecretVersionStage . For more
// information, see [IAM policy actions for Secrets Manager]and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [Concepts: Version]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/getting-started.html#term_version
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
func secretsmanager_UpdateSecretVersionStage(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.UpdateSecretVersionStageInput{
		// SecretId: *string, // Required
		// VersionStage: *string, // Required
	}

	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}
	if len(_secretsmanagerVersionStage) > 0 {
		input.VersionStage = aws.String(_secretsmanagerVersionStage)
	}
	if len(_secretsmanagerMoveToVersionId) > 0 {
		input.MoveToVersionId = aws.String(_secretsmanagerMoveToVersionId)
	}
	if len(_secretsmanagerRemoveFromVersionId) > 0 {
		input.RemoveFromVersionId = aws.String(_secretsmanagerRemoveFromVersionId)
	}

	if resp, err := client.UpdateSecretVersionStage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Validates that a resource policy does not grant a wide range of principals
// access to your secret. A resource-based policy is optional for secrets.
//
// The API performs three checks when validating the policy:
//
// - Sends a call to [Zelkova], an automated reasoning engine, to ensure your resource
// policy does not allow broad access to your secret, for example policies that use
// a wildcard for the principal.
//
// - Checks for correct syntax in a policy.
//
// - Verifies the policy does not lock out a caller.
//
// Secrets Manager generates a CloudTrail log entry when you call this action. Do
// not include sensitive information in request parameters because it might be
// logged. For more information, see [Logging Secrets Manager events with CloudTrail].
//
// Required permissions: secretsmanager:ValidateResourcePolicy and
// secretsmanager:PutResourcePolicy . For more information, see [IAM policy actions for Secrets Manager] and [Authentication and access control in Secrets Manager].
//
// [Authentication and access control in Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access.html
// [Logging Secrets Manager events with CloudTrail]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/retrieve-ct-entries.html
// [IAM policy actions for Secrets Manager]: https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#reference_iam-permissions_actions
// [Zelkova]: https://aws.amazon.com/blogs/security/protect-sensitive-data-in-the-cloud-with-automated-reasoning-zelkova/
func secretsmanager_ValidateResourcePolicy(cfg aws.Config, client *secretsmanager.Client) {
	input := &secretsmanager.ValidateResourcePolicyInput{
		// ResourcePolicy: *string, // Required
	}

	if len(_secretsmanagerResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_secretsmanagerResourcePolicy)
	}
	if len(_secretsmanagerSecretId) > 0 {
		input.SecretId = aws.String(_secretsmanagerSecretId)
	}

	if resp, err := client.ValidateResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_secretsmanagerCmd)
	_secretsmanagerCmd.Flags().SortFlags = false

	_secretsmanagerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_secretsmanagerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_secretsmanagerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerAddReplicaRegions, "add-replica-regions", "", "", "Add Replica Regions")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerBlockPublicPolicy, "block-public-policy", "", "", "Block Public Policy")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerDescription, "description", "", "", "Description")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExcludeCharacters, "exclude-characters", "", "", "Exclude Characters")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExcludeLowercase, "exclude-lowercase", "", "", "Exclude Lowercase")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExcludeNumbers, "exclude-numbers", "", "", "Exclude Numbers")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExcludePunctuation, "exclude-punctuation", "", "", "Exclude Punctuation")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExcludeUppercase, "exclude-uppercase", "", "", "Exclude Uppercase")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExternalSecretRotationMetadata, "external-secret-rotation-metadata", "", "", "External Secret Rotation Metadata")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerExternalSecretRotationRoleArn, "external-secret-rotation-role-arn", "", "", "External Secret Rotation Role ARN")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerFilters, "filters", "", "", "Filters")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerForceDeleteWithoutRecovery, "force-delete-without-recovery", "", "", "Force Delete Without Recovery")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerForceOverwriteReplicaSecret, "force-overwrite-replica-secret", "", "", "Force Overwrite Replica Secret")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerIncludeDeprecated, "include-deprecated", "", "", "Include Deprecated")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerIncludePlannedDeletion, "include-planned-deletion", "", "", "Include Planned Deletion")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerIncludeSpace, "include-space", "", "", "Include Space")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerMaxResults, "max-results", "", "", "Max Results")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerMoveToVersionId, "move-to-version-id", "", "", "Move To Version ID")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerName, "name", "", "", "Name")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerNextToken, "next-token", "", "", "Next Token")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerPasswordLength, "password-length", "", "", "Password Length")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRecoveryWindowInDays, "recovery-window-in-days", "", "", "Recovery Window In Days")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRemoveFromVersionId, "remove-from-version-id", "", "", "Remove From Version ID")
	_secretsmanagerCmd.Flags().StringSliceVarP(&_secretsmanagerRemoveReplicaRegions, "remove-replica-regions", "", nil, "Remove Replica Regions")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRequireEachIncludedType, "require-each-included-type", "", "", "Require Each Included Type")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRotateImmediately, "rotate-immediately", "", "", "Rotate Immediately")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRotationLambdaARN, "rotation-lambda-arn", "", "", "Rotation Lambda ARN")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRotationRules, "rotation-rules", "", "", "Rotation Rules")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerRotationToken, "rotation-token", "", "", "Rotation Token")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerSecretBinary, "secret-binary", "", "", "Secret Binary")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerSecretId, "secret-id", "", "", "Secret ID")
	_secretsmanagerCmd.Flags().StringSliceVarP(&_secretsmanagerSecretIdList, "secret-id-list", "", nil, "Secret ID List")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerSecretString, "secret-string", "", "", "Secret String")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerSortBy, "sort-by", "", "", "Sort By")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerSortOrder, "sort-order", "", "", "Sort Order")
	_secretsmanagerCmd.Flags().StringSliceVarP(&_secretsmanagerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerTags, "tags", "", "", "Tags")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerType, "type", "", "", "Type")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerVersionId, "version-id", "", "", "Version ID")
	_secretsmanagerCmd.Flags().StringVarP(&_secretsmanagerVersionStage, "version-stage", "", "", "Version Stage")
	_secretsmanagerCmd.Flags().StringSliceVarP(&_secretsmanagerVersionStages, "version-stages", "", nil, "Version Stages")

	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerBatchGetSecretValue, "batch-get-secret-value", "", false, "Batch Get Secret Value")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerCancelRotateSecret, "cancel-rotate-secret", "", false, "Cancel Rotate Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerCreateSecret, "create-secret", "", false, "Create Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerDeleteSecret, "delete-secret", "", false, "Delete Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerDescribeSecret, "describe-secret", "", false, "Describe Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerGetRandomPassword, "get-random-password", "", false, "Get Random Password")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerGetSecretValue, "get-secret-value", "", false, "Get Secret Value")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerListSecretVersionIds, "list-secret-version-ids", "", false, "List Secret Version Ids")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerListSecrets, "list-secrets", "", false, "List Secrets")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerPutSecretValue, "put-secret-value", "", false, "Put Secret Value")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerRemoveRegionsFromReplication, "remove-regions-from-replication", "", false, "Remove Regions From Replication")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerReplicateSecretToRegions, "replicate-secret-to-regions", "", false, "Replicate Secret To Regions")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerRestoreSecret, "restore-secret", "", false, "Restore Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerRotateSecret, "rotate-secret", "", false, "Rotate Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerStopReplicationToReplica, "stop-replication-to-replica", "", false, "Stop Replication To Replica")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerTagResource, "tag-resource", "", false, "Tag Resource")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerUntagResource, "untag-resource", "", false, "Untag Resource")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerUpdateSecret, "update-secret", "", false, "Update Secret")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerUpdateSecretVersionStage, "update-secret-version-stage", "", false, "Update Secret Version Stage")
	_secretsmanagerCmd.Flags().BoolVarP(&_secretsmanagerValidateResourcePolicy, "validate-resource-policy", "", false, "Validate Resource Policy")

}
