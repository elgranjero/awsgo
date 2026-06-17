package cmd

func Execute(args []string) error {
	if p := _licensemanagerlinuxsubscriptionsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_licensemanagerlinuxsubscriptionsCmd.Name()}, args...))
		return p.Execute()
	}
	_licensemanagerlinuxsubscriptionsCmd.SetArgs(args)
	return _licensemanagerlinuxsubscriptionsCmd.Execute()
}
