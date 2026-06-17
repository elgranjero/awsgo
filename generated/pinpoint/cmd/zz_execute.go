package cmd

func Execute(args []string) error {
	if p := _pinpointCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pinpointCmd.Name()}, args...))
		return p.Execute()
	}
	_pinpointCmd.SetArgs(args)
	return _pinpointCmd.Execute()
}
