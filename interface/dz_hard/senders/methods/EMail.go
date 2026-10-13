package methods

type EMail struct {
	from string
	text string
}

func NewEMail(from string, text string) *EMail {
	return &EMail{
		from: from,
		text: text,
	}
}

func (e *EMail) Send() {
	
}
