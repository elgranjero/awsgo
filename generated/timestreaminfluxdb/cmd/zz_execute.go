package cmd

func Execute(args []string) error {
	if p := _timestreaminfluxdbCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_timestreaminfluxdbCmd.Name()}, args...))
		return p.Execute()
	}
	_timestreaminfluxdbCmd.SetArgs(args)
	return _timestreaminfluxdbCmd.Execute()
}
