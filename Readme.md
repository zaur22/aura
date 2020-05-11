RPC API сервис для работы с инкрементом. В качестве реализации RPC
используется GRPC, потому что это стандарт де-факто, да и я сам просто
с ним знаком.

## Краткий обзор функционала
Первый метод API -- GetNumber, он возвращает число.
Сначала это число -- ноль. Второй метод API -- IncrementNumber,
он инкрементирует это число так, что при следующем вызове GetNumber
вернет на единицу больше. Третий метод SetSettings -- метод настроек.

Настройки cледующие:

* Размер инкремента (чтобы увеличивалось не на один, а на два, три или тысячу).
* Размер инкремента должен быть положительным.
* Размер верхней границы. Если я поставлю ее 1000, то при доинкременчивании до 1000 число сбрасывается на ноль само.

Подробнее [в proto файле](https://github.com/zaur22/aura/blob/master/api/grpc/aura.proto).

### Запуск и тесты

Команда запуска приложения:

    docker-compose up
    
Команда для запуска тестов:
    
    make tests

### Список настроек приложения и их дефолтные значения
 Все конфиги задаются через окружение. Подробный список
 [смотреть тут](https://github.com/zaur22/aura/blob/master/pkg/config/configs.go).