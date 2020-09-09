package service

import (
	"strings"

	"github.com/BinacsLee/server/config"
	"github.com/BinacsLee/server/service/db"
	"github.com/BinacsLee/server/types"

	"github.com/binacsgo/log"
	"github.com/binacsgo/tinyurl"
)

// WebTinyURLServiceImpl Web TinyURL service implement
type WebTinyURLServiceImpl struct {
	Config *config.Config `inject-name:"Config"`
	Logger log.Logger     `inject-name:"WebLogger"`

	RedisSvc db.RedisService `inject-name:"RedisService"`
	MysqlSvc db.MysqlService `inject-name:"MysqlService"`

	instance tinyurl.TinyURL
}

// AfterInject do inject
func (ts *WebTinyURLServiceImpl) AfterInject() error {
	// select a impl of TinyURL interface
	// choose md5 for now
	ts.instance = tinyurl.GetMD5Impl()
	return nil
}

// Encode return the encode string directly
func (ts *WebTinyURLServiceImpl) Encode(s string) string {
	return ts.instance.EncodeURL(s)
}

// URLEncode return the tiny-url of origin-url
func (ts *WebTinyURLServiceImpl) URLEncode(url string) (string, error) {
	ts.Logger.Info("WebTinyURLServiceImpl URLEncode Start", "url", url)
	if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
		ts.Logger.Error("WebTinyURLServiceImpl URLEncode, url illegal", "text", url)
		return "url illegal, http/https as prefix!", nil
	}
	turl := ts.instance.EncodeURL(url)
	err := ts.RedisSvc.Set(turl, url, types.TinyURLExpireDuration())
	if err != nil {
		ts.Logger.Error("WebTinyURLServiceImpl URLEncode RedisSvc.Set", "error", err, "url", url, "turl", turl)
		return "", err
	}
	return ts.Config.WebConfig.Host + "/r/" + turl, nil
}

// URLDecode return the origin-url of tiny-url
func (ts *WebTinyURLServiceImpl) URLDecode(text string) (string, error) {
	ts.Logger.Info("WebTinyURLServiceImpl URLDecode Start", "tiny-url", text)
	if !strings.HasPrefix(text, ts.Config.WebConfig.Host+"/r/") {
		ts.Logger.Error("WebTinyURLServiceImpl URLDecode, tiny-url illegal", "text", text)
		return "tiny-url illegal", nil
	}
	turl := strings.TrimPrefix(text, ts.Config.WebConfig.Host+"/r/")
	url, err := ts.RedisSvc.Get(turl)
	if err != nil {
		ts.Logger.Error("WebTinyURLServiceImpl URLDecode RedisSvc.Get", "error", err, "url", url, "turl", turl)
		return "", err
	}
	return url, nil
}

// URLSearch return the origin-url of tiny-url suffix
func (ts *WebTinyURLServiceImpl) URLSearch(turl string) (string, error) {
	ts.Logger.Info("WebTinyURLServiceImpl URLSearch Start", "tiny-url suffix", turl)
	url, err := ts.RedisSvc.Get(turl)
	if err != nil {
		ts.Logger.Error("WebTinyURLServiceImpl URLSearch RedisSvc.Get", "error", err, "url", url, "turl", turl)
		return "", err
	}
	return url, nil
}
