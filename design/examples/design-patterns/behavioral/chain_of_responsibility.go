package behavioral

import "fmt"

// Chain of Responsibility Pattern
// Allows passing requests along a chain of handlers until one handles it.

type Request struct {
	RequestType string
	Amount      int
}

type Handler interface {
	SetNext(Handler) Handler
	Handle(*Request) string
}

// Base handler
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

func (h *BaseHandler) Handle(req *Request) string {
	if h.next != nil {
		return h.next.Handle(req)
	}
	return "Request not handled"
}

// Concrete Handlers
type Manager struct {
	BaseHandler
}

func (m *Manager) Handle(req *Request) string {
	if req.RequestType == "leave" && req.Amount <= 3 {
		return fmt.Sprintf("Manager approved %d day leave", req.Amount)
	}
	return m.BaseHandler.Handle(req)
}

type Director struct {
	BaseHandler
}

func (d *Director) Handle(req *Request) string {
	if req.RequestType == "leave" && req.Amount <= 7 {
		return fmt.Sprintf("Director approved %d day leave", req.Amount)
	}
	if req.RequestType == "purchase" && req.Amount <= 10000 {
		return fmt.Sprintf("Director approved $%d purchase", req.Amount)
	}
	return d.BaseHandler.Handle(req)
}

type CEO struct {
	BaseHandler
}

func (c *CEO) Handle(req *Request) string {
	if req.RequestType == "leave" {
		return fmt.Sprintf("CEO approved %d day leave", req.Amount)
	}
	if req.RequestType == "purchase" {
		return fmt.Sprintf("CEO approved $%d purchase", req.Amount)
	}
	return c.BaseHandler.Handle(req)
}

func DemoChainOfResponsibility() {
	fmt.Println("=== Chain of Responsibility Pattern Demo ===\n")

	manager := &Manager{}
	director := &Director{}
	ceo := &CEO{}

	manager.SetNext(director).SetNext(ceo)

	requests := []*Request{
		{RequestType: "leave", Amount: 2},
		{RequestType: "leave", Amount: 5},
		{RequestType: "leave", Amount: 10},
		{RequestType: "purchase", Amount: 5000},
		{RequestType: "purchase", Amount: 50000},
	}

	for _, req := range requests {
		result := manager.Handle(req)
		fmt.Printf("%s request for %d: %s\n", req.RequestType, req.Amount, result)
	}
}
