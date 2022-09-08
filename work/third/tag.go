// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/9/6

package third

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// GetCropTagURL 获取企业标签库
	GetCropTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token=%s"
	// AddCropTagURL 添加企业客户标签
	AddCropTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%s"
	// EditCropTagURL 编辑企业客户标签
	EditCropTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=%s"
	// DelCropTagURL 删除企业客户标签
	DelCropTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=%s"
)

// GetCropTagRequest 获取企业标签库请求
type GetCropTagRequest struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
}

// GetCropTagListResponse 获取企业标签库响应
type GetCropTagListResponse struct {
	util.CommonError
	TagGroup []TagGroup `json:"tag_group"`
}

// TagGroup 企业标签组
type TagGroup struct {
	GroupID    string            `json:"group_id"`
	GroupName  string            `json:"group_name"`
	CreateTime int               `json:"create_time"`
	GroupOrder int               `json:"group_order"`
	Deleted    bool              `json:"deleted"`
	Tag        []TagGroupTagItem `json:"tag"`
}

// TagGroupTagItem 企业标签内的子项
type TagGroupTagItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

// GetCropTagList 获取企业标签库
// @see https://developer.work.weixin.qq.com/document/path/92696
func (r *Client) GetCropTagList(req *GetCropTagRequest) (*GetCropTagListResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(GetCropTagURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &GetCropTagListResponse{}
	if err = util.DecodeWithError(response, result, "GetCropTagList"); err != nil {
		return nil, err
	}
	return result, nil
}

// AddCropTagRequest 添加企业客户标签请求
type AddCropTagRequest struct {
	GroupID   string           `json:"group_id,omitempty"`
	GroupName string           `json:"group_name"`
	Order     int              `json:"order"`
	Tag       []AddCropTagItem `json:"tag"`
	AgentID   int              `json:"agentid"`
}

// AddCropTagItem 添加企业客户标签子项
type AddCropTagItem struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

// AddCropTagResponse 添加企业客户标签响应
type AddCropTagResponse struct {
	util.CommonError
	TagGroup TagGroup `json:"tag_group"`
}

// AddCropTag 添加企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92696
func (r *Client) AddCropTag(req *AddCropTagRequest) (*AddCropTagResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(AddCropTagURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &AddCropTagResponse{}
	if err = util.DecodeWithError(response, result, "AddCropTag"); err != nil {
		return nil, err
	}
	return result, nil
}

// EditCropTagRequest 编辑客户企业标签请求
type EditCropTagRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	AgentID string `json:"agent_id"`
}

// EditCropTag 编辑企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92696
func (r *Client) EditCropTag(req *EditCropTagRequest) error {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(EditCropTagURL, r.AccessToken), req); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "EditCropTag")
}

// DeleteCropTagRequest 删除企业客户标签请求
type DeleteCropTagRequest struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
	AgentID string   `json:"agent_id"`
}

// DeleteCropTag 删除企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92696
func (r *Client) DeleteCropTag(req *DeleteCropTagRequest) error {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(DelCropTagURL, r.AccessToken), req); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "DeleteCropTag")
}
