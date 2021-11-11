// Code generated by yaprotoc. DO NOT EDIT.
// source: oidb/0xef0.proto

package oidb0xef0

import (
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/proto"
)

type GroupInfoExt struct {
	GroupInfoExtSeq     *uint32 `protobuf:"varint,1,opt"`
	LuckyWordId         *uint64 `protobuf:"varint,2,opt"`
	LightCharNum        *uint32 `protobuf:"varint,3,opt"`
	LuckyWord           []byte  `protobuf:"bytes,4,opt"`
	StarId              *uint32 `protobuf:"varint,5,opt"`
	EssentialMsgSwitch  *uint32 `protobuf:"varint,6,opt"`
	TodoSeq             *uint32 `protobuf:"varint,7,opt"`
	BlacklistExpireTime *uint32 `protobuf:"varint,8,opt"`
}

func (x *GroupInfoExt) GetGroupInfoExtSeq() uint32 {
	if x != nil && x.GroupInfoExtSeq != nil {
		return *x.GroupInfoExtSeq
	}
	return 0
}

func (x *GroupInfoExt) GetLuckyWordId() uint64 {
	if x != nil && x.LuckyWordId != nil {
		return *x.LuckyWordId
	}
	return 0
}

func (x *GroupInfoExt) GetLightCharNum() uint32 {
	if x != nil && x.LightCharNum != nil {
		return *x.LightCharNum
	}
	return 0
}

func (x *GroupInfoExt) GetStarId() uint32 {
	if x != nil && x.StarId != nil {
		return *x.StarId
	}
	return 0
}

func (x *GroupInfoExt) GetEssentialMsgSwitch() uint32 {
	if x != nil && x.EssentialMsgSwitch != nil {
		return *x.EssentialMsgSwitch
	}
	return 0
}

func (x *GroupInfoExt) GetTodoSeq() uint32 {
	if x != nil && x.TodoSeq != nil {
		return *x.TodoSeq
	}
	return 0
}

func (x *GroupInfoExt) GetBlacklistExpireTime() uint32 {
	if x != nil && x.BlacklistExpireTime != nil {
		return *x.BlacklistExpireTime
	}
	return 0
}

func (x *GroupInfoExt) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type ReqBody struct {
	GroupCode    []uint64      `protobuf:"varint,1,rep"`
	GroupInfoExt *GroupInfoExt `protobuf:"bytes,2,opt"`
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspBody struct {
	RspDetail []*RspDetail `protobuf:"bytes,1,rep"`
}

func (x *RspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspDetail struct {
	GroupCode    *uint64       `protobuf:"varint,1,opt"`
	Result       *int32        `protobuf:"varint,2,opt"`
	GroupInfoExt *GroupInfoExt `protobuf:"bytes,3,opt"`
}

func (x *RspDetail) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *RspDetail) GetResult() int32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *RspDetail) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}