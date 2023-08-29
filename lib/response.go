package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	ACCESS_TOKEN = "EAADKZAAJHDtEBOzaZBWRstZCXzxpBeFZAzdpZCI7s20riY3Q8ZAMEXt97twTZABDDDl8wg7zsHRzAWVGQNM7NaYcu6hW24ZBf718UjJ7CQaZAp5OwZAQM9PgzC9UEVXpSiR6L59aFQNoIpsT3zui1QjCJGkEyZC4yqWDDc0Lb9NcFUTK2Tl5Rxe52ALcHntd5D9wQZDZD"
	FACEBOOK_API = "https://graph.facebook.com/v17.0/116527961541255/messages?access_token=%s"
	VERIFY_TOKEN = "test"
)

func handleMessage(event Messaging) {
	msgText := strings.TrimSpace(event.Message.Text)
	msgText = strings.ToLower(msgText)
	//thanksSentiment = event.
	var response Message
	feedbackScreens := event.Messaging_Feedback.Feedback_Screens
	if (len(feedbackScreens) > 0) {
		for _,fs := range event.Messaging_Feedback.Feedback_Screens {
			q := fs.Questions.Question
			rating := q.Payload
			follow_up := q.Follow_Up.Payload
			log.Println("Received rating " + rating + " with follow-up " + follow_up)
			feedback_store = append(feedback_store, Feedback{rating, follow_up})
			response = Message{Text: "Thank you for your review!"}
		}
	} else {
		switch msgText {
		case "thanks", "ty", "thank you":
			response = buildFeedbackTemplate()
		default:
		}
	}

	SendResponse(event, response)
}

func buildFeedbackTemplate() Message {
	return Message{
		Attachment: &Attachment{
			Type: "template",
			Payload: Payload {
				Template_Type: "customer_feedback",
				Title: "Rate your exeperience with Bottled Water.",
				Subtitle: "Let Bottled Water know how they are doing.",
				Button_Title: "Rate Experience",
				Feedback_Screens: []Feedback_Screen{
					Feedback_Screen {
						Questions: []Question{
							Question {
								ID: "q1",
								Type: "csat",
								Title: "How yould you rate your experience?",
								Score_Label: "neg_pos",
								Score_Option: "five_stars",
								Follow_Up: Follow_Up {
									Type: "free_form",
									Placeholder: "Give additional feedback",
								},
							},
						},
					},
				},
				Business_Privacy: Business_Privacy{
					URL: "https://www.example.com",
				},
			},
		},
	}
}

func SendResponse(event Messaging, responseMessage Message) {
	client := &http.Client{}
	response := Response{
		Recipient: User{
			ID: event.Sender.ID,
		},
		Message: responseMessage,
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)
	url := fmt.Sprintf(FACEBOOK_API, ACCESS_TOKEN)//os.Getenv("PAGE_ACCESS_TOKEN"))
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan error, 1)
    go func() {
        resp, err := client.Do(req)
        ch <- err
		defer resp.Body.Close()
    }()
    select {
		case err := <-ch:
			if err != nil {
				log.Fatal(err)
			}
		case <-time.After(5 * time.Second):
			log.Fatal("Timed out")
    }
}