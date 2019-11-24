package paynow

// Paynow ...
type Paynow struct {
	id, key, resultURL, returnURL string
	Payment Payment
}
// NewPaynow ...
func NewPaynow(id string, key string, resultURL string, returnURL string) *Paynow {
    p := new(Paynow)
    p.id = id
    p.key = key
    p.resultURL = resultURL
    p.returnURL = returnURL
    return p
}