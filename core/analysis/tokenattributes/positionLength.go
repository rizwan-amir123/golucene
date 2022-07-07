package tokenattributes

import (
	"github.com/rizwan-amir123/golucene/core/util"
)

type PositionLengthAttribute interface {
	util.Attribute
	SetPositionLength(int)
}
