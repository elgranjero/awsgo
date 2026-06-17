package cmd

func Execute(args []string) error {
	if p := _ecrpublicCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_ecrpublicCmd.Name()}, args...))
		return p.Execute()
	}
	_ecrpublicCmd.SetArgs(args)
	return _ecrpublicCmd.Execute()
}
