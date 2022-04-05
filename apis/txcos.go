package apis

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

type TxCos struct {
	SECRETID, SECRETKEY, REGIONID, BUCKETNAME string
}

func (txCos *TxCos) Up(path string) (string, error) {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", txCos.BUCKETNAME, txCos.REGIONID))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  txCos.SECRETID,
			SecretKey: txCos.SECRETKEY,
		},
	})

	key := "exampleobject"
	file := "test"

	result, _, err := client.Object.Upload(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return result.Location, nil
}
