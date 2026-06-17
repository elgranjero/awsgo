package cmd

func Execute(args []string) error {
	if p := _syntheticsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_syntheticsCmd.Name()}, args...))
		return p.Execute()
	}
	_syntheticsCmd.SetArgs(args)
	return _syntheticsCmd.Execute()
}
