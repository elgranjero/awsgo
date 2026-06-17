package cmd

func Execute(args []string) error {
	if p := _vpclatticeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_vpclatticeCmd.Name()}, args...))
		return p.Execute()
	}
	_vpclatticeCmd.SetArgs(args)
	return _vpclatticeCmd.Execute()
}
