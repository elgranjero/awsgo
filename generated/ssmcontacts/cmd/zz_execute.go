package cmd

func Execute(args []string) error {
	if p := _ssmcontactsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ssmcontactsCmd.Name()}, args...))
		return p.Execute()
	}
	_ssmcontactsCmd.SetArgs(args)
	return _ssmcontactsCmd.Execute()
}
