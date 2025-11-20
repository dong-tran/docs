package isp

// BAD: Violates ISP - fat interface
type WorkerBad interface {
	Work()
	Eat()
	Sleep()
	Code()
	Manage()
}

// Robot has to implement methods it doesn't need
type RobotBad struct{}

func (r *RobotBad) Work()   {}
func (r *RobotBad) Eat()    {} // Robots don't eat!
func (r *RobotBad) Sleep()  {} // Robots don't sleep!
func (r *RobotBad) Code()   {}
func (r *RobotBad) Manage() {} // Not all robots manage!

// GOOD: Follows ISP - segregated interfaces
type Worker interface {
Work()
}

type Eater interface {
Eat()
}

type Sleeper interface {
Sleep()
}

type Coder interface {
Code()
}

type Manager interface {
Manage()
}

// Human implements relevant interfaces
type Human struct {
name string
}

func (h *Human) Work()   {}
func (h *Human) Eat()    {}
func (h *Human) Sleep()  {}
func (h *Human) Code()   {}
func (h *Human) Manage() {}

// Robot only implements what it needs
type Robot struct {
id string
}

func (r *Robot) Work() {}
func (r *Robot) Code() {}

// Functions depend only on what they need
func DoWork(w Worker) {
w.Work()
}

func FeedWorker(e Eater) {
e.Eat()
}
