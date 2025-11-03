package handlers

import "github.com/gin-gonic/gin"

// POST /:company/admin/roster/upload  (Excel vendrá luego)
func AdminRosterUpload(c *gin.Context) {
	// TODO: validar rol GROUP_ADMIN, procesar Excel, crear/actualizar usuarios autorizados (sin exceder cupo)
	c.JSON(200, gin.H{"status": "accepted"})
}
