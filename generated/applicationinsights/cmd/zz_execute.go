package cmd

func Execute(args []string) error {
	if p := _applicationinsightsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_applicationinsightsCmd.Name()}, args...))
		return p.Execute()
	}
	_applicationinsightsCmd.SetArgs(args)
	return _applicationinsightsCmd.Execute()
}
