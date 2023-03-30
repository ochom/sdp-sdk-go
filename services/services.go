package services

import "os"

const devURL = "https://dtsvc.safaricom.com:8480/api"
const prodURL = "https://dsvc.safaricom.com:9480/api"

func getURL() string {
	val := os.Getenv("DEPLOYMENT_MODE")

	if val == "production" {
		return prodURL
	}

	return devURL
}
