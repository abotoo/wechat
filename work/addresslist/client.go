package addresslist

import (
	"github.com/abotoo/wechat/v3/work/context"
)

// Client 通讯录管理接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}
