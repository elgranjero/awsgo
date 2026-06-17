package cmd

func Execute(args []string) error {
	if p := _docdbelasticCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_docdbelasticCmd.Name()}, args...))
		return p.Execute()
	}
	_docdbelasticCmd.SetArgs(args)
	return _docdbelasticCmd.Execute()
}
