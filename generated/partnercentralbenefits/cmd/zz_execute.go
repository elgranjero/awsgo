package cmd

func Execute(args []string) error {
	if p := _partnercentralbenefitsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_partnercentralbenefitsCmd.Name()}, args...))
		return p.Execute()
	}
	_partnercentralbenefitsCmd.SetArgs(args)
	return _partnercentralbenefitsCmd.Execute()
}
