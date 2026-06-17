package cmd

func Execute(args []string) error {
	if p := _networkfirewallCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_networkfirewallCmd.Name()}, args...))
		return p.Execute()
	}
	_networkfirewallCmd.SetArgs(args)
	return _networkfirewallCmd.Execute()
}
