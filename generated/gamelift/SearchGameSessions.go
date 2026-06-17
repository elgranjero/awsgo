package gamelift

// SearchGameSessions is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves all active game sessions that match a set of search criteria and
// sorts them into a specified order.
//
// This operation is not designed to continually track game session status because
// that practice can cause you to exceed your API limit and generate errors.
// Instead, configure an Amazon Simple Notification Service (Amazon SNS) topic to
// receive notifications from a matchmaker or a game session placement queue.
//
// When searching for game sessions, you specify exactly where you want to search
// and provide a search filter expression, a sort expression, or both. A search
// request can search only one fleet, but it can search all of a fleet's locations.
//
// This operation can be used in the following ways:
//
// - To search all game sessions that are currently running on all locations in
// a fleet, provide a fleet or alias ID. This approach returns game sessions in the
// fleet's home Region and all remote locations that fit the search criteria.
//
// - To search all game sessions that are currently running on a specific fleet
// location, provide a fleet or alias ID and a location name. For location, you can
// specify a fleet's home Region or any remote location.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, a GameSession object is returned for each game session that
// matches the request. Search finds game sessions that are in ACTIVE status only.
// To retrieve information on game sessions in other statuses, use [DescribeGameSessions].
//
// To set search and sort criteria, create a filter expression using the following
// game session attributes. For game session search examples, see the Examples
// section of this topic.
//
// - gameSessionId -- A unique identifier for the game session. You can use
// either a GameSessionId or GameSessionArn value.
//
// - gameSessionName -- Name assigned to a game session. Game session names do
// not need to be unique to a game session.
//
// - gameSessionProperties -- A set of key-value pairs that can store custom
// data in a game session. For example: {"Key": "difficulty", "Value": "novice"}
// . The filter expression must specify the [https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameProperty]-- a Key and a string Value to search
// for the game sessions.
//
// For example, to search for the above key-value pair, specify the following
//
// search filter: gameSessionProperties.difficulty = "novice" . All game property
// values are searched as strings.
//
// For examples of searching game sessions, see the ones below, and also see [Search game sessions by game property].
//
// Avoid using periods (".") in property keys if you plan to search for game
//
// sessions by properties. Property keys containing periods cannot be searched and
// will be filtered out from search results due to search index limitations.
//
// - maximumSessions -- Maximum number of player sessions allowed for a game
// session.
//
// - creationTimeMillis -- Value indicating when a game session was created. It
// is expressed in Unix time as milliseconds.
//
// - playerSessionCount -- Number of players currently connected to a game
// session. This value changes rapidly as players join the session or drop out.
//
// - hasAvailablePlayerSessions -- Boolean value indicating whether a game
// session has reached its maximum number of players. It is highly recommended that
// all search requests include this filter attribute to optimize search performance
// and return only sessions that players can join.
//
// Returned values for playerSessionCount and hasAvailablePlayerSessions change
// quickly as players join sessions and others drop out. Results should be
// considered a snapshot in time. Be sure to refresh search results often, and
// handle sessions that fill up before a player can join.
//
// [All APIs by task]
//
// [Search game sessions by game property]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-client-api.html#game-properties-search
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameProperty]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_GameProperty
// [DescribeGameSessions]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeGameSessions.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
