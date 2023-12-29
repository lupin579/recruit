package serializer

type RankSerializer struct {
	Uid            int    `json:"uid" db:"uid"`
	Name           int    `json:"uname" db:"uname"`
	Identidication string `json:"identification" db:"identification"`
	Times          int    `json:"succeedTimes" db:"succeed_times"`
}
