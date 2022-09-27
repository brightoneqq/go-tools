package cons

var MsgFlags = map[int]string{
	SUCCESS:              "ok",
	ERROR:                "fail",
	NO_BASIC_ENV_SETTING: "NO BAIC ENV SETTING",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
