package services

type TttServiceType struct{}

var TttService TttServiceType

func init() {
	TttService = TttServiceType{}
}

func (*TttServiceType) DoSmth() string {
	return "Hello World!"
}
