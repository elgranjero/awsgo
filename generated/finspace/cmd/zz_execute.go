package cmd

func Execute(args []string) error {
	if p := _finspaceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_finspaceCmd.Name()}, args...))
		return p.Execute()
	}
	_finspaceCmd.SetArgs(args)
	return _finspaceCmd.Execute()
}
