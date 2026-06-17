package cmd

func Execute(args []string) error {
	if p := _lexruntimeserviceCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_lexruntimeserviceCmd.Name()}, args...))
		return p.Execute()
	}
	_lexruntimeserviceCmd.SetArgs(args)
	return _lexruntimeserviceCmd.Execute()
}
