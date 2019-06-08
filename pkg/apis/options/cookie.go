package options

import (
	"time"

	flag "github.com/spf13/pflag"
)

// CookieOptions contains configuration options relating to Cookie configuration
type CookieOptions struct {
	Name     string        `flag:"cookie-name" cfg:"cookie_name" env:"OAUTH2_PROXY_COOKIE_NAME" default:"_oauth2_proxy"`
	Secret   string        `flag:"cookie-secret" cfg:"cookie_secret" env:"OAUTH2_PROXY_COOKIE_SECRET"`
	Domain   string        `flag:"cookie-domain" cfg:"cookie_domain" env:"OAUTH2_PROXY_COOKIE_DOMAIN"`
	Path     string        `flag:"cookie-path" cfg:"cookie_path" env:"OAUTH2_PROXY_COOKIE_PATH" default:"/"`
	Expire   time.Duration `flag:"cookie-expire" cfg:"cookie_expire" env:"OAUTH2_PROXY_COOKIE_EXPIRE" default:"604800000000000"`
	Refresh  time.Duration `flag:"cookie-refresh" cfg:"cookie_refresh" env:"OAUTH2_PROXY_COOKIE_REFRESH" default:"0"`
	Secure   bool          `flag:"cookie-secure" cfg:"cookie_secure" env:"OAUTH2_PROXY_COOKIE_SECURE" default:"true"`
	HTTPOnly bool          `flag:"cookie-httponly" cfg:"cookie_httponly" env:"OAUTH2_PROXY_COOKIE_HTTPONLY" default:"true"`
}

// cookieFlagSet contains flags related to Cookie configuration
var cookieFlagSet = flag.NewFlagSet("cookie", flag.ExitOnError)

func init() {
	cookieFlagSet.String("cookie-name", "_oauth2_proxy", "the name of the cookie that the oauth_proxy creates")
	cookieFlagSet.String("cookie-secret", "", "the seed string for secure cookies (optionally base64 encoded)")
	cookieFlagSet.String("cookie-domain", "", "an optional cookie domain to force cookies to (ie: .yourcompany.com)*")
	cookieFlagSet.String("cookie-path", "/", "an optional cookie path to force cookies to (ie: /poc/)*")
	cookieFlagSet.Duration("cookie-expire", time.Duration(168)*time.Hour, "expire timeframe for cookie")
	cookieFlagSet.Duration("cookie-refresh", time.Duration(0), "refresh the cookie after this duration; 0 to disable")
	cookieFlagSet.Bool("cookie-secure", true, "set secure (HTTPS) cookie flag")
	cookieFlagSet.Bool("cookie-httponly", true, "set HttpOnly cookie flag")
}
