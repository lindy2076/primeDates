# primeDates
a small console app to play with dates and primality of numbers written in go

## что это
Это небольшое консольное приложение, определяющее **простые даты** и написанное на Go.

**Простая дата** - дата, в которой все числа простые. 

Например, 3 марта 2003 года - простая дата, т.к. 3, 3, 2003 - простые числа. 

### Пакет date

В пакете date есть структура date.Date с 3 полями: день (`uint`), месяц (`uint`) и год (`uint32`). Рекомендуется создавать объект через метод `BuildDate` или `BuildDateFromIso` - так проверяется корректность предоставленных данных.

### Пакет primes

Два метода - один определяет, простое число (`uint32`) или нет, второй возвращает строку вида `*число* is prime`, если число простое, и `*число* is not prime` в противном случае.
