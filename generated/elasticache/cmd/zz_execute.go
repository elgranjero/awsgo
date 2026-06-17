package cmd

func Execute(args []string) error {
	if p := _elasticacheCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_elasticacheCmd.Name()}, args...))
		return p.Execute()
	}
	_elasticacheCmd.SetArgs(args)
	return _elasticacheCmd.Execute()
}
