package registration

type RegistrationError int

const (
	TokenIssuedRecently RegistrationError = iota
	EmailAlreadyInUse
)

type UserRegistrationError struct {
	Kind RegistrationError
}

func NewUserRegistrationError(kind RegistrationError) *UserRegistrationError{
    return &UserRegistrationError{
        Kind: kind,
    }
}

func (u *UserRegistrationError)Error() string{
    switch u.Kind{
        case TokenIssuedRecently:
            return "Token issued recently, please wait"
        case EmailAlreadyInUse:
            return "Email already registered"
    }
    return "Unknown registration error"
} 

