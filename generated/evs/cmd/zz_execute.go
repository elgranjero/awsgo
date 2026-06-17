package cmd

func Execute(args []string) error {
	if p := _evsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_evsCmd.Name()}, args...))
		return p.Execute()
	}
	_evsCmd.SetArgs(args)
	return _evsCmd.Execute()
}
