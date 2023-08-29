package lib

type Callback struct {
	Object string `json:"object,omitempty"`
	Entry  []struct {
		ID        string      `json:"id,omitempty"`
		Time      int         `json:"time,omitempty"`
		Messaging []Messaging `json:"messaging,omitempty"`
	} `json:"entry,omitempty"`
}

type Messaging struct {
	Sender    User    `json:"sender,omitempty"`
	Recipient User    `json:"recipient,omitempty"`
	Timestamp int     `json:"timestamp,omitempty"`
	Message   Message `json:"message,omitempty"`
	Messaging_Feedback Messaging_Feedback `json:"messaging_feedback,omitempty"`
}

type User struct {
	ID string `json:"id,omitempty"`
}

type Message struct {
	MID        string `json:"mid,omitempty"`
	Text       string `json:"text,omitempty"`
	QuickReply *struct {
		Payload string `json:"payload,omitempty"`
	} `json:"quick_reply,omitempty"`
	Attachments *[]Attachment `json:"attachments,omitempty"`
	Attachment  *Attachment   `json:"attachment,omitempty"`
}

type Attachment struct {
	Type    string  `json:"type,omitempty"`
	Payload Payload `json:"payload,omitempty"`
}

type Response struct {
	Recipient User    `json:"recipient,omitempty"`
	Message   Message `json:"message,omitempty"`
}

type Messaging_Feedback struct {
	Feedback_Screens []Feedback_Screen_Response `json:"feedback_screens,omitempty"`
}

type Feedback_Screen_Response struct {
	Questions Question_Response `json:"questions,omitempty"`
}

type Question_Response struct {
	Question Question `json:"q1,omitempty"`
}

type Payload struct {
	Template_Type string `json:"template_type,omitempty"`
	Title string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	Button_Title string `json:"button_title,omitempty"`
	Feedback_Screens []Feedback_Screen `json:"feedback_screens,omitempty"`
	Business_Privacy Business_Privacy `json:"business_privacy,omitempty"`
}

type Feedback_Screen struct {
	Questions []Question `json:"questions,omitempty"`
}

type Question struct {
	ID string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	Title string `json:"title,omitempty"`
	Score_Label string `json:"score_label,omitempty"`
	Score_Option string `json:"score_option,omitempty"`
	Follow_Up Follow_Up `json:"follow_up,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type Business_Privacy struct {
	URL string `json:"url,omitempty"`
}

type Follow_Up struct {
	Type string `json:"type,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Payload string `json:"payload,omitempty"`
}

type Feedback struct {
	Rating int
	Follow_Up string
	Timestamp int64
}

var feedback_store []Feedback

type Input struct {
	message string `json:"message,omitempty"`
}

type Resp struct {
	isReview bool `json:"isReview,omitempty"`
	shouldAskForReview bool `json:"shouldAskForReview,omitempty"`
	reviewScore int `json:"reviewScore,omitempty"`
}

type Datapoint struct {
	x int64
	y int64
}

type Interval struct {
	feedback []Feedback
}
//intervals []Interval