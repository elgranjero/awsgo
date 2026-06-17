package cmd

func Execute(args []string) error {
	if p := _ebsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ebsCmd.Name()}, args...))
		return p.Execute()
	}
	_ebsCmd.SetArgs(args)
	return _ebsCmd.Execute()
}
