package postgre

type BookRepository struct{}

func (r *BookRepository) Get() {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) Create() {
	//TODO implement me
	panic("implement me")
}

func (r *BookRepository) Delete() {
	r.Get()
	//TODO implement me
	panic("implement me")
}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}
