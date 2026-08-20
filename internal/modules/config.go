package modules

import (
	"time"
)

const HTTP_CLIENT_TIMEOUT_INT_SECONDS = 15
const HTTP_CLIENT_TIMEOUT_SECONDS = time.Duration(HTTP_CLIENT_TIMEOUT_INT_SECONDS) * time.Second
