package cmd

func Execute(args []string) error {
	if p := _kmsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kmsCmd.Name()}, args...))
		return p.Execute()
	}
	_kmsCmd.SetArgs(args)
	return _kmsCmd.Execute()
}
