package arraysandslices

import (
	"learngowithtests/generics"
	"strings"
	"testing"
)

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	generics.AssertEqual(t, newBalanceFor(riya), 200)
	generics.AssertEqual(t, newBalanceFor(chris), 0)
	generics.AssertEqual(t, newBalanceFor(adil), 175)
}

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		if found != true {
			t.Errorf("Find(%d): expected true, got false", numbers[0])
		}
		generics.AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})
		if found != true {
			t.Errorf("Find: expected true, got false")
		}
		generics.AssertEqual(t, king, Person{Name: "Chris James"})
	})
}
