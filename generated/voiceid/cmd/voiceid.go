package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/voiceid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// voiceidCmd represents the voiceid command
var _voiceidCmd = &cobra.Command{
	Use:   "voiceid",
	Short: "AWS voiceid CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := voiceid.NewFromConfig(cfg)
		if _voiceidAssociateFraudster {
			voiceid_AssociateFraudster(cfg, client)
			return
		}
		if _voiceidCreateDomain {
			voiceid_CreateDomain(cfg, client)
			return
		}
		if _voiceidCreateWatchlist {
			voiceid_CreateWatchlist(cfg, client)
			return
		}
		if _voiceidDeleteDomain {
			voiceid_DeleteDomain(cfg, client)
			return
		}
		if _voiceidDeleteFraudster {
			voiceid_DeleteFraudster(cfg, client)
			return
		}
		if _voiceidDeleteSpeaker {
			voiceid_DeleteSpeaker(cfg, client)
			return
		}
		if _voiceidDeleteWatchlist {
			voiceid_DeleteWatchlist(cfg, client)
			return
		}
		if _voiceidDescribeDomain {
			voiceid_DescribeDomain(cfg, client)
			return
		}
		if _voiceidDescribeFraudster {
			voiceid_DescribeFraudster(cfg, client)
			return
		}
		if _voiceidDescribeFraudsterRegistrationJob {
			voiceid_DescribeFraudsterRegistrationJob(cfg, client)
			return
		}
		if _voiceidDescribeSpeaker {
			voiceid_DescribeSpeaker(cfg, client)
			return
		}
		if _voiceidDescribeSpeakerEnrollmentJob {
			voiceid_DescribeSpeakerEnrollmentJob(cfg, client)
			return
		}
		if _voiceidDescribeWatchlist {
			voiceid_DescribeWatchlist(cfg, client)
			return
		}
		if _voiceidDisassociateFraudster {
			voiceid_DisassociateFraudster(cfg, client)
			return
		}
		if _voiceidEvaluateSession {
			voiceid_EvaluateSession(cfg, client)
			return
		}
		if _voiceidListDomains {
			voiceid_ListDomains(cfg, client)
			return
		}
		if _voiceidListFraudsterRegistrationJobs {
			voiceid_ListFraudsterRegistrationJobs(cfg, client)
			return
		}
		if _voiceidListFraudsters {
			voiceid_ListFraudsters(cfg, client)
			return
		}
		if _voiceidListSpeakerEnrollmentJobs {
			voiceid_ListSpeakerEnrollmentJobs(cfg, client)
			return
		}
		if _voiceidListSpeakers {
			voiceid_ListSpeakers(cfg, client)
			return
		}
		if _voiceidListTagsForResource {
			voiceid_ListTagsForResource(cfg, client)
			return
		}
		if _voiceidListWatchlists {
			voiceid_ListWatchlists(cfg, client)
			return
		}
		if _voiceidOptOutSpeaker {
			voiceid_OptOutSpeaker(cfg, client)
			return
		}
		if _voiceidStartFraudsterRegistrationJob {
			voiceid_StartFraudsterRegistrationJob(cfg, client)
			return
		}
		if _voiceidStartSpeakerEnrollmentJob {
			voiceid_StartSpeakerEnrollmentJob(cfg, client)
			return
		}
		if _voiceidTagResource {
			voiceid_TagResource(cfg, client)
			return
		}
		if _voiceidUntagResource {
			voiceid_UntagResource(cfg, client)
			return
		}
		if _voiceidUpdateDomain {
			voiceid_UpdateDomain(cfg, client)
			return
		}
		if _voiceidUpdateWatchlist {
			voiceid_UpdateWatchlist(cfg, client)
			return
		}

	},
}

var (
	_voiceidAssociateFraudster               bool
	_voiceidCreateDomain                     bool
	_voiceidCreateWatchlist                  bool
	_voiceidDeleteDomain                     bool
	_voiceidDeleteFraudster                  bool
	_voiceidDeleteSpeaker                    bool
	_voiceidDeleteWatchlist                  bool
	_voiceidDescribeDomain                   bool
	_voiceidDescribeFraudster                bool
	_voiceidDescribeFraudsterRegistrationJob bool
	_voiceidDescribeSpeaker                  bool
	_voiceidDescribeSpeakerEnrollmentJob     bool
	_voiceidDescribeWatchlist                bool
	_voiceidDisassociateFraudster            bool
	_voiceidEvaluateSession                  bool
	_voiceidListDomains                      bool
	_voiceidListFraudsterRegistrationJobs    bool
	_voiceidListFraudsters                   bool
	_voiceidListSpeakerEnrollmentJobs        bool
	_voiceidListSpeakers                     bool
	_voiceidListTagsForResource              bool
	_voiceidListWatchlists                   bool
	_voiceidOptOutSpeaker                    bool
	_voiceidStartFraudsterRegistrationJob    bool
	_voiceidStartSpeakerEnrollmentJob        bool
	_voiceidTagResource                      bool
	_voiceidUntagResource                    bool
	_voiceidUpdateDomain                     bool
	_voiceidUpdateWatchlist                  bool

	_voiceidClientToken                       string
	_voiceidDataAccessRoleArn                 string
	_voiceidDescription                       string
	_voiceidDomainId                          string
	_voiceidEnrollmentConfig                  string
	_voiceidFraudsterId                       string
	_voiceidInputDataConfig                   string
	_voiceidJobId                             string
	_voiceidJobName                           string
	_voiceidJobStatus                         string
	_voiceidMaxResults                        string
	_voiceidName                              string
	_voiceidNextToken                         string
	_voiceidOutputDataConfig                  string
	_voiceidRegistrationConfig                string
	_voiceidResourceArn                       string
	_voiceidServerSideEncryptionConfiguration string
	_voiceidSessionNameOrId                   string
	_voiceidSpeakerId                         string
	_voiceidTagKeys                           []string
	_voiceidTags                              string
	_voiceidWatchlistId                       string
)

// Associates the fraudsters with the watchlist specified in the same domain.
func voiceid_AssociateFraudster(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.AssociateFraudsterInput{
		// DomainId: *string, // Required
		// FraudsterId: *string, // Required
		// WatchlistId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidFraudsterId) > 0 {
		input.FraudsterId = aws.String(_voiceidFraudsterId)
	}
	if len(_voiceidWatchlistId) > 0 {
		input.WatchlistId = aws.String(_voiceidWatchlistId)
	}

	if resp, err := client.AssociateFraudster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a domain that contains all Amazon Connect Voice ID data, such as
// speakers, fraudsters, customer audio, and voiceprints. Every domain is created
// with a default watchlist that fraudsters can be a part of.
func voiceid_CreateDomain(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.CreateDomainInput{
		// Name: *string, // Required
		// ServerSideEncryptionConfiguration: *types.ServerSideEncryptionConfiguration, // Required
	}

	if len(_voiceidName) > 0 {
		input.Name = aws.String(_voiceidName)
	}
	if len(_voiceidServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _voiceidServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_voiceidClientToken) > 0 {
		input.ClientToken = aws.String(_voiceidClientToken)
	}
	if len(_voiceidDescription) > 0 {
		input.Description = aws.String(_voiceidDescription)
	}
	if len(_voiceidTags) > 0 {
		if err := assignInputField(input, "Tags", _voiceidTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a watchlist that fraudsters can be a part of.
func voiceid_CreateWatchlist(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.CreateWatchlistInput{
		// DomainId: *string, // Required
		// Name: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidName) > 0 {
		input.Name = aws.String(_voiceidName)
	}
	if len(_voiceidClientToken) > 0 {
		input.ClientToken = aws.String(_voiceidClientToken)
	}
	if len(_voiceidDescription) > 0 {
		input.Description = aws.String(_voiceidDescription)
	}

	if resp, err := client.CreateWatchlist(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified domain from Voice ID.
func voiceid_DeleteDomain(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DeleteDomainInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified fraudster from Voice ID. This action disassociates the
// fraudster from any watchlists it is a part of.
func voiceid_DeleteFraudster(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DeleteFraudsterInput{
		// DomainId: *string, // Required
		// FraudsterId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidFraudsterId) > 0 {
		input.FraudsterId = aws.String(_voiceidFraudsterId)
	}

	if resp, err := client.DeleteFraudster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified speaker from Voice ID.
func voiceid_DeleteSpeaker(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DeleteSpeakerInput{
		// DomainId: *string, // Required
		// SpeakerId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidSpeakerId) > 0 {
		input.SpeakerId = aws.String(_voiceidSpeakerId)
	}

	if resp, err := client.DeleteSpeaker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified watchlist from Voice ID. This API throws an exception
// when there are fraudsters in the watchlist that you are trying to delete. You
// must delete the fraudsters, and then delete the watchlist. Every domain has a
// default watchlist which cannot be deleted.
func voiceid_DeleteWatchlist(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DeleteWatchlistInput{
		// DomainId: *string, // Required
		// WatchlistId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidWatchlistId) > 0 {
		input.WatchlistId = aws.String(_voiceidWatchlistId)
	}

	if resp, err := client.DeleteWatchlist(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified domain.
func voiceid_DescribeDomain(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DescribeDomainInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}

	if resp, err := client.DescribeDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified fraudster.
func voiceid_DescribeFraudster(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DescribeFraudsterInput{
		// DomainId: *string, // Required
		// FraudsterId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidFraudsterId) > 0 {
		input.FraudsterId = aws.String(_voiceidFraudsterId)
	}

	if resp, err := client.DescribeFraudster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified fraudster registration job.
func voiceid_DescribeFraudsterRegistrationJob(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DescribeFraudsterRegistrationJobInput{
		// DomainId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidJobId) > 0 {
		input.JobId = aws.String(_voiceidJobId)
	}

	if resp, err := client.DescribeFraudsterRegistrationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified speaker.
func voiceid_DescribeSpeaker(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DescribeSpeakerInput{
		// DomainId: *string, // Required
		// SpeakerId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidSpeakerId) > 0 {
		input.SpeakerId = aws.String(_voiceidSpeakerId)
	}

	if resp, err := client.DescribeSpeaker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified speaker enrollment job.
func voiceid_DescribeSpeakerEnrollmentJob(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DescribeSpeakerEnrollmentJobInput{
		// DomainId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidJobId) > 0 {
		input.JobId = aws.String(_voiceidJobId)
	}

	if resp, err := client.DescribeSpeakerEnrollmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the specified watchlist.
func voiceid_DescribeWatchlist(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DescribeWatchlistInput{
		// DomainId: *string, // Required
		// WatchlistId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidWatchlistId) > 0 {
		input.WatchlistId = aws.String(_voiceidWatchlistId)
	}

	if resp, err := client.DescribeWatchlist(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the fraudsters from the watchlist specified. Voice ID always
// expects a fraudster to be a part of at least one watchlist. If you try to
// disassociate a fraudster from its only watchlist, a ValidationException is
// thrown.
func voiceid_DisassociateFraudster(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.DisassociateFraudsterInput{
		// DomainId: *string, // Required
		// FraudsterId: *string, // Required
		// WatchlistId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidFraudsterId) > 0 {
		input.FraudsterId = aws.String(_voiceidFraudsterId)
	}
	if len(_voiceidWatchlistId) > 0 {
		input.WatchlistId = aws.String(_voiceidWatchlistId)
	}

	if resp, err := client.DisassociateFraudster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates a specified session based on audio data accumulated during a
// streaming Amazon Connect Voice ID call.
func voiceid_EvaluateSession(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.EvaluateSessionInput{
		// DomainId: *string, // Required
		// SessionNameOrId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidSessionNameOrId) > 0 {
		input.SessionNameOrId = aws.String(_voiceidSessionNameOrId)
	}

	if resp, err := client.EvaluateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the domains in the Amazon Web Services account.
func voiceid_ListDomains(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListDomainsInput{}

	if len(_voiceidMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _voiceidMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_voiceidNextToken) > 0 {
		input.NextToken = aws.String(_voiceidNextToken)
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

	var results []*voiceid.ListDomainsOutput
	p := voiceid.NewListDomainsPaginator(client, input)
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

// Lists all the fraudster registration jobs in the domain with the given JobStatus
// . If JobStatus is not provided, this lists all fraudster registration jobs in
// the given domain.
func voiceid_ListFraudsterRegistrationJobs(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListFraudsterRegistrationJobsInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _voiceidJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_voiceidMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _voiceidMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_voiceidNextToken) > 0 {
		input.NextToken = aws.String(_voiceidNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFraudsterRegistrationJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*voiceid.ListFraudsterRegistrationJobsOutput
	p := voiceid.NewListFraudsterRegistrationJobsPaginator(client, input)
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

// Lists all fraudsters in a specified watchlist or domain.
func voiceid_ListFraudsters(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListFraudstersInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _voiceidMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_voiceidNextToken) > 0 {
		input.NextToken = aws.String(_voiceidNextToken)
	}
	if len(_voiceidWatchlistId) > 0 {
		input.WatchlistId = aws.String(_voiceidWatchlistId)
	}

	if disablePaginator() {
		if resp, err := client.ListFraudsters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*voiceid.ListFraudstersOutput
	p := voiceid.NewListFraudstersPaginator(client, input)
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

// Lists all the speaker enrollment jobs in the domain with the specified JobStatus
// . If JobStatus is not provided, this lists all jobs with all possible speaker
// enrollment job statuses.
func voiceid_ListSpeakerEnrollmentJobs(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListSpeakerEnrollmentJobsInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidJobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _voiceidJobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_voiceidMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _voiceidMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_voiceidNextToken) > 0 {
		input.NextToken = aws.String(_voiceidNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSpeakerEnrollmentJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*voiceid.ListSpeakerEnrollmentJobsOutput
	p := voiceid.NewListSpeakerEnrollmentJobsPaginator(client, input)
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

// Lists all speakers in a specified domain.
func voiceid_ListSpeakers(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListSpeakersInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _voiceidMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_voiceidNextToken) > 0 {
		input.NextToken = aws.String(_voiceidNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSpeakers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*voiceid.ListSpeakersOutput
	p := voiceid.NewListSpeakersPaginator(client, input)
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

// Lists all tags associated with a specified Voice ID resource.
func voiceid_ListTagsForResource(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_voiceidResourceArn) > 0 {
		input.ResourceArn = aws.String(_voiceidResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all watchlists in a specified domain.
func voiceid_ListWatchlists(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.ListWatchlistsInput{
		// DomainId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _voiceidMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_voiceidNextToken) > 0 {
		input.NextToken = aws.String(_voiceidNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWatchlists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*voiceid.ListWatchlistsOutput
	p := voiceid.NewListWatchlistsPaginator(client, input)
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

// Opts out a speaker from Voice ID. A speaker can be opted out regardless of
// whether or not they already exist in Voice ID. If they don't yet exist, a new
// speaker is created in an opted out state. If they already exist, their existing
// status is overridden and they are opted out. Enrollment and evaluation
// authentication requests are rejected for opted out speakers, and opted out
// speakers have no voice embeddings stored in Voice ID.
func voiceid_OptOutSpeaker(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.OptOutSpeakerInput{
		// DomainId: *string, // Required
		// SpeakerId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidSpeakerId) > 0 {
		input.SpeakerId = aws.String(_voiceidSpeakerId)
	}

	if resp, err := client.OptOutSpeaker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new batch fraudster registration job using provided details.
func voiceid_StartFraudsterRegistrationJob(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.StartFraudsterRegistrationJobInput{
		// DataAccessRoleArn: *string, // Required
		// DomainId: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_voiceidDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_voiceidDataAccessRoleArn)
	}
	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _voiceidInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_voiceidOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _voiceidOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_voiceidClientToken) > 0 {
		input.ClientToken = aws.String(_voiceidClientToken)
	}
	if len(_voiceidJobName) > 0 {
		input.JobName = aws.String(_voiceidJobName)
	}
	if len(_voiceidRegistrationConfig) > 0 {
		if err := assignInputField(input, "RegistrationConfig", _voiceidRegistrationConfig); err != nil {
			log.Errorf("invalid --registration-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartFraudsterRegistrationJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new batch speaker enrollment job using specified details.
func voiceid_StartSpeakerEnrollmentJob(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.StartSpeakerEnrollmentJobInput{
		// DataAccessRoleArn: *string, // Required
		// DomainId: *string, // Required
		// InputDataConfig: *types.InputDataConfig, // Required
		// OutputDataConfig: *types.OutputDataConfig, // Required
	}

	if len(_voiceidDataAccessRoleArn) > 0 {
		input.DataAccessRoleArn = aws.String(_voiceidDataAccessRoleArn)
	}
	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidInputDataConfig) > 0 {
		if err := assignInputField(input, "InputDataConfig", _voiceidInputDataConfig); err != nil {
			log.Errorf("invalid --input-data-config: %s", err.Error())
			return
		}
	}
	if len(_voiceidOutputDataConfig) > 0 {
		if err := assignInputField(input, "OutputDataConfig", _voiceidOutputDataConfig); err != nil {
			log.Errorf("invalid --output-data-config: %s", err.Error())
			return
		}
	}
	if len(_voiceidClientToken) > 0 {
		input.ClientToken = aws.String(_voiceidClientToken)
	}
	if len(_voiceidEnrollmentConfig) > 0 {
		if err := assignInputField(input, "EnrollmentConfig", _voiceidEnrollmentConfig); err != nil {
			log.Errorf("invalid --enrollment-config: %s", err.Error())
			return
		}
	}
	if len(_voiceidJobName) > 0 {
		input.JobName = aws.String(_voiceidJobName)
	}

	if resp, err := client.StartSpeakerEnrollmentJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a Voice ID resource with the provided list of tags.
func voiceid_TagResource(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_voiceidResourceArn) > 0 {
		input.ResourceArn = aws.String(_voiceidResourceArn)
	}
	if len(_voiceidTags) > 0 {
		if err := assignInputField(input, "Tags", _voiceidTags); err != nil {
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

// Removes specified tags from a specified Amazon Connect Voice ID resource.
func voiceid_UntagResource(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_voiceidResourceArn) > 0 {
		input.ResourceArn = aws.String(_voiceidResourceArn)
	}
	if len(_voiceidTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _voiceidTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified domain. This API has clobber behavior, and clears and
// replaces all attributes. If an optional field, such as 'Description' is not
// provided, it is removed from the domain.
func voiceid_UpdateDomain(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.UpdateDomainInput{
		// DomainId: *string, // Required
		// Name: *string, // Required
		// ServerSideEncryptionConfiguration: *types.ServerSideEncryptionConfiguration, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidName) > 0 {
		input.Name = aws.String(_voiceidName)
	}
	if len(_voiceidServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _voiceidServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_voiceidDescription) > 0 {
		input.Description = aws.String(_voiceidDescription)
	}

	if resp, err := client.UpdateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified watchlist. Every domain has a default watchlist which
// cannot be updated.
func voiceid_UpdateWatchlist(cfg aws.Config, client *voiceid.Client) {
	input := &voiceid.UpdateWatchlistInput{
		// DomainId: *string, // Required
		// WatchlistId: *string, // Required
	}

	if len(_voiceidDomainId) > 0 {
		input.DomainId = aws.String(_voiceidDomainId)
	}
	if len(_voiceidWatchlistId) > 0 {
		input.WatchlistId = aws.String(_voiceidWatchlistId)
	}
	if len(_voiceidDescription) > 0 {
		input.Description = aws.String(_voiceidDescription)
	}
	if len(_voiceidName) > 0 {
		input.Name = aws.String(_voiceidName)
	}

	if resp, err := client.UpdateWatchlist(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_voiceidCmd)
	_voiceidCmd.Flags().SortFlags = false

	_voiceidCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_voiceidCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_voiceidCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_voiceidCmd.Flags().StringVarP(&_voiceidClientToken, "client-token", "", "", "Client Token")
	_voiceidCmd.Flags().StringVarP(&_voiceidDataAccessRoleArn, "data-access-role-arn", "", "", "Data Access Role ARN")
	_voiceidCmd.Flags().StringVarP(&_voiceidDescription, "description", "", "", "Description")
	_voiceidCmd.Flags().StringVarP(&_voiceidDomainId, "domain-id", "", "", "Domain ID")
	_voiceidCmd.Flags().StringVarP(&_voiceidEnrollmentConfig, "enrollment-config", "", "", "Enrollment Config")
	_voiceidCmd.Flags().StringVarP(&_voiceidFraudsterId, "fraudster-id", "", "", "Fraudster ID")
	_voiceidCmd.Flags().StringVarP(&_voiceidInputDataConfig, "input-data-config", "", "", "Input Data Config")
	_voiceidCmd.Flags().StringVarP(&_voiceidJobId, "job-id", "", "", "Job ID")
	_voiceidCmd.Flags().StringVarP(&_voiceidJobName, "job-name", "", "", "Job Name")
	_voiceidCmd.Flags().StringVarP(&_voiceidJobStatus, "job-status", "", "", "Job Status")
	_voiceidCmd.Flags().StringVarP(&_voiceidMaxResults, "max-results", "", "", "Max Results")
	_voiceidCmd.Flags().StringVarP(&_voiceidName, "name", "", "", "Name")
	_voiceidCmd.Flags().StringVarP(&_voiceidNextToken, "next-token", "", "", "Next Token")
	_voiceidCmd.Flags().StringVarP(&_voiceidOutputDataConfig, "output-data-config", "", "", "Output Data Config")
	_voiceidCmd.Flags().StringVarP(&_voiceidRegistrationConfig, "registration-config", "", "", "Registration Config")
	_voiceidCmd.Flags().StringVarP(&_voiceidResourceArn, "resource-arn", "", "", "Resource ARN")
	_voiceidCmd.Flags().StringVarP(&_voiceidServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_voiceidCmd.Flags().StringVarP(&_voiceidSessionNameOrId, "session-name-or-id", "", "", "Session Name Or ID")
	_voiceidCmd.Flags().StringVarP(&_voiceidSpeakerId, "speaker-id", "", "", "Speaker ID")
	_voiceidCmd.Flags().StringSliceVarP(&_voiceidTagKeys, "tag-keys", "", nil, "Tag Keys")
	_voiceidCmd.Flags().StringVarP(&_voiceidTags, "tags", "", "", "Tags")
	_voiceidCmd.Flags().StringVarP(&_voiceidWatchlistId, "watchlist-id", "", "", "Watchlist ID")

	_voiceidCmd.Flags().BoolVarP(&_voiceidAssociateFraudster, "associate-fraudster", "", false, "Associate Fraudster")
	_voiceidCmd.Flags().BoolVarP(&_voiceidCreateDomain, "create-domain", "", false, "Create Domain")
	_voiceidCmd.Flags().BoolVarP(&_voiceidCreateWatchlist, "create-watchlist", "", false, "Create Watchlist")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDeleteFraudster, "delete-fraudster", "", false, "Delete Fraudster")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDeleteSpeaker, "delete-speaker", "", false, "Delete Speaker")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDeleteWatchlist, "delete-watchlist", "", false, "Delete Watchlist")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDescribeDomain, "describe-domain", "", false, "Describe Domain")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDescribeFraudster, "describe-fraudster", "", false, "Describe Fraudster")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDescribeFraudsterRegistrationJob, "describe-fraudster-registration-job", "", false, "Describe Fraudster Registration Job")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDescribeSpeaker, "describe-speaker", "", false, "Describe Speaker")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDescribeSpeakerEnrollmentJob, "describe-speaker-enrollment-job", "", false, "Describe Speaker Enrollment Job")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDescribeWatchlist, "describe-watchlist", "", false, "Describe Watchlist")
	_voiceidCmd.Flags().BoolVarP(&_voiceidDisassociateFraudster, "disassociate-fraudster", "", false, "Disassociate Fraudster")
	_voiceidCmd.Flags().BoolVarP(&_voiceidEvaluateSession, "evaluate-session", "", false, "Evaluate Session")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListDomains, "list-domains", "", false, "List Domains")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListFraudsterRegistrationJobs, "list-fraudster-registration-jobs", "", false, "List Fraudster Registration Jobs")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListFraudsters, "list-fraudsters", "", false, "List Fraudsters")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListSpeakerEnrollmentJobs, "list-speaker-enrollment-jobs", "", false, "List Speaker Enrollment Jobs")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListSpeakers, "list-speakers", "", false, "List Speakers")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_voiceidCmd.Flags().BoolVarP(&_voiceidListWatchlists, "list-watchlists", "", false, "List Watchlists")
	_voiceidCmd.Flags().BoolVarP(&_voiceidOptOutSpeaker, "opt-out-speaker", "", false, "Opt Out Speaker")
	_voiceidCmd.Flags().BoolVarP(&_voiceidStartFraudsterRegistrationJob, "start-fraudster-registration-job", "", false, "Start Fraudster Registration Job")
	_voiceidCmd.Flags().BoolVarP(&_voiceidStartSpeakerEnrollmentJob, "start-speaker-enrollment-job", "", false, "Start Speaker Enrollment Job")
	_voiceidCmd.Flags().BoolVarP(&_voiceidTagResource, "tag-resource", "", false, "Tag Resource")
	_voiceidCmd.Flags().BoolVarP(&_voiceidUntagResource, "untag-resource", "", false, "Untag Resource")
	_voiceidCmd.Flags().BoolVarP(&_voiceidUpdateDomain, "update-domain", "", false, "Update Domain")
	_voiceidCmd.Flags().BoolVarP(&_voiceidUpdateWatchlist, "update-watchlist", "", false, "Update Watchlist")

}
