package Number

import (
	"math"
	"github.com/pangshu/toolbak/Convert"
)

// IsNan 是否为“非数值”.
func (*Number)IsNan(val interface{}) bool {
	var toolConvert Convert.Convert
	if toolConvert.IsFloat(val) {
		return math.IsNaN(toolConvert.ToFloat(val))
	}

	return !toolConvert.IsNumeric(val)
}
