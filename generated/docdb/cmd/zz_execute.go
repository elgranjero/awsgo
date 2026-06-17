package cmd

func Execute(args []string) error {
	if p := _docdbCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_docdbCmd.Name()}, args...))
		return p.Execute()
	}
	_docdbCmd.SetArgs(args)
	return _docdbCmd.Execute()
}
