package cmd

func Execute(args []string) error {
	if p := _signerdataCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_signerdataCmd.Name()}, args...))
		return p.Execute()
	}
	_signerdataCmd.SetArgs(args)
	return _signerdataCmd.Execute()
}
