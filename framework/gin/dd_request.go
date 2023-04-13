package gin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"mime/multipart"

	"github.com/spf13/cast"
)

//const defaultMultipartMemory = 32 << 20 // 32 MB
var _ = IRequest(&Context{})

// 代表请求包含的方法
type IRequest interface {
	// 请求地址url中带的参数
	// 形如: foo.com?a=1&b=bar&c[]=bar
	DefaultQueryInt(key string, def int) (int, bool)
	DefaultQueryInt64(key string, def int64) (int64, bool)
	DefaultQueryFloat64(key string, def float64) (float64, bool)
	DefaultQueryFloat32(key string, def float32) (float32, bool)
	DefaultQueryBool(key string, def bool) (bool, bool)
	DefaultQueryString(key string, def string) (string, bool)
	DefaultQueryStringSlice(key string, def []string) ([]string, bool)

	// 路由匹配中带的参数
	// 形如 /book/:id
	DefaultParamInt(key string, def int) (int, bool)
	DefaultParamInt64(key string, def int64) (int64, bool)
	DefaultParamFloat64(key string, def float64) (float64, bool)
	DefaultParamFloat32(key string, def float32) (float32, bool)
	DefaultParamBool(key string, def bool) (bool, bool)
	DefaultParamString(key string, def string) (string, bool)
	DefaultParam(key string) interface{}

	// form表单中带的参数
	DefaultFormInt(key string, def int) (int, bool)
	DefaultFormInt64(key string, def int64) (int64, bool)
	DefaultFormFloat64(key string, def float64) (float64, bool)
	DefaultFormFloat32(key string, def float32) (float32, bool)
	DefaultFormBool(key string, def bool) (bool, bool)
	DefaultFormString(key string, def string) (string, bool)
	DefaultFormStringSlice(key string, def []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	// json body
	BindJson(obj interface{}) error

	// xml body
	BindXml(obj interface{}) error

	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	GetHeaders() map[string][]string
	GetHeader(key string) string

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, error)
}

// 获取请求地址中所有参数
func (c *Context) QueryAll() map[string][]string {
	c.initQueryCache()
	return c.queryCache
}

// 请求地址url中带的参数
// 形如: foo.com?a=1&b=bar&c[]=bar

// 获取Int类型的请求参数
func (c *Context) DefaultQueryInt(key string, def int) (int, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			// 使用cast库将string转换为Int
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryInt64(key string, def int64) (int64, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryFloat64(key string, def float64) (float64, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryFloat32(key string, def float32) (float32, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryBool(key string, def bool) (bool, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryString(key string, def string) (string, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (c *Context) DefaultQueryStringSlice(key string, def []string) ([]string, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

// 路由匹配中带的参数
// 形如 /book/:id
func (c *Context) DefaultParamInt(key string, def int) (int, bool) {
	if val := c.Param(key); val != "" {
		// 通过cast进行类型转换
		return cast.ToInt(val), true
	}
	return def, false
}

func (c *Context) DefaultParamInt64(key string, def int64) (int64, bool) {
	if val := c.Param(key); val != "" {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (c *Context) DefaultParamFloat64(key string, def float64) (float64, bool) {
	if val := c.Param(key); val != "" {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (c *Context) DefaultParamFloat32(key string, def float32) (float32, bool) {
	if val := c.Param(key); val != "" {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (c *Context) DefaultParamBool(key string, def bool) (bool, bool) {
	if val := c.Param(key); val != "" {
		return cast.ToBool(val), true
	}
	return def, false
}

func (c *Context) DefaultParamString(key string, def string) (string, bool) {
	if val := c.Param(key); val != "" {
		return cast.ToString(val), true
	}
	return def, false
}

// 获取路由参数
func (c *Context) DefaultParam(key string) interface{} {
	if c.params != nil {
		for _, p := range *c.params {
			if p.Key == key {
				return p.Value
			}
		}
	}
	return nil
}

func (c *Context) FormAll() map[string][]string {
	if c.Request != nil {
		c.Request.ParseForm()
		return c.Request.PostForm
	}
	return map[string][]string{}
}

func (c *Context) DefaultFormInt(key string, def int) (int, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormInt64(key string, def int64) (int64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormFloat64(key string, def float64) (float64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormFloat32(key string, def float32) (float32, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormBool(key string, def bool) (bool, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) DefaultFormString(key string, def string) (string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0], true
	}
	return def, false
}

func (c *Context) DefaultFormStringSlice(key string, def []string) ([]string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}
func (c *Context) DefaultFormFile(key string) (*multipart.FileHeader, error) {
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(defaultMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.Request.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}
func (c *Context) DefaultForm(key string) interface{} {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

// 将body文本解析到obj结构体中
func (c *Context) BindJson(obj interface{}) error {
	if c.Request != nil {
		// 读取文本
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		// 重新填充request.Body，为后续的逻辑二次读取做准备
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 解析到obj结构体中
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("c.request empty")
	}
	return nil
}

// xml body
func (c *Context) BindXml(obj interface{}) error {
	if c.Request != nil {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = xml.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("c.request empty")
	}
	return nil
}

// 基础信息
func (c *Context) Uri() string {
	return c.Request.RequestURI
}

func (c *Context) Method() string {
	return c.Request.Method
}

func (c *Context) Host() string {
	return c.Request.URL.Host
}

func (c *Context) ClientIp() string {
	r := c.Request
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

// header
func (c *Context) GetHeaders() map[string][]string {
	return c.Request.Header
}

// cookie
func (c *Context) Cookies() map[string]string {
	cookies := c.Request.Cookies()
	ret := map[string]string{}
	for _, cookie := range cookies {
		ret[cookie.Name] = cookie.Value
	}
	return ret
}
