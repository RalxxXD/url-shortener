package shortener

type RedirectRespository interface {
	Find(code string) (*Redirect, error)
	Store(redirect *Redirect) error
}
