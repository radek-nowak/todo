package ui

import (
	"strings"
	"testing"
)

type TestColumn struct {
	name   string
	length int
}

func (c TestColumn) Name() string {
	return c.name
}
func (c TestColumn) MaxLen() int {
	return c.length
}

func Test_headerTop(t *testing.T) {
	type args struct {
		columns []Column
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "single column",
			args: args{
				columns: []Column{
					TestColumn{
						name:   "id",
						length: 2,
					},
				},
			},
			want: topLeft + topRight,
		},
		{
			name: "multiple columns",
			args: args{
				columns: []Column{
					TestColumn{
						name:   "col1",
						length: 2,
					},
					TestColumn{
						name:   "col2",
						length: 1,
					},
				},
			},
			want: topLeft + strings.Repeat(horizontal, 4) + topMiddle + strings.Repeat(horizontal, 3) + topRight,
		},
		{
			name: "no columns",
			args: args{
				columns: []Column{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// too fancy?
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("headerTop() panicked unexpectedly: %v", r)
					}
				} else {
					if tt.wantErr {
						t.Errorf("headerTop() did not panic, but expected to panic")
					}
				}
			}()

			if got := headerTop(tt.args.columns); got != tt.want {
				t.Errorf("headerTop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerMiddle(t *testing.T) {
	type args struct {
		columns []Column
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "single column",
			args: args{
				columns: []Column{
					TestColumn{
						name:   "ID",
						length: 2,
					},
				},
			},
			want: vertical + " ID " + vertical,
		},
		{
			name: "multiple columns",
			args: args{
				columns: []Column{
					TestColumn{
						name:   "ID",
						length: 2,
					},
					TestColumn{
						name:   "Task",
						length: 4,
					},
					TestColumn{
						name:   "Done",
						length: 4,
					},
				},
			},
			want: vertical + " ID " + vertical + " Task " + vertical + " Done " + vertical,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := headerMiddle(tt.args.columns); got != tt.want {
				t.Errorf("headerMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_headerBottom(t *testing.T) {
	type args struct {
		columns []Column
	}
	tests := []struct {
		name string
		args args
		want string
		wantErr bool
	}{
		{
			name: "single column",
			args: args{
				columns: []Column{
					TestColumn{
						name:   "id",
						length: 2,
					},
				},
			},
			want: bottomLeft + bottomRight,
		},
		{
			name: "multiple columns",
			args: args{
				columns: []Column{
					TestColumn{
						name:   "col1",
						length: 2,
					},
					TestColumn{
						name:   "col2",
						length: 1,
					},
				},
			},
			want: bottomLeft + strings.Repeat(horizontal, 4) + bottomMiddle + strings.Repeat(horizontal, 3) + bottomRight,
		},
		{
			name: "no columns",
			args: args{
				columns: []Column{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// too fancy?
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("headerTop() panicked unexpectedly: %v", r)
					}
				} else {
					if tt.wantErr {
						t.Errorf("headerTop() did not panic, but expected to panic")
					}
				}
			}()

			if got := headerBottom(tt.args.columns); got != tt.want {
				t.Errorf("headerTop() = %v, want %v", got, tt.want)
			}
		})
	}
}
