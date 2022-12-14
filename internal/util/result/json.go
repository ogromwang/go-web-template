package result

import (
	"github.com/gin-gonic/gin"
	"xiaohuazhu/internal/model"
)

func Ok(ctx *gin.Context, resp interface{}) {
	ctx.JSON(200, model.Result{
		Code:  200,
		Data:  resp,
		Error: "",
	})
}

func OkWithTotal(ctx *gin.Context, resp interface{}, total int64) {
	page := model.ResultWithPage{
		List:  resp,
		Total: total,
	}
	ctx.JSON(200, model.Result{
		Code:  200,
		Data:  page,
		Error: "",
	})
}

func OkWithMore(ctx *gin.Context, resp interface{}, more bool) {
	page := model.ResultWithMore{
		List: resp,
		More: more,
	}
	ctx.JSON(200, model.Result{
		Code:  200,
		Data:  page,
		Error: "",
	})
}

func Success(ctx *gin.Context) {
	ctx.JSON(200, model.Result{
		Code:  200,
		Data:  true,
		Error: "",
	})
}

func Fail(ctx *gin.Context, resp string) {
	ctx.JSON(200, model.Result{
		Code:  500,
		Data:  nil,
		Error: resp,
	})
}

func NoAuth(ctx *gin.Context, resp string) {
	ctx.JSON(200, model.Result{
		Code:  400,
		Data:  nil,
		Error: resp,
	})
}

func ServerError(ctx *gin.Context) {
	ctx.JSON(200, model.Result{
		Code:  505,
		Data:  nil,
		Error: "服务器内部错误",
	})
}
