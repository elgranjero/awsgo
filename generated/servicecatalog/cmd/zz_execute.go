package cmd

func Execute(args []string) error {
	if p := _servicecatalogCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_servicecatalogCmd.Name()}, args...))
		return p.Execute()
	}
	_servicecatalogCmd.SetArgs(args)
	return _servicecatalogCmd.Execute()
}
