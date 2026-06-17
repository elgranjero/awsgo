package cmd

func Execute(args []string) error {
	if p := _mailmanagerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mailmanagerCmd.Name()}, args...))
		return p.Execute()
	}
	_mailmanagerCmd.SetArgs(args)
	return _mailmanagerCmd.Execute()
}
