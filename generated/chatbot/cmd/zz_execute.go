package cmd

func Execute(args []string) error {
	if p := _chatbotCmd.Parent(); p != nil {
		p.SetArgs(append([]string{_chatbotCmd.Name()}, args...))
		return p.Execute()
	}
	_chatbotCmd.SetArgs(args)
	return _chatbotCmd.Execute()
}
