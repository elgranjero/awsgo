package cmd

func Execute(args []string) error {
	if p := _chimesdkmediapipelinesCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chimesdkmediapipelinesCmd.Name()}, args...))
		return p.Execute()
	}
	_chimesdkmediapipelinesCmd.SetArgs(args)
	return _chimesdkmediapipelinesCmd.Execute()
}
