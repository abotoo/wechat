package context

import (
	"github.com/abotoo/wechat/v3/credential"
	"github.com/abotoo/wechat/v3/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenContextHandle
}
