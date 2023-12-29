package xerr

type XCode int32

const (
	SUCCESS               XCode = 200
	UpdatePasswordSuccess XCode = 201
	NotExistInentifier    XCode = 202
	ERROR                 XCode = 500
	InvalidParams         XCode = 400

	//成员错误
	ErrorUserAuthFailed      XCode = 10001
	ErrorExistUser           XCode = 10002
	ErrorNotExistUser        XCode = 10003
	ErrorNotCompare          XCode = 10004
	ErrorNotComparePassword  XCode = 10005
	ErrorFailEncryption      XCode = 10006
	ErrorNotExistProduct     XCode = 10007
	ErrorNotExistAddress     XCode = 10008
	ErrorExistFavorite       XCode = 10009
	ErrorUserNotFound        XCode = 10010
	ErrorTokenExpired        XCode = 10011
	ErrorOperationForbidden  XCode = 10012
	ErrorUserOrgNotFound     XCode = 10013
	ErrorUserPointEmpty      XCode = 10014
	ErrorUserDuplicateVerify XCode = 10015

	//店家错误
	ErrorBossCheckTokenFail        XCode = 20001
	ErrorBossCheckTokenTimeout     XCode = 20002
	ErrorBossToken                 XCode = 20003
	ErrorBoss                      XCode = 20004
	ErrorBossInsufficientAuthority XCode = 20005
	ErrorBossProduct               XCode = 20006

	// 购物车
	ErrorProductExistCart       XCode = 20007
	ErrorProductMoreCart        XCode = 20008
	ErrorCartReduceMoreQuantity XCode = 20009

	//管理员错误
	ErrorAuthCheckTokenFail        XCode = 30001 //token 错误
	ErrorAuthCheckTokenTimeout     XCode = 30002 //token 过期
	ErrorAuthToken                 XCode = 30003
	ErrorAuth                      XCode = 30004
	ErrorAuthInsufficientAuthority XCode = 30005
	ErrorReadFile                  XCode = 30006
	ErrorSendEmail                 XCode = 30007
	ErrorCallApi                   XCode = 30008
	ErrorUnmarshalJson             XCode = 30009
	ErrorAdminFindUser             XCode = 30010
	ErrorVerifyEmpty               XCode = 30011

	//数据库错误
	ErrorDatabase         XCode = 40001
	ErrorRedis            XCode = 40002
	ErrorNotExistRecord   XCode = 40003
	ErrorInvalidFileExt   XCode = 40004
	ErrorInvalidCsvFormat XCode = 40005
	ErrorRedisLock        XCode = 40006

	//对象存储错误
	ErrorOss        XCode = 50001
	ErrorUploadFile XCode = 50002
	ErrorDeleteFile XCode = 50003

	//订单错误
	ErrorOrderNotFound  XCode = 60000
	ErrorOrderCreate    XCode = 60001
	ErrorStockNotEnough XCode = 60002
	ErrorSomeoneNotPaid XCode = 60003
)
