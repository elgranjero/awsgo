package cmd

func Execute(args []string) error {
	if p := _workspacesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workspacesCmd.Name()}, args...))
		return p.Execute()
	}
	_workspacesCmd.SetArgs(args)
	return _workspacesCmd.Execute()
}
