package cmd

func Execute(args []string) error {
	if p := _gameliftCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_gameliftCmd.Name()}, args...))
		return p.Execute()
	}
	_gameliftCmd.SetArgs(args)
	return _gameliftCmd.Execute()
}
