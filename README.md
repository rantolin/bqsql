# bqsql

**bqsql** is a command-line tool for easily running SQL queries against Google BigQuery tables. It simplifies the process of querying and managing BigQuery data by allowing you to write SQL queries directly in your terminal.

## Installation

To install **bqsql**, you need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

Once you have Go installed, you can install **bqsql** using `go get`:

```bash
go get github.com/rantolin/bqsql
```

This command will fetch the source code, build the executable, and place it in your Go bin directory. Make sure your Go bin directory is included in your system's PATH so you can run **bqsql** from anywhere in your terminal.

Alternatively, you can clone the repository and build the executable manually:

```bash
git clone https://github.com/rantolin/bqsql.git
cd bqsql
go build
```

This will create an executable named `bqsql` in the current directory. You can move this executable to a directory in your PATH to make it accessible from anywhere in your terminal.

After installation, you can verify that **bqsql** is installed correctly by running:

```bash
bqsql --version
```

This should display the version of **bqsql** installed on your system.

## Usage

### Authentication

Before using **bqsql**, make sure you have authenticated your Google Cloud account and have the necessary permissions to access your BigQuery datasets and tables. You can authenticate by running:

```bash
gcloud auth application-default login
```

This command will open a browser window for you to authenticate.

### Running Queries

To run a SQL query against a BigQuery table, use the following command:

```bash
bqsql --project_id <PROJECT_ID> --dataset <DATASET_ID> --query "<SQL_QUERY>"
```

Replace `<PROJECT_ID>` with your Google Cloud project ID, `<DATASET_ID>` with the BigQuery dataset ID containing the table you want to query, and `<SQL_QUERY>` with your SQL query enclosed in double quotes.

For example:

```bash
bqsql --project_id my-project --dataset my_dataset --query "SELECT * FROM my_table"
```

### Help

To see all available options and commands, you can use the `--help` flag:

```bash
bqsql --help
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

---

Feel free to modify or expand upon this README as needed!
