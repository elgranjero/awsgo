package cmd

func Execute(args []string) error {
	if p := _mwaaserverlessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_mwaaserverlessCmd.Name()}, args...))
		return p.Execute()
	}
	_mwaaserverlessCmd.SetArgs(args)
	return _mwaaserverlessCmd.Execute()
}
