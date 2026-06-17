package cmd

func Execute(args []string) error {
	if p := _controlcatalogCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_controlcatalogCmd.Name()}, args...))
		return p.Execute()
	}
	_controlcatalogCmd.SetArgs(args)
	return _controlcatalogCmd.Execute()
}
