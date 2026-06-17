package cmd

func Execute(args []string) error {
	if p := _codecatalystCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codecatalystCmd.Name()}, args...))
		return p.Execute()
	}
	_codecatalystCmd.SetArgs(args)
	return _codecatalystCmd.Execute()
}
