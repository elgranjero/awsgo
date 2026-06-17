package cmd

func Execute(args []string) error {
	if p := _codestarconnectionsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codestarconnectionsCmd.Name()}, args...))
		return p.Execute()
	}
	_codestarconnectionsCmd.SetArgs(args)
	return _codestarconnectionsCmd.Execute()
}
