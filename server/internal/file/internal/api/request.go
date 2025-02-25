package api

import "mime/multipart"

type UploadOnePictureRequest struct {
	Picture *multipart.FileHeader `form:"picture" binding:"required"`
}
