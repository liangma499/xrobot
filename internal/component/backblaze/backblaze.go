package backblaze

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
	"xbase/errors"
	"xbase/etc"
	"xbase/log"
	optionbaseconfig "xrobot/internal/option/option-base-config"
	"xrobot/internal/utils/xstr"
	"xrobot/internal/xtypes"

	"github.com/Backblaze/blazer/b2"
)

var (
	once     sync.Once
	instance *BackBlaze
)

type Config struct {
	Account  string `json:"account"`
	Key      string `json:"key"`
	Bucket   string `json:"bucket"`
	BasePath string `json:"basePath"`
}

func Instance() *BackBlaze {
	once.Do(func() {
		instance = NewInstance("etc.backblaze.default")
	})

	return instance
}

// NewInstance 新建实例
func NewInstance[T string | Config | *Config](config T) *BackBlaze {
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

	client, err := b2.NewClient(context.Background(), conf.Account, conf.Key)
	if err != nil {
		log.Fatalf("new a backblaze b2 instance failed: %v", err)
	}

	bucket, err := client.Bucket(context.Background(), conf.Bucket)
	if err != nil {
		log.Fatalf("load backblaze b2 bucket failed: %v", err)
	}

	return &BackBlaze{conf: conf, client: client, bucket: bucket}
}

type BackBlaze struct {
	conf   *Config
	client *b2.Client
	bucket *b2.Bucket
}

// BaseUrl 获取基础链接
func (b *BackBlaze) BaseUrl() string {
	return strings.TrimSuffix(optionbaseconfig.GetValue(xtypes.AvatarUrlKey), "/")
}

// PutObject 上传对象
func (b *BackBlaze) PutObject(ctx context.Context, file any) (*Result, error) {
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
	path := strings.TrimPrefix(strings.TrimSuffix(b.conf.BasePath, "/"), "/") + "/" + name

	w := b.bucket.Object(path).NewWriter(ctx, b2.WithAttrsOption(&b2.Attrs{
		ContentType: b.getContentType(body),
	}))
	if _, err := io.Copy(w, body); err != nil {
		w.Close()
		return nil, err
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	return &Result{
		baseUrl: b.BaseUrl(),
		name:    name,
		path:    path,
	}, nil
}

// DeleteObject 删除对象
func (b *BackBlaze) DeleteObject(ctx context.Context, file string) error {
	path := strings.TrimSuffix(b.conf.BasePath, "/") + "/" + strings.TrimPrefix(file, "/")

	return b.bucket.Object(path).Delete(ctx)
}

// 获取文件内容类型
func (b *BackBlaze) getContentType(body io.Reader) string {
	buf := make([]byte, 500)

	if _, err := body.Read(buf); err != nil {
		return "application/octet-stream"
	}

	return http.DetectContentType(buf)
}

type Result struct {
	baseUrl string
	name    string
	path    string
}

func (o *Result) Name() string {
	return o.name
}

func (o *Result) Path() string {
	return o.path
}

func (o *Result) Url() string {
	return strings.TrimSuffix(o.baseUrl, "/") + "/" + strings.TrimPrefix(o.path, "/")
}

func (b *BackBlaze) PutFileBase64(ctx context.Context, file string) (*Result, error) {
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
	w := b.bucket.Object(path).NewWriter(ctx)

	if _, err = io.Copy(w, body); err != nil {
		w.Close()
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	return &Result{
		baseUrl: b.BaseUrl(),
		name:    name,
		path:    path,
	}, nil
}
