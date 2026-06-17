package cmd

func Execute(args []string) error {
	if p := _chimesdkidentityCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chimesdkidentityCmd.Name()}, args...))
		return p.Execute()
	}
	_chimesdkidentityCmd.SetArgs(args)
	return _chimesdkidentityCmd.Execute()
}
