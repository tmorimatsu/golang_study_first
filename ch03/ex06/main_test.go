package main

import (
	"image/color"
	"testing"
)

func TestSuperSampling0(t *testing.T) {

	t.Log("上下左右の平均を返却すること")

	c := createSampleSlice()

	expected := color.RGBA{90, 100, 100, 255}
	actual := superSampling(c, 1, 1)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestSuperSampling1(t *testing.T) {

	t.Log("上下左右の要素がない場合には渡されたそのままの要素を返却すること")

	c := createSampleSlice()

	expected := c[2][1]
	actual := superSampling(c, 2, 1)
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

// テスト用のスライス作成
func createSampleSlice() [][]color.RGBA {
	c := make([][]color.RGBA, 3)
	for i := range c {
		c[i] = make([]color.RGBA, 3)
	}
	for i := range c {
		for j := range c[i] {
			c[i][j] = color.RGBA{uint8((i + 1) * (j + 1)), uint8((i + 1) * (j + 1)), uint8((i + 1) * (j + 1)), 255}
		}
	}

	c[0][1] = color.RGBA{200, 100, 20, 255}
	c[1][0] = color.RGBA{40, 100, 240, 255}
	c[2][1] = color.RGBA{80, 100, 60, 255}
	c[1][2] = color.RGBA{40, 100, 80, 255}

	return c
}
