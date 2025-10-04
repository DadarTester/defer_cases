//автозакрытие файла после чтения данных
func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Ошибка при открытии файла:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка чтения файла:", err)
    }
}

// освобождение ресурсов при работе с получением данных от сервера
func fetchData(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("err: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("invalid code status: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("err reading body: %v", err)
    }

    fmt.Printf("Got %d bite\n", len(body))
    fmt.Printf("Body: %s\n", string(body))
    
    return nil
}

func main() {
    err := fetchData("https://httpbin.org/json")
    if err != nil {
        log.Println("Ошибка:", err)
    }
}

// обработка паники с помощь функции recover
func safeDivide(a, b int) int {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("panic recover:", r)
        }
    }()

    return a / b
}

func main() {
    fmt.Println("result:", safeDivide(10, 0))
}

// обработка времени работы функции
func trackTime(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s took %v\n", name, time.Since(start))
    }
}

func slowOperation() {
    defer trackTime("slowOperation")()
    time.Sleep(100 * time.Millisecond)
}

// восстановление после паники
func calculateSum(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panic recover:", r)
            err = fmt.Errorf("panic occurred")
        }
    }()

    result = a + b
    if result > 100 {
        panic("Sum is too big")
    }

    return result, nil
}

func main() {
    sum, err := calculateSum(50, 60)
    if err != nil {
        fmt.Println("err:", err)
    } else {
        fmt.Println("Sum:", sum)
    }
}

// откат транзакции
func executeTransaction(db *sql.DB) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if err != nil {
            tx.Rollback()
            fmt.Println("transaction rolled back")
        } else {
            tx.Commit()
            fmt.Println("transaction ended")
        }
    }()

    _, err = tx.Exec("INSERT INTO users(name) VALUES('John')")
    if err != nil {
        return err
    }

    _, err = tx.Exec("INSERT INTO accounts(user_id) VALUES(LAST_INSERT_ID())")
    return err
}

func main() {
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        fmt.Println("err:", err)
        return
    }
    defer db.Close()

    err = executeTransaction(db)
    if err != nil {
        fmt.Println("err:", err)
    }
}

// изменение входящего значения, например, для применения скидок
func doubleReturn() (result int) {
    defer func() {
        result *= 2 // Меняем именованный возвращаемый параметр
    }()
    
    return 5 // Фактически вернет 10
}
