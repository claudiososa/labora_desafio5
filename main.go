package main

import (
	"fmt"
	"log"
	"os"

	"github.com/claudiososa/labora_desafio5/models"
	"github.com/claudiososa/labora_desafio5/services"
	"github.com/claudiososa/labora_desafio5/utils"
)

func showTitle(title string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println(title)
	fmt.Println("--------------------------------------------------------------------------------------")
}

func showMainMenu() {
	fmt.Println("1. Create (user, book, author)")
	fmt.Println("2. List (users, books, authors, loans)")
	fmt.Println("3. Search")
	fmt.Println("4. New Loan")
	fmt.Println("5. Delete (user, book, author)")
	fmt.Println("0. Exit")
}

func showCreateMenu() {
	clearScreen()
	showTitle("Select action (Create)")
	fmt.Println("1. User")
	fmt.Println("2. Book")
	fmt.Println("3. Author")
	fmt.Println("0. Goback")
}

func showSearchMenu() {
	clearScreen()
	showTitle("Select action (Search)")
	fmt.Println("1. User")
	fmt.Println("2. Book")
	fmt.Println("3. Author")
	fmt.Println("0. Goback")
}

func showListMenu() {
	clearScreen()
	showTitle("Select action (List)")
	fmt.Println("1. Users")
	fmt.Println("2. Books (available)")
	fmt.Println("3. Authors")
	fmt.Println("0. Goback")
}

func clearScreen() {
	fmt.Print("\033[2J") // Secuencia de escape ANSI para borrar la pantalla
	fmt.Print("\033[H")  // Mover el cursor a la esquina superior izquierda
}

func listBooks(books []models.Book, title string) {
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println(title)
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Printf("%-6s %-40s %-20s %-10s\n", "Id", "Title", "Status", "Author")
	fmt.Println("--------------------------------------------------------------------------------------")

	for _, book := range books {
		fmt.Printf("%-6d %-40s %-20s %-10s\n", book.Id, book.Title, book.Status, book.Author.Name)
	}
	fmt.Println()
}

func selectCreateOption() {
	for {
		var option int

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fmt.Println("1")
		case 2:
			//List Books
			clearScreen()
			serviceBook := &services.BookService{}

			books, err := serviceBook.GetBooksByStatus("available")

			if err != nil {
				showTitle("Error!! ")
			}

			listBooks(books, "List Books")
		case 3:
			//Create Author
			clearScreen()
			showTitle("Create Author")
			fmt.Println("Enter name:")

			var name string
			fmt.Scanf("%s", &name)

			serviceAuthor := &services.AuthorService{}

			err := serviceAuthor.Create(name)

			if err == nil {
				showTitle("Data was save.")
			}

		case 4:
			fmt.Println("4")
		case 0:
			return
		}
	}
}

func selectSearchOption() {
	for {
		var option int

		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			fmt.Println("1")
		case 2:
			//Search Books
			clearScreen()
			serviceBook := &services.BookService{}

			showTitle("Insert word to search (title or author)")
			var wordToSearch string

			fmt.Scan(&wordToSearch)

			books, err := serviceBook.SearchBooksByTitileAndAuthor(wordToSearch)

			if err != nil {
				showTitle("Error!! ")
			}

			listBooks(books, "List Books")
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
	err := utils.InitDB()
	if err != nil {
		log.Fatal(err)
	}
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
			showListMenu()
			selectCreateOption()
		case 3:
			showSearchMenu()
			selectSearchOption()
		case 4:
			fmt.Println("4")
		case 5:
			fmt.Println("5")
		case 0:
			showTitle("Goodbye. Thanks for use owner system...")
			os.Exit(0)
		}
	}
}
