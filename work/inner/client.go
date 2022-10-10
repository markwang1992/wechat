package inner

// Client 内部开发实例
type Client struct {
	AccessToken string
}

// NewClient 初始化实例
func NewClient(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
	}
}
