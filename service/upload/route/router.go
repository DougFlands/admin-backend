package route

import (
	"admin-server/service/apigateway/handler"
	"admin-server/service/upload/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// 支持跨域
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"}, // []string{"http://localhost:8090"},
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Range", "x-requested-with", "content-Type"},
		ExposeHeaders: []string{"Content-Length", "Accept-Ranges", "Content-Range", "Content-Disposition"},
		// AllowCredentials: true,
	}))

	router.Use(handler.Authorize())

	router.POST("/api/file/upload", api.UploadHandler)
	router.GET("/api/file/meta", api.GetFileMetaHandler)
	router.POST("/api/file/query", api.FileQueryHandler)
	router.POST("/api/file/rename", api.FileMetaRenameHandler)
	router.POST("/api/file/delete", api.FileDeleteHandle)

	router.POST("/api/file/fastupload", api.TryFastUploadHandler)

	// 从oss下载
	router.POST("/api/file/downloadurl", api.DownloadUrlHandler)

	// 分块上传 (暂未使用)
	router.POST("/api/file/mpupload/init", api.InitialMultipartUploadHandler)
	router.POST("/api/file/mpupload/uppart", api.UploadPartHandler)
	router.POST("/api/file/mpupload/complete", api.CompleteUploadHandler)

	return router
}
