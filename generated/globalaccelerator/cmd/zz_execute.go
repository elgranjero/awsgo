package cmd

func Execute(args []string) error {
	if p := _globalacceleratorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_globalacceleratorCmd.Name()}, args...))
		return p.Execute()
	}
	_globalacceleratorCmd.SetArgs(args)
	return _globalacceleratorCmd.Execute()
}
