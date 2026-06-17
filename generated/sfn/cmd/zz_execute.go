package cmd

func Execute(args []string) error {
	if p := _sfnCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_sfnCmd.Name()}, args...))
		return p.Execute()
	}
	_sfnCmd.SetArgs(args)
	return _sfnCmd.Execute()
}
