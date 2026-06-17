package cmd

func Execute(args []string) error {
	if p := _arcregionswitchCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_arcregionswitchCmd.Name()}, args...))
		return p.Execute()
	}
	_arcregionswitchCmd.SetArgs(args)
	return _arcregionswitchCmd.Execute()
}
