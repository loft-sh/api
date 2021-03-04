package token

type PrivateClaims struct {
	Loft Loft `json:"loft.sh,omitempty"`
}

type Loft struct {
	// The UID of the user that this token is for
	UID string `json:"uid,omitempty"`
	// The kubernetes name of the user that this token is signed for
	Name string `json:"name,omitempty"`
	// The generation of the token
	Gen int64 `json:"gen,omitempty"`
}
