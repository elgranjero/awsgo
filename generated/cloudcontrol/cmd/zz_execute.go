package cmd

func Execute(args []string) error {
	if p := _cloudcontrolCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudcontrolCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudcontrolCmd.SetArgs(args)
	return _cloudcontrolCmd.Execute()
}
