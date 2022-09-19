// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/12

package third

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// AddContactWayURL 配置客户联系「联系我」方式
	AddContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=%s"
	// GetContactWayURL 获取企业已配置的「联系我」方式
	GetContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way?access_token=%s"
	// ListContactWayURL 获取企业已配置的「联系我」列表
	ListContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list_contact_way?access_token=%s"
	// UpdateContactWayURL 更新企业已配置的「联系我」方式
	UpdateContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way?access_token=%s"
	// DelContactWayURL 删除企业已配置的「联系我」方式
	DelContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way?access_token=%s"
	// GetUserBehaviorDataURL 获取「联系客户统计」数据
	GetUserBehaviorDataURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_user_behavior_data?access_token=%s"
	// GetGroupChatStatURL 获取「群聊数据统计」数据 按群主聚合的方式
	GetGroupChatStatURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic?access_token=%s"
	// GetGroupChatStatByDayURL 获取「群聊数据统计」数据 按自然日聚合的方式
	GetGroupChatStatByDayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/statistic_group_by_day?access_token=%s"
	// FetchExternalContactUserListURL 获取客户列表
	FetchExternalContactUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=%s&userid=%s"
	// FetchExternalContactUserDetailURL 获取客户详情
	FetchExternalContactUserDetailURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s"
	// FetchBatchExternalContactUserDetailURL 批量获取客户详情
	FetchBatchExternalContactUserDetailURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=%s"
)

type (
	// ConclusionsRequest 结束语请求
	ConclusionsRequest struct {
		Text        ConclusionsText         `json:"text"`
		Image       ConclusionsImageRequest `json:"image"`
		Link        ConclusionsLink         `json:"link"`
		MiniProgram ConclusionsMiniProgram  `json:"miniprogram"`
	}
	// ConclusionsText 文本格式结束语
	ConclusionsText struct {
		Content string `json:"content"`
	}
	// ConclusionsImageRequest 图片格式结束语请求
	ConclusionsImageRequest struct {
		MediaID string `json:"media_id"`
	}
	// ConclusionsLink 链接格式结束语
	ConclusionsLink struct {
		Title  string `json:"title"`
		PicUrl string `json:"picurl"`
		Desc   string `json:"desc"`
		URL    string `json:"url"`
	}
	// ConclusionsMiniProgram 小程序格式结束语
	ConclusionsMiniProgram struct {
		Title      string `json:"title"`
		PicMediaID string `json:"pic_media_id"`
		AppID      string `json:"appid"`
		Page       string `json:"page"`
	}
	// ConclusionsResponse 结束语响应
	ConclusionsResponse struct {
		Text        ConclusionsText          `json:"text"`
		Image       ConclusionsImageResponse `json:"image"`
		Link        ConclusionsLink          `json:"link"`
		MiniProgram ConclusionsMiniProgram   `json:"miniprogram"`
	}
	// ConclusionsImageResponse 图片格式结束语响应
	ConclusionsImageResponse struct {
		PicUrl string `json:"pic_url"`
	}
)

type (
	// AddContactWayRequest 配置客户联系「联系我」方式请求
	AddContactWayRequest struct {
		Type          int                `json:"type"`
		Scene         int                `json:"scene"`
		Style         int                `json:"style"`
		Remark        string             `json:"remark"`
		SkipVerify    bool               `json:"skip_verify"`
		State         string             `json:"state"`
		User          []string           `json:"user"`
		Party         []int              `json:"party"`
		IsTemp        bool               `json:"is_temp"`
		ExpiresIn     int                `json:"expires_in"`
		ChatExpiresIn int                `json:"chat_expires_in"`
		UnionID       string             `json:"unionid"`
		Conclusions   ConclusionsRequest `json:"conclusions"`
	}
	// AddContactWayResponse 配置客户联系「联系我」方式响应
	AddContactWayResponse struct {
		util.CommonError
		ConfigID string `json:"config_id"`
		QrCode   string `json:"qr_code"`
	}
)

// AddContactWay 配置客户联系「联系我」方式
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) AddContactWay(req *AddContactWayRequest) (*AddContactWayResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(AddContactWayURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &AddContactWayResponse{}
	if err = util.DecodeWithError(response, result, "AddContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetContactWayRequest 获取企业已配置的「联系我」方式请求
	GetContactWayRequest struct {
		ConfigID string `json:"config_id"`
	}
	// GetContactWayResponse 获取企业已配置的「联系我」方式响应
	GetContactWayResponse struct {
		util.CommonError
		ContactWay ContactWayForGet `json:"contact_way"`
	}
	// ContactWayForGet 「联系我」配置
	ContactWayForGet struct {
		ConfigID      string              `json:"config_id"`
		Type          int                 `json:"type"`
		Scene         int                 `json:"scene"`
		Style         int                 `json:"style"`
		Remark        string              `json:"remark"`
		SkipVerify    bool                `json:"skip_verify"`
		State         string              `json:"state"`
		QrCode        string              `json:"qr_code"`
		User          []string            `json:"user"`
		Party         []int               `json:"party"`
		IsTemp        bool                `json:"is_temp"`
		ExpiresIn     int                 `json:"expires_in"`
		ChatExpiresIn int                 `json:"chat_expires_in"`
		UnionID       string              `json:"unionid"`
		Conclusions   ConclusionsResponse `json:"conclusions"`
	}
)

// GetContactWay 获取企业已配置的「联系我」方式
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) GetContactWay(req *GetContactWayRequest) (*GetContactWayResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(GetContactWayURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &GetContactWayResponse{}
	if err = util.DecodeWithError(response, result, "GetContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	//ListContactWayRequest 获取企业已配置的「联系我」列表请求
	ListContactWayRequest struct {
		StartTime int    `json:"start_time"`
		EndTime   int    `json:"end_time"`
		Cursor    string `json:"cursor"`
		Limit     int    `json:"limit"`
	}
	//ListContactWayResponse 获取企业已配置的「联系我」列表响应
	ListContactWayResponse struct {
		util.CommonError
		ContactWay []*ContactWayForList `json:"contact_way"`
		NextCursor string               `json:"next_cursor"`
	}
	// ContactWayForList 「联系我」配置
	ContactWayForList struct {
		ConfigId string `json:"config_id"`
	}
)

// ListContactWay 获取企业已配置的「联系我」列表
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) ListContactWay(req *ListContactWayRequest) (*ListContactWayResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(ListContactWayURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &ListContactWayResponse{}
	if err = util.DecodeWithError(response, result, "ListContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// UpdateContactWayRequest 更新企业已配置的「联系我」方式请求
	UpdateContactWayRequest struct {
		ConfigID      string             `json:"config_id"`
		Remark        string             `json:"remark"`
		SkipVerify    bool               `json:"skip_verify"`
		Style         int                `json:"style"`
		State         string             `json:"state"`
		User          []string           `json:"user"`
		Party         []int              `json:"party"`
		ExpiresIn     int                `json:"expires_in"`
		ChatExpiresIn int                `json:"chat_expires_in"`
		UnionID       string             `json:"unionid"`
		Conclusions   ConclusionsRequest `json:"conclusions"`
	}
	// UpdateContactWayResponse 更新企业已配置的「联系我」方式响应
	UpdateContactWayResponse struct {
		util.CommonError
	}
)

// UpdateContactWay 更新企业已配置的「联系我」方式
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) UpdateContactWay(req *UpdateContactWayRequest) (*UpdateContactWayResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(UpdateContactWayURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &UpdateContactWayResponse{}
	if err = util.DecodeWithError(response, result, "UpdateContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// DelContactWayRequest 删除企业已配置的「联系我」方式请求
	DelContactWayRequest struct {
		ConfigID string `json:"config_id"`
	}
	// DelContactWayResponse 删除企业已配置的「联系我」方式响应
	DelContactWayResponse struct {
		util.CommonError
	}
)

// DelContactWay 删除企业已配置的「联系我」方式
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) DelContactWay(req *DelContactWayRequest) (*DelContactWayResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(DelContactWayURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &DelContactWayResponse{}
	if err = util.DecodeWithError(response, result, "DelContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetUserBehaviorRequest 获取「联系客户统计」数据请求
	GetUserBehaviorRequest struct {
		UserID    []string `json:"userid"`
		PartyID   []int    `json:"partyid"`
		StartTime int      `json:"start_time"`
		EndTime   int      `json:"end_time"`
	}
	// GetUserBehaviorResponse 获取「联系客户统计」数据响应
	GetUserBehaviorResponse struct {
		util.CommonError
		BehaviorData []BehaviorData `json:"behavior_data"`
	}
	// BehaviorData 联系客户统计数据
	BehaviorData struct {
		StatTime            int     `json:"stat_time"`
		ChatCnt             int     `json:"chat_cnt"`
		MessageCnt          int     `json:"message_cnt"`
		ReplyPercentage     float64 `json:"reply_percentage"`
		AvgReplyTime        int     `json:"avg_reply_time"`
		NegativeFeedbackCnt int     `json:"negative_feedback_cnt"`
		NewApplyCnt         int     `json:"new_apply_cnt"`
		NewContactCnt       int     `json:"new_contact_cnt"`
	}
)

// GetUserBehaviorData 获取「联系客户统计」数据
// @see https://developer.work.weixin.qq.com/document/path/92132
func (r *Client) GetUserBehaviorData(req *GetUserBehaviorRequest) ([]BehaviorData, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(GetUserBehaviorDataURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	var result GetUserBehaviorResponse
	if err = util.DecodeWithError(response, &result, "GetUserBehaviorData"); err != nil {
		return nil, err
	}
	return result.BehaviorData, nil
}

type (
	// GetGroupChatStatRequest 获取「群聊数据统计」数据 按群主聚合的方式 请求
	GetGroupChatStatRequest struct {
		DayBeginTime int         `json:"day_begin_time"`
		DayEndTime   int         `json:"day_end_time"`
		OwnerFilter  OwnerFilter `json:"owner_filter"`
		OrderBy      int         `json:"order_by"`
		OrderAsc     int         `json:"order_asc"`
		Offset       int         `json:"offset"`
		Limit        int         `json:"limit"`
	}
	// GetGroupChatStatResponse 获取「群聊数据统计」数据 按群主聚合的方式 响应
	GetGroupChatStatResponse struct {
		util.CommonError
		Total      int                 `json:"total"`
		NextOffset int                 `json:"next_offset"`
		Items      []GroupChatStatItem `json:"items"`
	}
	// GroupChatStatItem 群聊数据统计(按群主聚合)条目
	GroupChatStatItem struct {
		Owner string                `json:"owner"`
		Data  GroupChatStatItemData `json:"data"`
	}
)

// OwnerFilter 群主过滤
type OwnerFilter struct {
	UseridList []string `json:"userid_list"`
}

// GroupChatStatItemData 群聊数据统计条目数据
type GroupChatStatItemData struct {
	NewChatCnt            int `json:"new_chat_cnt"`
	ChatTotal             int `json:"chat_total"`
	ChatHasMsg            int `json:"chat_has_msg"`
	NewMemberCnt          int `json:"new_member_cnt"`
	MemberTotal           int `json:"member_total"`
	MemberHasMsg          int `json:"member_has_msg"`
	MsgTotal              int `json:"msg_total"`
	MigrateTraineeChatCnt int `json:"migrate_trainee_chat_cnt"`
}

// GetGroupChatStat 获取「群聊数据统计」数据 按群主聚合的方式
// @see https://developer.work.weixin.qq.com/document/path/92133
func (r *Client) GetGroupChatStat(req *GetGroupChatStatRequest) (*GetGroupChatStatResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(GetGroupChatStatURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &GetGroupChatStatResponse{}
	if err = util.DecodeWithError(response, result, "GetGroupChatStat"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetGroupChatStatByDayRequest 获取「群聊数据统计」数据 按自然日聚合的方式 请求
	GetGroupChatStatByDayRequest struct {
		DayBeginTime int         `json:"day_begin_time"`
		DayEndTime   int         `json:"day_end_time"`
		OwnerFilter  OwnerFilter `json:"owner_filter"`
	}
	// GetGroupChatStatByDayResponse 获取「群聊数据统计」数据 按自然日聚合的方式 响应
	GetGroupChatStatByDayResponse struct {
		util.CommonError
		Items []GetGroupChatStatByDayItem `json:"items"`
	}
	// GetGroupChatStatByDayItem 群聊数据统计(按自然日聚合)条目
	GetGroupChatStatByDayItem struct {
		StatTime int                   `json:"stat_time"`
		Data     GroupChatStatItemData `json:"data"`
	}
)

// GetGroupChatStatByDay 获取「群聊数据统计」数据 按自然日聚合的方式
// @see https://developer.work.weixin.qq.com/document/path/92133
func (r *Client) GetGroupChatStatByDay(req *GetGroupChatStatByDayRequest) ([]GetGroupChatStatByDayItem, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(GetGroupChatStatByDayURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	var result GetGroupChatStatByDayResponse
	if err = util.DecodeWithError(response, &result, "GetGroupChatStatByDay"); err != nil {
		return nil, err
	}
	return result.Items, nil
}

// ExternalUserListResponse 外部联系人列表响应
type ExternalUserListResponse struct {
	util.CommonError
	ExternalUserID []string `json:"external_userid"`
}

// GetExternalUserList 获取客户列表
// @see https://developer.work.weixin.qq.com/document/path/92113
func (r *Client) GetExternalUserList(userID string) ([]string, error) {
	var (
		response []byte
		err      error
	)
	if response, err = util.HTTPGet(fmt.Sprintf(FetchExternalContactUserListURL, r.AccessToken, userID)); err != nil {
		return nil, err
	}
	result := &ExternalUserListResponse{}
	if err = util.DecodeWithError(response, result, "GetExternalUserList"); err != nil {
		return nil, err
	}
	return result.ExternalUserID, nil
}

// ExternalUserDetailResponse 外部联系人详情响应
type ExternalUserDetailResponse struct {
	util.CommonError
	ExternalUser
}

// ExternalUser 外部联系人
type ExternalUser struct {
	ExternalUserID  string       `json:"external_userid"`
	Name            string       `json:"name"`
	Avatar          string       `json:"avatar"`
	Type            int64        `json:"type"`
	Gender          int64        `json:"gender"`
	UnionID         string       `json:"unionid"`
	Position        string       `json:"position"`
	CorpName        string       `json:"corp_name"`
	CorpFullName    string       `json:"corp_full_name"`
	ExternalProfile string       `json:"external_profile"`
	FollowUser      []FollowUser `json:"follow_user"`
	NextCursor      string       `json:"next_cursor"`
}

// FollowUser 跟进用户（指企业内部用户）
type FollowUser struct {
	UserID         string        `json:"userid"`
	Remark         string        `json:"remark"`
	Description    string        `json:"description"`
	CreateTime     string        `json:"create_time"`
	Tags           []Tag         `json:"tags"`
	RemarkCorpName string        `json:"remark_corp_name"`
	RemarkMobiles  []string      `json:"remark_mobiles"`
	OperUserID     string        `json:"oper_userid"`
	AddWay         int64         `json:"add_way"`
	WeChatChannels WechatChannel `json:"wechat_channels"`
	State          string        `json:"state"`
}

// Tag 已绑定在外部联系人的标签
type Tag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	Type      int64  `json:"type"`
	TagID     string `json:"tag_id"`
}

// WechatChannel 视频号添加的场景
type WechatChannel struct {
	NickName string `json:"nickname"`
	Source   string `json:"source"`
}

// GetExternalUserDetail 获取外部联系人详情
// @see https://developer.work.weixin.qq.com/document/path/92265
func (r *Client) GetExternalUserDetail(externalUserID string) (*ExternalUser, error) {
	var (
		response []byte
		err      error
	)
	if response, err = util.HTTPGet(fmt.Sprintf(FetchExternalContactUserDetailURL, r.AccessToken, externalUserID)); err != nil {
		return nil, err
	}
	var result ExternalUserDetailResponse
	if err = util.DecodeWithError(response, &result, "GetExternalUserDetail"); err != nil {
		return nil, err
	}
	return &result.ExternalUser, nil
}

// BatchGetExternalUserDetailsRequest 批量获取外部联系人详情请求
type BatchGetExternalUserDetailsRequest struct {
	UserIDList []string `json:"userid_list"`
	Cursor     string   `json:"cursor"`
	Limit      int      `json:"limit"`
}

// ExternalUserDetailListResponse 批量获取外部联系人详情响应
type ExternalUserDetailListResponse struct {
	util.CommonError
	ExternalContactList []ExternalUser `json:"external_contact_list"`
}

// BatchGetExternalUserDetails 批量获取外部联系人详情
func (r *Client) BatchGetExternalUserDetails(req *BatchGetExternalUserDetailsRequest) (*ExternalUserDetailListResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(FetchBatchExternalContactUserDetailURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &ExternalUserDetailListResponse{}
	if err = util.DecodeWithError(response, result, "BatchGetExternalUserDetails"); err != nil {
		return nil, err
	}
	return result, nil
}
