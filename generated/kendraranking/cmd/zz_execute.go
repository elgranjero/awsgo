package cmd

func Execute(args []string) error {
	if p := _kendrarankingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kendrarankingCmd.Name()}, args...))
		return p.Execute()
	}
	_kendrarankingCmd.SetArgs(args)
	return _kendrarankingCmd.Execute()
}
