package cmd

func Execute(args []string) error {
	if p := _marketplaceentitlementserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplaceentitlementserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplaceentitlementserviceCmd.SetArgs(args)
	return _marketplaceentitlementserviceCmd.Execute()
}
