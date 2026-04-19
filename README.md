# bqsql

[![Go Version](https://img.shields.io/github/go-mod/go-version/rantolin/bqsql)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**bqsql** is a lightweight, profile-driven CLI tool designed to make Google BigQuery interactions seamless. Whether you're running ad-hoc SQL, exploring schemas, or managing datasets, `bqsql` provides a fast, terminal-centric experience.

---

## 🚀 Key Features

- **Profile Management**: Switch between projects and datasets easily using YAML-based profiles.
- **SQL at your Fingertips**: Execute complex queries directly from the terminal with formatted table outputs.
- **Resource Exploration**: Quickly list datasets, tables, and routines without the overhead of the Cloud Console.
- **Data Preview**: Use the `head` command to instantly peek at the first 10 rows of any table.
- **Developer Friendly**: Built with Cobra/Viper, making it highly extensible and script-friendly.

---

## 📦 Installation

### Using Go
```bash
go install github.com/rantolin/bqsql@latest
```

### From Source
```bash
git clone https://github.com/rantolin/bqsql.git
cd bqsql
go build -o bqsql
```

---

## 🛠 Setup & Configuration

### 1. Authentication
`bqsql` uses Google's Application Default Credentials (ADC). Authenticate your environment first:
```bash
gcloud auth application-default login
```

### 2. Configure Profiles
Set up your default project and dataset to avoid passing flags every time:
```bash
bqsql set --project_id your-project-id --dataset your_default_dataset --profile default
```
This creates a `bqsql.yaml` file in your home directory.

---

## 📖 Usage Examples

### Running a Query
```bash
bqsql query "SELECT name, count(*) FROM `my_project.my_dataset.users` GROUP BY 1 LIMIT 5"
```

### Previewing a Table
```bash
bqsql head users
```

### Listing Resources
```bash
bqsql list datasets
bqsql --dataset my_dataset list tables
```

### Describing Metadata
```bash
bqsql describe table users --schema --num_rows
```

---

## 🗺 Project Roadmap

We are currently transforming `bqsql` into a comprehensive BigQuery management suite.
- [ ] **Phase 1**: Streaming & Result Pagination (P0)
- [ ] **Phase 2**: JSON/CSV Output Formats (P0)
- [ ] **Phase 3**: Resource CRUD & IAM Management (P1)
- [ ] **Phase 4**: Interactive Shell (the "SnowSQL" experience) (P1)

Check out the full [ROADMAP.md](./ROADMAP.md) for more details.

---

## 🤝 Contributing

We welcome contributions! 
- **Developers**: See [DEVELOPMENT.md](./DEVELOPMENT.md) for our architectural guide and tutorial.
- **Instructions**: For AI agents or advanced context, see [GEMINI.md](./GEMINI.md).

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
