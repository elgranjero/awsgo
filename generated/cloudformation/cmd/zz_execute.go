package cmd

func Execute(args []string) error {
	if p := _cloudformationCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_cloudformationCmd.Name()}, args...))
		return p.Execute()
	}
	_cloudformationCmd.SetArgs(args)
	return _cloudformationCmd.Execute()
}
