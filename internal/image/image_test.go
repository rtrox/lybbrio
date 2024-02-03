package image

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcessFile(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    ImageFile
		wantErr bool
	}{
		{
			name: "test1.jpg",
			file: "test_fixtures/test1.jpg",
			want: ImageFile{
				Size:        665,
				Width:       35,
				Height:      35,
				ContentType: "image/jpeg",
			},
			wantErr: false,
		},
		{
			name: "test2.png",
			file: "test_fixtures/test2.png",
			want: ImageFile{
				Size:        181,
				Width:       49,
				Height:      43,
				ContentType: "image/png",
			},
			wantErr: false,
		},
		{
			name: "test3.gif",
			file: "test_fixtures/test3.gif",
			want: ImageFile{
				Size:        1201,
				Width:       26,
				Height:      66,
				ContentType: "image/gif",
			},
			wantErr: false,
		},
		{
			name: "test4.jpeg",
			file: "test_fixtures/test4.jpeg",
			want: ImageFile{
				Size:        634,
				Width:       8,
				Height:      9,
				ContentType: "image/jpeg",
			},
			wantErr: false,
		},
		{
			"broken.jpg",
			"test_fixtures/broken.jpg",
			ImageFile{},
			true,
		},
		{
			"doesnt_exist",
			"test_fixtures/doesnt_exist",
			ImageFile{},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require := require.New(t)
			got, err := ProcessFile(tt.file)
			if tt.wantErr {
				require.Error(err)
			} else {
				require.NoError(err)
			}
			require.EqualValues(tt.want, got)
		})
	}
}
