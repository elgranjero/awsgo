package cmd

func Execute(args []string) error {
	if p := _cloudsearchdomainCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudsearchdomainCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudsearchdomainCmd.SetArgs(args)
	return _cloudsearchdomainCmd.Execute()
}
