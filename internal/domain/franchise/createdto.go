package franchise

type CreateDTO struct {
	ID  string `validate:"required,uuid"`
	URL string `validate:"required,url"`
}
