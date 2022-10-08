package simple

//简单工厂

type EmailFactory interface {
	send() string
}

func newClient(provider string) EmailFactory {
	switch provider {
	case "qq":
		return &QQEamil{}
	case "google":
		return &Gmail{}
	default:
		return nil
	}
}

type QQEamil struct{}

func (qq *QQEamil) send() string {
	return "qq邮箱"
}

type Gmail struct{}

func (g *Gmail) send() string {
	return "google 邮箱"
}
