package quicksight

// CreateAccountCustomization is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Creates Amazon Quick Sight customizations. Currently, you can add a custom
// default theme by using the CreateAccountCustomization or
// UpdateAccountCustomization API operation. To further customize Amazon Quick
// Sight by removing Amazon Quick Sight sample assets and videos for all new users,
// see [Customizing Quick Sight]in the Amazon Quick Sight User Guide.
//
// You can create customizations for your Amazon Web Services account or, if you
// specify a namespace, for a Quick Sight namespace instead. Customizations that
// apply to a namespace always override customizations that apply to an Amazon Web
// Services account. To find out which customizations apply, use the
// DescribeAccountCustomization API operation.
//
// Before you use the CreateAccountCustomization API operation to add a theme as
// the namespace default, make sure that you first share the theme with the
// namespace. If you don't share it with the namespace, the theme isn't visible to
// your users even if you make it the default theme. To check if the theme is
// shared, view the current permissions by using the [DescribeThemePermissions]API operation. To share the
// theme, grant permissions by using the [UpdateThemePermissions]API operation.
//
// [UpdateThemePermissions]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UpdateThemePermissions.html
// [DescribeThemePermissions]: https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DescribeThemePermissions.html
// [Customizing Quick Sight]: https://docs.aws.amazon.com/quicksight/latest/user/customizing-quicksight.html
