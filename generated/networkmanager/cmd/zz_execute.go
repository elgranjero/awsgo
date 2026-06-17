package cmd

func Execute(args []string) error {
	if p := _networkmanagerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_networkmanagerCmd.Name()}, args...))
		return p.Execute()
	}
	_networkmanagerCmd.SetArgs(args)
	return _networkmanagerCmd.Execute()
}
