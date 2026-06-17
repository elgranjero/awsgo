package cmd

func Execute(args []string) error {
	if p := _shieldCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_shieldCmd.Name()}, args...))
		return p.Execute()
	}
	_shieldCmd.SetArgs(args)
	return _shieldCmd.Execute()
}
