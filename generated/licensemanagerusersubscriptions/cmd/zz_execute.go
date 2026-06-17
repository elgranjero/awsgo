package cmd

func Execute(args []string) error {
	if p := _licensemanagerusersubscriptionsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_licensemanagerusersubscriptionsCmd.Name()}, args...))
		return p.Execute()
	}
	_licensemanagerusersubscriptionsCmd.SetArgs(args)
	return _licensemanagerusersubscriptionsCmd.Execute()
}
