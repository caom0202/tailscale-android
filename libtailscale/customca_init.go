package libtailscale

import (
  _ "embed"
  "log"
  "tailscale.com/net/bakedroots"
)

//go:embed ca/ca_cert.pem
var myCA []byte

func init() {
  pool := bakedroots.Get()
  if ok := pool.AppendCertsFromPEM(myCA); !ok {
    log.Printf("customca_init: Failed to append custom CA")
  }
}
