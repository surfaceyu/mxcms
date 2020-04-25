package cache

import (
	"mxcms/app/models"
	mxconst "mxcms/app/rpcsvr/common"
	"strconv"
)

type Menu models.Menu
type AdminRolePriv models.AdminRolePriv

func (this *Menu) GetKey(parentId int) string {
	return "hash:"+mxconst.CACHE_MENU+":"+strconv.Itoa(parentId)
}

func (this *Menu) GetAdminKey(roleId int) string {
	return "hash:"+mxconst.CACHE_MENU+"_ADMIN"+":"+strconv.Itoa(roleId)
}

func (this *AdminRolePriv) GetKey(roleId int) string {
	return "hash:"+mxconst.CACHE_ADMINROLEPRIV+":"+strconv.Itoa(roleId)
}
