package leetcodeapi

type LeetcodeMeta struct {
	session   string
	csrfToken string
}

var leetcodeMeta = LeetcodeMeta{
	session:   "",
	csrfToken: "",
}

func SetCredentials(session string, csrfToken string) {
	leetcodeMeta.session = session
	leetcodeMeta.csrfToken = csrfToken
}

func RemoveCredentials() {
	leetcodeMeta.session = ""
	leetcodeMeta.csrfToken = ""
}
