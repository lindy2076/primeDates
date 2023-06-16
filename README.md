# primeDates
a small console app to play with dates and primality of numbers written in go

## что это
Это небольшое консольное приложение, определяющее **простые даты** и написанное на Go.

**Простая дата** - дата, в которой все числа простые. 

Например, 3 марта 2003 года - простая дата, т.к. 3, 3, 2003 - простые числа. 

Тут два пакета: date и primes.

### Пакет date

#### Структуры:

`Date`

поля:
- day `uint`
- month `uint`
- year `uint32`

методы:
- `(*Date).Day() uint` - возвращает поле day
- `(*Date).Month() uint` - возвращает поле month
- `(*Date).Year() uint32` - возвращает поле year
- `(*Date).IsYearLeap() bool` - проверяет, високосный год (Date.year) или нет
- `(*Date).ToIso() string` - переводит дату в iso формат (yyyy-mm-dd)

#### Публичные методы:

- `ValidateDate(day, month uint, year uint32) error` - возвращает ошибку, если переданы невозможные данные (типа 29.02.2001 или 32.01.0)
- `IsYearLeap(year uint32) bool` - проверяет, високосный год (year) или нет 
- `MonthToNumber(month string) (uint, error)` - переводит название месяца в номер месяца (индексирование с 1). Если month не соответствует английскому названию месяца, то возвращает ошибку
- `NumberToMonth(month uint) (string, error)` - переводит номер месяца (индексирование с 1) в английское название месяца. Если номер не в диапазоне от 1 до 12, то возвращает ошибку
- `BuildDate(day, month uint, year uint32) (*Date, error)` - возвращает ссылку на объект Date, если данные прошли валидацию, а в противном случае - ошибку.
- `BuildDateFromIso(dateIso string) (*Date, error)` - возвращает ссылку на объект Date, если строка валидная, в противном случае - ошибку.

Рекомендуется создавать объект Date через конструкторы BuildDate*.

### Пакет primes

Два метода:

- `IsPrime(num uint32) bool` - проверяет, простое число или нет
- `PrimalityString(num uint32) string` - возвращает строку вида `*num* is prime`/`*num* is not prime`. 

## todo
Date.Next()

Date.Previous()

main.go
