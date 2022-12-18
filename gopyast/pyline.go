package gopyast

type (
	PyLine struct {
		startColOffset int
		stopColOffset  int
		currColOffset  int

		fileContent string

		prevLine *PyLine
		nextLine *PyLine
	}
)

func NewPyLine(prevLine *PyLine, fileContent string) (lne *PyLine, err error) {
	lne = &PyLine{
		fileContent: fileContent,
	}
	if prevLine == nil {
		lne.startColOffset = 0
		lne.currColOffset = 0
	}
	return
}
