package cmd

func Execute(args []string) error {
	if p := _configserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_configserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_configserviceCmd.SetArgs(args)
	return _configserviceCmd.Execute()
}
