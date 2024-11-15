package service

type (
	IChaoxing interface {
		UserLogin(phone string, password string) (err error)
	}
)

var (
	localChaoxing IChaoxing
)

func Chaoxing() IChaoxing {
	if localChaoxing == nil {
		panic("implement not found for interface IChaoxing, forgot register?")
	}
	return localChaoxing
}

func RegisterChaoxing(i IChaoxing) {
	localChaoxing = i
}
