package VoteUtils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type Candidate struct { // 候选人
	ID           string
	Name         string
	Introduction string //候选人自我介绍
}

func (c *Candidate) SetCandidateInfo(Name string, Introduction string) {
	c.Name = Name
	c.Introduction = Introduction
	b, err := rand.Int(rand.Reader, new(big.Int).SetInt64(9999999999))
	if err != nil {
		return
	}
	c.ID = "Candidate_" + fmt.Sprintf("%s", b)
}
