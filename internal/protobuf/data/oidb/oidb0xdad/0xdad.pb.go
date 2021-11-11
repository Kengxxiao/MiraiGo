// Code generated by yaprotoc. DO NOT EDIT.
// source: oidb/0xdad.proto

package oidb0xdad

import (
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/proto"
)

type LoginSig struct {
	Type  *uint32 `protobuf:"varint,1,opt"`
	Sig   []byte  `protobuf:"bytes,2,opt"`
	Appid *uint32 `protobuf:"varint,3,opt"`
}

func (x *LoginSig) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *LoginSig) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *LoginSig) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type ReqBody struct {
	Client    *int64    `protobuf:"varint,1,opt"`
	ProductId *int64    `protobuf:"varint,2,opt"`
	Amount    *int64    `protobuf:"varint,3,opt"`
	ToUin     *uint64   `protobuf:"varint,4,opt"`
	Gc        *uint64   `protobuf:"varint,5,opt"`
	Ip        *string   `protobuf:"bytes,6,opt"`
	Vsersion  *string   `protobuf:"bytes,7,opt"`
	Sig       *LoginSig `protobuf:"bytes,8,opt"`
}

func (x *ReqBody) GetClient() int64 {
	if x != nil && x.Client != nil {
		return *x.Client
	}
	return 0
}

func (x *ReqBody) GetProductId() int64 {
	if x != nil && x.ProductId != nil {
		return *x.ProductId
	}
	return 0
}

func (x *ReqBody) GetAmount() int64 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *ReqBody) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *ReqBody) GetGc() uint64 {
	if x != nil && x.Gc != nil {
		return *x.Gc
	}
	return 0
}

func (x *ReqBody) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *ReqBody) GetVsersion() string {
	if x != nil && x.Vsersion != nil {
		return *x.Vsersion
	}
	return ""
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspBody struct {
	Retcode *int64  `protobuf:"varint,1,opt"`
	Errmsg  *string `protobuf:"bytes,2,opt"`
}

func (x *RspBody) GetRetcode() int64 {
	if x != nil && x.Retcode != nil {
		return *x.Retcode
	}
	return 0
}

func (x *RspBody) GetErrmsg() string {
	if x != nil && x.Errmsg != nil {
		return *x.Errmsg
	}
	return ""
}

func (x *RspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}