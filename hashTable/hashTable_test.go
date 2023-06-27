package hashTable

import (
	"reflect"
	"testing"
)

func TestHashTable_Exists(t1 *testing.T) {
	type fields struct {
		table     []*entry
		capacity  int
		threshold int
		size      int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test exists",
			fields: fields{
				table:     make([]*entry, 16),
				capacity:  16,
				threshold: 12,
				size:      0,
			},
			args: args{key: "test"},
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HashTable{
				table:     tt.fields.table,
				capacity:  tt.fields.capacity,
				threshold: tt.fields.threshold,
				size:      tt.fields.size,
			}
			if got := t.Exists(tt.args.key); got != tt.want {
				t1.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Get(t1 *testing.T) {
	type fields struct {
		table     []*entry
		capacity  int
		threshold int
		size      int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue interface{}
		wantExist bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HashTable{
				table:     tt.fields.table,
				capacity:  tt.fields.capacity,
				threshold: tt.fields.threshold,
				size:      tt.fields.size,
			}
			gotValue, gotExist := t.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t1.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotExist != tt.wantExist {
				t1.Errorf("Get() gotExist = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func TestHashTable_Put(t1 *testing.T) {
	type fields struct {
		table     []*entry
		capacity  int
		threshold int
		size      int
	}
	type args struct {
		key   interface{}
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HashTable{
				table:     tt.fields.table,
				capacity:  tt.fields.capacity,
				threshold: tt.fields.threshold,
				size:      tt.fields.size,
			}
			t.Put(tt.args.key, tt.args.value)
		})
	}
}

func TestHashTable_Remove(t1 *testing.T) {
	type fields struct {
		table     []*entry
		capacity  int
		threshold int
		size      int
	}
	type args struct {
		key interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HashTable{
				table:     tt.fields.table,
				capacity:  tt.fields.capacity,
				threshold: tt.fields.threshold,
				size:      tt.fields.size,
			}
			t.Remove(tt.args.key)
		})
	}
}

func TestHashTable_hash(t1 *testing.T) {
	type fields struct {
		table     []*entry
		capacity  int
		threshold int
		size      int
	}
	type args struct {
		key  interface{}
		size int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HashTable{
				table:     tt.fields.table,
				capacity:  tt.fields.capacity,
				threshold: tt.fields.threshold,
				size:      tt.fields.size,
			}
			if got := t.hash(tt.args.key, tt.args.size); got != tt.want {
				t1.Errorf("hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_resize(t1 *testing.T) {
	type fields struct {
		table     []*entry
		capacity  int
		threshold int
		size      int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &HashTable{
				table:     tt.fields.table,
				capacity:  tt.fields.capacity,
				threshold: tt.fields.threshold,
				size:      tt.fields.size,
			}
			t.resize()
		})
	}
}

func TestNewHashTable(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want *HashTable
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHashTable(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHashTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
