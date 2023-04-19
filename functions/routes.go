package functions

//
//import (
//	"github.com/sarailQAQ/faased-netpan/functions/file"
//	"github.com/sarailQAQ/faased-netpan/functions/user"
//	"github.com/sarailQAQ/faased-netpan/internal/svc"
//	"net/http"
//
//	"github.com/zeromicro/go-zero/rest"
//)
//
//func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
//	server.AddRoutes(
//		[]rest.Route{
//			{
//				Method:  http.MethodPost,
//				Path:    "/login",
//				Handler: LoginHandler(serverCtx),
//			},
//			{
//				Method:  http.MethodPost,
//				Path:    "/register",
//				Handler: RegisterHandler(serverCtx),
//			},
//			{
//				Method:  http.MethodGet,
//				Path:    "/registerCode",
//				Handler: RegisterCodeHandler(serverCtx),
//			},
//			{
//				Method:  http.MethodGet,
//				Path:    "/share/basic/detail",
//				Handler: ShareBasicDetailHandler(serverCtx),
//			},
//		},
//	)
//
//	server.AddRoutes(
//		rest.WithMiddlewares(
//			[]rest.Middleware{serverCtx.Auth},
//			[]rest.Route{
//				{
//					Method:  http.MethodGet,
//					Path:    "/user/detail",
//					Handler: user.UserDetailHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPost,
//					Path:    "/user/repository/link",
//					Handler: user.UserRepositoryLinkHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodGet,
//					Path:    "/user/file/list",
//					Handler: user.UserFileListHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPut,
//					Path:    "/user/fileName/edit",
//					Handler: user.UserFileNameEditHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPost,
//					Path:    "/user/dir/create",
//					Handler: user.UserDirCreateHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodDelete,
//					Path:    "/user/file/delete",
//					Handler: user.UserDeleteFileHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPut,
//					Path:    "/user/file/move",
//					Handler: user.UserFileMovedHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPost,
//					Path:    "/user/share/basic/create",
//					Handler: user.ShareBasicCreateHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPost,
//					Path:    "/user/share/basic/save",
//					Handler: user.ShareBasicSaveHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodPost,
//					Path:    "/user/refresh/authorization",
//					Handler: user.RefreshAuthorizationHandler(serverCtx),
//				},
//				{
//					Method:  http.MethodGet,
//					Path:    "/user/repository/by/id",
//					Handler: user.GetUserRepostoryByIdHandler(serverCtx),
//				},
//			}...,
//		),
//	)
//
//	server.AddRoutes(
//		rest.WithMiddlewares(
//			[]rest.Middleware{serverCtx.Auth},
//			[]rest.Route{
//				{
//					Method:  http.MethodPost,
//					Path:    "/file/upload",
//					Handler: file.UploadFileHandler(serverCtx),
//				},
//			}...,
//		),
//	)
//}
