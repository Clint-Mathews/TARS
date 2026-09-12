package accounting

import (
	"math"

	"github.com/Clint-Mathews/TARS/internal/model"
)

const DarwinBlockUnit int64 = 512

func AllocatedFromBlockss(blocks int64) (model.ByteValue, bool) {
	if blocks < 0 || blocks > math.MaxInt64/DarwinBlockUnit {
		return model.ByteValue{}, false
	}
	return model.ByteValue{
		Bytes: blocks * DarwinBlockUnit,
		Known: true,
	}, true
}

func Add(a, b model.ByteValue) (model.ByteValue, bool) {
	if !a.Known || !b.Known || a.Bytes < 0 || b.Bytes < 0 {
		return model.ByteValue{}, false
	}
	if b.Bytes > math.MaxInt64-a.Bytes {
		return model.ByteValue{}, false
	}
	return model.ByteValue{
		Bytes: a.Bytes + b.Bytes,
		Known: true,
	}, true
}

func Percent(part, whole model.ByteValue) (float64, bool) {
	if !part.Known || !whole.Known || part.Bytes <= 0 || whole.Bytes <= 0 {
		return 0, false
	}
	return float64(part.Bytes) / float64(whole.Bytes) * 100, true
}
