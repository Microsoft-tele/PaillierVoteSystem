package main

import (
	"RemoteRouter/CryptoUtils"
	"RemoteRouter/ShellUtils"
	"RemoteRouter/VoteUtils"
	"RemoteRouter/paillier"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
)

var PaillierPrivateKey *paillier.PrivateKey // 保存本地paillier公私钥
var Tickets []VoteUtils.BallotTicket        // 保存投票人上传的选票
var NameList []string
var IntroductionList []string

type ShowResultMap struct {
	Name      string
	Res       string
	ResCipher []string
}

var ShowResultList []ShowResultMap

func main() {
	http.Handle("/paillierKeys/pub/", http.StripPrefix("/paillierKeys/pub/", http.FileServer(http.Dir("../paillierKeys/pub"))))

	http.Handle("/css/img/", http.StripPrefix("/css/img/", http.FileServer(http.Dir("../css/img/"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("../css"))))

	http.Handle("/mod/", http.StripPrefix("/mod/", http.FileServer(http.Dir("../mod"))))

	http.HandleFunc("/loginIndex", LoginIndex)

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

	http.HandleFunc("/showResult", ShowResult)

	http.HandleFunc("/selectVoter", SelectVoter)

	http.HandleFunc("/verifySignature", VerifySignature)

	http.HandleFunc("/sendVerifyCode", SendVerifyCode)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
}

func SendVerifyCode(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	fmt.Println("监测到发送验证码按钮：")
	data := r.URL.RawQuery
	fmt.Println(data)
	rawMail := strings.Split(data, "=")
	fmt.Println(rawMail)
	mail := rawMail[1]

	VerifyNum := rand.Intn(900000) + 99999

}

func ShowResult(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("../mod/statistic.html")
	if err != nil {
		fmt.Println("解析错误:", err)
	}
	files.Execute(w, ShowResultList)
}

func LoginIndex(w http.ResponseWriter, r *http.Request) {
	files, _ := template.ParseFiles("../mod/login.html")
	files.Execute(w, "")
}

func Login(w http.ResponseWriter, r *http.Request) {

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
}

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
	//fmt.Println("Ticket对象转换成功:")
	//fmt.Println(Ticket)
	Tickets = append(Tickets, Ticket)

	w.Header().Set("Location", "http://127.0.0.1:12345/index")
	w.WriteHeader(302)
}

type Statistic struct {
	Name         string   // 记录候选人姓名
	Options      [][]byte //记录候选人得分情况
	OptionsStr   []string // 转换为字符串
	CalResCipher []byte   // 相加结果
	CalResByte   []byte   // 解密结果
	CalResInt    string
}

var StatisticShell []Statistic

func StatisticTickets(w http.ResponseWriter, r *http.Request) {
	PaillierPrivateKey = CryptoUtils.GetKeysFromJson("../paillierKeys/pri/key")
	fmt.Println("解析Paillier公钥成功:")

	StatisticShell = make([]Statistic, len(NameList)) // 每一个候选人的统计结果
	for i := 0; i < len(NameList); i++ {
		StatisticShell[i].Name = NameList[i]
	}

	for i := 0; i < len(Tickets); i++ { // 整合选票
		for k, v := range Tickets[i].NameAndOption {
			for j := 0; j < len(NameList); j++ {
				if k == StatisticShell[j].Name {
					StatisticShell[j].Options = append(StatisticShell[j].Options, v)
				}
			}
		}
	}
	ShowResultList = make([]ShowResultMap, 0)
	for i := 0; i < len(StatisticShell); i++ { //每一个候选人的统计结果
		res := new(big.Int).SetInt64(0)
		cRes, _ := paillier.Encrypt(&PaillierPrivateKey.PublicKey, res.Bytes())
		for j := 0; j < len(StatisticShell[0].Options); j++ { // 两个人投票
			StatisticShell[i].OptionsStr = append(StatisticShell[i].OptionsStr, string(StatisticShell[i].Options[j]))
			cRes = paillier.AddCipher(&PaillierPrivateKey.PublicKey, cRes, StatisticShell[i].Options[j])
		}
		StatisticShell[i].CalResCipher = cRes                      // 加运算后的密文结果
		decrypt, err := paillier.Decrypt(PaillierPrivateKey, cRes) // 解密
		if err != nil {
			return
		}
		StatisticShell[i].CalResByte = decrypt                                // 记录解密结果
		StatisticShell[i].CalResInt = new(big.Int).SetBytes(decrypt).String() // 转换为int型字符串
		fmt.Printf("[%v]的得票情况:[%v]\n", StatisticShell[i].Name, StatisticShell[i].CalResInt)
		resultMap := ShowResultMap{
			Name:      StatisticShell[i].Name,
			Res:       StatisticShell[i].CalResInt,
			ResCipher: StatisticShell[i].OptionsStr,
		}
		ShowResultList = append(ShowResultList, resultMap)
	}
	files, _ := template.ParseFiles("../mod/index.html")
	files.Execute(w, "开始解析投票结果，请稍后在结果栏中查看结果")
}

type tmpStruct struct {
	ID        string
	RSA       string
	Signature string
}

func SelectVoter(w http.ResponseWriter, r *http.Request) {
	VerifyList := make([]tmpStruct, 0)
	for i := 0; i < len(Tickets); i++ {
		tmp := tmpStruct{
			ID:        Tickets[i].ID,
			RSA:       string(Tickets[i].RSAPublicKey),
			Signature: string(Tickets[i].Signature),
		}
		VerifyList = append(VerifyList, tmp)
	}
	files, err := template.ParseFiles("../mod/selectVoter.html")
	if err != nil {
		fmt.Println("解析失败:", err)
	}
	files.Execute(w, VerifyList)
}
func VerifySignature(w http.ResponseWriter, r *http.Request) {
	fmt.Println("进入验证界面")
	r.ParseForm()
	form := r.PostForm
	var VoterId string
	for k, v := range form {
		fmt.Printf("[%v : %v]\n", k, v)
		VoterId = k
	}
	Ticket := VoteUtils.BallotTicket{
		ID:            "",
		CandidateNum:  0,
		NameAndOption: nil,
		RSAPublicKey:  nil,
		Signature:     nil,
	}
	for i := 0; i < len(Tickets); i++ {
		if Tickets[i].ID == VoterId {
			Ticket = Tickets[i]
			break
		}
	}
	var pubKey rsa.PublicKey
	err := json.Unmarshal(Ticket.RSAPublicKey, &pubKey)
	if err != nil {
		fmt.Println("转换RSA公钥失败:", err)
		return
	}
	var name []string
	name = strings.Split(Ticket.ID, "_")
	name1 := name[len(name)-1]
	fmt.Println("要验证的姓名是:", name1)
	hashed := sha256.Sum256([]byte(name1))
	err = rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, hashed[:], Ticket.Signature)
	flag := 1
	if err != nil {
		fmt.Println("验证签名失败，本次投票可能产生错误，请公证人慎重选择：", err)
		flag = 0
	}

	files, _ := template.ParseFiles("../mod/index.html")
	if flag == 1 {
		fmt.Println("验证签名成功:")
		files.Execute(w, "验证签名成功")
	} else {
		files.Execute(w, "验证签名失败")
	}
}
