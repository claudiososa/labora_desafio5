package main

import (
	"fmt"
	"os"
)

func showTitle(title string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println(title)
	fmt.Println("--------------------------------------------------------------------------------------")
}

func showMainMenu() {
	fmt.Println("1. Create (user, book, author)")
	fmt.Println("2. List (users, books, authors, loans)")
	fmt.Println("3. New Loan")
	fmt.Println("4. Delete (user, book, author)")
	fmt.Println("0. Exit")
}

func showCreateMenu() {
	clearScreen()
	showTitle("Select action")
	fmt.Println("1. User")
	fmt.Println("2. Book")
	fmt.Println("3. Author")
	fmt.Println("0. Goback")
}

func clearScreen() {
	fmt.Print("\033[2J") // Secuencia de escape ANSI para borrar la pantalla
	fmt.Print("\033[H")  // Mover el cursor a la esquina superior izquierda
}

func selectCreateOption() {
	for {
		var option int

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		case 4:
			fmt.Println("4")
		case 0:
			return
		}
	}
}

func main() {
	for {
		clearScreen()
		showTitle("Welcome to library")
		showMainMenu()
		var option int

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			showCreateMenu()
			selectCreateOption()
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		case 4:
			fmt.Println("4")
		case 0:
			showTitle("Goodbye. Thanks for use owner system...")
			os.Exit(0)
		}
	}
}

/*




Una biblioteca necesita un sistema para gestionar su colección de libros.
El programa debe permitir la adición de nuevos libros, la búsqueda de
libros por título o autor, la actualización del estado de un libro
(disponible o prestado) y la eliminación de libros.


// Implementar una estructura de datos para almacenar la información de cada
// libro (título, autor, género y estado).
type Book struct {
	title  string
	author string
	genre  string
	status string
}

// Constructor
func NewBook(title string, author string, genre string, status string) *Book {
	return &Book{
		title:  title,
		author: author,
		genre:  genre,
		status: status,
	}

}

func (b *Book) setStatus(status string) {
	b.status = status
}

func (b Book) String() string {
	return fmt.Sprintf("%-30s %-20s %-20s %-10s", b.title, b.author, b.genre, b.status)
}
func addBook(books *map[string]Book, b *Book) {
	(*books)[b.title] = *b
}

func searchBook(books map[string]Book, text string) {
	booksFound := make(map[string]Book)

	for _, book := range books {
		//Busco por el titulo
		if strings.Contains(book.title, text) {
			booksFound[book.title] = book
		} else {
			//Busco por el autor
			if strings.Contains(books[book.title].author, text) {
				booksFound[book.title] = book
			}
		}
	}

	listBooks(booksFound, "Books found to '"+text+"'")
}

func updateStatus(newStatus string, titleBook string, books *map[string]Book) {
	if book, ok := (*books)[titleBook]; ok {
		book.setStatus(newStatus)
		(*books)[titleBook] = book
	}
}

func listBooks(books map[string]Book, title string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println(title)
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf("%-30s %-20s %-20s %-10s\n", "Title", "Author", "Genre", "Status")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, book := range books {
		fmt.Println(book)
	}
	fmt.Println()
}

func deleteBook(titleBook string, books *map[string]Book) {
	delete((*books), titleBook)
}

func main() {

	books := make(map[string]Book)

	book1 := NewBook("The Hobbit", "J.R.R. Tolkien", "Fantasy", "Available")
	book2 := NewBook("The Girl on the Train", "Paula Hawkins", "Mystery/Thriller", "Borrowed")
	book3 := NewBook("The Martian", "Andy Weir", "Science Fiction", "Available")
	book4 := NewBook("Gone Girl", "Gillian Flynn", "Mystery/Thriller", "Borrowed")
	book5 := NewBook("The Alchemist", "Paulo Coelho", "Fiction", "Available")

	// Permitir la adición de nuevos libros a la colección.
	addBook(&books, book1)
	addBook(&books, book2)
	addBook(&books, book3)
	addBook(&books, book4)
	addBook(&books, book5)

	listBooks(books, "List Books")

	// Permitir la búsqueda de libros por título o autor.
	textToSearch := "ir"

	searchBook(books, textToSearch)

	// Permitir la actualización del estado de un libro a "disponible" o "prestado".
	//bookToUpdate := "The Hobbit"

	newStatus := "borrowed"

	titleBook := "The Hobbit"

	updateStatus(newStatus, titleBook, &books)

	listBooks(books, "Update status")

	// Permitir la eliminación de libros de la colección.

	titleBookToDelete := "The Hobbit"

	deleteBook(titleBookToDelete, &books)

	listBooks(books, "Book deleted : "+titleBookToDelete)

}
*/
