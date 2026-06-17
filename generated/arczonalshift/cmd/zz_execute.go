package cmd

func Execute(args []string) error {
	if p := _arczonalshiftCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_arczonalshiftCmd.Name()}, args...))
		return p.Execute()
	}
	_arczonalshiftCmd.SetArgs(args)
	return _arczonalshiftCmd.Execute()
}
