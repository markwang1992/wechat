// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/8/11

package third

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// DepartmentSimpleListURL 获取子部门ID列表
	DepartmentSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s&id=%d"
	// UserSimpleListURL 获取部门成员
	UserSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=%s&department_id=%d"
	// UserListIdURL 获取成员ID列表
	UserListIdURL = "https://qyapi.weixin.qq.com/cgi-bin/user/list_id?access_token=%s"
)

type (
	// DepartmentSimpleListResponse 获取子部门ID列表响应
	DepartmentSimpleListResponse struct {
		util.CommonError
		DepartmentId []*DepartmentID `json:"department_id"`
	}
	DepartmentID struct {
		ID       int `json:"id"`
		ParentID int `json:"parentid"`
		Order    int `json:"order"`
	}
)

// DepartmentSimpleList 获取子部门ID列表
// see https://developer.work.weixin.qq.com/document/path/95350
func (r *Client) DepartmentSimpleList(departmentId int) ([]*DepartmentID, error) {
	var (
		response []byte
		err      error
	)
	if response, err = util.HTTPGet(fmt.Sprintf(DepartmentSimpleListURL, r.AccessToken, departmentId)); err != nil {
		return nil, err
	}
	result := &DepartmentSimpleListResponse{}
	if err = util.DecodeWithError(response, result, "DepartmentSimpleList"); err != nil {
		return nil, err
	}
	return result.DepartmentId, nil
}

type (
	UserSimpleListResponse struct {
		util.CommonError
		UserList []*UserList
	}
	UserList struct {
		UserID     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
		OpenUserID string `json:"open_userid"`
	}
)

// UserSimpleList 获取部门成员
// @see https://developer.work.weixin.qq.com/document/path/90200
func (r *Client) UserSimpleList(departmentId int) ([]*UserList, error) {
	var (
		response []byte
		err      error
	)
	if response, err = util.HTTPGet(fmt.Sprintf(UserSimpleListURL, r.AccessToken, departmentId)); err != nil {
		return nil, err
	}
	result := &UserSimpleListResponse{}
	if err = util.DecodeWithError(response, result, "UserSimpleList"); err != nil {
		return nil, err
	}
	return result.UserList, nil
}

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
