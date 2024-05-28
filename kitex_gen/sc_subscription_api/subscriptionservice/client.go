// Code generated by Kitex v0.9.1. DO NOT EDIT.

package subscriptionservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	sc_subscription_api "github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SubscriptionPage(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest, callOptions ...callopt.Option) (r *sc_subscription_api.SubscriptionPageResponse, err error)
	SubscriptionPrepare(ctx context.Context, request *sc_subscription_api.SubscriptionPrepareRequest, callOptions ...callopt.Option) (r *sc_subscription_api.SubscriptionPrepareResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kSubscriptionServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSubscriptionServiceClient struct {
	*kClient
}

func (p *kSubscriptionServiceClient) SubscriptionPage(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest, callOptions ...callopt.Option) (r *sc_subscription_api.SubscriptionPageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SubscriptionPage(ctx, request)
}

func (p *kSubscriptionServiceClient) SubscriptionPrepare(ctx context.Context, request *sc_subscription_api.SubscriptionPrepareRequest, callOptions ...callopt.Option) (r *sc_subscription_api.SubscriptionPrepareResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SubscriptionPrepare(ctx, request)
}
