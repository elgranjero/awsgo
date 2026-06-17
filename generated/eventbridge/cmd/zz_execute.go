package cmd

func Execute(args []string) error {
	if p := _eventbridgeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_eventbridgeCmd.Name()}, args...))
		return p.Execute()
	}
	_eventbridgeCmd.SetArgs(args)
	return _eventbridgeCmd.Execute()
}
