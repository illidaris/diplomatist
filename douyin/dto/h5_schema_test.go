package dto

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestH5Schema(t *testing.T) {
	convey.Convey("TestH5Schema", t, func() {
		schema := H5Schema{
			NonceStr:  "Wm3WZYTPz0wzccnW",
			Timestamp: 1650941858,
			Ticket:    "@ml6sqYBGgTKmQNajnKNkaj8yksCAY++adIhlGIqfTiKyvBqOIkzdJ6WRgP+nO+wtVItqKbX4iZ+mFIYkyPJjpQ==",
		}

		convey.Convey("test GetSign success", func() {
			sign := schema.GetSign()
			right := "3f7b739a91a52cb7d85c4f89c5f611fe"
			convey.So(sign, convey.ShouldEqual, right)
		})
		convey.Convey("test GetScheme success", func() {
			scm := schema.GetScheme()
			right := "snssdk1128://openplatform/share?nonce_str=Wm3WZYTPz0wzccnW&signature=3f7b739a91a52cb7d85c4f89c5f611fe&timestamp=1650941858"
			convey.So(scm, convey.ShouldEqual, right)
		})
	})

}
