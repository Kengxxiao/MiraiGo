// Code generated by yaprotoc. DO NOT EDIT.
// source: oidb/0x6d7.proto

package oidb0x6d7

import (
	group_file_common "github.com/Mrs4s/MiraiGo/internal/protobuf/data/group_file_common"
	"github.com/pkg/errors"
	"github.com/segmentio/encoding/proto"
)

type CreateFolderReqBody struct {
	GroupCode      *uint64 `protobuf:"varint,1,opt"`
	AppId          *uint32 `protobuf:"varint,2,opt"`
	ParentFolderId *string `protobuf:"bytes,3,opt"`
	FolderName     *string `protobuf:"bytes,4,opt"`
}

func (x *CreateFolderReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *CreateFolderReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *CreateFolderReqBody) GetParentFolderId() string {
	if x != nil && x.ParentFolderId != nil {
		return *x.ParentFolderId
	}
	return ""
}

func (x *CreateFolderReqBody) GetFolderName() string {
	if x != nil && x.FolderName != nil {
		return *x.FolderName
	}
	return ""
}

func (x *CreateFolderReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type CreateFolderRspBody struct {
	RetCode       *int32                        `protobuf:"varint,1,opt"`
	RetMsg        *string                       `protobuf:"bytes,2,opt"`
	ClientWording *string                       `protobuf:"bytes,3,opt"`
	FolderInfo    *group_file_common.FolderInfo `protobuf:"bytes,4,opt"`
}

func (x *CreateFolderRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *CreateFolderRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *CreateFolderRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *CreateFolderRspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type DeleteFolderReqBody struct {
	GroupCode *uint64 `protobuf:"varint,1,opt"`
	AppId     *uint32 `protobuf:"varint,2,opt"`
	FolderId  *string `protobuf:"bytes,3,opt"`
}

func (x *DeleteFolderReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *DeleteFolderReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *DeleteFolderReqBody) GetFolderId() string {
	if x != nil && x.FolderId != nil {
		return *x.FolderId
	}
	return ""
}

func (x *DeleteFolderReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type DeleteFolderRspBody struct {
	RetCode       *int32  `protobuf:"varint,1,opt"`
	RetMsg        *string `protobuf:"bytes,2,opt"`
	ClientWording *string `protobuf:"bytes,3,opt"`
}

func (x *DeleteFolderRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *DeleteFolderRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *DeleteFolderRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *DeleteFolderRspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type MoveFolderReqBody struct {
	GroupCode      *uint64 `protobuf:"varint,1,opt"`
	AppId          *uint32 `protobuf:"varint,2,opt"`
	FolderId       *string `protobuf:"bytes,3,opt"`
	ParentFolderId *string `protobuf:"bytes,4,opt"`
	DestFolderId   *string `protobuf:"bytes,5,opt"`
}

func (x *MoveFolderReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *MoveFolderReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *MoveFolderReqBody) GetFolderId() string {
	if x != nil && x.FolderId != nil {
		return *x.FolderId
	}
	return ""
}

func (x *MoveFolderReqBody) GetParentFolderId() string {
	if x != nil && x.ParentFolderId != nil {
		return *x.ParentFolderId
	}
	return ""
}

func (x *MoveFolderReqBody) GetDestFolderId() string {
	if x != nil && x.DestFolderId != nil {
		return *x.DestFolderId
	}
	return ""
}

func (x *MoveFolderReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type MoveFolderRspBody struct {
	RetCode       *int32                        `protobuf:"varint,1,opt"`
	RetMsg        *string                       `protobuf:"bytes,2,opt"`
	ClientWording *string                       `protobuf:"bytes,3,opt"`
	FolderInfo    *group_file_common.FolderInfo `protobuf:"bytes,4,opt"`
}

func (x *MoveFolderRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *MoveFolderRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *MoveFolderRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *MoveFolderRspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RenameFolderReqBody struct {
	GroupCode     *uint64 `protobuf:"varint,1,opt"`
	AppId         *uint32 `protobuf:"varint,2,opt"`
	FolderId      *string `protobuf:"bytes,3,opt"`
	NewFolderName *string `protobuf:"bytes,4,opt"`
}

func (x *RenameFolderReqBody) GetGroupCode() uint64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *RenameFolderReqBody) GetAppId() uint32 {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return 0
}

func (x *RenameFolderReqBody) GetFolderId() string {
	if x != nil && x.FolderId != nil {
		return *x.FolderId
	}
	return ""
}

func (x *RenameFolderReqBody) GetNewFolderName() string {
	if x != nil && x.NewFolderName != nil {
		return *x.NewFolderName
	}
	return ""
}

func (x *RenameFolderReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RenameFolderRspBody struct {
	RetCode       *int32                        `protobuf:"varint,1,opt"`
	RetMsg        *string                       `protobuf:"bytes,2,opt"`
	ClientWording *string                       `protobuf:"bytes,3,opt"`
	FolderInfo    *group_file_common.FolderInfo `protobuf:"bytes,4,opt"`
}

func (x *RenameFolderRspBody) GetRetCode() int32 {
	if x != nil && x.RetCode != nil {
		return *x.RetCode
	}
	return 0
}

func (x *RenameFolderRspBody) GetRetMsg() string {
	if x != nil && x.RetMsg != nil {
		return *x.RetMsg
	}
	return ""
}

func (x *RenameFolderRspBody) GetClientWording() string {
	if x != nil && x.ClientWording != nil {
		return *x.ClientWording
	}
	return ""
}

func (x *RenameFolderRspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type ReqBody struct {
	CreateFolderReq *CreateFolderReqBody `protobuf:"bytes,1,opt"`
	DeleteFolderReq *DeleteFolderReqBody `protobuf:"bytes,2,opt"`
	RenameFolderReq *RenameFolderReqBody `protobuf:"bytes,3,opt"`
	MoveFolderReq   *MoveFolderReqBody   `protobuf:"bytes,4,opt"`
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}

type RspBody struct {
	CreateFolderRsp *CreateFolderRspBody `protobuf:"bytes,1,opt"`
	DeleteFolderRsp *DeleteFolderRspBody `protobuf:"bytes,2,opt"`
	RenameFolderRsp *RenameFolderRspBody `protobuf:"bytes,3,opt"`
	MoveFolderRsp   *MoveFolderRspBody   `protobuf:"bytes,4,opt"`
}

func (x *RspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil message")
	}
	return proto.Marshal(x)
}