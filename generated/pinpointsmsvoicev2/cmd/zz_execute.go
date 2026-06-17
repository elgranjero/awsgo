package cmd

func Execute(args []string) error {
	if p := _pinpointsmsvoicev2Cmd.Parent(); p != nil {
		p.SetArgs(append([]string{_pinpointsmsvoicev2Cmd.Name()}, args...))
		return p.Execute()
	}
	_pinpointsmsvoicev2Cmd.SetArgs(args)
	return _pinpointsmsvoicev2Cmd.Execute()
}
