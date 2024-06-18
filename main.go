package main

//导入多个包，一行一个
import (
	"errors"
	"fmt"
	"maps"
	"math"
	"strconv"
	"time"
	"unicode/utf8"
)

// 枚举定义
const (
	StateIdle = iota //iota表示未定义整数，默认为:0
	StateConnected
	StateError
	StateRetrying
	StateRunning
)

var stateMap = map[int]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
	StateRunning:   "Running",
}

// geometry :　定义接口
type geometry interface {
	area() float64  //函数名  返回类型
	perim() float64 //函数名  返回类型
}

// 定义结构area
type rect struct {
	width, height float64 //同类成员写一行，逗号隔开 / 成员名 类型
}

// 定义结构circle
type circle struct {
	radius float64 //成员名 类型
}

// area() : 实现接口方法area
func (a rect) area() float64 {
	return a.width * a.height
}

// perim ：实现接口方法perim
// func (参数名 类型) 函数名() 返回类型
func (a rect) perim() float64 {
	return 2*a.width + 2*a.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 封装接口调用方法
// func 函数名(参数名 类型)
func callInterfaceGeometry(g geometry) {
	fmt.Println(g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

// 同类开多个参数默认写法
func add2(a int, b int) int {
	return a + b
}

// 同类开多个参数简洁写法
func add3(a, b, c int) int {
	return a + b + c
}

// 多个返回值写法
func add(a int, b int) (int, int) {
	return a * 5, b * 5
}

// 多个返回值写法
func abc(a int, b int) (int, int, int) {
	return a, b, a + b
}

// 可变参数写法
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println("计算结果：", total)
	return total
}

// 闭包函数使用
func myFunc() func() int {
	i := 0
	return func() int {
		i++
		print(i)
		return i
	}
}

// 递归相乘
func fib(n int) int {
	if n == 0 {
		return 1
	}
	return n * fib(n-1)
}

// 递归前两个数相加
func fib1(n int) int {
	if n < 2 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

// 变量传参测试
func byVal(v int) { //作为参数传入是 999
	v = 0 //初始化变量 ，进入这个作用域后，v变成0 ，退出这个使用域后，会是999 ，没有修改到地址中的值
}

func byPointer(vp *int) { //作为参数传入是 999
	*vp = 0 //初始化指针 进入这个作用域后，v变成0  退出这个使用域后，也是 0 ，因为地址中的值已修改
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

// 结构定义
type person struct {
	name string
	age  int
}

// getName：实现结构方法getName
func (p person) getName() string {
	return p.name
}

// getAge：实现结构方法getAge
func (p person) getAge() int {
	return p.age
}

// 返回新的结构地址
func newPerson(name string) *person {
	p := person{name: name} //使用结构
	p.age = 18
	return &p //返回结构地址
}

// 结构定义
type VideoFrame struct {
	id   int
	head []byte
	len  int64
	data []byte
}

// 生成结构字段的get与set方法
// ================================
func (v *VideoFrame) Id() int {
	return v.id
}

func (v *VideoFrame) SetId(id int) {
	v.id = id
}

func (v *VideoFrame) Head() []byte {
	return v.head
}

func (v *VideoFrame) SetHead(head []byte) {
	v.head = head
}

func (v *VideoFrame) Len() int64 {
	return v.len
}

func (v *VideoFrame) SetLen(len int64) {
	v.len = len
}

func (v *VideoFrame) Data() []byte {
	return v.data
}

func (v *VideoFrame) SetData(data []byte) {
	v.data = data
}

//================================

// 接口定义
type myInterface interface {
	getVideoFrame() *VideoFrame
	FrameCount() int
}

//实现接口方法

func (v VideoFrame) getVideoFrame() *VideoFrame {
	mybyte := make([]byte, 32)
	copy(mybyte, "hello")
	vf := VideoFrame{
		id:   1111,
		head: mybyte,
		len:  128,
		data: []byte{01, 23, 45, 67, 89},
	}
	return &vf
}

func (v VideoFrame) FrameCount() int {
	return 999
}

// 定义基础结构
type base struct {
	num int //结构成员
}

func (b *base) Num() int {
	return b.num
}

func (b *base) SetNum(num int) {
	b.num = num
}

// 在新结构中嵌入结构
type container struct {
	base        //嵌入的结构
	str  string //新结构字段
}

// describe:base结构实现方法descirbe
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 生成器(泛型)
// any表示空接口interface{}
// comparable表示一个可比较的接口，实现了该接口的对象可使用==或者!=进行对象比较
// map[K]V 是根据传入的类型生成map对象
// []K 表示返回传入类型的集合
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m)) //根据map大小分配[]K内存并返回r
	//遍历map对象m并添加键到集合
	for k := range m {
		r = append(r, k) //添加到集合
	}
	return r //返回集合
}

// 泛型结构实现
type element[T any] struct {
	next *element[T] //下一元素指针
	val  T           //元素值
}

// 泛型队列实现
type List[T any] struct {
	head, tail *element[T] //队头元素与队尾元素指针
}

// 泛型队列-入队实现
func (lst *List[T]) Push(v T) {
	if lst.tail == nil { //空队
		lst.head = &element[T]{val: v} //队头
		lst.tail = lst.head            //队尾
	} else {
		lst.tail.next = &element[T]{val: v} //向队尾，入队元素
		lst.tail = lst.tail.next            //队尾指向该才入队的元素
	}
}

// 取队列中所有元素，并返回元素列表
func (lst *List[T]) GetAll() []T {
	var elems []T //元素列表
	//遍历队列，从队头到队尾
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val) //添加到列表
	}
	return elems //返回列表
}

func myFunEmitError(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return -1, errors.New("传入的参数a,b的值必须大于0")
	}
	return a + b, nil
}

// 定义错误
var custError = fmt.Errorf("这是一个格式化的错误信息：错误码：%v", 999)

// 在函数中返回错误
func throwErr() error {
	return custError //返回错误
}

// 自定义错误结构
type custErrorStruct struct {
	arg int    //错误码
	msg string //错误信息
}

// Error：实现结构方法Error
func (e custErrorStruct) Error() string {
	return fmt.Sprintf("自定义错误结构：错误码：%d 错误信息: %s", e.arg, e.msg) //格式化错误信息
}

// 在函数中抛出自定义错误结构
func testCustErr(i int) (int, error) {
	if i < 42 {
		return i, custErrorStruct{arg: 1001, msg: "输入数值小于42"}
	}
	return i, nil
}

func main() {
	fmt.Println("Go接口使用示例")
	//初始化结构，供接口调用
	r := rect{width: 5, height: 5}
	c := circle{radius: 5}
	//调用接口，传入结构
	callInterfaceGeometry(r)
	callInterfaceGeometry(c)
	//使用枚举
	s := 4
	switch s {
	case StateIdle:
		fmt.Println("StateIdle:", 0)
		break
	case StateConnected:
		fmt.Println("StateConnected:", 1)
		break
	case StateError:
		fmt.Println("StateError:", 2)
		break
	case StateRetrying:
		fmt.Println("StateRetrying:", 3)
		break
	case StateRunning:
		fmt.Println("StateRunning:", 4)
		break
	default:
		fmt.Println("StateIdle:", -1)
		break
	}
	fmt.Println("状态:", stateMap[s])

	//maps使用
	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 1
	m["c"] = 2
	m["d"] = 3
	m["e"] = 4
	m["f"] = 5
	m["g"] = 6
	m["h"] = 7
	fmt.Println("map:", m, "map长度:", len(m))
	//遍历map,无序输出
	for k, v := range m {
		fmt.Println("键：", k, "值：", v)
	}
	fmt.Println("取map中指定键f的值:", m["f"])
	v, ret := m["f"]
	fmt.Println("值:", v, "结果：", ret)
	delete(m, "f")
	fmt.Println("删除map中键为f的元素后:", m, "map长度:", len(m))
	clear(m)
	fmt.Println("调用clean函数清空map后:", m, "map长度:", len(m))
	//直接初始化map
	m1 := map[string]int{"a1": 1, "a2": 2}
	fmt.Println(m1["a1"], m1["a2"], m1, len(m1))
	m2 := map[string]int{"a1": 1, "a2": 2}
	//map对象比较
	tmp := ""
	if maps.Equal(m1, m2) {
		tmp = "YES"
	} else {
		tmp = "NO"
	}
	fmt.Println("map m1与m2是相等的：", tmp)
	//遍历map
	for k, v := range m1 {
		fmt.Println("键：", k, "值：", v)
	}

	//数组初始化
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	//数组遍历
	for k, v := range nums {
		fmt.Println("of nums: ===> ", "key:", k, "value:", v)
	}

	stringMap := make(map[string]string)
	stringArr := []string{"中国", "美国", "日本"}
	stringArr1 := []string{"ZH", "US", "JP"}
	for i, v := range stringArr1 {
		stringMap[v] = stringArr[i]
	}
	for k, v := range stringMap {
		fmt.Println(k, v)
	}

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fv := myFunc() //调用闭包函数,返回函数
	fv()           //调用返回的函数
	fmt.Println("递归函数调用-相乘：", fib(5))
	fmt.Println("递归函数调用-相加：", fib1(5))

	//变量声明与使用
	var x1 = 123
	x2 := 456
	fmt.Println(x1, x2)
	var v1, v2, v3 = 2017, 2019, 2022
	fmt.Println(v1, v2, v3)
	v4, v5, v6 := "VS2017", "VS2019", "VS2022"
	fmt.Println(v4, v5, v6)
	//常量声明与使用
	const YEAR, MONTH, WEEK = 12, 30, 7
	print("\n", YEAR, MONTH, WEEK, "\n")
	//循环使用
	fmt.Println(">>>>>>>>外部变量循环>>>>>>>>>>>>>")
	i := 10
	for i <= 20 {
		fmt.Println(i)
		i++
	}
	fmt.Println(">>>>>>>>常规循环>>>>>>>>>>>>>")
	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	fmt.Println("**********范围循环**********")
	for num := range 10 {
		fmt.Println(num)
	}

	//条件中使用多语句
	names := []string{"Apple", "Banana"}
	if name := "Apple"; name == names[0] {
		fmt.Println("===》相等")
	}
	//if多个条件判断
	if a := 10; a > 15 {
		fmt.Println("a>15", a)
	} else if a > 12 {
		fmt.Println("a>12,a<15", a)
	} else {
		fmt.Println("a<12", a)
	}

	//switch语句使用
	//匹配数值条件
	sw := 3
	switch sw {
	case 1:
		fmt.Println("one", sw)
		break
	case 2:
		fmt.Println("two", sw)
		break
	case 3:
		fmt.Println("three", sw)
		break
	default:
		fmt.Println("unknow", sw)
		break
	}
	//匹配日期条件
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("休息日")
		break
	case time.Friday:
		fmt.Println("工作日最后一天")
		break
	default:
		fmt.Println("工作日")
		break
	}
	//匹配时间条件
	wh := time.Now().Hour()
	switch wh {
	case 5:
		if wm := time.Now().Minute(); wm >= 30 {
			fmt.Println("准备下班...")
		}
		break
	default:
		fmt.Println("工作中...")
		break
	}
	//函数在表达式中使用
	typeChk := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool:", t)
		case int:
			fmt.Println("int:", t)
		default:
			fmt.Println("not a type:", t)
		}
	}
	typeChk(255)
	typeChk(!false)
	typeChk(m)

	//声明数组不初始化
	var arr1 [5]int
	//初始化数组
	arr1[0] = 100
	//取数组元素
	arr1_elm_0 := arr1[0]
	fmt.Println(arr1_elm_0)

	//声明数组并直接初始化
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	//声明可变长数组并初始化,数组大由初始化的元素个数决定
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr3, len(arr3))
	for k, v := range arr3 {
		fmt.Println(k, v)
	}
	//声明可变数组，长度由初始化值决定
	b := [...]int{100, 3: 400, 500, 9: 1000} //指定第3个元素值为400
	fmt.Println("idx:", b)
	for i, v := range b {
		fmt.Println("===>指定元素索引：", i, v)
	}
	//二维数组声明
	var arr2d [2][3]int
	//二维数组初始化
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr2d[i][j] = i + j + 1
		}
	}
	//遍历二维数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("===>输出二维数组元素值:[", i, j, "]", arr2d[i][j])
		}
	}
	//声明二维数组直接初始化
	arr2d_new := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	//遍历二维数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("******>输出二维数组元素值:[", i, j, "]", arr2d_new[i][j])
		}
	}

	//字符串数组使用，也叫分片(Slices)
	var s1 []string
	fmt.Println("for s1:", s1, len(s1), cap(s1), s1 == nil, len(s1) == 0)
	if s1 == nil {
		fmt.Println("空字符串数组(Slices)")
	}
	if len(s1) == 0 {
		fmt.Println("字符串数组(Slices)的大小为0")
	}
	//使用make为Slice分配内存
	s2 := make([]string, 3)
	fmt.Println("for s2:", s2, len(s2), cap(s2), s2 == nil, len(s2) == 0)
	//赋值Slice
	for i := 0; i < len(s2); i++ {
		s2[i] = strconv.Itoa((i + 1) * 5) //转换为字符串，并存储到元素中
	}
	//遍历
	for _, v := range s2 {
		fmt.Println("===> value of Slices : ", v)
	}
	//使用append向Slice中添加5个元素
	for i := 0; i < 5; i++ {
		s2 = append(s2, strconv.Itoa((i+1)*5+15))
	}
	fmt.Println("成功向Slices中添加元素后：", s2)
	//再次遍历
	for k, v := range s2 {
		fmt.Println("k,v in s2: ===> ", k, v)
	}
	//复制Slices副本
	s2_copy := make([]string, len(s2)) //根据s2大小分配新内存
	copy(s2_copy, s2)                  //复制s2到s2_copy
	fmt.Println(s2_copy)
	//遍历s2_copy
	for k, v := range s2_copy {
		fmt.Println("k,v in s2_copy: ===> ", k, v)
	}
	newSlice := s2[2:5] //取索引为2开始到索引为5结结束的所有元素并返回新Slices
	fmt.Println("取索引为2开始到索引为5结束的所有元素并返回新Slices:", newSlice)
	from2toend := s2[2:] //取索引为2开始,所有元素
	fmt.Println("取索引为2开始,所有元素:", from2toend)
	from0to5 := s2[:5] //取索引为0开始,索引为5结束的所有元素并返回新Slices
	fmt.Println("取索引为0开始,索引为5结束的所有元素并返回新Slices", from0to5)

	//使用make为二维数组分配内存
	new2darr := make([][]int, 5)
	for i := 0; i < 5; i++ {
		innerLen := i + 1
		new2darr[i] = make([]int, innerLen) //动态创建数组
		for j := 0; j < innerLen; j++ {
			new2darr[i][j] = i + j + 1 //赋值
		}
	}
	fmt.Println(new2darr)

	//map操作
	map2 := make(map[string]int) //先分配内存
	//再按KEY赋值
	map2["one"] = 1
	map2["two"] = 2
	map2["three"] = 3
	//直接声明并初始化键值对
	//map[键类型]值类型
	map1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(map1)
	//遍历
	for k, v := range map1 {
		fmt.Println("===>map1 顺序输出:", k, v)
	}
	for k, v := range map2 {
		fmt.Println("===>map2 乱序输出:", k, v)
	}
	map3 := map[string]string{"zh": "中国", "us": "美国", "jp": "日本"}
	for k, v := range map3 {
		fmt.Println(k, v)
	}

	val1 := 999
	fmt.Println("变量val1原始值:", val1)
	fmt.Println("变量val1原始地址:", &val1)           //输出变量地址
	byVal(val1)                                 //传入变量 ,因为变量在外部，所以值 没有被改变 些时val1值还是999
	fmt.Println("byVal函数初始化后，val1值:", val1)     //输出999
	byPointer(&val1)                            //传入变量地址后，会将地址址初始化为0,些时val1值为0
	fmt.Println("变量val1原始地址:", &val1)           //输出变量地址
	fmt.Println("byPointer函数初始化后，val1值:", val1) //输出0

	//字符串编码使用
	const sss = "สวัสดี"
	fmt.Println("Len:", len(sss))
	for i := 0; i < len(sss); i++ {
		fmt.Printf("%x ", sss[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(sss))
	for idx, runeValue := range sss {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(sss); i += w {
		runeValue, width := utf8.DecodeRuneInString(sss[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}

	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))

	mystr := "你好，中国"
	fmt.Println("bytes =", len(mystr))
	for i := 0; i < len(mystr); i++ {
		fmt.Printf("%x ", mystr[i])
	}
	fmt.Println("\nrunes =", utf8.RuneCountInString(mystr))

	//使用结构
	fmt.Println(person{name: "HELLO", age: 99})
	fmt.Println(newPerson("NICK"))
	//直接初始化结构
	cat := struct {
		name  string
		isAni bool
	}{"Tom", true}
	fmt.Println(cat, &cat) //使用结构
	//初始化结构
	person1 := person{name: "Tom", age: 99}
	//使用结构方法
	fmt.Println(person1.getName(), person1.getAge())

	//使用接口方法
	vf := VideoFrame{}
	fmt.Println("===> FrameCount:", vf.FrameCount())
	fmt.Println("===>VideoFrame:", vf.getVideoFrame())
	//使用结构方法
	fmt.Println("===>VideoFrame->ID:", vf.Id())
	vf.SetId(8888)
	fmt.Println("===>VideoFrame-ID:", vf.Id())

	//使用结构base与container
	cb := container{ //初始化结构container
		base: base{ //初始化结构base
			num: 888, //base结构成员num
		},
		str: "Hello Container", //container结构成员str
	}
	fmt.Println("结构Container:", cb)
	fmt.Printf("结构Container对象cb={container->base->num: %v, container->str: %v}\n", cb.num, cb.str)
	fmt.Println("结构container内嵌结构base的num成员:", cb.base.num)
	fmt.Println("结构base的成员方法describe:container->base->describe()", cb.describe())

	//临时定义接口
	type describer interface {
		describe() string //跳到base结构的describe方法
	}
	fmt.Println("--->describe:", cb.describe()) //调用内嵌结构的describe方法
	cb.SetNum(999)
	var d describer = cb                        //调用base结构的SetNum
	fmt.Println("===>describer:", d.describe()) //调用接口describer，接口再转到base结构对接口的实现describe方法

	//泛型队形测试
	//1.测试MapKeys函数
	var mmm = map[int]string{1: "2", 2: "4", 4: "8"}
	//map[int]string 相当于MapKeys中的[K comparable, V any],此时K=int,V=string
	//下面执行MapKeys(mmm)后会返回一个整形数组[]int ,
	//因为mmm类型是map[int]string 所以返回[]K相当于返回[]int
	fmt.Println("keys for int:", MapKeys(mmm)) //输入包含map键的集合,键为整数
	//再次测试
	//map[string]int ---> MapKeys中的[K comparable, V any]
	//K ---> string    V ---> int
	//[]K ---> []string
	var mmm1 = map[string]int{"one": 2, "two": 4, "four": 8}
	fmt.Println("keys for string:", MapKeys(mmm1)) //输入包含map键的集合,键为字符串

	//构造map
	mapStrStr := map[string]string{"zh": "中国", "us": "美国", "jp": "日本"}
	//指定泛型输入类型
	mapstrstrKEY := MapKeys[string, string](mapStrStr)
	fmt.Println("mapStrStr:", mapStrStr, "mapstrsrKey:", mapstrstrKEY)

	//使用泛型函数时直接初始化map
	makeKeysWithInit := MapKeys[int, string](map[int]string{
		1: "中国",
		2: "美国",
		3: "日本",
	})
	fmt.Println("makeKeysWithInit:", makeKeysWithInit)

	//泛型队列使用
	lst := List[int]{}                 //初始化整数队列
	lst.Push(10)                       //入队
	lst.Push(20)                       //入队
	lst.Push(30)                       //入队
	fmt.Println("list:", lst.GetAll()) //输出队列所有元素

	//泛型队列使用
	lst_str := List[string]{}                      //初始化整数队列
	lst_str.Push("zh")                             //入队
	lst_str.Push("us")                             //入队
	lst_str.Push("jp")                             //入队
	fmt.Println("lst_str list:", lst_str.GetAll()) //输出队列所有元素

	//入队自定义结构
	baseQueue := List[base]{} //队列初始化
	for i := range 10 {
		tmp := base{num: (i + 1) * 5} //初始化自定义结构
		baseQueue.Push(tmp)           //入队
	}
	fmt.Println("baseQueue-LIST:", baseQueue.GetAll())

	//错误处理
	err := errors.New("程序出错错误了") //直接创建错误输出对象
	fmt.Println("err:", err)
	//在函数中返回错误信息
	result, errmgs := myFunEmitError(-1, 2)
	//如果函数处理出错，输出错误信息
	if errmgs != nil {
		fmt.Println("错误信息：===> :", errmgs, "值:", result)
	}
	//使用格式库处理错误信息
	anErrA := fmt.Errorf("这是一个格式化的错误信息：错误码：%v", 999)
	fmt.Println("anErrA", anErrA)
	//错误信息内嵌
	anErrBIncludeA := fmt.Errorf("包含错误信息anErrA的内容: %w", anErrA)
	fmt.Println("anErrBIncludeA:===>", anErrBIncludeA)
	//错误信息判断
	errB := fmt.Errorf("这是一个格式化的错误信息：错误码：%v", 999)
	fmt.Println("errB:", &errB, "anErrA", &anErrA)
	if errors.Is(anErrA, errB) { //两个错误信息地址不一样
		fmt.Println("错误信息匹配")
	} else {
		fmt.Println("===>错误信息不匹配")
	}
	//判断函数中返回的错误
	if errors.Is(custError, throwErr()) {
		fmt.Println("错误信息匹配")
	}

	//定义临时函数
	checkValue := func(i int) (int, error) {
		if i < 10 {
			return 0, errors.New("输出值必须大于0")
		} else {
			return i * 5, nil
		}
	}
	//遍历数组
	for _, i := range []int{7, 42} {
		//调用临时函数并返回值与错误
		if r, e := checkValue(i); e != nil {
			fmt.Println("===>值检测不通过:", e) //出错
		} else {
			fmt.Println("PASS:值检测通过:", r) //正确
		}
	}

	//测试自定义错误结构
	rettmp, errtmp := testCustErr(38)
	if errtmp != nil {
		fmt.Println(errtmp.Error(), rettmp)
	}
}
