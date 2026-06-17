package cmd

func Execute(args []string) error {
	if p := _appstreamCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_appstreamCmd.Name()}, args...))
		return p.Execute()
	}
	_appstreamCmd.SetArgs(args)
	return _appstreamCmd.Execute()
}
