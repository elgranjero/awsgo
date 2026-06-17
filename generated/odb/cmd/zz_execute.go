package cmd

func Execute(args []string) error {
	if p := _odbCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_odbCmd.Name()}, args...))
		return p.Execute()
	}
	_odbCmd.SetArgs(args)
	return _odbCmd.Execute()
}
