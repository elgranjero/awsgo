package cmd

func Execute(args []string) error {
	if p := _personalizeruntimeCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_personalizeruntimeCmd.Name()}, args...))
		return p.Execute()
	}
	_personalizeruntimeCmd.SetArgs(args)
	return _personalizeruntimeCmd.Execute()
}
