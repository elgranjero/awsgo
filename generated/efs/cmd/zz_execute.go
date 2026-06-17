package cmd

func Execute(args []string) error {
	if p := _efsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_efsCmd.Name()}, args...))
		return p.Execute()
	}
	_efsCmd.SetArgs(args)
	return _efsCmd.Execute()
}
