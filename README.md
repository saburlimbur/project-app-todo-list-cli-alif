# To-Do List CLI (Golang)

A simple Command Line Interface (CLI) application for managing a To-Do List using Golang and Cobra.  
Data disimpan dalam file JSON: `data/todos.json`.

---

## Menjalankan Aplikasi

Jalankan aplikasi utama:

```sh
go run .
```

### Command Todo Features

- [x] Add Todo

  ```sh
  go run . add --title="<judul>" --desc="<deskripsi>" --status=<status> --priority=<priority>
  ```

- [x] List Todo

  ```sh
  go run . lists
  ```

- [x] Update Todo Status

  ```sh
  go run .
  ```

- [x] Delete Todo

  ```sh
  go run .
  ```
