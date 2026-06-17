package budgets

// CreateBudget is generated as a reference stub.
// Executable command wiring lives under cmd/budgets.go.
//
// Creates a budget and, if included, notifications and subscribers.
//
// Only one of BudgetLimit or PlannedBudgetLimits can be present in the syntax at
// one time. Use the syntax that matches your use case. The Request Syntax section
// shows the BudgetLimit syntax. For PlannedBudgetLimits , see the [Examples] section.
//
// Similarly, only one set of filter and metric selections can be present in the
// syntax at one time. Either FilterExpression and Metrics or CostFilters and
// CostTypes , not both or a different combination. We recommend using
// FilterExpression and Metrics as they provide more flexible and powerful
// filtering capabilities. The Request Syntax section shows the FilterExpression /
// Metrics syntax.
//
// [Examples]: https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_CreateBudget.html#API_CreateBudget_Examples
