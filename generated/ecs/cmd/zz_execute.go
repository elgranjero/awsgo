package cmd

func Execute(args []string) error {
	if p := _ecsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ecsCmd.Name()}, args...))
		return p.Execute()
	}
	_ecsCmd.SetArgs(args)
	return _ecsCmd.Execute()
}
