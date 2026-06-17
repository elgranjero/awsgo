package cmd

func Execute(args []string) error {
	if p := _supplychainCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_supplychainCmd.Name()}, args...))
		return p.Execute()
	}
	_supplychainCmd.SetArgs(args)
	return _supplychainCmd.Execute()
}
