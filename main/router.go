package main

import (
	"SockGo/CryptoUtils"
	"SockGo/ShellUtils"
	"SockGo/VoteUtils"
	"SockGo/paillier"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var PaillierPrivateKey *paillier.PrivateKey // 保存本地paillier公私钥
var Tickets []VoteUtils.BallotTicket        // 保存投票人上传的选票
var NameList []string
var IntroductionList []string

func main() {
	http.Handle("/paillierKeys/pub/", http.StripPrefix("/paillierKeys/pub/", http.FileServer(http.Dir("../paillierKeys/pub"))))

	http.Handle("/css/img/", http.StripPrefix("/css/img/", http.FileServer(http.Dir("../css/img/"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css"))))

	http.Handle("/mod/", http.StripPrefix("/mod/", http.FileServer(http.Dir("../mod"))))

	http.HandleFunc("/login", Login)

	http.HandleFunc("/register", Register)

	http.HandleFunc("/entryCandidateInfo", EntryCandidateInfo)

	http.HandleFunc("/recvCandidateInfo", RecvCandidateInfo)

	http.HandleFunc("/", Index)

	http.HandleFunc("/index", Index) // 首页

	http.HandleFunc("/init", Init) // 初始化系统，包括删除系统中已经有的paillier公钥,将NameList和IntroductionList清空

	http.HandleFunc("/ticket", SendTickets)

	http.HandleFunc("/recvTicket", RecvTicket)

	http.HandleFunc("/statistic", StatisticTickets)

	http.HandleFunc("/downloadPaillierPublicKey", DownloadPaillierPublicKey) // 用户必须先下载公钥，再上传选票之前必须上传本公钥到服务器

	http.HandleFunc("/createPublicKey", CreatePublicKey) // 只有公证人才有访问的权利，使用本函数之前要对公证人做身份验证
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	files, _ := template.ParseFiles("../mod/login.html")
	files.Execute(w, "")
}

func Register(w http.ResponseWriter, r *http.Request) {
	files, _ := template.ParseFiles("../mod/register.html")
	files.Execute(w, "")
}

func EntryCandidateInfo(w http.ResponseWriter, r *http.Request) {
	files, _ := template.ParseFiles("../mod/entryCandidateInfo.html")
	files.Execute(w, "")
}

func RecvCandidateInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	NameList = make([]string, 0)
	IntroductionList = make([]string, 0)
	//fmt.Println(r.PostFormValue("Name"))
	form := r.PostForm
	for k, v := range form {
		if k == "Name" {
			for _, name := range v {
				//fmt.Printf("v is %T : %v:\n", name, name)
				NameList = append(NameList, name)
			}
		} else {
			for _, intro := range v {
				IntroductionList = append(IntroductionList, intro)
			}
		}
	}
	fmt.Println("姓名", NameList)
	fmt.Println("简介", IntroductionList)

	files, _ := template.ParseFiles("../mod/index.html")
	files.Execute(w, "设置选票成功")
}

func Init(w http.ResponseWriter, r *http.Request) {
	ShellUtils.GetOutFromStdout("rm ../paillierKeys/pri/*") // 清空paillier密钥文件
	ShellUtils.GetOutFromStdout("rm ../paillierKeys/pub/*")
	NameList = make([]string, 0)
	IntroductionList = make([]string, 0)

	fmt.Println("初始化完毕")
	files, _ := template.ParseFiles("../mod/index.html")
	files.Execute(w, "初始化完毕")
}

func CreatePublicKey(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/createSuccess.html", "../mod/index.html")
	if err != nil {
		return
	}
	CryptoUtils.CreateKeys(1024)
	err = files.Execute(w, "创建Paillier公钥成功")
	if err != nil {
		fmt.Println("创建失败，请查看原因:", err)
		return
	}
}

func DownloadPaillierPublicKey(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/downloadPaillierPublicKey.html")
	if err != nil {
		return
	}
	err = files.Execute(w, "")
	if err != nil {
		return
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/index.html") // 这里应该加载菜单
	if err != nil {
		fmt.Println("加载模版错误:", err)
		return
	}
	err = files.Execute(w, "")
	if err != nil {
		fmt.Println("Execute err:", err)
		return
	}
}

func SendTickets(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/ticket.html") // 加载选票

	//can1 := VoteUtils.Candidate{
	//	ID:           "1",
	//	Name:         "李为君",
	//	Introduction: "20",
	//}
	//can2 := VoteUtils.Candidate{
	//	ID:           "2",
	//	Name:         "何俭涛",
	//	Introduction: "20",
	//}
	//can3 := VoteUtils.Candidate{
	//	ID:           "3",
	//	Name:         "徐许越",
	//	Introduction: "20",
	//}
	//can4 := VoteUtils.Candidate{
	//	ID:           "4",
	//	Name:         "闵浩哲",
	//	Introduction: "20",
	//}
	cans := make([]VoteUtils.Candidate, 0)
	for i := 0; i < len(NameList); i++ {
		can := VoteUtils.Candidate{}
		can.SetCandidateInfo(NameList[i], IntroductionList[i])
		cans = append(cans, can)
	}
	if err != nil {
		fmt.Println("加载模版失败:", err)
		return
	}
	err = files.Execute(w, cans)
	if err != nil {
		return
	}
}

func RecvData(w http.ResponseWriter, r *http.Request) {
	length := r.ContentLength
	body := make([]byte, length)
	read, err := r.Body.Read(body)
	if err != nil {
		if err == io.EOF {
			fmt.Println("读取完毕:", read)
		} else {
			return
		}
	}
	bodyStr := string(body)
	fmt.Println("Data:", bodyStr)

	//obj, err := regexp.Compile(`.*?name="(?P<name>op.*?)"`)
	//if err != nil {
	//	fmt.Println("编译正则失败:", err)
	//	return
	//}
	//res := make([]byte, 0)
	//for _, s := range obj.FindAllSubmatchIndex(body, -1) {
	//	res = obj.Expand(res, []byte("$name"), body, s)
	//	fmt.Println(string(res))
	//	res = make([]byte, 0)
	//}
	w.Header().Set("Location", "/mod/paillier")
	w.WriteHeader(302)
} // 使用缓存接收整个数据包

func RecvTicket(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}
	TicketData := r.PostFormValue("Ticket")
	//fmt.Println(Ticket)
	TicketJson := []byte(TicketData)
	//fmt.Println(TicketJson)
	Ticket := VoteUtils.BallotTicket{}
	err = json.Unmarshal(TicketJson, &Ticket)
	if err != nil {
		fmt.Println("转换为Ticket对象失败:", err)
	}
	//fmt.Println(Ticket)
	Tickets = append(Tickets, Ticket)
	//res := new(big.Int).SetInt64(0)
	//cRes, _ := paillier.Encrypt(&PaillierPrivateKey.PublicKey, res.Bytes())
	//for i := 0; i < Ticket.CandidateNum; i++ {
	//	cRes = paillier.AddCipher(&PaillierPrivateKey.PublicKey, cRes, Ticket.Option[i])
	//}
	//fmt.Println("调试：相加结果:", cRes)
	//decrypt, err := paillier.Decrypt(PaillierPrivateKey, cRes)
	//if err != nil {
	//	return
	//}
	//fmt.Println("解密的数字:", new(big.Int).SetBytes(decrypt).String())

}

func StatisticTickets(w http.ResponseWriter, r *http.Request) {
	PaillierPrivateKey = CryptoUtils.GetKeysFromJson("../paillierKeys/pri/key")
	fmt.Println("解析Paillier公钥成功:")
	fmt.Println("打印调试信息:")
	for i, v := range Tickets {
		fmt.Printf("[%d : %v]\n", i, v.CandidateNameList)
	}
}
