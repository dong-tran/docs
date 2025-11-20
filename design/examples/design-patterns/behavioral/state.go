package behavioral

import "fmt"

// State Pattern - Allows an object to alter its behavior when its internal state changes.

type State interface {
	InsertCoin(machine *VendingMachine)
	EjectCoin(machine *VendingMachine)
	PressButton(machine *VendingMachine)
	Dispense(machine *VendingMachine)
}

type VendingMachine struct {
	noCoinState    State
	hasCoinState   State
	soldState      State
	soldOutState   State
	currentState   State
	count          int
}

func NewVendingMachine(count int) *VendingMachine {
	vm := &VendingMachine{count: count}
	vm.noCoinState = &NoCoinState{}
	vm.hasCoinState = &HasCoinState{}
	vm.soldState = &SoldState{}
	vm.soldOutState = &SoldOutState{}
	if count > 0 {
		vm.currentState = vm.noCoinState
	} else {
		vm.currentState = vm.soldOutState
	}
	return vm
}

func (vm *VendingMachine) InsertCoin() {
	vm.currentState.InsertCoin(vm)
}

func (vm *VendingMachine) EjectCoin() {
	vm.currentState.EjectCoin(vm)
}

func (vm *VendingMachine) PressButton() {
	vm.currentState.PressButton(vm)
	vm.currentState.Dispense(vm)
}

func (vm *VendingMachine) SetState(state State) {
	vm.currentState = state
}

func (vm *VendingMachine) ReleaseItem() {
	if vm.count > 0 {
		fmt.Println("Item dispensed")
		vm.count--
	}
}

func (vm *VendingMachine) GetCount() int {
	return vm.count
}

type NoCoinState struct{}
func (s *NoCoinState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Coin inserted")
	vm.SetState(vm.hasCoinState)
}
func (s *NoCoinState) EjectCoin(vm *VendingMachine) {
	fmt.Println("No coin to eject")
}
func (s *NoCoinState) PressButton(vm *VendingMachine) {
	fmt.Println("Insert coin first")
}
func (s *NoCoinState) Dispense(vm *VendingMachine) {
	fmt.Println("Pay first")
}

type HasCoinState struct{}
func (s *HasCoinState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Coin already inserted")
}
func (s *HasCoinState) EjectCoin(vm *VendingMachine) {
	fmt.Println("Coin ejected")
	vm.SetState(vm.noCoinState)
}
func (s *HasCoinState) PressButton(vm *VendingMachine) {
	fmt.Println("Button pressed")
	vm.SetState(vm.soldState)
}
func (s *HasCoinState) Dispense(vm *VendingMachine) {
	fmt.Println("Press button first")
}

type SoldState struct{}
func (s *SoldState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Please wait, dispensing item")
}
func (s *SoldState) EjectCoin(vm *VendingMachine) {
	fmt.Println("Too late, item being dispensed")
}
func (s *SoldState) PressButton(vm *VendingMachine) {
	fmt.Println("Dispensing...")
}
func (s *SoldState) Dispense(vm *VendingMachine) {
	vm.ReleaseItem()
	if vm.GetCount() > 0 {
		vm.SetState(vm.noCoinState)
	} else {
		fmt.Println("Machine sold out")
		vm.SetState(vm.soldOutState)
	}
}

type SoldOutState struct{}
func (s *SoldOutState) InsertCoin(vm *VendingMachine) {
	fmt.Println("Machine sold out")
}
func (s *SoldOutState) EjectCoin(vm *VendingMachine) {
	fmt.Println("No coin to eject")
}
func (s *SoldOutState) PressButton(vm *VendingMachine) {
	fmt.Println("Machine sold out")
}
func (s *SoldOutState) Dispense(vm *VendingMachine) {
	fmt.Println("No items available")
}

func DemoState() {
	fmt.Println("=== State Pattern Demo ===\n")
	vm := NewVendingMachine(2)
	fmt.Printf("Items in machine: %d\n\n", vm.GetCount())
	vm.InsertCoin()
	vm.PressButton()
	fmt.Println()
	vm.InsertCoin()
	vm.PressButton()
	fmt.Println()
	vm.InsertCoin()
	vm.PressButton()
}
