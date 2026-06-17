package cmd

func Execute(args []string) error {
	if p := _workspacesinstancesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workspacesinstancesCmd.Name()}, args...))
		return p.Execute()
	}
	_workspacesinstancesCmd.SetArgs(args)
	return _workspacesinstancesCmd.Execute()
}
