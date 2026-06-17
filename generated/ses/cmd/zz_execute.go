package cmd

func Execute(args []string) error {
	if p := _sesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sesCmd.Name()}, args...))
		return p.Execute()
	}
	_sesCmd.SetArgs(args)
	return _sesCmd.Execute()
}
