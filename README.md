Небольшая программа для непрерывной периодической очистки выбранной папки.

Реализована проверка наличия параметров в командной строке; если их нет, значения параметров берутся из конфигурационного файла (.yaml).

### Структура конфигурационного файла

```YAML
timeClean: 2s
filePath:
```

где `timeClean` - время, определяющее частоту запуска программы, `filePath` - путь до папки.
