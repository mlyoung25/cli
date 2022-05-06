package api

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
	hgql "github.com/hasura/go-graphql-client"
	"github.com/spf13/viper"
	"github.com/zeet-dev/pkg/utils"
	"k8s.io/client-go/transport"
)

type Client struct {
	GQL          graphql.Client
	Subscription *hgql.SubscriptionClient
}

func New(host string, accessToken string) *Client {
	client := newGraphQLClient(host, accessToken)
	subscription := newSubscriptionClient(host, accessToken)

	return &Client{GQL: client, Subscription: subscription}
}

func newGraphQLClient(server, token string) graphql.Client {
	tp := http.DefaultTransport
	if viper.GetBool("debug") {
		tp = utils.LoggingHttpTransport
	}

	return graphql.NewClient(utils.URLJoin(server, "graphql"), &http.Client{
		Transport: transport.NewBearerAuthRoundTripper(token, tp),
	})
}

func newSubscriptionClient(server, token string) *hgql.SubscriptionClient {
	client := hgql.NewSubscriptionClient(server).WithConnectionParams(map[string]interface{}{
		"headers": map[string]string{
			"Authentication": "Bearer " + token,
		},
	})

	return client
}