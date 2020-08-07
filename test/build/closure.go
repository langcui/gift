package main

func myFunc(message int) {
    println(message)
}

func main() {
    f := func(message int) {
        println(message)
    }
    f(0x100)
    myFunc(0x200)
}
