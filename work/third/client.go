package third

// Client 第三方应用实例
type Client struct {
	AccessToken string
}

// NewClient 初始化实例
func NewClient(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
	}
}
