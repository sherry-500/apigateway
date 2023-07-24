package main

import (
	"context"
	demo "github.com/Supernova114514/kitex_teacher/kitex_gen/demo"
	"fmt"
	"errors"
	"log"
	//"time"
	"strings"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TeacherServiceImpl implements the last service interface defined in the IDL.
type TeacherServiceImpl struct{
	db *gorm.DB
}
type Teacher struct {
	Id int32
	Name string
	Email string
	CollegeName string
	CollegeAddress string
	Sex string
	//CreatedAt time.Time `gorm:"default:CURRENT_TMESTAMP`
}
// Register implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Register(ctx context.Context, teacher *demo.Teacher) (resp *demo.RegisterResp, err error) {
	result := s.db.Table("teachers").Create(teacher2Model(teacher))
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(json.Marshal(resp))
	return
}

// Query implements the TeacherServiceImpl interface.
func (s *TeacherServiceImpl) Query(ctx context.Context, req *demo.QueryReq) (resp *demo.Teacher, err error) {
	var teacherRes Teacher

	result := s.db.Table("teachers").First(&teacherRes, req.Id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		resp = &demo.Teacher{}
		panic("the result of query is null")
	}else{
		resp = &demo.Teacher{
			Id : req.Id,
			Name : teacherRes.Name,
			College: &demo.College{
				Name: teacherRes.CollegeName,
				Address: teacherRes.CollegeAddress,
			},
			Email: strings.Split(teacherRes.Email, ","),
			Sex: teacherRes.Sex,
		}
	}
	return
}

func (s *TeacherServiceImpl) InitDB() {
    db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
    if err != nil {
       panic(err)
    }
    // drop table
    db.Migrator().DropTable(Teacher{})
    // create table
    err = db.Migrator().CreateTable(Teacher{})
    if err != nil {
       panic(err)
    }
    s.db = db
}

func teacher2Model(teacher *demo.Teacher) *Teacher{
	teacherData, err := json.Marshal(teacher)
	if err != nil {
		panic(err)
	}

	fmt.Println("teacher2model: " + string(teacherData))
	teacherInfo := &Teacher{
		Id : teacher.Id,
		Name: teacher.Name,
		CollegeName: teacher.College.Name,
		CollegeAddress: teacher.College.Address,
		Email : strings.Join(teacher.Email, ","),
		Sex: teacher.Sex,
	}
	return teacherInfo
}

//实现GenericCall接口
func (s *TeacherServiceImpl)GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error){
	if method == "Register"{
		reqStr, ok := request.(string)
		if !ok {
			return nil, errors.New("Invalid request type, cannot transfer it to json string")
		}

		fmt.Println(reqStr)

		var req demo.Teacher
		err = json.Unmarshal([]byte(reqStr), &req)
		if err != nil {
			panic("反序列化错误")
		}

		resp, err := s.Register(ctx, &req)
		if err != nil{
			panic("get register response")
		}
		respData, err := json.Marshal(resp)
		if err != nil {
			panic("get respData failed")
		}
		respStr := string(respData)
		fmt.Println(respStr)

		return respStr, nil
	}else{
		fmt.Println("query!!!")
		reqStr, ok := request.(string)
		if !ok {
			return nil, errors.New("Invalid request type, cannot transfer it to json string")
		}

		fmt.Println(reqStr)

		var req demo.QueryReq
		err = json.Unmarshal([]byte(reqStr), &req)
		if err != nil {
			panic("反序列化错误")
		}
		resp, err := s.Query(ctx, &req)
		if err != nil{
			panic("get register response")
		}

		respData, err := json.Marshal(resp)
		if err != nil {
			panic("get respData failed")
		}
		respStr := string(respData)
		fmt.Println(respStr)

		return respStr, nil
	}

	//序列化
	// respData, err := json.Marshal(resp)
	// if err != nil {
	// 	panic("get respData failed")
	// }
	// respStr := string(respData)
	// fmt.Println(respStr)

	// return respStr, nil
}