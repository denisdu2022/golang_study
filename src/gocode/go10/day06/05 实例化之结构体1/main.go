package main

import "fmt"

type Student struct {
	sid, age     int
	name, gender string
	course       []string
}

func CourseInit(stu Student) {
	stu.course = []string{"chinese", "math", "english"}
	fmt.Println(stu)
}

func CourseInit2(stu *Student) {
	(*stu).course = []string{"chinese", "math", "english"}
}

func main() {

	//案例1:
	s1 := Student{sid: 1001, name: "daerwen", age: 23}
	s2 := s1                //值拷贝
	fmt.Println("s2: ", s2) //s2:  {1001 23 daerwen  []}
	s1.age = 100
	fmt.Println(s2.age) //23

	//如果希望s3的值跟随s1保持一致
	s3 := &s1
	//s1.age = 100
	fmt.Println(s3)        //&{1001 100 daerwen  []}
	fmt.Println((*s3).age) //100

	//案例2
	var s4 = Student{sid: 1001, name: "dafenqi", age: 18}
	CourseInit(s4)
	fmt.Println("s4 报的课程: ", s4.course) //s4 报的课程:  []  初始化不成功
	var s5 = &Student{sid: 1001, name: "bijiasuo", age: 20}
	CourseInit2(s5)
	//fmt.Println("s5报的课程: ", *s5.course) //*s5.course的写法是错误的
	fmt.Println("s5报的课程: ", (*s5).course) //s5报的课程:  [chinese math english]
	fmt.Println("s5报的课程: ", s5.course)    //s5报的课程:  [chinese math english]

}
