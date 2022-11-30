# mdns

### Описание
Mdns - простой dns proxy сервер-балансир.
Распределяет dns запросы на несколько серверов. 
Используется в задачах с большим количеством dns трафика(dns fuzzing и т.д), когда 1 сервер не справляется.
Имеет простейший веб интерфейс для простмотра статистики.

### Пример конфиг файла

```
servaddr: ':1234' #ip:port dns proxy
logport: ':8080'  #ip:port web interface
servers:          #ip:port servers
  - '1.1.1.1:53'
  - '8.8.8.8:53'
  - '8.8.4.4:53'

```

### Примеры запуска
Запуск с gobuster на стандартном конфиге
```
go build
./mdns
./gobuster dns -r "127.0.0.1:1234" -d "google.com" -w /tmp/wordlist.txt -t 30
```

### Просмотр логов
Вывод логов реализован простейшим web интерфейсом.
По умолчанию http://127.0.0.1:8080


### Скриншоты/Screenshots
![image](https://user-images.githubusercontent.com/119516169/204873673-2f5af56d-9715-4698-8129-8ca9297d7cb8.png)
![image](https://user-images.githubusercontent.com/119516169/204869299-a8470980-1904-43ac-87e3-5041609b2127.png)


