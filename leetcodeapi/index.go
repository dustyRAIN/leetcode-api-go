package leetcodeapi

var credentials = leetcodeMeta{
	session:   "",
	csrfToken: "",
}

func SetCredentials(session string, csrfToken string) {
	credentials.session = session
	credentials.csrfToken = csrfToken
}

func RemoveCredentials() {
	credentials.session = ""
	credentials.csrfToken = ""
}
