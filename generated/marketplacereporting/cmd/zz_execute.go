package cmd

func Execute(args []string) error {
	if p := _marketplacereportingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_marketplacereportingCmd.Name()}, args...))
		return p.Execute()
	}
	_marketplacereportingCmd.SetArgs(args)
	return _marketplacereportingCmd.Execute()
}
