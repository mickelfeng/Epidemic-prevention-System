package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh-pwh/Epidemic-prevention-System/controller"
	"github.com/pwh-pwh/Epidemic-prevention-System/middlewares"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}
	r := gin.New()
	r.Use(middlewares.Recover)
	r.Use(middlewares.CorsMiddleware())
	r.GET("/captcha", controller.GetCaptcha)
	r.POST("/login", middlewares.CaptchaMiddleware(), controller.LoginHander)
	r.GET("/userInfo", middlewares.JwtAuth(""), controller.UserInfo)
	//news 相关接口
	r.GET("/news", middlewares.JwtAuth(""), controller.News)
	r.GET("/chinaData", middlewares.JwtAuth(""), controller.ChinaData)
	r.GET("/riskarea", middlewares.JwtAuth(""), controller.GetRiskArea)
	r.GET("/history", middlewares.JwtAuth(""), controller.GetHistory)
	r.GET("/infiniteNews", middlewares.JwtAuth(""), controller.InfiniteNews)
	r.GET("/vaccineTopData", middlewares.JwtAuth(""), controller.VaccineTopData)
	r.GET("/chinaVaccineTrendData", middlewares.JwtAuth(""), controller.ChinaVaccineTrendData)
	r.GET("/rumor", middlewares.JwtAuth(""), controller.Rumor)

	r.POST("/upload", middlewares.JwtAuth(""), controller.Upload)
	//注册路由
	menuGroup := r.Group("/sys/menu")
	menuGroup.GET("/nav", middlewares.JwtAuth(""), controller.Nav)
	menuGroup.GET("/list", middlewares.JwtAuth("sys:menu:list"), controller.ListMenu)
	menuGroup.DELETE("/:id", middlewares.JwtAuth("sys:menu:delete"), controller.DeleteMenu)
	menuGroup.POST("", middlewares.JwtAuth("sys:menu:save"), controller.SaveMenu)
	menuGroup.PUT("", middlewares.JwtAuth("sys:menu:update"), controller.EditMenu)
	//公告路由
	noticeGroup := r.Group("/sys/notice")
	noticeGroup.GET("", middlewares.JwtAuth(""), controller.GetNotice)
	noticeGroup.GET("/list", middlewares.JwtAuth("monitor:notice:list"), controller.ListNotice)
	noticeGroup.POST("", middlewares.JwtAuth("monitor:notice:save"), controller.SaveNotice)
	noticeGroup.PUT("", middlewares.JwtAuth("monitor:notice:update"), controller.UpdateNotice)
	noticeGroup.DELETE("", middlewares.JwtAuth("monitor:notice:delete"), controller.DeleteNotice)
	noticeGroup.GET("/:id", middlewares.JwtAuth("monitor:notice:set"), controller.SetNotice)
	//未回归路由
	arGroup := r.Group("/access/return")
	arGroup.GET("/list", middlewares.JwtAuth("access:return:list"), controller.AccessReturnList)
	//出入登记路由
	aRgistGroup := r.Group("/access/register")
	aRgistGroup.GET("/list", middlewares.JwtAuth("access:register:list"), controller.GetAccessRegisterList)
	aRgistGroup.POST("", middlewares.JwtAuth("access:register:save"), controller.SaveAccessRegister)
	//goodsinfo 路由
	goodsInfoGroup := r.Group("/good/info")
	goodsInfoGroup.GET("/all", middlewares.JwtAuth("good:info:list"), controller.GetAllGoodsInfo)
	goodsInfoGroup.GET("/total", middlewares.JwtAuth("good:info:list"), controller.GetTotalGoodsInfo)
	goodsInfoGroup.GET("/list", middlewares.JwtAuth("good:info:list"), controller.GetListGoodsInfo)
	goodsInfoGroup.POST("", middlewares.JwtAuth("good:info:save"), controller.SaveGoodsInfo)
	goodsInfoGroup.PUT("", middlewares.JwtAuth("good:info:update"), controller.UpdateGoodsInfo)
	goodsInfoGroup.DELETE("", middlewares.JwtAuth("good:info:delete"), controller.DeleteGoodsInfo)
	//goodstype 路由
	goodsTypeGroup := r.Group("/good/type")
	goodsTypeGroup.GET("", middlewares.JwtAuth("good:type:list"), controller.GetSimpleListGoodsType)
	goodsTypeGroup.GET("/list", middlewares.JwtAuth("good:type:list"), controller.GetListGoodsType)
	goodsTypeGroup.POST("", middlewares.JwtAuth("good:type:save"), controller.SaveGoodsType)
	goodsTypeGroup.PUT("", middlewares.JwtAuth("good:type:update"), controller.UpdateGoodsType)
	goodsTypeGroup.DELETE("", middlewares.JwtAuth("good:type:delete"), controller.DeleteGoodsType)
	//物资出入库路由
	goodsStockGroup := r.Group("/good/stock")
	goodsStockGroup.GET("/list", middlewares.JwtAuth("good:stock:list"), controller.GetListGoodsStock)
	//good:stock:operate
	goodsStockGroup.POST("", middlewares.JwtAuth("good:stock:operate"), controller.SaveGoodsStock)
	///health/clock
	healthClockGroup := r.Group("/health/clock")
	healthClockGroup.GET("/list", middlewares.JwtAuth("health:clock:list"), controller.GetListHealthClock)
	healthClockGroup.GET("", middlewares.JwtAuth(""), controller.CheckHealthClock)
	healthClockGroup.POST("", middlewares.JwtAuth("health:clock:save"), controller.SaveHealthClock)
	// dept route
	deptGroup := r.Group("/sys/dept")
	deptGroup.GET("/list/:flag", controller.ListDept)
	deptGroup.DELETE("/:id", middlewares.JwtAuth("sys:dept:delete"), controller.DeleteDeptById)
	deptGroup.PUT("", middlewares.JwtAuth("sys:dept:update"), controller.UpdateDept)
	deptGroup.POST("", middlewares.JwtAuth("sys:dept:save"), controller.SaveDept)
	// health/report route
	healthReportGroup := r.Group("/health/report")
	healthReportGroup.GET("", middlewares.JwtAuth(""), controller.CheckHealthReport)
	healthReportGroup.GET("/list", middlewares.JwtAuth("health:report:list"), controller.ListHealthReport)
	healthReportGroup.POST("", middlewares.JwtAuth("health:report:save"), controller.SaveHealthReport)
	// leave/apply route
	leaveApplyGroup := r.Group("/leave/apply")
	leaveApplyGroup.GET("/list", middlewares.JwtAuth("leave:apply:list,leave:record:list"), controller.GetListLeaveApply)
	leaveApplyGroup.POST("", middlewares.JwtAuth("leave:apply:save"), controller.SaveLeaveApply)
	leaveApplyGroup.PUT("", middlewares.JwtAuth("leave:apply:update,leave:record:examine"), controller.UpdateLeaveApply)
	// /monitor/redis
	monitorGroup := r.Group("/monitor")
	monitorGroup.GET("/redis", middlewares.JwtAuth("monitor:redis:list"), controller.GetRedisInfo)

	// /register
	registerGroup := r.Group("/register")
	registerGroup.POST("", middlewares.CaptchaMiddleware(), controller.Register)
	registerGroup.GET("/deptList", controller.DeptList)

	// /sys/loginInfo
	loginInfoGroup := r.Group("/sys/loginInfo")
	loginInfoGroup.POST("", middlewares.JwtAuth("sys:login:clear"), controller.ClearLoginInfo)
	loginInfoGroup.DELETE("", middlewares.JwtAuth("sys:login:delete"), controller.DeleteLoginInfo)
	loginInfoGroup.GET("/list", middlewares.JwtAuth("sys:login:list"), controller.ListLoginInfo)

	// /sys/operateLog
	opLogGroup := r.Group("/sys/operateLog")
	opLogGroup.POST("", middlewares.JwtAuth("monitor:operate:clear"), controller.ClearOpLog)
	opLogGroup.GET("/list", middlewares.JwtAuth("monitor:operate:list"), controller.ListOpLog)
	opLogGroup.DELETE("", middlewares.JwtAuth("monitor:operate:delete"), controller.DeleteOpLog)

	// /sys/role
	roleGroup := r.Group("/sys/role")
	roleGroup.DELETE("", middlewares.JwtAuth("sys:role:delete"), controller.DeleteRole)
	roleGroup.GET("/info/:id", middlewares.JwtAuth(""), controller.InfoRole)
	roleGroup.GET("/list", middlewares.JwtAuth("sys:role:list"), controller.ListRole)
	roleGroup.POST("", middlewares.JwtAuth("sys:role:save"), controller.AddRole)
	roleGroup.PUT("", middlewares.JwtAuth("sys:role:update"), controller.EditRole)

	// /sys/user
	userGroup := r.Group("/sys/user")
	userGroup.POST("/avatar", middlewares.JwtAuth(""), controller.Avatar)
	userGroup.GET("/updatePassword", middlewares.JwtAuth("sys:user:update"), controller.UpdatePassword)
	userGroup.POST("/updateInfo", middlewares.JwtAuth("sys:user:update"), controller.UpdateInfo)
	userGroup.POST("/reset", middlewares.JwtAuth("sys:user:repass"), controller.ResetPwd)
	userGroup.POST("/userRole/:id", middlewares.JwtAuth("sys:user:role"), controller.ApplyUserRole)
	userGroup.DELETE("", middlewares.JwtAuth("sys:user:delete"), controller.DeleteUesr)
	userGroup.PUT("", middlewares.JwtAuth("sys:user:update"), controller.UpdateUser)
	userGroup.POST("", middlewares.JwtAuth("sys:user:save"), controller.AddUser)
	userGroup.GET("/list", middlewares.JwtAuth("sys:user:list"), controller.ListUser)
	userGroup.GET("/info/:id", controller.Info)
	return r
}
