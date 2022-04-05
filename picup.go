package picup

import "github.com/c3b2a7/picup/apis"

type PicUp struct {
	apis []apis.API
}

type UpResult struct {
}

type Interceptor interface {
	intercept()
}
