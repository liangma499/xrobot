package backblazeAvatar

import (
	"bytes"
	"context"

	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"tron_robot/internal/code"
	optionbaseconfig "tron_robot/internal/option/option-base-config"
	"tron_robot/internal/utils/xfile"
	"tron_robot/internal/utils/xhash"
	"tron_robot/internal/utils/xstr"
	"tron_robot/internal/xtelegram/telegram/types"
	"tron_robot/internal/xtypes"
	"xbase/errors"
	"xbase/etc"
	"xbase/log"
)

var (
	once     sync.Once
	instance *Avatar
)

type Config struct {
	BasePath    string `json:"basePath"`
	BasePathUrl string `json:"basePathUrl"`
}

func Instance() *Avatar {
	once.Do(func() {
		instance = NewInstance("etc.backblaze.avatar")
	})

	return instance
}
func (c *Config) makeBasePath() error {
	if !c.fileIsExisted(c.BasePath) {
		if err := os.MkdirAll(c.BasePath, 0777); err != nil { //os.ModePerm
			return err
		}
	}
	return nil
}

func (c *Config) fileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *Avatar {
	var (
		conf *Config
		v    any = config
	)

	switch c := v.(type) {
	case string:
		conf = &Config{}
		if err := etc.Get(c).Scan(conf); err != nil {
			log.Fatalf("load backblaze b2 config failed: %v", err)
		}
	case Config:
		conf = &c
	case *Config:
		conf = c
	}

	if err := conf.makeBasePath(); err != nil {
		log.Fatalf("load backblaze b2 config failed: %v", err)
	}

	return &Avatar{conf: conf}
}

type Avatar struct {
	conf *Config
}

// BaseUrl 获取基础链接
func (b *Avatar) BaseUrl() string {
	return strings.TrimSuffix(optionbaseconfig.GetValue(xtypes.AvatarUrlKey), "/")
}

// PutObject 上传对象
func (b *Avatar) PutObject(ctx context.Context, file any) (*Result, error) {
	var (
		body     io.Reader
		filename string
	)

	switch f := file.(type) {
	case string:
		ff, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		defer ff.Close()

		body = ff
		filename = ff.Name()
	case *os.File:
		body = f
		filename = f.Name()
	case *multipart.FileHeader:
		ff, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer ff.Close()

		body = ff
		filename = f.Filename
	default:
		fmt.Println(f)
		return nil, errors.New("invalid file type")
	}

	name := xstr.SerialNO() + filepath.Ext(filename)
	//path := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePath, "/"), "/") + "/" + name
	path := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePath, "/"), "/") + "/" + name
	// 创建目标文件
	fileData, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create dest file failed: %s", err.Error())
	}
	defer fileData.Close()

	// copy内容到目标文件
	_, err = io.Copy(fileData, body)
	if err != nil {
		return nil, fmt.Errorf("copy file failed: %s", err.Error())
	}

	staticPath := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePathUrl, "/"), "/") + "/" + name

	return &Result{
		baseUrl: b.BaseUrl(),
		name:    name,
		path:    staticPath,
	}, nil

}

// DeleteObject 删除对象
func (b *Avatar) DeleteObject(ctx context.Context, file string) error {
	path := strings.TrimSuffix(b.conf.BasePath, "/") + "/" + strings.TrimPrefix(file, "/")

	return os.Remove(path)
}

type Result struct {
	baseUrl string
	name    string
	path    string
	mime    string
}

func (o *Result) Name() string {
	return o.name
}

func (o *Result) Path() string {
	return o.path
}
func (o *Result) Mime() string {
	return o.mime
}
func (o *Result) BaseUrl() string {
	return o.baseUrl
}
func (o *Result) Url() string {
	return strings.TrimSuffix(o.baseUrl, "/") + "/" + strings.TrimPrefix(o.path, "/")
}

// Upload 上传文件
func (b *Avatar) PutFileBase64(ctx context.Context, file string) (*Result, error) {
	if len(file) < 11 || (file[11] != 'p' && file[11] != 'j') {
		return nil, nil
	}
	dataArr := strings.Split(file, ",")
	if len(dataArr) != 2 {
		return nil, nil
	}
	content, err := base64.StdEncoding.DecodeString(dataArr[1])
	if err != nil {
		return nil, err
	}
	ext := "." + strings.Split(strings.Split(dataArr[0], ";")[0], "/")[1]
	name := xstr.SerialNO() + ext
	path := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePath, "/"), "/") + "/" + name
	body := bytes.NewReader(content)
	// 创建目标文件
	fileData, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create dest file failed: %s", err.Error())
	}
	defer fileData.Close()

	// copy内容到目标文件
	_, err = io.Copy(fileData, body)
	if err != nil {
		return nil, fmt.Errorf("copy file failed: %s", err.Error())
	}
	staticPath := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePathUrl, "/"), "/") + "/" + name

	return &Result{
		baseUrl: b.BaseUrl(),
		name:    name,
		path:    staticPath,
	}, nil
}

// Upload 上传文件
func (b *Avatar) PutFile(file *multipart.FileHeader) (*Result, error) {

	mime := file.Header.Get("Content-Type")
	// 读取文件、文件后缀
	_, ext := xfile.GetFileAndExt(file.Filename)
	// 转换后的文件名
	name := xstr.SerialNO() + ext
	path := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePath, "/"), "/") + "/" + name

	// 读取文件内容
	f, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("open file failed: %s", err.Error())
	}
	defer f.Close()

	// 创建目标文件
	fileData, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create dest file failed: %s", err.Error())
	}
	defer fileData.Close()

	// copy内容到目标文件
	_, err = io.Copy(fileData, f)
	if err != nil {
		return nil, fmt.Errorf("copy file failed: %s", err.Error())
	}

	staticPath := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePathUrl, "/"), "/") + "/" + name

	return &Result{
		baseUrl: b.BaseUrl(),
		name:    name,
		path:    staticPath,
		mime:    mime,
	}, nil
}
func (b *Avatar) fileLink(file *types.File, token string) string {
	return "https://api.telegram.org/file/bot" + token + "/" + file.FilePath
}

// https://api.telegram.org/file/bot6867997452:AAFYZXHAC_TDvcfBiYto2ShutRSiUcboa04/photos/file_4.jpg
// 拉取TG头像
func (b *Avatar) PutFileTgAvarl(file *types.File, token string) (*Result, error) {

	if file == nil {
		return nil, errors.NewError(code.NotFound)
	}
	if file.FilePath == "" {
		return nil, errors.NewError(code.NotFound)
	}

	url := b.fileLink(file, token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.NewError(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.NewError(code.NotFound)
	}

	fileSplilt := strings.Split(file.FilePath, ".")
	if len(fileSplilt) < 2 {
		return nil, errors.NewError(code.NotFound)
	}
	// 转换后的文件名
	name := xhash.MD5(string(data)) + "." + fileSplilt[len(fileSplilt)-1]
	//name := "file_1.jpg"
	path := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePath, "/"), "/") + "/" + name

	// 创建目标文件
	fileData, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("create dest file failed: %s", err.Error())
	}
	defer fileData.Close()

	// copy内容到目标文件
	_, err = fileData.Write(data)
	//_, err = io.Copy(fileData, copyData.Write())
	if err != nil {
		return nil, fmt.Errorf("copy file failed: %s", err.Error())
	}

	staticPath := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePathUrl, "/"), "/") + "/" + name

	return &Result{
		baseUrl: b.BaseUrl(),
		name:    name,
		path:    staticPath,
	}, nil
}
