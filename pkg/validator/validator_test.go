package validator

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func validHTMLFile() []byte {
	f, err := os.Open("./fixtures/valid.html")
	if err != nil {
		panic("Cannot open fixtures")
	}
	data, _ := ioutil.ReadAll(f)

	return data
}

func invalidHTMLFile() []byte {
	f, err := os.Open("./fixtures/invalid.html")
	if err != nil {
		panic("Cannot open fixtures")
	}
	data, _ := ioutil.ReadAll(f)

	return data
}

func validCSSFile() []byte {
	f, err := os.Open("./fixtures/valid.css")
	if err != nil {
		panic("Cannot open fixtures")
	}
	data, _ := ioutil.ReadAll(f)

	return data
}

func TestValidator_ValidateBytes(t *testing.T) {
	type args struct {
		content     []byte
		contentType string
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		{
			name: "Test valid HTML file",
			args: args{
				content:     validHTMLFile(),
				contentType: "text/html; charset=utf-8",
			},
			wantErr: false,
			want: Response{
				Messages: []ResponseMessage{},
			},
		},
		{
			name: "Test invalid HTML file",
			args: args{
				content:     invalidHTMLFile(),
				contentType: "text/html; charset=utf-8",
			},
			wantErr: false,
			want: Response{
				Messages: []ResponseMessage{
					ResponseMessage{
						Type:         "error",
						LastLine:     7,
						LastColumn:   7,
						Message:      "Element “head” is missing a required instance of child element “title”.",
						Extract:      "ie=edge\">\n</head>\n<body",
						HiliteStart:  10,
						HiliteLength: 7,
					},
					ResponseMessage{
						Type:         "error",
						LastLine:     9,
						LastColumn:   11,
						Message:      "Element “title” not allowed as child of element “body” in this context. (Suppressing further errors from this subtree.)",
						Extract:      "body>\n    <title>Docume",
						HiliteStart:  10,
						HiliteLength: 7,
					},
				},
			},
		},
		{
			name: "Test valid CSS file",
			args: args{
				content:     validCSSFile(),
				contentType: "text/css; charset=utf-8",
			},
			wantErr: false,
			want: Response{
				Messages: []ResponseMessage{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewValidator()
			got, err := v.ValidateBytes(tt.args.content, tt.args.contentType)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validator.ValidateBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validator.ValidateBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
