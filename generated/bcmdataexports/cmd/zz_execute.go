package cmd

func Execute(args []string) error {
	if p := _bcmdataexportsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_bcmdataexportsCmd.Name()}, args...))
		return p.Execute()
	}
	_bcmdataexportsCmd.SetArgs(args)
	return _bcmdataexportsCmd.Execute()
}
