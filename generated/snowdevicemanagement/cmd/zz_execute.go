package cmd

func Execute(args []string) error {
	if p := _snowdevicemanagementCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_snowdevicemanagementCmd.Name()}, args...))
		return p.Execute()
	}
	_snowdevicemanagementCmd.SetArgs(args)
	return _snowdevicemanagementCmd.Execute()
}
