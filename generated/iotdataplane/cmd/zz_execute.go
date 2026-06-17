package cmd

func Execute(args []string) error {
	if p := _iotdataplaneCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotdataplaneCmd.Name()}, args...))
		return p.Execute()
	}
	_iotdataplaneCmd.SetArgs(args)
	return _iotdataplaneCmd.Execute()
}
