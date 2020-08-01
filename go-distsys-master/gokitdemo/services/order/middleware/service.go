package middleware

import "github.com/dilipbaviskar/go-distsys/gokitdemo/services/order"

// Middleware describes a service middleware.
type Middleware func(service order.Service) order.Service
