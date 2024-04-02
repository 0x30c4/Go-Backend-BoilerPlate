package handlers

import (
	"strings"
)

func CheckIfBrowser(userAgent string) bool {
  if strings.Contains(userAgent, "Mozilla") || strings.Contains(userAgent, "Chrome") || strings.Contains(userAgent, "Safari") {
    // Client is likely a web browser
    return true
  } else {
    // Client may not be a web browser
    return false
  }
}
