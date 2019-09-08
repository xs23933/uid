package uid

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Tx struct {
	ID UID    `json:"id"`
	OX string `json:"ox"`
}

type x []byte

func Test_uid(t *testing.T) {
	b := make([]UID, 0)
	for i := 0; i < 100; i++ {
		b = append(b, New())
	}
	fmt.Printf("%v\n", b)
}

func Test_parse(t *testing.T) {
	buf := []byte(`["bn3wqd23rqx8","tq4edzas4wai","d4g1gbi3b0j6","f3dfqe6m99p2","pofuzqhzp2vr","5d7b87wnwxu4","e5bnxohduhwa","p3lj2i9ue2sa","j7uj7maklyi5","bgphy1358aer","bnhzlbdz34m1","yhqmr83ldx18","69mhhj82pib2","bnu9laynuz2w","4vsd97t5d6zp","blimnjuefoem","r84jythxdf24","bbp4gscd7pql","jsyw2tq6pyel","z8vd0584cbz2","obesamlfmwnp","bgyotvkh7kva","bnvrmryhjtzu","6llcp7ku6tlb","gg1clq8a3c4e","bahvfa0v4cjf","836deazezn38","bm4nskvfkpcr","76lrerj8xuzb","gk7xwsmefv45","bleh0x6943ag","546p69taa8m4","blw7waml3h00","bfssdjtiths5","fy8dh8m6dz6c","wddzq5b13eyl","bej0l0bhl169","lrem25cz2n2e","xgvg05unvps0","bpachawf7a0q","wibqno4bupdk","ewkfremc5oey","6y8y99nci2nd","546svj119319","sqku8onlkmjr","umfhmhhw8o4f","yd5fswt3wsz5","zdi2dj9tto0e","pbmkhfr0h7kz","x8l0s9lcas0g","2f6yil9f5f2s","oxrxi3501j34","bengp78g728m","0z2b5k9wyc2p","bj85jhad0452","x5sbqq5ao1jh","bhc7rgxoelu3","v9gw6u7ot8kx","bbf0p61ix5ng","4xh1xwws54xd","g5rjehnr7u39","k39en1kvhba8","700sa63s05cg","bfrbc4pjgi3l","vv6fcw0sa499","24pwi1ic0v8k","blk9bixsjui4","boptx5zdgo83","97hrstsm0j7f","bdlbc6ho8cbr","rs5j9umrzfvc","8eq6kcemhplx","bh8lj7pyx7iv","fpmhd2b5183n","bdnguunc9xuo","mggtpw9y6sku","bfqqm6q8xx6w","bblywa3qsrzj","r3706vqk6sd8","90ysbcctqtyg","5m4y73e0926p","bhuavs4qxkh5","u2oyb4n4v39v","ipc0z2x9d6ds","bdpid2o68atm","padmiak6lfta","9ewj58kkm4on","9sie02znp0ev","txynrcnc36zo","bdamk6f21hfb","zl4k6k4xvmzw","oosp2ga07kk6","vq6dtdj0mm9b","la2i643b39r0","bbeuyv59h5nx","bmqejg5kp705","bn2spqnvca6p","bk9wlcyst95x","blnvq6cke0tv","bbxwjl20gr6n"]`)
	x := make([]UID, 0)
	json.Unmarshal(buf, &x)
	for _, item := range x {
		fmt.Printf("%v\t", item)
	}
	fmt.Printf("%v", x)
}
