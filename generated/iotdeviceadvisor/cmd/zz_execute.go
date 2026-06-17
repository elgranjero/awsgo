package cmd

func Execute(args []string) error {
	if p := _iotdeviceadvisorCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotdeviceadvisorCmd.Name()}, args...))
		return p.Execute()
	}
	_iotdeviceadvisorCmd.SetArgs(args)
	return _iotdeviceadvisorCmd.Execute()
}
