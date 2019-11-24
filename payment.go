package paynow

// Payment ...
type Payment struct {
	reference string
	email string
	items map[string]float64
}

// NewPayment ...
func NewPayment(reference string, email string) *Payment {
    p := new(Payment)
    p.reference = reference
    p.email = email
    return p
}

// Add ...
func (p *Payment) Add(item string, amount float64) {
	if (p.items == nil) {
		p.items = make(map[string]float64)
	}
	p.items[item] = amount
}