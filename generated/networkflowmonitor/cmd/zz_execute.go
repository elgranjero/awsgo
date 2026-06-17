package cmd

func Execute(args []string) error {
	if p := _networkflowmonitorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_networkflowmonitorCmd.Name()}, args...))
		return p.Execute()
	}
	_networkflowmonitorCmd.SetArgs(args)
	return _networkflowmonitorCmd.Execute()
}
