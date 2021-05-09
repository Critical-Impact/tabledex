package mangodex

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	GetMDHomeURLPath = "at-home/server/%s"
)

// GetMDHomeURL : Get MangaDex@Home URL.
// https://api.mangadex.org/docs.html#operation/get-at-home-server-chapterId
func (dc *DexClient) GetMDHomeURL(chapId string, forcePort443 bool) (string, error) {
	return dc.GetMDHomeURLContext(context.Background(), chapId, forcePort443)
}

// GetMDHomeURLContext : GetMDHomeURL with custom context.
func (dc *DexClient) GetMDHomeURLContext(ctx context.Context, chapId string, forcePort443 bool) (string, error) {
	u, _ := url.Parse(BaseAPI)
	u.Path = fmt.Sprintf(GetMDHomeURLPath, chapId)

	// Set query parameters
	q := u.Query()
	q.Add("forcePort443", strconv.FormatBool(forcePort443))
	u.RawQuery = q.Encode()

	var r map[string]string
	_, err := dc.RequestAndDecode(ctx, http.MethodGet, u.String(), nil, &r)
	return r["baseUrl"], err
}
