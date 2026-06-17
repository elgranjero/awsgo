package cmd

func Execute(args []string) error {
	if p := _ioteventsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ioteventsCmd.Name()}, args...))
		return p.Execute()
	}
	_ioteventsCmd.SetArgs(args)
	return _ioteventsCmd.Execute()
}
