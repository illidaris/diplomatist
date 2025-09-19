package dto

import "testing"

func TestGetSign(t *testing.T) {
	b := &BaseRequest{
		AppId:     "56721",
		AppSecret: "secret",
		Nonce:     "1234",
		Ts:        1722863913108,
		Body:      `{"key":"value"}`,
	}
	right := "+gZLQ2PN5xezyDHVcjNVQWE4qQVEc+6pLQ4uOEFsfvs="
	v := b.GetSigns("POST", "webcast/gamecp/attribute/channel/upload_test_payment_summary", nil)
	println(v)
	if b.Sign != right {
		t.Error("faild")
	}
}
