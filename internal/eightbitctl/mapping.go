package eightbitctl

type KeyCode string

const (
	KeyCodeA     KeyCode = "A"
	KeyCodeB     KeyCode = "B"
	KeyCodeX     KeyCode = "X"
	KeyCodeY     KeyCode = "Y"
	KeyCodeL     KeyCode = "L"
	KeyCodeR     KeyCode = "R"
	KeyCodeL2    KeyCode = "2L"
	KeyCodeR2    KeyCode = "2R"
	KeyCodeUP    KeyCode = "UP"
	KeyCodeDOWN  KeyCode = "DOWN"
	KeyCodeLEFT  KeyCode = "LEFT"
	KeyCodeRIGHT KeyCode = "RIGHT"
)

var keymap = map[int]KeyCode{
	304: KeyCodeA,
	305: KeyCodeB,
	307: KeyCodeX,
	308: KeyCodeY,
	310: KeyCodeL,
	311: KeyCodeR,
	312: KeyCodeL2,
	313: KeyCodeR2,
}

var verticalmap = map[int]KeyCode{
	0:   KeyCodeUP,
	255: KeyCodeDOWN,
}

var horizontalmap = map[int]KeyCode{
	0:   KeyCodeLEFT,
	255: KeyCodeRIGHT,
}
