package domain

type Transaction struct {
	Id             int64
	PaymentGateway string
	Amount         int64
	Status         string
}
