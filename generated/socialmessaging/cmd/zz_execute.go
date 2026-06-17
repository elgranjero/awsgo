package cmd

func Execute(args []string) error {
	if p := _socialmessagingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_socialmessagingCmd.Name()}, args...))
		return p.Execute()
	}
	_socialmessagingCmd.SetArgs(args)
	return _socialmessagingCmd.Execute()
}
