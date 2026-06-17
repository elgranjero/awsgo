package cmd

func Execute(args []string) error {
	if p := _stsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_stsCmd.Name()}, args...))
		return p.Execute()
	}
	_stsCmd.SetArgs(args)
	return _stsCmd.Execute()
}
