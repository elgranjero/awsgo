package cmd

func Execute(args []string) error {
	if p := _connectcontactlensCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_connectcontactlensCmd.Name()}, args...))
		return p.Execute()
	}
	_connectcontactlensCmd.SetArgs(args)
	return _connectcontactlensCmd.Execute()
}
