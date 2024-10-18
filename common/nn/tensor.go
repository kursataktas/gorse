package nn

import (
	"fmt"
	"github.com/chewxy/math32"
	"math/rand"
	"strings"
)

type Tensor struct {
	data  []float32
	shape []int
}

func LinSpace(start, end float32, shape ...int) *Tensor {
	n := 1
	for _, s := range shape {
		n *= s
	}
	data := make([]float32, n)
	delta := (end - start) / float32(n-1)
	for i := range data {
		data[i] = start + delta*float32(i)
	}
	return &Tensor{
		data:  data,
		shape: shape,
	}
}

func RandN(shape ...int) *Tensor {
	n := 1
	for _, s := range shape {
		n *= s
	}
	data := make([]float32, n)
	for i := range data {
		data[i] = rand.Float32()
	}
	return &Tensor{
		data:  data,
		shape: shape,
	}
}

func (t *Tensor) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	if len(t.data) <= 10 {
		for i := 0; i < len(t.data); i++ {
			builder.WriteString(fmt.Sprint(t.data[i]))
			if i != len(t.data)-1 {
				builder.WriteString(", ")
			}
		}
	} else {
		for i := 0; i < 5; i++ {
			builder.WriteString(fmt.Sprint(t.data[i]))
			builder.WriteString(", ")
		}
		builder.WriteString("..., ")
		for i := len(t.data) - 5; i < len(t.data); i++ {
			builder.WriteString(fmt.Sprint(t.data[i]))
			if i != len(t.data)-1 {
				builder.WriteString(", ")
			}
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (t *Tensor) clone() *Tensor {
	newData := make([]float32, len(t.data))
	copy(newData, t.data)
	return &Tensor{
		data:  newData,
		shape: t.shape,
	}
}

func (t *Tensor) add(other *Tensor) *Tensor {
	if len(t.data) != len(other.data) {
		panic("tensors must have the same size")
	}
	for i := range t.data {
		t.data[i] += other.data[i]
	}
	return t
}

func (t *Tensor) mul(other *Tensor) *Tensor {
	if len(t.data) != len(other.data) {
		panic("tensors must have the same size")
	}
	for i := range t.data {
		t.data[i] *= other.data[i]
	}
	return t
}

func (t *Tensor) pow(n float32) *Tensor {
	for i := range t.data {
		t.data[i] = math32.Pow(t.data[i], n)
	}
	return t
}

func (t *Tensor) sin() *Tensor {
	for i := range t.data {
		t.data[i] = math32.Sin(t.data[i])
	}
	return t
}
