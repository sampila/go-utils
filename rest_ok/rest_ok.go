package rest_ok

import()

type RestOk interface {
	MetaHead() metaHead
	Data()     interface{}
  Total()    []int
}

type metaHead struct {
  Message    string   `json:"message"`
	Status     int      `json:"status"`
	Error      bool     `json:"error"`
  Reason     string   `json:"reason"`
}

type restOk struct {
	RestMetaHead metaHead     `json:"meta"`
	RestData     interface{}  `json:"data"`
  RestTotal    []int       `json:"total,omitempty"`
}

func (r restOk) MetaHead() metaHead {
	return r.RestMetaHead
}

func (r restOk) Data() interface{} {
	return r.RestData
}

func (r restOk) Total() []int {
	return r.RestTotal
}

func NewRestOK(message string, status int, err bool, reason string, payload interface{}, total ...int) (RestOk) {
  mh := metaHead{ Message: message,
							     Status: status,
							     Error: err,
							     Reason: reason }

  return restOk{ RestMetaHead: mh,
								 RestData:  payload,
							   RestTotal: total }
}
