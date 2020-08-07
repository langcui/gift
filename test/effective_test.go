package effective_test

import "log"
import "testing"
import "io"
import "os"
import "time"
import "fmt"
import "runtime"

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


func escape() *int {
    i := 10
    log.Printf("%v %T\n", i, i)
    log.Printf("%v %T\n", &i, &i)

    j := 10
    log.Printf("j %v %T\n", j, j)
    log.Printf("j %v %T\n", &j, &j)
    return &i
}

func TestEscape(t *testing.T) {
    a := escape()
    log.Printf("%v %T\n", a, a)
    log.Printf("%v %T\n", &a, &a)

}

func TestMap(t *testing.T) {
    m := map[string]bool {
        "a":true,
        "b":true,
    }

    c, ok := m["c"]
    log.Println(c)
    log.Println(ok)
}

const (
    a = iota
    b
    c = iota
    d
)

func TestIota(t *testing.T) {
    log.Printf("a:%v, b:%v, c:%v, d:%v.\n", a, b, c, d)
}

func TestChan(t *testing.T) {
    arr := []int{1,2,3}
    for i := range arr {
        log.Println("i:", i)
        j := i
        go func (){log.Println("in go:", j)}()
   //     time.Sleep(time.Second * 1)
        log.Println("i:", i)
    }

    fmt.Println("cpu num:", runtime.NumCPU())
    fmt.Println("cpu GOMAXPROCS:", runtime.GOMAXPROCS(1))
    time.Sleep(time.Second * 2)
}

func F() (r int) {
    defer func() {
        r++
    }()
    return 1
}

func TestClosure(t *testing.T) {
    fmt.Println(F())
}



















