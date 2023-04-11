package device

import (
	"confluence_fake/internal/service"
	"context"
)

/**
 * @Date: 2023/3/2 20:54
 * @Desc:
 */

type sDevice struct{}

func init() {
	service.RegisterDevice(New())
}

func New() *sDevice {
	return &sDevice{}
}

func (this *sDevice) GetList(ctx context.Context, in string) (out string, err error) {

	return
}
