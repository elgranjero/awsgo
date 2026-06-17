package cmd

func Execute(args []string) error {
	if p := _connectcampaignsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_connectcampaignsCmd.Name()}, args...))
		return p.Execute()
	}
	_connectcampaignsCmd.SetArgs(args)
	return _connectcampaignsCmd.Execute()
}
