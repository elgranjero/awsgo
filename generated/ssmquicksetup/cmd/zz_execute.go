package cmd

func Execute(args []string) error {
	if p := _ssmquicksetupCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssmquicksetupCmd.Name()}, args...))
		return p.Execute()
	}
	_ssmquicksetupCmd.SetArgs(args)
	return _ssmquicksetupCmd.Execute()
}
