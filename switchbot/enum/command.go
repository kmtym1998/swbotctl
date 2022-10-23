package enum

type Command string

var (
	TurnOn  Command = "turnOn"
	TurnOff Command = "turnOff"
)

func (c Command) String() string {
	return string(c)
}
