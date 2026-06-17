package cmd

func Execute(args []string) error {
	if p := _freetierCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_freetierCmd.Name()}, args...))
		return p.Execute()
	}
	_freetierCmd.SetArgs(args)
	return _freetierCmd.Execute()
}
