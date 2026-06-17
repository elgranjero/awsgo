package cmd

func Execute(args []string) error {
	if p := _personalizeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_personalizeCmd.Name()}, args...))
		return p.Execute()
	}
	_personalizeCmd.SetArgs(args)
	return _personalizeCmd.Execute()
}
