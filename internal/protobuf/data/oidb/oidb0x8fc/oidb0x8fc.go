// Code generated by Proto2Go.Net. Only used for MiraiGo. DO NOT EDIT.
// Source: oidb0x8fc.proto

package oidb0x8fc

import (
	"github.com/pkg/errors"
	"go.dedis.ch/protobuf"
)

type (
	ReqBody struct {
		GroupCode      *int64        `protobuf:"1,opt"`
		ShowFlag       *int32        `protobuf:"2,opt"`
		MemLevelInfo   []*MemberInfo `protobuf:"3"`
		LevelName      []*LevelName  `protobuf:"4"`
		UpdateTime     *int32        `protobuf:"5,opt"`
		OfficeMode     *int32        `protobuf:"6,opt"`
		GroupOpenAppid *int32        `protobuf:"7,opt"`
		MsgClientInfo  *ClientInfo   `protobuf:"8,opt"`
		AuthKey        []byte        `protobuf:"9,opt"`
	}

	MemberInfo struct {
		Uin                    *int64          `protobuf:"1,opt"`
		Point                  *int32          `protobuf:"2,opt"`
		ActiveDay              *int32          `protobuf:"3,opt"`
		Level                  *int32          `protobuf:"4,opt"`
		SpecialTitle           []byte          `protobuf:"5,opt"`
		SpecialTitleExpireTime *int32          `protobuf:"6,opt"`
		UinName                []byte          `protobuf:"7,opt"`
		MemberCardName         []byte          `protobuf:"8,opt"`
		Phone                  []byte          `protobuf:"9,opt"`
		Email                  []byte          `protobuf:"10,opt"`
		Remark                 []byte          `protobuf:"11,opt"`
		Gender                 *int32          `protobuf:"12,opt"`
		Job                    []byte          `protobuf:"13,opt"`
		TribeLevel             *int32          `protobuf:"14,opt"`
		TribePoint             *int32          `protobuf:"15,opt"`
		RichCardName           []*CardNameElem `protobuf:"16"`
		CommRichCardName       []byte          `protobuf:"17,opt"`
	}

	CardNameElem struct {
		EnumCardType *int32 `protobuf:"1,opt"`
		Value        []byte `protobuf:"2,opt"`
	}

	LevelName struct {
		Level *int32  `protobuf:"1,opt"`
		Name  *string `protobuf:"2,opt"`
	}

	ClientInfo struct {
		Implat       *int32  `protobuf:"1,opt"`
		IngClientver *string `protobuf:"2,opt"`
	}

	CommCardNameBuf struct {
		RichCardName []*RichCardNameElem `protobuf:"1"`
	}

	RichCardNameElem struct {
		Ctrl []byte `protobuf:"1,opt"`
		Text []byte `protobuf:"2,opt"`
	}
)

func (x *ReqBody) GetGroupCode() int64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *ReqBody) GetShowFlag() int32 {
	if x != nil && x.ShowFlag != nil {
		return *x.ShowFlag
	}
	return 0
}

func (x *ReqBody) GetUpdateTime() int32 {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return 0
}

func (x *ReqBody) GetOfficeMode() int32 {
	if x != nil && x.OfficeMode != nil {
		return *x.OfficeMode
	}
	return 0
}

func (x *ReqBody) GetGroupOpenAppid() int32 {
	if x != nil && x.GroupOpenAppid != nil {
		return *x.GroupOpenAppid
	}
	return 0
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *MemberInfo) GetUin() int64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *MemberInfo) GetPoint() int32 {
	if x != nil && x.Point != nil {
		return *x.Point
	}
	return 0
}

func (x *MemberInfo) GetActiveDay() int32 {
	if x != nil && x.ActiveDay != nil {
		return *x.ActiveDay
	}
	return 0
}

func (x *MemberInfo) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *MemberInfo) GetSpecialTitleExpireTime() int32 {
	if x != nil && x.SpecialTitleExpireTime != nil {
		return *x.SpecialTitleExpireTime
	}
	return 0
}

func (x *MemberInfo) GetGender() int32 {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return 0
}

func (x *MemberInfo) GetTribeLevel() int32 {
	if x != nil && x.TribeLevel != nil {
		return *x.TribeLevel
	}
	return 0
}

func (x *MemberInfo) GetTribePoint() int32 {
	if x != nil && x.TribePoint != nil {
		return *x.TribePoint
	}
	return 0
}

func (x *MemberInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *CardNameElem) GetEnumCardType() int32 {
	if x != nil && x.EnumCardType != nil {
		return *x.EnumCardType
	}
	return 0
}

func (x *CardNameElem) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *LevelName) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *LevelName) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *LevelName) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *ClientInfo) GetImplat() int32 {
	if x != nil && x.Implat != nil {
		return *x.Implat
	}
	return 0
}

func (x *ClientInfo) GetIngClientver() string {
	if x != nil && x.IngClientver != nil {
		return *x.IngClientver
	}
	return ""
}

func (x *ClientInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *CommCardNameBuf) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *RichCardNameElem) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}
