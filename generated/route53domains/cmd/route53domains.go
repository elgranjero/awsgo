package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// route53domainsCmd represents the route53domains command
var _route53domainsCmd = &cobra.Command{
	Use:   "route53domains",
	Short: "AWS route53domains CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := route53domains.NewFromConfig(cfg)
		if _route53domainsAcceptDomainTransferFromAnotherAwsAccount {
			route53domains_AcceptDomainTransferFromAnotherAwsAccount(cfg, client)
			return
		}
		if _route53domainsAssociateDelegationSignerToDomain {
			route53domains_AssociateDelegationSignerToDomain(cfg, client)
			return
		}
		if _route53domainsCancelDomainTransferToAnotherAwsAccount {
			route53domains_CancelDomainTransferToAnotherAwsAccount(cfg, client)
			return
		}
		if _route53domainsCheckDomainAvailability {
			route53domains_CheckDomainAvailability(cfg, client)
			return
		}
		if _route53domainsCheckDomainTransferability {
			route53domains_CheckDomainTransferability(cfg, client)
			return
		}
		if _route53domainsDeleteDomain {
			route53domains_DeleteDomain(cfg, client)
			return
		}
		if _route53domainsDeleteTagsForDomain {
			route53domains_DeleteTagsForDomain(cfg, client)
			return
		}
		if _route53domainsDisableDomainAutoRenew {
			route53domains_DisableDomainAutoRenew(cfg, client)
			return
		}
		if _route53domainsDisableDomainTransferLock {
			route53domains_DisableDomainTransferLock(cfg, client)
			return
		}
		if _route53domainsDisassociateDelegationSignerFromDomain {
			route53domains_DisassociateDelegationSignerFromDomain(cfg, client)
			return
		}
		if _route53domainsEnableDomainAutoRenew {
			route53domains_EnableDomainAutoRenew(cfg, client)
			return
		}
		if _route53domainsEnableDomainTransferLock {
			route53domains_EnableDomainTransferLock(cfg, client)
			return
		}
		if _route53domainsGetContactReachabilityStatus {
			route53domains_GetContactReachabilityStatus(cfg, client)
			return
		}
		if _route53domainsGetDomainDetail {
			route53domains_GetDomainDetail(cfg, client)
			return
		}
		if _route53domainsGetDomainSuggestions {
			route53domains_GetDomainSuggestions(cfg, client)
			return
		}
		if _route53domainsGetOperationDetail {
			route53domains_GetOperationDetail(cfg, client)
			return
		}
		if _route53domainsListDomains {
			route53domains_ListDomains(cfg, client)
			return
		}
		if _route53domainsListOperations {
			route53domains_ListOperations(cfg, client)
			return
		}
		if _route53domainsListPrices {
			route53domains_ListPrices(cfg, client)
			return
		}
		if _route53domainsListTagsForDomain {
			route53domains_ListTagsForDomain(cfg, client)
			return
		}
		if _route53domainsPushDomain {
			route53domains_PushDomain(cfg, client)
			return
		}
		if _route53domainsRegisterDomain {
			route53domains_RegisterDomain(cfg, client)
			return
		}
		if _route53domainsRejectDomainTransferFromAnotherAwsAccount {
			route53domains_RejectDomainTransferFromAnotherAwsAccount(cfg, client)
			return
		}
		if _route53domainsRenewDomain {
			route53domains_RenewDomain(cfg, client)
			return
		}
		if _route53domainsResendContactReachabilityEmail {
			route53domains_ResendContactReachabilityEmail(cfg, client)
			return
		}
		if _route53domainsResendOperationAuthorization {
			route53domains_ResendOperationAuthorization(cfg, client)
			return
		}
		if _route53domainsRetrieveDomainAuthCode {
			route53domains_RetrieveDomainAuthCode(cfg, client)
			return
		}
		if _route53domainsTransferDomain {
			route53domains_TransferDomain(cfg, client)
			return
		}
		if _route53domainsTransferDomainToAnotherAwsAccount {
			route53domains_TransferDomainToAnotherAwsAccount(cfg, client)
			return
		}
		if _route53domainsUpdateDomainContact {
			route53domains_UpdateDomainContact(cfg, client)
			return
		}
		if _route53domainsUpdateDomainContactPrivacy {
			route53domains_UpdateDomainContactPrivacy(cfg, client)
			return
		}
		if _route53domainsUpdateDomainNameservers {
			route53domains_UpdateDomainNameservers(cfg, client)
			return
		}
		if _route53domainsUpdateTagsForDomain {
			route53domains_UpdateTagsForDomain(cfg, client)
			return
		}
		if _route53domainsViewBilling {
			route53domains_ViewBilling(cfg, client)
			return
		}

	},
}

var (
	_route53domainsAcceptDomainTransferFromAnotherAwsAccount bool
	_route53domainsAssociateDelegationSignerToDomain         bool
	_route53domainsCancelDomainTransferToAnotherAwsAccount   bool
	_route53domainsCheckDomainAvailability                   bool
	_route53domainsCheckDomainTransferability                bool
	_route53domainsDeleteDomain                              bool
	_route53domainsDeleteTagsForDomain                       bool
	_route53domainsDisableDomainAutoRenew                    bool
	_route53domainsDisableDomainTransferLock                 bool
	_route53domainsDisassociateDelegationSignerFromDomain    bool
	_route53domainsEnableDomainAutoRenew                     bool
	_route53domainsEnableDomainTransferLock                  bool
	_route53domainsGetContactReachabilityStatus              bool
	_route53domainsGetDomainDetail                           bool
	_route53domainsGetDomainSuggestions                      bool
	_route53domainsGetOperationDetail                        bool
	_route53domainsListDomains                               bool
	_route53domainsListOperations                            bool
	_route53domainsListPrices                                bool
	_route53domainsListTagsForDomain                         bool
	_route53domainsPushDomain                                bool
	_route53domainsRegisterDomain                            bool
	_route53domainsRejectDomainTransferFromAnotherAwsAccount bool
	_route53domainsRenewDomain                               bool
	_route53domainsResendContactReachabilityEmail            bool
	_route53domainsResendOperationAuthorization              bool
	_route53domainsRetrieveDomainAuthCode                    bool
	_route53domainsTransferDomain                            bool
	_route53domainsTransferDomainToAnotherAwsAccount         bool
	_route53domainsUpdateDomainContact                       bool
	_route53domainsUpdateDomainContactPrivacy                bool
	_route53domainsUpdateDomainNameservers                   bool
	_route53domainsUpdateTagsForDomain                       bool
	_route53domainsViewBilling                               bool

	_route53domainsAccountId                       string
	_route53domainsAdminContact                    string
	_route53domainsAdminPrivacy                    string
	_route53domainsAuthCode                        string
	_route53domainsAutoRenew                       string
	_route53domainsBillingContact                  string
	_route53domainsBillingPrivacy                  string
	_route53domainsConsent                         string
	_route53domainsCurrentExpiryYear               string
	_route53domainsDomainName                      string
	_route53domainsDurationInYears                 string
	_route53domainsEnd                             string
	_route53domainsFIAuthKey                       string
	_route53domainsFilterConditions                string
	_route53domainsId                              string
	_route53domainsIdnLangCode                     string
	_route53domainsMarker                          string
	_route53domainsMaxItems                        string
	_route53domainsNameservers                     string
	_route53domainsOnlyAvailable                   string
	_route53domainsOperationId                     string
	_route53domainsPassword                        string
	_route53domainsPrivacyProtectAdminContact      string
	_route53domainsPrivacyProtectBillingContact    string
	_route53domainsPrivacyProtectRegistrantContact string
	_route53domainsPrivacyProtectTechContact       string
	_route53domainsRegistrantContact               string
	_route53domainsRegistrantPrivacy               string
	_route53domainsSigningAttributes               string
	_route53domainsSortBy                          string
	_route53domainsSortCondition                   string
	_route53domainsSortOrder                       string
	_route53domainsStart                           string
	_route53domainsStatus                          string
	_route53domainsSubmittedSince                  string
	_route53domainsSuggestionCount                 string
	_route53domainsTagsToDelete                    []string
	_route53domainsTagsToUpdate                    string
	_route53domainsTarget                          string
	_route53domainsTechContact                     string
	_route53domainsTechPrivacy                     string
	_route53domainsTld                             string
	_route53domainsType                            string
)

// Accepts the transfer of a domain from another Amazon Web Services account to
// the currentAmazon Web Services account. You initiate a transfer between Amazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// If you use the CLI command at [accept-domain-transfer-from-another-aws-account], use JSON format as input instead of text
// because otherwise CLI will throw an error from domain transfer input that
// includes single quotes.
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [accept-domain-transfer-from-another-aws-account]: https://docs.aws.amazon.com/cli/latest/reference/route53domains/accept-domain-transfer-from-another-aws-account.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func route53domains_AcceptDomainTransferFromAnotherAwsAccount(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.AcceptDomainTransferFromAnotherAwsAccountInput{
		// DomainName: *string, // Required
		// Password: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsPassword) > 0 {
		input.Password = aws.String(_route53domainsPassword)
	}

	if resp, err := client.AcceptDomainTransferFromAnotherAwsAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a delegation signer (DS) record in the registry zone for this domain
// name.
//
// Note that creating DS record at the registry impacts DNSSEC validation of your
// DNS records. This action may render your domain name unavailable on the internet
// if the steps are completed in the wrong order, or with incorrect timing. For
// more information about DNSSEC signing, see [Configuring DNSSEC signing]in the Route 53 developer guide.
//
// [Configuring DNSSEC signing]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html
func route53domains_AssociateDelegationSignerToDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.AssociateDelegationSignerToDomainInput{
		// DomainName: *string, // Required
		// SigningAttributes: *types.DnssecSigningAttributes, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsSigningAttributes) > 0 {
		if err := assignInputField(input, "SigningAttributes", _route53domainsSigningAttributes); err != nil {
			log.Errorf("invalid --signing-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateDelegationSignerToDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the transfer of a domain from the current Amazon Web Services account
// to another Amazon Web Services account. You initiate a transfer betweenAmazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// You must cancel the transfer before the other Amazon Web Services account
// accepts the transfer using [AcceptDomainTransferFromAnotherAwsAccount].
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func route53domains_CancelDomainTransferToAnotherAwsAccount(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.CancelDomainTransferToAnotherAwsAccountInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.CancelDomainTransferToAnotherAwsAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation checks the availability of one domain name. Note that if the
// availability status of a domain is pending, you must submit another request to
// determine the availability of the domain name.
func route53domains_CheckDomainAvailability(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.CheckDomainAvailabilityInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsIdnLangCode) > 0 {
		input.IdnLangCode = aws.String(_route53domainsIdnLangCode)
	}

	if resp, err := client.CheckDomainAvailability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks whether a domain name can be transferred to Amazon Route 53.
func route53domains_CheckDomainTransferability(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.CheckDomainTransferabilityInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsAuthCode) > 0 {
		input.AuthCode = aws.String(_route53domainsAuthCode)
	}

	if resp, err := client.CheckDomainTransferability(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes the specified domain. This action is permanent. For more
// information, see [Deleting a domain name registration].
//
// To transfer the domain registration to another registrar, use the transfer
// process that’s provided by the registrar to which you want to transfer the
// registration. Otherwise, the following apply:
//
// - You can’t get a refund for the cost of a deleted domain registration.
//
// - The registry for the top-level domain might hold the domain name for a
// brief time before releasing it for other users to register (varies by registry).
//
// - When the registration has been deleted, we'll send you a confirmation to
// the registrant contact. The email will come from
// noreply(at)domainnameverification.net or noreply(at)registrar.amazon.com .
//
// [Deleting a domain name registration]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-delete.html
func route53domains_DeleteDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.DeleteDomainInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes the specified tags for a domain.
// All tag operations are eventually consistent; subsequent operations might not
// immediately represent all issued operations.
func route53domains_DeleteTagsForDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.DeleteTagsForDomainInput{
		// DomainName: *string, // Required
		// TagsToDelete: []string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsTagsToDelete) > 0 {
		input.TagsToDelete = append([]string(nil), _route53domainsTagsToDelete...)
	}

	if resp, err := client.DeleteTagsForDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation disables automatic renewal of domain registration for the
// specified domain.
func route53domains_DisableDomainAutoRenew(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.DisableDomainAutoRenewInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.DisableDomainAutoRenew(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation removes the transfer lock on the domain (specifically the
// clientTransferProhibited status) to allow domain transfers. We recommend you
// refrain from performing this action unless you intend to transfer the domain to
// a different registrar. Successful submission returns an operation ID that you
// can use to track the progress and completion of the action. If the request is
// not completed successfully, the domain registrant will be notified by email.
func route53domains_DisableDomainTransferLock(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.DisableDomainTransferLockInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.DisableDomainTransferLock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a delegation signer (DS) record in the registry zone for this domain
// name.
func route53domains_DisassociateDelegationSignerFromDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.DisassociateDelegationSignerFromDomainInput{
		// DomainName: *string, // Required
		// Id: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsId) > 0 {
		input.Id = aws.String(_route53domainsId)
	}

	if resp, err := client.DisassociateDelegationSignerFromDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation configures Amazon Route 53 to automatically renew the specified
// domain before the domain registration expires. The cost of renewing your domain
// registration is billed to your Amazon Web Services account.
//
// The period during which you can renew a domain name varies by TLD. For a list
// of TLDs and their renewal policies, see [Domains That You Can Register with Amazon Route 53]in the Amazon Route 53 Developer Guide.
// Route 53 requires that you renew before the end of the renewal period so we can
// complete processing before the deadline.
//
// [Domains That You Can Register with Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html
func route53domains_EnableDomainAutoRenew(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.EnableDomainAutoRenewInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.EnableDomainAutoRenew(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation sets the transfer lock on the domain (specifically the
// clientTransferProhibited status) to prevent domain transfers. Successful
// submission returns an operation ID that you can use to track the progress and
// completion of the action. If the request is not completed successfully, the
// domain registrant will be notified by email.
func route53domains_EnableDomainTransferLock(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.EnableDomainTransferLockInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.EnableDomainTransferLock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For operations that require confirmation that the email address for the
// registrant contact is valid, such as registering a new domain, this operation
// returns information about whether the registrant contact has responded.
//
// If you want us to resend the email, use the ResendContactReachabilityEmail
// operation.
func route53domains_GetContactReachabilityStatus(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.GetContactReachabilityStatusInput{}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.GetContactReachabilityStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns detailed information about a specified domain that is
// associated with the current Amazon Web Services account. Contact information for
// the domain is also returned as part of the output.
func route53domains_GetDomainDetail(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.GetDomainDetailInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.GetDomainDetail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetDomainSuggestions operation returns a list of suggested domain names.
func route53domains_GetDomainSuggestions(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.GetDomainSuggestionsInput{
		// DomainName: *string, // Required
		// OnlyAvailable: *bool, // Required
		// SuggestionCount: int32, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsOnlyAvailable) > 0 {
		if err := assignInputField(input, "OnlyAvailable", _route53domainsOnlyAvailable); err != nil {
			log.Errorf("invalid --only-available: %s", err.Error())
			return
		}
	}
	if len(_route53domainsSuggestionCount) > 0 {
		if err := assignInputField(input, "SuggestionCount", _route53domainsSuggestionCount); err != nil {
			log.Errorf("invalid --suggestion-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDomainSuggestions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns the current status of an operation that is not completed.
func route53domains_GetOperationDetail(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.GetOperationDetailInput{
		// OperationId: *string, // Required
	}

	if len(_route53domainsOperationId) > 0 {
		input.OperationId = aws.String(_route53domainsOperationId)
	}

	if resp, err := client.GetOperationDetail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns all the domain names registered with Amazon Route 53 for
// the current Amazon Web Services account if no filtering conditions are used.
func route53domains_ListDomains(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ListDomainsInput{}

	if len(_route53domainsFilterConditions) > 0 {
		if err := assignInputField(input, "FilterConditions", _route53domainsFilterConditions); err != nil {
			log.Errorf("invalid --filter-conditions: %s", err.Error())
			return
		}
	}
	if len(_route53domainsMarker) > 0 {
		input.Marker = aws.String(_route53domainsMarker)
	}
	if len(_route53domainsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53domainsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53domainsSortCondition) > 0 {
		if err := assignInputField(input, "SortCondition", _route53domainsSortCondition); err != nil {
			log.Errorf("invalid --sort-condition: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDomains(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53domains.ListDomainsOutput
	p := route53domains.NewListDomainsPaginator(client, input)
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

// Returns information about all of the operations that return an operation ID and
// that have ever been performed on domains that were registered by the current
// account.
//
// This command runs only in the us-east-1 Region.
func route53domains_ListOperations(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ListOperationsInput{}

	if len(_route53domainsMarker) > 0 {
		input.Marker = aws.String(_route53domainsMarker)
	}
	if len(_route53domainsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53domainsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53domainsSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _route53domainsSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_route53domainsSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _route53domainsSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}
	if len(_route53domainsStatus) > 0 {
		if err := assignInputField(input, "Status", _route53domainsStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_route53domainsSubmittedSince) > 0 {
		if err := assignInputField(input, "SubmittedSince", _route53domainsSubmittedSince); err != nil {
			log.Errorf("invalid --submitted-since: %s", err.Error())
			return
		}
	}
	if len(_route53domainsType) > 0 {
		if err := assignInputField(input, "Type", _route53domainsType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListOperations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53domains.ListOperationsOutput
	p := route53domains.NewListOperationsPaginator(client, input)
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

// Lists the following prices for either all the TLDs supported by Route 53, or
// the specified TLD:
//
// - Registration
//
// - Transfer
//
// - Owner change
//
// - Domain renewal
//
// - Domain restoration
func route53domains_ListPrices(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ListPricesInput{}

	if len(_route53domainsMarker) > 0 {
		input.Marker = aws.String(_route53domainsMarker)
	}
	if len(_route53domainsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53domainsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53domainsTld) > 0 {
		input.Tld = aws.String(_route53domainsTld)
	}

	if disablePaginator() {
		if resp, err := client.ListPrices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53domains.ListPricesOutput
	p := route53domains.NewListPricesPaginator(client, input)
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

// This operation returns all of the tags that are associated with the specified
// domain.
//
// All tag operations are eventually consistent; subsequent operations might not
// immediately represent all issued operations.
func route53domains_ListTagsForDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ListTagsForDomainInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.ListTagsForDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Moves a domain from Amazon Web Services to another registrar.
// Supported actions:
//
// - Changes the IPS tags of a .uk domain, and pushes it to transit. Transit
// means that the domain is ready to be transferred to another registrar.
func route53domains_PushDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.PushDomainInput{
		// DomainName: *string, // Required
		// Target: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsTarget) > 0 {
		input.Target = aws.String(_route53domainsTarget)
	}

	if resp, err := client.PushDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation registers a domain. For some top-level domains (TLDs), this
// operation requires extra parameters.
//
// When you register a domain, Amazon Route 53 does the following:
//
// - Creates a Route 53 hosted zone that has the same name as the domain. Route
// 53 assigns four name servers to your hosted zone and automatically updates your
// domain registration with the names of these name servers.
//
// - Enables auto renew, so your domain registration will renew automatically
// each year. We'll notify you in advance of the renewal date so you can choose
// whether to renew the registration.
//
// - Optionally enables privacy protection, so WHOIS queries return contact for
// the registrar or the phrase "REDACTED FOR PRIVACY", or "On behalf of owner." If
// you don't enable privacy protection, WHOIS queries return the information that
// you entered for the administrative, registrant, and technical contacts.
//
// # While some domains may allow different privacy settings per contact, we
//
// recommend specifying the same privacy setting for all contacts.
//
// - If registration is successful, returns an operation ID that you can use to
// track the progress and completion of the action. If the request is not completed
// successfully, the domain registrant is notified by email.
//
// - Charges your Amazon Web Services account an amount based on the top-level
// domain. For more information, see [Amazon Route 53 Pricing].
//
// [Amazon Route 53 Pricing]: http://aws.amazon.com/route53/pricing/
func route53domains_RegisterDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.RegisterDomainInput{
		// AdminContact: *types.ContactDetail, // Required
		// DomainName: *string, // Required
		// DurationInYears: *int32, // Required
		// RegistrantContact: *types.ContactDetail, // Required
		// TechContact: *types.ContactDetail, // Required
	}

	if len(_route53domainsAdminContact) > 0 {
		if err := assignInputField(input, "AdminContact", _route53domainsAdminContact); err != nil {
			log.Errorf("invalid --admin-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsDurationInYears) > 0 {
		if err := assignInputField(input, "DurationInYears", _route53domainsDurationInYears); err != nil {
			log.Errorf("invalid --duration-in-years: %s", err.Error())
			return
		}
	}
	if len(_route53domainsRegistrantContact) > 0 {
		if err := assignInputField(input, "RegistrantContact", _route53domainsRegistrantContact); err != nil {
			log.Errorf("invalid --registrant-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsTechContact) > 0 {
		if err := assignInputField(input, "TechContact", _route53domainsTechContact); err != nil {
			log.Errorf("invalid --tech-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsAutoRenew) > 0 {
		if err := assignInputField(input, "AutoRenew", _route53domainsAutoRenew); err != nil {
			log.Errorf("invalid --auto-renew: %s", err.Error())
			return
		}
	}
	if len(_route53domainsBillingContact) > 0 {
		if err := assignInputField(input, "BillingContact", _route53domainsBillingContact); err != nil {
			log.Errorf("invalid --billing-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsIdnLangCode) > 0 {
		input.IdnLangCode = aws.String(_route53domainsIdnLangCode)
	}
	if len(_route53domainsPrivacyProtectAdminContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectAdminContact", _route53domainsPrivacyProtectAdminContact); err != nil {
			log.Errorf("invalid --privacy-protect-admin-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectBillingContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectBillingContact", _route53domainsPrivacyProtectBillingContact); err != nil {
			log.Errorf("invalid --privacy-protect-billing-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectRegistrantContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectRegistrantContact", _route53domainsPrivacyProtectRegistrantContact); err != nil {
			log.Errorf("invalid --privacy-protect-registrant-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectTechContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectTechContact", _route53domainsPrivacyProtectTechContact); err != nil {
			log.Errorf("invalid --privacy-protect-tech-contact: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects the transfer of a domain from another Amazon Web Services account to
// the current Amazon Web Services account. You initiate a transfer betweenAmazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func route53domains_RejectDomainTransferFromAnotherAwsAccount(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.RejectDomainTransferFromAnotherAwsAccountInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.RejectDomainTransferFromAnotherAwsAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation renews a domain for the specified number of years. The cost of
// renewing your domain is billed to your Amazon Web Services account.
//
// We recommend that you renew your domain several weeks before the expiration
// date. Some TLD registries delete domains before the expiration date if you
// haven't renewed far enough in advance. For more information about renewing
// domain registration, see [Renewing Registration for a Domain]in the Amazon Route 53 Developer Guide.
//
// [Renewing Registration for a Domain]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-renew.html
func route53domains_RenewDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.RenewDomainInput{
		// CurrentExpiryYear: int32, // Required
		// DomainName: *string, // Required
	}

	if len(_route53domainsCurrentExpiryYear) > 0 {
		if err := assignInputField(input, "CurrentExpiryYear", _route53domainsCurrentExpiryYear); err != nil {
			log.Errorf("invalid --current-expiry-year: %s", err.Error())
			return
		}
	}
	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsDurationInYears) > 0 {
		if err := assignInputField(input, "DurationInYears", _route53domainsDurationInYears); err != nil {
			log.Errorf("invalid --duration-in-years: %s", err.Error())
			return
		}
	}

	if resp, err := client.RenewDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For operations that require confirmation that the email address for the
// registrant contact is valid, such as registering a new domain, this operation
// resends the confirmation email to the current email address for the registrant
// contact.
func route53domains_ResendContactReachabilityEmail(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ResendContactReachabilityEmailInput{}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.ResendContactReachabilityEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resend the form of authorization email for this operation.
func route53domains_ResendOperationAuthorization(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ResendOperationAuthorizationInput{
		// OperationId: *string, // Required
	}

	if len(_route53domainsOperationId) > 0 {
		input.OperationId = aws.String(_route53domainsOperationId)
	}

	if resp, err := client.ResendOperationAuthorization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns the authorization code for the domain. To transfer a
// domain to another registrar, you provide this value to the new registrar.
func route53domains_RetrieveDomainAuthCode(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.RetrieveDomainAuthCodeInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.RetrieveDomainAuthCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transfers a domain from another registrar to Amazon Route 53.
// For more information about transferring domains, see the following topics:
//
// - For transfer requirements, a detailed procedure, and information about
// viewing the status of a domain that you're transferring to Route 53, see [Transferring Registration for a Domain to Amazon Route 53]in
// the Amazon Route 53 Developer Guide.
//
// - For information about how to transfer a domain from one Amazon Web Services
// account to another, see [TransferDomainToAnotherAwsAccount].
//
// - For information about how to transfer a domain to another domain registrar,
// see [Transferring a Domain from Amazon Route 53 to Another Registrar]in the Amazon Route 53 Developer Guide.
//
// During the transfer of any country code top-level domains (ccTLDs) to Route 53,
// except for .cc and .tv, updates to the owner contact are ignored and the owner
// contact data from the registry is used. You can update the owner contact after
// the transfer is complete. For more information, see [UpdateDomainContact].
//
// If the registrar for your domain is also the DNS service provider for the
// domain, we highly recommend that you transfer your DNS service to Route 53 or to
// another DNS service provider before you transfer your registration. Some
// registrars provide free DNS service when you purchase a domain registration.
// When you transfer the registration, the previous registrar will not renew your
// domain registration and could end your DNS service at any time.
//
// If the registrar for your domain is also the DNS service provider for the
// domain and you don't transfer DNS service to another provider, your website,
// email, and the web applications associated with the domain might become
// unavailable.
//
// If the transfer is successful, this method returns an operation ID that you can
// use to track the progress and completion of the action. If the transfer doesn't
// complete successfully, the domain registrant will be notified by email.
//
// [Transferring a Domain from Amazon Route 53 to Another Registrar]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-from-route-53.html
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [Transferring Registration for a Domain to Amazon Route 53]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html
// [UpdateDomainContact]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainContact.html
func route53domains_TransferDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.TransferDomainInput{
		// AdminContact: *types.ContactDetail, // Required
		// DomainName: *string, // Required
		// DurationInYears: *int32, // Required
		// RegistrantContact: *types.ContactDetail, // Required
		// TechContact: *types.ContactDetail, // Required
	}

	if len(_route53domainsAdminContact) > 0 {
		if err := assignInputField(input, "AdminContact", _route53domainsAdminContact); err != nil {
			log.Errorf("invalid --admin-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsDurationInYears) > 0 {
		if err := assignInputField(input, "DurationInYears", _route53domainsDurationInYears); err != nil {
			log.Errorf("invalid --duration-in-years: %s", err.Error())
			return
		}
	}
	if len(_route53domainsRegistrantContact) > 0 {
		if err := assignInputField(input, "RegistrantContact", _route53domainsRegistrantContact); err != nil {
			log.Errorf("invalid --registrant-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsTechContact) > 0 {
		if err := assignInputField(input, "TechContact", _route53domainsTechContact); err != nil {
			log.Errorf("invalid --tech-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsAuthCode) > 0 {
		input.AuthCode = aws.String(_route53domainsAuthCode)
	}
	if len(_route53domainsAutoRenew) > 0 {
		if err := assignInputField(input, "AutoRenew", _route53domainsAutoRenew); err != nil {
			log.Errorf("invalid --auto-renew: %s", err.Error())
			return
		}
	}
	if len(_route53domainsBillingContact) > 0 {
		if err := assignInputField(input, "BillingContact", _route53domainsBillingContact); err != nil {
			log.Errorf("invalid --billing-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsIdnLangCode) > 0 {
		input.IdnLangCode = aws.String(_route53domainsIdnLangCode)
	}
	if len(_route53domainsNameservers) > 0 {
		if err := assignInputField(input, "Nameservers", _route53domainsNameservers); err != nil {
			log.Errorf("invalid --nameservers: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectAdminContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectAdminContact", _route53domainsPrivacyProtectAdminContact); err != nil {
			log.Errorf("invalid --privacy-protect-admin-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectBillingContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectBillingContact", _route53domainsPrivacyProtectBillingContact); err != nil {
			log.Errorf("invalid --privacy-protect-billing-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectRegistrantContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectRegistrantContact", _route53domainsPrivacyProtectRegistrantContact); err != nil {
			log.Errorf("invalid --privacy-protect-registrant-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsPrivacyProtectTechContact) > 0 {
		if err := assignInputField(input, "PrivacyProtectTechContact", _route53domainsPrivacyProtectTechContact); err != nil {
			log.Errorf("invalid --privacy-protect-tech-contact: %s", err.Error())
			return
		}
	}

	if resp, err := client.TransferDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Transfers a domain from the current Amazon Web Services account to another
// Amazon Web Services account. Note the following:
//
// - The Amazon Web Services account that you're transferring the domain to must
// accept the transfer. If the other account doesn't accept the transfer within 3
// days, we cancel the transfer. See [AcceptDomainTransferFromAnotherAwsAccount].
//
// - You can cancel the transfer before the other account accepts it. See [CancelDomainTransferToAnotherAwsAccount].
//
// - The other account can reject the transfer. See [RejectDomainTransferFromAnotherAwsAccount].
//
// When you transfer a domain from one Amazon Web Services account to another,
// Route 53 doesn't transfer the hosted zone that is associated with the domain.
// DNS resolution isn't affected if the domain and the hosted zone are owned by
// separate accounts, so transferring the hosted zone is optional. For information
// about transferring the hosted zone to another Amazon Web Services account, see [Migrating a Hosted Zone to a Different Amazon Web Services Account]
// in the Amazon Route 53 Developer Guide.
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [RejectDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_RejectDomainTransferFromAnotherAwsAccount.html
// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
// [Migrating a Hosted Zone to a Different Amazon Web Services Account]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/hosted-zones-migrating.html
// [CancelDomainTransferToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_CancelDomainTransferToAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func route53domains_TransferDomainToAnotherAwsAccount(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.TransferDomainToAnotherAwsAccountInput{
		// AccountId: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_route53domainsAccountId) > 0 {
		input.AccountId = aws.String(_route53domainsAccountId)
	}
	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}

	if resp, err := client.TransferDomainToAnotherAwsAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates the contact information for a particular domain. You
// must specify information for at least one contact: registrant, administrator, or
// technical.
//
// If the update is successful, this method returns an operation ID that you can
// use to track the progress and completion of the operation. If the request is not
// completed successfully, the domain registrant will be notified by email.
func route53domains_UpdateDomainContact(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.UpdateDomainContactInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsAdminContact) > 0 {
		if err := assignInputField(input, "AdminContact", _route53domainsAdminContact); err != nil {
			log.Errorf("invalid --admin-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsBillingContact) > 0 {
		if err := assignInputField(input, "BillingContact", _route53domainsBillingContact); err != nil {
			log.Errorf("invalid --billing-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsConsent) > 0 {
		if err := assignInputField(input, "Consent", _route53domainsConsent); err != nil {
			log.Errorf("invalid --consent: %s", err.Error())
			return
		}
	}
	if len(_route53domainsRegistrantContact) > 0 {
		if err := assignInputField(input, "RegistrantContact", _route53domainsRegistrantContact); err != nil {
			log.Errorf("invalid --registrant-contact: %s", err.Error())
			return
		}
	}
	if len(_route53domainsTechContact) > 0 {
		if err := assignInputField(input, "TechContact", _route53domainsTechContact); err != nil {
			log.Errorf("invalid --tech-contact: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation updates the specified domain contact's privacy setting. When
// privacy protection is enabled, your contact information is replaced with contact
// information for the registrar or with the phrase "REDACTED FOR PRIVACY", or "On
// behalf of owner."
//
// While some domains may allow different privacy settings per contact, we
// recommend specifying the same privacy setting for all contacts.
//
// This operation affects only the contact information for the specified contact
// type (administrative, registrant, or technical). If the request succeeds, Amazon
// Route 53 returns an operation ID that you can use with [GetOperationDetail]to track the progress
// and completion of the action. If the request doesn't complete successfully, the
// domain registrant will be notified by email.
//
// By disabling the privacy service via API, you consent to the publication of the
// contact information provided for this domain via the public WHOIS database. You
// certify that you are the registrant of this domain name and have the authority
// to make this decision. You may withdraw your consent at any time by enabling
// privacy protection using either UpdateDomainContactPrivacy or the Route 53
// console. Enabling privacy protection removes the contact information provided
// for this domain from the WHOIS database. For more information on our privacy
// practices, see [https://aws.amazon.com/privacy/].
//
// [https://aws.amazon.com/privacy/]: https://aws.amazon.com/privacy/
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
func route53domains_UpdateDomainContactPrivacy(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.UpdateDomainContactPrivacyInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsAdminPrivacy) > 0 {
		if err := assignInputField(input, "AdminPrivacy", _route53domainsAdminPrivacy); err != nil {
			log.Errorf("invalid --admin-privacy: %s", err.Error())
			return
		}
	}
	if len(_route53domainsBillingPrivacy) > 0 {
		if err := assignInputField(input, "BillingPrivacy", _route53domainsBillingPrivacy); err != nil {
			log.Errorf("invalid --billing-privacy: %s", err.Error())
			return
		}
	}
	if len(_route53domainsRegistrantPrivacy) > 0 {
		if err := assignInputField(input, "RegistrantPrivacy", _route53domainsRegistrantPrivacy); err != nil {
			log.Errorf("invalid --registrant-privacy: %s", err.Error())
			return
		}
	}
	if len(_route53domainsTechPrivacy) > 0 {
		if err := assignInputField(input, "TechPrivacy", _route53domainsTechPrivacy); err != nil {
			log.Errorf("invalid --tech-privacy: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainContactPrivacy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation replaces the current set of name servers for the domain with the
// specified set of name servers. If you use Amazon Route 53 as your DNS service,
// specify the four name servers in the delegation set for the hosted zone for the
// domain.
//
// If successful, this operation returns an operation ID that you can use to track
// the progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
func route53domains_UpdateDomainNameservers(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.UpdateDomainNameserversInput{
		// DomainName: *string, // Required
		// Nameservers: []types.Nameserver, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsNameservers) > 0 {
		if err := assignInputField(input, "Nameservers", _route53domainsNameservers); err != nil {
			log.Errorf("invalid --nameservers: %s", err.Error())
			return
		}
	}
	if len(_route53domainsFIAuthKey) > 0 {
		input.FIAuthKey = aws.String(_route53domainsFIAuthKey)
	}

	if resp, err := client.UpdateDomainNameservers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation adds or updates tags for a specified domain.
// All tag operations are eventually consistent; subsequent operations might not
// immediately represent all issued operations.
func route53domains_UpdateTagsForDomain(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.UpdateTagsForDomainInput{
		// DomainName: *string, // Required
	}

	if len(_route53domainsDomainName) > 0 {
		input.DomainName = aws.String(_route53domainsDomainName)
	}
	if len(_route53domainsTagsToUpdate) > 0 {
		if err := assignInputField(input, "TagsToUpdate", _route53domainsTagsToUpdate); err != nil {
			log.Errorf("invalid --tags-to-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTagsForDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all the domain-related billing records for the current Amazon Web
// Services account for a specified period
func route53domains_ViewBilling(cfg aws.Config, client *route53domains.Client) {
	input := &route53domains.ViewBillingInput{}

	if len(_route53domainsEnd) > 0 {
		if err := assignInputField(input, "End", _route53domainsEnd); err != nil {
			log.Errorf("invalid --end: %s", err.Error())
			return
		}
	}
	if len(_route53domainsMarker) > 0 {
		input.Marker = aws.String(_route53domainsMarker)
	}
	if len(_route53domainsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _route53domainsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_route53domainsStart) > 0 {
		if err := assignInputField(input, "Start", _route53domainsStart); err != nil {
			log.Errorf("invalid --start: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ViewBilling(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*route53domains.ViewBillingOutput
	p := route53domains.NewViewBillingPaginator(client, input)
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
	_rootCmd.AddCommand(_route53domainsCmd)
	_route53domainsCmd.Flags().SortFlags = false

	_route53domainsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_route53domainsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_route53domainsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_route53domainsCmd.Flags().StringVarP(&_route53domainsAccountId, "account-id", "", "", "Account ID")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsAdminContact, "admin-contact", "", "", "Admin Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsAdminPrivacy, "admin-privacy", "", "", "Admin Privacy")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsAuthCode, "auth-code", "", "", "Auth Code")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsAutoRenew, "auto-renew", "", "", "Auto Renew")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsBillingContact, "billing-contact", "", "", "Billing Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsBillingPrivacy, "billing-privacy", "", "", "Billing Privacy")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsConsent, "consent", "", "", "Consent")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsCurrentExpiryYear, "current-expiry-year", "", "", "Current Expiry Year")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsDomainName, "domain-name", "", "", "Domain Name")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsDurationInYears, "duration-in-years", "", "", "Duration In Years")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsEnd, "end", "", "", "End")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsFIAuthKey, "fi-auth-key", "", "", "Fi Auth Key")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsFilterConditions, "filter-conditions", "", "", "Filter Conditions")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsId, "id", "", "", "ID")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsIdnLangCode, "idn-lang-code", "", "", "Idn Lang Code")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsMarker, "marker", "", "", "Marker")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsMaxItems, "max-items", "", "", "Max Items")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsNameservers, "nameservers", "", "", "Nameservers")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsOnlyAvailable, "only-available", "", "", "Only Available")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsOperationId, "operation-id", "", "", "Operation ID")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsPassword, "password", "", "", "Password")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsPrivacyProtectAdminContact, "privacy-protect-admin-contact", "", "", "Privacy Protect Admin Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsPrivacyProtectBillingContact, "privacy-protect-billing-contact", "", "", "Privacy Protect Billing Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsPrivacyProtectRegistrantContact, "privacy-protect-registrant-contact", "", "", "Privacy Protect Registrant Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsPrivacyProtectTechContact, "privacy-protect-tech-contact", "", "", "Privacy Protect Tech Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsRegistrantContact, "registrant-contact", "", "", "Registrant Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsRegistrantPrivacy, "registrant-privacy", "", "", "Registrant Privacy")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsSigningAttributes, "signing-attributes", "", "", "Signing Attributes")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsSortBy, "sort-by", "", "", "Sort By")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsSortCondition, "sort-condition", "", "", "Sort Condition")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsSortOrder, "sort-order", "", "", "Sort Order")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsStart, "start", "", "", "Start")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsStatus, "status", "", "", "Status")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsSubmittedSince, "submitted-since", "", "", "Submitted Since")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsSuggestionCount, "suggestion-count", "", "", "Suggestion Count")
	_route53domainsCmd.Flags().StringSliceVarP(&_route53domainsTagsToDelete, "tags-to-delete", "", nil, "Tags To Delete")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsTagsToUpdate, "tags-to-update", "", "", "Tags To Update")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsTarget, "target", "", "", "Target")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsTechContact, "tech-contact", "", "", "Tech Contact")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsTechPrivacy, "tech-privacy", "", "", "Tech Privacy")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsTld, "tld", "", "", "Tld")
	_route53domainsCmd.Flags().StringVarP(&_route53domainsType, "type", "", "", "Type")

	_route53domainsCmd.Flags().BoolVarP(&_route53domainsAcceptDomainTransferFromAnotherAwsAccount, "accept-domain-transfer-from-another-aws-account", "", false, "Accept Domain Transfer From Another AWS Account")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsAssociateDelegationSignerToDomain, "associate-delegation-signer-to-domain", "", false, "Associate Delegation Signer To Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsCancelDomainTransferToAnotherAwsAccount, "cancel-domain-transfer-to-another-aws-account", "", false, "Cancel Domain Transfer To Another AWS Account")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsCheckDomainAvailability, "check-domain-availability", "", false, "Check Domain Availability")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsCheckDomainTransferability, "check-domain-transferability", "", false, "Check Domain Transferability")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsDeleteTagsForDomain, "delete-tags-for-domain", "", false, "Delete Tags For Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsDisableDomainAutoRenew, "disable-domain-auto-renew", "", false, "Disable Domain Auto Renew")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsDisableDomainTransferLock, "disable-domain-transfer-lock", "", false, "Disable Domain Transfer Lock")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsDisassociateDelegationSignerFromDomain, "disassociate-delegation-signer-from-domain", "", false, "Disassociate Delegation Signer From Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsEnableDomainAutoRenew, "enable-domain-auto-renew", "", false, "Enable Domain Auto Renew")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsEnableDomainTransferLock, "enable-domain-transfer-lock", "", false, "Enable Domain Transfer Lock")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsGetContactReachabilityStatus, "get-contact-reachability-status", "", false, "Get Contact Reachability Status")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsGetDomainDetail, "get-domain-detail", "", false, "Get Domain Detail")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsGetDomainSuggestions, "get-domain-suggestions", "", false, "Get Domain Suggestions")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsGetOperationDetail, "get-operation-detail", "", false, "Get Operation Detail")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsListDomains, "list-domains", "", false, "List Domains")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsListOperations, "list-operations", "", false, "List Operations")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsListPrices, "list-prices", "", false, "List Prices")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsListTagsForDomain, "list-tags-for-domain", "", false, "List Tags For Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsPushDomain, "push-domain", "", false, "Push Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsRegisterDomain, "register-domain", "", false, "Register Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsRejectDomainTransferFromAnotherAwsAccount, "reject-domain-transfer-from-another-aws-account", "", false, "Reject Domain Transfer From Another AWS Account")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsRenewDomain, "renew-domain", "", false, "Renew Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsResendContactReachabilityEmail, "resend-contact-reachability-email", "", false, "Resend Contact Reachability Email")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsResendOperationAuthorization, "resend-operation-authorization", "", false, "Resend Operation Authorization")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsRetrieveDomainAuthCode, "retrieve-domain-auth-code", "", false, "Retrieve Domain Auth Code")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsTransferDomain, "transfer-domain", "", false, "Transfer Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsTransferDomainToAnotherAwsAccount, "transfer-domain-to-another-aws-account", "", false, "Transfer Domain To Another AWS Account")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsUpdateDomainContact, "update-domain-contact", "", false, "Update Domain Contact")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsUpdateDomainContactPrivacy, "update-domain-contact-privacy", "", false, "Update Domain Contact Privacy")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsUpdateDomainNameservers, "update-domain-nameservers", "", false, "Update Domain Nameservers")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsUpdateTagsForDomain, "update-tags-for-domain", "", false, "Update Tags For Domain")
	_route53domainsCmd.Flags().BoolVarP(&_route53domainsViewBilling, "view-billing", "", false, "View Billing")

}
