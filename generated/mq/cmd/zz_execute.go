package cmd

func Execute(args []string) error {
	if p := _mqCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mqCmd.Name()}, args...))
		return p.Execute()
	}
	_mqCmd.SetArgs(args)
	return _mqCmd.Execute()
}
