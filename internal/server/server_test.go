package server

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	mockclient "gomock-showcase/test/mocks/magic"
)

func newSentientMock(t *testing.T) Client {
	t.Helper()

	ctrl := gomock.NewController(t)
	c := mockclient.NewMockClient(ctrl)

	c.EXPECT().Magic(gomock.Eq("anything1")).Return("", errors.New("things are on fire")).AnyTimes()
	c.EXPECT().Magic(gomock.Eq("anything2")).Return("Queen of Hearts", nil).AnyTimes()
	c.EXPECT().Magic(gomock.Eq("anything3")).Return("King of Spades", nil).AnyTimes()
	c.EXPECT().Magic(gomock.Eq("anything4")).Return("Oops", nil).AnyTimes()
	return c
}

func TestServer_OnFire_Inline(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		magicClient Client
	}
	type args struct {
		something string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "on fire",
			fields: fields{
				magicClient: func() Client {
					c := mockclient.NewMockClient(ctrl)
					c.EXPECT().Magic(gomock.Any()).Return("", errors.New("things are on fire")).Times(1)
					return c
				}(),
			},
			args: args{
				something: "anything1",
			},
			want:    "On fire",
			wantErr: true,
		},
		{
			name: "magic well done",
			fields: fields{
				magicClient: func() Client {
					c := mockclient.NewMockClient(ctrl)
					c.EXPECT().Magic(gomock.Any()).Return("Queen of Hearts", nil).Times(1)
					return c
				}(),
			},
			args: args{
				something: "anything2",
			},
			want:    "How do you know I picked that?!",
			wantErr: false,
		},
		{
			name: "magic failed",
			fields: fields{
				magicClient: func() Client {
					c := mockclient.NewMockClient(ctrl)
					c.EXPECT().Magic(gomock.Any()).Return("King of Spades", nil).Times(1)
					return c
				}(),
			},
			args: args{
				something: "anything3",
			},
			want:    "No, you're wrong!",
			wantErr: false,
		},
		{
			name: "nothing's work but everything is fine",
			fields: fields{
				magicClient: func() Client {
					c := mockclient.NewMockClient(ctrl)
					c.EXPECT().Magic(gomock.Any()).Return("Oops", nil).Times(1)
					return c
				}(),
			},
			args: args{
				something: "anything4",
			},
			want:    "It's on fire but it's fine.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				magicClient: tt.fields.magicClient,
			}
			got, err := s.OnFire(tt.args.something)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnFire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OnFire() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_OnFire_Stub(t *testing.T) {
	stub := newSentientMock(t)
	type fields struct {
		magicClient Client
	}
	type args struct {
		something string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "on fire",
			fields: fields{
				magicClient: stub,
			},
			args: args{
				something: "anything1",
			},
			want:    "On fire",
			wantErr: true,
		},
		{
			name: "magic well done",
			fields: fields{
				magicClient: stub,
			},
			args: args{
				something: "anything2",
			},
			want:    "How do you know I picked that?!",
			wantErr: false,
		},
		{
			name: "magic failed",
			fields: fields{
				magicClient: stub,
			},
			args: args{
				something: "anything3",
			},
			want:    "No, you're wrong!",
			wantErr: false,
		},
		{
			name: "nothing's work but everything is fine",
			fields: fields{
				magicClient: stub,
			},
			args: args{
				something: "anything4",
			},
			want:    "It's on fire but it's fine.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				magicClient: tt.fields.magicClient,
			}
			got, err := s.OnFire(tt.args.something)
			if (err != nil) != tt.wantErr {
				t.Errorf("OnFire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OnFire() got = %v, want %v", got, tt.want)
			}
		})
	}
}
