package sebastion

type Input struct {
	Name, Description string
	Type              InputType
}

type InputType int

const (
	InputTypeInt InputType = iota
	InputTypeString
	InputTypeBool
)

type InputValues map[int]interface{}

func (i InputValues) GetInt(idx int) int {
	return i[idx].(int)
}
func (i InputValues) GetString(idx int) string {
	return i[idx].(string)
}
func (i InputValues) GetBool(idx int) bool {
	return i[idx].(bool)
}

type Action interface {
	Details() (name, description string)
	Inputs() []Input
	Run(InputValues) error
}
