package handlers

import (
	"github.com/blocktop/mp-common/server"
	"github.com/blocktop/mp-transfer-server/config"
	"io/ioutil"
	"net/http"
)

func HandleGetInfo(w http.ResponseWriter, r *http.Request) {
	cfg := config.GetConfig()
	info, err := ioutil.ReadFile(cfg.TransferServerInfoPath)
	if err != nil {
		server.ResponseError(w, http.StatusInternalServerError, server.NOFILE, err)
		return
	}

	server.ResponseJSONString(w, string(info))
}
