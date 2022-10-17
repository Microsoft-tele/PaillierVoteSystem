package VoteUtils

type BallotTicket struct {
	ID            string      //选票ID
	CandidateNum  int         //参选人数
	CandidateList []Candidate //候选人列表
	Option        []int       // 选项
	Signature     string      // 电子签名
}
