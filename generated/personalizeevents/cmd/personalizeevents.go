package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// personalizeeventsCmd represents the personalizeevents command
var _personalizeeventsCmd = &cobra.Command{
	Use:   "personalizeevents",
	Short: "AWS personalizeevents CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := personalizeevents.NewFromConfig(cfg)
		if _personalizeeventsPutActionInteractions {
			personalizeevents_PutActionInteractions(cfg, client)
			return
		}
		if _personalizeeventsPutActions {
			personalizeevents_PutActions(cfg, client)
			return
		}
		if _personalizeeventsPutEvents {
			personalizeevents_PutEvents(cfg, client)
			return
		}
		if _personalizeeventsPutItems {
			personalizeevents_PutItems(cfg, client)
			return
		}
		if _personalizeeventsPutUsers {
			personalizeevents_PutUsers(cfg, client)
			return
		}

	},
}

var (
	_personalizeeventsPutActionInteractions bool
	_personalizeeventsPutActions            bool
	_personalizeeventsPutEvents             bool
	_personalizeeventsPutItems              bool
	_personalizeeventsPutUsers              bool

	_personalizeeventsActionInteractions string
	_personalizeeventsActions            string
	_personalizeeventsDatasetArn         string
	_personalizeeventsEventList          string
	_personalizeeventsItems              string
	_personalizeeventsSessionId          string
	_personalizeeventsTrackingId         string
	_personalizeeventsUserId             string
	_personalizeeventsUsers              string
)

// Records action interaction event data. An action interaction event is an
// interaction between a user and an action. For example, a user taking an action,
// such a enrolling in a membership program or downloading your app.
//
// For more information about recording action interactions, see [Recording action interaction events]. For more
// information about actions in an Actions dataset, see [Actions dataset].
//
// [Recording action interaction events]: https://docs.aws.amazon.com/personalize/latest/dg/recording-action-interaction-events.html
// [Actions dataset]: https://docs.aws.amazon.com/personalize/latest/dg/actions-datasets.html
func personalizeevents_PutActionInteractions(cfg aws.Config, client *personalizeevents.Client) {
	input := &personalizeevents.PutActionInteractionsInput{
		// ActionInteractions: []types.ActionInteraction, // Required
		// TrackingId: *string, // Required
	}

	if len(_personalizeeventsActionInteractions) > 0 {
		if err := assignInputField(input, "ActionInteractions", _personalizeeventsActionInteractions); err != nil {
			log.Errorf("invalid --action-interactions: %s", err.Error())
			return
		}
	}
	if len(_personalizeeventsTrackingId) > 0 {
		input.TrackingId = aws.String(_personalizeeventsTrackingId)
	}

	if resp, err := client.PutActionInteractions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more actions to an Actions dataset. For more information see [Importing actions individually].
//
// [Importing actions individually]: https://docs.aws.amazon.com/personalize/latest/dg/importing-actions.html
func personalizeevents_PutActions(cfg aws.Config, client *personalizeevents.Client) {
	input := &personalizeevents.PutActionsInput{
		// Actions: []types.Action, // Required
		// DatasetArn: *string, // Required
	}

	if len(_personalizeeventsActions) > 0 {
		if err := assignInputField(input, "Actions", _personalizeeventsActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_personalizeeventsDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeeventsDatasetArn)
	}

	if resp, err := client.PutActions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Records item interaction event data. For more information see [Recording item interaction events].
//
// [Recording item interaction events]: https://docs.aws.amazon.com/personalize/latest/dg/recording-item-interaction-events.html
func personalizeevents_PutEvents(cfg aws.Config, client *personalizeevents.Client) {
	input := &personalizeevents.PutEventsInput{
		// EventList: []types.Event, // Required
		// SessionId: *string, // Required
		// TrackingId: *string, // Required
	}

	if len(_personalizeeventsEventList) > 0 {
		if err := assignInputField(input, "EventList", _personalizeeventsEventList); err != nil {
			log.Errorf("invalid --event-list: %s", err.Error())
			return
		}
	}
	if len(_personalizeeventsSessionId) > 0 {
		input.SessionId = aws.String(_personalizeeventsSessionId)
	}
	if len(_personalizeeventsTrackingId) > 0 {
		input.TrackingId = aws.String(_personalizeeventsTrackingId)
	}
	if len(_personalizeeventsUserId) > 0 {
		input.UserId = aws.String(_personalizeeventsUserId)
	}

	if resp, err := client.PutEvents(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more items to an Items dataset. For more information see [Importing items individually].
//
// [Importing items individually]: https://docs.aws.amazon.com/personalize/latest/dg/importing-items.html
func personalizeevents_PutItems(cfg aws.Config, client *personalizeevents.Client) {
	input := &personalizeevents.PutItemsInput{
		// DatasetArn: *string, // Required
		// Items: []types.Item, // Required
	}

	if len(_personalizeeventsDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeeventsDatasetArn)
	}
	if len(_personalizeeventsItems) > 0 {
		if err := assignInputField(input, "Items", _personalizeeventsItems); err != nil {
			log.Errorf("invalid --items: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutItems(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more users to a Users dataset. For more information see [Importing users individually].
//
// [Importing users individually]: https://docs.aws.amazon.com/personalize/latest/dg/importing-users.html
func personalizeevents_PutUsers(cfg aws.Config, client *personalizeevents.Client) {
	input := &personalizeevents.PutUsersInput{
		// DatasetArn: *string, // Required
		// Users: []types.User, // Required
	}

	if len(_personalizeeventsDatasetArn) > 0 {
		input.DatasetArn = aws.String(_personalizeeventsDatasetArn)
	}
	if len(_personalizeeventsUsers) > 0 {
		if err := assignInputField(input, "Users", _personalizeeventsUsers); err != nil {
			log.Errorf("invalid --users: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutUsers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_personalizeeventsCmd)
	_personalizeeventsCmd.Flags().SortFlags = false

	_personalizeeventsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_personalizeeventsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_personalizeeventsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsActionInteractions, "action-interactions", "", "", "Action Interactions")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsActions, "actions", "", "", "Actions")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsDatasetArn, "dataset-arn", "", "", "Dataset ARN")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsEventList, "event-list", "", "", "Event List")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsItems, "items", "", "", "Items")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsSessionId, "session-id", "", "", "Session ID")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsTrackingId, "tracking-id", "", "", "Tracking ID")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsUserId, "user-id", "", "", "User ID")
	_personalizeeventsCmd.Flags().StringVarP(&_personalizeeventsUsers, "users", "", "", "Users")

	_personalizeeventsCmd.Flags().BoolVarP(&_personalizeeventsPutActionInteractions, "put-action-interactions", "", false, "Put Action Interactions")
	_personalizeeventsCmd.Flags().BoolVarP(&_personalizeeventsPutActions, "put-actions", "", false, "Put Actions")
	_personalizeeventsCmd.Flags().BoolVarP(&_personalizeeventsPutEvents, "put-events", "", false, "Put Events")
	_personalizeeventsCmd.Flags().BoolVarP(&_personalizeeventsPutItems, "put-items", "", false, "Put Items")
	_personalizeeventsCmd.Flags().BoolVarP(&_personalizeeventsPutUsers, "put-users", "", false, "Put Users")

}
