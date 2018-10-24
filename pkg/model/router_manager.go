package model

type RouterManager interface {
	GetRoutes() []Route
	GetDefaultRoute() Route
	SetRoute(destination string, mask string, metric int) bool
}
