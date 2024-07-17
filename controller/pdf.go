package controller

import (
	"bytes"
	"candidate-backend/models"
	"io"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func DecryptPdfFile(c *gin.Context) {

	var req models.PDFDecryptRequest
	
    if err := c.ShouldBind(&req); err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}

	//read file
	file, err := req.PdfFile.Open()
    if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
    }
    defer file.Close()

	
	var buf bytes.Buffer
	var decryptedBuf bytes.Buffer
	w := &decryptedBuf

    _, err = io.Copy(&buf, file)
    if err != nil {
		boom.BadRequest(c.Writer, err.Error());
        return
    }

	rs := bytes.NewReader(buf.Bytes())
	
	conf := model.NewAESConfiguration("", req.PdfPassword, 256)
	err = api.Decrypt(rs, w, conf)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
        return
    }

	c.Writer.Header().Set("Content-Type", "application/pdf");
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(decryptedBuf.Bytes())
}