Так как 
приложение разрабатывалось в рамках монолитной 
архитектуры, был создан всего один основной проект, содержащий всю логику 
приложения. Согласно чистой архитектуре, кодовая база бекенда была разбита 
на слои: домен, пользовательские сценарии, хэндлеры и репозитории. Для 
фронтенда были выделены папки для стиля и для js-файлов css и js 
соответственно. Структура приложения показана на рисунке ниже.
backend - папка бекенд части:
-cmd - папка с точкой входа приложения
-internal - папка с логикой сервиса
-pkg - папка с кодом, который можно переиспользовать
-go.mod - файл для сохранения версий сторонних пакетов
frontend - папка фронтенд части:
-css - стили фронтенда
-js - методы JS проекта
-dashboard.html и index.html - страницы проекта
tests - папка с тестами
Dockerfile - файл для запуска приложения
![image](https://github.com/user-attachments/assets/cc0ee304-639e-4297-b69a-9388d1c0a407)

для запуска приложения в корневой папке необходимо прописать docker-compose up
