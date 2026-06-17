package cmd

func Execute(args []string) error {
	if p := _iamCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iamCmd.Name()}, args...))
		return p.Execute()
	}
	_iamCmd.SetArgs(args)
	return _iamCmd.Execute()
}
