package cmd

func Execute(args []string) error {
	if p := _marketplacecommerceanalyticsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplacecommerceanalyticsCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplacecommerceanalyticsCmd.SetArgs(args)
	return _marketplacecommerceanalyticsCmd.Execute()
}
