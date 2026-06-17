package cmd

func Execute(args []string) error {
	if p := _chimesdkmessagingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chimesdkmessagingCmd.Name()}, args...))
		return p.Execute()
	}
	_chimesdkmessagingCmd.SetArgs(args)
	return _chimesdkmessagingCmd.Execute()
}
