package cmd

func Execute(args []string) error {
	if p := _machinelearningCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_machinelearningCmd.Name()}, args...))
		return p.Execute()
	}
	_machinelearningCmd.SetArgs(args)
	return _machinelearningCmd.Execute()
}
