package cmd

func Execute(args []string) error {
	if p := _licensemanagerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_licensemanagerCmd.Name()}, args...))
		return p.Execute()
	}
	_licensemanagerCmd.SetArgs(args)
	return _licensemanagerCmd.Execute()
}
