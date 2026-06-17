package cmd

func Execute(args []string) error {
	if p := _gameliftstreamsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_gameliftstreamsCmd.Name()}, args...))
		return p.Execute()
	}
	_gameliftstreamsCmd.SetArgs(args)
	return _gameliftstreamsCmd.Execute()
}
