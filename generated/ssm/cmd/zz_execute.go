package cmd

func Execute(args []string) error {
	if p := _ssmCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssmCmd.Name()}, args...))
		return p.Execute()
	}
	_ssmCmd.SetArgs(args)
	return _ssmCmd.Execute()
}
