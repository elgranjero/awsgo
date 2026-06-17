package cmd

func Execute(args []string) error {
	if p := _bedrockagentruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bedrockagentruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_bedrockagentruntimeCmd.SetArgs(args)
	return _bedrockagentruntimeCmd.Execute()
}
