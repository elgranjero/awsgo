package cmd

func Execute(args []string) error {
	if p := _iotthingsgraphCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotthingsgraphCmd.Name()}, args...))
		return p.Execute()
	}
	_iotthingsgraphCmd.SetArgs(args)
	return _iotthingsgraphCmd.Execute()
}
