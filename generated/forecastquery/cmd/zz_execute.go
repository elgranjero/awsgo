package cmd

func Execute(args []string) error {
	if p := _forecastqueryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_forecastqueryCmd.Name()}, args...))
		return p.Execute()
	}
	_forecastqueryCmd.SetArgs(args)
	return _forecastqueryCmd.Execute()
}
