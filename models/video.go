package models

type Video struct {
	VideoID     string
	UserID      string
	Title       string
	Description string
	Category    string
	Tags        []string
}

func GetRecommendations(userID string, videos []Video, userVideos []Video) []Video {
	recommendations := []Video{}
	threshold := 0.5

	for _, userVideo := range userVideos {
		for _, video := range videos {
			if video.UserID != userID {
				categorySimilarity := jaccard([]string{userVideo.Category}, []string{video.Category})
				tagsSimilarity := jaccard(userVideo.Tags, video.Tags)
				averageSimilarity := (categorySimilarity + tagsSimilarity) / 2

				if averageSimilarity >= threshold {
					recommendations = append(recommendations, video)
				}
			}
		}
	}

	return recommendations
}

func jaccard(setA, setB []string) float64 {
	intersection := 0
	union := len(setA) + len(setB)
	for _, a := range setA {
		for i, b := range setB {
			if a == b {
				intersection++
				setB = append(setB[:i], setB[i+1:]...)
				break
			}
		}
	}
	union -= intersection
	return float64(intersection) / float64(union)
}
