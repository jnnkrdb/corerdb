package mail

// mailobjects of this package

// mail class, contains information about the mail, receivers, etc.
type MailMessage struct {
	From       string   `json:"from"`
	To         []string `json:"to"`
	Cc         []string `json:"cc"`
	Bcc        []string `json:"bcc"`
	Subject    string   `json:"subject"`
	Body       string   `json:"body"`
	TextFormat string   `json:"textformat"`
	/*
		_Attachements struct {
			FileName string `json:"filename"`
			Content  string `json:"content"`
		} `json:"attachments"`
	*/
}

// smtp server class
type SMTPServer struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	// must be base64 encoded
	Password string `json:"password"`
}
