package effective_test

import "log"
import "testing"
import "io"
import "os"

func nextInt(pos int) (nextPos int) {
    nextPos = 10
    return
    return pos + 1
}

func TestNamedFunc(t *testing.T) {
    r := nextInt(1)
    log.Println(r)
}

// 内容返回文件的内容作为字符串。
func TestReadContents(t *testing.T) {
    filename := "./a.txt"
    f, err := os.Open(filename)
    if err != nil {
        panic(err)
        return
    }
    defer f.Close()  // 我们结束后就关闭了f

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        log.Println(n, err)
        result = append(result, buf[0:n]...) // append稍后讨论。
        if err != nil {
            if err == io.EOF {
                break
            }
            return // 如果我们回到这里，f就关闭了。
        }
    }
    log.Println(len(result), cap(result))
    return // 如果我们回到这里，f就关闭了。
}
