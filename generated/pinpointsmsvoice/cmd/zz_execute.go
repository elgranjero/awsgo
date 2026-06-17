package cmd

func Execute(args []string) error {
	if p := _pinpointsmsvoiceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pinpointsmsvoiceCmd.Name()}, args...))
		return p.Execute()
	}
	_pinpointsmsvoiceCmd.SetArgs(args)
	return _pinpointsmsvoiceCmd.Execute()
}
