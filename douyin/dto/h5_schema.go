package dto

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/illidaris/aphrodite/pkg/convert"
	"github.com/illidaris/diplomatist/douyin/vao"
	"github.com/spf13/cast"
)

func NewH5Schema() H5Schema {
	return H5Schema{
		ShareType: vao.SHARE_TYPE_H5,
		NonceStr:  randomString(12),
		Timestamp: time.Now().Unix(),
	}
}

type H5Schema struct {
	Ticket           string   `json:"ticket" form:"ticket" url:"ticket,omitempty"`                                     // 抖音返回的 ticket
	ShareType        string   `json:"share_type" form:"share_type" url:"share_type,omitempty"`                         // 固定值为 h5
	ClientKey        string   `json:"client_key" form:"client_key" url:"client_key,omitempty"`                         // 应用唯一标识
	NonceStr         string   `json:"nonce_str" form:"nonce_str" url:"nonce_str,omitempty"`                            // 随机字符串
	Timestamp        int64    `json:"timestamp" form:"timestamp" url:"timestamp,omitempty"`                            // 时间戳
	Signature        string   `json:"signature" form:"signature" url:"-"`                                              // 签名
	State            string   `json:"state" form:"state" url:"state,omitempty"`                                        // 建议填写, 按照查询视频发布结果获取 share_id，可以获取视频发布情况
	ImagePath        string   `json:"image_path" form:"image_path" url:"image_path,omitempty"`                         // 图片地址（单个，不能超过 20M）当 video_path 存在时优先使用 video_path 当前支持的格式包含 png/jpg/gif关于该参数请阅读表格下面的注意事项。
	ImageListPath    string   `json:"image_list_path" form:"image_list_path" url:"image_list_path,omitempty"`          // 图片文件路径（多个），图集模式分享。示例：'["https://douyin.com"]' 当 video_path 存在时优先使用 video_path 当前支持的格式包含 png/jpg。注意抖音 22.2.0 版本以上支持该参数。关于该参数请阅读表格下面的注意事项。
	VideoPath        string   `json:"video_path" form:"video_path" url:"video_path,omitempty"`                         // 视频文件路径（单个，不能超过 128M)。当前支持的格式包含 mp4/mov。关于该参数请阅读表格下面的注意事项。
	HashTagList      []string `json:"hash_tag_list" form:"hash_tag_list" url:"hash_tag_list,omitempty"`                // 支持有第三方预设内容分享抖音时默认携带的话题，指定的话题会展现在发布页面。示例：'["hashtag1","hashtag2"]' 用户可自行删除该话题，该话题类型支持商业化话题和普通话题。发布后和抖音原生话题没有差别。
	MicroAppInfo     any      `json:"micro_app_info" form:"micro_app_info" url:"micro_app_info,omitempty"`             // 添加小程序。视频成功发布视频后，在视频左下角带有小程序入口。示例：'{"appId":"小程序appId","appTitle":"小程序标题","description":"小程序描述语","appUrl":"小程序中生成该页面时写的 path 地址"}'
	ShareToPublish   int      `json:"share_to_publish" form:"share_to_publish" url:"share_to_publish,omitempty"`       // 为 1 时直接分享到抖音发布页（仅视频）
	Title            string   `json:"title" form:"title" url:"title,omitempty"`                                        // 视频标题
	TitleHashTagList []any    `json:"title_hashtag_list" form:"title_hashtag_list" url:"title_hashtag_list,omitempty"` // 插入title的hashtag列表，可指定在title中位置。
	PoiId            string   `json:"poi_id" form:"poi_id" url:"poi_id,omitempty"`                                     // 地理位置信息锚点 id，与小程序 appId 互斥，优先展示小程序。
	ShareToType      int      `json:"share_to_type" form:"share_to_type" url:"share_to_type,omitempty"`                // 0：投稿 1：转发到日常
	ShortTitle       string   `json:"short_title" form:"short_title" url:"short_title,omitempty"`                      // 短标题
	DownloadType     int      `json:"download_type" form:"download_type" url:"download_type,omitempty"`                // 是否允许下载 1：允许，2：不允许
	PrivateStatus    int      `json:"private_status" form:"private_status" url:"private_status,omitempty"`             // 视频公开范围 0：全部人可见，1：自己可见，2：好友可见
	Feature          string   `json:"feature" form:"feature" url:"feature,omitempty"`                                  // 可选值： note：当整体投稿流程为"普通投稿"，且投稿内容为"多图"，投稿类型为“图文”时，可以设置这个值，将体裁转变为"笔记"。
}

func (r H5Schema) GetScheme() string {
	u, err := query.Values(r)
	if err != nil {
		return ""
	}
	sign := r.GetSign()
	u.Add("signature", sign)
	u.Del("ticket")
	raw := u.Encode()
	return fmt.Sprintf("%s?%s", vao.H5SCHEMA, raw)
}

func (r H5Schema) GetSign() string {
	u := url.Values{}
	u.Add("ticket", r.Ticket)
	u.Add("timestamp", cast.ToString(r.Timestamp))
	u.Add("nonce_str", r.NonceStr)

	raw := convert.ValuesRawEncode(u)
	bs := md5.Sum([]byte(raw))
	return fmt.Sprintf("%x", bs)
}

func randomString(length int) string {
	b := make([]byte, (length+2)/3*4)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	// 使用base64编码，这样可以包含URL安全的字符和数字，并且长度会是原始长度的4/3（向上取整）
	return base64.URLEncoding.EncodeToString(b)[:length]
}
