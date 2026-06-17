package cmd

func Execute(args []string) error {
	if p := _b2biCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_b2biCmd.Name()}, args...))
		return p.Execute()
	}
	_b2biCmd.SetArgs(args)
	return _b2biCmd.Execute()
}
