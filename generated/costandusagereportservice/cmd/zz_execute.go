package cmd

func Execute(args []string) error {
	if p := _costandusagereportserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_costandusagereportserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_costandusagereportserviceCmd.SetArgs(args)
	return _costandusagereportserviceCmd.Execute()
}
