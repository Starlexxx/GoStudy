package methodless_users

import (
	"reflect"
	"testing"
)

func TestOldest(t *testing.T) {
	tests := []struct {
		name  string
		users []interface{}
		want  interface{}
	}{
		{
			name: "normal",
			users: []interface{}{
				Customer{age: 10},
				Employee{age: 20},
				Customer{age: 30},
				Employee{age: 40},
			},
			want: Employee{age: 40},
		},
		{
			name: "one",
			users: []interface{}{
				Customer{age: 10},
			},
			want: Customer{age: 10},
		},
		{
			name: "same",
			users: []interface{}{
				Customer{age: 10},
				Customer{age: 10},
				Customer{age: 10},
			},
			want: Customer{age: 10},
		},
		{
			name: "reversed",
			users: []interface{}{
				Customer{age: 40},
				Customer{age: 30},
				Customer{age: 20},
				Customer{age: 10},
			},
			want: Customer{age: 40},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Oldest(tt.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Oldest() = %v, want %v", got, tt.want)
			}
		})
	}
}
