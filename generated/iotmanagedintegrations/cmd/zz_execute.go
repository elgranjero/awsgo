package cmd

func Execute(args []string) error {
	if p := _iotmanagedintegrationsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_iotmanagedintegrationsCmd.Name()}, args...))
		return p.Execute()
	}
	_iotmanagedintegrationsCmd.SetArgs(args)
	return _iotmanagedintegrationsCmd.Execute()
}
