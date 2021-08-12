package store

import (
	"testing"

	"github.com/golang/mock/gomock"

	mockclient "gomock-showcase/test/mocks/store/source"
)

func newSentientMock(t *testing.T) Client {
	t.Helper()

	ctrl := gomock.NewController(t)
	c := mockclient.NewMockClient(ctrl)

	c.EXPECT().DarkMagic(gomock.Eq("Harry Porter")).Return("Owl").AnyTimes()
	c.EXPECT().DarkMagic(gomock.Eq("Ron Weasley")).Return("Mouse").AnyTimes()
	c.EXPECT().DarkMagic(gomock.Eq("Voldemort")).Return("Bad guy").AnyTimes()
	return c
}

func TestStore_Verify(t *testing.T) {
	sc := newSentientMock(t)
	type fields struct {
		c Client
	}
	type args struct {
		something string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "verify harry porter",
			fields: fields{
				c: sc,
			},
			args: args{
				something: "Harry Porter",
			},
			want: "Hello Mr. Porter. Welcome!",
		},
		{
			name: "verify ron weasley",
			fields: fields{
				c: sc,
			},
			args: args{
				something: "Ron Weasley",
			},
			want: "You're in the wrong place.",
		},
		{
			name: "verify voldemort",
			fields: fields{
				c: sc,
			},
			args: args{
				something: "Voldemort",
			},
			want: "Unknown creature. Avada Kedavra!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{
				c: tt.fields.c,
			}
			if got := s.Verify(tt.args.something); got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
