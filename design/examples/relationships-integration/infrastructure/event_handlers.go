package infrastructure

import (
"fmt"
"github.com/dong-tran/docs/integration-example/shared/patterns"
)

// Event handlers demonstrating Observer pattern

type EmailNotificationHandler struct{}

func (h *EmailNotificationHandler) OnEvent(event patterns.Event) {
	fmt.Printf("📧 Email Handler: %s - %+v\n", event.Type, event.Data)
}

type LoggingHandler struct{}

func (h *LoggingHandler) OnEvent(event patterns.Event) {
	fmt.Printf("📝 Logger: %s - %+v\n", event.Type, event.Data)
}

type AnalyticsHandler struct{}

func (h *AnalyticsHandler) OnEvent(event patterns.Event) {
	fmt.Printf("📊 Analytics: %s - %+v\n", event.Type, event.Data)
}
