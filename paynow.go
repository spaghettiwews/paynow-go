package paynow

type Paynow struct {
	id, key, resultUrl, returnUrl string
	Payment Payment
}

func NewPaynow(id string, key string, resultUrl string, returnUrl string) *Paynow {
    p := new(Paynow)
    p.id = id
    p.key = key
    p.resultUrl = resultUrl
    p.returnUrl = returnUrl
    return p
}