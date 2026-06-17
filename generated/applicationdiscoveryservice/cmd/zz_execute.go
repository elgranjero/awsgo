package cmd

func Execute(args []string) error {
	if p := _applicationdiscoveryserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_applicationdiscoveryserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_applicationdiscoveryserviceCmd.SetArgs(args)
	return _applicationdiscoveryserviceCmd.Execute()
}
