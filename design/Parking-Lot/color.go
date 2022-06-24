package design

type Color int

const (
	White Color = 0
	Black Color = 1
	Blue  Color = 2
)

var s2c = map[string]Color{
	"White": 0,
	"Black": 1,
	"Blue":  2,
}

var c2s = [...]string{
	"White",
	"Black",
	"Blue",
}

func StringToColor(color string) Color {
	return s2c[color]
}

func (c Color) String() string {
	return c2s[c]
}
