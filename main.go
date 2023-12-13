package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://leetcode.com/graphql/"
	payload := []byte(`{
		"query":"query questionOfToday {activeDailyCodingChallengeQuestion {date userStatus link question {acRate difficulty freqBar frontendQuestionId: questionFrontendId isFavor paidOnly: isPaidOnly status title titleSlug hasVideoSolution hasSolution topicTags {name id slug}}}}",
		"variables":{}
	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("authority", "leetcode.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-TW,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "__stripe_mid=cc5fc0b0-67b3-4165-b248-b95594dfd451387eee; csrftoken=RVVZZ3cV1z0Wl6CcVOnIl0F4AbGC3nehiTfNuNPdRP9iXCK77VuKucMTVPBVvtxX; LEETCODE_SESSION=YOUR_SESSION_COOKIE")
	req.Header.Set("origin", "https://leetcode.com")
	req.Header.Set("random-uuid", "902421e3-3df4-1156-928a-78a723aaa092")
	req.Header.Set("referer", "https://leetcode.com/problemset/all/?page=2")
	req.Header.Set("sec-ch-ua", `"Microsoft Edge";v="117", "Not;A=Brand";v="8", "Chromium";v="117"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sentry-trace", "381007ccbace4d63a90cc36e8b2e990d-b3fdba938bcc22f4-0")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.60")
	req.Header.Set("x-csrftoken", "RVVZZ3cV1z0Wl6CcVOnIl0F4AbGC3nehiTfNuNPdRP9iXCK77VuKucMTVPBVvtxX")

	type ResponseData struct {
		Data struct {
			ActiveDailyCodingChallengeQuestion struct {
				Question struct {
					Difficulty string `json:"difficulty"`
				} `json:"question"`
			} `json:"activeDailyCodingChallengeQuestion"`
		} `json:"data"`
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	difficulty := data.Data.ActiveDailyCodingChallengeQuestion.Question.Difficulty
	fmt.Println("LeetCode每日一题的難度為:", difficulty)
}
