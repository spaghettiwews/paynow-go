package paynow-go

type Payment struct {
	reference string
	email string
	items map[string]float64
}

func NewPayment(reference string, email string) *Payment {
    p := new(Payment)
    p.reference = reference
    p.email = email
    return p
}

func (p *Payment) Add(item string, amount float64) {
	if (p.items == nil) {
		p.items = make(map[string]float64)
	}
	p.items[item] = amount
}