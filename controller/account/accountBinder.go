package account

type LoginBinder struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterBinder struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type MailBinder struct {
	Email string `json:"email"`
}
