package cmd

func Execute(args []string) error {
	if p := _codeguruprofilerCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codeguruprofilerCmd.Name()}, args...))
		return p.Execute()
	}
	_codeguruprofilerCmd.SetArgs(args)
	return _codeguruprofilerCmd.Execute()
}
