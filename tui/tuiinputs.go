package ui

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joerdav/sebastion"
)

// stringHandler is a TUIInputHandler that can handle string inputs.
type stringHandler struct{}

func (stringHandler) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[string])
	return ok
}

func (stringHandler) Get(i sebastion.Input) error {
	inp := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
		Default: i.Value.DefaultString(),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}
	return i.Value.Set(inp)
}

// boolHandler is a TUIInputHandler that can handle boolean inputs.
type boolHandler struct{}

func (boolHandler) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[bool])
	return ok
}

func (boolHandler) Get(i sebastion.Input) error {
	inp := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
		Default: i.Value.DefaultString() == "true",
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}
	return i.Value.Set(inp)
}

// intHandler is a TUIInputHandler that can handle int inputs.
type intHandler struct{}

func (intHandler) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[int])
	return ok
}

func (intHandler) Get(i sebastion.Input) error {
	inp := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
		Default: i.Value.DefaultString(),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(func(ans interface{}) error {
		str, ok := ans.(string)
		if !ok {
			return errors.New("This response must be a number.")
		}
		_, err := strconv.Atoi(str)
		if err != nil {
			return errors.New("This response must be a number.")
		}
		return nil
	}))
	if err != nil {
		return err
	}
	s, err := strconv.Atoi(inp)
	if err != nil {
		return errors.New("This response must be a number.")
	}
	return i.Value.Set(s)
}

// multiStringHandler is a TUIInputHandler that can handle string inputs with set options.
type multiStringHandler struct{}

func (multiStringHandler) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.MultiStringSelect)
	return ok
}

func (multiStringHandler) Get(i sebastion.Input) error {
	s, ok := i.Value.(sebastion.MultiStringSelect)
	if !ok {
		return errors.New("input is not a MultiStringSelect")
	}
	inp := ""
	prompt := &survey.Select{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
		Options: s.Options,
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(survey.Required))
	if err != nil {
		return err
	}
	return i.Value.Set(inp)
}

// float32Handler is a TUIInputHandler that can handle float32 inputs.
type float32Handler struct{}

func (float32Handler) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[float32])
	return ok
}

func (float32Handler) Get(i sebastion.Input) error {
	inp := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
		Default: i.Value.DefaultString(),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(func(ans interface{}) error {
		str, ok := ans.(string)
		if !ok {
			return errors.New("This response must be a number.")
		}
		_, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return errors.New("This response must be a number.")
		}
		return nil
	}))
	if err != nil {
		return err
	}
	s, err := strconv.ParseFloat(inp, 32)
	if err != nil {
		return errors.New("This response must be a number.")
	}
	return i.Value.Set(s)
}

// float32Handler is a TUIInputHandler that can handle float64 inputs.
type float64Handler struct{}

func (float64Handler) CanHandle(i sebastion.Input) bool {
	_, ok := i.Value.(sebastion.InputReference[float32])
	return ok
}

func (float64Handler) Get(i sebastion.Input) error {
	inp := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("%s - %s\n", i.Name, i.Description),
		Default: i.Value.DefaultString(),
	}
	err := survey.AskOne(prompt, &inp, survey.WithValidator(func(ans interface{}) error {
		str, ok := ans.(string)
		if !ok {
			return errors.New("This response must be a number.")
		}
		_, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return errors.New("This response must be a number.")
		}
		return nil
	}))
	if err != nil {
		return err
	}
	s, err := strconv.ParseFloat(inp, 64)
	if err != nil {
		return errors.New("This response must be a number.")
	}
	return i.Value.Set(s)
}
