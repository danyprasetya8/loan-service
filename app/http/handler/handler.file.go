package handler

import (
	"fmt"
	"loan-service/pkg/responsehelper"

	"github.com/gin-gonic/gin"
)

// DownloadFile
//
//	@Summary	Download file
//	@Tags		File
//	@Produce	octet-stream
//	@Param		id	path	string	true	"File ID"
//	@Success	200	{file}	binary
//	@Router		/api/v1/file/{id}/_download [POST]
func (h *Handler) DownloadFile(c *gin.Context) {
	fileID := c.Param("id")
	fileModel := h.fileService.Find(fileID)

	if fileModel == nil {
		responsehelper.BadRequest(c, "file not found")
		return
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileModel.OriginalName))
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(fileModel.Path)
}
