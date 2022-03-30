package request

import "os"

var DEBUG = os.Getenv("DEBUG") == "true"
