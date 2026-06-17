package cloudwatch

// PutDashboard is generated as a reference stub.
// Executable command wiring lives under cmd/cloudwatch.go.
//
// Creates a dashboard if it does not already exist, or updates an existing
// dashboard. If you update a dashboard, the entire contents are replaced with what
// you specify here.
//
// All dashboards in your account are global, not region-specific.
//
// A simple way to create a dashboard using PutDashboard is to copy an existing
// dashboard. To copy an existing dashboard using the console, you can load the
// dashboard and then use the View/edit source command in the Actions menu to
// display the JSON block for that dashboard. Another way to copy a dashboard is to
// use GetDashboard , and then use the data returned within DashboardBody as the
// template for the new dashboard when you call PutDashboard .
//
// When you create a dashboard with PutDashboard , a good practice is to add a text
// widget at the top of the dashboard with a message that the dashboard was created
// by script and should not be changed in the console. This message could also
// point console users to the location of the DashboardBody script or the
// CloudFormation template used to create the dashboard.
