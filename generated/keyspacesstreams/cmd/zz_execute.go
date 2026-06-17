package cmd

func Execute(args []string) error {
	if p := _keyspacesstreamsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_keyspacesstreamsCmd.Name()}, args...))
		return p.Execute()
	}
	_keyspacesstreamsCmd.SetArgs(args)
	return _keyspacesstreamsCmd.Execute()
}
