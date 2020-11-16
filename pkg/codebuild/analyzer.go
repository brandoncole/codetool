package codebuild

import (
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

type (
	BuildStatistics struct {
		BuildsCount             int64
		BuildDurationInSeconds  int64
		PhaseDurationsInSeconds map[string]int64
	}
	Metric struct {
		Name       string
		Value      int64
		Percentage float32
	}
	Metrics []Metric
)

func NewBuildStatistics() *BuildStatistics {
	return &BuildStatistics{
		PhaseDurationsInSeconds: make(map[string]int64),
	}
}

func GetBuildStatistics(region string) (*BuildStatistics, error) {

	result := NewBuildStatistics()

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSessionWithOptions(
		session.Options{SharedConfigState: session.SharedConfigEnable},
	)
	if nil != err {
		return nil, err
	}

	// Create CodeBuild service client
	svc := codebuild.New(sess)
	var token *string

	// Get the list of builds
	for {

		listBuildsInput := &codebuild.ListBuildsInput{
			NextToken: token,
			SortOrder: aws.String("ASCENDING"),
		}
		names, err := svc.ListBuilds(listBuildsInput)
		if nil != err {
			return nil, err
		}

		getBuildsInput := &codebuild.BatchGetBuildsInput{Ids: names.Ids}
		builds, err := svc.BatchGetBuilds(getBuildsInput)
		if nil != err {
			return nil, err
		}

		for _, build := range builds.Builds {
			result.BuildsCount++
			result.BuildDurationInSeconds += int64(build.EndTime.Sub(*build.StartTime) / time.Second)
			for _, phase := range build.Phases {
				if nil == phase {
					continue
				}
				if duration, found := result.PhaseDurationsInSeconds[*phase.PhaseType]; found {
					result.PhaseDurationsInSeconds[*phase.PhaseType] = duration + *phase.DurationInSeconds
				} else {
					if nil != phase.DurationInSeconds {
						result.PhaseDurationsInSeconds[*phase.PhaseType] = *phase.DurationInSeconds
					}
				}
			}
		}

		token = names.NextToken
		if nil == token {
			break
		}

	}

	return result, nil

}

func (s *BuildStatistics) SortedMetrics() Metrics {

	var metrics Metrics

	for k, v := range s.PhaseDurationsInSeconds {
		metrics = append(metrics, Metric{
			Name:       k,
			Value:      v,
			Percentage: float32(v) / float32(s.BuildDurationInSeconds) * 100.0,
		})
	}

	sort.Sort(metrics)

	return metrics

}

func (l Metrics) Len() int           { return len(l) }
func (l Metrics) Less(i, j int) bool { return l[i].Value > l[j].Value }
func (l Metrics) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
