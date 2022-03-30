package sign

import "os"

var DEBUG = os.Getenv("DEBUG") == "true"
