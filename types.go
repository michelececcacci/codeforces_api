package codeforces

type JudgeProtocol struct {
	Protocol string `json:"protocol"`
	Manual   string `json:"manual"`
	Verdict  string `json:"verdict"`
}

type Problem struct {
	ContestID int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Points    float64  `json:"points"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}

type ProblemStatistic struct {
	ContestID   int    `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}

type User struct {
	Country                 string `json:"country"`
	City                    string `json:"city"`
	LastName                string `json:"lastName"`
	LastOnlineTimeSeconds   int    `json:"lastOnlineTimeSeconds"`
	Rating                  int    `json:"rating"`
	FriendOfCount           int    `json:"friendOfCount"`
	TitlePhoto              string `json:"titlePhoto"`
	Handle                  string `json:"handle"`
	Avatar                  string `json:"avatar"`
	FirstName               string `json:"firstName"`
	Contribution            int    `json:"contribution"`
	Organization            string `json:"organization"`
	Rank                    string `json:"rank"`
	MaxRating               int    `json:"maxRating"`
	RegistrationTimeSeconds int    `json:"registrationTimeSeconds"`
	MaxRank                 string `json:"maxRank"`
}

type BlogEntry struct {
	OriginalLocale          string   `json:"originalLocale"`
	AllowViewHistory        bool     `json:"allowViewHistory"`
	CreationTimeSeconds     int      `json:"creationTimeSeconds"`
	Rating                  int      `json:"rating"`
	AuthorHandle            string   `json:"authorHandle"`
	ModificationTimeSeconds int      `json:"modificationTimeSeconds"`
	ID                      int      `json:"id"`
	Title                   string   `json:"title"`
	Locale                  string   `json:"locale"`
	Tags                    []string `json:"tags"`
}

type Member struct {
	Handle string `json:"handle"`
}

type Hacker struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	Ghost            bool     `json:"ghost"`
	Room             int      `json:"room"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}

type Defender struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	Ghost            bool     `json:"ghost"`
	Room             int      `json:"room"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}

type ContestHack []struct {
	ID                  int           `json:"id"`
	CreationTimeSeconds int           `json:"creationTimeSeconds"`
	Hacker              Hacker        `json:"hacker"`
	Defender            Defender      `json:"defender"`
	Verdict             string        `json:"verdict"`
	Problem             Problem       `json:"problem"`
	JudgeProtocol       JudgeProtocol `json:"judgeProtocol"`
}

type RatingChange struct {
	ContestID               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

// used to parse away the status, and isolate the result
type ResultWrapper[T any] struct {
	Status string `json:"status"`
	Result T      `json:"result"`
}

type Contest struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Type                string `json:"type"`
	Phase               string `json:"phase"`
	Frozen              bool   `json:"frozen"`
	DurationSeconds     int    `json:"durationSeconds"`
	StartTimeSeconds    int    `json:"startTimeSeconds"`
	RelativeTimeSeconds int    `json:"relativeTimeSeconds"`
}

type Party struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	Ghost            bool     `json:"ghost"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}

type ProblemResult struct {
	Points                    float64 `json:"points"`
	RejectedAttemptCount      int     `json:"rejectedAttemptCount"`
	Type                      string  `json:"type"`
	BestSubmissionTimeSeconds int     `json:"bestSubmissionTimeSeconds,omitempty"`
}

type Row struct {
	Party                 Party           `json:"party"`
	Rank                  int             `json:"rank"`
	Points                float64         `json:"points"`
	Penalty               int             `json:"penalty"`
	SuccessfulHackCount   int             `json:"successfulHackCount"`
	UnsuccessfulHackCount int             `json:"unsuccessfulHackCount"`
	ProblemResults        []ProblemResult `json:"problemResults"`
}

type ContestStandings struct {
	Contest  Contest   `json:"contest"`
	Problems []Problem `json:"problems"`
	Rows     []Row     `json:"rows"`
}

type Author struct {
	ContestID        int      `json:"contestId"`
	Members          []Member `json:"members"`
	ParticipantType  string   `json:"participantType"`
	Ghost            bool     `json:"ghost"`
	StartTimeSeconds int      `json:"startTimeSeconds"`
}

type ContestStatus struct {
	ID                  int     `json:"id"`
	ContestID           int     `json:"contestId"`
	CreationTimeSeconds int     `json:"creationTimeSeconds"`
	RelativeTimeSeconds int64   `json:"relativeTimeSeconds"`
	Problem             Problem `json:"problem"`
	Author              Author  `json:"author"`
	ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict             string  `json:"verdict"`
	Testset             string  `json:"testset"`
	PassedTestCount     int     `json:"passedTestCount"`
	TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
}

type Problemset struct {
	Problems          []Problem          `json:"problems"`
	ProblemStatistics []ProblemStatistic `json:"problemStatistics"`
}

type Comment struct {
	ID                  int    `json:"id"`
	CreationTimeSeconds int    `json:"creationTimeSeconds"`
	CommentatorHandle   string `json:"commentatorHandle"`
	Locale              string `json:"locale"`
	Text                string `json:"text"`
	Rating              int    `json:"rating"`
	ParentCommentID     int    `json:"parentCommentId,omitempty"`
}
