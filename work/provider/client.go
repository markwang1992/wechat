package provider

// Client 服务商凭证相关接口实例
type Client struct {
	AccessToken string
}

// NewClient 初始化实例
func NewClient(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
	}
}
