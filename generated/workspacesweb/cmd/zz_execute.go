package cmd

func Execute(args []string) error {
	if p := _workspaceswebCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workspaceswebCmd.Name()}, args...))
		return p.Execute()
	}
	_workspaceswebCmd.SetArgs(args)
	return _workspaceswebCmd.Execute()
}
