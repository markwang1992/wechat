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
