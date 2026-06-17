package cmd

func Execute(args []string) error {
	if p := _marketplaceagreementCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplaceagreementCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplaceagreementCmd.SetArgs(args)
	return _marketplaceagreementCmd.Execute()
}
