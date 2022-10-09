package auth

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func buildAuthUrl() string {
	authUrl, _ := url.Parse("https://linear.app/oauth/authorize")
	query := authUrl.Query()
	query.Set("client_id", "b3835b3ea649fe44b268ec8b2fd2f903")
	query.Set("redirect_uri", "http://localhost:8787")
	query.Set("response_type", "code")
	query.Set("scope", "read")
	authUrl.RawQuery = query.Encode()
	return authUrl.String()
}

func OpenAuthScreen() error {
	println("Opened Auth screen in browser. Please continue there")
	return openBrowser(buildAuthUrl())
}
