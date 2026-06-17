package cmd

func Execute(args []string) error {
	if p := _xrayCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_xrayCmd.Name()}, args...))
		return p.Execute()
	}
	_xrayCmd.SetArgs(args)
	return _xrayCmd.Execute()
}
