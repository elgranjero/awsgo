package cmd

func Execute(args []string) error {
	if p := _mgnCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mgnCmd.Name()}, args...))
		return p.Execute()
	}
	_mgnCmd.SetArgs(args)
	return _mgnCmd.Execute()
}
