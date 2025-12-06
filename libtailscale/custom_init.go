package libtailscale

import (
  _ "embed"
  "log"
  "tailscale.com/net/bakedroots"
	"tailscale.com/envknob"
)

//go:embed ca/ca_cert.pem
var myCA []byte

func init() {
  // Append custom CA
  pool := bakedroots.Get()
  if ok := pool.AppendCertsFromPEM(myCA); !ok {
    log.Printf("custom_init: Failed to append custom CA")
  }
  log.Printf("custom_init: Custom CA appended")

  // Disable log uploading
  envknob.SetNoLogsNoSupport()
  log.Printf("custom_init: Log uploading disabled")
}
