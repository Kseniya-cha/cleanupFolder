Небольшая программа для непрерывной периодической очистки выбранной папки.

Реализована проверка наличия параметров в командной строке; если их нет, значения параметров берутся из конфигурационного файла (.yaml).

### Структура конфигурационного файла

```YAML
timeClean: 2s    
filePath:        
logger:
  loglevel:                  #- уровень логирования
  logFileEnable: true        #- писать ли логи в файл
  logStdoutEnable: true      #- писать ли логи в консоль
  logFile:                   #- путь до файла с логами
  rewriteLog: true           #- перезаписывать ли файл для логирования
```

`timeClean` - время, определяющее частоту запуска программы, `filePath` - путь до папки.

Доступные уровни логирования: `DEBUG`, `INFO`, `WARNING`, `ERROR`, `CRITICAL`, `PANIC`, `FATAL`. Если `rewriteLog = true` и файл, указанный в поле `logFile`, существует, файл будет перезаписан, иначе логи продолжат записываться в уже существующий. Если указанный файл не существует, он будет создан. Параметры `logFileEnable` и `logStdoutEnable` определяют, будут ли писаться логи в файл и консоль соответственно (если `true`).

Чтение конфигурационного файла реализовано с помощью `github.com/spf13/viper`, логгер — с библиотекой `go.uber.org/zap`.
