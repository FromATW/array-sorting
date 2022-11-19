package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func sort(a []int, l int, r int) {
	for i := l; i < r; i++ {
		for j := i + 1; j <= r; j++ {
			if a[i] > a[j] {
				swap(&a[i], &a[j])
			}
		}
	}
}

func main() {
	var n, l, r int
	a := make([]int, 1000)
	for {
		_, err := fmt.Scanf("%d %d %d\n", &n, &l, &r)
		if err != nil {
			fmt.Println("格式错误，请重新输入")
		}
		if l > r {
			fmt.Println("格式错误，请重新输入")
		} else {
			break
		}
	}
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &a[i])
		if err != nil {
			fmt.Println("格式错误，请重新输入")
		}
	}
	sort(a, l, r)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
}
