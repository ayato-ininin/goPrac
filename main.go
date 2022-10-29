/*
アスタリスク（アステリスク）は古ギリシャ語で「星」って意味
*→スターって読むことがある。
() →パレンティス
[]→ブラケット
{}→ブレース括弧
| →バーティカルバー(垂直な線)
~ →チルダ
`` →バッククォート
^　→キャレット記号
*/

//gofmt -w main.go(これで正しい記法に変えてくれる)
//go install golang.org/x/tools/cmd/goimports@latest
//goimports -l -w .
//上記でgo のimportをきれいにするコマンドをインストール可

//コントロールRでデバッグ実行のショートカット

package main

//https://pkg.go.dev/std
//標準ライブラリのdocument
import (
	//アルファベット順
	"fmt"
	"log"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	Pi   = 3.14 // コンストは基本グローバルだそう。
	user = "A"
	pass = "###"
)

func main() {
	section3()
	section4()
	section5()
	section6()
}

/*
セクション3(定義)
*/
func section3()  {
	fmt.Println("section3====================================================")
		// msg := "hello world"
	// fmt.Println(msg, time.Now())
	// fmt.Println(user.Current())
	var (
		i    int     = 1   //デフォルトは0
		f64  float64 = 1.2 //デフォルトは0
		s    string  = "test"
		t, f bool    = true, false //デフォルトはfalse
	)
	fmt.Println(i, f64, s, t, f)
	fmt.Printf("%T\n", f64) //型をプリントできる(自分で改行入れる必要あり)
	fmt.Printf("%T\n", s)

	//２つとも同じ変数宣言
	//var i int = 1(varは関数外でも宣言可)
	//xi := 1(関数内のみ使用！、ショートな記法)→デフォルトはfloat64！float32とかにしたいなら、varが必要。

	numeric()//数値型
	stringPrac()//文字列型
	booleanPrac()//論理値型
	kataHenkan()//型変換
	hairetsu()//配列
	suraisu()//スライス
	suraisuMakeCap()//スライスのmakeとcap
	mapPrac()//map
	bytePrac()//バイト型
	functionPrac()//関数
	clojya()//クロージャ
	kahentyouHikisuu()//可変長引数
	question()//演習
}

/*
数値型について
int32とかint64というのは、bitのこと。
32bitか64bitか。
uintはunsigned long integer(符号なし長整数)
intはsigned integer(符号付き整数)
uint8→8bit→2のゼロ乗から2の7乗まで→符号なしなので、0スタートの合計255まで表示可！！
int8→-127~127(符号をまたいで255bit分になる！！)
*/
func numeric() {
	var (
		//goの記法として、一番長いとこに=とか型宣言をあわせるそう。
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	//値の型を%Tで表示し、値そのものは%vで表示。(godocに記載あり)
	fmt.Printf("type=%T value=%v\n", u8, u8)
}

func stringPrac() {
	fmt.Println(string("hello world"[0]))

	var s string = "xcodex"
	// s[0] = "T"
	s = strings.Replace(s, "x", "h", 1) //文字の置換
	fmt.Println(s)

	fmt.Println(`test
test`) //改行を残すなら、バッククォート
}

func booleanPrac() {
	t, f := true, false
	fmt.Printf("%T %v\n", t, t)
	fmt.Printf("%T %v\n", f, f)
}

func kataHenkan() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx) //%fでfloat表示

	var y float64 = 1.3
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy) //%dで整数表示

	var s string = "14"
	//z = int(s)//goはこんな感じで文字列から数値には変換ができない
	i, _ := strconv.Atoi(s) // Atoi→ascii to integerの略、Atoiはintとerrorを返すので_なければエラーがでてくる。アンスコは_にするとerrの引数は使わないということになる。
	/*
		asciiコードとは、コンピュータが使う文字と、その文字に割当てた番号の対応表
		コンピュータに「0」と「1」以外の文字を表現させるためのものが文字コード(0,1しか理解できないから。)
		fmt.Println(string("Hello world"[0]))
		これはasciiコードで72とでるが、Hが72番であるということ。
		例えば、英大文字の「A」は ASCIIコードでは2進数で1000001（16進数で0x41、10進数で65）で表現します
		10進数の0番から127番までの番号（2進数では0000000から1111111まで）に、128文字が割り当てられています。→8bit(1バイトで表される!!)
	*/
	fmt.Printf("%T %v\n", i, i) //%dで整数表示
}

/*
配列
*/
func hairetsu() {
	var a [2]int //２個のint配列を入れるという定義
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200} //直接いれる方法(2個ってサイズは変えれない。)
	fmt.Println(b)
}

/*
スライス([]の配列に長さを指定しない形。)
*/
func suraisu() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])

	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)
}

/*
スライスのメイクとキャパシティ
メモリに書き込む？？
*/
func suraisuMakeCap() {
	n := make([]int, 3, 5) //キャパシティは5入る配列で、3つ中身入れた配列を作る。
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	var c []int
	//c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}

// pythonの辞書型と似ている。
func mapPrac() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2) // 値がなければ0が帰ってきて、真偽値も受け取れる。２つのvalueがくる。

	m2 := make(map[string]int) //mapの初期化でメモリ上に作成。
	m2["PC"] = 5000
	fmt.Println(m2)
}

func bytePrac() {
	b := []byte{72, 73} //asciiコードから文字列を判定する
	fmt.Println(b)
	fmt.Println(string(b)) //stringでキャスト

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}

func functionPrac() {
	r1, r2 := add(10, 30)
	fmt.Println(r1, r2)
	r3 := cal(10, 30)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)
}

func add(x int, y int) (int, int) {
	return x + y, x - y
}

func cal(price int, item int) (result int) {
	result = price * item //:=で再定義はできない
	return result
}

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

/*
クロージャは、こんな感じでπが違う値になろうとも、
一つの関数で独立して返すことができるから優秀。
パイを自由な値にしつつ、面積を求める関数を使える。
*/
func clojya() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}

func foo(params ...int) {
	fmt.Println(len(params), params)
}

/*
可変長引数
何個引数が入るかわからない場合に、複数引数が入れれるような実装
*/
func kahentyouHikisuu() {
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)
}

func question() {
	f := 1.11
	fmt.Println(int(f))

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v\n", m, m)
}

func section4()  {
	fmt.Println("section4====================================================")
	ifPrac()
	rangePrac()
	switchPrac()
	deferPrac()//遅延実行
	logPrac()
	errorHandlePrac()
	//そもそもpanicを書くこと自体は推奨されてなくて、基本はエラーハンドルすること。panicは自分が何したらいいかわかってない状態であるからよくない。
	panicAndRecoverPrac()
	enshuuMondai()
}

func by2(num int) string {
	if num % 2 == 0{
		return "ok"
	} else {
		return "no"
	}
}

func ifPrac()  {
	num:=9
	if num % 2 == 0{
		fmt.Println("by 2")
	} else if num %3 == 0  {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x,y := 10,10
	if x ==10 && y == 10 {
		fmt.Println("true")
	}

	result := by2(10)
	if result == "ok"{
		fmt.Println("great")
	}

	//変数をしまうのと、条件分岐を簡略化できる。
	if result2 := by2(10); result2 == "ok"{
		fmt.Println("great")
	}
	//ただ、result2の使い方だと、この上のブレースの中でしか使えないから注意。

}

func rangePrac()  {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i,l[i])
	}

	//配列の中身を出すのは上記と同義
	for i,v := range l{
		fmt.Println(i,v)
	}

	//iのindexがいらない場合、中身だけ取り出す
	for _,v := range l{
		fmt.Println(v)
	}

	m:= map[string]int{"apple":100,"banana":200}

	for k,v := range m{
		fmt.Println(k,v)
	}

	//キーのみ取り出せる。
	for k := range m{
		fmt.Println(k)
	}

	//valueだけ取り出す。
	for _,v := range m{
		fmt.Println(v)
	}
}

func getOsName() string {
	return "mac"
}

func switchPrac()  {
	// os:= getOsName()
	// switch os {
	// case "mac":
	// 	fmt.Println("mac")
	// default:
	// 	fmt.Println("default")
	// }

	//基本定義は上記だが、osはここでしか使わないなら省略記法
	switch os:= getOsName() ; os {
	case "mac":
		fmt.Println("mac")
	default:
		fmt.Println("default", os)
	}
	//これは使えない！！
	//fmt.Println("default", os)

	t:= time.Now()
	fmt.Println(t.Hour())
	fmt.Printf("type=%T value=%v\n", t.Hour(),t.Hour())
	//条件を書かなくてもかける。
	// switch {
	// case t.Hour < 12:
	// 	fmt.Println("morning")
	// case t.Hour < 17:
	// 	fmt.Println("afternoon")
	// }
}

func foohello() {
	defer fmt.Println("world foo")//この関数の最後に実行する、という意味

	fmt.Println("hello foo")
}

//遅延実行
func deferPrac()  {
	// foohello()
	// defer fmt.Println("world")//この関数の最後に実行する、という意味

	// fmt.Println("hello")
	// fmt.Println("hello")

	// fmt.Println("run")
	// //スタッキングdeferといって、複数ある場合最初のdeferが最後に呼ばれる！！
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("success")

	//用途！！
	file, _ := os.Open("./main.go")//main.goのファイルを開く
	defer file.Close()//最後にちゃんと実行してくれるから、closeし忘れることもない！！
	data := make([]byte,100)//100バイトのバイト配列
	file.Read(data)
	fmt.Println(string(data))//stringにキャスト
}

const (
	Ldate         = 1 << iota  // 日付
	Ltime                      // 時刻
	Lmicroseconds              // 時刻のマイクロ秒
	Llongfile                  // ソースファイル（ディレクトリパスを含む）
	Lshortfile                 // ソースファイル（ファイル名のみ）
	LUTC                       // タイムゾーンに依らない UTC 時刻
	LstdFlags     = Ldate | Ltime  // 日付 (Ldata) と時刻 (Ltime) ：デフォルト
)

//ログをファイルに書き込むためのloggingの設定
func LoggingSettings(logFile string)  {
	//O_RDWR : ファイルの読み込みと書き込み両方
	//O_CREATE: ファイルがなければ作成
	//O_APPEND: 上書きではなく追記する
	//0666: ファイルのパーミッション、「rw-rw-rw-」なら「0666」、自分-グループ-他人
	//r:4(読み)、w:2(書き)、x:1(実行)
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	//Stdout→standard output(標準出力:コンソールに出る！！これがないと、ログファイルにしかかかれない。multiwriterは標準出力とログファイルの両方に書き込むという設定)
	multiLogFile := io.MultiWriter(os.Stdout,logfile)
//出力時の情報を追加で付加したい場合.
//定数はビットフラグで定義されているので、| でまとめて設定できます：
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)//出力先
}

func logPrac()  {
	LoggingSettings("test.log")//どのファイルに書くか

	_, err := os.Open("dddd")
	if err != nil {
		//log.Fatalln("Exit",err)
	}

	//goのloggingはあまりdebugとかinfoみたいに別れてるものがない。
	log.Println("logging")//よく使われるプリントライン
	log.Printf("%T %v","test","test")//型を出力できる

	//fatal系はos.Exit(1) を呼び出してプログラムを終了させます。　
	//log.Fatalf("%T %v","test","test")//型を出力できる
	//log.Fatalln("error")
	fmt.Println("ok")//okが表示されず、fatallnでコードが終了する！！
}

func errorHandlePrac()  {
	file, err := os.Open("./main.go")//initialize
	if err != nil {
		log.Fatal("Error")
	}
	defer file.Close()
	data := make([]byte,100)
	count,err := file.Read(data)//initialize(上のerrを上書き)
	if err != nil{
		log.Fatal("Error")
	}
	fmt.Println(count,string(data))

	//err := os.Chdir("test")initializeしようとするとエラー出るから
	//下記の用にoverrideする必要あり。
	err = os.Chdir("test")
	if err != nil {
		//log.Fatal("Error")
	}

	//エラーのみしか返さないなら、こんなふうに一行で書くこともできる。
	if err = os.Chdir("test") ; err != nil {
		//log.Fatal("Error")
	}
}

func thirdPratyConnectionDB(){
	//panicは自分で例外を作成できる！！
	//詳細なエラー分が標準出力されて強制終了する。
	panic("unable to connect database")
}

func save()  {
	defer func() {
		//エラーが出たあと、deferでrecoverをすることでパニックのメッセージはでるが、正常に動作が進む。
		s:= recover()
		fmt.Println(s)
	}()
	thirdPratyConnectionDB()
}

func panicAndRecoverPrac()  {
	save()
	fmt.Println("ok")
}

func enshuuMondai()  {
	l := []int{100,300,23,11,23,2,4,6,4}
	// sort.Sort(sort.IntSlice(l))
	// fmt.Println(l[0])
	var min int

	for i, num := range l{
		if i == 0{
			min = num
			continue
		}

		if min > num {
			min = num
		}
	}
	fmt.Println(min)

	m := map[string]int{
		"apple": 200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi": 90,
	}
	sum := 0
	//valueだけ取り出す。
	for _,v := range m{
		sum += v
	}
	fmt.Println(sum)
}

func section5()  {
	fmt.Println("section5====================================================")
	pointaPrac()//ポインタ
	pointaDifferenceNewAndMake()//makeとnewの違い
	structPrac()//構造体
}

func one(x *int)  {
	*x = 1//参照渡し(xというポインタの参照先に1を入れる。)
}

func pointaPrac()  {
	var n int = 100
	// fmt.Println(n)

	// fmt.Println(&n)//100を入れたメモリのアドレスになる！！
	// var p *int = &n//*がポインタ型のintにしている。
	// fmt.Println(p)//アドレス
	// fmt.Println(*p)//参照先

	one(&n)
	fmt.Println(n)
}

func pointaDifferenceNewAndMake()  {
	//値を入れる前に、ポインタが入るメモリだけ確保したい。
	var p *int = new(int)//これでポインタのメモリをnew
	fmt.Println(p)
	fmt.Println(*p)//0が初期値で入る。
	*p++//1

	var p2 *int//これはポインタを確保してるわけでないから、nilが返る。
	fmt.Println(p2)
	//*p2++//エラー

	//じゃあmapとかに使っていたmakeとnewは何が違うの？
	//そもそもmakeはslice, map, channelのみで、データ構造体を生成するために使う。
	s := make([]int,0)
	fmt.Printf("%T\n",s)

	m:= make(map[string]int)
	fmt.Printf("%T\n",m)

	ch:= make(chan int)//チャネル
	fmt.Printf("%T\n",ch)

	//newは指定した型のポインタ型を生成するために使う。
	//ほとんどの場合で，構造体（struct）のポインタを生成するためにしか使われません
	var pp *int = new(int)
	fmt.Printf("%T\n",pp)//*intが帰ってくる。

	var st = new(struct{})//ストラクト、*structが帰ってくる。
	fmt.Printf("%T\n",st)
}

type Vertex struct{
	//大文字じゃなくて小文字だとプライベートになって、小文字だと他のとこからアクセスできないから注意
	X int
	Y int//デフォルトは0
	S string//デフォルトは空
}

func changeVertex(v Vertex)  {
	v.X = 1000
}

func changeVertex2(v *Vertex)  {
	//(*v).X = 1000→こっちでもいい
	v.X = 1000//ストラクトの場合、勝手に実態を指すから、*はいらない。
}

func structPrac()  {
	v := Vertex{X:1, Y:2}
	fmt.Println(v)
	fmt.Println(v.X,v.Y)
	v.X = 100
	fmt.Println(v.X,v.Y)
	v2:= Vertex{X:1}
	fmt.Println(v2)

	v3:= Vertex{1,2,"test"}
	fmt.Println(v3)

	v4:= Vertex{}
	fmt.Println(v4)


	var v5 Vertex
	fmt.Println(v5)//nilではない

	//v6とv7は同じ意味で初期化できる！&のほうが、アドレス返るのわかりやすいからこっち使われやすいかも。
	v6 := new(Vertex)
	fmt.Printf("%T %v\n",v6,v6)//ポインタが返る

	v7 := &Vertex{}
	fmt.Printf("%T %v\n",v7,v7)//ポインタが返る

	//ssとsも同じ意味であるが、スライスとかであればmakeのほうが推奨されてるそう。
	ss:=make([]int,0)
	fmt.Println(ss)

	s:=[]int{}
	fmt.Println(s)

	advance()
}

func advance()  {
	v:=Vertex{1,2,"test"}
	changeVertex(v)//これは値渡しだから、中身が変わらない。
	fmt.Println(v)

	v2:= &Vertex{1,2,"test"}
	changeVertex2(v2)
	fmt.Println(v2)
	fmt.Println(*v2)
}

func section6()  {
	fmt.Println("section6====================================================")
	//まず、関数とメソッドは別物という認識！！
	//メソッドはクラスの中にある「操作」のこと。
	//ただgoにはクラスがないから、メソッドというものはないんやけど、
	//structがもつメソッドみたいに扱う技術が下記。
	methodPointaPrac()//メソッド
	//単語の意味としては埋め込み、
	//goには継承もないから、代わりにエンベデッドがある。
	embeddedPrac()
	non_struct_mehodPrac()//レアであまり使わないかも？
	interfacePrac()
	typeAssersionPrac()
	stringerPrac()//ストラクトの出力内容を変えれる
	errorPrac()//自分なりのエラーを作れる。
	question6()
}

/*pythonでメソッドやコンストラクタを書く例
class Vertex2(object):
	def __init__(self,x,y):
		self._x = x
		self._y = y

	def area(self):
		return self._x * self._y

	def scale(self, i):
		self._x = self._x *i
		self._y = self.y * i

v = Vertex2(3,4)
v.scale(10)
print(v.area())
*/


type Vertex2 struct {
	x, y int//同じパッケージないからでしかアクセスできない
}

// func Area(v Vertex2) int {
// 	return v.x * v.y
// }

//普通上記みたいにするが、下記みたいにするとオブジェクトのメソッド的に扱える。
//メソッドの作り方。
func (v Vertex2) Area() int {
	//値を渡す、値レシーバー
	return v.x * v.y
}

//メソッドの中でポインタを使用、変数の中身を書き換える
func (v *Vertex2) Scale(i int) {
	//ポインタを使用するポインタレシーバー
	v.x = v.x * i
	v.y = v.y * i
}

//コンストラクタ作成で、ストラクトをreturnする。
func New(x,y int) *Vertex2 {
	return &Vertex2{x,y}
}

func methodPointaPrac()  {
	//v := Vertex2{3,4}
	//fmt.Println(Area(v))

	v:= New(3,4)//コンストラクタで生成
	v.Scale(10)
	fmt.Println(v.Area())//メソッドを利用した書き方
}

//Vertex2の継承
type Vertex3D struct {
	Vertex2
	z int
}

func (v Vertex3D) Area3D() int {
	//値を渡す、値レシーバー
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

//コンストラクタ作成で、ストラクトをreturnする。
func New3D(x,y,z int) *Vertex3D {
	return &Vertex3D{Vertex2{x,y}, z}
}

func embeddedPrac()  {
	v:= New3D(3,4,5)
	v.Scale3D(10)
	fmt.Println(v.Area3D())
}

//structではなく、intの名前を変更するようなイメージ。
type MyInt int

//自分で定義したintに、メソッドを持たすことも可能。
func (i MyInt) Double() int {
	return int(i*2)//intでキャストする必要あり
}
func non_struct_mehodPrac()  {
	//あまり使わない記法かも。
	myInt := MyInt(10)
	fmt.Println(myInt.Double)
}

//interfaceは実際のコードではなく、型。関数名とかだけ定義
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p Person) Say() string{
	//p.Name = "MR." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

//interfaceのダックタイピング
func DriveCar(human Human) {
	if human.Say() == "MIKE" {
		fmt.Println("ok")
	} else {
		fmt.Println("get out")
	}
}

func interfacePrac()  {
	var mike Human = Person{"MIKE"}
	//interfaceがあるし、メソッドも呼び出しやすいかも
	mike.Say()
	DriveCar(mike)
}

//異なる方を渡すとき、下記のinterfaceの空にするとどんな型でも受け付ける。
func do(i interface{})  {
/*
	ii:= i.(int)//iをintにタイプアサーション(必須)
	ii *= 2
	fmt.Println(ii)

	ss := i.(string)//iをstringにタイプアサーション
	fmt.Println(ss+"!")
*/
	//switchとtypeはセットで、スイッチタイプ文。型の仕分け
	switch v:= i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v+"!")
	default:
		fmt.Printf("i dont know %T\n",v)
	}
}

//タイプアサーションとタイプの仕分け処理(スイッチタイプ文)
func typeAssersionPrac()  {
	//var i interface{} = 10 //この段階ではiはまだintではない！！interfaceの型
	do(10)
	do("mike")
	do(true)

	var i int = 10
	ii := float64(10)//これはタイプコンバージョン、タイプアサーションではない
	fmt.Println(i,ii)
}

type Person2 struct {
	Name string
	Age int
}

//出力の内容を変えれる！！これしないと、mikeの出力は{mike 22}となる。
//mikeのageを隠す事もできる。
func (p Person2) String() string {
	return fmt.Sprintf("my name is %v",p.Name)
}

func stringerPrac()  {
	mike := Person2{"Mike",22}
	fmt.Println(mike)
}

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v",e.Username)
}
func myFunc() error {
	ok := false
	if ok {
		return nil
	} else {
		//理由は、ポインタであれば同じエラー名でも違うエラーとして扱える。
		return &UserNotFound{Username:"mike"}//errorではポインタでやるべきだそう。
	}
	return nil
}

func errorPrac()  {
	//ポインタであれば同じエラー名でも違うエラーとして扱える。
	e1 := &UserNotFound{"mike"}
	e2 := &UserNotFound{"mike"}
	fmt.Println(e1==e2)//false
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}

type VertexPrac struct{
	X,Y int
}

func (v VertexPrac) Plus() int {
	return v.X + v.Y
}

func (v VertexPrac) String() string {
	return fmt.Sprintf("X is %d! Y is %d!",v.X,v.Y)
}

func question6()  {
	v:= VertexPrac{3,4}
	fmt.Println(v.Plus())
	fmt.Println(v)
}
