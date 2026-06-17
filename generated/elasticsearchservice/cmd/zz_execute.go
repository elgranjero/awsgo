package cmd

func Execute(args []string) error {
	if p := _elasticsearchserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_elasticsearchserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_elasticsearchserviceCmd.SetArgs(args)
	return _elasticsearchserviceCmd.Execute()
}
