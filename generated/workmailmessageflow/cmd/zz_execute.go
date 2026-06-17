package cmd

func Execute(args []string) error {
	if p := _workmailmessageflowCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workmailmessageflowCmd.Name()}, args...))
		return p.Execute()
	}
	_workmailmessageflowCmd.SetArgs(args)
	return _workmailmessageflowCmd.Execute()
}
