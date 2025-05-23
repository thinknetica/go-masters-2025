В Go мьютексы (mutexes) используются для **безопасного доступа к общим данным** из нескольких горутин, предотвращая **гонки данных** (data races) и обеспечивая **консистентность состояния**. Они являются частью пакета `sync`.

---

## 🔹 **Когда нужны мьютексы?**
Мьютексы применяются, когда:
1. **Несколько горутин читают и пишут в одну переменную** (например, счётчик, кэш, структуру данных).
2. **Требуется атомарное изменение состояния** (например, банковский перевод).
3. **Каналы неудобны или избыточны** (например, для защиты редких записей в мапу).

---

## 🔹 **Типы мьютексов в Go**
### 1. **`sync.Mutex` (обычный мьютекс)**
   - Блокирует доступ для всех горутин, кроме одной.
   - Пример:
     ```go
     package main

     import (
         "fmt"
         "sync"
     )

     type SafeCounter struct {
         mu    sync.Mutex
         value int
     }

     func (c *SafeCounter) Increment() {
         c.mu.Lock()         // Блокируем доступ
         defer c.mu.Unlock() // Разблокируем при выходе
         c.value++
     }

     func main() {
         counter := SafeCounter{}
         var wg sync.WaitGroup

         for i := 0; i < 1000; i++ {
             wg.Add(1)
             go func() {
                 counter.Increment()
                 wg.Done()
             }()
         }

         wg.Wait()
         fmt.Println("Итоговое значение:", counter.value) // 1000
     }
     ```

### 2. **`sync.RWMutex` (читающий-пишущий мьютекс)**
   - Позволяет **множественное чтение** но **эксклюзивную запись**.
   - Эффективен, когда чтений больше, чем записей.
   - Пример:
     ```go
     type Cache struct {
         mu    sync.RWMutex
         items map[string]string
     }

     func (c *Cache) Get(key string) string {
         c.mu.RLock()         // Блокировка для чтения
         defer c.mu.RUnlock() // Разблокировка
         return c.items[key]
     }

     func (c *Cache) Set(key, value string) {
         c.mu.Lock()          // Эксклюзивная блокировка
         defer c.mu.Unlock()
         c.items[key] = value
     }
     ```

---

## 🔹 **Почему не всегда каналы?**
Каналы — это идиоматичный способ общения горутин, но:
- **Мьютексы проще** для защиты конкретной переменной.
- **Каналы избыточны** для простых операций (например, инкремент счётчика).
- **Мьютексы могут быть быстрее** в высоконагруженных сценариях.

> **Правило:** *"Используйте каналы для координации, мьютексы для защиты состояния"* (Rob Pike).

---

## 🔹 **Опасности мьютексов**
1. **Взаимная блокировка (deadlock)**:
   ```go
   var mu sync.Mutex
   mu.Lock()
   mu.Lock() // Программа зависнет
   ```
2. **Забытая разблокировка**:
   ```go
   mu.Lock()
   if err != nil {
       return // Забыли Unlock!
   }
   mu.Unlock()
   ```
   **Решение:** Всегда используйте `defer mu.Unlock()`.

---

## 🔹 **Когда мьютексы не нужны?**
1. **Данные только читаются** (без изменений).
2. **Каждая горутина работает со своей копией данных**.
3. **Атомарные операции** (например, `atomic.AddInt64`).

---

## 🔹 **Вывод**
- **Мьютексы** в Go защищают общие ресурсы от гонок.
- **`sync.Mutex`** — для эксклюзивного доступа.
- **`sync.RWMutex`** — для оптимизации частых чтений.
- **Альтернативы:** Каналы, `atomic`, иммутабельные структуры.

Используйте мьютексы аккуратно, чтобы избежать deadlock и сохранить производительность! 🔒