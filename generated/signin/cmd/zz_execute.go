package cmd

func Execute(args []string) error {
	if p := _signinCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_signinCmd.Name()}, args...))
		return p.Execute()
	}
	_signinCmd.SetArgs(args)
	return _signinCmd.Execute()
}
