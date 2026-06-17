package cmd

func Execute(args []string) error {
	if p := _snsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_snsCmd.Name()}, args...))
		return p.Execute()
	}
	_snsCmd.SetArgs(args)
	return _snsCmd.Execute()
}
