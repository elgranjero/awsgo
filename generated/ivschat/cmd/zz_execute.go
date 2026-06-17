package cmd

func Execute(args []string) error {
	if p := _ivschatCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ivschatCmd.Name()}, args...))
		return p.Execute()
	}
	_ivschatCmd.SetArgs(args)
	return _ivschatCmd.Execute()
}
