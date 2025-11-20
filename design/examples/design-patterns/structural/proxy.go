package structural

import "fmt"

// Proxy Pattern
// Provides a surrogate or placeholder for another object to control access to it.

// Subject interface
type Image interface {
	Display()
}

// RealSubject
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	img := &RealImage{filename: filename}
	img.loadFromDisk()
	return img
}

func (img *RealImage) loadFromDisk() {
	fmt.Printf("Loading image from disk: %s\n", img.filename)
}

func (img *RealImage) Display() {
	fmt.Printf("Displaying image: %s\n", img.filename)
}

// Proxy
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (proxy *ProxyImage) Display() {
	if proxy.realImage == nil {
		proxy.realImage = NewRealImage(proxy.filename)
	}
	proxy.realImage.Display()
}

// Protection Proxy example
type Document interface {
	View()
	Edit(content string)
}

type RealDocument struct {
	content string
}

func (d *RealDocument) View() {
	fmt.Printf("Viewing document: %s\n", d.content)
}

func (d *RealDocument) Edit(content string) {
	d.content = content
	fmt.Printf("Document edited to: %s\n", content)
}

type ProtectedDocument struct {
	doc      *RealDocument
	password string
	user     string
}

func NewProtectedDocument(content, password string) *ProtectedDocument {
	return &ProtectedDocument{
		doc:      &RealDocument{content: content},
		password: password,
	}
}

func (p *ProtectedDocument) Authenticate(user, password string) bool {
	if password == p.password {
		p.user = user
		fmt.Printf("User '%s' authenticated successfully\n", user)
		return true
	}
	fmt.Printf("Authentication failed for user '%s'\n", user)
	return false
}

func (p *ProtectedDocument) View() {
	if p.user == "" {
		fmt.Println("Access denied: Please authenticate first")
		return
	}
	p.doc.View()
}

func (p *ProtectedDocument) Edit(content string) {
	if p.user == "" {
		fmt.Println("Access denied: Please authenticate first")
		return
	}
	p.doc.Edit(content)
}

// Caching Proxy example
type DatabaseQuery interface {
	Execute(query string) []string
}

type RealDatabase struct{}

func (db *RealDatabase) Execute(query string) []string {
	fmt.Printf("Executing expensive query on database: %s\n", query)
	return []string{"result1", "result2", "result3"}
}

type CachingDatabaseProxy struct {
	db    *RealDatabase
	cache map[string][]string
}

func NewCachingDatabaseProxy() *CachingDatabaseProxy {
	return &CachingDatabaseProxy{
		db:    &RealDatabase{},
		cache: make(map[string][]string),
	}
}

func (proxy *CachingDatabaseProxy) Execute(query string) []string {
	if result, exists := proxy.cache[query]; exists {
		fmt.Printf("Returning cached result for: %s\n", query)
		return result
	}

	result := proxy.db.Execute(query)
	proxy.cache[query] = result
	return result
}

func DemoProxy() {
	fmt.Println("=== Proxy Pattern Demo ===\n")

	// Virtual Proxy (lazy loading)
	fmt.Println("1. Virtual Proxy (Lazy Loading):")
	img1 := NewProxyImage("photo1.jpg")
	img2 := NewProxyImage("photo2.jpg")

	fmt.Println("\nFirst display calls (images not loaded yet):")
	img1.Display()
	img2.Display()

	fmt.Println("\nSecond display calls (images already loaded):")
	img1.Display()
	img2.Display()

	// Protection Proxy
	fmt.Println("\n\n2. Protection Proxy (Access Control):")
	doc := NewProtectedDocument("Confidential Information", "secret123")

	fmt.Println("\nTrying to access without authentication:")
	doc.View()
	doc.Edit("Modified content")

	fmt.Println("\nAuthenticating and accessing:")
	doc.Authenticate("john", "secret123")
	doc.View()
	doc.Edit("Updated by John")

	// Caching Proxy
	fmt.Println("\n\n3. Caching Proxy:")
	dbProxy := NewCachingDatabaseProxy()

	fmt.Println("\nFirst query (hits database):")
	result1 := dbProxy.Execute("SELECT * FROM users")
	fmt.Printf("Results: %v\n", result1)

	fmt.Println("\nSame query again (cached):")
	result2 := dbProxy.Execute("SELECT * FROM users")
	fmt.Printf("Results: %v\n", result2)

	fmt.Println("\nDifferent query (hits database):")
	result3 := dbProxy.Execute("SELECT * FROM products")
	fmt.Printf("Results: %v\n", result3)
}
