package cmd

func Execute(args []string) error {
	if p := _kinesisvideosignalingCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_kinesisvideosignalingCmd.Name()}, args...))
		return p.Execute()
	}
	_kinesisvideosignalingCmd.SetArgs(args)
	return _kinesisvideosignalingCmd.Execute()
}
