// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/10/10

package inner

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// UserListIdURL 获取成员ID列表
	UserListIdURL = "https://qyapi.weixin.qq.com/cgi-bin/user/list_id?access_token=%s"
)

// UserListIdRequest 获取成员ID列表请求
type UserListIdRequest struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

// UserListIdResponse 获取成员ID列表响应
type UserListIdResponse struct {
	util.CommonError
	NextCursor string      `json:"next_cursor"`
	DeptUser   []*DeptUser `json:"dept_user"`
}

// DeptUser 用户-部门关系
type DeptUser struct {
	UserID     string `json:"userid"`
	Department int    `json:"department"`
}

// UserListId 获取成员ID列表
// see https://developer.work.weixin.qq.com/document/path/96067
func (r *Client) UserListId(req *UserListIdRequest) (*UserListIdResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostJSON(fmt.Sprintf(UserListIdURL, r.AccessToken), req); err != nil {
		return nil, err
	}
	result := &UserListIdResponse{}
	if err = util.DecodeWithError(response, result, "UserListId"); err != nil {
		return nil, err
	}
	return result, nil
}
