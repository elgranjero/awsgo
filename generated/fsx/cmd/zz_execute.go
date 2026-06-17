package cmd

func Execute(args []string) error {
	if p := _fsxCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_fsxCmd.Name()}, args...))
		return p.Execute()
	}
	_fsxCmd.SetArgs(args)
	return _fsxCmd.Execute()
}
