package cmd

func Execute(args []string) error {
	if p := _pinpointemailCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pinpointemailCmd.Name()}, args...))
		return p.Execute()
	}
	_pinpointemailCmd.SetArgs(args)
	return _pinpointemailCmd.Execute()
}
