package cmd

func Execute(args []string) error {
	if p := _iotjobsdataplaneCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotjobsdataplaneCmd.Name()}, args...))
		return p.Execute()
	}
	_iotjobsdataplaneCmd.SetArgs(args)
	return _iotjobsdataplaneCmd.Execute()
}
