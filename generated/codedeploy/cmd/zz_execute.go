package cmd

func Execute(args []string) error {
	if p := _codedeployCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codedeployCmd.Name()}, args...))
		return p.Execute()
	}
	_codedeployCmd.SetArgs(args)
	return _codedeployCmd.Execute()
}
