package cmd

func Execute(args []string) error {
	if p := _emrserverlessCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_emrserverlessCmd.Name()}, args...))
		return p.Execute()
	}
	_emrserverlessCmd.SetArgs(args)
	return _emrserverlessCmd.Execute()
}
