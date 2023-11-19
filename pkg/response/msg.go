package response

import "xfd-backend/pkg/xerr"

var MsgFlags = map[xerr.XCode]string{
	xerr.SUCCESS:               "ok",
	xerr.UpdatePasswordSuccess: "修改密码成功",
	xerr.NotExistInentifier:    "该第三方账号未绑定",
	xerr.ERROR:                 "fail",
	xerr.InvalidParams:         "请求参数错误",

	xerr.ErrorUserAuthFailed:     "用户token失效",
	xerr.ErrorExistUser:          "已存在该用户名",
	xerr.ErrorNotExistUser:       "该用户不存在",
	xerr.ErrorNotCompare:         "账号密码错误",
	xerr.ErrorNotComparePassword: "两次密码输入不一致",
	xerr.ErrorFailEncryption:     "加密失败",
	xerr.ErrorNotExistProduct:    "该商品不存在",
	xerr.ErrorNotExistAddress:    "该收获地址不存在",
	xerr.ErrorExistFavorite:      "已收藏该商品",
	xerr.ErrorUserNotFound:       "用户不存在",
	xerr.ErrorTokenExpired:       "用户登录过期",
	xerr.ErrorOperationForbidden: "请求无效",

	xerr.ErrorBossCheckTokenFail:        "商家的Token鉴权失败",
	xerr.ErrorBossCheckTokenTimeout:     "商家Token已超时",
	xerr.ErrorBossToken:                 "商家的Token生成失败",
	xerr.ErrorBoss:                      "商家Token错误",
	xerr.ErrorBossInsufficientAuthority: "商家权限不足",
	xerr.ErrorBossProduct:               "商家读文件错误",

	xerr.ErrorProductExistCart:       "商品已经在购物车了，数量+1",
	xerr.ErrorProductMoreCart:        "超过最大上限",
	xerr.ErrorCartReduceMoreQuantity: "减少数量超过购物车数量",

	xerr.ErrorAuthCheckTokenFail:        "Token鉴权失败",
	xerr.ErrorAuthCheckTokenTimeout:     "Token已超时",
	xerr.ErrorAuthToken:                 "Token生成失败",
	xerr.ErrorAuth:                      "Token错误",
	xerr.ErrorAuthInsufficientAuthority: "权限不足",
	xerr.ErrorReadFile:                  "读文件失败",
	xerr.ErrorSendEmail:                 "发送邮件失败",
	xerr.ErrorCallApi:                   "调用接口失败",
	xerr.ErrorUnmarshalJson:             "解码JSON失败",

	xerr.ErrorUploadFile:    "上传失败",
	xerr.ErrorDeleteFile:    "删除失败",
	xerr.ErrorAdminFindUser: "管理员查询用户失败",

	xerr.ErrorDatabase: "数据库操作出错,请重试",

	xerr.ErrorOss:            "COS配置错误",
	xerr.ErrorNotExistRecord: "记录不存在",
	xerr.ErrorOrderNotFound:  "订单不存在",
	xerr.ErrorOrderCreate:    "订单创建失败",
	xerr.ErrorStockNotEnough: "商品库存不足",
	xerr.ErrorSomeoneNotPaid: "库存被锁定，有人未支付，请等待",
}

// GetMsg 获取状态码对应信息
func GetMsg(code xerr.XCode) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[xerr.ERROR]
}
