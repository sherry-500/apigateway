package main

import (
	"context"
	demo "github.com/Supernova114514/kitex_teacher/kitex_gen/demo"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct{}

// Register implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Register(ctx context.Context, teacher *demo.Teacher) (resp *demo.RegisterResp, err error) {
	// TODO: Your code here...
	resp = &demo.RegisterResp{
		Message: "add teacher",
		Success: true,
	}
	return
}

// Query implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Teacher, err error) {
	// TODO: Your code here...
	return
}
