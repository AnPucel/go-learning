package pointers

import (
	"errors"
	"fmt"
)

// Take aways:
// - Pointers
//   Go copies values when passing them to functions/methods
//   For functions that need to mutate the original pass a pointer
// - nil
//   Pointers can be nil. If a function returns a pointer you
//   need to check that it's not nil
// - Errors
//   Don't just check errors, handle them gracefully?
// - Create new types from existing ones
//   Useful for adding more domain specific meaning to values
//   Can let you implement interfaces

// A new type. This type & int are not interchangeable
// We can declare methods on this type
type Bitcoin int

// This is defined in the fmt package
// It lets you define how your type is printed when using %s
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol
// then it is private outside the package it's defined in.
// Methods with the wallet receiver can access it
type Wallet struct {
	balance Bitcoin
}

// You want pointers as the receiver versus a copy so the actual value gets changed (duh)
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Error is an interface
// If you see a function that takes arguments or returns values that are interfaces,
// they can be nillable.
// Must import errors

// Errors are values, so we can set it to a var here
// var keyword allows us to define values global to the package
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// This is also valid: (*w).balance
// language permits us to write w.balance, without an explicit dereference
// "struct pointers" are automatically dereferenced
// Returning a copy of w.balance is fine, but the receiver types should be consistent
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
