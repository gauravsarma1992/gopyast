package gopyast

type (
	PyBlock struct {
		startLineno int
		stopLineno  int
		currLineno  int
		colOffset   int

		prevBlock *PyBlock
		nextBlock *PyBlock

		depth int

		fileArr []string
	}
)

func NewPyBlock(prevBlock *PyBlock) (blk *PyBlock, err error) {
	blk = &PyBlock{}
	if prevBlock == nil {
		blk.startLineno = 0
		blk.colOffset = 0
		blk.depth = 0
	}
	return
}

func (blk *PyBlock) hasEnded(lineno int) (ended bool) {
	if blk.findColOffset(blk.currLineno) != blk.findColOffset(lineno) {
		ended = true
		return
	}
	return
}

func (blk *PyBlock) findColOffset(lineno int) (colOffset int) {
	line = fileArr[lineno]
	for _, chr := range line {
		if chr == "" {
			colOffset += 1
			continue
		}
		break
	}
	return
}

func (blk *PyBlock) startNewBlock() (err error) {
	var (
		newBlk *PyBlock
	)
	if newBlk, err = NewPyBlock(blk); err != nil {
		return
	}

	if err = newBlk.Walk(); err != nil {
		return
	}
	return
}

func (blk *PyBlock) Walk() (err error) {
	if blk.hasEnded() {
		if err = blk.updateEndValues(); err != nil {
			return
		}
		if err = blk.startNewBlock(); err != nil {
			return
		}
		return
	}
	return
}
