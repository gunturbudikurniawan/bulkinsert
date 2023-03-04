package middleware

import (
	"bpjs/config"
	"bytes"
	"context"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"
)

var SysErr string

type Log struct {
	Uri       string `json:"uri"`
	Method    string `json:"method"`
	ReqFrom   string `json:"req_from"`
	ReqTo     string `json:"req_to"`
	Status    int    `json:"status"`
	Code      int    `json:"code"`
	Data      Data   `json:"data"`
	IsSuccess bool   `json:"is_success"`
	Message   string `json:"message"`
	SysErr    string `json:"sys_err"`
	ClientIP  string `json:"client_ip"`
	HostIP    string `json:"host_ip"`
	StartedOn string `json:"started_on"`
	EndedOn   string `json:"ended_on"`
	RespTime  int    `json:"resp_time"`
	ReqId     string `json:"req_id"`
}

type Data struct {
	Req  interface{} `json:"req"`
	Resp interface{} `json:"resp"`
}

type LogResponseWriter struct {
	http.ResponseWriter
	statusCode int
	buf        bytes.Buffer
}

func NewLogResponseWriter(w http.ResponseWriter) *LogResponseWriter {
	return &LogResponseWriter{ResponseWriter: w}
}

func (w *LogResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *LogResponseWriter) Write(body []byte) (int, error) {
	w.buf.Write(body)
	return w.ResponseWriter.Write(body)
}

type resp struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	IsSuccess bool        `json:"is_success"`
	Message   string      `json:"message"`
}

func InternalLogger() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Reset sys_err value for each trx in new relic
			SysErr = ""
			starttime := time.Now().In(config.App.TLoc)

			var reqdata interface{}
			var respdata resp
			var reqid string

			// If request method is GET then reqdata obtained from query params
			// Else reqdata obtained from body
			if r.Method == "GET" {
				reqdata = r.URL.Query()
			} else {
				// Reading body into memory and reassigning to r.body to avoid EOF error
				bodyBytes, _ := io.ReadAll(r.Body)
				r.Body.Close()
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

				if len(bodyBytes) > 0 {
					err := json.Unmarshal(bodyBytes, &reqdata)
					if err != nil {
						log.Println("Error: Internal request log failed to Unmarshall -", err)
					}
				}
			}

			logRespWriter := NewLogResponseWriter(w)

			// Generate UUID as request unique identifier
			if u, err := uuid.NewV4(); err != nil {
				reqid = err.Error()
			} else {
				reqid = u.String()
			}
			ctx := context.WithValue(r.Context(), "reqid", reqid)

			// Serving http request
			next.ServeHTTP(logRespWriter, r.WithContext(ctx))

			// Converting resp body to struct
			err := json.Unmarshal(logRespWriter.buf.Bytes(), &respdata)
			if err != nil {
				log.Println("Error: Internal response log failed to Unmarshall -", err)
			}

			endtime := time.Now().In(config.App.TLoc)

			ilog := Log{
				Uri:     r.URL.Path,
				Method:  r.Method,
				ReqFrom: r.Header.Get("x-consumer-username"),
				ReqTo:   config.AppName,
				Status:  logRespWriter.statusCode,
				Code:    respdata.Code,
				Data: Data{
					Req:  reqdata,
					Resp: respdata.Data,
				},
				IsSuccess: respdata.IsSuccess,
				Message:   respdata.Message,
				SysErr:    SysErr,
				ClientIP:  r.RemoteAddr,
				HostIP:    config.Glb.Serv.Host + ":" + config.Glb.Serv.Port,
				StartedOn: starttime.Format(config.AppTLayout),
				EndedOn:   endtime.Format(config.AppTLayout),
				RespTime:  int(endtime.Sub(starttime).Milliseconds()),
				ReqId:     reqid,
			}

			// Convert log to json
			l, err := json.Marshal(ilog)
			if err != nil {
				log.Println("Error: Internal log failed to Marshall -", err)
			}

			// Print log
			fmt.Println(string(l))
		})
	}
}
