package mailer

type MailAction int

const (
	SendValidationToken MailAction = iota
	RecoverPassword     MailAction = iota
)

type MailOptions struct {
	Email string
	Token string
	Host  string
	Port  string
}

type Mailer struct {
	Feed   chan chan MailOrder
	Config *MailOptions
}

type MailOrder struct {
	Destination string
	MailType    MailAction
}

func NewMailer(config *MailOptions) *Mailer {
	return &Mailer{
        Feed: make(chan chan MailOrder, 1024),
        Config: config,
    }
}

func (m *Mailer)Subscribe() chan MailOrder{
    newChan := make(chan MailOrder)
    m.Feed <- newChan
    return newChan
}

func (m *Mailer)Run() error{
    for {}

    return nil
}
