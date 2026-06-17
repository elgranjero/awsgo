package cmd

func Execute(args []string) error {
	if p := _iotCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotCmd.Name()}, args...))
		return p.Execute()
	}
	_iotCmd.SetArgs(args)
	return _iotCmd.Execute()
}
