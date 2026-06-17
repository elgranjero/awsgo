package cmd

func Execute(args []string) error {
	if p := _codeartifactCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codeartifactCmd.Name()}, args...))
		return p.Execute()
	}
	_codeartifactCmd.SetArgs(args)
	return _codeartifactCmd.Execute()
}
