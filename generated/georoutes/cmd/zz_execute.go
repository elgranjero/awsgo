package cmd

func Execute(args []string) error {
	if p := _georoutesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_georoutesCmd.Name()}, args...))
		return p.Execute()
	}
	_georoutesCmd.SetArgs(args)
	return _georoutesCmd.Execute()
}
