package cache

import (
	"mxcms/app/models"
	mxconst "mxcms/app/rpcsvr/common"
)

type AdminUser models.AdminUser

func (this *AdminUser) GetKey() string {
	return "hash"+mxconst.CACHE_ADMIN + this.Adminname
}
