// Code generated by Kitex v0.6.1. DO NOT EDIT.

package studentservice

import (
	server "github.com/cloudwego/kitex/server"
	demo "github.com/sherry-500/kitex_student/kitex_gen/demo"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler demo.StudentService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
