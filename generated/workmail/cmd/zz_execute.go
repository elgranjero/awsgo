package cmd

func Execute(args []string) error {
	if p := _workmailCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_workmailCmd.Name()}, args...))
		return p.Execute()
	}
	_workmailCmd.SetArgs(args)
	return _workmailCmd.Execute()
}
