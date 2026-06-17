package cmd

func Execute(args []string) error {
	if p := _tnbCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_tnbCmd.Name()}, args...))
		return p.Execute()
	}
	_tnbCmd.SetArgs(args)
	return _tnbCmd.Execute()
}
