package cmd

func Execute(args []string) error {
	if p := _acmCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_acmCmd.Name()}, args...))
		return p.Execute()
	}
	_acmCmd.SetArgs(args)
	return _acmCmd.Execute()
}
