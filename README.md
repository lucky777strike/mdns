Mdns - простой dns proxy сервер-балансир.
Распределяет dns запросы на несколько серверов. 
Рекомендуется использовать для таких задач,как dns fuzzing.
Имеет простейший веб интерфейс для простмотра статистики.

Настройка configs/config.yml


Запуск с gobuster на стандартном конфиге 
./gobuster dns -r "127.0.0.1:1234" -d "google.com" -w /tmp/wordlist.txt -t 30

![image](https://user-images.githubusercontent.com/119516169/204873673-2f5af56d-9715-4698-8129-8ca9297d7cb8.png)
![image](https://user-images.githubusercontent.com/119516169/204869299-a8470980-1904-43ac-87e3-5041609b2127.png)


