package pkg

import (
	"fmt"
	"time"
)

// FormatTime 將時間格式化為標準格式
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// ParseTime 將字符串解析為時間
func ParseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// TruncateString 截斷字符串到指定長度
func TruncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// FormatPrice 格式化價格顯示
func FormatPrice(price float64) string {
	return fmt.Sprintf("$%.2f", price)
}
