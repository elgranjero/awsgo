package cmd

func Execute(args []string) error {
	if p := _devicefarmCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_devicefarmCmd.Name()}, args...))
		return p.Execute()
	}
	_devicefarmCmd.SetArgs(args)
	return _devicefarmCmd.Execute()
}
