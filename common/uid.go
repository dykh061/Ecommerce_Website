package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/btcsuite/btcutil/base58"
)

// UID là phương thức dùng để tạo một định danh duy nhất (unique identifier) ảo cho toàn bộ hệ thống
// Cấu trúc của nó gồm 62 bit: LocalID - ObjectType - ShardID
// 32 bit dành cho Local ID, giá trị tối đa (2^32) - 1
// 10 bit dành cho loại đối tượng (Object Type)
// 18 bit dành cho Shard ID
type UID struct {
	localID    uint32
	objectType int
	shardID    uint32
}

func NewUID(localID uint32, objectType int, shardID uint32) *UID {
	return &UID{
		localID:    localID,
		objectType: objectType,
		shardID:    shardID,
	}
}

// Ví dụ: Shard = 1, ObjectType = 1, LocalID = 1
// Khi ghép bit lại ta được: 0001 0001 0001
//
// Dịch trái bit:
// 1 << 8  : đưa Shard lên vị trí bit cao      => 0001 0000 0000
// 1 << 4  : đưa ObjectType vào giữa           =>      1 0000
// 1 << 0  : LocalID giữ nguyên                =>            1
//
// OR các giá trị lại với nhau:
// => 0001 0001 0001
func (uid UID) String() string {
	val := uint64(uid.localID)<<28 | uint64(uid.objectType)<<18 | uint64(uid.shardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}
func (uid UID) GetLoacalID() uint32 {
	return uid.localID
}
func (uid UID) GetLoacalType() int {
	return uid.objectType
}
func (uid UID) GetShardID() uint32 {
	return uid.shardID
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return UID{}, err
	}
	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}

	// Ví dụ thao tác bit:
	// x = 1110 1110 0101
	// Dịch phải 4 bit: x >> 4 = 1110 1110
	// Mask 4 bit thấp: (x >> 4) & 0000 1111 = 1110
	//
	// Lý do phải AND (&) với 0x3FFF:
	// 0x3FFF = 0011 1111 1111 1111 (14 bit 1)
	// Phép AND dùng để loại bỏ các bit không liên quan ở bên trái,
	// chỉ giữ lại đúng số bit cần thiết cho giá trị (ví dụ: ObjectType / LocalID),
	// tránh việc các bit của shard hoặc phần khác làm sai kết quả.
	u := UID{
		localID:    uint32(uid >> 28),
		objectType: int(uid >> 18 & 0x3FF),
		shardID:    uint32(uid & 0x3FFFF),
	}
	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}
func (uid UID) UnmarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))
	if err != nil {
		return err
	}
	uid.localID = decodeUID.localID
	uid.objectType = decodeUID.objectType
	uid.shardID = decodeUID.shardID
	return nil
}

// đoạn hỗ trợ sinh sku cho variant (product variant)
var seq uint32

func GenLocalID() uint32 {
	t := uint32(time.Now().Unix()) // giây
	s := atomic.AddUint32(&seq, 1) & 0xFFFF
	return (t << 16) | s
}
