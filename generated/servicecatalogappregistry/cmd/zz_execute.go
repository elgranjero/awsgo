package cmd

func Execute(args []string) error {
	if p := _servicecatalogappregistryCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_servicecatalogappregistryCmd.Name()}, args...))
		return p.Execute()
	}
	_servicecatalogappregistryCmd.SetArgs(args)
	return _servicecatalogappregistryCmd.Execute()
}
