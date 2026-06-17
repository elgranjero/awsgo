package cmd

func Execute(args []string) error {
	if p := _connectparticipantCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_connectparticipantCmd.Name()}, args...))
		return p.Execute()
	}
	_connectparticipantCmd.SetArgs(args)
	return _connectparticipantCmd.Execute()
}
