# Task Tracker CLI

Task Tracker CLI — это простое приложение командной строки на языке Go для управления задачами. Оно позволяет добавлять, обновлять, удалять задачи, а также отслеживать их статус.

### Функционал

- Добавление задачи
- Обновление описания задачи
- Удаление задачи
- Отметка задачи как выполняемой
- Отметка задачи как выполненной
- Вывод списка всех задач
- Вывод списка задач по статусу (выполненные, невыполненные, в процессе выполнения)
- Хранение задач в JSON-файле

### Установка

1. Убедитесь, что у вас установлен Go (>=1.18).

2. Клонируйте репозиторий:
   ```sh
   git clone https://github.com/che1nov/task-tracker-cli.git
   cd task-tracker-cli-cli
   ```

3. Скомпилируйте и установите:
   ```sh
   go build -o task-cli
   ```

### Использование

#### Добавление новой задачи

```sh
task-cli add "Купить продукты"
# Вывод: Задача успешно добавлена (ID: 1)
```

#### Обновление задачи

```sh
task-cli update 1 "Купить продукты и приготовить ужин"
```

#### Удаление задачи

```sh
task-cli delete 1
```

#### Изменение статуса задачи

```sh
task-cli mark-in-progress 1
```

```sh
task-cli mark-done 1
```

#### Вывод списка задач

```sh
task-cli list
```

#### Вывод списка задач по статусу

```sh
task-cli list done
```

```sh
task-cli list todo
```

```sh
task-cli list in-progress
```

### Структура задачи

Каждая задача представлена в JSON-файле с полями:

```json
{
  "id": 1,
  "description": "Купить продукты",
  "status": "todo",
  "createdAt": "2025-03-16T12:00:00Z",
  "updatedAt": "2025-03-16T12:00:00Z"
}
```

### Хранение данных

Файл `tasks.json` используется для хранения списка задач. Если файл отсутствует, он создается автоматически.

### Ошибки и обработка исключений

- Проверка наличия задачи перед обновлением или удалением
- Валидация ввода пользователя
- Обработка ошибок при работе с файлом JSON

### Лицензия

Этот проект распространяется под лицензией MIT.

````markdown name=LICENSE
MIT License

Copyright (c) 2025 Ilya Chernov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.