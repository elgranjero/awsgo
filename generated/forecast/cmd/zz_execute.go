package cmd

func Execute(args []string) error {
	if p := _forecastCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_forecastCmd.Name()}, args...))
		return p.Execute()
	}
	_forecastCmd.SetArgs(args)
	return _forecastCmd.Execute()
}
