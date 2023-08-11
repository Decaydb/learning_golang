/* Для записи строк из цикла в файл есть полезная штука "strings.Builder" из пакета strings
Вот рабочий пример:
*/

package main
import(
  "fmt"
  "strings"
  )
func main() {
  var builder strings.Builder 
for i:=1; <=5; i++{
  builder.WriteString(fmt.Sprinf("Какой-либо текст, который выводится в цикле несколько раз %d\n"))
  }
  result:= builder.String()
println(result)
  }
