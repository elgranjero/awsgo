package cmd

func Execute(args []string) error {
	if p := _marketplacecatalogCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplacecatalogCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplacecatalogCmd.SetArgs(args)
	return _marketplacecatalogCmd.Execute()
}
