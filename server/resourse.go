package server

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/zhengxiaoyao0716/util/zip"

	"gopkg.in/macaron.v1"
)

// DownloadApp 下载app
func DownloadApp(ctx *macaron.Context) (int, []byte) {
	buffer := bytes.Buffer{}
	writer := zip.NewWriter(&buffer)

	writer.Prefix = "nssk-emulate/"
	if err := writer.WriteFiles("html"); err != nil {
		code, content := MakeErr(ctx, 403, err)
		return code, []byte(content)
	}
	if err := writer.WriteFiles("nssk-emulate.exe"); err != nil {
		code, content := MakeErr(ctx, 403, err)
		return code, []byte(content)
	}

	host := ctx.RemoteAddr()
	portIndex := strings.Index(host, ":")
	if portIndex != -1 {
		host = host[0:portIndex]
	}

	batBytes := []byte("@echo off\n" +
		fmt.Sprintln("cd", writer.Prefix) +
		fmt.Sprintln("set", "hour=%time:~0,2%") +
		fmt.Sprintln(
			"start", "/min", "nssk-emulate.exe",
			"-host", host, "-master", GetStrCache("address"),
			"%1 %2 %3 %4 %5 %6 %7 %8 %9", "2 >>",
			"nssk-emulate-%date:~0,4%_%date:~5,2%_%date:~8,2%-%hour: =0%_%time:~3,2%_%time:~6,2%.log",
		),
	)
	writer.Prefix = ""
	if err := writer.WriteBytes("nssk-emulate.bat", batBytes); err != nil {
		code, content := MakeErr(ctx, 403, err)
		return code, []byte(content)
	}

	writer.Close()
	header := ctx.Header()
	header.Add("Content-Disposition", "filename=nssk-emulate.zip")
	header.Add("Content-Type", "application/octet-stream")
	return 200, buffer.Bytes()
}
