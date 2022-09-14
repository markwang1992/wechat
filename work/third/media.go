// @Author markwang <wangyu@uniondrug.cn>
// @Date   2022/9/14

package third

import (
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// UploadImgURL 上传图片
	UploadImgURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"
)

type UploadImgResponse struct {
	util.CommonError
	URL string `json:"url"`
}

// UploadImg 上传图片
// @see https://developer.work.weixin.qq.com/document/path/90392
func (r *Client) UploadImg(filename string) (*UploadImgResponse, error) {
	var (
		err      error
		response []byte
	)
	if response, err = util.PostFile("media", filename, fmt.Sprintf(UploadImgURL, r.AccessToken)); err != nil {
		return nil, err
	}
	result := &UploadImgResponse{}
	if err = util.DecodeWithError(response, result, "UploadImg"); err != nil {
		return nil, err
	}
	return result, nil
}
