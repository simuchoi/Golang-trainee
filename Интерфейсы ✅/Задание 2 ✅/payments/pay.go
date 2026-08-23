package payments

type PaymentMethod interface {
	Pay(usd float64) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentInfo   map[int]PaymentInfo
	paymentMethod PaymentMethod
}

// Конструктор
func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentInfo:   make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p PaymentModule) Pay(description string, usd float64) int {
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Amount:      usd,
		IsCancel:    false,
	}

	p.paymentInfo[id] = info

	return id
}

func (p *PaymentModule) Cancel(id int) {
	info, ok := p.paymentInfo[id]

	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.IsCancel = true

	p.paymentInfo[id] = info
}

func (p PaymentModule) GetInfo(id int) PaymentInfo {
	info, ok := p.paymentInfo[id]

	if !ok {
		return PaymentInfo{}
	}

	return info
}

func (p *PaymentModule) GetAllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentInfo))

	for k, v := range p.paymentInfo {
		tempMap[k] = v
	}

	return tempMap
}
