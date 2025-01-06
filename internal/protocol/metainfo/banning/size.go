package banning

import (
	"bitmagnet-io/bitmagnet/internal/protocol/metainfo"
	"errors"
)

type sizeChecker struct {
	min int64
}

func (c sizeChecker) Check(info metainfo.Info) error {
	if info.TotalLength() < c.min {
		return errors.New("size too small")
	}
	return nil
}
