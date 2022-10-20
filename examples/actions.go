package examples

import (
	"fmt"
	"math"
	"time"

	"github.com/joerdav/sebastion"
)

// Validators

func greaterThan(than int) func(int) error {
	return func(i int) error {
		if i <= than {
			return fmt.Errorf("must be greater than %v", than)
		}
		return nil
	}
}

type Spam struct {
	message string
	repeat  int
}

func (cp *Spam) Details() sebastion.ActionDetails {
	return sebastion.ActionDetails{Name: "Spam", Description: "Print a message multiple times"}
}
func (cp *Spam) Inputs() []sebastion.Input {
	return []sebastion.Input{
		sebastion.NewMultiStringInput("Message", "The message to print", &cp.message,
			[]string{
				"Hello, World!",
				"Hello, Sebastion!",
				"Third Option",
			}, nil),
		sebastion.NewInput("Repetitions", "How many times to repeat", &cp.repeat,
			&sebastion.InputProps[int]{
				Default:   1,
				Validator: greaterThan(0),
			},
		),
	}
}

func (cp *Spam) Run(ctx sebastion.Context) error {
	for i := 0; i < cp.repeat; i++ {
		ctx.Logger.Println(cp.message)
		time.Sleep(time.Second)
	}
	return nil
}

type Primes struct {
	from, to int
}

func (cp *Primes) Details() sebastion.ActionDetails {
	return sebastion.ActionDetails{Name: "Primes", Description: "Calculate Prime Numbers"}
}
func (c *Primes) Inputs() []sebastion.Input {
	return []sebastion.Input{
		sebastion.NewInput("From", "Where to start searching for primes", &c.from, &sebastion.InputProps[int]{
			Default:   3,
			Validator: greaterThan(2),
		}),
		sebastion.NewInput("To", "Where to stop searching for primes", &c.to, &sebastion.InputProps[int]{
			Default:   100,
			Validator: greaterThan(2),
		}),
	}
}
func (c *Primes) Run(ctx sebastion.Context) error {
	from, to := c.from, c.to
	if from < 2 || to < 2 {
		return fmt.Errorf("Numbers must be greater than 2.")
	}
	for from <= c.to {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(from))); i++ {
			if from%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ctx.Logger.Printf("%d\n", from)
		}
		from++
	}
	return nil
}
