package terraformanalyzer

import (
	"context"
	"fmt"
	"io"
	"regexp"

	"github.com/meyskens/lookout-html-analyzer/pkg/validator"

	"gopkg.in/meyskens/lookout-sdk.v0/pb"

	"gopkg.in/src-d/go-log.v1"
)

// this regex checks if the file is a css file
var cssNameRegex = regexp.MustCompile(`\.css$`)

type Analyzer struct {
	DataClient pb.DataClient
	Version    string
}

func (a Analyzer) NotifyReviewEvent(ctx context.Context, review *pb.ReviewEvent) (*pb.EventResponse, error) {
	log.Infof("got review request %v", review)

	changes, err := a.DataClient.GetChanges(ctx, &pb.ChangesRequest{
		Head:            &review.Head,
		Base:            &review.Base,
		WantContents:    true,
		WantLanguage:    false,
		WantUAST:        false,
		ExcludeVendored: true,
		IncludePattern:  `\.(css|htm|html)$`,
	})

	if err != nil {
		log.Errorf(err, "GetChanges from DataServer failed")
		return nil, err
	}

	var comments []*pb.Comment
	hadFiles := map[string]bool{}

	v := validator.NewValidator()

	for {
		change, err := changes.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Errorf(err, "GetChanges from DataServer failed")
			return nil, err
		}

		log.Infof("analyzing '%s'", change.Head.Path)

		if change.Head == nil {
			log.Infof("ignoring deleted '%s'", change.Base.Path)
			continue
		}

		if _, hasAnalyzed := hadFiles[change.Head.Path]; hasAnalyzed {
			log.Infof("ignoring already analyzed '%s'", change.Head.Path)
			continue
		}

		hadFiles[change.Head.Path] = true

		contentType := "text/html; charset=utf-8" // default to HTML
		if cssNameRegex.MatchString(change.Head.Path) {
			contentType = "text/css; charset=utf-8"
		}
		resp, err := v.ValidateBytes(change.Head.Content, contentType)
		if err != nil {
			comments = append(comments, &pb.Comment{
				File: change.Head.Path,
				Line: 0,
				Text: fmt.Sprintf("Validator errored on validating:\n%s", err),
			})
			continue
		}

		for _, message := range resp.Messages {
			comments = append(comments, &pb.Comment{
				File: change.Head.Path,
				Line: int32(message.LastLine),
				Text: message.Message,
			})
		}

	}

	return &pb.EventResponse{AnalyzerVersion: a.Version, Comments: comments}, nil
}

func (a Analyzer) NotifyPushEvent(context.Context, *pb.PushEvent) (*pb.EventResponse, error) {
	return &pb.EventResponse{}, nil
}
