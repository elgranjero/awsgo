package cmd

func Execute(args []string) error {
	if p := _glueCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_glueCmd.Name()}, args...))
		return p.Execute()
	}
	_glueCmd.SetArgs(args)
	return _glueCmd.Execute()
}
