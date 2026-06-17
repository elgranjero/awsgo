package cmd

func Execute(args []string) error {
	if p := _taxsettingsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_taxsettingsCmd.Name()}, args...))
		return p.Execute()
	}
	_taxsettingsCmd.SetArgs(args)
	return _taxsettingsCmd.Execute()
}
