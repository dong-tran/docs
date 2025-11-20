# Gang of Four Design Patterns in Go
## Complete Implementation of All 23 Classic Design Patterns

This directory contains working implementations of all 23 Gang of Four design patterns in Go, demonstrating their practical usage with real-world examples.

## 📚 Pattern Categories

### Creational Patterns (5)
**Focus**: Object creation mechanisms

| Pattern | Purpose | File |
|---------|---------|------|
| **Singleton** | Ensure only one instance exists with global access | `creational/singleton.go` |
| **Factory Method** | Create objects without specifying exact class | `creational/factory.go` |
| **Abstract Factory** | Create families of related objects | `creational/abstract_factory.go` |
| **Builder** | Construct complex objects step by step | `creational/builder.go` |
| **Prototype** | Clone objects without coupling to their classes | `creational/prototype.go` |

### Structural Patterns (7)
**Focus**: Object composition and relationships

| Pattern | Purpose | File |
|---------|---------|------|
| **Adapter** | Make incompatible interfaces work together | `structural/adapter.go` |
| **Bridge** | Decouple abstraction from implementation | `structural/bridge.go` |
| **Composite** | Compose objects into tree structures | `structural/composite.go` |
| **Decorator** | Add behavior to objects dynamically | `structural/decorator.go` |
| **Facade** | Provide simplified interface to complex subsystem | `structural/facade.go` |
| **Flyweight** | Share common state to support large numbers of objects | `structural/flyweight.go` |
| **Proxy** | Control access to objects | `structural/proxy.go` |

### Behavioral Patterns (11)
**Focus**: Object collaboration and responsibilities

| Pattern | Purpose | File |
|---------|---------|------|
| **Chain of Responsibility** | Pass requests along handler chain | `behavioral/chain_of_responsibility.go` |
| **Command** | Encapsulate requests as objects | `behavioral/command.go` |
| **Interpreter** | Define grammar and interpreter for language | `behavioral/interpreter.go` |
| **Iterator** | Access collection elements sequentially | `behavioral/iterator.go` |
| **Mediator** | Reduce coupling by centralizing communication | `behavioral/mediator.go` |
| **Memento** | Save and restore object state | `behavioral/memento.go` |
| **Observer** | Notify multiple objects of state changes | `behavioral/observer.go` |
| **State** | Change behavior when internal state changes | `behavioral/state.go` |
| **Strategy** | Select algorithm at runtime | `behavioral/strategy.go` |
| **Template Method** | Define algorithm skeleton, defer steps to subclasses | `behavioral/template_method.go` |
| **Visitor** | Add operations without modifying classes | `behavioral/visitor.go` |

## 🚀 Quick Start

Each pattern file is self-contained with:
- Clear explanation of the pattern
- Interface definitions
- Concrete implementations
- Demo function showing usage
- Real-world examples

### Running Examples

```bash
# View all patterns
ls -R creational/ structural/ behavioral/

# To test a specific pattern (after adding tests)
go test ./creational -v
go test ./structural -v
go test ./behavioral -v
```

## 📖 Pattern Selection Guide

### When to Use Creational Patterns
- **Singleton**: Database connections, configuration managers, logging
- **Factory Method**: Object creation when exact type is determined at runtime
- **Abstract Factory**: Need to create families of related objects (UI themes, database drivers)
- **Builder**: Complex object construction with many optional parameters
- **Prototype**: Cloning objects more efficient than creating from scratch

### When to Use Structural Patterns
- **Adapter**: Integrate legacy code or third-party libraries
- **Bridge**: Platform-independent abstractions (GUI frameworks, device drivers)
- **Composite**: Tree structures (file systems, UI components, org charts)
- **Decorator**: Add responsibilities dynamically (logging, caching, validation)
- **Facade**: Simplify complex libraries or subsystems
- **Flyweight**: Many similar objects (game particles, text characters, map tiles)
- **Proxy**: Lazy loading, access control, logging, caching

### When to Use Behavioral Patterns
- **Chain of Responsibility**: Request handling pipeline (middleware, event bubbling)
- **Command**: Undo/redo, task scheduling, transactional systems
- **Interpreter**: Domain-specific languages, query parsers, rule engines
- **Iterator**: Traverse collections without exposing structure
- **Mediator**: Complex communication between components (chat rooms, air traffic control)
- **Memento**: Undo/redo, snapshots, transaction rollback
- **Observer**: Event systems, pub/sub, model-view synchronization
- **State**: State machines, workflow engines, protocol handlers
- **Strategy**: Algorithm selection (payment methods, sorting, compression)
- **Template Method**: Framework extension points, algorithm skeletons
- **Visitor**: Operations on object structures (compilers, export formats)

## 🎯 Real-World Examples Included

Each pattern includes practical examples:

- **Singleton**: Database connection pool, logger
- **Factory/Abstract Factory**: Database drivers, UI component creation
- **Builder**: HTTP request builder, query builder
- **Prototype**: Document templates, configuration cloning
- **Adapter**: Legacy system integration, third-party API adaptation
- **Bridge**: Device abstraction (TV/Radio with different remote types)
- **Composite**: File system, graphics rendering
- **Decorator**: HTTP middleware, stream wrappers
- **Facade**: Computer startup, video conversion
- **Flyweight**: Forest rendering, text editor character styles
- **Proxy**: Virtual proxy (lazy loading), protection proxy (access control), cache proxy
- **Chain of Responsibility**: Approval workflow, middleware pipeline
- **Command**: Text editor operations, remote control
- **Interpreter**: Mathematical expression evaluator
- **Iterator**: Collection traversal (forward, reverse, filtered)
- **Mediator**: Chat room, air traffic control
- **Memento**: Text editor undo/redo
- **Observer**: Event system, notification system
- **State**: Vending machine, TCP connection
- **Strategy**: Payment processing, sorting algorithms
- **Template Method**: Data processing pipeline (CSV/JSON)
- **Visitor**: Shape operations (area, perimeter, export)

## 📚 Learning Path

### Beginner (Start Here)
1. **Strategy** - Simplest behavioral pattern
2. **Factory Method** - Basic creational pattern
3. **Observer** - Common event pattern
4. **Decorator** - Useful structural pattern
5. **Singleton** - Most well-known pattern

### Intermediate
6. **Builder** - Complex object creation
7. **Adapter** - Interface compatibility
8. **Command** - Action encapsulation
9. **Template Method** - Algorithm skeleton
10. **State** - State machine implementation

### Advanced
11. **Abstract Factory** - Factory of factories
12. **Composite** - Tree structures
13. **Proxy** - Access control
14. **Chain of Responsibility** - Request pipeline
15. **Visitor** - Double dispatch

### Expert
16. **Prototype** - Object cloning
17. **Bridge** - Abstraction separation
18. **Facade** - Subsystem simplification
19. **Flyweight** - Memory optimization
20. **Mediator** - Complex interactions
21. **Memento** - State preservation
22. **Iterator** - Collection traversal
23. **Interpreter** - Language implementation

## 🎓 Pattern Relationships

```
Creation → Structure → Behavior
    ↓           ↓          ↓
Singleton   Adapter    Observer
Factory     Decorator  Strategy
Builder     Composite  Command
Prototype   Proxy      State
Abstract    Bridge     Template
Factory     Facade     Chain
            Flyweight  Iterator
                       Mediator
                       Memento
                       Visitor
                       Interpreter
```

## 💡 Best Practices

1. **Don't Overuse**: Patterns solve specific problems - don't force them
2. **Combine Patterns**: Many real systems use multiple patterns together
3. **Understand Trade-offs**: Every pattern has pros and cons
4. **Code Clarity**: Pattern should make code clearer, not more complex
5. **Go Idioms**: Adapt patterns to Go's style (interfaces, composition over inheritance)

## 🔗 Related Resources

- Original Book: "Design Patterns: Elements of Reusable Object-Oriented Software" (1994)
- Presentation: See `/design/GoF-Design-Patterns-Presentation.md`
- Integration Example: See `/design/examples/relationships-integration`

## 📊 Pattern Statistics

- **Total Patterns**: 23
- **Creational**: 5 patterns (22%)
- **Structural**: 7 patterns (30%)
- **Behavioral**: 11 patterns (48%)
- **Lines of Code**: ~3000+ lines
- **Real-world Examples**: 23+ scenarios

---

**All 23 Gang of Four patterns implemented and ready to use!** 🎉
