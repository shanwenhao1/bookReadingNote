package school

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

/*
	题目链接: https://www.nowcoder.com/test/question/3897c2bcc87943ed98d8e0b9e18c4666?pid=260145&tid=48545660

	老师想知道从某某同学当中，分数最高的是多少，现在请你编程模拟老师的询问。当然，老师有时候需要更新某位同学的成绩


		输入描述:
		输入包括多组测试数据。
		每组输入第一行是两个正整数N和M（0 < N <= 30000,0 < M < 5000）,分别代表学生的数目和操作的数目。
		学生ID编号从1编到N。
		第二行包含N个整数，代表这N个学生的初始成绩，其中第i个数代表ID为i的学生的成绩
		接下来又M行，每一行有一个字符C（只取‘Q’或‘U’），和两个正整数A,B,当C为'Q'的时候, 表示这是一条询问操作，他询问ID从A到B（包括A,B）的学生当中，成绩最高的是多少
		当C为‘U’的时候，表示这是一条更新操作，要求把ID为A的学生的成绩更改为B。

		输出描述:
		对于每一次询问操作，在一行里面输出最高成绩.

		输入例子1:
			5 7
			1 2 3 4 5
			Q 1 5
			U 3 6
			Q 3 4
			Q 4 5
			U 4 5
			U 2 9
			Q 1 5

		输出例子1:
			5
			6
			5
			9
*/

type Student struct {
	grade []int
}

func (stu *Student) change(stuId int, grade int) error {
	if stuId > len(stu.grade) || stuId < 1 {
		return errors.New("无该学生")
	}
	stu.grade[stuId-1] = grade
	return nil
}

func (stu *Student) handle(action string, par ...int) (int, error) {
	if strings.ToUpper(action) == "Q" {
		if len(par) != 2 {
			return 0, errors.New("parameter error")
		}
		if par[0] > par[1] || par[1] > len(stu.grade) {
			return 0, errors.New("parameter error")
		}
		// 截取所选段
		var midSlice = make([]int, par[1]-par[0]+1)
		copy(midSlice, stu.grade[par[0]-1:par[1]])
		// 排序
		sort.Slice(midSlice, func(i, j int) bool {
			return midSlice[i] > midSlice[j]
		})
		fmt.Println(stu.grade)
		fmt.Println(midSlice)
		return midSlice[0], nil
	} else if strings.ToUpper(action) == "U" {
		if par[0] < 1 || par[0] > len(stu.grade) || par[1] < 0 {
			return 0, errors.New("parameter error")
		}
		stu.grade[par[0]-1] = par[1]
		return 0, nil
	} else {
		return 0, errors.New("parameter error")
	}
}

func GradeRun() {
	var (
		stuNum int
		stuAct int
		sGrade int
	)
	var stu = new(Student)
	fmt.Println("请输入学生数量及要操作的次数: ")
	fmt.Scanf("%d %d", &stuNum, &stuAct)
	fmt.Println("请输入学生分数")
	var stuGradeS = make([]int, 0, stuNum)
	for i := 0; i < stuNum; i++ {
		fmt.Scanf("%d", &sGrade)
		stuGradeS = append(stuGradeS, sGrade)
	}
	//copy(stu.grade, stuGradeS)
	stu.grade = stuGradeS

	var (
		act  string
		par1 int
		par2 int
	)
	for i := 0; i < stuAct; i++ {
		fmt.Println("请输入操作及行为, 示例Q 1 5")
		fmt.Scanf("%s %d %d", &act, &par1, &par2)
		result, err := stu.handle(act, par1, par2)
		if err != nil {
			panic(err)
		}
		if result != 0 {
			fmt.Println(fmt.Sprintf("最高分为: %d", result))
		}
	}
}
