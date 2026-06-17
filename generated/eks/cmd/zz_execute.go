package cmd

func Execute(args []string) error {
	if p := _eksCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_eksCmd.Name()}, args...))
		return p.Execute()
	}
	_eksCmd.SetArgs(args)
	return _eksCmd.Execute()
}
