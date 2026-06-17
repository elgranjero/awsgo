package cmd

func Execute(args []string) error {
	if p := _codegurusecurityCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_codegurusecurityCmd.Name()}, args...))
		return p.Execute()
	}
	_codegurusecurityCmd.SetArgs(args)
	return _codegurusecurityCmd.Execute()
}
