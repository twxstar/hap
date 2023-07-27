package hap

import (
	"github.com/twxstar/hap/accessory"
	"github.com/twxstar/hap/characteristic"
	"github.com/twxstar/hap/log"
	"github.com/xiam/to"

	"encoding/json"
	"net/http"
	"strings"
)

type characteristicData struct {
	Aid   uint64      `json:"aid"`
	Iid   uint64      `json:"iid"`
	Value interface{} `json:"value"`

	// optional values
	Type        *string     `json:"type,omitempty"`
	Permissions []string    `json:"perms,omitempty"`
	Status      *int        `json:"status,omitempty"`
	Events      *bool       `json:"ev,omitempty"`
	Format      *string     `json:"format,omitempty"`
	Unit        *string     `json:"unit,omitempty"`
	MinValue    interface{} `json:"minValue,omitempty"`
	MaxValue    interface{} `json:"maxValue,omitempty"`
	MinStep     interface{} `json:"minStep,omitempty"`
	MaxLen      *int        `json:"maxLen,omitempty"`
	ValidValues []int       `json:"valid-values,omitempty"`
	ValidRange  []int       `json:"valid-values-range,omitempty"`
}

type putCharacteristicData struct {
	Aid uint64 `json:"aid"`
	Iid uint64 `json:"iid"`

	Value  interface{} `json:"value,omitempty"`
	Status *int        `json:"status,omitempty"`
	Events *bool       `json:"ev,omitempty"`

	Remote   *bool `json:"remote,omitempty"`
	Response *bool `json:"r,omitempty"`
}

func (srv *Server) getCharacteristics(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	// id=1.4,1.5
	v := req.FormValue("id")
	if len(v) == 0 {
		JsonError(res, JsonStatusInvalidValueInRequest)
		return
	}

	meta := req.FormValue("meta") == "1"
	perms := req.FormValue("perms") == "1"
	typ := req.FormValue("type") == "1"
	ev := req.FormValue("ev") == "1"

	arr := []*characteristicData{}
	err := false
	for _, str := range strings.Split(v, ",") {
		ids := strings.Split(str, ".")
		if len(ids) != 2 {
			continue
		}
		cdata := &characteristicData{
			Aid: to.Uint64(ids[0]),
			Iid: to.Uint64(ids[1]),
		}
		arr = append(arr, cdata)

		c := srv.findC(cdata.Aid, cdata.Iid)
		if c == nil {
			err = true
			status := JsonStatusServiceCommunicationFailure
			cdata.Status = &status
			continue
		}

		v, s := c.ValueRequest(req)
		if s != 0 {
			err = true
			cdata.Status = &s
		} else {
			cdata.Value = v
		}

		if meta {
			cdata.Format = &c.Format
			cdata.Unit = &c.Unit
			if c.MinVal != nil {
				cdata.MinValue = c.MinVal
			}
			if c.MaxVal != nil {
				cdata.MaxValue = c.MaxVal
			}
			if c.StepVal != nil {
				cdata.MinStep = c.StepVal
			}

			if c.MaxLen > 0 {
				cdata.MaxLen = &c.MaxLen
			}

			if len(c.ValidVals) > 0 {
				cdata.ValidValues = c.ValidVals
			}

			if len(c.ValidRange) > 0 {
				cdata.ValidRange = c.ValidRange
			}
		}

		// Should the response include the events flag?
		if ev {
			var ev bool
			if v, ok := c.Events[req.RemoteAddr]; ok {
				ev = v
			}
			cdata.Events = &ev
		}

		if perms {
			cdata.Permissions = c.Permissions
		}

		if typ {
			cdata.Type = &c.Type
		}
	}

	resp := struct {
		Characteristics []*characteristicData `json:"characteristics"`
	}{arr}

	log.Debug.Println(toJSON(resp))

	if err {
		JsonMultiStatus(res, resp)
	} else {
		JsonOK(res, resp)
	}
}

func (srv *Server) putCharacteristics(res http.ResponseWriter, req *http.Request) {
	if !srv.IsAuthorized(req) {
		log.Info.Printf("request from %s not authorized\n", req.RemoteAddr)
		JsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	data := struct {
		Cs []putCharacteristicData `json:"characteristics"`
	}{}
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		JsonError(res, JsonStatusInvalidValueInRequest)
		return
	}

	log.Debug.Println(toJSON(data))

	arr := []*putCharacteristicData{}
	for _, d := range data.Cs {
		c := srv.findC(d.Aid, d.Iid)
		cdata := &putCharacteristicData{
			Aid: d.Aid,
			Iid: d.Iid,
		}

		if c == nil {
			status := JsonStatusServiceCommunicationFailure
			cdata.Status = &status
			arr = append(arr, cdata)
			continue
		}

		var value interface{}
		var status int
		if d.Value != nil {
			value, status = c.SetValueRequest(d.Value, req)
		}

		if status != 0 {
			cdata.Status = &status
		}

		if d.Response != nil && value != nil {
			cdata.Value = value
		}

		if d.Events != nil {
			if !c.IsObservable() {
				status := JsonStatusNotificationNotSupported
				cdata.Status = &status
				arr = append(arr, cdata)
			} else {
				c.Events[req.RemoteAddr] = *d.Events
			}
		}

		if cdata.Status != nil || cdata.Value != nil {
			arr = append(arr, cdata)
		}
	}

	if len(arr) == 0 {
		res.WriteHeader(http.StatusNoContent)
		return
	}

	resp := struct {
		Characteristics []*putCharacteristicData `json:"characteristics"`
	}{arr}

	log.Debug.Println(toJSON(resp))
	JsonMultiStatus(res, resp)
}

func (srv *Server) findC(aid, iid uint64) *characteristic.C {
	var as []*accessory.A
	as = append(as, srv.a)
	as = append(as, srv.as[:]...)

	for _, a := range as {
		if a.Id == aid {
			for _, s := range a.Ss {
				for _, c := range s.Cs {
					if c.Id == iid {
						return c
					}
				}
			}
		}
	}

	return nil
}
