package token

type PrivateClaims struct {
	Loft Loft `json:"loft.sh,omitempty"`
}

type Loft struct {
	UID  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
}
