package cmd

func Execute(args []string) error {
	if p := _iotwirelessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotwirelessCmd.Name()}, args...))
		return p.Execute()
	}
	_iotwirelessCmd.SetArgs(args)
	return _iotwirelessCmd.Execute()
}
