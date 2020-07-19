package client

import (
	"errors"
	"github.com/Mrs4s/MiraiGo/binary"
	"github.com/Mrs4s/MiraiGo/binary/jce"
	"github.com/Mrs4s/MiraiGo/client/pb"
	"github.com/Mrs4s/MiraiGo/client/pb/msg"
	"github.com/Mrs4s/MiraiGo/client/pb/structmsg"
	"github.com/golang/protobuf/proto"
	"sync"
	"time"
)

var (
	groupJoinLock  = new(sync.Mutex)
	groupLeaveLock = new(sync.Mutex)
)

func decodeLoginResponse(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	reader := binary.NewReader(payload)
	reader.ReadUInt16() // sub command
	t := reader.ReadByte()
	reader.ReadUInt16()
	m := reader.ReadTlvMap(2)
	if t == 0 { // login success
		if t150, ok := m[0x150]; ok {
			c.t150 = t150
		}
		if t161, ok := m[0x161]; ok {
			c.decodeT161(t161)
		}
		c.decodeT119(m[0x119])
		return LoginResponse{
			Success: true,
		}, nil
	}
	if t == 2 {
		c.t104, _ = m[0x104]
		if m.Exists(0x192) { // slider, not supported yet
			return LoginResponse{
				Success: false,
				Error:   UnknownLoginError,
			}, nil
		}
		if m.Exists(0x165) { // image
			imgData := binary.NewReader(m[0x105])
			signLen := imgData.ReadUInt16()
			imgData.ReadUInt16()
			sign := imgData.ReadBytes(int(signLen))
			return LoginResponse{
				Success:      false,
				Error:        NeedCaptcha,
				CaptchaImage: imgData.ReadAvailable(),
				CaptchaSign:  sign,
			}, nil
		} else {
			return LoginResponse{
				Success: false,
				Error:   UnknownLoginError,
			}, nil
		}
	} // need captcha

	if t == 160 {
		return LoginResponse{
			Success:      false,
			Error:        UnsafeDeviceError,
			VerifyUrl:    string(m[0x204]),
			ErrorMessage: "",
		}, nil
	}

	if t == 204 {
		c.t104 = m[0x104]
		return c.sendAndWait(c.buildDeviceLockLoginPacket(m[0x402]))
	} // drive lock

	if t149, ok := m[0x149]; ok {
		t149r := binary.NewReader(t149)
		t149r.ReadBytes(2)
		t149r.ReadStringShort() // title
		return LoginResponse{
			Success:      false,
			Error:        OtherLoginError,
			ErrorMessage: t149r.ReadStringShort(),
		}, nil
	}

	if t146, ok := m[0x146]; ok {
		t146r := binary.NewReader(t146)
		t146r.ReadBytes(4)      // ver and code
		t146r.ReadStringShort() // title
		return LoginResponse{
			Success:      false,
			Error:        OtherLoginError,
			ErrorMessage: t146r.ReadStringShort(),
		}, nil
	}

	return nil, nil // ?
}

func decodeClientRegisterResponse(_ *QQClient, _ uint16, payload []byte) (interface{}, error) {
	request := &jce.RequestPacket{}
	request.ReadFrom(jce.NewJceReader(payload))
	data := &jce.RequestDataVersion2{}
	data.ReadFrom(jce.NewJceReader(request.SBuffer))
	return nil, nil
}

func decodePushReqPacket(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	request := &jce.RequestPacket{}
	request.ReadFrom(jce.NewJceReader(payload))
	data := &jce.RequestDataVersion2{}
	data.ReadFrom(jce.NewJceReader(request.SBuffer))
	r := jce.NewJceReader(data.Map["PushReq"]["ConfigPush.PushReq"][1:])
	jceBuf := []byte{}
	t := r.ReadInt32(1)
	r.ReadSlice(&jceBuf, 2)
	seq := r.ReadInt64(3)
	_, pkt := c.buildConfPushRespPacket(t, seq, jceBuf)
	return nil, c.send(pkt)
}

func decodeMessageSvcPacket(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	rsp := msg.GetMessageResponse{}
	err := proto.Unmarshal(payload, &rsp)
	if err != nil {
		return nil, err
	}
	if rsp.Result != 0 {
		return nil, errors.New("message svc result unsuccessful")
	}
	c.syncCookie = rsp.SyncCookie
	c.pubAccountCookie = rsp.PubAccountCookie
	c.msgCtrlBuf = rsp.MsgCtrlBuf
	if rsp.UinPairMsgs == nil {
		return nil, nil
	}
	var delItems []*pb.MessageItem
	for _, pairMsg := range rsp.UinPairMsgs {
		for _, message := range pairMsg.Messages {
			// delete message
			delItem := &pb.MessageItem{
				FromUin: message.Head.FromUin,
				ToUin:   message.Head.ToUin,
				MsgType: 187,
				MsgSeq:  message.Head.MsgSeq,
				MsgUid:  message.Head.MsgUid,
			}
			delItems = append(delItems, delItem)
			if message.Head.ToUin != c.Uin {
				continue
			}
			switch message.Head.MsgType {
			case 33: // 加群同步
				groupJoinLock.Lock()
				group := c.FindGroupByUin(message.Head.FromUin)
				if message.Head.AuthUin == c.Uin {
					if group == nil && c.ReloadGroupList() == nil {
						c.dispatchJoinGroupEvent(c.FindGroupByUin(message.Head.FromUin))
					}
				} else {
					if group != nil && group.FindMember(message.Head.AuthUin) == nil {
						mem := &GroupMemberInfo{
							Uin: message.Head.AuthUin,
							Nickname: func() string {
								if message.Head.AuthNick == "" {
									return message.Head.FromNick
								}
								return message.Head.AuthNick
							}(),
							JoinTime:   time.Now().Unix(),
							Permission: Member,
						}
						group.Members = append(group.Members, mem)
						c.dispatchNewMemberEvent(&MemberJoinGroupEvent{
							Group:  group,
							Member: mem,
						})
					}
				}
				groupJoinLock.Unlock()
			case 84, 87:
				_, pkt := c.buildSystemMsgNewGroupPacket()
				_ = c.send(pkt)
			case 141: // 临时会话
				if message.Head.C2CTmpMsgHead == nil {
					continue
				}
				group := c.FindGroupByUin(message.Head.C2CTmpMsgHead.GroupUin)
				if group == nil {
					continue
				}
				mem := group.FindMember(message.Head.FromUin)
				if mem == nil || message.Head.FromUin == c.Uin {
					continue
				}
				lastSeq, ok := c.lastMessageSeqTmp.Load(mem.Uin)
				if !ok {
					c.lastMessageSeqTmp.Store(mem.Uin, int32(-1))
					lastSeq = int32(-1)
				}
				if message.Head.MsgSeq > lastSeq.(int32) {
					c.lastMessageSeqTmp.Store(mem.Uin, message.Head.MsgSeq)
					c.dispatchTempMessage(c.parseTempMessage(message))
				}
			case 166: // 好友消息
				if message.Body.RichText == nil || message.Body.RichText.Elems == nil {
					continue
				}
				if c.lastMessageSeq >= message.Head.MsgSeq {
					continue
				}
				c.lastMessageSeq = message.Head.MsgSeq
				c.dispatchFriendMessage(c.parsePrivateMessage(message))
			case 187:
				_, pkt := c.buildSystemMsgNewFriendPacket()
				_ = c.send(pkt)
			}
		}
	}
	_, _ = c.sendAndWait(c.buildDeleteMessageRequestPacket(delItems))
	if rsp.SyncFlag != msg.SyncFlag_STOP {
		seq, pkt := c.buildGetMessageRequestPacket(rsp.SyncFlag, time.Now().Unix())
		_, _ = c.sendAndWait(seq, pkt)
	}
	return nil, err
}

func decodeGroupMessagePacket(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	pkt := msg.PushMessagePacket{}
	err := proto.Unmarshal(payload, &pkt)
	if err != nil {
		return nil, err
	}
	if pkt.Message.Head.FromUin == c.Uin {
		c.dispatchGroupMessageReceiptEvent(&groupMessageReceiptEvent{
			Rand: pkt.Message.Body.RichText.Attr.Random,
			Seq:  pkt.Message.Head.MsgSeq,
		})
		return nil, nil
	}
	if pkt.Message.Content != nil && pkt.Message.Content.PkgNum > 1 {
		var builder *groupMessageBuilder
		i, ok := c.groupMsgBuilders.Load(pkt.Message.Content.DivSeq)
		if !ok {
			builder = &groupMessageBuilder{
				MessageSeq:   pkt.Message.Content.DivSeq,
				MessageCount: pkt.Message.Content.PkgNum,
			}
			c.groupMsgBuilders.Store(pkt.Message.Content.DivSeq, builder)
		} else {
			builder = i.(*groupMessageBuilder)
		}
		builder.MessageSlices = append(builder.MessageSlices, pkt.Message)
		if int32(len(builder.MessageSlices)) >= builder.MessageCount {
			c.groupMsgBuilders.Delete(pkt.Message.Content.DivSeq)
			c.dispatchGroupMessage(c.parseGroupMessage(builder.build()))
		}
		return nil, nil
	}
	c.dispatchGroupMessage(c.parseGroupMessage(pkt.Message))
	return nil, nil
}

func decodeSvcNotify(c *QQClient, _ uint16, _ []byte) (interface{}, error) {
	_, pkt := c.buildGetMessageRequestPacket(msg.SyncFlag_START, time.Now().Unix())
	return nil, c.send(pkt)
}

func decodeFriendGroupListResponse(_ *QQClient, _ uint16, payload []byte) (interface{}, error) {
	request := &jce.RequestPacket{}
	request.ReadFrom(jce.NewJceReader(payload))
	data := &jce.RequestDataVersion3{}
	data.ReadFrom(jce.NewJceReader(request.SBuffer))
	r := jce.NewJceReader(data.Map["FLRESP"][1:])
	totalFriendCount := r.ReadInt16(5)
	friends := []jce.FriendInfo{}
	r.ReadSlice(&friends, 7)
	var l []*FriendInfo
	for _, f := range friends {
		l = append(l, &FriendInfo{
			Uin:      f.FriendUin,
			Nickname: f.Nick,
			Remark:   f.Remark,
			FaceId:   f.FaceId,
		})
	}
	rsp := FriendListResponse{
		TotalCount: int32(totalFriendCount),
		List:       l,
	}
	return rsp, nil
}

func decodeGroupListResponse(_ *QQClient, _ uint16, payload []byte) (interface{}, error) {
	request := &jce.RequestPacket{}
	request.ReadFrom(jce.NewJceReader(payload))
	data := &jce.RequestDataVersion3{}
	data.ReadFrom(jce.NewJceReader(request.SBuffer))
	r := jce.NewJceReader(data.Map["GetTroopListRespV2"][1:])
	groups := []jce.TroopNumber{}
	r.ReadSlice(&groups, 5)
	var l []*GroupInfo
	for _, g := range groups {
		l = append(l, &GroupInfo{
			Uin:            g.GroupUin,
			Code:           g.GroupCode,
			Name:           g.GroupName,
			Memo:           g.GroupMemo,
			OwnerUin:       g.GroupOwnerUin,
			MemberCount:    uint16(g.MemberNum),
			MaxMemberCount: uint16(g.MaxGroupMemberNum),
		})
	}
	return l, nil
}

func decodeGroupMemberListResponse(_ *QQClient, _ uint16, payload []byte) (interface{}, error) {
	request := &jce.RequestPacket{}
	request.ReadFrom(jce.NewJceReader(payload))
	data := &jce.RequestDataVersion3{}
	data.ReadFrom(jce.NewJceReader(request.SBuffer))
	r := jce.NewJceReader(data.Map["GTMLRESP"][1:])
	members := []jce.TroopMemberInfo{}
	r.ReadSlice(&members, 3)
	next := r.ReadInt64(4)
	var l []*GroupMemberInfo
	for _, m := range members {
		l = append(l, &GroupMemberInfo{
			Uin:                    m.MemberUin,
			Nickname:               m.Nick,
			CardName:               m.Name,
			Level:                  uint16(m.MemberLevel),
			JoinTime:               m.JoinTime,
			LastSpeakTime:          m.LastSpeakTime,
			SpecialTitle:           m.SpecialTitle,
			SpecialTitleExpireTime: m.SpecialTitleExpireTime,
			Permission: func() MemberPermission {
				if m.Flag == 1 {
					return Administrator
				}
				return Member
			}(),
		})
	}
	return groupMemberListResponse{
		NextUin: next,
		list:    l,
	}, nil
}

func decodeGroupImageStoreResponse(_ *QQClient, _ uint16, payload []byte) (interface{}, error) {
	pkt := pb.D388RespBody{}
	err := proto.Unmarshal(payload, &pkt)
	if err != nil {
		return nil, err
	}
	rsp := pkt.MsgTryupImgRsp[0]
	if rsp.Result != 0 {
		return groupImageUploadResponse{
			ResultCode: rsp.Result,
			Message:    rsp.FailMsg,
		}, nil
	}
	if rsp.BoolFileExit {
		return groupImageUploadResponse{IsExists: true}, nil
	}
	return groupImageUploadResponse{
		UploadKey:  rsp.UpUkey,
		UploadIp:   rsp.Uint32UpIp,
		UploadPort: rsp.Uint32UpPort,
	}, nil
}

func decodeOnlinePushReqPacket(c *QQClient, seq uint16, payload []byte) (interface{}, error) {
	request := &jce.RequestPacket{}
	request.ReadFrom(jce.NewJceReader(payload))
	data := &jce.RequestDataVersion2{}
	data.ReadFrom(jce.NewJceReader(request.SBuffer))
	jr := jce.NewJceReader(data.Map["req"]["OnlinePushPack.SvcReqPushMsg"][1:])
	msgInfos := []jce.PushMessageInfo{}
	uin := jr.ReadInt64(0)
	jr.ReadSlice(&msgInfos, 2)
	_ = c.send(c.buildDeleteOnlinePushPacket(uin, seq, msgInfos))
	seqExists := func(ms int16) bool {
		for _, s := range c.onlinePushCache {
			if ms == s {
				return true
			}
		}
		return false
	}
	for _, m := range msgInfos {
		if seqExists(m.MsgSeq) {
			continue
		}
		c.onlinePushCache = append(c.onlinePushCache, m.MsgSeq)
		if m.MsgType == 732 {
			r := binary.NewReader(m.VMsg)
			groupId := int64(uint32(r.ReadInt32()))
			iType := r.ReadByte()
			r.ReadByte()
			switch iType {
			case 0x0c: // 群内禁言
				operator := int64(uint32(r.ReadInt32()))
				if operator == c.Uin {
					continue
				}
				r.ReadBytes(6)
				target := int64(uint32(r.ReadInt32()))
				t := r.ReadInt32()
				c.dispatchGroupMuteEvent(&GroupMuteEvent{
					GroupCode:   groupId,
					OperatorUin: operator,
					TargetUin:   target,
					Time:        t,
				})
			case 0x11: // 撤回消息
				r.ReadByte()
				b := pb.NotifyMsgBody{}
				_ = proto.Unmarshal(r.ReadAvailable(), &b)
				if b.OptMsgRecall == nil {
					continue
				}
				for _, rm := range b.OptMsgRecall.RecalledMsgList {
					c.dispatchGroupMessageRecalledEvent(&GroupMessageRecalledEvent{
						GroupCode:   groupId,
						OperatorUin: b.OptMsgRecall.Uin,
						AuthorUin:   rm.AuthorUin,
						MessageId:   rm.Seq,
						Time:        rm.Time,
					})
				}
			}
		}
		if m.MsgType == 528 {
			vr := jce.NewJceReader(m.VMsg)
			subType := vr.ReadInt64(0)
			probuf := vr.ReadAny(10).([]byte)
			switch subType {
			case 0xD4:
				d4 := pb.SubD4{}
				if err := proto.Unmarshal(probuf, &d4); err != nil {
					return nil, err
				}
				groupLeaveLock.Lock()
				if g := c.FindGroupByUin(d4.Uin); g != nil {
					if err := c.ReloadGroupList(); err != nil {
						groupLeaveLock.Unlock()
						return nil, err
					}
					c.dispatchLeaveGroupEvent(&GroupLeaveEvent{Group: g})
				}
				groupLeaveLock.Unlock()
			}
		}
	}

	return nil, nil
}

func decodeOnlinePushTransPacket(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	info := msg.TransMsgInfo{}
	err := proto.Unmarshal(payload, &info)
	if err != nil {
		return nil, err
	}
	data := binary.NewReader(info.MsgData)
	if info.MsgType == 34 {
		data.ReadInt32()
		data.ReadByte()
		target := int64(uint32(data.ReadInt32()))
		typ := int32(data.ReadByte())
		operator := int64(uint32(data.ReadInt32()))
		if g := c.FindGroupByUin(info.FromUin); g != nil {
			switch typ {
			case 0x03:
				groupLeaveLock.Lock()
				defer groupLeaveLock.Unlock()
				if err = c.ReloadGroupList(); err != nil {
					return nil, err
				}
				c.dispatchLeaveGroupEvent(&GroupLeaveEvent{
					Group:    g,
					Operator: g.FindMember(operator),
				})
			case 0x82:
				if m := g.FindMember(target); m != nil {
					g.removeMember(m.Uin)
					c.dispatchMemberLeaveEvent(&MemberLeaveGroupEvent{
						Group:  g,
						Member: m,
					})
				}
			case 0x83:
				if m := g.FindMember(target); m != nil {
					g.removeMember(m.Uin)
					c.dispatchMemberLeaveEvent(&MemberLeaveGroupEvent{
						Group:    g,
						Member:   m,
						Operator: g.FindMember(operator),
					})
				}
			}
		}

	}
	if info.MsgType == 44 {
		data.ReadBytes(5)
		var4 := int32(data.ReadByte())
		var var5 int64 = 0
		target := int64(uint32(data.ReadInt32()))
		if var4 != 0 && var4 != 1 {
			var5 = int64(uint32(data.ReadInt32()))
		}
		if g := c.FindGroupByUin(info.FromUin); g != nil {
			if var5 == 0 && data.Len() == 1 {
				newPermission := func() MemberPermission {
					if data.ReadByte() == 1 {
						return Administrator
					}
					return Member
				}()
				mem := g.FindMember(target)
				if mem.Permission != newPermission {
					old := mem.Permission
					mem.Permission = newPermission
					c.dispatchPermissionChanged(&MemberPermissionChangedEvent{
						Group:         g,
						Member:        mem,
						OldPermission: old,
						NewPermission: newPermission,
					})
				}
			}
		}
	}
	return nil, nil
}

func decodeSystemMsgGroupPacket(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	rsp := structmsg.RspSystemMsgNew{}
	if err := proto.Unmarshal(payload, &rsp); err != nil {
		return nil, err
	}
	if len(rsp.Groupmsgs) == 0 {
		return nil, nil
	}
	st := rsp.Groupmsgs[0]
	if st.Msg != nil {
		// 其他SubType不关心
		if st.Msg.SubType == 1 {
			switch st.Msg.C2CInviteJoinGroupFlag {
			case 0: //成员申请
				c.dispatchJoinGroupRequest(&UserJoinGroupRequest{
					RequestId:     st.MsgSeq,
					Message:       st.Msg.MsgAdditional,
					RequesterUin:  st.ReqUin,
					RequesterNick: st.Msg.ReqUinNick,
					GroupCode:     st.Msg.GroupCode,
					GroupName:     st.Msg.GroupName,
					client:        c,
				})
			case 1: // 被邀请
				c.dispatchGroupInvitedEvent(&GroupInvitedRequest{
					RequestId:   st.MsgSeq,
					InvitorUin:  st.Msg.ActionUin,
					InvitorNick: st.Msg.ActionUinNick,
					GroupCode:   st.Msg.GroupCode,
					GroupName:   st.Msg.GroupName,
					client:      c,
				})
			}
		}
	}
	return nil, nil
}

func decodeSystemMsgFriendPacket(c *QQClient, _ uint16, payload []byte) (interface{}, error) {
	rsp := structmsg.RspSystemMsgNew{}
	if err := proto.Unmarshal(payload, &rsp); err != nil {
		return nil, err
	}
	if len(rsp.Friendmsgs) == 0 {
		return nil, nil
	}
	st := rsp.Friendmsgs[0]
	if st.Msg != nil {
		c.dispatchNewFriendRequest(&NewFriendRequest{
			RequestId:     st.MsgSeq,
			Message:       st.Msg.MsgAdditional,
			RequesterUin:  st.ReqUin,
			RequesterNick: st.Msg.ReqUinNick,
			client:        c,
		})
	}
	return nil, nil
}