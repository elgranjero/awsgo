package cmd

func Execute(args []string) error {
	if p := _braketCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_braketCmd.Name()}, args...))
		return p.Execute()
	}
	_braketCmd.SetArgs(args)
	return _braketCmd.Execute()
}
