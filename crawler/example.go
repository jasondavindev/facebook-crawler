package crawler

import (
	"fmt"
	"os"
)

// CheckURL search for pending comments example
func CheckURL() {
	url := "https://www.facebook.com/plugins/feedback.php?info=put_your_iframe_url_here"
	user := "YOUR_OFFICIAL_USER_ID"
	hasPendingComments, err := VerifyFacebookPage(url, user)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(hasPendingComments)
}
