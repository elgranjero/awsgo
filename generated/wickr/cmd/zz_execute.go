package cmd

func Execute(args []string) error {
	if p := _wickrCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_wickrCmd.Name()}, args...))
		return p.Execute()
	}
	_wickrCmd.SetArgs(args)
	return _wickrCmd.Execute()
}
