package model

// MethodCallHolder defins all graph that should store method calls.
type MethodCallHolder interface {
	AcceptMethodCall(methodCall *MethodCall) error
}
