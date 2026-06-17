package cmd

func Execute(args []string) error {
	if p := _notificationsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_notificationsCmd.Name()}, args...))
		return p.Execute()
	}
	_notificationsCmd.SetArgs(args)
	return _notificationsCmd.Execute()
}
