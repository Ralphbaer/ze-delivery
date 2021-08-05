package zmock

import (
	"testing"

	"github.com/Ralphbaer/ze-delivery/common/zpointers"
)

type s1 struct {
	ID string
}

type sP1 struct {
	ID *string
}

func TestFieldValue_Matches(t *testing.T) {
	type fields struct {
		FieldName string
		Value     interface{}
	}
	type args struct {
		x interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "eq1",
			fields: fields{
				FieldName: "ID",
				Value:     "1",
			},
			args: args{
				x: s1{
					ID: "1",
				},
			},
			want: true,
		},
		{
			name: "eq1",
			fields: fields{
				FieldName: "ID",
				Value:     zpointers.String("1"),
			},
			args: args{
				x: sP1{
					ID: zpointers.String("1"),
				},
			},
			want: true,
		},
		{
			name: "neq1",
			fields: fields{
				FieldName: "ID",
				Value:     "1",
			},
			args: args{
				x: sP1{
					ID: zpointers.String("1"),
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := FieldValue{
				FieldName: tt.fields.FieldName,
				Value:     tt.fields.Value,
			}
			if got := s.Matches(tt.args.x); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}