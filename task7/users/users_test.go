package users

import (
	"testing"
)

func TestCustomer_Age(t *testing.T) {
	tests := []struct {
		name string
		c    Customer
		want int
	}{
		{
			name: "normal",
			c:    Customer{age: 10},
			want: 10,
		},
		{
			name: "zero",
			c:    Customer{age: 0},
			want: 0,
		},
		{
			name: "negative",
			c:    Customer{age: -10},
			want: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Age(); got != tt.want {
				t.Errorf("Customer.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_Age(t *testing.T) {
	tests := []struct {
		name string
		e    Employee
		want int
	}{
		{
			name: "normal",
			e:    Employee{age: 10},
			want: 10,
		},
		{
			name: "zero",
			e:    Employee{age: 0},
			want: 0,
		},
		{
			name: "negative",
			e:    Employee{age: -10},
			want: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Age(); got != tt.want {
				t.Errorf("Employee.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name  string
		users []Age
		want  int
	}{
		{
			name: "normal",
			users: []Age{
				&Customer{age: 10},
				&Employee{age: 20},
				&Customer{age: 30},
				&Employee{age: 40},
			},
			want: 40,
		},
		{
			name: "one",
			users: []Age{
				&Customer{age: 10},
			},
			want: 10,
		},
		{
			name: "same",
			users: []Age{
				&Customer{age: 10},
				&Customer{age: 10},
				&Customer{age: 10},
			},
			want: 10,
		},
		{
			name: "reversed",
			users: []Age{
				&Customer{age: 40},
				&Customer{age: 30},
				&Customer{age: 20},
				&Customer{age: 10},
			},
			want: 40,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.users); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
