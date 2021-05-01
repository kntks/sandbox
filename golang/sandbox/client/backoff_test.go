package client

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"go.uber.org/goleak"
)

func ok(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{result: ok}`))
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{result: badrequest}`))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`{result: internal server error}`))
}

func httpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/200", ok)
	mux.HandleFunc("/4xx", badRequest)
	mux.HandleFunc("/5xx", internalServerError)
	return mux
}

func TestMain(m *testing.M) {
	code := m.Run()
	goleak.VerifyTestMain(m)
	os.Exit(code)
}

func TestGet(t *testing.T) {

	ts := httptest.NewServer(httpHandler())
	defer ts.Close()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name:    "StatusCode200のとき意図したデータが取得できる",
			args:    args{url: ts.URL + "/200"},
			want:    &Result{Status: http.StatusOK, Body: []byte(`{result: ok}`)},
			wantErr: false,
		},
		{
			name:    "StatusCode400のとき意図したエラーコードが取得できる",
			args:    args{url: ts.URL + "/4xx"},
			want:    &Result{Status: http.StatusBadRequest, Body: []byte(`{result: badrequest}`)},
			wantErr: true,
		},
		{
			name:    "StatusCode500のとき意図したエラーコードが取得できる",
			args:    args{url: ts.URL + "/5xx"},
			want:    &Result{Status: http.StatusInternalServerError, Body: []byte(`{result: internal server error}`)},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.url, 1, RetryDelayGenerator(1))
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
