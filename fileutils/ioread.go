package fileutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/trungtvq/go-utils/service/cs"
)

func ReadNum(message string) int {
	var num int
	cs.Notice(message)
	fmt.Scanf("%d", &num)
	return num
}
func ReadString(message string) string {
	cs.Notice(message)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)
	return str
}
