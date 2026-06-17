package cmd

func Execute(args []string) error {
	if p := _bcmrecommendedactionsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bcmrecommendedactionsCmd.Name()}, args...))
		return p.Execute()
	}
	_bcmrecommendedactionsCmd.SetArgs(args)
	return _bcmrecommendedactionsCmd.Execute()
}
