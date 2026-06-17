package cmd

func Execute(args []string) error {
	if p := _appintegrationsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appintegrationsCmd.Name()}, args...))
		return p.Execute()
	}
	_appintegrationsCmd.SetArgs(args)
	return _appintegrationsCmd.Execute()
}
