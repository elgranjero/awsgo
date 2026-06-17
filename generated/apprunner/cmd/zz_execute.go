package cmd

func Execute(args []string) error {
	if p := _apprunnerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_apprunnerCmd.Name()}, args...))
		return p.Execute()
	}
	_apprunnerCmd.SetArgs(args)
	return _apprunnerCmd.Execute()
}
