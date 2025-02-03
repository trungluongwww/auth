package social

type Social interface {
	GetFacebookInfo(accessToken string) (res *FacebookInfo, err error)
}

type socialImpl struct {
}

func NewSocial() Social {
	return &socialImpl{}
}
