package pathoperate

import (
	"testing"
)

func TestStatOperate(t *testing.T) {
	path1 := "./"
	StatOperate(path1)
}

func TestFilePathOperate(t *testing.T) {
	path1 := "./pathOperate"
	FilePathOperate(path1)

}

func TestDirWalk(t *testing.T) {
	path1 := "../pathOperate"
	DirWalk(path1)

}
