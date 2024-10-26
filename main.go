package main

import (
        "bufio"
        "fmt"
        "math/rand"
        "os"
        "time"
        "sync"
)

type BankClient struct {
        balance int
        mu      sync.Mutex
}

func (bc *BankClient) Deposit(amount int) {
        bc.mu.Lock()
        defer bc.mu.Unlock()
        bc.balance += amount
}

func (bc *BankClient) Withdrawal(amount int) error {
        bc.mu.Lock()
        defer bc.mu.Unlock()
        if bc.balance < amount {
                return fmt.Errorf("недостаточно средств")
        }
        bc.balance -= amount
        return nil
}

func (bc *BankClient) Balance() int {
        bc.mu.Lock()
        defer bc.mu.Unlock()
        return bc.balance
}

func main() {
        client := &BankClient{}

        // Запуск горутин для автоматических операций
        go func() {
                for {
                        amount := rand.Intn(10) + 1
                        time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
                        client.Deposit(amount)
                        //fmt.Printf("Автоматический депозит: +%d\n", amount)
                }
        }()

        go func() {
                for {
                        amount := rand.Intn(5) + 1
                        time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
                        err := client.Withdrawal(amount)
                        if err != nil {
                                fmt.Println(err)
                        } else {
                                //fmt.Printf("Автоматическое снятие: -%d\n", amount)
                        }
                }
        }()

        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
                command := scanner.Text()
                switch command {
                case "balance":
                        fmt.Println("Ваш баланс:", client.Balance())
                case "deposit":
                        fmt.Print("Введите сумму для пополнения: ")
                        var amount int
                        fmt.Scanln(&amount)
                        client.Deposit(amount)
                        fmt.Println("Счет пополнен.")
                case "withdrawal":
                        fmt.Print("Введите сумму для снятия: ")
                        var amount int
                        fmt.Scanln(&amount)
                        err := client.Withdrawal(amount)
                        if err != nil {
                                fmt.Println(err)
                        } else {
                                fmt.Println("Сумма снята.")
                        }
                case "exit":
                        fmt.Println("Выход...")
                        return
                default:
                        fmt.Println("Неизвестная команда. Доступные команды: balance, deposit, withdrawal, exit")
                }
        }







        fmt.Println("First commit")   
}