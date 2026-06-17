package cmd

func Execute(args []string) error {
	if p := _ivsrealtimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ivsrealtimeCmd.Name()}, args...))
		return p.Execute()
	}
	_ivsrealtimeCmd.SetArgs(args)
	return _ivsrealtimeCmd.Execute()
}
