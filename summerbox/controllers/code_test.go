package controllers

import (
	"github.com/skip2/go-qrcode"
	"testing"
)

func Test(t *testing.T)  {
	path:="https://github.com/ITAleftback"
	_ = qrcode.WriteFile(path, qrcode.Medium, 256, "qr.png")
}
