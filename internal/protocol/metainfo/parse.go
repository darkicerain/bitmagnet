package metainfo

import (
	"bitmagnet-io/bitmagnet/internal/protocol"
	"errors"
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	mi "github.com/anacrolix/torrent/metainfo"
)

func ParseMetaInfoBytes(infoHash protocol.ID, metaInfoBytes []byte) (Info, error) {
	if protocol.ID(mi.HashBytes(metaInfoBytes)) != infoHash {
		return Info{}, errors.New("info bytes have wrong hash")
	}
	var info Info
	if unmarshalErr := bencode.Unmarshal(metaInfoBytes, &info); unmarshalErr != nil {
		return Info{}, fmt.Errorf("error unmarshaling info bytes: %s", unmarshalErr)
	}
	return info, nil
}
