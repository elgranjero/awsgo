package cmd

func Execute(args []string) error {
	if p := _chimesdkmeetingsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chimesdkmeetingsCmd.Name()}, args...))
		return p.Execute()
	}
	_chimesdkmeetingsCmd.SetArgs(args)
	return _chimesdkmeetingsCmd.Execute()
}
