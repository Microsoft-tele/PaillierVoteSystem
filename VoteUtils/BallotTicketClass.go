package VoteUtils

type BallotTicket struct {
	ID                string   //选票ID
	CandidateNum      int      //参选人数
	CandidateNameList []string //候选人列表
	Option            [][]byte // 选项
	RSAPublicKey      []byte   // RSA公钥，由投票者写入
	Signature         []byte   // 电子签名
}

//func (b *BallotTicket) InitBallotTicket(CandidateNum int, CandidateNameList []Candidate, PaillierPublicKey paillier.PublicKey) {
//	bigNum, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
//	if err != nil {
//		return
//	}
//	b.ID = "Notary_" + fmt.Sprintf("%s", bigNum)
//	b.Option = nil
//
//}
