package cmd

func Execute(args []string) error {
	if p := _identitystoreCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_identitystoreCmd.Name()}, args...))
		return p.Execute()
	}
	_identitystoreCmd.SetArgs(args)
	return _identitystoreCmd.Execute()
}
