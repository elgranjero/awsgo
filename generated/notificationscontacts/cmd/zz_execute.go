package cmd

func Execute(args []string) error {
	if p := _notificationscontactsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_notificationscontactsCmd.Name()}, args...))
		return p.Execute()
	}
	_notificationscontactsCmd.SetArgs(args)
	return _notificationscontactsCmd.Execute()
}
