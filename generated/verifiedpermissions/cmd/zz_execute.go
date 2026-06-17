package cmd

func Execute(args []string) error {
	if p := _verifiedpermissionsCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_verifiedpermissionsCmd.Name()}, args...))
		return p.Execute()
	}
	_verifiedpermissionsCmd.SetArgs(args)
	return _verifiedpermissionsCmd.Execute()
}
