package css

import (
	"encoding/base64"
	"log"
	"net/http"
)

func SigninCSS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting page: %s", r.URL.Path)
	css_b64 := "Ym9keSB7CiAgcGFkZGluZy10b3A6IDQwcHg7CiAgcGFkZGluZy1ib3R0b206IDQwcHg7CiAgYmFja2dyb3VuZC1jb2xvcjogI2VlZTsKfQoKLmZvcm0tc2lnbmluIHsKICBtYXgtd2lkdGg6IDUzMHB4OwogIHBhZGRpbmc6IDE1cHg7CiAgbWFyZ2luOiAwIGF1dG87Cn0KLmZvcm0tc2lnbmluIC5mb3JtLXNpZ25pbi1oZWFkaW5nLAouZm9ybS1zaWduaW4gLmNoZWNrYm94IHsKICBtYXJnaW4tYm90dG9tOiAxMHB4Owp9Ci5mb3JtLXNpZ25pbiAuY2hlY2tib3ggewogIGZvbnQtd2VpZ2h0OiBub3JtYWw7Cn0KLmZvcm0tc2lnbmluIC5mb3JtLWNvbnRyb2wgewogIHBvc2l0aW9uOiByZWxhdGl2ZTsKICBoZWlnaHQ6IGF1dG87CiAgLXdlYmtpdC1ib3gtc2l6aW5nOiBib3JkZXItYm94OwogICAgIC1tb3otYm94LXNpemluZzogYm9yZGVyLWJveDsKICAgICAgICAgIGJveC1zaXppbmc6IGJvcmRlci1ib3g7CiAgcGFkZGluZzogMTBweDsKICBmb250LXNpemU6IDE2cHg7Cn0KLmZvcm0tc2lnbmluIC5mb3JtLWNvbnRyb2w6Zm9jdXMgewogIHotaW5kZXg6IDI7Cn0KLmZvcm0tc2lnbmluIGlucHV0W3R5cGU9ImVtYWlsIl0gewogIG1hcmdpbi1ib3R0b206IC0xcHg7CiAgYm9yZGVyLWJvdHRvbS1yaWdodC1yYWRpdXM6IDA7CiAgYm9yZGVyLWJvdHRvbS1sZWZ0LXJhZGl1czogMDsKfQouZm9ybS1zaWduaW4gaW5wdXRbdHlwZT0icGFzc3dvcmQiXSB7CiAgbWFyZ2luLWJvdHRvbTogMTBweDsKICBib3JkZXItdG9wLWxlZnQtcmFkaXVzOiAwOwogIGJvcmRlci10b3AtcmlnaHQtcmFkaXVzOiAwOwp9"
	css, _ := base64.StdEncoding.DecodeString(css_b64)
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(css))
}
