package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// personalizeruntimeCmd represents the personalizeruntime command
var _personalizeruntimeCmd = &cobra.Command{
	Use:   "personalizeruntime",
	Short: "AWS personalizeruntime CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := personalizeruntime.NewFromConfig(cfg)
		if _personalizeruntimeGetActionRecommendations {
			personalizeruntime_GetActionRecommendations(cfg, client)
			return
		}
		if _personalizeruntimeGetPersonalizedRanking {
			personalizeruntime_GetPersonalizedRanking(cfg, client)
			return
		}
		if _personalizeruntimeGetRecommendations {
			personalizeruntime_GetRecommendations(cfg, client)
			return
		}

	},
}

var (
	_personalizeruntimeGetActionRecommendations bool
	_personalizeruntimeGetPersonalizedRanking   bool
	_personalizeruntimeGetRecommendations       bool

	_personalizeruntimeCampaignArn     string
	_personalizeruntimeContext         string
	_personalizeruntimeFilterArn       string
	_personalizeruntimeFilterValues    string
	_personalizeruntimeInputList       []string
	_personalizeruntimeItemId          string
	_personalizeruntimeMetadataColumns string
	_personalizeruntimeNumResults      string
	_personalizeruntimePromotions      string
	_personalizeruntimeRecommenderArn  string
	_personalizeruntimeUserId          string
)

// Returns a list of recommended actions in sorted in descending order by
// prediction score. Use the GetActionRecommendations API if you have a custom
// campaign that deploys a solution version trained with a PERSONALIZED_ACTIONS
// recipe.
//
// For more information about PERSONALIZED_ACTIONS recipes, see [PERSONALIZED_ACTIONS recipes]. For more
// information about getting action recommendations, see [Getting action recommendations].
//
// [Getting action recommendations]: https://docs.aws.amazon.com/personalize/latest/dg/get-action-recommendations.html
// [PERSONALIZED_ACTIONS recipes]: https://docs.aws.amazon.com/personalize/latest/dg/nexts-best-action-recipes.html
func personalizeruntime_GetActionRecommendations(cfg aws.Config, client *personalizeruntime.Client) {
	input := &personalizeruntime.GetActionRecommendationsInput{}

	if len(_personalizeruntimeCampaignArn) > 0 {
		input.CampaignArn = aws.String(_personalizeruntimeCampaignArn)
	}
	if len(_personalizeruntimeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeruntimeFilterArn)
	}
	if len(_personalizeruntimeFilterValues) > 0 {
		if err := assignInputField(input, "FilterValues", _personalizeruntimeFilterValues); err != nil {
			log.Errorf("invalid --filter-values: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeNumResults) > 0 {
		if err := assignInputField(input, "NumResults", _personalizeruntimeNumResults); err != nil {
			log.Errorf("invalid --num-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeUserId) > 0 {
		input.UserId = aws.String(_personalizeruntimeUserId)
	}

	if resp, err := client.GetActionRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Re-ranks a list of recommended items for the given user. The first item in the
// list is deemed the most likely item to be of interest to the user.
//
// The solution backing the campaign must have been created using a recipe of type
// PERSONALIZED_RANKING.
func personalizeruntime_GetPersonalizedRanking(cfg aws.Config, client *personalizeruntime.Client) {
	input := &personalizeruntime.GetPersonalizedRankingInput{
		// CampaignArn: *string, // Required
		// InputList: []string, // Required
		// UserId: *string, // Required
	}

	if len(_personalizeruntimeCampaignArn) > 0 {
		input.CampaignArn = aws.String(_personalizeruntimeCampaignArn)
	}
	if len(_personalizeruntimeInputList) > 0 {
		input.InputList = append([]string(nil), _personalizeruntimeInputList...)
	}
	if len(_personalizeruntimeUserId) > 0 {
		input.UserId = aws.String(_personalizeruntimeUserId)
	}
	if len(_personalizeruntimeContext) > 0 {
		if err := assignInputField(input, "Context", _personalizeruntimeContext); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeruntimeFilterArn)
	}
	if len(_personalizeruntimeFilterValues) > 0 {
		if err := assignInputField(input, "FilterValues", _personalizeruntimeFilterValues); err != nil {
			log.Errorf("invalid --filter-values: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeMetadataColumns) > 0 {
		if err := assignInputField(input, "MetadataColumns", _personalizeruntimeMetadataColumns); err != nil {
			log.Errorf("invalid --metadata-columns: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPersonalizedRanking(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of recommended items. For campaigns, the campaign's Amazon
// Resource Name (ARN) is required and the required user and item input depends on
// the recipe type used to create the solution backing the campaign as follows:
//
// - USER_PERSONALIZATION - userId required, itemId not used
//
// - RELATED_ITEMS - itemId required, userId not used
//
// Campaigns that are backed by a solution created using a recipe of type
// PERSONALIZED_RANKING use the API.
//
// For recommenders, the recommender's ARN is required and the required item and
// user input depends on the use case (domain-based recipe) backing the
// recommender. For information on use case requirements see [Choosing recommender use cases].
//
// [Choosing recommender use cases]: https://docs.aws.amazon.com/personalize/latest/dg/domain-use-cases.html
func personalizeruntime_GetRecommendations(cfg aws.Config, client *personalizeruntime.Client) {
	input := &personalizeruntime.GetRecommendationsInput{}

	if len(_personalizeruntimeCampaignArn) > 0 {
		input.CampaignArn = aws.String(_personalizeruntimeCampaignArn)
	}
	if len(_personalizeruntimeContext) > 0 {
		if err := assignInputField(input, "Context", _personalizeruntimeContext); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeFilterArn) > 0 {
		input.FilterArn = aws.String(_personalizeruntimeFilterArn)
	}
	if len(_personalizeruntimeFilterValues) > 0 {
		if err := assignInputField(input, "FilterValues", _personalizeruntimeFilterValues); err != nil {
			log.Errorf("invalid --filter-values: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeItemId) > 0 {
		input.ItemId = aws.String(_personalizeruntimeItemId)
	}
	if len(_personalizeruntimeMetadataColumns) > 0 {
		if err := assignInputField(input, "MetadataColumns", _personalizeruntimeMetadataColumns); err != nil {
			log.Errorf("invalid --metadata-columns: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeNumResults) > 0 {
		if err := assignInputField(input, "NumResults", _personalizeruntimeNumResults); err != nil {
			log.Errorf("invalid --num-results: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimePromotions) > 0 {
		if err := assignInputField(input, "Promotions", _personalizeruntimePromotions); err != nil {
			log.Errorf("invalid --promotions: %s", err.Error())
			return
		}
	}
	if len(_personalizeruntimeRecommenderArn) > 0 {
		input.RecommenderArn = aws.String(_personalizeruntimeRecommenderArn)
	}
	if len(_personalizeruntimeUserId) > 0 {
		input.UserId = aws.String(_personalizeruntimeUserId)
	}

	if resp, err := client.GetRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_personalizeruntimeCmd)
	_personalizeruntimeCmd.Flags().SortFlags = false

	_personalizeruntimeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_personalizeruntimeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_personalizeruntimeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeCampaignArn, "campaign-arn", "", "", "Campaign ARN")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeContext, "context", "", "", "Context")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeFilterArn, "filter-arn", "", "", "Filter ARN")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeFilterValues, "filter-values", "", "", "Filter Values")
	_personalizeruntimeCmd.Flags().StringSliceVarP(&_personalizeruntimeInputList, "input-list", "", nil, "Input List")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeItemId, "item-id", "", "", "Item ID")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeMetadataColumns, "metadata-columns", "", "", "Metadata Columns")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeNumResults, "num-results", "", "", "Num Results")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimePromotions, "promotions", "", "", "Promotions")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeRecommenderArn, "recommender-arn", "", "", "Recommender ARN")
	_personalizeruntimeCmd.Flags().StringVarP(&_personalizeruntimeUserId, "user-id", "", "", "User ID")

	_personalizeruntimeCmd.Flags().BoolVarP(&_personalizeruntimeGetActionRecommendations, "get-action-recommendations", "", false, "Get Action Recommendations")
	_personalizeruntimeCmd.Flags().BoolVarP(&_personalizeruntimeGetPersonalizedRanking, "get-personalized-ranking", "", false, "Get Personalized Ranking")
	_personalizeruntimeCmd.Flags().BoolVarP(&_personalizeruntimeGetRecommendations, "get-recommendations", "", false, "Get Recommendations")

}
