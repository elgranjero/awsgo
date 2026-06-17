package cmd

func Execute(args []string) error {
	if p := _redshiftserverlessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_redshiftserverlessCmd.Name()}, args...))
		return p.Execute()
	}
	_redshiftserverlessCmd.SetArgs(args)
	return _redshiftserverlessCmd.Execute()
}
