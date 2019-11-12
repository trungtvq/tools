package format
import (
	"context"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func syncCookie(sd runtime.ServerMetadata, w http.ResponseWriter, key authKey) {
	md := sd.HeaderMD
	k := string(key)
	v := md[k]
	if len(v) > 0 && v[0] != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     k,
			HttpOnly: true,
			Value:    v[0],
			Path:     "/",
		})
	}
}

// FormatHTTPResponse format http response from proto messages
func FormatHTTPResponse(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	// fmt.Println("format", ctx, w, resp)
	// w.Header().Set("X-My-Tracking-Token", resp.Token)
	// resp.Token = ""

	md, _ := runtime.ServerMetadataFromContext(ctx)
	syncCookie(md, w, DeviceIDCookie)
	syncCookie(md, w, AuthCookie)
	return nil
}
