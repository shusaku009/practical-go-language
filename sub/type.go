package sub

import "fmt"

type (
	HTTPStatus    int
	NationalRoute int
)

const (
	StatusOK              HTTPStatus = 200
	StatusUnauthorized    HTTPStatus = 401
	StatusPaymentRequired HTTPStatus = 402
	StatusForbidden       HTTPStatus = 403
)

func (s HTTPStatus) String() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusUnauthorized:
		return "Unauthorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	default:
		return fmt.Sprintf("HTTPStatus(%d)", s)
	}
}

const (
	Nagasakikaido   NationalRoute = 200
	AizuNumatakaido NationalRoute = 401
	HokurikuDo      NationalRoute = 402
	KurinokiBypass  NationalRoute = 403
)

func (n NationalRoute) String() string {
	switch n {
	case Nagasakikaido:
		return "長崎街道"
	case AizuNumatakaido:
		return "会津沼田街道"
	case HokurikuDo:
		return "北陸道"
	case "KurinokiBypass":
		return "栗の木バイパス"
	default:
		return fmt.Sprintf("国道%d号線", n)
	}
}

type Consumer struct {
	ActiveFlg bool
}

type Consumers []Consumer

func (c Consumers) ActiveConsumer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.ActiveFlg {
			resp = append(resp, v)
		}
	}
	return resp
}
