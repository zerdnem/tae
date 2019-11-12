package utils

func HashType(hash string) string {
	var hashtype string
	switch len(hash) {
	case 32:
		hashtype = "md5"
	case 40:
		hashtype = "sha1"
	case 64:
		hashtype = "sha256"
	case 96:
		hashtype = "sha384"
	case 128:
		hashtype = "sha512"
	default:
		hashtype = ""
	}
	return hashtype
}
