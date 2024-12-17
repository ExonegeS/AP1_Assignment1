package service

func (s *service) ListBooks() ListBooksResponse {
	books := s.books.ListBooks()
	return ListBooksResponse{Books: books}
}
