package service

import "database/sql"

type Book struct {
	ID     int
	Title  string
	Author string
	Genre  string
}

type BookService struct {
	db *sql.DB
}

// Inserir Book
func (s *BookService) CreateBook(book *Book) error {
	query := "Insert into book (title, author, genre) value(?,?,?)"
	result, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(lastInsertID)
	return nil
}

// Array
func (s *BookService) GetBooks() ([]Book, error) {
	query := "Select id, title, author, genre from books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	// Criando lista de books e alimentando com dados
	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// LISTAR BOOK POR ID
func (s *BookService) GetBookByID(id int) (*Book, error) {
	query := "select id, title, author, genre from books where id = ?"
	row := s.db.QueryRow(query, id)

	var book Book
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return nil, err
	}
	// Retornando ponteiro de book
	return &book, nil
}

// ATUALIZAR BOOK
func (s *BookService) UpdateBook(book *Book) error {
	query := "update books set title=?, author=?, genre=? where id=?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.Title)
	return err
}

// DELETAR BOOK
func (s *BookService) DeleteBook(id int) error {
	query := "delete from books where id=?"
	_, err := s.db.Exec(query, id)
	return err
}
