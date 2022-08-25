package examples

import "github.com/joerdav/sebastion"

type Spam struct {
	message string
	repeat  int
}

func (cp *Spam) Details() (string, string) { return "Spam a message", "" }
func (cp *Spam) Inputs() []sebastion.Input {
	return []sebastion.Input{
		sebastion.NewMultiStringInput("Message", "The message to print", &cp.message,
			"Hello, world!",
			"Goodbye, world!",
			"Hello, Sebastion!"),
		sebastion.NewInput("Repetitions", "How man times to repeat", &cp.repeat),
	}
}
func (cp *Spam) Run(ctx sebastion.Context) error {
	for i := 0; i < cp.repeat; i++ {
		ctx.Logger.Println(cp.message)
	}
	return nil
}

type EchoSomething struct {
	text string
}

func (c *EchoSomething) Details() (string, string) { return "Echo", "Repeat whatever is passed in" }
func (c *EchoSomething) Inputs() []sebastion.Input {
	return []sebastion.Input{
		sebastion.NewInput("Text", "Some text to be echo-ed.", &c.text),
	}
}
func (c *EchoSomething) Run(ctx sebastion.Context) error {
	ctx.Logger.Println(c.text)
	return nil
}

type Panic struct{}

func (Panic) Details() (string, string) { return "Panic", "" }
func (Panic) Inputs() []sebastion.Input {
	return []sebastion.Input{}
}
func (Panic) Run(sebastion.Context) error {
	panic("panic")
}
