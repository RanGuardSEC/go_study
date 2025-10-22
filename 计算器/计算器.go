package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func is(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func calculate(num1, num2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("除数不能为 0")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("不支持的运算符 '%s'", op)
	}
}

func baocuo(input string) (float64, float64, string, error) {
	var num1, num2 float64
	var c rune
	count, err := fmt.Sscanf(input, "%f %c %f", &num1, &c, &num2)
	if err != nil || count != 3 {
		return 0, 0, "", fmt.Errorf("输入格式错误,请使用：数字 运算符 数字，例如：3 + 5")
	}
	op := string(c)
	if !is(op) {
		return 0, 0, "", fmt.Errorf("不支持的运算符 '%s'，请使用 +, -, *, /", op)
	}

	return num1, num2, op, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("请输入表达式 (输入 exit 退出): \n")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "exit" {
			fmt.Println("程序已退出")
			return
		}

		num1, num2, op, err := baocuo(input)

		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculate(num1, num2, op)
		if err != nil {
			fmt.Printf("计算错误: %v\n", err)
			continue
		}
		//格式化输出
		if result == float64(int(result)) && num1 == float64(int(num1)) && num2 == float64(int(num2)) {
			//三者都是整数，显示整数格式
			fmt.Printf("计算结果:%d %s %d = %d\n", int(num1), op, int(num2), int(result))
		} else {
			//否则保留两位小数
			fmt.Printf("计算结果:%.2f %s %2f = %.2f\n", num1, op, num2, result)
		}
		fmt.Println()
	}
}
