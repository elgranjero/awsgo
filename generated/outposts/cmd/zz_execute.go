package cmd

func Execute(args []string) error {
	if p := _outpostsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_outpostsCmd.Name()}, args...))
		return p.Execute()
	}
	_outpostsCmd.SetArgs(args)
	return _outpostsCmd.Execute()
}
