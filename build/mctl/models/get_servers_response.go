package models

import "github.com/skkrimon/mc/mctl/util"

type GetServersResponse struct {
	Servers []util.Server `json:"servers"`
}
