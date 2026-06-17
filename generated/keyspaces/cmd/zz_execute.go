package cmd

func Execute(args []string) error {
	if p := _keyspacesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_keyspacesCmd.Name()}, args...))
		return p.Execute()
	}
	_keyspacesCmd.SetArgs(args)
	return _keyspacesCmd.Execute()
}
