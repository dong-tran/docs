package behavioral

import "fmt"

// Iterator Pattern
// Provides a way to access elements of a collection sequentially without exposing its underlying representation.

type Iterator interface {
	HasNext() bool
	Next() interface{}
	Current() interface{}
}

type Collection interface {
	CreateIterator() Iterator
}

// Concrete Collection
type BookShelf struct {
	books []string
}

func (b *BookShelf) AddBook(book string) {
	b.books = append(b.books, book)
}

func (b *BookShelf) CreateIterator() Iterator {
	return &BookIterator{
		shelf: b,
		index: 0,
	}
}

// Concrete Iterator
type BookIterator struct {
	shelf *BookShelf
	index int
}

func (i *BookIterator) HasNext() bool {
	return i.index < len(i.shelf.books)
}

func (i *BookIterator) Next() interface{} {
	if i.HasNext() {
		book := i.shelf.books[i.index]
		i.index++
		return book
	}
	return nil
}

func (i *BookIterator) Current() interface{} {
	if i.index > 0 && i.index <= len(i.shelf.books) {
		return i.shelf.books[i.index-1]
	}
	return nil
}

// Real-world example: Different iteration strategies
type User struct {
	Name string
	Age  int
}

type UserCollection struct {
	users []*User
}

func (uc *UserCollection) CreateIterator() Iterator {
	return &UserIterator{collection: uc, index: 0}
}

func (uc *UserCollection) CreateReverseIterator() Iterator {
	return &ReverseUserIterator{collection: uc, index: len(uc.users) - 1}
}

func (uc *UserCollection) CreateFilteredIterator(minAge int) Iterator {
	return &FilteredUserIterator{collection: uc, index: 0, minAge: minAge}
}

type UserIterator struct {
	collection *UserCollection
	index      int
}

func (i *UserIterator) HasNext() bool {
	return i.index < len(i.collection.users)
}

func (i *UserIterator) Next() interface{} {
	if i.HasNext() {
		user := i.collection.users[i.index]
		i.index++
		return user
	}
	return nil
}

func (i *UserIterator) Current() interface{} {
	if i.index > 0 && i.index <= len(i.collection.users) {
		return i.collection.users[i.index-1]
	}
	return nil
}

type ReverseUserIterator struct {
	collection *UserCollection
	index      int
}

func (i *ReverseUserIterator) HasNext() bool {
	return i.index >= 0
}

func (i *ReverseUserIterator) Next() interface{} {
	if i.HasNext() {
		user := i.collection.users[i.index]
		i.index--
		return user
	}
	return nil
}

func (i *ReverseUserIterator) Current() interface{} {
	if i.index >= -1 && i.index < len(i.collection.users)-1 {
		return i.collection.users[i.index+1]
	}
	return nil
}

type FilteredUserIterator struct {
	collection *UserCollection
	index      int
	minAge     int
}

func (i *FilteredUserIterator) HasNext() bool {
	for i.index < len(i.collection.users) {
		if i.collection.users[i.index].Age >= i.minAge {
			return true
		}
		i.index++
	}
	return false
}

func (i *FilteredUserIterator) Next() interface{} {
	if i.HasNext() {
		user := i.collection.users[i.index]
		i.index++
		return user
	}
	return nil
}

func (i *FilteredUserIterator) Current() interface{} {
	if i.index > 0 && i.index <= len(i.collection.users) {
		return i.collection.users[i.index-1]
	}
	return nil
}

func DemoIterator() {
	fmt.Println("=== Iterator Pattern Demo ===\n")

	fmt.Println("1. Book Collection:")
	shelf := &BookShelf{}
	shelf.AddBook("Design Patterns")
	shelf.AddBook("Clean Code")
	shelf.AddBook("Refactoring")

	iterator := shelf.CreateIterator()
	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println("Book:", book)
	}

	fmt.Println("\n2. User Collection with Different Iterators:")
	users := &UserCollection{
		users: []*User{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 17},
			{Name: "Charlie", Age: 30},
			{Name: "David", Age: 16},
			{Name: "Eve", Age: 28},
		},
	}

	fmt.Println("\nForward iteration:")
	iter := users.CreateIterator()
	for iter.HasNext() {
		user := iter.Next().(*User)
		fmt.Printf("%s (age %d)\n", user.Name, user.Age)
	}

	fmt.Println("\nReverse iteration:")
	reverseIter := users.CreateReverseIterator()
	for reverseIter.HasNext() {
		user := reverseIter.Next().(*User)
		fmt.Printf("%s (age %d)\n", user.Name, user.Age)
	}

	fmt.Println("\nFiltered iteration (age >= 18):")
	filteredIter := users.CreateFilteredIterator(18)
	for filteredIter.HasNext() {
		user := filteredIter.Next().(*User)
		fmt.Printf("%s (age %d)\n", user.Name, user.Age)
	}
}
