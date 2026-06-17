package cmd

func Execute(args []string) error {
	if p := _connectCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_connectCmd.Name()}, args...))
		return p.Execute()
	}
	_connectCmd.SetArgs(args)
	return _connectCmd.Execute()
}
