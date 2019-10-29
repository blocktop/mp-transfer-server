package handlers

import (
	"fmt"
	"github.com/blocktop/mp-common/server"
	"github.com/blocktop/mp-transfer-server/config"
	"io/ioutil"
	"net/http"
)

const (
	// NOFILE is the error code for a file not found
	NOFILE = "NOFILE"
)

func HandleGetInfo(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()
	info, err := ioutil.ReadFile(cfg.TransferServerInfoPath)
	if err != nil {
		server.ResponseError(w, http.StatusInternalServerError, NOFILE, fmt.Errorf("error reading the MP_TRANSFER_SERVER_INFO_PATH at %s", cfg.TransferServerInfoPath))
		return
	}

	server.ResponseJSONString(w, string(info))
}
