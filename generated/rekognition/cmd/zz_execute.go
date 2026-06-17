package cmd

func Execute(args []string) error {
	if p := _rekognitionCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_rekognitionCmd.Name()}, args...))
		return p.Execute()
	}
	_rekognitionCmd.SetArgs(args)
	return _rekognitionCmd.Execute()
}
