package main

import (
	"github.com/asiainfoLDP/datafoundry_payment/api"
	"github.com/asiainfoLDP/datafoundry_payment/api/account"
	"github.com/asiainfoLDP/datafoundry_payment/api/amount"
	"github.com/asiainfoLDP/datafoundry_payment/api/balance"
	"github.com/asiainfoLDP/datafoundry_payment/api/checkout"
	"github.com/asiainfoLDP/datafoundry_payment/api/coupon"
	"github.com/asiainfoLDP/datafoundry_payment/api/integration"
	"github.com/asiainfoLDP/datafoundry_payment/api/market"
	"github.com/asiainfoLDP/datafoundry_payment/api/recharge"
	"github.com/julienschmidt/httprouter"
)

const (
	PAYMENT_API_PREFIX     = "/payment/v1"
	INTEGRATION_API_PREFIX = "/integration/v1"
)

func createRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", api.Index)

	router.POST(PAYMENT_API_PREFIX+"/recharge", recharge.Recharge)
	router.POST(PAYMENT_API_PREFIX+"/checkout", checkout.Checkout)
	router.GET(PAYMENT_API_PREFIX+"/balance", balance.Balance)
	router.GET(PAYMENT_API_PREFIX+"/market", market.Market)
	router.GET(PAYMENT_API_PREFIX+"/amounts", amount.AmountList)
	router.GET(PAYMENT_API_PREFIX+"/amounts/:tid", amount.Amount)
	router.GET(PAYMENT_API_PREFIX+"/account", account.Account)
	router.GET(PAYMENT_API_PREFIX+"/coupon/:serial", coupon.Coupon)
	router.POST(PAYMENT_API_PREFIX+"/redeem", coupon.Redeem)
	router.GET(PAYMENT_API_PREFIX+"/orders", checkout.Order)
	router.POST(PAYMENT_API_PREFIX+"/notification", recharge.Notification)
	router.GET(PAYMENT_API_PREFIX+"/regions", market.ListRegion)

	router.GET(INTEGRATION_API_PREFIX+"/repos", integration.ListRepos)
	router.GET(INTEGRATION_API_PREFIX+"/repos/:repo", integration.GetRepo)
	router.GET(INTEGRATION_API_PREFIX+"/repos/:repo/items/:item", integration.GetItem)

	router.NotFound = &api.Mux{}

	return router
}
